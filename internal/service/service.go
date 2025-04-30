package service

import (
	"strings"

	"github.com/AeddGynvael3110/sprint6/pkg/morse"
)

func DetectAndConvert(input string) (string, error) {
	if strings.Contains(input, ".") || strings.Contains(input, "-") {
		// Предполагаем, что это код Морзе
		return morse.ToText(input), nil
	} else {
		// Предполагаем, что это обычный текст
		return morse.ToMorse(input), nil
	}
}
