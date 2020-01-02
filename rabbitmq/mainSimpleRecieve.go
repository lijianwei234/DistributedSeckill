//Small_HG
package main

import "rabbitmq/RabbitMq"

func main() {
	rabbitmq := RabbitMq.NewRabbitMQSimple("" +
		"imoocSimple")
	rabbitmq.ConsumeSimple()
}
