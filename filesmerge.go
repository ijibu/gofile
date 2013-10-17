// 文件合并filesmerge.go
package main

import (
	"fmt"
	"io"
	"os"
)

func mergeFile(files []string, fileName string) {
	num := len(files)
	//每次最多拷贝10m
	bufsize := 1024 * 1024 * 10
	buf := make([]byte, bufsize)

	newfile, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("failed to create file", fileName)
	} else {
		fmt.Println("create file:", fileName)
	}

	totalLen := 0
	for i := 0; i < num; i++ {
		copylen := 0
		file, err := os.Open(files[i])
		if err != nil {
			fmt.Println("open file failed:", files[i])
		}
		finfo, err := file.Stat()
		if err != nil {
			fmt.Println("get file info failed:", file)
		}
		fillsize := int(finfo.Size())
		size := (fillsize + bufsize - 1) / bufsize
		fmt.Printf("file size:", size)

		for copylen < size {
			n, err2 := file.Read(buf)
			if err2 != nil && err2 != io.EOF {
				fmt.Println(err2, "failed to read from:", file)
				break
			}

			if n <= 0 {
				break
			}

			//fmt.Println(n, len(buf))
			//写文件
			w_buf := buf[:n]
			var off1 int64 = int64(bufsize * totalLen)
			newfile.WriteAt(w_buf, off1)
			copylen += 1
			totalLen += 1
		}
	}

	return
}

func main() {
	files := []string{"navicat100_sqlserver_en.rar0", "navicat100_sqlserver_en.rar1", "navicat100_sqlserver_en.rar2", "navicat100_sqlserver_en.rar3", "navicat100_sqlserver_en.rar4", "navicat100_sqlserver_en.rar5", "navicat100_sqlserver_en.rar6", "navicat100_sqlserver_en.rar7", "navicat100_sqlserver_en.rar8"}
	fileName := "ijibu.rar"

	mergeFile(files, fileName)
}
