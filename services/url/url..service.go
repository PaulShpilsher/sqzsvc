package url

import (
	"fmt"
	"sqzsvc/models"
)

func SubmitUrl(url string, clientAddress string) (string, error) {

	var err error = nil

	userUrl := &models.UrlData{
		ClientAddress: clientAddress,
		Url:           url,
	}

	if _, ok := userUrl.GetByUrl(url); !ok {
		err = userUrl.Save()
	}

	return userUrl.ShortCode, err
}

func GetUrl(shortCode string) (string, error) {

	userUrl := &models.UrlData{}

	if _, ok := userUrl.GetByShortCode(shortCode); !ok {
		return "", fmt.Errorf("short code '%s' not found", shortCode)
	}

	return userUrl.Url, nil
}
