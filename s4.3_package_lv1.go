package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const textPath = "testTXT.txt"

func main() {
	//decortation
	newOsDec := fileDec(newOs)
	newBufDec := fileDec(newBuf)
	readByOsDec := fileDec(readByOs)
	readByBufDec := fileDec(readByBuf)

	ioTimeA := time.Now()
	_ = newOs()
	ioTimeB := time.Since(ioTimeA)

	bufTimeA := time.Now()

	bufTimeB := time.Since(bufTimeA)

	fmt.Printf("ioReader running time: %v\nbuf running time: %v\n", ioTimeB, bufTimeB)

}

// file operation decoration
func fileDec(fn func(file *os.File) int) func() int {
	return func() int {
		f, errOpen := os.OpenFile(textPath, os.O_RDWR|os.O_CREATE, 0666)
		if errOpen != nil {
			panic(errOpen)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)
		result := fn(f)
		return result
	}
}

func newBuf(f *os.File) int {
	//initializing
	strOrigin := []byte("abcd")
	bufWriter := bufio.NewWriter(f)
	//writing
	for i := 1; i <= 1024; i++ {
		_, errWr := bufWriter.Write(strOrigin)
		if errWr != nil {
			println("写入错误", errWr)
		}
	}
	//flushing
	err := bufWriter.Flush()
	if err != nil {
		panic(err)
	}
	return 0
}

func newOs(f *os.File) int {
	strOrigin := []byte("abcd")
	for i := 1; i <= 1024; i++ {
		_, err := f.Write(strOrigin)
		if err != nil {
			panic(err)
		}
	}
	return 0
}

func readByBuf(f *os.File) int {
	reader := bufio.NewReader(f)
	b := make([]byte, 2)
	count := 0
	for {
		_, err := reader.Read(b)

		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		//fmt.Printf("%s\n", block)
		if rune(b[0]) == 'a' || rune(b[1]) == 'a' {
			count++
		}
	}
	return count
}

func readByOs(file *os.File) int {
	buf := make([]byte, 2)
	count := 0
	for {
		_, err := file.Read(buf)

		if err != nil && err != io.EOF {
			//fmt.Println("读取文件出错:", err)
			panic(err)
		} else if err == io.EOF {
			//println("osreader读完了")
			break
		}
		//fmt.Printf("%s\n", buf)
		if buf == []byte("a") {
			count++
		}
		return count
	}
}
