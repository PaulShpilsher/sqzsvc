package services

import models "sqzsvc/models"

type ShortCodeService struct {
	Identity *models.Identity
}

func (s *ShortCodeService) RegisterLongUrl(longUrl string) error {

	// TODO: validate URL
	user := &models.User{}
	if _, err := user.GetUserById(s.Identity.UserId); err != nil {
		return err
	}

	userUrl := &models.UserUrl{
		UserId:    s.Identity.UserId,
		LongUrl:   longUrl,
		ShortCode: "aaa",
	}

	if _, err := userUrl.SaveUserUrl(); err != nil {
		return err
	}

	return nil
}

// func newShortCode() {

// }
