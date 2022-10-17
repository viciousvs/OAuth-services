package argon2

var chArgon = make(chan struct{}, 2)

type argonHasher struct {
	params *Params
}

func NewArgonHasher() argonHasher {
	params := DefaultParams
	return argonHasher{params: params}
}
func (h argonHasher) GenerateHash(password string) (string, error) {
	chArgon <- struct{}{}
	defer func() {
		<-chArgon
	}()
	hash, err := createHash(password, h.params)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (h argonHasher) CompareHashAndPassword(passwordHash, password string) (bool, error) {
	chArgon <- struct{}{}
	defer func() { <-chArgon }()
	return comparePasswordAndHash(password, passwordHash)
}
