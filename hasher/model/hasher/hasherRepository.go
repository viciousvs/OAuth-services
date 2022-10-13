package hasher

type Repository interface {
	GenerateHash(password string) (string, error)
	CompareHashAndPassword(passwordHash, password string) error
}
