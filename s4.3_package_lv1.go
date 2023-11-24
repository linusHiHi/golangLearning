package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func fileOpen() *os.File {
	//initialize
	file, errOp := os.OpenFile("testTXT.txt", os.O_RDWR|os.O_CREATE, 0666)
	if errOp != nil {
		fmt.Println("打开错误", errOp)
		return nil
	}
	return file
}

func newTestTXT() *os.File {
	strOrigin := []byte("abcd")
	file := fileOpen()

	//use "bufio" write text for testing
	bufWriter := bufio.NewWriter(file)

	for i := 1; i <= 1024; i++ {
		_, errWr := bufWriter.Write(strOrigin)
		if errWr != nil {
			println("写入错误", errWr)
		}
	}
	err := bufWriter.Flush()
	if err != nil {
		return nil
	}
	return file
}

func readInBuf(file *os.File) {
	reader := bufio.NewReader(file)
	block := make([]byte, 2)
	for {
		n, err := reader.Read(block)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
	}

}

func readInOs(file *os.File) {
	buf := make([]byte, 2)

	for {
		_, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("读取文件出错:", err)
			} else if err == io.EOF {
				println("读完了")
				break
			}
		}
	}
}

func main() {
	testTxt := newTestTXT()

	println("hello")
	ioTimeA := time.Now()
	readInOs(testTxt)
	ioTimeB := time.Since(ioTimeA)

	println("hello")
	bufTimeA := time.Now()
	readInBuf(testTxt)
	bufTimeB := time.Since(bufTimeA)

	fmt.Printf("ioReader running time: %v\nbuf running time: %v\n", ioTimeB, bufTimeB)
	testTxt.Close()
}
