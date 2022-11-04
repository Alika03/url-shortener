package utils

import uuid "github.com/satori/go.uuid"

func GenerateUuid() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

func CheckUuid(ids ...string) error {
	for _, id := range ids {
		_, err := uuid.FromString(id)
		if err != nil {
			return err
		}
	}

	return nil
}
