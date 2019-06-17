package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	// "github.com/labstack/echo"
)


type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func Swervin(c echo.Context) error {
	return c.String(http.StatusOK, "Yallo Wurld")
}
func getCats(c echo.Context) error {

	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat's name is: %s\nand his type is: %s", catName, catType))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "did you want json or string data?",
	})
	
}

func getDogs(c echo.Context) error {
	dogName := c.QueryParam("name")
	dogType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("This dog %s is a good boy \nand it's type is %s", dogName, dogType))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": dogName,
			"type": dogType,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "did you mean json or string data?",
	})
}

func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("Failed reading the request body for addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &cat)
	if err != nil{
	log.Printf("Failed unmarshaling in addCats: %s", err)
	return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is your cat: %#v", cat)
	return c.String(http.StatusOK, "wee got your cat!!")
}
func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
		if err != nil{
		log.Printf("Failed processing addDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "wee got your dog!!")
}

func main() {
	fmt.Println("Sup Earth")

	e := echo.New()

	e.GET("/", Swervin)
	e.GET("/cats/:data", getCats)
	e.GET("/dogs/:data", getDogs)

	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)

	e.Start(":8000")
}
