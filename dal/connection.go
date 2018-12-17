package dal

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/aneri/gqlgen-authentication/models"
	auth "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	Db      *gorm.DB
	Context context.Context
	Auth    *auth.Token
}

var Once sync.Once
var instance *DbConnection

func Connect() *DbConnection {

	Once.Do(func() {
		context := context.Background()
		connectionString := "postgresql://root@localhost:26257/training?sslmode=disable"
		db, err := gorm.Open("postgres", connectionString)
		if err != nil {
			log.Fatal("Error while initializing database", err)
		}
		fmt.Println("Database successfully connected")
		instance = &DbConnection{
			Db:      db,
			Context: context,
		}
		db.AutoMigrate(&models.Job{}, &models.User{}, &models.Application{})
	})
	return instance
}
