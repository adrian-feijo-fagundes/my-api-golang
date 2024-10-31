package api

import (
	"fmt"
	"net/http"

	"github.com/adrian-feijo-fagundes/my-api-golang/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{Echo: e, DB: db}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/students", createStudent)
	api.Echo.GET("/students/:id", getStudent)
	api.Echo.PUT("/students/:id", updateStudent)
	api.Echo.DELETE("/students/:id", deleteStudent)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

// Handler
func getStudents(c echo.Context) error {
	students, err := db.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}
func createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := db.AddStudent(student); err != nil { // FORMA "simplificada de tratar um erro"
		return c.String(http.StatusInternalServerError, "Error to create students\n")
	}

	return c.String(http.StatusOK, "Create students\n")
}
func getStudent(c echo.Context) error {
	id := c.Param("id") // Pega o parametro que foi passado
	message := "GET " + getId(id)
	return c.String(http.StatusOK, message)
}
func updateStudent(c echo.Context) error {
	id := c.Param("id")
	message := "UPDATE " + getId(id)
	return c.String(http.StatusOK, message)
}
func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	message := "DELETE " + getId(id)
	return c.String(http.StatusOK, message)
}

func getId(id string) string {
	return fmt.Sprintf("%s student \n", id)
}
