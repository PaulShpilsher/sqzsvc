package transition

import (
	"log"
	"sqzsvc/models"
)

func LogTransition(shortCode string, clientAddress string) {
	transiton := &models.Transition{
		ShortCode:     shortCode,
		ClientAddress: clientAddress,
	}

	if err := transiton.Save(); err != nil {
		log.Println("Failed to log transiton", err.Error())
	}
}
