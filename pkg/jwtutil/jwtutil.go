package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

type Claims struct {
	UserID      string   `json:"user_id"`
	Username    string   `json:"username"`
	RoleIDs     []int    `json:"role_ids"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

// HasPermission reports whether the claims grant the given "resource:action" permission.
func (c Claims) HasPermission(resource, action string) bool {
	required := resource + ":" + action
	for _, p := range c.Permissions {
		if p == required {
			return true
		}
	}
	return false
}

func GenerateToken(secret string, expiresIn time.Duration, userID, username string, roleIDs []int, permissions []string) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:      userID,
		Username:    username,
		RoleIDs:     roleIDs,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiresIn)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
