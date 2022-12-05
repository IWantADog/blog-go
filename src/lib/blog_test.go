package lib

import (
	"testing"

	"github.com/blog/global/global"
	"github.com/blog/global/model"
)

func TestGetBlogs(t *testing.T) {
	auther := model.Author{
		Name: "auther-1",
		Desc: "xxxxxx",
	}
	global.DB.Create(&auther)

	blog := model.Blog{
		Title:    "xxxx",
		Content:  "xxxx",
		AuthorID: int(auther.ID),
	}
	result := global.DB.Create(&blog)
	if result.Error != nil {
		panic(result.Error)
	}

	res := GetAllBlog()
	if len(res) != 1 {
		t.Fatal("get blog error")
	}
	resFirst := res[0]
	if resFirst.Title != "xxxx" {
		t.Fatal("get error data")
	}
}
