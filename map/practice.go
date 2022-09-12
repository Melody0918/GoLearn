package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	count := 0
	for key, _ := range users {
		count++
		if key == name {
			users[key]["pwd"] = "888888"
			break
		}
		if count == len(users) && key != name {
			users[name] = make(map[string]string, 2)
			users[name]["nickname"] = name
			users[name]["pwd"] = "123456"
		}
	}
}

func main() {
	var users map[string]map[string]string
	users = make(map[string]map[string]string, 5)

	users["张三"] = make(map[string]string, 2)
	users["张三"]["nickname"] = "小张"
	users["张三"]["pwd"] = "xiaozhang123"

	users["李四"] = make(map[string]string)
	users["李四"]["nickname"] = "小李"
	users["李四"]["pwd"] = "lisi123"

	users["王五"] = make(map[string]string)
	users["王五"]["nickname"] = "小王"
	users["王五"]["pwd"] = "wangwu123"

	fmt.Println(users)
	modifyUser(users, "kp")
	fmt.Println(users)
}
