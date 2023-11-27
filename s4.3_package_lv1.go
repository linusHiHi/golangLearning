package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const textPath = "testTXT.txt"

type re

func main() {
	ioTimeA := time.Now()
	readInOs(os.OpenFile("testTXT.txt"))
	ioTimeB := time.Since(ioTimeA)

	bufTimeA := time.Now()
	readByBuf(fileOpen("testTXT.txt"))
	bufTimeB := time.Since(bufTimeA)

	fmt.Printf("ioReader running time: %v\nbuf running time: %v\n", ioTimeB, bufTimeB)

}

//file operation decoration
func fileDec(fn func(file *os.File)int)  {
	f, errOpen := os.OpenFile(textPath, os.O_RDWR|os.O_CREATE, 0666)
	if errOpen != nil{panic(errOpen)}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {panic(err)}
	}(f)

	fn(f)
}

func newTestTXT(f *os.File)int{
	//initializing
	strOrigin := []byte("abcd")
	bufWriter := bufio.NewWriter(f)
	//writing
	for i := 1; i <= 1024; i++ {
		_, errWr := bufWriter.Write(strOrigin)
		if errWr != nil {println("写入错误", errWr)}
	}
	//flushing
	err := bufWriter.Flush()
	if err != nil {panic(err)}
	return 0
}


func readByBuf(f *os.File) int{
	reader := bufio.NewReader(f)
	b := make([]byte, 2)
	count := 0
	for {
		_, err := reader.Read(b)

		if err != nil && err != io.EOF {panic(err)}
		if err == io.EOF {break}
		//fmt.Printf("%s\n", block)
		if b == []byte("a"){
			count++
		}
	}
	return count
}

func readInOs(file *os.File) int{
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
		if buf == []byte("a"){count++}
		return count
	}
}
