package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ahmedprusevic/testovi-server/auth"
	"github.com/ahmedprusevic/testovi-server/models"
	"github.com/ahmedprusevic/testovi-server/responses"
	errorformatter "github.com/ahmedprusevic/testovi-server/utils"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.FillFields()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := s.UsersToken(user.Email, user.Password)

	if err != nil {
		formattedError := errorformatter.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)

}

func (s *Server) UsersToken(email, password string) (string, error) {
	var err error

	user := models.User{}
	err = s.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(int32(user.ID))

}
