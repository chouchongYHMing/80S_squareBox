package main

import (
	"Mygin/model"
	"Mygin/routes"
)

func main() {
	model.InitDb()

	routes.InitRouter()
}
