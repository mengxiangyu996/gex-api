package response

import "isme-go/types"

type Role struct {
	Id     int           `json:"id"`
	Code   string        `json:"code"`
	Name   string        `json:"name"`
	Enable types.Boolean `json:"enable"`
}
