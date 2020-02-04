## Docker commands
- Build image: `docker build -t [image name] -f [docker file] .`
- Run image: `docker run [image name]`
- Access container: `docker exec -it [image name] sh`

## Dev project
- Navigate to project source folder
- Generate api `./proto/_gen_corepb.sh`
- Generate swagger `make gen-swagger`
- Implement api
- Testing api

## Run project
### 1. Requirements
- MySQL: database
- Redis: Cache
- Kafka: Queue
- Jaeger: Tracing
### 2. Run
- Setup environment (MySQL/Redis/Kafka/Jaeger): `docker-compose up`
- Run database script `src/[project name]/migrations/*.sql`
- Build image: `docker build -t [image name] -f [docker file] .`
- Run image: `docker run --network=host --volume [path]/config.yml:/app/config.yaml [image name]`