package crypto

type CryptoService interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}

type Crypto struct {
	RSA RSAHelper
}

func New() Crypto {
	rsa := RSAHelper{}
	rsa.New()
	return Crypto{RSA: rsa}
}
