package main

import (
	"customer-family-crud-backend/driver/database"
	"customer-family-crud-backend/driver/migration"
	"customer-family-crud-backend/interfaces/handler"
	repository "customer-family-crud-backend/repository/impl"
	service "customer-family-crud-backend/service/impl"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env file, using system environment variable")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("env DATABASE_URL is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	allowedOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		log.Fatal("env CORS_ALLOWED_ORIGIN variable is not set")
	}

	// db connection
	db, err := database.NewPostgresDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// running migrations
	migration.RunMigration(dsn)
	migration.InsertInitialNationalities(db)

	customerRepo := repository.NewCustomerRepositoryImpl(db)
	nationalityRepo := repository.NewNationalityRepositoryImpl(db)

	customerService := service.NewCustomerService(customerRepo)
	nationalityService := service.NewNationalityService(nationalityRepo)

	customerHandler := handler.NewCustomerHandler(customerService)
	nationalityHandler := handler.NewNationalityHandler(nationalityService)

	router := mux.NewRouter()

	// customer handlers
	router.HandleFunc("/api/customers", customerHandler.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/customers/{id}", customerHandler.UpdateCustomer).Methods(http.MethodPut)
	router.HandleFunc("/api/customers/{id}", customerHandler.GetCustomerByID).Methods(http.MethodGet)
	router.HandleFunc("/api/customers/{id}", customerHandler.DeleteCustomer).Methods(http.MethodDelete)
	router.HandleFunc("/api/customers", customerHandler.GetAllCustomers).Methods(http.MethodGet).Queries("page", "{page}", "limit", "{limit}")

	// nationality handlers
	router.HandleFunc("/api/nationalities", nationalityHandler.GetAllNationalities).Methods(http.MethodGet)

	// cors configuration
	c := cors.New(cors.Options{
		AllowedOrigins: []string{allowedOrigin},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
		Debug:          true,
	})

	handlerWithCORS := c.Handler(router)

	log.Printf("backend server running on port: %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), handlerWithCORS))
}
