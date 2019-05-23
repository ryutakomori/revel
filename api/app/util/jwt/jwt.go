package jwt

import (
	db "api/app/util/db"
	"api/app/util/define"
	"api/app/util/msg"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	dbmodels "xo/models"

	"github.com/jinzhu/gorm"
)

var (
	secretKey interface{}
	publicKey interface{}
)

func Initialize() {
	rawSecretKey, _ := ioutil.ReadFile("keys/revel.pri")
	rawPublicKey, _ := ioutil.ReadFile("keys/revel.pub.pkcs8")

	err := ParseKeys(rawSecretKey, rawPublicKey)
	if err != nil {
		fmt.Println(err)
	}
}

func ParseKeys(rawSK []byte, rawPK []byte) error {
	privateKeyBlock, _ := pem.Decode(rawSK)
	if privateKeyBlock == nil {
		return errors.New("Private key cannot decode")
	}
	if privateKeyBlock.Type != "RSA PRIVATE KEY" {
		return errors.New("Private key type is not rsa")
	}
	secretKey, _ = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)

	publicKeyBlock, _ := pem.Decode(rawPK)
	if publicKeyBlock == nil {
		return errors.New("Public key cannot decode")
	}
	if publicKeyBlock.Type != "PUBLIC KEY" {
		return errors.New("Public key type is invalid")
	}
	publicKey, _ = x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)

	return nil
}
func Verify(token string) (string, error) {
	tokenString := strings.Replace(token, "Bearer ", "", 1)

	parsedToken, err := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwtgo.SigningMethodRSA); !ok {
			return "", errors.New("Unexpected signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		err = errors.Wrap(err, "Token is invalid")
		return "", err
	}
	if !parsedToken.Valid {
		return "", errors.New("Token is invalid")
	}

	return tokenString, nil
}

func TokenGenerate(userID string, now int64) (tokenString string, err error) {

	claims := jwtgo.StandardClaims{
		Subject:  userID,
		IssuedAt: now,
		//	ExpiresAt: time.Unix(now, 0).AddDate(0, 0, expiryDays).Unix(),
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)

	// sign by secret key
	tokenString, err = token.SignedString(secretKey)

	if err != nil {
		err = errors.Wrap(err, "Failed to sign token")
		return
	}
	return tokenString, nil
}

func UserAuth(gorm *gorm.DB, tokenStr string) (int, dbmodels.User, string) {
	user := dbmodels.User{}

	token, err := Verify(tokenStr)
	if err != nil {
		return http.StatusUnauthorized, user, msg.Tpl("Jwt", "認証エラーです。")
	}

	errors := gorm.Where("token = ?", token).Where("is_deleted = false").First(&user).GetErrors()
	if db.Handler(errors) == define.DB_SQL_FAILURE {
		return http.StatusInternalServerError, user, ""
	}
	if user.UserID == 0 {
		return http.StatusUnauthorized, user, msg.Tpl("User", "存在しないユーザーです。")
	}
	return 0, user, ""
}
