package url

import (
	"fmt"
	"sqzsvc/models"
)

func SubmitUrl(url string, clientAddress string) (string, error) {

	var err error = nil

	urlData := &models.Url{
		ClientAddress: clientAddress,
		LongUrl:       url,
	}

	if _, ok := urlData.GetByLongUrl(url); !ok {
		err = urlData.Save()
	}

	return urlData.ShortCode, err
}

func GetUrl(shortCode string) (string, error) {

	userUrl := &models.Url{}

	if _, ok := userUrl.GetByShortCode(shortCode); !ok {
		return "", fmt.Errorf("short code '%s' not found", shortCode)
	}

	return userUrl.LongUrl, nil
}
