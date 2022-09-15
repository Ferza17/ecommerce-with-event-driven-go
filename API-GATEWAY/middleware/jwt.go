package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/RoseRocket/xerrs"
	"github.com/dgrijalva/jwt-go"

	"github.com/Ferza17/event-driven-api-gateway/helper/response"
	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type TokenIdentity struct {
	UserId  string
	ExpTime int
}

// Expiration Time in Minute
const ExpTimeDay = 60
const signature = "SHOULD BE REPLACE"
const secret = "SHOULD BE REPLACE"

func CreateToken(userId string) (string, error) {
	expTime := time.Now().AddDate(0, 0, ExpTimeDay)

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = expTime.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	resultToken, err := at.SignedString([]byte(signature))
	if err != nil {
		err = xerrs.Mask(err, utils.ErrCreateToken)
		return "", err
	}

	return resultToken, nil
}

func JwtRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

		if tokenString == "" {
			response.Nay(w, r, http.StatusUnauthorized, utils.ErrJwtRequired)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, utils.ErrInvalidToken
			}
			return []byte(secret), nil
		})
		if err != nil {
			response.Nay(w, r, http.StatusUnauthorized, err)
			return
		}

		// Check Validation Token
		if err != nil {
			err = utils.ErrTokenInvalid
			response.Nay(w, r, http.StatusUnauthorized, err)
			return
		}

		ctx := context.WithValue(r.Context(), utils.TokenIdentityContextKey, SetTokenIdentity(token))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetTokenIdentity(token *jwt.Token) TokenIdentity {
	claims := token.Claims.(jwt.MapClaims)
	identity := TokenIdentity{}
	identity.UserId = claims["user_id"].(string)
	identity.ExpTime = int(claims["exp"].(float64))
	return identity
}

func GetTokenIdentityFromContext(ctx context.Context) TokenIdentity {
	return ctx.Value(utils.TokenIdentityContextKey).(TokenIdentity)
}
