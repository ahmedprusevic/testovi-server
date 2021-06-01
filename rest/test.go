package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ahmedprusevic/testovi-server/models"
	"github.com/ahmedprusevic/testovi-server/responses"
	errorformatter "github.com/ahmedprusevic/testovi-server/utils"
	"github.com/gorilla/mux"
)

func (s *Server) CreateTest(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	test := models.Test{}

	err = json.Unmarshal(body, &test)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	test.FillFields()

	err = test.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	testCreated, err := test.SaveTest(s.DB)

	if err != nil {
		formattedError := errorformatter.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusCreated, testCreated)

}

func (s *Server) GetTests(w http.ResponseWriter, r *http.Request) {
	test := models.Test{}

	tests, err := test.FindAllTests(s.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, tests)
}

func (s *Server) GetTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	test := models.Test{}

	testRes, err := test.FindTestById(s.DB, uint32(tid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, testRes)
}

func (s *Server) UpdateTest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	tid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	test := models.Test{}

	err = json.Unmarshal(body, &test)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	test.FillFields()

	err = test.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedTest, err := test.UpdateTest(s.DB, uint32(tid))

	if err != nil {
		formattedError := errorformatter.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedTest)
}

func (s *Server) DeleteTest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	test := models.Test{}

	tid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = test.DeleteTest(s.DB, uint32(tid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", tid))

	responses.JSON(w, http.StatusNoContent, "")
}
