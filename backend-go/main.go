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

	customerService := service.NewCustomerService(customerRepo)

	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/api/customers", customerHandler.CreateCustomer).Methods(http.MethodPost)

	log.Printf("backend server running on port: %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
