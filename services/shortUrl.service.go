package services

import models "sqzsvc/models"

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
