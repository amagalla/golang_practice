package physicians

import (
	"golang_practice/cmd/main/controllers/appointments"
	"golang_practice/cmd/main/structs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertPhysicians (c *gin.Context) {
	var requestData structs.InsertPhysiciansRequestData
	
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

func GetPhysicians(c *gin.Context) {
	physiciansList, err := appointments.GetPhysiciansList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve physicians"})
		return
	}

	response := structs.PhysicianListResponse{
		PhysicianList: physiciansList,
	}

	c.JSON(http.StatusOK, response)
}
