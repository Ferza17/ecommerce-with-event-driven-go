package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/RoseRocket/xerrs"
	"github.com/dgrijalva/jwt-go"
	"github.com/vektah/gqlparser/v2/gqlerror"

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

func RegisterTokenHTTPContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), utils.TokenContextKey, tokenString)))
	})
}

func DirectiveJwtRequired(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	token := GetTokenFromContext(ctx)
	if token == "" {
		return nil, &gqlerror.Error{
			Message: utils.ErrInvalidToken.Error(),
		}
	}
	rawToken, err := validateToken(token)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
		}
	}
	ctx = context.WithValue(ctx, utils.TokenIdentityContextKey, SetTokenIdentity(rawToken))
	return next(ctx)
}

func validateToken(token string) (identity *jwt.Token, err error) {
	if token == "" {
		err = xerrs.Mask(err, utils.ErrJwtRequired)
		return
	}
	identity, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.ErrInvalidToken
		}
		return []byte(secret), nil
	})
	if err != nil {
		err = xerrs.Mask(err, utils.ErrForbidden)
		return
	}
	return
}

func GetTokenFromContext(ctx context.Context) string {
	return ctx.Value(utils.TokenContextKey).(string)
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
