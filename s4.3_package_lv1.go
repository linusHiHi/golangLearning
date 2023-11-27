package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func fileOpen(nameFile string) *os.File {
	//initialize
	file, errOp := os.OpenFile(nameFile, os.O_RDWR|os.O_CREATE, 0666)
	if errOp != nil {
		fmt.Println("打开错误", errOp)
		return nil
	}
	return file
}

func newTestTXT() *os.File {
	strOrigin := []byte("abcd")
	file := fileOpen("testTXT.txt")
	defer file.Close()

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

// type readerF func (file *os.File)
//
// func reader(f readerF, url string)  {
//
// }
func readInBuf(file *os.File) {
	defer file.Close()
	reader := bufio.NewReader(file)
	block := make([]byte, 2*512)
	for {
		_, err := reader.Read(block)
		fmt.Printf("%s\n", block)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
	}

}

func readInOs(file *os.File) {
	buf := make([]byte, 2*512)
	defer file.Close()

	for {
		_, err := file.Read(buf)
		fmt.Printf("%s\n", buf)
		if err != nil && err != io.EOF {
			fmt.Println("读取文件出错:", err)
			panic(err)
		} else if err == io.EOF {
			println("osreader读完了")
			break
		}
	}
}

func main() {

	//println("hello")
	ioTimeA := time.Now()
	readInOs(fileOpen("testTXT.txt"))
	ioTimeB := time.Since(ioTimeA)

	//println("hello")
	bufTimeA := time.Now()
	readInBuf(fileOpen("testTXT.txt"))
	bufTimeB := time.Since(bufTimeA)

	fmt.Printf("ioReader running time: %v\nbuf running time: %v\n", ioTimeB, bufTimeB)

}
