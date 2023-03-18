package authjwt

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Decode(tokenString string) (*jwt.Token, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("AUTH_SECRET")), nil
	})

	return token, err
}

func Encode(exp int64, userID uint) (string, error){
	claims := jwt.MapClaims{
		"exp": exp,
		"user_id": userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	return tokenString, err
}
