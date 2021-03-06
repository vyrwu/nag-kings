/*
 * Kings API
 *
 * This is a sample Kings Server based on the OpenAPI 3.0 specification.
 *
 * API version: 0.1.0
 * Contact: dev.anowak@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"

	sw "example/com/nag/kings/go"
)

func main() {
	// TODO: HTTP to HTTPS redirect, TLS with my own certs

	log.Println("Server started")

	router := sw.NewRouter()

	fs := http.FileServer(http.Dir("./swaggerui"))
	router.PathPrefix("/swaggerui").Handler(http.StripPrefix("/swaggerui/", fs))

	log.Fatal(http.ListenAndServe(":8081", router))
}
