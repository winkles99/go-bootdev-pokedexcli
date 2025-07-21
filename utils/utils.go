package utils

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func GetCmdFromPrompt(prompt string) ([]string, error) {
	if prompt == "" {
		return nil, fmt.Errorf("prompt cannot be empty")
	}
	return strings.Fields(strings.ToLower(prompt)), nil
}

func GetOffsetFromUrl(uri *string) int {
	if uri == nil || *uri == "" {
		return 0
	}
	u, err := url.Parse(*uri)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	offsetStr := q.Get("offset")

	if offsetStr == "" {
		return 0
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		fmt.Println("Error during conversion")
		return 0
	}

	return offset
}