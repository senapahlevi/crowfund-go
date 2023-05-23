package user

// urutan
// 	1input user input
// 	2.handler hasil dari input akan dimapping ke struct
// 	3.service mapping struct ke struct Model misal User
// 	4.repository save ke struct ke db
// 	5.db db mysql menerima dari golang

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
