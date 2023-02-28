package itest

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load environment file with %q", err)
	}
}

func GetAnnouncementsUrl() string {

	return fmt.Sprintf("%s/announcements", getEnvOrExit("BASE_URL"))
}

func getEnvOrExit(variable string) string {

	res := os.Getenv(variable)
	if res == "" {
		log.Fatalf("Please, set %q", variable)
	}
	log.Debugf("%q: %q", variable, res)

	return res
}
