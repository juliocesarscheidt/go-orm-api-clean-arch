## Running unit tests

### With Docker

```bash
docker image build \
  --tag juliocesarmidia/go-orm-api-test:latest \
  -f ./src/test.Dockerfile ./src

# run unit tests only
docker container run --rm \
  --name go-orm-api-test juliocesarmidia/go-orm-api-test:latest

# run vet and unit tests
docker container run --rm \
  --name go-orm-api-test juliocesarmidia/go-orm-api-test:latest \
  sh -c "go vet && go test tests/**/**/*_test.go -v"
```

### With go CLI

```bash
cd ./src

go vet

go test tests/**/**/*_test.go -v
go test -race tests/**/**/*_test.go -v

go test -cover \
  -coverpkg=github.com/juliocesarscheidt/go-orm-api/application/usecase \
  -coverprofile cover.out tests/**/**/*_test.go -v

go tool cover -html=cover.out -o coverage.html
```
