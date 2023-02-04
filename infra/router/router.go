package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliocesarscheidt/go-orm-api/domain/repository"
	"github.com/juliocesarscheidt/go-orm-api/infra/controller"
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// LogMiddleware - custom logger middleware method
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Logger.Infof(r.RequestURI)
		// call next handler
		next.ServeHTTP(w, r)
	})
}

// GetRouter - it returns the mux Router with the injected middlewares
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(LogMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
	return router
}

func InjectRoutes(router *mux.Router, userRepository repository.UserRepository) {
	router.Path("/api/v1/users").HandlerFunc(controller.CreateUser(userRepository)).Methods(http.MethodPost)
	router.Path("/api/v1/users").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.GetUsers(userRepository)).Methods(http.MethodGet)
	router.Path("/api/v1/users/{id:[0-9]+}").HandlerFunc(controller.GetUser(userRepository)).Methods(http.MethodGet)
	router.Path("/api/v1/users/{id:[0-9]+}").HandlerFunc(controller.UpdateUser(userRepository)).Methods(http.MethodPut)
	router.Path("/api/v1/users/{id:[0-9]+}").HandlerFunc(controller.DeleteUser(userRepository)).Methods(http.MethodDelete)
	router.Path("/metrics").Handler(promhttp.Handler()).Methods(http.MethodGet)
	router.Path("/healthcheck").Handler(controller.Healthcheck()).Methods(http.MethodGet)
}
