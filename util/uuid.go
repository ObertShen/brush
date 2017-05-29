package util

import (
	"log"

	u "github.com/nu7hatch/gouuid"
)

// GetUUID 生成并返回 v4 UUID
func GetUUID() string {
	uuid, err := u.NewV4()
	if err != nil {
		log.Panic(err)
	}
	return uuid.String()
}
