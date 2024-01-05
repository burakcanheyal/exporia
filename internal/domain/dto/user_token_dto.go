package dto

type UserToken struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
