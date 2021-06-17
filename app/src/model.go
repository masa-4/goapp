package src

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Fish struct {
	gorm.Model
	Name   string
	Origin string
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "goapp"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbInit() {
	db := gormConnect()

	defer db.Close()
	db.AutoMigrate(&Fish{})
}

func GetFishAll() []Fish {
	db := gormConnect()

	defer db.Close()
	var fish []Fish

	db.Order("ID asc").Find(&fish)
	return fish
}

func GetFishforEdit(fishID int) Fish {
	db := gormConnect()
	defer db.Close()

	var fish Fish
	db.First(&fish, fishID)
	return fish
}

func InsertFish(fish *Fish) {
	db := gormConnect()

	defer db.Close()
	db.Create(&fish)
}

func DeleteFish(fishID int) {
	fish := []Fish{}
	db := gormConnect()

	db.Delete(&fish, fishID)
	defer db.Close()
}

func EditFish(fishID int, changedStatus map[string]string) {
	fish := []Fish{}
	db := gormConnect()
	db.Model(&fish).Where(fishID).Update(map[string]string{"Name": changedStatus["Name"], "Origin": changedStatus["Origin"]})
	defer db.Close()
}
