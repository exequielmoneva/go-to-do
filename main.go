package main

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/structs"
	"net/http"
)

func main() {
	// Para crear una Api necesitamos delarar Mux, que maneja las respuestas de los endpoint
	// se declaran mediante HandleFunc("ruta del endpoint", funcion)
	// La funcion requiere declarar ResponseWriter para la respuesta y Request para leer datos del request
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet{
			data := structs.Response{
				Code: http.StatusOK,
				Body: "Json of all entries",
			}
			json.NewEncoder(w).Encode(data)
		}

	})

	fmt.Println("Server is running on localhost:3000")
	http.ListenAndServe("localhost:3000",mux)
}
