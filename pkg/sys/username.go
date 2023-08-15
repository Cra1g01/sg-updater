package sys

import (
	"log"
	"os/user"
)

func GetUsername() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to fetch user, %v", err)
	}
	return user.Username
}
