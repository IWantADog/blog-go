package lib

import (
	"errors"
	"log"

	"github.com/blog/src/common"
	"github.com/blog/src/global"
	"github.com/blog/src/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAuthor(name, password, desc string) (*model.Author, error) {
	if is_repeat := checkAuthorNameRepeat(name); is_repeat {
		return nil, &common.AuthorNameRepeatError
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	author := model.Author{Name: name, Desc: desc, Password: string(hashPassword)}
	result := global.GetDB().Create(&author)
	if result.Error != nil {
		log.Fatalf("author create error: %v", result.Error)
	}
	return &author, nil
}

func checkAuthorNameRepeat(name string) bool {
	author := model.Author{}
	err := global.GetDB().Where("name = ?", name).First(&author).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func GetAuthors() []model.Author {
	var authorList []model.Author
	result := global.GetDB().Find(&authorList)
	if result.Error != nil {
		log.Fatalf("find author error: %v", result.Error)
	}
	return authorList
}

func UpdateAuthor(id uint, name, desc string) uint {
	var author model.Author
	db := global.GetDB()
	result := db.First(&author, id)
	if result.Error != nil {
		log.Fatalf("find author error %v", result.Error)
	}

	author.Name = name
	author.Desc = desc
	db.Save(&author)
	return id
}
