package api_test

import (
	"GittSwap/pkg/api"
	"GittSwap/pkg/schema"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	schema.TestDBInit()
	return router
}

func TestListFamilyMembers(t *testing.T) {
	router := setUpRouter()

	router.GET("/members", api.ListMembers)
	req, _ := http.NewRequest("GET", "/members", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestAddMemberHandler(t *testing.T) {
	r := setUpRouter()
	r.POST("/members", api.AddMember)
	familyMemberId := 1
	member := schema.FamilyMember{
		ID:   uint(familyMemberId),
		Name: "Demo Family Member",
	}
	jsonValue, _ := json.Marshal(member)
	req, _ := http.NewRequest("POST", "/members", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestListFamilyMembersWithData(t *testing.T) {
	r := setUpRouter()

	// Adding data for Family Member
	r.POST("/members", api.AddMember)
	familyMemberId := 1
	member := schema.FamilyMember{
		ID:   uint(familyMemberId),
		Name: "Demo Family Member",
	}
	jsonValue, _ := json.Marshal(member)
	req, _ := http.NewRequest("POST", "/members", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the added data
	r.GET("/members", api.ListMembers)
	req2, _ := http.NewRequest("GET", "/members", nil)
	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req2)

	assert.Equal(t, http.StatusOK, w1.Code)
	assert.NotEmpty(t, w1.Body.Bytes(), "Empty response received")
}
