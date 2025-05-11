package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CookieTtl = time.Minute * 60 // 1 hour
)

var todos []Todo
var loggedInUser string

// Add a new global variable for the secret key
var secretKey = []byte("123")

// Add routes for the ToDo App
func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Todos":    todos,
			"LoggedIn": loggedInUser != "",
			"Username": loggedInUser,
			"Role":     getRole(loggedInUser),
		})
	})

	router.POST("/add", authenticate, func(c *gin.Context) {
		text := c.PostForm("todo")
		todo := Todo{Text: text, Done: false, CreatedAt: time.Now()}
		todos = append(todos, todo)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.POST("/toggle", authenticate, func(c *gin.Context) {
		index := c.PostForm("index")
		toggleIndex(index)
		c.Redirect(http.StatusSeeOther, "/")
	})

	// Create route for user authentication
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Dummy credential check
		if (username == "employee" && password == "password") || (username == "senior" && password == "password") {
			tokenString, err := createToken(username)
			if err != nil {
				c.String(http.StatusInternalServerError, "Error creating token")
				return
			}

			loggedInUser = username
			fmt.Printf("Token created: %s\n", tokenString)
			c.SetCookie("token", tokenString, int(CookieTtl), "/", "localhost", false, true)
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			c.String(http.StatusUnauthorized, "Invalid credentials")
		}
	})

	router.GET("/logout", func(ctx *gin.Context) {
		loggedInUser = ""
		ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
		ctx.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
