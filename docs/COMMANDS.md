### Run Postgres in Docker

docker run --name my-postgres-db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=case -p 5432:5432 -d
postgres

### Dockefile

1. Build the Docker Image:

docker build -t se-school-case .

2. Run the Docker Container:

docker run -p 3000:3000 --name se-school-case-container se-school-case

### Docker Compose

docker-compose up --build

### RabbitMQ

docker run -d --hostname my-rabbit --name some-rabbit rabbitmq
