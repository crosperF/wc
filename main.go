package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println("app started..")
	cflag := flag.Bool("c", false, "counts the number the bytes in the file")
	lflag := flag.Bool("l", false, "counts the number the lines in the file")
	wflag := flag.Bool("w", false, "counts the number the words in the file")
	flag.Parse()

	// fmt.Println(*cflag)

	// get the file name
	args := flag.Args()
	// fmt.Println(args)

	//todo - check if the file name is present

	// get file name and read file
	file_name := args[0]
	file, _ := os.Open(file_name)
	b, _ := io.ReadAll(file)

	if *cflag {
		fmt.Println(len(b))
	}

	if *lflag {
		counter := 1
		for i := 0; i < len(b); i++ {
			if b[i] == 10 {
				counter++
			}
		}
		fmt.Println(counter)
	}

	if *wflag {

		word_counter := 0

		for i := 0; i < len(b); i++ {
			if b[i] == 32 { // space
				word_counter += 1
			}

			if b[i] == 10 { // new line
				word_counter += 1
			}
		}
		// word_counter++
		fmt.Println(word_counter)
	}
}
