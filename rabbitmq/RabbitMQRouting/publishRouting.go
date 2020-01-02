//Small_HG
package main

import (
	"strconv"
	"time"
	"fmt"
	"rabbitmq/RabbitMq"
)

func main()  {
	imoocOne:=RabbitMq.NewRabbitMQRouting("exImooc","imooc_one")
	imoocTwo:=RabbitMq.NewRabbitMQRouting("exImooc","imooc_two")
	for i := 0; i <= 10; i++ {
		imoocOne.PublishRouting("Hello imooc one!" + strconv.Itoa(i))
		imoocTwo.PublishRouting("Hello imooc Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
