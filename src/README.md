## Consuming API

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

curl -i --silent -X PATCH \
  --url 'http://localhost:8000/api/v1/users/1' \
  --header 'content-type: application/json' \
  --data '{
	"name": "user changed",
	"password": "PASSWORD_CHANGED"
}'

curl -i --silent -X DELETE --url 'http://localhost:8000/api/v1/users/1'

# health
curl --silent -X GET --url 'http://localhost:8000/api/v1/health/live'
curl --silent -X GET --url 'http://localhost:8000/api/v1/health/ready'

# metrics
curl --silent -X GET --url 'http://localhost:8000/metrics'
curl --silent -X GET --url 'http://localhost:8000/metrics' | grep api_http_request
```
