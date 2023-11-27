package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

/*
因为之前写了测读取时间的程序，所以本程序测试了写入然后再读取耗费的时间。
主要过程就是分别用两种方法写入“abcd\n”*1024,然后读取并数出其中“a”的个数。
*/
//定义文件名，方便随时修改
const textPathOs = "testTXTos.txt"
const textPathBuf = "testTXTBuf.txt"

// ”读写函数“类型，在装饰器中作为参数类型
type function func(file *os.File) int

func main() {
	//decoration
	writeOsDec := fileDec(writeOs, textPathOs)
	writeBufDec := fileDec(writeBuf, textPathBuf)
	readByOsDec := fileDec(readByOs, textPathOs)
	readByBufDec := fileDec(readByBuf, textPathBuf)
	//os包测速
	ioTimeA := time.Now()
	_ = writeOsDec()
	//count* 变量是接受数得“a”的个数的
	CountOs := readByOsDec()
	ioTimeB := time.Since(ioTimeA)
	//bufio测速
	bufTimeA := time.Now()
	_ = writeBufDec()
	CountBuf := readByBufDec()
	bufTimeB := time.Since(bufTimeA)
	//打印耗时和a的数量
	fmt.Printf("ioReadWriter running time: %v\nbufReadWriter running time: %v\n", ioTimeB, bufTimeB)
	fmt.Printf("ioReader count a: %d\nbuf count a: %d\n", CountOs, CountBuf)
}

// file read&write operation decoration 作用是打开关闭文件
func fileDec(fn function, path string) func() int {
	fnDec := func() int {
		f, errOpen := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
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
	return fnDec
}

// 用bufio 来写入数据，以下类推。
// 返回值是为了与function类型匹配。
func writeBuf(f *os.File) int {
	//initializing
	strOrigin := []byte("abcd\n")
	bufWriter := bufio.NewWriter(f)
	//writing
	for i := 1; i <= 1024; i++ {
		_, errWr := bufWriter.Write(strOrigin)
		if errWr != nil {
			panic(errWr)
		}
	}
	//flushing
	err := bufWriter.Flush()
	if err != nil {
		panic(err)
	}
	return 0
}

func writeOs(f *os.File) int {
	strOrigin := []byte("abcd\n")
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
			panic(err)
		}
		if err == io.EOF {
			break
		}
		//数a的个数
		if buf[0] == 'a' || buf[1] == 'a' {
			count++
		}
	}
	return count
}
