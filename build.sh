# Installing Go language

sudo apt install golang-go

# Docker build for publisher/sender

sudo docker build -f sender/Dockerfile .

# Docker build for consumer/reciever

sudo docker build -f reciever/Dockerfile .

# Docker command to start the rabbit message queue server

sudo docker run -it --rm --name rabbitmq -p 5672:5672 -p 8080:15672 rabbitmq:3.11-management

# Start the publisher/sender to push the messages to the rabbitmq server

go run sender/publisher.go

# Start the consumer/reciever to pull the message to the rabbitmq server

go run reciever/consumer.go
