package helper

// format nya seperti ini sesuai di video
// 	"meta" : {
// 		"code":
// 		"message":  ""
// 		"status":  ""
// 	},
// 	"data" : {
// 		"name" : ,
// 		"occupation" : ,
// 		"token" : ,
// 		"email" : ,dll
// 	}
type Response struct {
	Meta Meta        `json:"meta"` // biar huruf kecil sesuai video
	Data interface{} `json:"data"` // biar huruf kecil sesuai video dan interface supaya flexibel mau "data" :{[]} juga bisa
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {

	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}
