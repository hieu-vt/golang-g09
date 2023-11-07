package jwt

import (
	"flag"
	"fmt"
	"g09-to-do-list/common"
	"g09-to-do-list/plugin/tokenprovider"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtProvider struct {
	name   string
	secret string
}

type myClaims struct {
	Payload common.TokenPayload `json:"payload"`
	jwt.MapClaims
}

type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func NewJwtProvider(name string) *jwtProvider {
	return &jwtProvider{name: name}
}

func (j *jwtProvider) GetPrefix() string {
	return j.name
}

func (j *jwtProvider) Get() interface{} {
	return j
}

func (j *jwtProvider) Name() string {
	return j.name
}

func (j *jwtProvider) InitFlags() {
	flag.StringVar(&j.secret, fmt.Sprintf("%s-secret", j.name), "hieu-dz", "Secret key of jwt provider")
}

func (j *jwtProvider) Configure() error {
	return nil
}

func (j *jwtProvider) Run() error {
	return nil
}

func (j *jwtProvider) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()
	return c
}

func (t *token) GetToken() string {
	return t.Token
}

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	// generate the JWT
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		common.TokenPayload{
			UId:   data.UserId(),
			URole: data.Role(),
		},
		jwt.MapClaims{
			"ExpiresAt": now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			"IssuedAt":  now.Local().Unix(),
			"Id":        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)

	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return claims.Payload, nil
}
