package src

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func StartServer() {
	router := gin.Default()
	dbInit()
	router.LoadHTMLGlob("../app/views/*")

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth_session", store))

	router.GET("/", home)

	router.GET("/signup", signup)
	router.POST("/signup", user_create)

	router.GET("/user_index", user_index)

	router.GET("/login", getlogin)
	router.POST("/login", postlogin)

	router.GET("/fish_index", fish_index)

	router.GET("/fish_register", fish_register)

	router.POST("/create", create)

	router.POST("/fish_edit/:id", edit_fish)

	router.POST("/update/:id", updatefish)

	router.POST("/fish_delete/:id", fish_delete)

	menu := router.Group("/menu")
	menu.Use(sessionCheck())
	{
		menu.GET("/top", getMenu)
	}
	router.POST("/logout", postLogout)

	router.Run()
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "top.html", gin.H{})
}

func signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func fish_index(c *gin.Context) {
	all_fish := GetFishAll()
	c.HTML(http.StatusOK, "fish_index.html", gin.H{
		"fish": all_fish,
	})
}

func fish_register(c *gin.Context) {
	c.HTML(http.StatusOK, "fish_register.html", gin.H{})
}

func create(c *gin.Context) {
	name := c.PostForm("name")
	origin := c.PostForm("origin")
	registered_fish := Fish{
		Name:   name,
		Origin: origin,
	}
	InsertFish(&registered_fish)

	all_fish := GetFishAll()
	c.HTML(http.StatusPermanentRedirect, "fish_index.html", gin.H{
		"fish": all_fish,
	})
}

func edit_fish(c *gin.Context) {
	id := c.Param("id")
	fishid, _ := strconv.Atoi(id)
	fish_for_edit := GetFishforEdit(fishid)
	c.HTML(http.StatusOK, "fish_edit.html", gin.H{
		"ID":     fish_for_edit.ID,
		"Name":   fish_for_edit.Name,
		"Origin": fish_for_edit.Origin,
	})
}

func updatefish(c *gin.Context) {
	id := c.Param("id")
	fish_id, _ := strconv.Atoi(id)
	name := c.PostForm("name")
	origin := c.PostForm("origin")
	changestatus := map[string]string{"Name": name, "Origin": origin}
	EditFish(fish_id, changestatus)

	c.Redirect(http.StatusFound, "/fish_index")
}

func fish_delete(c *gin.Context) {
	id := c.Param("id")
	fishID, _ := strconv.Atoi(id)

	DeleteFish(fishID)
	c.Redirect(http.StatusFound, "/fish_index")
}

func user_create(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	hassedpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	created_user := User{
		Username: username,
		Email:    email,
		Password: hassedpassword,
	}
	InsertUser(&created_user)
	c.Redirect(http.StatusFound, "/user_index")
}

func user_index(c *gin.Context) {
	all_users := GetUserAll()
	c.HTML(http.StatusOK, "user_index.html", gin.H{
		"users": all_users,
	})
}

func getMenu(c *gin.Context) {
	c.HTML(http.StatusAccepted, "user_menu.html", gin.H{})
}
