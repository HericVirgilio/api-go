package database

import (
	"fmt"
	"log"

	"github.com/HericVirgilio/api-go/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){

	var err error

	dsn := "host=localhost user=admin password=admin dbname=apigo port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")

	fmt.Println("Executando migrações...")
	err = DB.AutoMigrate(&domain.User{})
	if err != nil{
		log.Fatal("Falha ao executar migrações: ", err)
	}
	fmt.Println("Migrações concluídas.")
}