package routes

import (
	"github.com/julienschmidt/httprouter"
	"jameesjohn.com/hngStageTwo/src/controllers"
	"net/http"
)

func Router() http.Handler {
	router := httprouter.New()

	router.GET("/api/persons", controllers.GetAllPersons)
	router.POST("/api/persons", controllers.CreatePerson)
	router.GET("/api/persons/:personId", controllers.GetPerson)
	router.PUT("/api/persons/:personId", controllers.UpdatePerson)
	router.DELETE("/api/persons/:personId", controllers.DeletePerson)

	return router
}
