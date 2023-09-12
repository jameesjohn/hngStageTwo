package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
	"io"
	"jameesjohn.com/hngStageTwo/src/database"
	"jameesjohn.com/hngStageTwo/src/models"
	"jameesjohn.com/hngStageTwo/src/utils"
	"net/http"
	"strconv"
	"strings"
)

func getModel() models.PersonModel {
	pModel := models.PersonModel{Db: database.Db}

	return pModel
}

func CreatePerson(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type createRequest struct {
		Name string `json:"name"`
	}

	pModel := getModel()

	response, err := io.ReadAll(r.Body)
	var body createRequest
	err = json.Unmarshal(response, &body)
	if err != nil {
		utils.Fail(w, http.StatusUnprocessableEntity, map[string]interface{}{
			"name": "Invalid data supplied. The name field should be a string",
		})

		return
	}
	if strings.TrimSpace(body.Name) == "" {
		utils.Fail(w, http.StatusUnprocessableEntity, map[string]interface{}{
			"name": "Name field is required",
		})

		return
	}

	person, err := pModel.Create(models.Person{Name: body.Name})
	if err != nil {
		utils.Fail(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

	utils.Success(w, http.StatusCreated, map[string]interface{}{
		"person": person,
	})

}

func GetPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("personId")

	personId, err := strconv.ParseUint(id, 10, 64)

	fmt.Println(personId)
	if err != nil {
		utils.Fail(w, http.StatusNotFound, map[string]interface{}{
			"error": "Person not found",
		})
		return
	}

	pModel := getModel()
	person, err := pModel.Find(personId)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Fail(w, http.StatusNotFound, map[string]interface{}{
				"error": "Person not found",
			})
			return
		}

		utils.Fail(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

	utils.Success(w, http.StatusOK, map[string]interface{}{
		"person": person,
	})

}

func UpdatePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	type updateRequest struct {
		Name string `json:"name"`
	}

	id := p.ByName("personId")
	personId, err := strconv.ParseUint(id, 10, 64)

	fmt.Println(personId)
	if err != nil {
		utils.Fail(w, http.StatusNotFound, map[string]interface{}{
			"error": "Person not found",
		})
		return
	}

	response, err := io.ReadAll(r.Body)
	var body updateRequest
	err = json.Unmarshal(response, &body)

	// Validate request
	if err != nil {
		utils.Fail(w, http.StatusUnprocessableEntity, map[string]interface{}{
			"name": "Invalid data supplied. The name field should be a string",
		})

		return
	}
	if strings.TrimSpace(body.Name) == "" {
		utils.Fail(w, http.StatusUnprocessableEntity, map[string]interface{}{
			"name": "Name field is required",
		})

		return
	}

	toUpdate := models.Person{Id: personId, Name: body.Name}

	pModel := getModel()
	updatedPerson, err := pModel.Update(toUpdate)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Fail(w, http.StatusNotFound, map[string]interface{}{
				"error": "Person not found",
			})
			return
		}

		utils.Fail(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

	utils.Success(w, http.StatusOK, map[string]interface{}{
		"person": updatedPerson,
	})
}

func DeletePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("personId")
	personId, err := strconv.ParseUint(id, 10, 64)

	fmt.Println(personId)
	if err != nil {
		utils.Fail(w, http.StatusNotFound, map[string]interface{}{
			"error": "Person not found",
		})
		return
	}

	pModel := getModel()
	err = pModel.Delete(personId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Fail(w, http.StatusNotFound, map[string]interface{}{
				"error": "Person not found",
			})
			return
		}

		utils.Fail(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

	utils.Success(w, http.StatusNoContent, map[string]interface{}{})
}

func GetAllPersons(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pModel := getModel()
	persons, err := pModel.GetAll()
	if err != nil {
		utils.Fail(w, http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("Unable to get persons: %s", err.Error()),
		})
	}

	utils.Success(w, http.StatusOK, map[string]interface{}{
		"persons": persons,
	})
}
