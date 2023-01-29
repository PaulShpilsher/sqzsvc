package url

import (
	"fmt"
	"sqzsvc/models"
)

func SubmitUrl(url string, clientAddress string) (string, error) {

	var err error = nil

	urlData := &models.UrlEntry{
		ClientAddress: clientAddress,
		Url:           url,
	}

	if _, ok := urlData.GetByUrl(url); !ok {
		err = urlData.Save()
	}

	return urlData.ShortCode, err
}

func GetUrl(shortCode string) (string, error) {

	userUrl := &models.UrlEntry{}

	if _, ok := userUrl.GetByShortCode(shortCode); !ok {
		return "", fmt.Errorf("short code '%s' not found", shortCode)
	}

	return userUrl.Url, nil
}
