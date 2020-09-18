package main

import (
	"SecondKill/config"
	"SecondKill/database"
	"SecondKill/mq"
	"SecondKill/repositories"
	"SecondKill/services"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Setup()
	db := database.GetMysqlInstance()
	//创建product数据库操作实例
	product := repositories.NewProduct(db)
	//创建product serivce
	productService := services.NewProductService(product)
	//创建Order数据库实例
	order := repositories.NewOrder(db)
	//创建order Service
	orderService := services.NewOrderService(order)


	if err:=mq.InitConsume(orderService, productService);err!=nil{
		fmt.Println(err)
	}
	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c
}
