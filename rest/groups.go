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

func (s *Server) CreateGroup(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	group := models.Group{}

	err = json.Unmarshal(body, &group)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	group.FillFields()

	err = group.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	groupCreated, err := group.SaveGroup(s.DB)

	if err != nil {
		formattedError := errorformatter.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
	}

	responses.JSON(w, http.StatusCreated, groupCreated)
}

func (s *Server) GetGroups(w http.ResponseWriter, r *http.Request) {
	group := models.Group{}

	groups, err := group.FindAllGroups(s.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, groups)
}

func (s *Server) GetGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	gid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	group := models.Group{}

	groupRes, err := group.FindGroupById(s.DB, uint32(gid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, groupRes)
}

func (s *Server) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	gid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	group := models.Group{}

	err = json.Unmarshal(body, &group)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	group.FillFields()

	err = group.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedGroup, err := group.UpdateGroup(s.DB, uint32(gid))

	if err != nil {
		formattedError := errorformatter.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatedGroup)

}

func (s *Server) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	group := models.Group{}

	gid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = group.DeleteGroup(s.DB, uint32(gid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", gid))
	responses.JSON(w, http.StatusNoContent, "")
}
