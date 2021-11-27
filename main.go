package main

import (
	"fmt"

	"github.com/suren-m/go-server/models"
	"github.com/suren-m/go-server/routes"
)

func main() {
	fmt.Println("Main package - main file")
	models.AllUsers()
	routes.APIPostRoute()
}
