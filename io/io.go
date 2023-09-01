package split

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	splitType "split/splittype"
	"strings"
)

const (
	alphabet             = "abcdefghijklmnopqrstuvwxyz"
	carryForwardAlphabet = len(alphabet)
)

func decimalToAlphabet(decimal, maxDigit int) string {
	result := ""
	for decimal > 0 {
		remainder := decimal % carryForwardAlphabet
		result = string(alphabet[remainder]) + result
		decimal /= carryForwardAlphabet
	}

	for len(result) < maxDigit {
		result = "a" + result
	}

	return result
}

func createNewOutputFile(fileNamePrefix string, decimal, alphabetMaxDigit int) (io.Writer, error) {
	outputFilePath := fmt.Sprintf("%s%s", fileNamePrefix, decimalToAlphabet(decimal, alphabetMaxDigit))
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return nil, err
	}
	return outputFile, nil
}

func SplitDataToFiles(splitType splitType.SplitType) error {
	switch {
	case strings.Contains(splitType.ThresholdType, "line"):
		return SplitDataToFilesByLine(splitType)
	case splitType.ThresholdType == "byte":
		return SplitDataToFilesByByte(splitType)
	default:
		return errors.New("unsupported split type")
	}
}

func SplitDataToFilesByLine(splitType splitType.SplitType) error {
	data, err := os.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	lines := strings.Split(string(data), "\n")

	var thresholdPerFile int
	switch splitType.ThresholdType {
	case "lineEquallyDivided":
		thresholdPerFile = len(lines) / splitType.Value
	case "line":
		if splitType.Value > 0 && splitType.Value <= len(lines) {
			thresholdPerFile = splitType.Value
		} else {
			return errors.New("invalid threshold value")
		}
	}

	alphabetMaxDigit := 2
	fileNamePrefix := "x"

	outputFileCounter := 0
	decimalCounter := 0

	outputFile, err := createNewOutputFile(fileNamePrefix, decimalCounter, alphabetMaxDigit)
	if err != nil {
		return err
	}

	writeCounter := 0
	for _, line := range lines {
		fmt.Fprintln(outputFile, line)
		writeCounter++
		if writeCounter == thresholdPerFile {
			writeCounter = 0
			if outputFileCounter == int(math.Pow(float64(carryForwardAlphabet), float64(alphabetMaxDigit))-1) {
				decimalCounter = 0
				alphabetMaxDigit++
				fileNamePrefix += "z"
			} else {
				decimalCounter++
			}
			if outputFile, err = createNewOutputFile(fileNamePrefix, decimalCounter, alphabetMaxDigit); err != nil {
				return err
			}
		}
	}

	return nil
}

func SplitDataToFilesByByte(splitType splitType.SplitType) error {
	data, err := os.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	var thresholdPerFile int

	if splitType.Value > 0 && splitType.Value <= len(data) {
		thresholdPerFile = splitType.Value
	} else {
		return fmt.Errorf("invalid threshold value")
	}

	alphabetMaxDigit := 2
	fileNamePrefix := "x"

	outputFileCounter := 0
	decimalCounter := 0

	outputFile, err := createNewOutputFile(fileNamePrefix, decimalCounter, alphabetMaxDigit)
	if err != nil {
		return err
	}

	writeCounter := 0
	for _, b := range data {
		fmt.Fprintf(outputFile, "%c", b)
		writeCounter++
		if writeCounter == thresholdPerFile {
			writeCounter = 0
			if outputFileCounter == int(math.Pow(float64(carryForwardAlphabet), float64(alphabetMaxDigit))-1) {
				decimalCounter = 0
				alphabetMaxDigit++
				fileNamePrefix += "z"
			} else {
				decimalCounter++
			}
			if outputFile, err = createNewOutputFile(fileNamePrefix, decimalCounter, alphabetMaxDigit); err != nil {
				return err
			}
		}
	}

	return nil
}
