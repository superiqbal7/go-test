package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type Numbers struct {
	Number1   float32           `bson:"number1"`
	Number2   float32           `bson:"number2"`
	Operator  string 				`bson:"operator"`
}

func calculate(c *gin.Context) {

	reqBody := Numbers{}
	err := c.Bind(&reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid input data"})
    return
	}

	var result float32

	fmt.Println(reqBody.Number1);
	fmt.Println(reqBody.Number2);
	fmt.Println(reqBody.Operator)

	switch reqBody.Operator {
		case "*":
			result = reqBody.Number1 * reqBody.Number2
		case "/":
			if reqBody.Number2 == 0.0 {
				c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid input data"})
        return
			}
			result =  reqBody.Number1 / reqBody.Number2
		case "+":
			result =  reqBody.Number1 + reqBody.Number2
		case "-":
			result =  reqBody.Number1 - reqBody.Number2
	}

	c.IndentedJSON(http.StatusOK, result)
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length"},
			// AllowCredentials: true,
			// AllowOriginFunc: func(origin string) bool {
			// 		return origin == "https://github.com"
			// },
			// MaxAge: 12 * time.Hour,
	}))
  r.POST("/calculate", calculate)
  r.Run()
}
