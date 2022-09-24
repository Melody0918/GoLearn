package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	sort2 "sort"
	"strconv"
	"time"
)

func writeDataToFile(writeFileName string, exitChan chan int) {
	rand.Seed(time.Now().UnixNano())
	file, err := os.OpenFile(writeFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 2000; i++ {
		writer.WriteString(strconv.Itoa(rand.Intn(1000)+1) + "\n")
	}
	writer.Flush()
	exitChan <- 1
}

func sort(writeFileName string, storeFileName string, exitChan chan int) {
	file, err := os.OpenFile(writeFileName, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var slice []int
	//循环读取文件中的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = str[:len(str)-1]
		i, err1 := strconv.Atoi(str)
		fmt.Println(i)
		if err1 != nil {
			fmt.Println("Atoi转换err=", err1)
			break
		}
		slice = append(slice, i)
	}
	fmt.Println(slice)
	s := sort2.IntSlice(slice)
	s.Sort() //递增排序
	fmt.Println(slice)
	storeFile, err2 := os.OpenFile(storeFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}
	defer storeFile.Close()
	writer := bufio.NewWriter(storeFile)
	for _, v := range slice {
		writer.WriteString(strconv.Itoa(v) + "\n")
	}
	writer.Flush()
	exitChan <- 1
}

//作业2
func main() {
	//开一个协程writeDataToFile,随机生成1000个数据，存放到文件中
	//var writeChan chan int = make(chan int, 2000)
	start := time.Now().Nanosecond()
	var exitChan1 = make(chan int, 1)
	var exitChan2 = make(chan int, 1)
	writeFileName := "./Concurrency/channelpractice3/temp.txt"
	storeFileName := "./Concurrency/channelpractice3/result.txt"
	//当writeDataToFile完成写100个数据到文件后，让sort协程从文件中读取1000个文件，并完成排序，重新写入到另一个文件中
	//优化前单协程
	//go writeDataToFile(writeFileName, exitChan1)
	//for i := 0; i < 1; i++ {
	//	<-exitChan1
	//}

	//多协程优化速度
	for i := 0; i < 10; i++ {
		go writeDataToFile(writeFileName, exitChan1)
	}
	for i := 0; i < 10; i++ {
		<-exitChan1
	}
	//优化前单线程
	go sort(writeFileName, storeFileName, exitChan2)
	for i := 0; i < 1; i++ {
		<-exitChan2
	}
	//优化后多线程
	//for i := 0; i < 2; i++ {
	//	go sort(writeFileName, storeFileName, exitChan2)
	//}
	//
	//for i := 0; i < 2; i++ {
	//	<-exitChan2
	//}//效果反而变差
	end := time.Now().Nanosecond()
	fmt.Printf("程序运行时间是%vns", end-start) //一个协程:236900900ns,优化后:217969400ns
}
