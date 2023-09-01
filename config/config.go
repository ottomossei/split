package split

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	splitType "split/splittype"
)

const (
	InitConfigB = ""
	InitConfigL = 1000
	InitConfigN = ""
)

type Config struct {
	B string
	L int
	N string
}

type SplitType = splitType.SplitType

func invalidBFlag(input string) bool {
	if input == "" {
		return false
	}
	matched, _ := regexp.MatchString(`^\d+[mMkKGTP]?[B]?$`, input)
	return !matched
}

func (config *Config) validate() error {
	fmt.Println(config.B, config.L, config.N)
	switch {
	case config.B != InitConfigB && config.L != InitConfigL && config.N == InitConfigN:
		fallthrough
	case config.B == InitConfigB && config.L != InitConfigL && config.N != InitConfigN:
		fallthrough
	case config.B != InitConfigB && config.L == InitConfigL && config.N != InitConfigN:
		fallthrough
	case config.B != InitConfigB && config.L != InitConfigL && config.N != InitConfigN:
		return errors.New("cannot split in more than one way")
	case config.L <= 0:
		return errors.New("invalid number of lines: ‘" + fmt.Sprint(config.L) + "’")
	case invalidBFlag(config.B):
		return errors.New("invalid number of lines: ‘" + fmt.Sprint(config.B) + "’")
	default:
		return nil
	}
}

func ParseFlags() (*Config, error) {
	config := &Config{}
	flag.StringVar(&config.B, "b", InitConfigB, "-b flag")
	flag.IntVar(&config.L, "l", InitConfigL, "-l flag")
	flag.StringVar(&config.N, "n", InitConfigN, "-n flag")
	flag.Parse()
	return config, config.validate()
}

func (config *Config) ConvertSplitType() (SplitType, error) {
	var result SplitType
	var err error
	switch {
	case config.N != InitConfigN:
		outputLine, value, e := splitType.ExtractNumbers(config.N)
		err = e
		result = SplitType{ThresholdType: "lineEquallyDivided", Value: value, OutputLine: outputLine}
	case config.B != InitConfigB:
		flagValue, e := splitType.ParseByteSize(config.B)
		err = e
		result = SplitType{ThresholdType: "byte", Value: flagValue, OutputLine: 0}
	default:
		result = SplitType{ThresholdType: "line", Value: config.L, OutputLine: 0}
	}
	return result, err
}
