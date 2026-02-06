package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Protected(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		next(res, req)
	}
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		email := req.FormValue("email")
		password := req.FormValue("password")

		exist, err := UserExists(db, email, password)
		if err != nil {
			log.Println(err)

			http.Redirect(
				res,
				req,
				"/mark/login?error=invalid credintials",
				http.StatusSeeOther,
			)
			return
		}
		if exist {
			jwtSecret := []byte(os.Getenv("JWT_SECRET"))
			if len(jwtSecret) == 0 {
				log.Fatal("some secerts aren't set")
				return
			}
			token, err := createJwt(email, jwtSecret)
			if err != nil {
				log.Println(err)
				http.Redirect(
					res,
					req,
					"/mark/login?error=internal server error",
					http.StatusSeeOther,
				)
				return
			}

			cookie := http.Cookie{
				Name:     "jwt_token",
				Value:    token,
				Path:     "/",
				Expires:  time.Now().Add(15 * time.Minute),
				Secure:   false,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			}

			http.SetCookie(res, &cookie)
			http.Redirect(res, req, "/mark/add", http.StatusSeeOther)
			return
		} else {
			http.Redirect(res, req, "/mark/login?error=invalid credintials", http.StatusSeeOther)
			return
		}
	}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createJwt(email string, secret []byte) (string, error) {
	claims := jwt.MapClaims{
		"sub":   email,
		"email": email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func verifyJwt(tokenString string) (*jwt.RegisteredClaims, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	tkn, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected method : %v", t.Header["alg"])
			}
			return jwtSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := tkn.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}
	return claims, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
