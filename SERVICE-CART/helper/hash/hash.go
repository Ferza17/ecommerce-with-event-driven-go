package hash

import (
	"github.com/RoseRocket/xerrs"
	"golang.org/x/crypto/bcrypt"

	"github.com/Ferza17/event-driven-cart-service/utils"
)

func Hashed(password string) (response string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	if err != nil {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	response = string(hashed)
	return
}

func Compare(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
