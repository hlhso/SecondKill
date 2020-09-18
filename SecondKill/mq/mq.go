package mq

import (
	"SecondKill/datamodels"
	"SecondKill/services"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

var producer *nsq.Producer
const (
	addPro = "127.0.0.1:4150"
	addCon = "127.0.0.1:4161"
	topic = "topic_demo"
	channel ="first"
)


//初始化及连接
func InitProduct ()(err error) {
	cfg := nsq.NewConfig()
	producer, err = nsq.NewProducer(addPro, cfg)
	if err != nil {
		fmt.Printf("create producer failed, err:%v\n", err)
	}
	return nil
}

//发布消息
func Publish(message string) error {
	err := producer.Publish(topic, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

//消费者类型
type MyHandler struct {
	OrderService   services.IOrderService
	ProductService services.IProductService
}

//消费者处理函数
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	log.Printf("Received a message: %s", msg.Body)
	message := &datamodels.Message{}
	err = json.Unmarshal([]byte(string(msg.Body)), message)
	if err != nil {
		fmt.Println(err)
	}

	//插入订单
	order:=&datamodels.Order{
		Model:       gorm.Model{},
		UserId:      message.UserID,
		ProductId:   message.ProductID,
	}
	_, err = m.OrderService.InsertOrder(order)
	if err != nil {
		fmt.Println(err)
	}

	//扣除商品数量
	err = m.ProductService.SubNumberOne(message)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	return
}

//初始化消费者
func InitConsume(orderService services.IOrderService,productService services.IProductService) error {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic,channel , config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return err
	}
	consumer := &MyHandler{
		OrderService: orderService,
		ProductService: productService,
	}
	c.AddHandler(consumer)

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err := c.ConnectToNSQLookupd(addCon); err != nil { // 通过lookupd查询
		return err
	}
	return nil

}
