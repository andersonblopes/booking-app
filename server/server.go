package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

func main() {
	var g = singleflight.Group{}

	r := gin.Default()

	r.GET("/api/v1/get_something", func(c *gin.Context) {
		name := c.DefaultQuery("name", "")
		response, _, _ := g.Do(name, func() (interface{}, error) {
			result := processingRequest(name)
			return result, nil
		})
		fmt.Fprint(c.Writer, response)
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func processingRequest(name string) string {
	fmt.Println("[DEBUG] processing request..")
	return "Hi there! You requested " + name
}
