package sdk

// Credential is used to sign the request
type Credential struct {
	SecertID  string
	SecretKey string
}

func NewCredentials(secretID, secretKey string) *Credential {
	return &Credential{
		secretID,
		secretKey,
	}
}
