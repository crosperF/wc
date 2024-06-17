package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
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
	input_type := false
	var input_text []byte

	if len(args) == 0 {
		input_type = true
		input, _ := io.ReadAll(os.Stdin)
		input_text = input
	}

	main_str := ""
	var file *os.File

	if !input_type {
		// get file name and read file
		file_name := args[0]
		file, _ = os.Open(file_name)
		main_str = ""
	}

	if *cflag {
		ans := 0
		if !input_type {
			b, _ := io.ReadAll(file)
			ans = len(b)
		} else {
			ans = len(input_text)
		}
		main_str += " c: " + strconv.Itoa(ans)
	}

	if *lflag {
		b := []byte{}
		if !input_type {
			file.Seek(0, io.SeekStart)
			by, _ := io.ReadAll(file)
			b = by
		} else {
			b = input_text
		}
		counter := 1
		for i := 0; i < len(b); i++ {
			if b[i] == 10 {
				counter++
			}
		}
		main_str += " l: " + strconv.Itoa(counter)
	}

	if *wflag {
		counter := 0
		if !input_type {
			file.Seek(0, io.SeekStart)
			file_scanner := bufio.NewScanner(file)
			file_scanner.Split(bufio.ScanWords)
			counter = 0
			for file_scanner.Scan() {
				counter++
			}
		} else {
			s := string(input_text)
			counter = len(strings.Fields(s))
		}

		main_str += " w: " + strconv.Itoa(counter)
	}

	if *mflag {
		bytes := []byte{}
		if !input_type {
			file.Seek(0, io.SeekStart)
			b, _ := io.ReadAll(file)
			bytes = b
		} else {
			bytes = input_text
		}
		str := string(bytes)
		counter := 0

		for range str {
			counter++
		}
		main_str += " m: " + strconv.Itoa(counter)
	}
	fmt.Println(main_str)
}
