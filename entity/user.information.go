package entity

// UserInformation 用户信息
type UserInformation struct {
	Id      int    `json:"id"`      // 用户 ID
	Account string `json:"account"` // 用户账号
	Name    string `json:"name"`    // 用户姓名
	Balance string `json:"balance"` // 用户余额
	Token   string `json:"token"`   // 用户令牌
}
