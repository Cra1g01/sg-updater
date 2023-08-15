package sys

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Cra1g01/sg-updater/pkg/config"
)

func GetIpAddr() string {
	ip, err := ipRequest(config.IpURL)
	if err != nil {
		log.Fatalf("Failed to get IP address, %v", err)
	}
	return ip
}

func ipRequest(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Error: status code %d", res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
        return "", err
	}
	sb := string(body)
    trimmedSb := strings.TrimSpace(sb)
	return trimmedSb, nil
}
