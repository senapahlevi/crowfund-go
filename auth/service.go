package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(UserID int) (string, error)
	ValidateToken(Token string) (*jwt.Token, error)
	// return selain error juga *jwt.Token karena buat pakai method bawaan package jwt.token yang mana
	//balikan nya juga jwt token
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SecretKey = []byte("Tokone123")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err

	}
	return signedToken, nil
}

//validate token itu untuk validasi apakah bener hasil generate jwt yang panjang itu dibuat oleh secret key kita ?
//kalo bukan ya auto ditolak karena jwt bisa dibikin siapa aja dengan payload yang sama yakni misal user_id = 28

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok { //jika method nya gasama disini casenya gw make HS256
			return nil, errors.New("invalid token")
		}
		return []byte(SecretKey), nil //jika method sama dan secret key nya sama maka ini valid lancar semua
	})
	if err != nil {
		return token, err
	}
	return token, nil //jika sukses
}
