package model

//定义一个结构体

type student struct {
	Name string
	age  int
}

//因为student结构体首字母是小写，因此是只能在model使用
//我们通过工厂模式来解决

func Newstudent(n string, a int) *student {
	return &student{
		Name: n,
		age:  a,
	}
}

//如果age字段小写,则在其他包不可以直接方法，我们可以提供一个方法
func (s *student) GetScore() int {
	return s.age
}
