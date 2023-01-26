package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/MicBun/go-file-manager/config"
	"github.com/MicBun/go-file-manager/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Setenv("DB_HOST", "tcp(localhost:3306)")
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/login", nil)

	database := config.ConnectDataBase()
	ctx.Set("db", database)

	// test valid login
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
	assert.NotEmpty(t, response["token"])

	// test invalid login
	user.Password = "wrongpassword"
	jsonValue, _ = json.Marshal(user)
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	ctx.Request = req
	Login(ctx)
	assert.Equal(t, http.StatusUnauthorized, ctx.Writer.Status())

	// test invalid json
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	ctx.Request = req
	Login(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}
