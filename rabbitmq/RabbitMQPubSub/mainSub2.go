//Small_HG
package main

import "rabbitmq/RabbitMq"

func main() {
	rabbitmq := RabbitMq.NewRabbitMQPubSub("" +
		"newProduct")
	rabbitmq.RecieveSub()
}
