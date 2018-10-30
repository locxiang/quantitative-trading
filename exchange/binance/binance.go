package binance

import (
	"time"
	"github.com/locxiang/quantitative-trading/app/pkg/setting"
	"github.com/gorilla/websocket"
	"github.com/lexkong/log"
	"github.com/buger/jsonparser"
	"strings"
	"encoding/json"
	"strconv"
	tradeService "github.com/locxiang/quantitative-trading/app/services/trade"
	"github.com/locxiang/quantitative-trading/app/models"
	"fmt"
	poolService "github.com/locxiang/quantitative-trading/app/services/pool"
)

type Trade struct {
	Type          string `json:"e"`
	Time          int64  `json:"E"`
	Symbol        string `json:"s"`
	TradeId       int64  `json:"t"`
	Price         string `json:"p"`
	Quantity      string `json:"q"`
	BuyerOrderId  int64  `json:"b"`
	SellerOrderId int64  `json:"a"`
	Timestamp     int64  `json:"T"`
	IsMaker       bool   `json:"m"`
	Ignore        bool   `json:"M"`
}

func (t *Trade) insertDB() (s *models.Trade, err error) {
	s = new(models.Trade)
	s.Exchange = "binance"
	s.Symbol = t.Symbol
	s.TradeId = t.TradeId
	s.Price, err = strconv.ParseFloat(t.Price, 64)
	if err != nil {
		return
	}

	s.Quantity, err = strconv.ParseFloat(t.Quantity, 64)
	if err != nil {
		return
	}

	s.IsMaker = t.IsMaker
	s.Ignore = t.Ignore
	s.BuyerOrderId = t.BuyerOrderId
	s.SellerOrderId = t.SellerOrderId
	s.Timestamp = t.Timestamp

	if t.BuyerOrderId < t.SellerOrderId {
		s.Direction = models.OrderSell
	} else {
		s.Direction = models.OrderBuy
	}

	tradeService.Add(s)
	return
}

type Depth struct {
	Type    string     `json:"e"` // Event type
	Time    int64      `json:"E"` // Event time
	Symbol  string     `json:"s"` // Symbol
	FirstId int64      `json:"U"` // First update ID in event
	FinalId int64      `json:"u"` //Final update ID in event
	Bids    [][]string `json:"b"`
	Asks    [][]string `json:"a"`
}

func (d *Depth) OrderBids() (orders []*Order) {
	for _, b := range d.Bids {
		var err error
		o := new(Order)
		o.Price, err = strconv.ParseFloat(b[0], 64)
		o.Quantity, err = strconv.ParseFloat(b[1], 64)
		if err != nil {
			log.Errorf(err, "转换价格出错")
		}
		orders = append(orders, o)
	}

	return
}

type Order struct {
	Price    float64
	Quantity float64
}

const (
	TRADE = "trade"
	DEPTH = "depth"
)

type streamEvents []byte

func (s streamEvents) Type() string {
	stream, _, _, _ := jsonparser.Get(s, "stream")
	str := string(stream)
	arr := strings.Split(str, "@")
	return arr[1]
}

func (s streamEvents) Data() []byte {
	data, _, _, _ := jsonparser.Get(s, "data")
	return data
}

//转换成结构体
func (s streamEvents) Unmarshal() interface{} {
	switch s.Type() {
	case TRADE:
		var r Trade
		json.Unmarshal(s.Data(), &r)
		return r
	case DEPTH:
		var d Depth
		json.Unmarshal(s.Data(), &d)
		return d
	default:
		log.Errorf(nil, "结构体转换失败：%s", s.Type())
	}

	return nil
}

func Init() (done chan struct{}) {

	done = make(chan struct{})
	cfg := setting.Env().Exchange

	url := cfg.Url

	for _, t := range cfg.Spot {
		t = strings.ToLower(t)
		url += fmt.Sprintf("%s@depth/%s@trade/", t, t)
	}

	log.Infof("连接交易所地址：%s", url)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		log.Fatal("dial:", err)
	}

	//循环
	go func() {
		defer c.Close()
		defer close(done)
		for {
			select {
			case <-done:
				return
			case <-time.After(1 * time.Second):
				err := c.WriteMessage(websocket.PingMessage, []byte(""))
				if err != nil {
					log.Infof("write:", err)
					return
				}
			default:
				//TODO 防止堵塞
			}

			_, message, err := c.ReadMessage()
			if err != nil {
				log.Errorf(err, "读取数据出错")
				return
			}
			s := streamEvents(message)

			switch data := s.Unmarshal().(type) {
			case Depth:
			case Trade:

				mt, err := data.insertDB()
				if err != nil {
					log.Errorf(err, "trade 写入数据库失败")
					continue
				}

				if err := mt.InsertTSDB(); err != nil {
					log.Errorf(err, "trade 写入TSDB失败")
				}
				poolService.EventTrade(mt)
			}
		}
	}()

	return
}
