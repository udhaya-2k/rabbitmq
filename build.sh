# Installing Go language

sudo apt install golang-go

# Installing Docker

sudo apt install docker.io

# Docker build for publisher/sender

sudo docker build -f sender/Dockerfile .

# Docker build for consumer/reciever

sudo docker build -f reciever/Dockerfile .

# Start the publisher/sender and reciever/consumer

gnome-terminal -- sh -c "cd sender && go run publisher.go" && gnome-terminal -- sh -c "cd reciever && go run consumer.go"
