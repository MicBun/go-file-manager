package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/MicBun/go-file-manager/config"
	"github.com/MicBun/go-file-manager/middleware"
	"github.com/MicBun/go-file-manager/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	t.Setenv("DB_HOST", "tcp(localhost:3306)")
	r := gin.Default()
	fileRoutes := r.Group("/file")
	fileRoutes.Use(middleware.JwtAuthMiddleware())
	db := config.ConnectDataBase()
	fileRoutes.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	fileRoutes.POST("/upload", UploadFile)

	// Create a test file
	file, err := os.Create("testfile.png")
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}
	defer file.Close()

	// Create a test multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "testfile.png")
	if _, err := io.Copy(part, file); err != nil {
		t.Errorf("Error copying test file to form: %v", err)
	}
	writer.Close()

	// Create a test request
	req, _ := http.NewRequest("POST", "/file/upload", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// Create a test recorder
	w := httptest.NewRecorder()

	// Do Login first then get the token
	token := getTokenByLogin(t)
	req.Header.Add("Authorization", "Bearer "+token)

	// Perform the test
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// close the test file
	err = file.Close()

	// Delete the test file
	err = os.Remove("testfile.png")
	if err != nil {
		t.Errorf("Error deleting test file: %v", err)
	}
}

func getTokenByLogin(t *testing.T) string {
	t.Setenv("DB_HOST", "tcp(localhost:3306)")
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/login", nil)

	database := config.ConnectDataBase()
	ctx.Set("db", database)

	user := models.User{
		Username: "user1",
		Password: "password1",
	}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	ctx.Request = req
	Login(ctx)
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	var response map[string]string
	body, _ := io.ReadAll(w.Body)
	_ = json.Unmarshal(body, &response)

	return response["token"]
}
