package Utils

import (
	"NFUShop/Config"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const EmptyString = ""

func GetJSONStringArray(jsonString string) []string {
	var ret []string
	json.Unmarshal([]byte(jsonString), &ret)
	return ret
}
func StringArrayToJSON(strArray []string) string {
	if bytes, err := json.Marshal(strArray); err == nil {
		return string(bytes)
	}
	return EmptyString
}

func StrToInt(str string) int {
	var ret int
	ret = 0
	if data, err := strconv.Atoi(str); err == nil {
		ret = data
	}
	return ret
}

type Claims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

func GenerateJWT(obj interface{}) string {
	claims := Claims{
		obj,
		jwt.StandardClaims{ExpiresAt: time.Now().Add(148 * time.Hour).Unix()},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString([]byte(Config.GetJWTSecret()))
	return token
}

func ParseJWT(token string) interface{} {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.GetJWTSecret()), nil
	})
	if tokenClaims != nil && err == nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims.Data
		}
	}
	return nil
}
