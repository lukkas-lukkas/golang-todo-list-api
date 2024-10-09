package helpers

type Encryptor interface {
	Encrypt(value string) (string, error)
	Check(encrypted string, string string) bool
}
