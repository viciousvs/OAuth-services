package getUUID

import "github.com/viciousvs/OAuth-services/user/utils/hashPassword"

//getUUIDDTO
type getUUIDDTO struct {
	Login        string
	PasswordHash string
}

//makeGetUUIDDTO
func makeGetUUIDDTO(login, password string) getUUIDDTO {
	return getUUIDDTO{
		Login:        login,
		PasswordHash: hashPassword.Hash(password),
	}
}
