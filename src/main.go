package main

import (
	"fmt"
	"github.com/FabianSteven/ChatbotWithGo/src/database"
)

func main() {
	db := database.GetConnection()
	fmt.Printf("Hola Mundo")
}
