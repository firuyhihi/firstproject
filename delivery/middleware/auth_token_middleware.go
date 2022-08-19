package middleware

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"ticket.narindo.com/utils"

// 	"github.com/gin-gonic/gin"
// )

// type authHeader struct {
// 	AuthorizationHeader string `header:"Authorization"`
// }

// type AuthTokenMiddleware interface {
// 	RequireToken() gin.HandlerFunc
// }

// type authTokenMiddleware struct {
// 	acctToken utils.Token
// }

// // RequireToken implements AuthTokenMiddleware
// func (a *authTokenMiddleware) RequireToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		h := authHeader{}
// 		if err := c.ShouldBindHeader(&h); err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Unauthorized",
// 			})
// 			c.Abort()
// 		}

// 		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
// 		fmt.Println("tokenString: ", tokenString)
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		token, _ := a.acctToken.VerifyAccessToken(tokenString)
// 		userId, err := a.acctToken.FetchAccessToken(token)
// 		if userId == "" || err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		if token != nil { // token berisi, brarti dia sudah login
// 			c.Set("user-id", userId)
// 			c.Next()
// 		} else {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"message": "Unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		}
// 	}
// }

// func GetUserAccDet(email, userId string) (*model.AccessDetail, error) {
// 	var user model.UserTes
// 	cfg := config.NewConfig()
// 	infra := manager.NewInfra(&cfg)
// 	conn := infra.SqlDb()
// 	err := conn.Where("email = ? and user_id = ?", email, userId).Find(&user).Error
// 	if err != nil {
// 		fmt.Println("aaa")
// 	}
// 	tokenString := user.Token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("signing method invalid")
// 		} else if method != cfg.JwtSigningMethod {
// 			return nil, fmt.Errorf("signing method invalid")
// 		}

// 		return []byte(cfg.JwtSignatureKey), nil
// 	})

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
// 		log.Println("Token invalid...")
// 		return nil, errors.New("aaa")
// 	}
// 	accessUuid := claims["AccessUUID"].(string)
// 	userName := claims["Username"].(string)
// 	accDetail := &model.AccessDetail{
// 		AccessUuid: accessUuid,
// 		Username:   userName,
// 	}
// 	return accDetail, nil
// }

// func NewTokenValidator(acctToken utils.Token) AuthTokenMiddleware {
// 	return &authTokenMiddleware{acctToken: acctToken}
// }
