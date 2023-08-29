package main

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// file handling will live here
func handleWCCmd(cmd *cobra.Command, args []string) {
	// first check if file exists
	// for _, fileName := range args {
	// 	return fileName

	// }
	// loop over all files if multiple is passed
	// permissions around opening files
	// what is the output if you give multiple files?

	// output
	// treats files independently err if DNE
	// also return total line count

}

// this may return something else based on flag
// l -> # of lines
// c -> # of bytes
// w -> # of words
// m -> # of chars
func handleFileLines(r io.Reader) (int, error) {
	return 0, nil
}

func openFile(fileName string) (*os.File, error) {
	// TODO: if the file is large, we probs don't want this in memory
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)

	if err != nil {
		return nil, fmt.Errorf("%s: open: %s", fileName, err)
	}

	return file, nil
}
