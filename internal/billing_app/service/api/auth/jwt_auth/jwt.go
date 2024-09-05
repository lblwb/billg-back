package jwt_auth

import (
	modelUser "backend/internal/billing_app/models/user"
	"backend/internal/database"
	"crypto/rand"
	"encoding/base64"
	"errors"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"time"
)

//var (
//	SecretToken string
//)

type userCredential struct {
	Username []byte `json:"username"`
	Password []byte `json:"password"`
	jwt.RegisteredClaims
}

type JwtAuths struct {
	SecretToken string
	usersEntity *modelUser.UsersEntity
	db          *database.StorageDb
}

var JwtPairKey = ""

func NewJwtAuths(db *database.StorageDb) *JwtAuths {

	jwtAuths := &JwtAuths{
		db:          db,
		usersEntity: modelUser.NewUsersEntity(db),
	}

	jwtAuths.SecretToken = JwtPairKey

	//jwtAuths.generateNewPairs()
	return jwtAuths
}

//const SECRET_TOKEN = "kdjfdjklejrkrej"

func GenerateNewPairs() string {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("GenerateKey Pair: %v", err)
	}
	return string(key)
}

func (ja JwtAuths) GenerateJwtTokenWithClaims(username string, password []byte) (string, jwt.Claims, error) {
	// Create the Claims
	//claims := jwt.MapClaims{
	//	"name":  username,
	//	"admin": true,
	//	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	//}

	expireAt := time.Now().Add(36 * time.Hour)

	claims := &userCredential{
		Username: []byte(username),
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	}

	var secretToken = ja.SecretToken

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secretToken))
	return t, claims.RegisteredClaims, err
}

func (ja JwtAuths) VerifyJWTToken(tokenString string) (jwt.MapClaims, error) {
	var secretToken = ja.SecretToken
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretToken), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// MiddlewareJwtAuthRoute - Middleware авторизации JWT
func MiddlewareJwtAuthRoute(app *fiber.App, router fiber.Router, db *database.StorageDb) *fiber.App {
	jwa := NewJwtAuths(db)

	// JWT Middleware
	router.Use(jwtware.New(
		jwtware.Config{
			//Обработка успешной авторизации
			//SuccessHandler: func(c *fiber.Ctx) error {
			//	//fmt.Println("Auth->Success->Handle-->", c)
			//	//log.Println("user", "claims:", c)
			//	//return c
			//	return c.Next()
			//},
			//Обработка ошибок
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": "Not Auth!",
				})
			},
			SigningKey: jwtware.SigningKey{JWTAlg: jwtware.HS512, Key: []byte(jwa.SecretToken)},
			//TokenLookup: "header:ex",
			//TokenLookup: "header:Authorization",
		},
	))

	return app

}

//Tool

// GetClaimByUser - Получения сессионных полей пользователя
func (ja JwtAuths) GetClaimByUser(c *fiber.Ctx, field string) interface{} {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims[field]
}

func (ja JwtAuths) GetUserDataByClaim(c *fiber.Ctx) (modelUser.Users, error) {
	//var userData modelUser.Users

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	usernameBase64 := claims["username"].(string)
	usernameByte, err := base64.StdEncoding.DecodeString(usernameBase64)
	if err != nil {
		log.Fatalln("username usernameByte", usernameByte)
		return modelUser.Users{}, err
	}
	username := string(usernameByte)
	//

	//fmt.Println("JWT-Claim:", username)

	//
	//db, err := ja.db.GetDB()
	//if err != nil {
	//	log.Fatalln("username dbconnect", db)
	//	return modelUser.Users{}, err
	//}

	if username != "" {
		userByUsername, err := ja.usersEntity.GetUserByUsername(username)
		if err != nil {
			return modelUser.Users{}, err
		}

		return userByUsername, nil

		//db := db.First(&userData, modelUser.Users{Username: username})
		//if db.Error != nil {
		//	log.Fatalln(db)
		//	return modelUser.Users{}, db.Error
		//}
		//return userData, nil
	} else {
		//log.Fatalln("username check", db)
		return modelUser.Users{}, err
	}

	//userData, err := ja.usersEntity.GetUserByUsername(string(username))
	//if err != nil {
	//	return modelUser.Users{}, err
	//} else {
	//	return userData, nil
	//}
}
