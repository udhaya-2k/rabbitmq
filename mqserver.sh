# Running rabbitmq server in the docker container

echo "tarus" | sudo -S docker run -it --rm --name rabbitmq -p 5672:5672 -p 8080:15672 rabbitmq:3.11-management
