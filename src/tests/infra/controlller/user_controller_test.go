package controller

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path/filepath"
	"testing"

	dto "github.com/juliocesarscheidt/go-orm-api/application/dto"
	controller "github.com/juliocesarscheidt/go-orm-api/infra/controller"
	infrarepository "github.com/juliocesarscheidt/go-orm-api/infra/repository"
)

type MysqlContainer struct {
	Container        *mysql.MySQLContainer
	ConnectionString string
}

func CreateMysqlContainer(ctx context.Context) (*MysqlContainer, error) {
	mysqlContainer, err := mysql.RunContainer(ctx,
		testcontainers.WithImage("mysql:8.0"),
		mysql.WithDatabase("go_orm_api"),
		mysql.WithUsername("root"),
		mysql.WithPassword("admin"),
		mysql.WithScripts(filepath.Join("..", "testdata", "init-db.sql")),
	)
	connString, err := mysqlContainer.ConnectionString(ctx, "parseTime=True")
	if err != nil {
		return nil, err
	}
	return &MysqlContainer{
		Container:        mysqlContainer,
		ConnectionString: connString,
	}, err
}

func TestCreateUserSuccess(t *testing.T) {
	ctx := context.Background()
	mysqlContainer, err := CreateMysqlContainer(ctx)
	if err != nil {
		t.Errorf("Expected err to be nil, got %v", err)
	}
	defer mysqlContainer.Container.Terminate(ctx)

	db, _ := gorm.Open(mysqlDriver.Open(mysqlContainer.ConnectionString), &gorm.Config{})
	userRepository := infrarepository.UserRepositoryDatabase{Db: db}

	userController := controller.NewUserController(userRepository)

	// fake HTTP request
	createUserDto := &dto.CreateUserDto{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "PASSWORD",
	}
	id, err := userController.CreateUserUsecase.Execute(createUserDto)
	assert.NoError(t, err)
	assert.NotNil(t, id)

	// get the just created user
	user, err := userController.GetUserUsecase.Execute(&dto.GetUserDto{Id: id})
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, createUserDto.Email, user.Email)
	assert.Equal(t, createUserDto.Name, user.Name)
}
