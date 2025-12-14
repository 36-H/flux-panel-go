package model

type BaseRes struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	TimeStamp int64  `json:"ts"`
	Data      any    `json:"data"`
}
