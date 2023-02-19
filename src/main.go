package main

import (
	"log"
	"net/http"
	"time"

	applicationrepository "github.com/juliocesarscheidt/go-orm-api/application/repository"
	infrarepository "github.com/juliocesarscheidt/go-orm-api/infra/repository"
	"github.com/juliocesarscheidt/go-orm-api/infra/router"
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var userRepository applicationrepository.UserRepository
	// check the usage of repository in memory
	if utils.GetInMemoryDbConfig() {
		userRepository = infrarepository.UserRepositoryMemory{}
	} else {
		db, _ := gorm.Open(mysql.Open(utils.GetDbConnectionStringConfig()), &gorm.Config{})
		userRepository = infrarepository.UserRepositoryDatabase{Db: db}
	}
	// call migrations
	if err := userRepository.MigrateUser(); err != nil {
		panic(err)
	}
	// create router and its routes
	r := router.GetRouter()
	router.InjectRoutes(r, userRepository)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	utils.Logger.Info("Server listening on 0.0.0.0:8000")
	log.Fatal(srv.ListenAndServe())
}
