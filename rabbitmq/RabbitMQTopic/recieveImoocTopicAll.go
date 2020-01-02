//Small_HG
package main

import "rabbitmq/RabbitMq"

func main()  {
	imoocOne:=RabbitMq.NewRabbitMQTopic("exImoocTopic","#")
	imoocOne.RecieveTopic()
}
