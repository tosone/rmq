package rmq

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	// simple keys
	Set(key string, value string, expiration time.Duration) error
	Del(key string) (affected int64, err error)
	TTL(key string) (ttl time.Duration, err error)

	// lists
	LPush(key string, value ...string) (total int64, err error)
	LLen(key string) (affected int64, err error)
	LRem(key string, count int64, value string) (affected int64, err error)
	LTrim(key string, start, stop int64) error
	RPopLPush(source, destination string) (value string, err error)
	RPop(source string) (value string, err error)

	// sets
	SAdd(key, value string) (total int64, err error)
	SMembers(key string) (members []string, err error)
	SRem(key, value string) (affected int64, err error)

	// special
	FlushDb() error

	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
	EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
	EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd
	ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd
	ScriptLoad(ctx context.Context, script string) *redis.StringCmd
}
