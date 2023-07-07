package main

import (
	"example.com/gin_wire/di"
)

func main() {
	server := di.InitializeServer()
	engine := server.GetEngine()
	engine.Run(":8080")
}
