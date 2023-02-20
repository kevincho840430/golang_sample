package models

func (User) TableName() string {
	return "user"
}

type User struct {
	Id     int    `json:"id"`
	UserId string `gorm:"type:varchar(100);not null;index:user_id"`
	Name   string `gorm:"type:varchar(100);not null"`
	//Gender   string `gorm:"type:tinyint;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(50);not null"`
	//DeleteTime int    `gorm:"type:int(10);default:null"`
}

// func (user *User) check()error{
// 	if(user.UserId==""){
// 		error := errors.New("model_struct_error")
// 		return error
// 	}
// 	if(user.Name==nil){
// 		error := errors.New("model_struct_error")
// 	}
// }
