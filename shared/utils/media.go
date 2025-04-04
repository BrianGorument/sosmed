package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

// IsValidURL checks if a string is a valid URL
func IsValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

// DecodeBase64ToFile decodes a base64 string and saves it as a file
func DecodeBase64ToFile(base64Str string) (string, error) {
	// Remove data URL prefix (data:image/png;base64, etc.)
	parts := strings.SplitN(base64Str, ",", 2)
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid base64 string")
	}

	// Decode base64
	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	// Save to file (here we save to a temp file)
	filePath := "media_" + fmt.Sprintf("%d", time.Now().Unix()) + ".jpg" // Can be dynamic based on media type
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
