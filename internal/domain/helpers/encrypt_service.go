package helpers

type EncryptService interface {
	Encrypt(value string) (string, error)
	Check(encrypted string, string string) bool
}
