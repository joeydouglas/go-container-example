package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := "80"
	enviro := os.Getenv("APP_ENV")

	switch enviro {
	case "dev":
		http.HandleFunc("/", devFunction)
		fmt.Println("Environment set to " + enviro)
	case "prod":
		http.HandleFunc("/", prodFunction)
		fmt.Println("Environment set to " + enviro)
	default:
		fmt.Println("No environment specified. Closing server.")
		return

	}

	fmt.Println("Server listening on port: " + port)
	http.ListenAndServe(":"+port, nil)

}

func devFunction(w http.ResponseWriter, r *http.Request) {
	appEnv := "This is DEV."
	fmt.Fprint(w, "", appEnv)
}

func prodFunction(w http.ResponseWriter, r *http.Request) {
	appEnv := "This is PROD."
	fmt.Fprint(w, "", appEnv)
}
