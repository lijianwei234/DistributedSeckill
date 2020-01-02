//Small_HG
package main

import (
	"fmt"
	"rabbitmq/RabbitMq"
)

func main()  {
	rabbitmq := RabbitMq.NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.PublishSimple("Hello imooc!")
	fmt.Println("发送成功！")
}