package midleware

import (
	"go_Muhammad-Wahyu-Yudiansyah/22_Midleware/Praktikum/Alltugas/constant"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(UserId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = UserId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}
