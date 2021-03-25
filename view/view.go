package view

import (
	"GO-GITHUB/config/db"
	"GO-GITHUB/models"
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterView(response http.ResponseWriter, request *http.Request) {
	// response.Header().Add("Content-Type", "application/json")

	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	if request.Method != "POST" {
		tmpl.Execute(response, nil)
		return
	}

	var user models.User
	//json.NewDecoder(request.Body).Decode(&user)
	user = models.User{
		Username:  request.FormValue("username"),
		Firstname: request.FormValue("firstname"),
		Lastname:  request.FormValue("lastname"),
		Email:     request.FormValue("email"),
		Password:  request.FormValue("password"),
	}

	coll, err := db.GetDBCollection()
	if err != nil {
		log.Fatal(err)
	}

	// Before inserting check, if the username already exists or not.
	var check models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"username", user.Username}}
	err = coll.FindOne(ctx, filter).Decode(&check)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
			if err != nil {
				//message := "Error While Hashing Password, Try Again"
				//json.NewEncoder(response).Encode(message)
				return
			}
			user.Password = string(hash)

			_, err = coll.InsertOne(ctx, user)
			if err != nil {
				//message := "Error While Creating User, Try Again"
				//json.NewEncoder(response).Encode(message)
				return
			}
			//message := "Registration Successful"
			//json.NewEncoder(response).Encode(message)
			tmpl.Execute(response, struct{ Success bool }{true})
			return
		}

		// message := err.Error()
		//json.NewEncoder(response).Encode(message)
		return
	}

	message := "Username already Exists!!"
	json.NewEncoder(response).Encode(message)

}

func LoginView(response http.ResponseWriter, request *http.Request) {
	// response.Header().Add("Content-Type", "application/json")

	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	if request.Method != "POST" {
		tmpl.Execute(response, nil)
		return
	}
	var user models.User
	// json.NewDecoder(request.Body).Decode(&user)
	user = models.User{
		Username: request.FormValue("username"),
		Password: request.FormValue("password"),
	}

	coll, err := db.GetDBCollection()
	if err != nil {
		log.Fatal(err)
	}

	var check models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"username", user.Username}}
	err = coll.FindOne(ctx, filter).Decode(&check)

	if err != nil {
		message := "Invalid username"
		json.NewEncoder(response).Encode(message)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(check.Password), []byte(user.Password))

	if err != nil {
		message := "Invalid password"
		json.NewEncoder(response).Encode(message)
		return
	}

	// message := "Login Successful"
	// json.NewEncoder(response).Encode(message)
	tmpl.Execute(response, struct{ Success bool }{true})
}
