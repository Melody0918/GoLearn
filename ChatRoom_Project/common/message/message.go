package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息
}

//定义两种消息(登陆消息+登陆返回消息)..后面需要再增加

type LoginMes struct {
	UserId   string `json:"userId"`   //用户Id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回错误码 500表示该用户未注册 200表示登陆成功
	Error string `json:"error"` // 返回错误信息
}
