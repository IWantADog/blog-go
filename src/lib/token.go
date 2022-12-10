package lib

import (
	"context"
	"fmt"
	"time"

	"github.com/blog/src/global"
)

const TokenExpiration time.Duration = 60 * 60

type token string

func CreateToken() token {

}

func (t *token) genTokenKey() string {
	return fmt.Sprintf("token:%s", *t)
}

func (t *token) Find() (uint, error) {
	rdb := global.GetRedis()
	ctx := context.Background()
	value, err := rdb.Get(ctx, t.genTokenKey()).Int()
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

// TODO: 如果一个用户持续在线，考虑如何延长用户token的过期时间
func (t *token) Set(id uint, ttl time.Duration) {
	rdb := global.GetRedis()
	ctx := context.Background()
	rdb.Set(ctx, t.genTokenKey(), id, ttl)
}
