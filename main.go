package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("app started..")
	cflag := flag.Bool("c", false, "counts the number the bytes in the file")
	lflag := flag.Bool("l", false, "counts the number the lines in the file")
	wflag := flag.Bool("w", false, "counts the number the words in the file")
	mflag := flag.Bool("m", false, "counts the number the multibyte in the file")
	flag.Parse()

	if !*cflag && !*lflag && !*wflag && !*mflag {
		true_val := true
		cflag = &true_val
		lflag = &true_val
		wflag = &true_val
		mflag = &true_val
	}

	// get the file name
	args := flag.Args()
	// fmt.Println(args)

	//todo - check if the file name is present

	// get file name and read file
	file_name := args[0]
	file, _ := os.Open(file_name)
	main_str := ""

	if *cflag {
		b, _ := io.ReadAll(file)
		main_str += " c: " + strconv.Itoa(len(b))
	}

	if *lflag {
		file.Seek(0, io.SeekStart)
		b, _ := io.ReadAll(file)
		counter := 1
		for i := 0; i < len(b); i++ {
			if b[i] == 10 {
				counter++
			}
		}
		main_str += " l: " + strconv.Itoa(counter)
	}

	if *wflag {
		file.Seek(0, io.SeekStart)
		file_scanner := bufio.NewScanner(file)
		file_scanner.Split(bufio.ScanWords)
		counter := 0
		for file_scanner.Scan() {
			counter++
		}
		// fmt.Println(counter)
		main_str += " w: " + strconv.Itoa(counter)
	}

	if *mflag {
		file.Seek(0, io.SeekStart)
		bytes, _ := io.ReadAll(file)
		str := string(bytes)
		counter := 0

		for range str {
			counter++
		}
		main_str += " m: " + strconv.Itoa(counter)
	}
	fmt.Println(main_str)
}
