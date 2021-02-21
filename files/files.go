package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SetFile(path, name string, data interface{}) (err error) {
	t := time.Now()
	// t.Format("2006")
	// t.Format("01")
	// t.Format("02")
	date := GetCurrentDirectory() + "/" + path + "/" + t.Format("2006") + "/" + t.Format("01") + "/" + t.Format("02")
	//fmt.Println(date)
	err = os.MkdirAll(date, os.ModePerm)
	if err != nil {
		fmt.Println(date, err)
		return err
	}
	return bufferWrite(date+"/"+name, data)

}

func bufferWrite(w3 string, param interface{}) error {
	fileHandle, err := os.OpenFile(w3, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(w3, "open file error :", err)
		return err
	}
	defer fileHandle.Close()
	// NewWriter 默认缓冲区大小是 4096
	// 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
	buf := bufio.NewWriter(fileHandle)
	// 字节写入
	if data, ok := param.([]byte); ok {
		buf.Write(data)
		// 将缓冲中的数据写入
		err := buf.Flush()
		if err != nil {
			log.Println(w3, "flush error :", err)
		}
		return err
	}

	// 字符串写入
	if data, ok := param.(string); ok {
		buf.WriteString(data)
		// 将缓冲中的数据写入
		err := buf.Flush()
		if err != nil {
			log.Println(w3, "flush error :", err)
		}
		return err
	}
	return fmt.Errorf("传入的写入内容不是字符串或字符数组.写入失败!")

}

func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}
