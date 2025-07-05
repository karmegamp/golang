package main

import (
	"bytes"
	"container/list"
	"crypto/sha1"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// 1. String functions

	str := "Karmegam"

	fmt.Println(strings.Contains(str, "Kar"))
	fmt.Println(strings.Count(str, "m"))
	fmt.Println(strings.HasPrefix(str, "Kar"))
	fmt.Println(strings.HasSuffix(str, "megam"))
	fmt.Println(strings.Index(str, "m"))
	fmt.Println(strings.LastIndex(str, "m"))

	st1 := []string{"h", "hh"} // custom init slice of string
	fmt.Println(strings.Join(st1, "-"))

	str2 := make([]string, 1) // string slice default init using make, size mandatory
	str2[0] = "g"
	fmt.Println(strings.Repeat(str2[0], 10))

	str2 = append(str2, "ff") // grow slice
	fmt.Println(str2[1])

	fmt.Println(strings.Replace("aaaa", "a", "b", 3)) // bbba, replace a with b in first 3 occurance
	slstr := strings.Split("a-b-c-d", "-")
	fmt.Printf("%v, %T\n", slstr, slstr)

	fmt.Println(strings.ToUpper("abcd"))
	fmt.Println(strings.ToLower("Abcd"))

	fmt.Println([]byte("test"))                     // string to slice of byte
	fmt.Println(string([]byte{'t', 'e', 's', 't'})) // slice of byte to string
	fmt.Println(string([]byte{116, 101, 115, 116}))

	// 2. I/O  and Buffer handling
	// Buffer is always character/byte based, so convert to []byte is mandatory
	// Buffer bytes can be read using Bytes() method and converted to string for display
	// Buffer need to be streammed for direct read or write
	// All file contents are nothing but sequence of byte/charater

	var buf bytes.Buffer
	buf.Write([]byte("test"))                // string to []byte to buffer
	fmt.Println(string([]byte(buf.Bytes()))) // buffer to string  buf->[]byte->string
	io.Copy(os.Stdout, &buf)                 // buffer to stream

	// 3. Files and folder
	file, _ := os.Open("go_packages.go")
	// defer file.Close()

	// allocate bytes and read file
	stat, _ := file.Stat()
	bs := make([]byte, stat.Size())
	file.Read(bs)

	// convert bytes and print
	str = string(bs)
	fmt.Println(str)
	file.Close()

	// shorthand read file
	// bs, _ = ioutil.ReadFile("go_packages.go")
	// str = string(bs)
	// fmt.Println(str)

	// create a file
	f, _ := os.Create("test.txt")
	f.WriteString("Hello")
	f.Close()

	// Get directory content
	dir, _ := os.Open(".")
	fileInfo, _ := dir.Readdir(0) // value <=0 means all entries
	for _, fi := range fileInfo {
		fmt.Println(fi.Name(), fi.Size())
	}

	// Directory walk
	filepath.WalkDir("..", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})

	// custom error
	err := errors.New("custom error")
	fmt.Println(err)
	fmt.Println(err.Error())

	var b bytes.Buffer
	b.Write([]byte(err.Error()))
	io.Copy(os.Stdout, &b)

	// Container
	var l list.List // doubly linked list
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// for e := l.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value.(int)) // type assertion
	// }

	l.PushBack("\nfour")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	// Hash and Cryptography
	getHash := func(filename string) uint32 {
		f, _ := os.Open(filename)
		defer f.Close()
		h := crc32.NewIEEE() // get hash object
		io.Copy(h, f)
		return h.Sum32() // get hash for file
	}

	h1 := getHash("test.txt")
	h2 := getHash("test.txt")
	h3 := getHash("go_packages.go")

	fmt.Println(h1, h2, h1 == h2)
	fmt.Println(h1, h3, h1 == h3)

	// Sha
	sh := sha1.New()

	// need bytes to create sha
	f, _ = os.Open("test.txt")
	fs, _ := f.Stat()
	bf := make([]byte, fs.Size())
	f.Read(bf)
	fmt.Println(string(bf))
	sh.Write(bf)
	// io.Copy(sh, bf)

	fmt.Println("sha of", string(bf), " ", sh.Sum([]byte{}), sh.Size()) // return sha bytes
}
