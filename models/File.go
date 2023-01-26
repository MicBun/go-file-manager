package models

import (
	"errors"
	"gorm.io/gorm"
)

// File
type File struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Uploader   string `json:"uploader"`
	Type       string `json:"type"`
}

type FileInterface interface {
	UploadFile(db *gorm.DB, file *File) error
	DownloadFile(db *gorm.DB, file *File) error
}

func NewFile() FileInterface {
	return &File{}
}

func (f *File) UploadFile(db *gorm.DB, file *File) error {
	var count int64
	db.Model(&File{}).Where("name = ? AND uploader = ?", file.Name, file.Uploader).Count(&count)
	if count > 0 {
		return errors.New("file already exists")
	}
	err := db.Create(&file).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *File) DownloadFile(db *gorm.DB, file *File) error {
	err := db.Where("name = ? AND uploader = ?", file.Name, file.Uploader).First(&file).Error
	if err != nil {
		return err
	}
	return nil
}
