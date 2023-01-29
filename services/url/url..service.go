package url

import (
	"fmt"
	"sqzsvc/models"
)

func SubmitLongUrl(longUrl string, clientAddress string) (string, error) {

	var err error = nil

	url := &models.Url{
		ClientAddress: clientAddress,
		LongUrl:       longUrl,
	}

	if _, ok := url.GetByLongUrl(longUrl); !ok {
		err = url.Save()
	}

	return url.ShortCode, err
}

func GetUrl(shortCode string) (string, error) {

	url := &models.Url{}

	if _, ok := url.GetByShortCode(shortCode); !ok {
		return "", fmt.Errorf("short code '%s' not found", shortCode)
	}

	return url.LongUrl, nil
}
