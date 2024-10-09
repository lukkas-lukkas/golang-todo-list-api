package helpers

import "golang.org/x/crypto/bcrypt"

type BcryptEncryptor struct{}

func NewBcryptEncryptor() *BcryptEncryptor {
	return &BcryptEncryptor{}
}

func (e BcryptEncryptor) Encrypt(value string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil

}

func (e BcryptEncryptor) Check(encrypted string, string string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(string))
	return err == nil
}
