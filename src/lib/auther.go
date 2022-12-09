package lib

import (
	"errors"
	"log"

	"github.com/blog/src/common"
	"github.com/blog/src/global"
	"github.com/blog/src/model"
	"gorm.io/gorm"
)

func CreateAuthor(name, desc string) (uint, error) {
	if is_repeat := checkAuthorNameRepeat(name); is_repeat {
		return 0, &common.AuthorNameRepeatError
	}
	author := model.Author{Name: name, Desc: desc}
	result := global.DB.Create(&author)
	if result.Error != nil {
		log.Fatalf("author create error: %v", result.Error)
	}
	return author.ID, nil
}

func checkAuthorNameRepeat(name string) bool {
	author := model.Author{Name: name}
	err := global.DB.First(&author).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func GetAuthors() []model.Author {
	var authorList []model.Author
	result := global.DB.Find(&authorList)
	if result.Error != nil {
		log.Fatalf("find author error: %v", result.Error)
	}
	return authorList
}

func UpdateAuthor(id uint, name, desc string) uint {
	var author model.Author
	result := global.DB.First(&author, id)
	if result.Error != nil {
		log.Fatalf("find author error %v", result.Error)
	}

	author.Name = name
	author.Desc = desc
	global.DB.Save(&author)
	return id
}
