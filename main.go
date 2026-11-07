package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

const addr = "localhost:8080"

type Vinil struct{
	ID 		string  `json: "id"`
	Title 	string  `json: "title"`
	Artist 	string  `json: "artist"`
	Price	float64 `json: "price"`
}

var vinils = []Vinil{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
	router := gin.Default()
	router.GET("/vinils", getVinil)
	router.GET("/vinils/:id", getVinilbyID)
	router.POST("/vinils", postVinil)

	router.Run(addr)
}

func getVinil(c *gin.Context){
	c.IndentedJSON(http.StatusOK, vinils)
}

func postVinil(c *gin.Context){
	var newVinil Vinil;

	if err := c.BindJSON(&newVinil); err != nil{
		return;
	}

	vinils = append(vinils, newVinil)
	c.IndentedJSON(http.StatusCreated, newVinil)
}

func getVinilbyID(c *gin.Context){
	id := c.Param("id")

	for _, v := range vinils{
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return;
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vinil doesn't find"})
}