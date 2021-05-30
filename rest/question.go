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

func (s *Server) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	question := models.Question{}

	err = json.Unmarshal(body, &question)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	question.FillFields()

	err = question.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	questionCreated, err := question.SaveQuestion(s.DB)

	if err != nil {
		formattedError := errorformatter.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(w, http.StatusCreated, questionCreated)

}

func (s *Server) GetQuestions(w http.ResponseWriter, r *http.Request) {
	queston := models.Question{}

	questions, err := queston.FindAllQuestions(s.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, questions)
}

func (s *Server) GetQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	qid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	question := models.Question{}

	questionRes, err := question.FindQuestionById(s.DB, uint32(qid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, questionRes)
}

func (s *Server) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	qid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	question := models.Question{}

	err = json.Unmarshal(body, &question)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	question.FillFields()

	err = question.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedQuestion, err := question.UpdateQuestion(s.DB, uint32(qid))

	if err != nil {
		formattedError := errorformatter.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatedQuestion)

}

func (s *Server) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	question := models.Question{}

	qid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = question.DeleteQuestion(s.DB, uint32(qid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", qid))
	responses.JSON(w, http.StatusNoContent, "")
}
