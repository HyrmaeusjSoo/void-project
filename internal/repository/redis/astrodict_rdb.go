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

// 保存 - 按有效期
func (ad *AstroDict) Save(astrodict model.AstroDictJson) error {
	val, err := json.Marshal(astrodict)
	if err != nil {
		return err
	}
	err = ad.db.Set(ad.ctx, astrodict.Name(), string(val), time.Duration(global.Config.System.AstroDictCacheExpire)*time.Hour).Err()
	return err
}

// 查询
func (ad *AstroDict) Fetch() (astro *model.AstroDictJson, err error) {
	val, redisErr := ad.db.Get(ad.ctx, astro.Name()).Result()
	if redisErr != nil && redisErr != redis.Nil {
		return nil, redisErr
	}

	if val != "" {
		astro = &model.AstroDictJson{}
		err = json.Unmarshal([]byte(val), astro)
	}
	return
}
