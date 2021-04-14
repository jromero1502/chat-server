package server

import "net/http"

type Middleware func(handler http.HandlerFunc) http.HandlerFunc

func LoginMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		PrintServerInfo("Logging middleware...")
		//TODO: Write the logic of the login middleware
		handler(w, request)
	}
}
