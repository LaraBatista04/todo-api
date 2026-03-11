// main.go
package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"todo_API/database"
	"todo_API/routes"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Iniciando conexão com o banco de dados...")
	if err := database.Connect(); err != nil {
		slog.Error("Falha ao inicializar conexão com o banco de dados", "error", err)
		os.Exit(1)
	}

	r := gin.Default()

	slog.Info("Configurando rotas da API...")
	routes.SetupRoutes(r)

	slog.Info("Servidor rodando na porta :8080")
	if err := r.Run(":8080"); err != nil {
		slog.Error("Falha ao iniciar o servidor", "error", err)
		os.Exit(1)
	}
}