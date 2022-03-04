package api

import (
	"discount-service/handlers"
	"discount-service/infra"
	"discount-service/manager"
	"discount-service/resources/response"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server interface {
	Run()
}

type server struct {
	router         chi.Router
	infra          infra.Infra
	serviceManager manager.ServiceManager
}

// NewServer construct new API server
func NewServer(infra infra.Infra) Server {

	return &server{
		router:         chi.NewRouter(),
		infra:          infra,
		serviceManager: manager.NewServiceManager(infra),
	}
}

func (c *server) Run() {
	c.handlers()
	c.run()
}

func (c *server) run() {
	apiConfig := c.infra.Config().Sub("api")
	host := apiConfig.GetString("host")
	port := apiConfig.GetInt("port")
	addr := fmt.Sprintf("%s:%d", host, port)
	server := &http.Server{
		Addr:         addr,
		Handler:      c.router,
		ReadTimeout:  time.Duration(apiConfig.GetInt("read_timeout")) * time.Second,
		WriteTimeout: time.Duration(apiConfig.GetInt("write_timeout")) * time.Second,
		IdleTimeout:  time.Duration(apiConfig.GetInt("idle_timeout")) * time.Second,
	}

	log.Printf("already running in %s:%d \n", host, port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server start error : %v", err)
	}
}

func (c *server) handlers() {
	checkoutHandler := handlers.NewCheckoutHandler(c.serviceManager.CheckoutService())

	c.router.Post("/checkout", checkoutHandler.HandlerCheckoutOrder)

	// set default handler
	c.router.Get("/", response.Index)
	c.router.NotFound(response.NotFound)
	c.router.MethodNotAllowed(response.MethodNotAllowed)
}
