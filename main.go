package main

import (
	"log"
	"net/http"
	"time"

	"github.com/juliocesarscheidt/go-orm-api/domain/entity"
	"github.com/juliocesarscheidt/go-orm-api/infra/repository"
	"github.com/juliocesarscheidt/go-orm-api/infra/router"
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connectionString := utils.GetDbConnectionString()
	db, _ := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	// db, _ := gorm.Open(sqlite.Open("go_orm_api.db"), &gorm.Config{})
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		utils.Logger.Errorf("Err %v", err)
	}
	// create repositories
	userRepository := repository.UserRepositoryDatabase{Db: db}
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
