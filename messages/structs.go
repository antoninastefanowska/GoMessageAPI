package Message

import (
	"github.com/gocql/gocql"
)

type GetByEmailRequest struct {
	Email string `json:"email"`
}

type GetMessageResponse struct {
	Message Message `json:"message"`
}

type GetByEmailResponse struct {
	Messages []Message `json:"messages"`
}

type PostResponse struct {
	ID gocql.UUID `json:"id"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
