// cmd/api/main.go
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/HericVirgilio/api-go/internal/handler/http"
	"github.com/HericVirgilio/api-go/internal/repository/postgres"
	"github.com/HericVirgilio/api-go/internal/usecase"
	"github.com/HericVirgilio/api-go/pkg/database"
)

func main() {
	// 1. Inicializar a conexão com o banco de dados
	database.InitDB()
	db := database.DB

	// 2. Injetar as dependências (montando as camadas)
	// Camada de Repositório
	userRepository := postgres.NewUserPostgresRepository(db)
	// Camada de Caso de Uso
	userUseCase := usecase.NewUserUseCase(userRepository)
	// Camada de Handler
	userHandler := http.NewUserHandler(userUseCase)

	// 3. Configurar o Roteador Gin
	router := gin.Default()

	// 4. Definir as rotas da API
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	// 5. Iniciar o servidor
	router.Run(":8080")
}