package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件进行读写操作（如果文件不存在则创建，文件权限为0666）
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// 初始化 bufio.Writer
	writer := bufio.NewWriter(file)
	// 写入数据到缓冲区
	data := []byte("Hello, Golang!")
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println("写入数据到缓冲区出错:", err)
		return
	}

	// 获取缓冲区可用字节数
	availableBytes := writer.Available()
	fmt.Printf("缓冲区可用字节数: %d\n", availableBytes)

	// 获取底层缓冲区的大小
	bufferSize := writer.Size()
	fmt.Printf("底层缓冲区的大小: %d\n", bufferSize)

	// 刷新缓冲区到底层的 io.Writer（即文件）
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区出错:", err)
		return
	}
	fmt.Println("刷新缓冲区成功")

	// 初始化 bufio.Reader    file.Close()
	file, _ = os.Open("example.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := bufio.NewReader(file)

	// 从缓冲区读取数据
	readData := make([]byte, 12)
	n, err := reader.Read(readData)
	if err != nil {
		if err != io.EOF {
			fmt.Println("从缓冲区读取数据出错:", err)
			return
		}
	}
	fmt.Printf("从缓冲区读取 %d 字节的数据: %s\n", n, readData)

	// 使用 Peek 方法
	peekData, err := reader.Peek(5)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Peek 出错:", err)
			return
		}
	}
	fmt.Printf("Peek 到的数据: %s\n", peekData)

	// 使用 Reset 方法
	newBuffer := bytes.NewBufferString("Reset Buffer")
	reader.Reset(newBuffer)

	resetData, err := reader.Read(readData)
	if err != nil {
		fmt.Println("从重置后的缓冲区读取数据出错:", err)
		return
	}
	fmt.Printf("从重置后的缓冲区读取 %d 字节的数据: %s\n", resetData, readData[:resetData])
}
