package src

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title    string
	Content  string
	AuthorID int
	Author   Author `gorm:"constraint:OnDelete:SET NULL"`
	Tag      []*Tag `gorm:"many2many:tag_blog_rel"`
}

type Tag struct {
	gorm.Model
	Name  string
	Desc  string
	Blogs []*Blog `gorm:"many2many:tag_blog_rel"`
}

type Author struct {
	gorm.Model
	Name string
	Desc string
}
