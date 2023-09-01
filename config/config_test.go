package split

import (
	"testing"
)

func TestInvalidBFlag(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"123B", false},
		{"123kB", false},
		{"123mB", false},
		{"123M", false},
		{"123kG", true},
		{"ABC", true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := invalidBFlag(test.input)
			if actual != test.expected {
				t.Errorf("For input %q, expected %v, but got %v", test.input, test.expected, actual)
			}
		})
	}
}

func TestConfigValidation(t *testing.T) {
	testCases := []struct {
		desc   string
		config Config
		errMsg string
	}{
		{
			desc: "Valid Config",
			config: Config{
				B: "",
				L: 1000,
				N: "",
			},
			errMsg: "",
		},
		{
			desc: "Invalid L Value",
			config: Config{
				B: "",
				L: -1,
				N: "",
			},
			errMsg: "invalid number of lines: ‘-1’",
		},
		{
			desc: "Invalid L Value",
			config: Config{
				B: "10AA",
				L: 1000,
				N: "",
			},
			errMsg: "invalid number of lines: ‘10AA’",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.config.validate()
			if tc.errMsg != "" {
				if err == nil || err.Error() != tc.errMsg {
					t.Errorf("Expected error message: '%s', but got: '%v'", tc.errMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: '%v'", err)
				}
			}
		})
	}
}

func TestConfig_ConvertSplitType(t *testing.T) {
	tests := []struct {
		name     string
		config   Config
		expected SplitType
		err      error
	}{
		{
			name: "Test Case 1",
			config: Config{
				B: "100",
			},
			expected: SplitType{
				ThresholdType: "byte",
				Value:         100,
				OutputLine:    0,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.config.ConvertSplitType()

			if err != test.err {
				t.Errorf("Expected error: %v, but got: %v", test.err, err)
			}

			if actual != test.expected {
				t.Errorf("Expected: %+v, but got: %+v", test.expected, actual)
			}
		})
	}
}
