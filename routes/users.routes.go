package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/scch94/apirest/db"
	"github.com/scch94/apirest/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	cratedUser := db.DB.Create(&user)
	err := cratedUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	fmt.Println(user)
	json.NewEncoder(w).Encode(&user)
}
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
