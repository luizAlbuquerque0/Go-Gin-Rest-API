package main

import (
	"github.com/luizAlbuquerque0/go-api-rest/database"
	"github.com/luizAlbuquerque0/go-api-rest/routes"
)



func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
