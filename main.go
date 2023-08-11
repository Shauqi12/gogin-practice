package main

import "github.com/gin-gonic/gin"



func main() {
	app := gin.Default()

	


	app.GET("/", func(ctx *gin.Context) {

		isValidate := false

		if !isValidate {
			ctx.AbortWithStatusJSON(400, gin.H{
				"message": "Not valid",
			})	

			return
		}

		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	app.Run(": 8000")
}