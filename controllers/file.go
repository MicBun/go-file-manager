package controllers

import (
	"fmt"
	"github.com/MicBun/go-file-manager/models"
	jwtauth "github.com/MicBun/go-file-manager/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type ListFileOutput struct {
	models.File
	Path string `json:"path"`
}

func ListFile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var files []models.File
	err := db.Find(&files).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	output := make([]ListFileOutput, 0)
	for _, file := range files {
		output = append(output, ListFileOutput{
			File: file,
			Path: fmt.Sprintf("/download?uploader=%v&filename=%v", file.Uploader, file.Name),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"files": output,
	})
}

func UploadFile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	username := func() string {
		claims, _ := jwtauth.ExtractTokenID(c)
		username, _ := models.NewUser().GetUsernameByID(db, claims)
		return username
	}()

	// Check file type
	if !isValidFileType(file) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
		return
	}

	// Store the file on the local storage
	filename := generateFilename()
	// create folder with username name on assets folder
	os.Mkdir("./assets/"+username, 0755)
	// create file with filename on assets folder
	path := "./assets/" + username + "/" + filename
	localFile, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer localFile.Close()

	fileOpen, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer fileOpen.Close()

	io.Copy(localFile, fileOpen)

	// Save file to database
	fileModel := models.NewFile().UploadFile(db,
		&models.File{
			Name:     filename,
			Uploader: username,
			Type:     file.Header.Get("Content-Type"),
		},
	)

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": fileModel})
}

func generateFilename() string {
	t := time.Now()
	return t.Format("20060102150405") + ".png"
}

func isValidFileType(file *multipart.FileHeader) bool {
	validFileTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"video/mp4":  true,
		"video/webm": true,
	}

	fileType := file.Header.Get("Content-Type")
	return validFileTypes[fileType]
}

func DownloadFile(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)
	//var file models.File
	uploader := c.Query("uploader")
	filename := c.Query("filename")

	// Check if file exists in assets folder
	if _, err := os.Stat("./assets/" + uploader + "/" + filename); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// download file
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%v", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./assets/" + uploader + "/" + filename)
}
