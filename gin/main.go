package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bind struct {
	Name string `json:"name" binding:"required"`
}

type ExStruct struct {
	Name string `json:"name" binding:"required"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// post example
	r.POST("/examplePost", func(c *gin.Context) {

		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err.Error())
		}

		var data map[string]interface{}
		json.Unmarshal([]byte(value), &data)

		doc, _ := json.Marshal(data)
		c.String(http.StatusOK, string(doc))
	})

	// post to chaincode example
	r.POST("/testChaincode", func(c *gin.Context) {

		// if you want request instance you have to process extra.
		// var bodyBytes []byte
		// if c.Request.Body != nil {
		// 	bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		// }
		// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		body := c.Request.Body
		req, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println(err.Error())
		}

		res, err := ingress.Routing(req)
		if err != nil {
			doc, _ := json.Marshal(err)
			c.String(http.StatusOK, string(doc))
		} else {
			doc, _ := json.Marshal(res)
			c.String(http.StatusOK, string(doc))
		}
	})

	// post form example
	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)
	})

	// post bind example
	r.POST("/bind", func(c *gin.Context) {

		req := &Bind{}
		err := c.ShouldBind(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.Status(http.StatusOK)

		var requestBody ExStruct
		if err := c.BindJSON(&requestBody); err != nil {
			// DO SOMETHING WITH THE ERROR
		}

		fmt.Println(requestBody.Name)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
