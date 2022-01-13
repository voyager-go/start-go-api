package entities

import "github.com/voyager-go/start-go-api/entities/internal"

type CasbinRule internal.CasbinRule

type CasbinRuleServiceCreateReq struct {
	RoleID string `json:"role_id"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
