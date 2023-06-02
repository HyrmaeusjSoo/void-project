package redis

import (
	"chat/internal/model"
	"chat/internal/repository/driver"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type AstroDict struct{}

var ctx = context.Background()

func (*AstroDict) Save(astrodict model.AstroDict) error {
	val, err := json.Marshal(astrodict)
	if err != nil {
		return err
	}
	err = driver.Redis.Set(ctx, "astrodict_ce", string(val), 2*time.Hour).Err()
	return err
}

func (*AstroDict) Fetch() (ad *model.AstroDict, err error) {
	val, err1 := driver.Redis.Get(ctx, "astrodict_ce").Result()
	if err1 != nil && err1 != redis.Nil {
		return nil, err1
	}

	if val != "" {
		ad = &model.AstroDict{}
		err = json.Unmarshal([]byte(val), ad)
	}
	return
}
