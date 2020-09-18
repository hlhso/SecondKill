package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	//已购数量
	sum int64 = 0
	//预存商品数量
	productNum int64 = 100000
	//互斥锁
	mutex sync.Mutex
	//计数
	count int64 = 0
)

//获取秒杀商品
func GetOneProduct() bool {
	//加锁
	mutex.Lock()
	defer mutex.Unlock()
	count += 1
	// 限流，没100个值允许成功一个
	fmt.Printf("sum: %d  count: %d\n", sum, count)
	//if count%100 == 0 {
		// 判断是否超过限制，防止超卖
		if sum < productNum {
			sum += 1
			fmt.Printf("sum: %d  count: %d\n", sum, count)
			return true
		}
	//}
	return false
}

func GetProduct(w http.ResponseWriter, req *http.Request) {
	if GetOneProduct() {
		w.Write([]byte("true"))
		return
	}
	w.Write([]byte("false"))
	return
}

func main() {
	http.HandleFunc("/getOne", GetProduct)
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		log.Fatal("Err:", err)
	}
}
