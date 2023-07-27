package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func Login(c *gin.Context) {
	// 获取用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 验证用户名和密码
	m := md5.New()
	m.Write([]byte(username + "qqzl" + password))
	hm := hex.EncodeToString(m.Sum(nil))
	fmt.Println(hm)
	if hm != "bcfc3329cc0847c9c8289cb7ce9ab824" {
		c.JSON(http.StatusOK, gin.H{
			"token": "",
			"code":  "0",
		})
	} else {

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = username
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
