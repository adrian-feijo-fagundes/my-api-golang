package api

import (
	"fmt"
	"net/http"

	"github.com/adrian-feijo-fagundes/my-api-golang/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{Echo: e, DB: studentDB}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

// Handler
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}
func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(student); err != nil { // FORMA "simplificada de tratar um erro"
		return c.String(http.StatusInternalServerError, "Error to create students\n")
	}

	return c.String(http.StatusOK, "Create students\n")
}
func (api *API) getStudent(c echo.Context) error {
	id := c.Param("id") // Pega o parametro que foi passado
	message := "GET " + getId(id)
	return c.String(http.StatusOK, message)
}
func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	message := "UPDATE " + getId(id)
	return c.String(http.StatusOK, message)
}
func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	message := "DELETE " + getId(id)
	return c.String(http.StatusOK, message)
}

func getId(id string) string {
	return fmt.Sprintf("%s student \n", id)
}
