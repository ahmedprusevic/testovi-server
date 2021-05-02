package rest

import (
	"net/http"

	"github.com/ahmedprusevic/testovi-server/responses"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to driver' testing API")
}
