package lib

import (
	"errors"

	"github.com/blog/src/common"
	"github.com/blog/src/global"
	"github.com/blog/src/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(name, password string) (*LoginRespInfo, error) {
	author, err := getAuthorByName(name)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(password))
	if err != nil {
		return nil, &common.WrongPassword
	}

	token := CreateToken()
	authorBaseInfo := AuthorBaseInfo{
		Name: author.Name,
		ID:   author.ID,
		Desc: author.Desc,
	}
	token.Set(&authorBaseInfo, TokenExpiration)
	return &LoginRespInfo{
		Token:          string(token),
		AuthorBaseInfo: &authorBaseInfo,
	}, nil
}

func getAuthorByName(name string) (*model.Author, error) {
	author := model.Author{}
	result := global.GetDB().Where("name = ?", name).First(&author)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &common.RecordNotExist
		}
		panic(result.Error)
	}
	return &author, nil
}

func Logout(token string) {
	t := Token(token)
	t.Delete()
}
