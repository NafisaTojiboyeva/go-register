package models

type User struct {
	Id int				`json:"id" gorm:"column:user_id";"auto_increment:true"` 
	Fullname string		`json:"fullName"` 
	Phone string		`json:"phone"`
	Password string		`json:"password"`
	SmsConfirm string	`json:"smsConfirm"`
	IsVerified bool		`json:"-"`
}

type RegisterInfo struct {
	Fullname string 	`json:"fullName"` 
	Phone string		`json:"phone"`
	Password string		`json:"password"`
}

type LoginInfo struct {
	Phone string		`json:"phone"`
	Password string		`json:"password"`
	// SmsConfirm string	`json:"smsConfirm"`
}

type ChangePasswordInfo struct {
	Phone string		`json:"phone"`
	Password string		`json:"password"`
	SmsConfirm string	`json:"smsConfirm"`
}

type Code struct {
	Code string			`json:"code"`
}

type SMSRequestBody struct {
	APIKey string		`json:"api_key"`
	APISecret string	`json:"api_secret"`
	From string			`json:"from"`
	To string			`json:"to"`
	Text string			`json:"text"`
}