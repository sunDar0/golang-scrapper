package common

import (
	"log"
	"net/http"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CheckCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status:", res.StatusCode)
	}
}

// CleanString string sanitize
func CleanString(txt string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(txt)), " ")
}
