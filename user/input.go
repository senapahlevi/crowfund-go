package user

// urutan
// 	1input user input
// 	2.handler hasil dari input akan dimapping ke struct
// 	3.service mapping struct ke struct Model misal User
// 	4.repository save ke struct ke db
// 	5.db db mysql menerima dari golang

type RegisterUserInput struct {
	Name       string
	Occupation string
	Email      string
	Password   string
}
