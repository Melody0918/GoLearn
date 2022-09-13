package main

import "fmt"

//编写一个学生考试系统

type Student struct {
	Name  string
	Age   int
	Score float64
}

//将Pupil和Graduate共有的方法也绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Println("学生姓名：", stu.Name, "学生年龄：", stu.Age, "学生分数：", stu.Score)
}

func (stu *Student) SetScore(score float64) {
	if score > 0 && score <= 100 {
		stu.Score = score
	} else {
		fmt.Println("分数设置错误...")
	}
}

//小学生
type Pupil struct {
	Student //嵌入了Student匿名结构体
}

//显示成绩
//这是Pupil结构体特有的方法，保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

type Graduate struct {
	Student
}

//这是Graduate独有的方法
func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

func main() {
	pupil := Pupil{}
	pupil.Student.Name = "kp"
	pupil.Student.Age = 10
	pupil.Student.Score = 60
	pupil.testing()
	pupil.SetScore(80)
	pupil.ShowInfo()
	pupil.Student.ShowInfo()

	graduate := Graduate{}
	graduate.Student.Name = "kp"
	graduate.Student.Age = 22
	//graduate.Student.Score = 60
	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
	graduate.Student.ShowInfo()
}
