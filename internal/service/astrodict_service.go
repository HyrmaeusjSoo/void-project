package service

import (
	"errors"
	"strings"
	"void-project/internal/model"
	"void-project/internal/repository/redis"
	"void-project/internal/repository/request"
	"void-project/internal/repository/sqlite"
)

type AstroDictService struct {
	rdb *redis.AstroDict
	db  *sqlite.AstrodictRepository
}

func NewAstroDictService() *AstroDictService {
	return &AstroDictService{redis.NewAstroDict(), sqlite.NewAstrodictRepository()}
}

// 从远程查询 - 短期内缓存到Redis
func (ad *AstroDictService) FetchRemote(name string) (res *model.AstroDictJson, err error) {
	astro, err := ad.rdb.Fetch()
	if err != nil {
		return
	}
	if astro == nil || len(astro.AstroDict) == 0 {
		astro, err = request.NewAD("ce").GetAstroDict()
		if err != nil {
			return
		}
		err = ad.rdb.Save(*astro)
		if err != nil {
			return
		}
	}

	name = strings.ToLower(name)
	res = &model.AstroDictJson{}
	for _, v := range astro.AstroDict {
		c, e := strings.ToLower(v.C), strings.ToLower(v.E)
		if strings.Contains(c, name) || strings.Contains(e, name) {
			res.AstroDict = append(res.AstroDict, v)
		}
	}
	return
}

// 查询
func (ad *AstroDictService) Fetch(name string) ([]model.Astrodict, error) {
	return ad.db.GetList(name)
}

// 同步到本地库
func (ad *AstroDictService) Sync(lang string) error {
	// 获取
	sl := make([]*model.Astrodict, 0, 100)
	astro, err := request.NewAD(lang).GetAstroDict()
	if err != nil {
		return errors.New("")
	}
	for _, v := range astro.AstroDict {
		sl = append(sl, &model.Astrodict{v.C, v.E})
	}

	return ad.db.Create(lang, sl)
}
