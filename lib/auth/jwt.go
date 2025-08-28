package auth

import (
	"fmt"
	"news-portal/config"
	"news-portal/internal/core/domain/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt interface {
	GenerateToken(data *entity.JwtData) (string, int64, error)
	VerifyToken(token string) (*entity.JwtData, error)
}

type Options struct {
	SigningKey string
	issuer     string
}

func (o *Options) GenerateToken(data *entity.JwtData) (string, int64, error) {
	now := time.Now().Local()
	expiresAt := now.Add(time.Hour * 24)
	data.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiresAt)
	data.RegisteredClaims.Issuer = o.issuer
	data.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)
	acToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	accesToken, err := acToken.SignedString([]byte(o.SigningKey))
	if err != nil {
		return "", 0, err
	}
	return accesToken, expiresAt.Unix(), nil
}

func (o *Options) VerifyToken(token string) (*entity.JwtData, error) {
	parsedToken, err := jwt.Parse(token, func(t*jwt.Token)(interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(o.SigningKey), nil
	})

	if err!= nil {
		return nil, jwt.ErrECDSAVerification
	}

	if parsedToken.Valid {
		claim, ok := parsedToken.Claims.(jwt.MapClaims)
		if!ok || !parsedToken.Valid {
			return nil, err
		}

		jwtData := &entity.JwtData{
			UserId: claim["user_id"].(float64),
			RegisteredClaims: jwt.RegisteredClaims{},
		}

		return jwtData, nil
	}
	return nil, fmt.Errorf("Token is invalid")
}

func NewJwt(cfg *config.Config) Jwt {
	opt := new(Options)
	opt.SigningKey = cfg.JwtSecret
	opt.issuer = cfg.JwtIssuer

	return opt
}