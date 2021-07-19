package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Common way to work with files is to use the "os" package
	// It provides unifom behavior across all OSes. The design
	// of the package is unix-like. However, the errpr handling
	// is go-like.

	// Creating new File
	var newFile *os.File
	fmt.Printf("newFile is of type %T\n", newFile)
	var err error

	newFile, err = os.Create("scratchpad/a.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Create file failed with err:", err)
		os.Exit(1)
	}
	newFile.Close()
	// This fails because the path "scratchpad/new/" isn't present
	_, err = os.Create("scratchpad/new/b.txt")
	if err != nil {
		// SHOULD enter here
		fmt.Println("Create file failed with err:", err)
		// commeting call to exit to prevent program from exiting here
		//os.Exit(1)
	}

	// Opening existing file. Opens in read mode by default
	newFile, err = os.Open("scratchpad/a.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file:", err)
		os.Exit(1)
	}
	newFile.Close()

	// Opening a file in append mode
	newFile, err = os.OpenFile("scratchpad/a.txt", os.O_APPEND, 0777)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file:", err)
		os.Exit(1)
	}
	newFile.Close()

	// Get stat of opened file
	fileInfo, err := os.Stat("scratchpad/a.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to get file info:", err)
		os.Exit(1)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Is dir?:", fileInfo.IsDir())
	fmt.Println("Permission:", fileInfo.Mode())
	fmt.Println("File size:", fileInfo.Size())

	// Stat can also be used to check if file doesn't exist
	fileInfo, err = os.Stat("scratchpad/new/a.txt")
	if err != nil {
		// SHOULD enter here
		fmt.Println("File doesn't exist:", err)
		//os.Exit(1)
	}

	// Renaming a file
	oldPath := "scratchpad/a.txt"
	newPath := "scratchpad/a_new.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Rename failed:", err)
	}

	// Remove a file
	err = os.Remove(newPath)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Remove failed:", err)
	}

	// Writing byte slices to a file
	newFile, err = os.OpenFile(
		"scratchpad/b.txt",
		// create file if it doesn't exist and open in write mode
		// or truncate file and open in write mode
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0655)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file:", err)
	}
	// Create a byteslice from string to write into file
	byteSlice := []byte("Hello World!\n")
	//Notice we're calling write on newFile, the handle opened for file
	bytesWritten, err := newFile.Write(byteSlice)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to write to file:", err)
	}
	fmt.Println("Wrote", bytesWritten, "bytes to file")
	newFile.Close()

	// Writing byte slices to a file using ioutil
	// ioutil.WriteFile takes care of the following:
	// 1. opening file in write mode if it doesn't exist
	// 2. truncating and opening file in write mode if it does exist
	// 3. writing byte slice to file
	// 4. closing the file
	newByteSlice := []byte("foo bar\n")
	err = ioutil.WriteFile("scratchpad/c.txt", newByteSlice, 0655)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to write to file:", err)
	}

	// ioutil.WriteFile & os.Write both write into a file, leading to disk i/o
	// while performing multiple writes to a file. Instead bufio provides a
	// means to write to memory and flush i.e. write to file either when the
	// buffer is full or when explicit flush is called
	newFile, err = os.OpenFile(
		"scratchpad/d.txt",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0655)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file:", err)
	}

	buffWriter := bufio.NewWriter(newFile)
	bytesWritten, err = buffWriter.Write([]byte("first line\n"))
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to write to buffer:", err)
	}
	fmt.Println("Wrote", bytesWritten, "bytes to buffer")
	bytesWritten, err = buffWriter.WriteString("second line\n")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to write to buffer:", err)
	}
	fmt.Println("Bytes available in buffer:", buffWriter.Available())
	fmt.Println("Bytes unflushed from buffer:", buffWriter.Buffered())
	fileInfo, err = os.Stat("scratchpad/d.txt")
	fmt.Println("File size before flush:", fileInfo.Size())
	// note: any content in the buffer will not be flushed to file if the file
	// is closed without calling Flush
	buffWriter.Flush()
	fileInfo, err = os.Stat("scratchpad/d.txt")
	fmt.Println("File size after flush:", fileInfo.Size())
	// We can reset buffer if we've not flushed data into file
	bytesWritten, err = buffWriter.Write([]byte("third line\n"))
	if err != nil {
		// SHOULD NOT enter here
	}
	fmt.Println("Bytes unflushed from buffer before reset:", buffWriter.Buffered())
	buffWriter.Reset(buffWriter)
	fmt.Println("Bytes unflushed from buffer after reset:", buffWriter.Buffered())
	newFile.Close()

	// Reading from file
	newFile, err = os.Open("scratchpad/d.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file for read:", err)
	}
	// Read 200 bytes using file.Read
	var content = make([]byte, 200)
	bytesRead, err := newFile.Read(content)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to read file:", err)
	}
	fmt.Println("Read", bytesRead, "bytes")
	fmt.Println("file content is:")
	fmt.Println(strings.Repeat("#", 20))
	fmt.Println(string(content))
	fmt.Println(strings.Repeat("#", 20))
	newFile.Close()
	fmt.Println()

	// io.ReadFull achieves a similar behavior
	newFile, err = os.Open("scratchpad/d.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file for read:", err)
	}
	var content1 = make([]byte, 20)
	bytesRead, err = io.ReadFull(newFile, content1)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to read file:", err)
	}
	fmt.Println("Read", bytesRead, "bytes")
	fmt.Println("file content is:")
	fmt.Println(strings.Repeat("#", 20))
	fmt.Println(string(content1))
	fmt.Println(strings.Repeat("#", 20))
	newFile.Close()
	fmt.Println()

	// ioutil.ReadAll reads all contents of the file
	newFile, err = os.Open("scratchpad/d.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file for read:", err)
	}
	data, err := ioutil.ReadAll(newFile)
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to read file:", err)
	}
	fmt.Println("Read", len(data), "bytes")
	fmt.Println("file content is:")
	fmt.Println(strings.Repeat("#", 20))
	fmt.Println(string(data))
	fmt.Println(strings.Repeat("#", 20))
	newFile.Close()
	fmt.Println()

	// ioutil.ReadFile takes care of opening the file, reading all data
	// and closing the file
	data, err = ioutil.ReadFile("scratchpad/d.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to read file:", err)
	}
	fmt.Println("Read", len(data), "bytes")
	fmt.Println("file content is:")
	fmt.Println(strings.Repeat("#", 20))
	fmt.Println(string(data))
	fmt.Println(strings.Repeat("#", 20))
	fmt.Println()

	// Sometimes you need to read the file line by line or read the
	// contents until you hit a delimiter. bufio.Scanner us useful in
	// such scenarios.
	newFile, err = os.Open("scratchpad/d.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file for read:", err)
	}
	bufScanner := bufio.NewScanner(newFile)
	idx := 1
	for bufScanner.Scan() {
		fmt.Println("line no", idx, ":", bufScanner.Text())
		idx += 1
	}
	if bufScanner.Err() != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to scan file:", bufScanner.Err())
	} else {
		fmt.Println("Scanner completed and reached end of file")
	}
	newFile.Close()

	// To scane word by word
	newFile, err = os.Open("scratchpad/d.txt")
	if err != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to open file for read:", err)
	}
	bufScanner = bufio.NewScanner(newFile)
	// bufio.ScanRunes scans rune by rune (one utf8 char per scan)
	bufScanner.Split(bufio.ScanWords)
	idx = 1
	for bufScanner.Scan() {
		fmt.Println("word no", idx, ":", bufScanner.Text())
		idx += 1
	}
	if bufScanner.Err() != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to scan file:", bufScanner.Err())
	} else {
		fmt.Println("Scanner completed and reached end of file")
	}
	newFile.Close()

	// Reading from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your name:")
	scanner.Scan()
	fmt.Println("Hello", scanner.Text())
	if scanner.Err() != nil {
		// SHOULD NOT enter here
		fmt.Println("Failed to scan file:", scanner.Err())
	} else {
		fmt.Println("Successfully scanned input from STDIN")
	}

}
