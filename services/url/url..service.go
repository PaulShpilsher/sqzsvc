package url

import (
	"fmt"
	models "sqzsvc/models"
)

func RegisterLongUrl(identity *models.Identity, longUrl string) (string, error) {

	// TODO: validate URL
	// TODO: normalize URL

	userUrl := &models.UserUrl{
		UserID:  identity.UserID,
		LongUrl: longUrl,
	}

	var err error
	if _, ok := userUrl.GetByUserAndUrl(); !ok {
		err = userUrl.Save()
	}

	return userUrl.ShortCode, err
}

func GetLongUrl(shortCode string) (string, error) {

	userUrl := &models.UserUrl{}
	if _, ok := userUrl.GetByShortCode(shortCode); !ok {
		return "", fmt.Errorf("short code '%s' not found", shortCode)
	}

	return userUrl.LongUrl, nil
}
