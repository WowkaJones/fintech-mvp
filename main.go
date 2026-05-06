package main

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

var users = map[string]string{}
var balance = map[string]float64{}

func main() {
    r := gin.Default()

    r.POST("/register", register)
    r.POST("/login", login)
    r.POST("/add-income", auth(), addIncome)

    r.Run(":8080")
}

func register(c *gin.Context) {
    email := c.PostForm("email")
    pass := c.PostForm("password")
    hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
    users[email] = string(hash)
    balance[email] = 0
    c.JSON(200, gin.H{"ok": true})
}

func login(c *gin.Context) {
    email := c.PostForm("email")
    pass := c.PostForm("password")
    hash := users[email]
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
    if err != nil {
        c.JSON(401, gin.H{"error": "wrong password"})
        return
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour).Unix(),
    })
    tokenString, _ := token.SignedString([]byte("secret"))
    c.JSON(200, gin.H{"token": tokenString})
}

func auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.GetHeader("Authorization")
        token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        })
        if err != nil {
            c.JSON(401, gin.H{"error": "invalid token"})
            c.Abort()
            return
        }
        claims := token.Claims.(jwt.MapClaims)
        c.Set("email", claims["email"])
        c.Next()
    }
}

func addIncome(c *gin.Context) {
    email, _ := c.Get("email")
    var req struct {
        Amount float64 `json:"amount"`
    }
    c.BindJSON(&req)
    balance[email.(string)] += req.Amount
    c.JSON(200, gin.H{"balance": balance[email.(string)]})
}