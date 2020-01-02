//Small_HG
package main

import "rabbitmq/RabbitMq"

func main()  {
	imoocOne:=RabbitMq.NewRabbitMQRouting("exImooc","imooc_two")
	imoocOne.RecieveRouting()
}
