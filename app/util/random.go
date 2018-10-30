package util

import (
	"math/rand"
	"strconv"
	"time"
	"math"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

func init() {
	rand.NewSource(time.Now().Unix())
}

//保留小数位
func Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
}

//Krand 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//随机数
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

//随机生成加减.
func RandomIncOrDec() string {
	list := []string{
		"+",
		"-",
	}
	return string(list[RandInt(0, len(list))])
}

//处理精度
func PrecisionHandle(nr float64, n int) (result float64, err error) {
	sf := strconv.FormatFloat(nr, 'f', n, 64)
	result, err = strconv.ParseFloat(sf, 64)
	if err != nil {
		return
	}

	return
}

//随机生成 float64的浮点数
func RandomFloat64(numberRange ...float64) float64 {
	nr := 1.0
	nr = rand.Float64()*(float64(numberRange[1])-float64(numberRange[0])) + float64(numberRange[0])

	if len(numberRange) > 2 {
		nr, _ = PrecisionHandle(nr, int(numberRange[2]))
	}

	return nr
}
