package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}
type StudentSlice []Student

func (stu StudentSlice) Len() int {
	return len(stu)
}
func (stu StudentSlice) Less(i, j int) bool {
	return stu[i].Score > stu[j].Score
}
func (stu StudentSlice) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {
	var students StudentSlice
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	r2 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 10; i++ {
		student := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   r1.Intn(16) + 16,
			Score: r2.Float64() * 100,
		}
		students = append(students, student)
	}
	fmt.Println("----------------排序前----------------")
	for _, x := range students {
		fmt.Printf("学生：%v;年龄：%v;成绩：%.2f\n", x.Name, x.Age, x.Score)
	}
	sort.Sort(students)
	fmt.Println("----------------排序后----------------")
	//for j := 0; j < len(students); j++ {
	//	fmt.Println(students[j])
	//}或者
	for i, j := range students {
		fmt.Printf("学生：%v;年龄：%v;成绩：%.2f;排名：%v\n", j.Name, j.Age, j.Score, i+1)
	}
}
