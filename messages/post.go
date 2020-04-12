package Message

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
)

func Post(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		return
	}

	var errors []string

	var uuid gocql.UUID
	var message Message
	var success bool
	var err error

	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&message)

	success = false

	if err == nil {
		fmt.Println("Creating a new message")

		uuid, err = InsertMessage(message)

		if err != nil {
			errors = append(errors, err.Error())
		} else {
			success = true
		}
	} else {
		errors = append(errors, "Decoding error: "+err.Error())
	}

	if success {
		fmt.Println("message_id", uuid)
		json.NewEncoder(w).Encode(PostResponse{ID: uuid})
	} else {
		fmt.Println("errors", errors)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errors})
	}
}
