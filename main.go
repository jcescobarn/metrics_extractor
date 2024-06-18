package main

import (
	"fmt"
	"log"
	"net/http"

	"metricsExtractor/internal"
	"metricsExtractor/services"
)

func main() {

	systemService := services.NewSystemService()
	info := internal.NewInfo(systemService)
	websocketService := internal.NewWebsocket(info)

	http.HandleFunc("/ws", websocketService.HandleConnections)

	var port string = "8090"
	fmt.Printf("Servidor iniciado en el puerto %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v\n", err)
	}
}
