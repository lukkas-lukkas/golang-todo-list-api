package encrypt

import "golang.org/x/crypto/bcrypt"

type BcryptService struct{}

func NewBcryptService() *BcryptService {
	return &BcryptService{}
}

func (s BcryptService) Encrypt(value string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil

}

func (s BcryptService) Check(encrypted string, string string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(string))
	return err == nil
}
