package datamodels

import (
	"SecondKill/config"
	"encoding/json"
)

//简单的消息体
type Message struct {
	ProductID int64
	UserID    int64
}

func NewMessage(productID int64, userID int64) *Message {
	return &Message{ProductID: productID, UserID: userID}
}

func (m *Message) JsonToStr() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		config.AppSetting.Logger.Error("json 解析出错：" + err.Error())
	}

	return string(bytes)
}

func (m *Message) StrToJson(dataStr []byte) *Message {
	if err := json.Unmarshal(dataStr, m); err != nil {
		config.AppSetting.Logger.Error("json 转换出错：" + err.Error())
	}
	return m
}