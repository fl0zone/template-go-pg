package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	var db *sql.DB
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	databaseURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		log.Fatal("missing DATABASE_URL")
		return
	}
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "80"
	}

	shouldCreateTable := os.Getenv("GO_ENV") != "production"

	db, err = ConnectDB(databaseURL, shouldCreateTable)
	if err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
		os.Exit(1)
	}

	router := gin.Default()

	router.GET("/todo", getAllTodoHandler(db))
	router.GET("/todo/:id", getTodoHandler(db))
	router.POST("/todo", createTodoHandler(db))
	router.PUT("/todo/:id", toggleTodoHandler(db))
	router.DELETE("/todo/:id", deleteTodoHandler(db))

	router.Run(":" + port)
}
