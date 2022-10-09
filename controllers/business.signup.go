package controllers

import (
	"foodDeliveryAppServer/models"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignUpBody struct {
	OrganizationName string `form:"organizationName" json:"organizationName" binding:"required"`
	Email            string `form:"email" json:"email" binding:"required" valid:"email"`
	Password         string `form:"password" json:"password" binding:"required"`
	Address          string `form:"address" json:"address" binding:"required"`
	PhoneNumber      int    `form:"phoneNumber" json:"phoneNumber" binding:"required"`
}

func SignUpHandler(ctx *gin.Context) {
	// get db instance
	dbCtx, _ := ctx.Get("db")
	db, ok := dbCtx.(*gorm.DB)
	// DB error if any
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "DB instance not found!"})
		return
	}

	// Sign-up request body
	var reqBody SignUpBody
	// Validation if any field is missing
	if err := ctx.BindJSON(&reqBody); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email is already exists in DB
	var LocalBusinesses []models.LocalBusiness
	hasRegisteredResult := db.Where("email = ?", reqBody.Email).Find(&LocalBusinesses)

	if hasRegisteredResult.RowsAffected > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is already registered, please login instead"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	// Handle has error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong hashing password!"})
		return
	}

	// Map body with model
	business := models.LocalBusiness{
		OrganizationName: reqBody.OrganizationName,
		Email:            reqBody.Email,
		Password:         string(hashedPassword),
		Address:          reqBody.Address,
		PhoneNumber:      reqBody.PhoneNumber,
	}
	// Create the new account
	result := db.Create(&business)
	// Handle error if any from DB
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong creating business account"})
		return
	}

	// Success message
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Business registered successfully!",
	})
}
