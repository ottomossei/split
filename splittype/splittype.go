package split

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type SplitType struct {
	ThresholdType string
	Value         int
	OutputLine    int
}

var (
	byteUnitMap = map[string]int{
		"k": 1, "K": 1,
		"m": 2, "M": 2,
		"G": 3,
		"T": 4,
		"P": 5,
	}
)

func ParseByteSize(input string) (int, error) {
	re := regexp.MustCompile(`^(\d+)([mMkKGTP])?B?$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) < 2 {
		return 0, fmt.Errorf("invalid input: %s", input)
	}

	size, err := strconv.Atoi(matches[1])

	if err != nil {
		return 0, fmt.Errorf("invalid size: %s", matches[1])
	}
	unit := matches[2]
	exponent, found := byteUnitMap[unit]

	if found {
		size *= int(math.Pow(1024, float64(exponent)))
	}

	return size, nil
}

func ExtractNumbers(input string) (int, int, error) {
	if onlyNumber, err := parseOnlyNumber(input); err == nil {
		return 0, onlyNumber, nil
	}

	if firstNumber, secondNumber, err := parseRange(input); err == nil {
		return firstNumber, secondNumber, nil
	}

	return 0, 0, errors.New("extractNumbers: error")
}

func parseOnlyNumber(input string) (int, error) {
	numericPattern := "^[0-9]+$"
	if matched, err := regexp.MatchString(numericPattern, input); err == nil && matched {
		return strconv.Atoi(input)
	}
	return 0, fmt.Errorf("invalid format: %s", input)
}

func parseRange(input string) (int, int, error) {
	pattern := `(\d+)/(\d+)`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(input)

	if len(match) >= 3 {
		firstNumber, err1 := strconv.Atoi(match[1])
		secondNumber, err2 := strconv.Atoi(match[2])
		if err1 == nil && err2 == nil && firstNumber <= secondNumber {
			return firstNumber, secondNumber, nil
		}
	}
	return 0, 0, fmt.Errorf("invalid format: %s", input)
}
