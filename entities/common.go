package entities

import (
	"github.com/voyager-go/start-go-api/global"
	"github.com/voyager-go/start-go-api/pkg/util"
)

type PageResult struct {
	util.Pagination
	List interface{} `json:"list"`
}

type PageReq struct {
	util.Pagination
	Searches []global.Search `json:"searches"`
}
