package main

//

import (
	_ "encoding/json"
	_ "net/http"

	"./route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//routes define
	hello := route.Hello
	api1 := route.Api1
	param := route.GetParam
	que := route.QueryStr

	//route
	r.GET("/HELLO", hello)     //default get
	r.GET("param:name", param) //get with param

	r.GET("Query", que)

	api := r.Group("/api")
	{
		api.GET("/api1", api1)

	}

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
