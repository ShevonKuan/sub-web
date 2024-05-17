package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTConfig struct {
	SigningKey []byte
	ExpiresAt  time.Duration
}

var jwtConfig = &JWTConfig{
	SigningKey: []byte("your-secret-key"),
	ExpiresAt:  24 * time.Hour,
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	// 获取用户名和密码
	user := User{}
	c.ShouldBind(&user)
	// 验证用户名和密码
	m := md5.New()

	io.WriteString(m, user.Username+"qqzl"+user.Password)
	hm := hex.EncodeToString(m.Sum(nil))
	if hm != "6eb0cc571ceb8be8e8785f51959c29aa" {
		c.JSON(http.StatusOK, gin.H{
			"token": "",
			"code":  "0",
		})
	} else {

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(jwtConfig.ExpiresAt).Unix()
		tokenString, _ := token.SignedString(jwtConfig.SigningKey)
		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
			"code":  "1",
		})
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "未登录",
			})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtConfig.SigningKey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token无效",
			})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token无效",
			})
			return
		}
		c.Next()
	}
}
