package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/NayronFerreira/microservice_products/configs"
	"github.com/NayronFerreira/microservice_products/internal/infra/database"
	"github.com/NayronFerreira/microservice_products/internal/infra/web/handlers"
)

type Handler interface {
	RegisterRoutes(mux *http.ServeMux)
}

type Server struct {
	config *configs.Config
	db     *sql.DB
}

func NewServer(config *configs.Config, db *sql.DB) *Server {
	return &Server{config: config, db: db}
}

func (s Server) SetupServer() *http.Server {
	mux := http.NewServeMux()

	s.registerHandlers(mux)

	addr := fmt.Sprintf(":%s", s.config.WebServerPort)
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func (s Server) registerHandlers(mux *http.ServeMux) {
	productRepo := database.NewProductRepository(s.db)
	handlers := s.createHandlers(productRepo)

	for _, handler := range handlers {
		handler.RegisterRoutes(mux)
	}
}

func (s Server) createHandlers(productRepo *database.ProductRepository) []Handler {
	return []Handler{
		handlers.NewCreateProductHandler(productRepo),
		handlers.NewGetAllProductsHandler(productRepo),
		handlers.NewUpdateProductHandler(productRepo),
		handlers.NewGetProductByIDHandler(productRepo),
		handlers.NewDeleteProductHandler(productRepo),
	}
}
