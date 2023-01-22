package services

import models "sqzsvc/models"

type ShortCodeService struct {
	Identity *models.Identity
}

func (s *ShortCodeService) RegisterLongUrl(longUrl string) (string, error) {

	// TODO: validate URL
	// TODO: normalize URL

	user := &models.User{}
	if _, err := user.GetUserById(s.Identity.UserId); err != nil {
		return "", err
	}

	userUrl := &models.UserUrl{
		UserId:  s.Identity.UserId,
		LongUrl: longUrl,
	}

	// TODO: check if this user already registered this URL

	// save / create
	if _, err := userUrl.Save(); err != nil {
		return "", err
	}

	return userUrl.ShortCode, nil
}
