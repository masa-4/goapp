package src

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID       int
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Email    string `form:"email" binding:"required" gorm:"unique;not null"`
	Password []byte `form:"password" binding:"required" gorm:"unique;not null"`
}

func InsertUser(user *User) {
	db := gormConnect()
	db.Create(&user)
	defer db.Close()
}

func DeleteUser(userID int) {
	user := []User{}
	db := gormConnect()
	db.Delete(&user, userID)
	defer db.Close()
}

func GetUserAll() []User {
	var users []User
	db := gormConnect()
	db.Order("ID asc").Find(&users)
	defer db.Close()

	return users
}

func GetUserFromEmail(posted_email string) User {
	var user User
	db := gormConnect()
	db.Where("email = ?", posted_email).First(&user)
	return user
}
