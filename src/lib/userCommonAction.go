package lib

import (
	"errors"

	"github.com/blog/src/common"
	"github.com/blog/src/global"
	"github.com/blog/src/model"
	"gorm.io/gorm"
)

func Login(name, password string) (*LoginRespInfo, error) {
	author, err := getAuthorByName(name)
	if err != nil {
		return nil, err
	}

}

func getAuthorByName(name string) (*model.Author, error) {
	author := model.Author{Name: name}
	result := global.GetDB().First(&author)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &common.AuthorNameRepeatError
	}
	return &author, nil
}
