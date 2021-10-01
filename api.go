package main

import (
	"net/http"
	"ptz/onvif"

	"github.com/gin-gonic/gin"
)

type Input struct {
	X float64 `form:"x" json:"x" binding:"required"`
	Y float64 `form:"y" json:"y" binding:"required"`
	Z float64 `form:"z" json:"z" default:0 binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/api/ptz/controll", func(c *gin.Context) {
		var json Input
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		onvif.ControllPTZ(json.X, json.Y, json.Z, "192.168.100.64", "admin", "ftidev@123")

		c.JSON(http.StatusOK, gin.H{
			"status":  "PTZ activate",
			"VectorX": json.X,
			"VectorY": json.Y,
			"VectorZ": json.Z,
		})
	})
	r.Run(":3000")
}
