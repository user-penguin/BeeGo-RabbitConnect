package controllers

import (
	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
	"log"
	"fmt"
	_ "github.com/streadway/amqp"
)

type MainController struct {
	beego.Controller
}

var text []string

func (this *MainController) MessageCatcher() {
	text = this.GetStrings("text")

	conn, err := amqp.Dial("amqp://diman:1@172.17.17.200:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"transfer", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := text[0]
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.TplName = "default/hello-sitepoint.tpl"
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}