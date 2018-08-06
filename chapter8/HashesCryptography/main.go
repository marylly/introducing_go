package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	// remember to always close opened files
	defer f.Close()

	// create a hasher
	h := crc32.NewIEEE()

	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err1 := io.Copy(h, f)
	// we don't care about how many bytes were written, but we do want to 
	// handle the error
	if err1 != nil {
		return 0, err1
	}
	return h.Sum32(), nil
}

func main() {
	// create a hasher
	h := crc32.NewIEEE()

	// write our data to it
	h.Write([]byte("test"))

	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)

	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}
	h2, err1 := getHash("test2.txt")
	if err1 != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}