package testing_prac

import (
	"fmt"
	"os"
	"testing"
)

func TestMonster_Store(t *testing.T) {
	monster := Monster{
		Name:  "张三",
		Age:   24,
		Skill: "Coding",
	}
	//序列化前
	fmt.Println("序列化前", monster)
	b := monster.Store()
	if b != true {
		t.Fatalf("Monster结构体序列化保存失败！")
	}
	t.Logf("Monster结构体序列化保存成功！")
}

func TestReStore(t *testing.T) {
	var monster Monster
	filepath := "../testprac.txt"
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		t.Fatalf("打开文件错误！")
	}
	t.Logf("打开文件成功！")
	defer file.Close()
	err1 := monster.ReStore(file)
	if err1 != true {
		t.Fatalf("反序列化错误！")
	}
	t.Logf("反序列化成功,反序列化后为：%v", monster)
}
