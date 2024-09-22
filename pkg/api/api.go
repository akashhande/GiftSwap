package api

import (
	"GittSwap/pkg/schema"
	"GittSwap/pkg/security"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	// Define API endpoints
	r.GET("/members", ListMembers)
	r.GET("/members/:id", getMember)
	r.POST("/members", AddMember)
	r.PUT("/members/:id", updateMember)
	r.DELETE("/members/:id", deleteMember)
	r.GET("/gift_exchange", getGiftExchange)
}

func ListMembers(c *gin.Context) {
	var members []schema.FamilyMember
	schema.DB.Find(&members)

	c.JSON(http.StatusOK, gin.H{"data": members})
}

func getMember(c *gin.Context) {
	var member schema.FamilyMember
	if err := schema.DB.Where("id = ?", c.Param("id")).First(&member); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}

func AddMember(c *gin.Context) {
	var newMember schema.FamilyMember
	if err := c.ShouldBindJSON(&newMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member := schema.FamilyMember{
		Name: newMember.Name}
	schema.DB.Create(&member)

	c.JSON(http.StatusCreated, gin.H{"data": member})
}

func updateMember(c *gin.Context) {
	var member schema.FamilyMember
	if err := schema.DB.Where("id = ?", c.Param("id")).First(&member); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input schema.FamilyMember
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schema.DB.Model(&member).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": member})
}

func deleteMember(c *gin.Context) {
	// Extract the username and password from the Basic Auth header
	username, password, _ := c.Request.BasicAuth()

	// Check if the provided credentials are valid
	if !security.IsValidCredentials(username, password) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Proceed with the delete operation if credentials are valid
	var member schema.FamilyMember
	if err := schema.DB.Where("id = ?", c.Param("id")).First(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	schema.DB.Delete(&member)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func getGiftExchange(c *gin.Context) {
	// Retrieve all members from the database
	var members []schema.FamilyMember
	schema.DB.Find(&members)

	// Shuffle members randomly
	rand.Shuffle(len(members), func(i, j int) {
		members[i], members[j] = members[j], members[i]
	})

	// Retrieve all existing gift exchanges (optional for efficiency)
	var existingExchanges []schema.GiftExchange
	schema.DB.Find(&existingExchanges)

	assignGifts(members, existingExchanges)

}

func assignGifts(members []schema.FamilyMember, existingExchanges []schema.GiftExchange) ([]schema.GiftExchange, error) {
	giftExchanges := []schema.GiftExchange{}
	seenPairs := map[string]bool{}

	for len(members) > 0 {
		giver := members[rand.Intn(len(members))]
		members = remove(members, giver)

		receiver := members[rand.Intn(len(members))]
		members = remove(members, receiver)

		// Check for self-assignment
		if giver.ID == receiver.ID {
			members = append(members, giver, receiver)
			continue
		}

		// Check for duplicate pairs
		key := fmt.Sprintf("%d-%d", giver.ID, receiver.ID)
		if seenPairs[key] {
			members = append(members, giver, receiver)
			continue
		}
		seenPairs[key] = true

		// Create a new gift exchange
		giftExchanges = append(giftExchanges, schema.GiftExchange{
			AssignerID:  giver.ID,
			RecipientID: receiver.ID,
		})

		schema.DB.Create(&giftExchanges)
	}

	return giftExchanges, nil
}

func remove(slice []schema.FamilyMember, element schema.FamilyMember) []schema.FamilyMember {
	i := 0
	for i < len(slice) {
		if slice[i] == element {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
		i++
	}
	return slice
}
