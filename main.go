/*
package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/scch94/apirest/db"
	"github.com/scch94/apirest/models"
	"github.com/scch94/apirest/routes"
)

func main() {
	//creas conexion
	db.DBconnction()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	//creamos el router donde va a estar levantado el servidor
	router := mux.NewRouter()
	//esta funcion de la libreria max nos permite en el primer parametro decir la ruta que tomara y en la segunda una funcion
	//la cual tiene dos parametros el primero un writer =como respondemos y read como obtenemos los parametros
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")
	//aqui inicialisamos el servidor para eso necesitamos el puerto y el router que emos creado
	http.ListenAndServe(":3000", router)

}
*/

//************************************

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"log"
	"net/http"
)

const configFile = "configuration/config.json"

type Config struct {
	Key1 string `json:"nombre"`
	Key2 string `json:"saludo"`
	Key3 uint   `json:"edad"`
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error al obtener la ruta del ejecutables: ", err)
		return
	}
	fmt.Println(exePath)
	exeDir := filepath.Dir(exePath)
	configFilePath := filepath.Join(exeDir, configFile)
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error al leer el archivo de configuración:", err)
		return
	}
	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error al parser el archivo json:", err)
		return
	}
	fmt.Println("Key1:", config.Key1)
	fmt.Println("Key2:", config.Key2)
	fmt.Println("Key3:", config.Key3)
	http.HandleFunc("/", handleRequest)
	fmt.Println("Servidor escuchando en el puerto :311131...")
	log.Fatal(http.ListenAndServe(":3111", nil))

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nueva solicitud recibida:")
	fmt.Println("Método:", r.Method)
	fmt.Println("URL:", r.URL.Path)
	fmt.Println("Dirección IP del cliente:", r.RemoteAddr)

	// Leer y registrar el cuerpo de la solicitud
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error al leer el cuerpo de la solicitud:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	fmt.Println("Cuerpo de la solicitud:", string(body))
	jsonResponse := map[string]string{
		"status":      "0",
		"forwardRef":  "3",
		"Description": "No errors",
	}
	// Leer y registrar el cuerpo de la solicitud
	// Serializar el objeto JSON
	jsonResponseBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, "Error al serializar el objeto JSON", http.StatusInternalServerError)
		return
	}

	// Establecer el tipo de contenido de la respuesta como "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Escribir el objeto JSON en el cuerpo de la respuesta
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponseBytes)
}

////////////////////////////////////////////////////////////////////////////////

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/gin-gonic/gin"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// )

// const configFile = "configuration/config.json"

// type Config struct {
// 	Key1 string `json:"nombre"`
// 	Key2 string `json:"saludo"`
// 	Key3 uint   `json:"edad"`
// }

// // Mapa de usuarios y contraseñas
// var users = map[string]string{
// 	"usuario1": "password1",
// 	"usuario2": "password2",
// }

// func main() {
// 	// Cargar configuración
// 	exePath, err := os.Executable()
// 	if err != nil {
// 		fmt.Println("Error al obtener la ruta del ejecutable:", err)
// 		return
// 	}
// 	exeDir := filepath.Dir(exePath)
// 	configFilePath := filepath.Join(exeDir, configFile)
// 	data, err := ioutil.ReadFile(configFilePath)
// 	if err != nil {
// 		fmt.Println("Error al leer el archivo de configuración:", err)
// 		return
// 	}
// 	var config Config
// 	err = json.Unmarshal(data, &config)
// 	if err != nil {
// 		fmt.Println("Error al parsear el archivo JSON:", err)
// 		return
// 	}
// 	fmt.Println("Key1:", config.Key1)
// 	fmt.Println("Key2:", config.Key2)
// 	fmt.Println("Key3:", config.Key3)

// 	// Crear un enrutador Gin
// 	r := gin.Default()

// 	// Middleware de autenticación básica
// 	r.Use(AuthMiddleware())

// 	// Ruta principal
// 	r.POST("/", handleRequest)

// 	// Iniciar el servidor
// 	fmt.Println("Servidor escuchando en el puerto 3111...")
// 	log.Fatal(r.Run(":3111"))
// }

// // Middleware para autenticación básica
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		username, password, ok := c.Request.BasicAuth()
// 		if !ok {
// 			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			return
// 		}
// 		expectedPassword, ok := users[username]
// 		if !ok || password != expectedPassword {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			return
// 		}
// 		c.Set(gin.AuthUserKey, username)
// 		c.Next()
// 	}
// }

// func handleRequest(c *gin.Context) {
// 	fmt.Println("Nueva solicitud recibida:")
// 	fmt.Println("Método:", c.Request.Method)
// 	fmt.Println("URL:", c.Request.URL.Path)
// 	fmt.Println("Dirección IP del cliente:", c.ClientIP())

// 	// Leer y registrar el cuerpo de la solicitud
// 	body, err := ioutil.ReadAll(c.Request.Body)
// 	if err != nil {
// 		fmt.Println("Error al leer el cuerpo de la solicitud:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
// 		return
// 	}
// 	fmt.Println("Cuerpo de la solicitud:", string(body))

// 	// Crear la respuesta JSON
// 	jsonResponse := map[string]string{
// 		"status":      "0",
// 		"forwardRef":  "3",
// 		"Description": "No errors",
// 	}

// 	// Enviar la respuesta JSON
// 	c.JSON(http.StatusOK, jsonResponse)
// }
