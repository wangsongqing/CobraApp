package bootstrap

import (
	"CobraApp/pkg/config"
	"CobraApp/pkg/rabbitmq"
	"fmt"
)

func SetupRabbitMQ() {
	connUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.GetString("rabbitmq.user"), config.GetString("rabbitmq.password"), config.GetString("rabbitmq.host"), config.GetString("rabbitmq.port"))
	rabbitmq.ConnectRabbitMQ(connUrl)
}
