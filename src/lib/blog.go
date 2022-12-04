package lib

import (
	"log"

	"github.com/blog/src"
	"github.com/blog/src/model"
)

func GetAllBlog() []model.Blog {
	var blogs []model.Blog
	result := src.DB.Find(&blogs)
	if result.Error != nil {
		log.Fatal("get blogs error")
	}
	return blogs
}
