package service

import (
	"chat/internal/model"
	"chat/internal/repository/redis"
	"chat/internal/repository/request"
	"strings"
)

type AstroDictService struct {
	db *redis.AstroDict
}

func NewAstroDictService() *AstroDictService {
	return &AstroDictService{redis.NewAstroDict()}
}

func (ad *AstroDictService) Fetch(name string) (res *model.AstroDict, err error) {
	astro, err := ad.db.Fetch()
	if err != nil {
		return
	}
	if astro == nil || len(astro.AstroDict) == 0 {
		astro, err = request.NewAD("ce").GetAstroDict()
		if err != nil {
			return
		}
		err = ad.db.Save(*astro)
		if err != nil {
			return
		}
	}

	name = strings.ToLower(name)
	res = &model.AstroDict{}
	for _, v := range astro.AstroDict {
		c := strings.ToLower(v.C)
		e := strings.ToLower(v.E)
		if strings.Contains(c, name) || strings.Contains(e, name) {
			res.AstroDict = append(res.AstroDict, v)
		}
	}
	return
}
