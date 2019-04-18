package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

type LoginResMes struct {
	Code  int    `json:"code"` //返回的状态码500,404等
	Error string `json:"error"`
}

type RegisterMes struct {
	//...
}
