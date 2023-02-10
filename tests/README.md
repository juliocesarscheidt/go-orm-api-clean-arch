## Running unit tests

```bash
go vet
go test tests/**/**/*_test.go -v

go test -cover \
  -coverpkg=github.com/juliocesarscheidt/go-orm-api/application/usecase \
  -coverprofile cover.out tests/**/**/*_test.go -v

go tool cover -html=cover.out -o coverage.html

go test tests/application/usecase/create_user_test.go -v
go test tests/application/usecase/get_user_test.go -v
go test tests/application/usecase/get_users_test.go -v
go test tests/application/usecase/update_user_test.go -v
```
