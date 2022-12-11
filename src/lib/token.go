package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/blog/src/global"
	"github.com/google/uuid"
)

const TokenExpiration time.Duration = time.Hour

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

func (t *Token) Find() (*AuthorBaseInfo, error) {
	rdb := global.GetRedis()
	ctx := context.Background()
	value, err := rdb.Get(ctx, t.genTokenKey()).Result()
	if err != nil {
		return nil, err
	}

	var authorInfo AuthorBaseInfo
	err = json.Unmarshal([]byte(value), &authorInfo)
	if err != nil {
		panic(err)
	}
	return &authorInfo, nil
}

// TODO: 如果一个用户持续在线，考虑如何延长用户token的过期时间
func (t *Token) Set(authorInfo *AuthorBaseInfo, ttl time.Duration) {
	rdb := global.GetRedis()
	ctx := context.Background()

	data, err := json.Marshal(authorInfo)
	if err != nil {
		panic(err)
	}
	key := t.genTokenKey()
	rdb.Set(ctx, key, string(data), ttl)
}

func (t *Token) Delete() {
	rdb := global.GetRedis()
	rdb.Del(context.Background(), t.genTokenKey()).Result()
}
