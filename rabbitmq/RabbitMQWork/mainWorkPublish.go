//Small_HG
package main

import (
	"strconv"
	"time"
	"fmt"
	"rabbitmq/RabbitMq"
)

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("" +
		"imoocSimple")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello imooc!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}