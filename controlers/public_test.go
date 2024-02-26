package controlers_test

import (
	"dcardAssignment/controlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNormalQuery(t *testing.T) {

	r := gin.Default()
	
	r.GET("/public/advertisement", controlers.PublicControler{}.PublicAdvertisement)

	query := "offset=0&limit=10&age=25&gender=M&country=TW&platform=android"

	req, err := http.NewRequest("GET", "/public/advertisement?"+query, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t,response["items"])
}

func TestVaildQuery(t *testing.T) {

	r := gin.Default()
	
	r.GET("/public/advertisement", controlers.PublicControler{}.PublicAdvertisement)

	query := "offset=0&limit=0&age=25&gender=M&country=TW&platform=android"

	req, err := http.NewRequest("GET", "/public/advertisement?"+query, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t,response["Error Message"])
}

func TestMissingQuery(t *testing.T) {

	r := gin.Default()
	
	r.GET("/public/advertisement", controlers.PublicControler{}.PublicAdvertisement)

	query := "offset=0&limit=10&age=25&gender=F&country=TW"

	req, err := http.NewRequest("GET", "/public/advertisement?"+query, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t,response["items"])
}

func TestEmptyQuery(t *testing.T) {

	r := gin.Default()
	
	r.GET("/public/advertisement", controlers.PublicControler{}.PublicAdvertisement)

	query := "offset=0&limit=10"

	req, err := http.NewRequest("GET", "/public/advertisement?"+query, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	
	assert.NotNil(t,response["items"])
}