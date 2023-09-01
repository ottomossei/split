package main

import (
	"fmt"
	"os"
	splitConfig "split/config"
	splitIo "split/io"
)

func main() {
	config, err := splitConfig.ParseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", err)
		os.Exit(1)
	}

	// for debug
	fmt.Println("-b:", config.B)
	fmt.Println("-l:", config.L)
	fmt.Println("-n:", config.N)

	splitType, err := config.ConvertSplitType()
	if err != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", err)
		os.Exit(1)
	}

	err = splitIo.SplitDataToFiles(splitType)

}
