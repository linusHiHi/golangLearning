package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
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

// timestampWriter 是一个实现 io.Writer 接口的结构体，它在写入数据前添加时间戳
type timestampWriter struct {
	timestamp time.Time
	of        *os.File
}

func (tw *timestampWriter) Write(p []byte) (n int, err error) {
	// 添加时间戳和时间
	println(tw.timestamp.Unix())
	stamp := testBinaryWrite(tw.timestamp.Unix())
	for _, v := range stamp {
		p = append(p, v)
	}
	//p = append(p, '\n')
	// 输出到文件
	n, err = tw.of.Write(p)
	return n, err
}

func testBinaryWrite(x interface{}) []byte {
	buf := new(bytes.Buffer)
	// for int32, the resulting size of buf will be 4 bytes
	// for int64, the resulting size of buf will be 8 bytes
	err := binary.Write(buf, binary.BigEndian, x)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("%x\n", buf.Bytes())
	return buf.Bytes()
}
