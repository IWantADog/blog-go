package lib

import (
	"log"

	"github.com/blog/src/global"
	"github.com/blog/src/model"
)

func GetAllBlog() []model.Blog {
	var blogs []model.Blog
	result := global.GetDB().Find(&blogs)
	if result.Error != nil {
		log.Fatal("get blogs error")
	}
	return blogs
}
