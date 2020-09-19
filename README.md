# SecondKill
**基于nsq的分布式秒杀系统**
# 文件目录（部分前端文件夹压缩了）
.<br>
├── backend<br>
│   ├── main.go<br>
│   └── web.rar<br>
├── common<br>
│   ├── common.go<br>
│   ├── consistent.go<br>
│   ├── filter.go<br>
│   └── ip.go<br>
├── config<br>
│   ├── conf.ini<br>
│   └── setting.go<br>
├── consumer.go<br>
├── database<br>
│   ├── mysql.go<br>
│   └── redis.go<br>
├── datamodels<br>
│   ├── message.go<br>
│   ├── order.go<br>
│   ├── product.go<br>
│   └── user.go<br>
├── encrypt<br>
│   └── aes.go<br>
├── getOne.go<br>
├── go.mod<br>
├── go.sum<br>
├── mq<br>
│   └── mq.go<br>
├── repositories<br>
│   ├── order_repository.go<br>
│   ├── product_repository.go<br>
│   └── user_repository.go<br>
├── services<br>
│     ├── order_service.go<br>
│     ├── product_services.go<br>
│     └── user_service.go<br>
├── validate.go<br>
└── web<br>
　├── controllers<br>
　│    ├── product_controller.go<br>
　│    └── user_controller.go<br>
　├── htmlProductShow<br>
　│    └── htmlProduct.html<br>
　├── main.go<br>
　├── middleware<br>
　│    └── auth.go<br>
　├── public.rar<br>
　├── tool<br>
　│    └── cookie.go<br>
　└── views<br>
　　　├── product<br>
　　　│    ├── result.html<br>
　　　│    └── view.html<br>
　　　├── shared<br>
　　　│    ├── error.html<br>
　　　│    ├── layout.html<br>
　　　│    └── productLayout.html<br>
　　　├── template<br>
　　　│    └── product.html<br>
　　　└── user<br>
　　　　├── login.html<br>
　　　　└── register.html<br>
# 基本组件
* 前端框架:后台管理beagle 电商系统namira
* web框架:iris
* 消息队列:nsq
* 数据库:mysql
* 数据库操作框架:gorm
* 配置管理:ini
* cookie加密:aes cbc模式
* 一致性hash算法和虚拟节点实现分布式验证

