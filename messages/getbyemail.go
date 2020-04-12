package Message

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetByEmail(w http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		return
	}

	var messageList []Message
	var email string

	vars := mux.Vars(request)
	email = vars["email"]

	messageList = FindByEmail(email)

	json.NewEncoder(w).Encode(GetByEmailResponse{Messages: messageList})
}
