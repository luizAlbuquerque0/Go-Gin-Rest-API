package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizAlbuquerque0/go-api-rest/database"
	"github.com/luizAlbuquerque0/go-api-rest/routes"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Luiz",
	})
}

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
