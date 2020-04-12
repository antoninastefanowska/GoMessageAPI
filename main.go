package main

import (
	"log"
	"net/http"

	Cassandra "../../src/messageapi/cassandra"
	Message "../../src/messageapi/messages"
	"github.com/gorilla/mux"
)

func main() {
	cassandraSession := Cassandra.Session
	defer cassandraSession.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/message", Message.Post)
	router.HandleFunc("/api/messages/{email}", Message.GetByEmail)
	log.Fatal(http.ListenAndServe(":8080", router))
}
