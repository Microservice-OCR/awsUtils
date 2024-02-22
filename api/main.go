package main

import (
	"awsutils/awsconnect"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Main() {

	AWS_PORT, ok := os.LookupEnv("AWS_PORT")
	if !ok {
		log.Fatal("saving port not found")
	}

	router := gin.Default()
	router.Use(gin.Logger())

	router.POST("/upload", func(c *gin.Context) {
		// Récupération du fichier image de la requête
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		trueName, err := awsconnect.Put(header)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Image uploaded and path saved successfully", "name": trueName})
	})

	router.GET("/image/:filename", func(c *gin.Context) {
		// Récupération du fichier image de la requête
		filename := c.Param("filename")
		
		file, err := awsconnect.Get(filename)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.Data(200,"image/png",file)
	})

	router.Run(fmt.Sprintf(":%s", AWS_PORT))
}
