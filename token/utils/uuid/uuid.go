package uuid

import "github.com/google/uuid"

// GenerateUUID create uuid with type string or panic
func GenerateUUID() string {
	return uuid.New().String()
}

//IsValidUUID
func IsValidUUID(uuidInput string) bool {
	_, err := uuid.Parse(uuidInput)
	return err == nil
}
