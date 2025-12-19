package typings

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	Token string `json:"token"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtSessionInfo struct {
	AccountId  []byte `json:"accountId"`
	EmployeeId []byte `json:"employeeId"`
	Username   string `json:"username"`
}

type JwtSessionParse struct {
	jwt.RegisteredClaims
	AccountId  []byte `json:"accountId"`
	EmployeeId []byte `json:"employeeId"`
	Username   string `json:"username"`
	Iat        int64  `json:"iat"`
}

type AuthenticateUpdate struct {
	AccountId    []byte `json:"accountId"`
	LoginCount   int    `json:"loginCount"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	UpdatedAt    int64  `json:"updatedAt"`
}
