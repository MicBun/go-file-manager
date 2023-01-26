package route

import (
	"gorm.io/gorm"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	db := &gorm.DB{}
	r := SetupRouter(db)

	// Test that the router has the correct routes
	//pokemonRoute := r.Routes()[0]
	loginRoute := r.Routes()[0]
	if loginRoute.Method != "POST" || loginRoute.Path != "/login" {
		t.Errorf("Expected POST /login, got %s %s", loginRoute.Method, loginRoute.Path)
	}

	//findRoute := r.Routes()[1]
	resetRoute := r.Routes()[1]
	if resetRoute.Method != "POST" || resetRoute.Path != "/resetUserDatabase" {
		t.Errorf("Expected POST /resetUserDatabase, got %s %s", resetRoute.Method, resetRoute.Path)
	}

	uploadRoute := r.Routes()[2]
	if uploadRoute.Method != "POST" || uploadRoute.Path != "/file/upload" {
		t.Errorf("Expected POST /file/upload, got %s %s", uploadRoute.Method, uploadRoute.Path)
	}

	// check length of routes
	if len(r.Routes()) < 0 {
		t.Errorf("Expected routes, got %d", len(r.Routes()))
	}
}
