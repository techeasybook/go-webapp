package handler

import (
	"fmt"
	"net/http"
)

// HealthCheck this method returns current status of the service.
func HealthCheck(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("go-webapp is up and running..!\n")
	fmt.Fprintf(writer, "go-webapp is up and running..!")
}

// GetUser this method returns the username.
func GetUser(writer http.ResponseWriter, request *http.Request) {
	userName := "TechBook"
	fmt.Printf("returning the user name: %s \n", userName)
	fmt.Fprintf(writer, "User name is: "+userName)
}
