package model

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model `json:"-"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorID   int    `json:"-"`
	Author     Author `gorm:"constraint:OnDelete:SET NULL" json:"author"`
	Tag        []*Tag `gorm:"many2many:tag_blog_rel" json:"tag"`
}

type Tag struct {
	gorm.Model `json:"-"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Blogs      []*Blog `gorm:"many2many:tag_blog_rel" json:"-"`
}

type Author struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"uniqueIndex;size:50" json:"name"`
	Desc       string `json:"desc"`
	Password   string `gorm:"size:64" json:"-"`
}
