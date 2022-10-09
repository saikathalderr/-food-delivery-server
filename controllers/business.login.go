package controllers

import (
	"foodDeliveryAppServer/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JwtKey = []byte("${jwt_secret_key}")

type LoginBody struct {
	Email    string `form:"email" json:"email" binding:"required" valid:"email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Claims struct {
	jwt.StandardClaims
	ID               uint
	Email            string
	Address          string
	PhoneNumber      int
	OrganizationName string
}

func LoginHandler(ctx *gin.Context) {
	// get db instance
	dbCtx, _ := ctx.Get("db")
	db, ok := dbCtx.(*gorm.DB)
	// DB error if any
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "DB instance not found!"})
		return
	}

	// Login request body
	var reqBody LoginBody
	// Validation if any field is missing
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if email not exists in DB
	var LocalBusiness models.LocalBusiness
	hasRegisteredResult := db.Where("email = ?", reqBody.Email).First(&LocalBusiness)

	if hasRegisteredResult.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	// Check if password matches or not
	err := bcrypt.CompareHashAndPassword([]byte(LocalBusiness.Password), []byte(reqBody.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	// here, I have kept it as 10 minutes
	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		ID:               LocalBusiness.ID,
		Email:            LocalBusiness.Email,
		Address:          LocalBusiness.Address,
		PhoneNumber:      LocalBusiness.PhoneNumber,
		OrganizationName: LocalBusiness.OrganizationName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Success message
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
