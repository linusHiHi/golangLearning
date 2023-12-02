package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// 打开一个日志文件，如果文件不存在则创建，追加写入
	loggerTxt, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败！", err)
	}
	defer loggerTxt.Close()
	// 创建一个带时间戳的写入器
	logWriter := &timestampWriter{
		timestamp: time.Now(),
		of:        loggerTxt,
	}

	// 模拟用户操作并记录日志
	fmt.Fprintln(logWriter, "用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作B")
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

// timestampWriter 是一个实现 io.Writer 接口的结构体，它在写入数据前添加时间戳
type timestampWriter struct {
	timestamp time.Time
	of        *os.File
}

func (tw *timestampWriter) Write(p []byte) (n int, err error) {
	// 添加时间戳和时间
	//println(tw.timestamp.Unix())
	var clock [3]int
	clockStr := []byte("")
	clock[0], clock[1], clock[2] = tw.timestamp.Clock()
	for _, v := range clock {
		clockStr = strconv.AppendInt(clockStr, int64(v), 10)
	}
	stamp := []byte(strconv.FormatInt(tw.timestamp.Unix(), 10))
	p = BytesCombine([]byte("时间："), clockStr, []byte("时间戳："), stamp, p)
	//p = append(p, '\n')
	// 输出到文件
	n, err = tw.of.Write(p)
	return n, err
}
