package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"sistema-libros-electronicos/internal/config"
	"sistema-libros-electronicos/internal/database"
	"sistema-libros-electronicos/internal/routes"
)

func main() {
	config.LoadEnv()

	err := database.Connect()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	} else {
		log.Println("Conexión exitosa a la base de datos")
	}

	//Crea el servidor web (Instancia de Gin)
	router := gin.Default()

	// Registrar rutas desde el paquete routes
	routes.RegisterRoutes(router)

	// Configuración del Servidor
	port := config.GetServerPort()

	log.Println("Servidor corriendo en http://localhost:" + port)

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Error al iniciar el servidor", err)
	}
}
