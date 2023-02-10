# Simple Go API implementing clean architecture, using Mux, Go ORM and Prometheus


## Running with docker compose

```bash
docker-compose up -d mysql
docker-compose logs -f --tail 100 mysql

docker-compose up -d --build go-orm-api
docker-compose logs -f --tail 100 go-orm-api
```

## Running with Docker

```bash
docker image build --tag juliocesarmidia/go-orm-api:latest .

docker container run --rm --name go-orm-api juliocesarmidia/go-orm-api:latest
```

## Testing API

```bash
curl --silent -X POST \
  --url 'http://localhost:8000/api/v1/users' \
  --header 'accept: application/json' \
  --header 'content-type: application/json' \
  --data '{
	"name": "user",
	"email": "user@mail.com",
	"password": "PASSWORD"
}'

curl --silent -X GET --url 'http://localhost:8000/api/v1/users?page=0&size=10'

curl --silent -X GET --url 'http://localhost:8000/api/v1/users/1'

curl -i --silent -X PUT \
  --url 'http://localhost:8000/api/v1/users/1' \
  --header 'content-type: application/json' \
  --data '{
	"name": "user changed",
	"password": "PASSWORD_CHANGED"
}'

curl -i --silent -X DELETE --url 'http://localhost:8000/api/v1/users/1'

# health
curl --silent -X GET --url 'http://localhost:8000/healthcheck'

# metrics
curl --silent -X GET --url 'http://localhost:8000/metrics'
curl --silent -X GET --url 'http://localhost:8000/metrics' | grep api_http_request
```
