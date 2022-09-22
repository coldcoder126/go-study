package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func file1() {
	//创建一个文件
	file, err := os.Create("file.txt")
	defer file.Close()
	// 写文件
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}

// bufio实现了带缓冲的读写，不用再手动设置缓存区
// 写文件
func filebuf1() {
	file, err := os.OpenFile("filebuf.txt", os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(fmt.Sprintf("this number is : %v \n", i))
	}
	writer.Flush()
}

// 读文件
func file2() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	var buf [128]byte
	var content []byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file err")
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

// 读文件，带缓存
func filebuf2() {
	file, err := os.Open("filebuf.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}

// 拷贝文件
func file3() {
	srcFile, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	dstFile, err := os.Create("abc.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 缓冲读取
	buf := make([]byte, 32)
	count := 0
	for {
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		dstFile.Write(buf[:n])
		count++
		fmt.Println(count)
	}
	srcFile.Close()
	dstFile.Close()
}
