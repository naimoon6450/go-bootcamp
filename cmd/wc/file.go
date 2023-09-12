package main

import (
	"bufio"
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

  fileName := args[0]
  f, err := openFile(fileName)
  if err != nil {
    fmt.Printf("error opening file %w", err)
  }


  defer f.Close()

  lines, err := handleFileLines(f)
  if err != nil {
    fmt.Printf("error counting lines %w", err)
  }

	// output
	// treats files independently err if DNE
	// also return total line count

  fmt.Println(lines)
}

// this may return something else based on flag
// l -> # of lines
// c -> # of bytes
// w -> # of words
// m -> # of chars
func handleFileLines(r io.Reader) (int, error) {
  // check if flags were changed
  // call the required function
  // flag := cmd.Flags().Lookup("line count").Changed
  // call CountLines or CountWords depending on which flag is changed
  var lines int
  scanner := bufio.NewScanner(r)

  for scanner.Scan() {
    if len(scanner.Text()) > 0 {
      lines++
    }
  }

	return lines, nil
}

func openFile(fileName string) (*os.File, error) {
	// TODO: if the file is large, we probs don't want this in memory
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)

	if err != nil {
		return nil, fmt.Errorf("%s: open: %s", fileName, err)
	}

	return file, nil
}
