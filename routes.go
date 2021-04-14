package server

import (
	"fmt"
	"net/http"
)

func HomeRoute(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Red social backend :)")
}
