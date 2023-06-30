package physicians

import (
	"golang_practice/cmd/main/controllers/appointments"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var requestData struct {
	PhysicianFirstName string `json:"physician_first_name"`
	PhysicianLastName string `json:"physician_last_name"`
}


func InsertPhysicians (c *gin.Context) {
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

type PhysicianListResponse struct {
	PhysicianList []appointments.Physicians `json:"results"`
}

func GetPhysicians(c *gin.Context) {
	physiciansList, err := appointments.GetPhysiciansList()
	if err != nil {
		// Handle the error, such as returning an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve physicians"})
		return
	}

	response := PhysicianListResponse{
		PhysicianList: physiciansList,
	}

	// Return the response as JSON
	c.JSON(http.StatusOK, response)
}
