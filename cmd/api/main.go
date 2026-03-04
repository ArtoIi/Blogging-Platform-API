package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ArtoIi/Blogging-Platform-API/internal/application"
	"github.com/ArtoIi/Blogging-Platform-API/internal/infrastructure/mysql"
	"github.com/ArtoIi/Blogging-Platform-API/internal/interfaces/database"
	routeshttp "github.com/ArtoIi/Blogging-Platform-API/internal/interfaces/http"
	"github.com/ArtoIi/Blogging-Platform-API/internal/interfaces/http/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.NewMySQLConnection("root", "senha_poderosa", "localhost:3306", "blog_db")

	if err != nil {
		log.Fatal("Erro ao conectar no banco através do pacote database:", err)
	}
	defer db.Close()

	fmt.Println("✅ Banco de dados conectado com sucesso!")

	repo := mysql.NewMySQLPostRepository(db)
	service := application.NewPostService(repo)
	handler := handlers.NewPostHandler(service)

	mux := routeshttp.SetupRoutes(handler)

	port := ":8080"
	fmt.Printf("Servidor iniciado em http://localhost%s\n", port)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
