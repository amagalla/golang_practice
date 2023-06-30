package physicians

import (
	"golang_practice/cmd/main/controllers/appointments"
	"log"

	"github.com/gin-gonic/gin"
)

func InsertPhysicians (c *gin.Context) {
	var requestData struct {
		PhysicianFirstName string `json:"physician_first_name"`
		PhysicianLastName string `json:"physician_last_name"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		log.Println("Failed to parse request body: ", err)
		c.JSON(400, gin.H{
			"Error: ": "Failed to parse request body",
		})

		return
	}

	if err := appointments.InsertPhysiciansData(requestData.PhysicianFirstName, requestData.PhysicianLastName); err != nil {
		log.Println("Failed to insert physician data: ", err)

		c.JSON(400, gin.H {
			"Error": "Failed to insert physician data",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Physician data inserted successfully",
	})
}