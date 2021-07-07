package controller

import (
	"fmt"
	"net/http"
)

// HomeController provides the controller functionalities
type HomeController struct {
}

// GetHomePage serves the main page
func (hController HomeController) GetHomePage(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
}
