package main

import (
	"horas-golang/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eg := r.Group("/v1")
	{
		eg.GET("/horas", getHoras)
		eg.POST("/horas", getHorasComEducacao)
	}

	r.Run(":8080")
}

func getHoras(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensagem": time.Now().Format("15:00:00"),
	})
}

func getHorasComEducacao(c *gin.Context) {
	var horasReq models.HorasReq

	if err := c.ShouldBindJSON(&horasReq); err != nil {
		c.JSON(400, gin.H{"error": "mensagem mal formatada"})
	}

	if horasReq.Mensagem != "por_favor" {
		c.JSON(400, gin.H{"error": "Palavra mágica?"})
	} else {
		c.JSON(200, gin.H{
			"hora":     time.Now().Format("15:00:00"),
			"mensagem": "Aqui está a hora, como solicitado",
		})
	}
}
