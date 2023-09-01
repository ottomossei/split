package split

import (
	"testing"
)

func TestParseByteSize(t *testing.T) {
	const unit int = 1024
	testCases := []struct {
		input    string
		expected int
		wantErr  bool
	}{
		{"1000", 1000, false},
		{"2k", 2 * unit, false},
		{"5M", 5 * unit * unit, false},
		{"1G", 1 * unit * unit * unit, false},
		{"10T", 10 * unit * unit * unit * unit, false},
		{"15P", 15 * unit * unit * unit * unit * unit, false},
		{"aaaa", 0, true},
		{"", 0, true},
		{"1KB", unit, false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := ParseByteSize(tc.input)
			if err != nil && !tc.wantErr {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, result)
			}
		})
	}
}

func TestExtractNumbers(t *testing.T) {
	testCases := []struct {
		input          string
		expectedFirst  int
		expectedSecond int
		wantErr        bool
	}{
		{"123", 0, 123, false},
		{"5/10", 5, 10, false},
		{"abc", 0, 0, true},
		{"12xyz", 0, 0, true},
		{"9876.", 0, 0, true},
		{"invalid", 0, 0, true},
		{"1KB", 0, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			first, second, err := ExtractNumbers(tc.input)
			if err != nil && !tc.wantErr {
				t.Errorf("Unexpected error: %v", err)
			}
			if first != tc.expectedFirst || second != tc.expectedSecond {
				t.Errorf("Expected: (%d, %d), Got: (%d, %d)", tc.expectedFirst, tc.expectedSecond, first, second)
			}
		})
	}
}

func TestParseOnlyNumber(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
		wantErr  bool
	}{
		{"123", 123, false},
		{"456789", 456789, false},
		{"abc", 0, true},
		{"12xyz", 0, true},
		{"9876.", 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := parseOnlyNumber(tc.input)
			if err != nil && !tc.wantErr {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, result)
			}
		})
	}
}

func TestParseRange(t *testing.T) {
	testCases := []struct {
		input          string
		expectedFirst  int
		expectedSecond int
		wantErr        bool
	}{
		{"1/5", 1, 5, false},
		{"10/20", 10, 20, false},
		{"5/3", 0, 0, true},
		{"invalid", 0, 0, true},
		{"abc/def", 0, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			first, second, err := parseRange(tc.input)
			if err != nil && !tc.wantErr {
				t.Errorf("Unexpected error: %v", err)
			}
			if first != tc.expectedFirst || second != tc.expectedSecond {
				t.Errorf("Expected: (%d, %d), Got: (%d, %d)", tc.expectedFirst, tc.expectedSecond, first, second)
			}
		})
	}
}
