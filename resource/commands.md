### Run Postgres in Docker
docker run --name my-postgres-db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=case -p 5432:5432 -d postgres