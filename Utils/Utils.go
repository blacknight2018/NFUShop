package Utils

import (
	"NFUShop/Config"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const EmptyString = ""

type Time time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func JSON2IntArray(jsonString string) []int {
	var ret []int
	json.Unmarshal([]byte(jsonString), &ret)
	return ret
}

func Any2JSON(in interface{}) string {
	var ret string
	if bytes, err := json.Marshal(in); err == nil {
		ret = string(bytes)
	}
	return ret
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

/**
 * @Description: 将数据解析成JWTToken
 * @param obj
 * @return string
 */
func GenerateJWT(obj interface{}) string {
	claims := Claims{
		obj,
		jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Duration(Config.GetTokenValidTime()) * time.Second).Unix()},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString([]byte(Config.GetJWTSecret()))
	return token
}

/**
 * @Description: 将JWTToken解析为对象
 * @param token
 * @return interface{}
 */
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

func UploadImg(base64Data string) string {
	var ret string
	type name struct {
		Data string `json:"data"`
	}
	var tmp name
	tmp.Data = base64Data
	defaultClient := http.DefaultClient
	request, err := http.NewRequest("POST", Config.GetImgServer()+"/img", strings.NewReader(Any2JSON(tmp)))
	if err != nil {
		return ret
	}
	request.Header.Set("Content-Type", "Application/json")
	resp, err := defaultClient.Do(request)
	if err != nil {
		return ret
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	ret = string(bytes)
	return ret
}

func ContextGetInt(context *gin.Context, key string) int {
	var ret int
	if v, ok := context.Get(key); ok {
		switch data := v.(type) {
		case int:
			ret = data
			break
		case int64:
			ret = int(data)
			break
		case float64:
			ret = int(data)
			break
		}
	}
	return ret
}
func ContextQueryInt(context *gin.Context, key string) int {
	data := context.Query(key)
	ret, _ := strconv.Atoi(data)
	return ret
}

func Int2IntPtr(v int) *int {
	return &v
}

func Bool2BoolPtr(v bool) *bool {
	return &v
}
