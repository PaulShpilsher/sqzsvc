package services

import (
	"fmt"
	models "sqzsvc/models"
)

type ShortCodeService struct {
	Identity *models.Identity
}

func (s *ShortCodeService) RegisterLongUrl(longUrl string) (string, error) {

	// TODO: validate URL
	// TODO: normalize URL

	userUrl := &models.UserUrl{
		UserID:  s.Identity.UserID,
		LongUrl: longUrl,
	}

	var err error
	if _, ok := userUrl.GetByUserAndUrl(); !ok {
		err = userUrl.Save()
	}

	return userUrl.ShortCode, err
}

func (s *ShortCodeService) GetLongUrl(shortCode string) (string, error) {

	userUrl := &models.UserUrl{}
	if _, ok := userUrl.GetByShortCode(shortCode); !ok {
		return "", fmt.Errorf("short code '%s' not found", shortCode)
	}

	return userUrl.LongUrl, nil
}
