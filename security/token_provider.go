package security

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/segmentio/ksuid"
	"log"
	"mini-market-go/models"
	"strings"
	"time"
)

type User struct {
	UserId      int64               `json:"userId"`
	Authorities []*models.Authority `json:"authorities"`
	FacebookId  string              `json:"facebookId"`
}

type PrincipleClaims struct {
	User *User `json:"user,omitempty"`
}

type JwtCustomClaims struct {
	jwt.StandardClaims
	Principle PrincipleClaims `json:"principle,omitempty"`
}

func VerifyRequest(ctx *context.Context) error {
	//glog.Infof("URL: %s", ctx.Input.URL())
	isGet := ctx.Input.IsGet()
	if strings.HasPrefix(ctx.Input.URL(), "/core/v1/accounts/authenticate") {
		return nil
	}
	if strings.HasPrefix(ctx.Input.URL(), "/core/v1/images/upload") {
		return nil
	}
	if strings.HasPrefix(ctx.Input.URL(), "/core/v1/items/get-by-customer") {
		return nil
	}
	if strings.HasPrefix(ctx.Input.URL(), "/core/v1/items/get-by-customer") {
		return nil
	}
	if strings.HasPrefix(ctx.Input.URL(), "/core/v1/items") && isGet {
		return nil
	}
	return VerifyToken(ctx)
}

func VerifyToken(ctx *context.Context) error {
	userAgent := ctx.Input.UserAgent()
	if userAgent == "" {
		return errors.New("auth.error.missing.useragent")
	}
	strAuthorization := ctx.Request.Header.Get("Authorization")
	if strAuthorization == "" {
		return errors.New("auth.error.missing.token")
	}
	token, _ := GetTokenFromStrAuthorization(strAuthorization)
	if claims, err := ParseToken(token); err == nil {
		ctx.Input.SetData("user", claims)
		return nil
	} else {
		return err
	}
}

func ParseToken(tk string) (*JwtCustomClaims, error) {
	claims := JwtCustomClaims{}
	tokenStatus, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("JwtUserSecret")), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenStatus == nil || !tokenStatus.Valid {
		return nil, errors.New("auth.error.tokeninvalid")
	}
	return &claims, nil
}

func GetTokenFromStrAuthorization(token string) (string, error) {
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:], nil
	}
	return token, nil
}

func CreateUserToken(u *models.User, provider string) string {
	// Embed Custom information to `token`
	tokenId := ksuid.New().String()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        tokenId,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		Principle: PrincipleClaims{
			User: &User{
				UserId:      u.Id,
				Authorities: u.Authorities,
			},
		},
	})
	// token -> string. Only server knows this secret (foobar).
	strToken, err := token.SignedString([]byte(beego.AppConfig.String("JwtUserSecret")))
	if err != nil {
		log.Fatalln(err)
	}

	return strToken
}

func GetCurrentUser(c *beego.Controller) (*User, error) {
	u := c.Ctx.Input.GetData("user").(*JwtCustomClaims)
	if u == nil {
		return nil, errors.New("auth.error.missing.principal")
	}

	return u.Principle.User, nil
}

func GetCurrentUserId(c *beego.Controller) (int64, error) {
	u, err := GetCurrentUser(c)
	if u == nil || err != nil {
		return -1, errors.New("auth.error.missing.principal")
	}

	return u.UserId, nil
}
