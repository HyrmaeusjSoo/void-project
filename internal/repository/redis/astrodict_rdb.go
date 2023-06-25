package redis

import (
	"context"
	"encoding/json"
	"time"
	"void-project/global"
	"void-project/internal/model"
	"void-project/internal/repository/driver"

	"github.com/redis/go-redis/v9"
)

type AstroDict struct {
	db  *redis.Client
	ctx context.Context
}

func NewAstroDict() *AstroDict {
	return &AstroDict{
		db:  driver.Redis,
		ctx: context.Background(),
	}
}

func (ad *AstroDict) Save(astrodict model.AstroDict) error {
	val, err := json.Marshal(astrodict)
	if err != nil {
		return err
	}
	err = ad.db.Set(ad.ctx, "astrodict_ce", string(val), global.Config.System.AstroDictCacheExpire*time.Hour).Err()
	return err
}

func (ad *AstroDict) Fetch() (astro *model.AstroDict, err error) {
	val, err1 := ad.db.Get(ad.ctx, "astrodict_ce").Result()
	if err1 != nil && err1 != redis.Nil {
		return nil, err1
	}

	if val != "" {
		astro = &model.AstroDict{}
		err = json.Unmarshal([]byte(val), astro)
	}
	return
}
