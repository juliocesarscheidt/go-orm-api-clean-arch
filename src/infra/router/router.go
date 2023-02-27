package router

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/juliocesarscheidt/go-orm-api/application/repository"
	"github.com/juliocesarscheidt/go-orm-api/infra/controller"
	"github.com/juliocesarscheidt/go-orm-api/shared/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func ExtractIpFromRemoteAddr(remoteAddr string) string {
	addressParts := strings.Split(remoteAddr, ":")
	if len(addressParts) > 0 {
		return addressParts[0]
	}
	return ""
}

// LogMiddleware - custom logger middleware method
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Logger.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.RequestURI,
			"host":   r.Host,
			"ip":     ExtractIpFromRemoteAddr(r.RemoteAddr),
		}).Infof("")
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
	// create controllers, injecting repositories on them
	userController := controller.NewUserController(userRepository)
	healthcheckController := controller.NewHealthcheckController()
	// user routes
	router.Path("/api/v1/users").HandlerFunc(userController.CreateUser()).Methods(http.MethodPost)
	router.Path("/api/v1/users/{id:[0-9]+}").HandlerFunc(userController.GetUser()).Methods(http.MethodGet)
	router.Path("/api/v1/users").Queries("page", "{page}", "size", "{size}").HandlerFunc(userController.GetUsers()).Methods(http.MethodGet)
	router.Path("/api/v1/users/{id:[0-9]+}").HandlerFunc(userController.UpdateUser()).Methods(http.MethodPatch)
	router.Path("/api/v1/users/{id:[0-9]+}").HandlerFunc(userController.DeleteUser()).Methods(http.MethodDelete)
	// crosscutting routes
	router.Path("/api/v1/health/live").HandlerFunc(healthcheckController.CheckLiveness()).Methods(http.MethodGet)
	router.Path("/api/v1/health/ready").HandlerFunc(healthcheckController.CheckReadiness()).Methods(http.MethodGet)
	router.Path("/metrics").Handler(promhttp.Handler()).Methods(http.MethodGet)
}
