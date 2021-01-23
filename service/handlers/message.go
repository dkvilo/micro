// Package classification Message API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: David Kviloria<dkviloria@gmail.com> https://katakuraa.dev/
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dkvilo/micro/service/models"
	"github.com/julienschmidt/httprouter"
)

// swagger:response messageResponse
type messageResponse struct {
	// in: body
	Body models.Message
}

// Messenger structure for message
type Messenger struct {
	l	*log.Logger
	port string
}

// NewMessenger creates new message instance
func NewMessenger(l *log.Logger, port string) *Messenger {
	return &Messenger{
		l,
		port,
	}
}

// swagger:route GET /message message MessageScheme
// Returns the message
// responses:
//	200: messageResponse

// Handler func for the messenger [GET] controller
func (m *Messenger) Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")	
	
	data := models.Message {
		Message: "Hello, World",
		Port: m.port,
	}
	
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		m.l.Println(err.Error())
	}
}

