package rest

import (
	"io/ioutil"
	"net/http"

	"github.com/ahmedprusevic/testovi-server/models"
	"github.com/ahmedprusevic/testovi-server/responses"
)

func (s *Server) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	question := models.Question{}

}
