package common

import "github.com/google/uuid"

func GenerateUUID() (newUUID string, er error) {
	uuid, er := uuid.NewRandom()
	if er != nil {
		return "", er
	}
	newUUID = uuid.String()
	return
}
