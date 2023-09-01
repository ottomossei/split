package split

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestDecimalToAlphabet(t *testing.T) {
	tests := []struct {
		decimal          int
		maxDigit         int
		expectedAlphabet string
	}{
		{1, 2, "ab"},
		{25, 2, "az"},
		{26, 2, "ba"},
		{27, 2, "bb"},
	}

	for _, test := range tests {
		result := decimalToAlphabet(test.decimal, test.maxDigit)
		if result != test.expectedAlphabet {
			t.Errorf("For decimal %d and maxDigit %d, expected %s but got %s", test.decimal, test.maxDigit, test.expectedAlphabet, result)
		}
	}
}

func TestCreateNewOutputFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_output")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		fileNamePrefix   string
		decimal          int
		alphabetMaxDigit int
		expectedContent  string
	}{
		{"test_file_", 1, 2, "test_file_ab"},
		{"test_file_", 25, 2, "test_file_az"},
	}

	for _, test := range tests {
		outputFile, err := createNewOutputFile(tempDir+string(filepath.Separator)+test.fileNamePrefix, test.decimal, test.alphabetMaxDigit)
		if err != nil {
			t.Errorf("Error creating output file: %v", err)
			continue
		}
		defer outputFile.(io.Closer).Close()

		expectedFilePath := filepath.Join(tempDir, test.expectedContent)
		if outputFile.(*os.File).Name() != expectedFilePath {
			t.Errorf("For fileNamePrefix %s, decimal %d, and alphabetMaxDigit %d, expected file path %s but got %s",
				test.fileNamePrefix, test.decimal, test.alphabetMaxDigit, expectedFilePath, outputFile.(*os.File).Name())
		}
	}
}
