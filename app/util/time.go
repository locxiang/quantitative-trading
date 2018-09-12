package util

import (
	"time"
)

//毫秒转换时间类型
func Millisecond2Time(timeUtc int64) time.Time {
	return time.Unix(timeUtc/1000, timeUtc%1000*1e6)
}

//输出毫秒值
func Time2Millisecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}


type JsonTimeMillisecond time.Time

//实现它的json序列化方法
func (t JsonTimeMillisecond) MarshalJSON() int64 {
	var stamp = Time2Millisecond(time.Time(t))
	return stamp
}
