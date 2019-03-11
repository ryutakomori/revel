package hash

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

func HashFromScrypt(str string) string {
	salt := []byte("km126z6zwh")
	converted, _ := scrypt.Key([]byte(str), salt, 16384, 8, 1, 32)
	return hex.EncodeToString(converted[:])
}

func HashFromBcrypt(str string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(str), 10)
	return string(converted)
}

// func main() {
// 	fmt.Println(HashFromScrypt("boee"))
// 	fmt.Println(HashFromBcrypt("boee"))
//}
