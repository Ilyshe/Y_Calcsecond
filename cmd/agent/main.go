package main

import (
	"log"

	"github.com/Ilushe/Y_Calcsecond/internal/application"
)

func main() {
	agent := application.NewAgent()
	log.Println("Starting Agent...")
	agent.Run()
}
