package helper

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key") 

func GenerateJWT(userId int, email string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userId,
        "email":   email,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return nil, err
    }
    return token.Claims.(jwt.MapClaims), nil
}