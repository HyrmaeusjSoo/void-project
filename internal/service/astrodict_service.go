package service

import (
	"chat/internal/model"
	"chat/internal/repository/redis"
	"chat/internal/repository/request"
	"strings"
)

type AstroDictService struct{}

var rdb = &redis.AstroDict{}

func (*AstroDictService) Fetch(name string) (res *model.AstroDict, err error) {
	ad, err := rdb.Fetch()
	if err != nil {
		return
	}
	if ad == nil || len(ad.AstroDict) == 0 {
		ad, err = request.New("ce").GetAstroDict()
		if err != nil {
			return
		}
		err = rdb.Save(*ad)
		if err != nil {
			return
		}
	}

	name = strings.ToLower(name)
	res = &model.AstroDict{}
	for _, v := range ad.AstroDict {
		c := strings.ToLower(v.C)
		e := strings.ToLower(v.E)
		if strings.Contains(c, name) || strings.Contains(e, name) {
			res.AstroDict = append(res.AstroDict, v)
		}
	}
	return
}
