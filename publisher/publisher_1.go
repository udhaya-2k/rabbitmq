package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

var rabbit_host = os.Getenv("RABBIT_HOST")
var rabbit_port = os.Getenv("RABBIT_PORT")
var rabbit_user = "guest"
var rabbit_password = "guest"

func main() {

	// Create a new RabbitMQ connection with the required credentials of the RabbitMQ instance
	rabbitmqConn, err := amqp.Dial("amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/")
	if err != nil {
		log.Panic(err, "Failed to create RabbitMQ connection")
	} else {
		log.Println("RabbitMQ connection created!!!")
	}
	defer rabbitmqConn.Close()

	// Open a channel to the RabbitMQ instance.
	ch, err := rabbitmqConn.Channel()
	if err != nil {
		logError(err, "Failed to open a channel")
	} else {
		log.Println("Channel opened!!")
	}
	defer ch.Close()

	// Declare a queue to publish messages
	queue, err := ch.QueueDeclare(
		"Queue_1",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logError(err, "Unable to declare queue")
	}
	log.Println("Queue_1 created")

	// Create a new Fiber server.
	app := fiber.New()

	// Define a handler or route to get messages and publish to the queue.
	app.Get("/publish", func(c *fiber.Ctx) error {
		msg := c.Query("msg")
		if msg == "" {
			log.Println("msg parameter missing!!")
			return c.SendStatus(500)
		}

		// Publish a message to the queue
		err = ch.Publish(
			"",
			queue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			},
		)
		if err != nil {
			logError(err, "Failed to put message in queue")
		}
		log.Println("Message added")

		return c.SendStatus(201)
	})

	log.Fatal(app.Listen(":5050"))
}

func logError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
