package db

import (
	"github.com/glebarez/sqlite"
	"github.com/krau/favpics-helper/internal/models"
	"github.com/krau/favpics-helper/pkg/util"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	util.Log.Info("init db")
	var err error
	db, err = gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
	if err != nil {
		util.Log.Fatal("init db error:", err)
	}
	db.AutoMigrate(&models.Pic{})
}

func SavePics(pics []models.Pic) {
	for _, pic := range pics {
		db.Create(&pic)
	}
}

func IsPicExist(p models.Pic) bool {
	// 查询图片是否存在
	var count int64
	util.Log.Debug("check pic exist:", p.Link)
	db.Model(&models.Pic{}).Where("Link = ?", p.Link).Count(&count)
	return count > 0
}

func AddPics(pics []models.Pic) {
	for _, pic := range pics {
		db.Model(&pic).Create(&pic)
		util.Log.Debugf("add pic to db: %s", pic.Link)
	}
}

func AddPic(pic models.Pic) {
	db.Model(&pic).Create(&pic)
	util.Log.Debugf("add pic to db: %s", pic.Link)
}

func DeletePic(id int) {
	var pic models.Pic
	db.First(&pic, id)
	db.Delete(&pic)
}

func DeletePics(ids []int) {
	for _, id := range ids {
		DeletePic(id)
	}
}

func GetPicCount() int64 {
	var count int64
	db.Model(&models.Pic{}).Count(&count)
	return count
}

func GetPicCountBySource(source string) int64 {
	var count int64
	db.Model(&models.Pic{}).Where("source = ?", source).Count(&count)
	return count
}
