package lib

import (
	"context"
	"fmt"
	"time"

	"github.com/blog/src/global"
	"github.com/google/uuid"
)

const TokenExpiration time.Duration = 60 * 60

type Token string

func CreateToken() Token {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return Token(uuid.String())
}

func (t *Token) genTokenKey() string {
	return fmt.Sprintf("token:%s", *t)
}

func (t *Token) Find() (uint, error) {
	rdb := global.GetRedis()
	ctx := context.Background()
	value, err := rdb.Get(ctx, t.genTokenKey()).Int()
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

// TODO: 如果一个用户持续在线，考虑如何延长用户token的过期时间
func (t *Token) Set(id uint, ttl time.Duration) {
	rdb := global.GetRedis()
	ctx := context.Background()
	rdb.Set(ctx, t.genTokenKey(), id, ttl)
}

func (t *Token) Delete() {
	rdb := global.GetRedis()
	rdb.Del(context.Background(), t.genTokenKey()).Result()
}
