package helpers

import (
	"strings"
)

func IsValidImageExtension(imageStr string) bool {
	validExtensions := []string{".jpg", ".jpeg", ".png", ".svg"}
	for _, ext := range validExtensions {
		if strings.HasSuffix(strings.ToLower(imageStr), ext) {
			return true
		}
	}
	return false
}
