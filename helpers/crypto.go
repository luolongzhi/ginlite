package helpers

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//bcrypt hash and valid
func HashPassword(rawPassword string) string {
	password := []byte(rawPassword)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(hashedPassword))

	return string(hashedPassword)
}

func VerifyPassword(rawPassword string, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))
}

//jwt
const JwtSecret = "kZtqItd8GLGtHQ92J5Qmt2QyZrHKymeI"
const JwtExpiresSecond = 7 * 24 * 60 * 60

type JwtClaim struct {
	Env       string
	UserID    uint
	AccountID string
	Timestamp uint
	ExpiredAt uint
}

func GenerateJwt(env string, userID uint, accountID string) (string, error) {
	t := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"env":        env,
		"user_id":    userID,
		"account_id": accountID,
		"timestamp":  t.Unix(),
		"expired_at": t.Unix() + JwtExpiresSecond,
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(JwtSecret))
}

func VerifyJwt(tokenString string, claim *JwtClaim) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JwtSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["timestamp"], claims["user_id"], claims["account_id"])

		claim.Env = claims["env"].(string)
		claim.UserID = claims["user_id"].(uint)
		claim.AccountID = claims["account_id"].(string)
		claim.Timestamp = claims["timestamp"].(uint)
		claim.ExpiredAt = claims["expired_at"].(uint)

		return nil
	} else {
		fmt.Println(err)
		return err
	}
}
