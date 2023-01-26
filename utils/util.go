package utils

import (
	"errors"
	"fmt"
	"time"
)

func CalculateDuration(timeValue uint64, mode uint64) (time.Duration, error) {
	duration := time.Duration(timeValue)
	switch mode {
	case 1:
		duration *= time.Microsecond
	case 2:
		duration *= time.Millisecond
	case 3:
		duration *= time.Second
	case 4:
		duration *= time.Minute
	case 5:
		duration *= time.Hour
	default:
		return 0, errors.New("wrong mode")
	}
	return duration, nil
}

func GetDocumentation(generalDocumentation string, flagsDocumentation []string) string {
	documentation := fmt.Sprintf("%s.\n", generalDocumentation)
	for _, v := range flagsDocumentation {
		documentation += v
	}
	return documentation
}

func FormatFlagDocumentation(flagName string, documentationText string) string {
	return fmt.Sprintf("-%s: %s.\n", flagName, documentationText)
}

func FormatArgDocumentation(argumentName string, documentationText string) string {
	return fmt.Sprintf("%s: %s.\n", argumentName, documentationText)
}
