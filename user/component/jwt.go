package component

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/types"
)

type jwtComponentImpl struct {
	SigningKey []byte
	ValidTime  time.Duration
	us         database.UserStore
}

type JwtComponent interface {
	// GenerateToken generate a jwt token, and return the token and signed string
	GenerateToken(ctx context.Context, req types.CreateJWTReq) (claims *types.JWTClaims, signed string, err error)
	ParseToken(ctx context.Context, token string) (user *types.User, err error)
}

func NewJwtComponent(signKey string, validHour int) JwtComponent {
	return &jwtComponentImpl{
		SigningKey: []byte(signKey),
		ValidTime:  time.Duration(validHour) * time.Hour,
		us:         database.NewUserStore(),
	}
}

// GenerateToken generate a jwt token, and return the token and signed string
func (c *jwtComponentImpl) GenerateToken(ctx context.Context, req types.CreateJWTReq) (claims *types.JWTClaims, signed string, err error) {
	u, err := c.us.FindByUUID(ctx, req.UUID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to find user by uuid '%s',error: %w", req.UUID, err)
	}
	expireAt := jwt.NewNumericDate(time.Now().Add(c.ValidTime))
	claims = &types.JWTClaims{
		UUID:        u.UUID,
		CurrentUser: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expireAt,
			Issuer:    "OpenCSG",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err = token.SignedString(c.SigningKey)
	if err != nil {
		return nil, "", fmt.Errorf("generate jwt token failed: %w", err)
	}

	return claims, signed, nil
}

func (c *jwtComponentImpl) ParseToken(ctx context.Context, token string) (user *types.User, err error) {
	claims := &types.JWTClaims{}
	_, err = jwt.ParseWithClaims(token, claims,
		func(token *jwt.Token) (interface{}, error) {
			return c.SigningKey, nil
		},
		jwt.WithIssuer("OpenCSG"),
	)

	if err != nil {
		return nil, fmt.Errorf("parse jwt token failed: %w", err)
	}

	dbu, err := c.us.FindByUsername(ctx, claims.CurrentUser)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by name '%s', %w", claims.CurrentUser, err)
	}

	// create new user object
	u := &types.User{
		UUID:              dbu.UUID,
		Username:          dbu.Username,
		Email:             dbu.Email,
		Roles:             dbu.Roles(),
		CanChangeUserName: dbu.CanChangeUserName,
	}
	return u, nil
}
