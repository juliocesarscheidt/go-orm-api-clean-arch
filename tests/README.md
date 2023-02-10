## Running unit tests

#### With Docker

```bash
docker image build --tag juliocesarmidia/go-orm-api-test:latest -f test.Dockerfile .

docker container run --rm --name go-orm-api-test juliocesarmidia/go-orm-api-test:latest
```

#### With go CLI

```bash
go vet
go test tests/**/**/*_test.go -v

go test -cover \
  -coverpkg=github.com/juliocesarscheidt/go-orm-api/application/usecase \
  -coverprofile cover.out tests/**/**/*_test.go -v

go tool cover -html=cover.out -o coverage.html
```
