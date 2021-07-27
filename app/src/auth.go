package src

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SessionInfo struct {
	UserName interface{}
}

func getlogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func postlogin(c *gin.Context) {
	user_name := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	password_from_db := GetUserFromEmail(email).Password
	hashd_password := []byte(password_from_db)
	err := bcrypt.CompareHashAndPassword(hashd_password, []byte(password))
	if err != nil {
		log.Println("failed login")
		c.Redirect(403, "/login")
	} else {
		log.Println("successed login")
		login(c, user_name)
		c.Redirect(302, "/menu/top")
	}
}

func login(c *gin.Context, UserName string) {
	session := sessions.Default(c)
	session.Set("UserName", UserName)
	session.Save()
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func postLogout(c *gin.Context) {
	logout(c)
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		var session_info SessionInfo

		session := sessions.Default(c)
		session_info.UserName = session.Get("UserName")

		if session_info.UserName == nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		} else {
			c.Set("UserName", session_info.UserName)
			c.Next()
		}
		log.Print("finished session check")
	}
}
