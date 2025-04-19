package types

import "strconv"

type ID string

func (id ID) Uint64() uint64 {
	i, _ := strconv.ParseUint(string(id), 10, 64)
	return i
}

func Uint64ID(u uint64) ID {
	return ID(strconv.FormatUint(u, 10))
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"` // Account name
	Password string `json:"password" validate:"required"` // Password
}

type UserLoginResponse struct {
	ID    ID     `json:"id"`    // Account ID
	Token string `json:"token"` // JWT token
}
