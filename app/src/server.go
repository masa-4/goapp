package src

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	dbInit()
	router.LoadHTMLGlob("../app/views/*")
	router.GET("/", home)
	router.GET("/fish_index", fish_index)
	router.GET("/fish_register", fish_register)
	router.POST("/create", create)
	router.POST("/fish_edit/:id", edit_fish)
	router.POST("/update/:id", updatefish)
	router.POST("/fish_delete/:id", fish_delete)
	router.Run()
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "top.html", gin.H{})
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
