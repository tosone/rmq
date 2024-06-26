package rmq

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var unusedContext = context.TODO()

type RedisWrapper struct {
	rawClient redis.UniversalClient
}

func (wrapper RedisWrapper) Set(key string, value string, expiration time.Duration) error {
	// NOTE: using Err() here because Result() string is always "OK"
	return wrapper.rawClient.Set(unusedContext, key, value, expiration).Err()
}

func (wrapper RedisWrapper) Del(key string) (affected int64, err error) {
	return wrapper.rawClient.Del(unusedContext, key).Result()
}

func (wrapper RedisWrapper) TTL(key string) (ttl time.Duration, err error) {
	return wrapper.rawClient.TTL(unusedContext, key).Result()
}

func (wrapper RedisWrapper) LPush(key string, value ...string) (total int64, err error) {
	return wrapper.rawClient.LPush(unusedContext, key, value).Result()
}

func (wrapper RedisWrapper) LLen(key string) (affected int64, err error) {
	return wrapper.rawClient.LLen(unusedContext, key).Result()
}

func (wrapper RedisWrapper) LRem(key string, count int64, value string) (affected int64, err error) {
	return wrapper.rawClient.LRem(unusedContext, key, int64(count), value).Result()
}

func (wrapper RedisWrapper) LTrim(key string, start, stop int64) error {
	// NOTE: using Err() here because Result() string is always "OK"
	return wrapper.rawClient.LTrim(unusedContext, key, int64(start), int64(stop)).Err()
}

func (wrapper RedisWrapper) RPopLPush(source, destination string) (value string, err error) {
	value, err = wrapper.rawClient.RPopLPush(unusedContext, source, destination).Result()
	// println("RPopLPush", source, destination, value, err)
	switch err {
	case nil:
		return value, nil
	case redis.Nil:
		return value, ErrorNotFound
	default:
		return value, err
	}
}

func (wrapper RedisWrapper) RPop(source string) (value string, err error) {
	value, err = wrapper.rawClient.RPop(unusedContext, source).Result()
	// println("RPop", source, value, err)
	switch err {
	case nil:
		return value, nil
	case redis.Nil:
		return value, ErrorNotFound
	default:
		return value, err
	}
}

func (wrapper RedisWrapper) SAdd(key, value string) (total int64, err error) {
	return wrapper.rawClient.SAdd(unusedContext, key, value).Result()
}

func (wrapper RedisWrapper) SMembers(key string) (members []string, err error) {
	return wrapper.rawClient.SMembers(unusedContext, key).Result()
}

func (wrapper RedisWrapper) SRem(key, value string) (affected int64, err error) {
	return wrapper.rawClient.SRem(unusedContext, key, value).Result()
}

func (wrapper RedisWrapper) FlushDb() error {
	// NOTE: using Err() here because Result() string is always "OK"
	return wrapper.rawClient.FlushDB(unusedContext).Err()
}

func (wrapper RedisWrapper) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return wrapper.rawClient.Eval(ctx, script, keys, args...)
}

func (wrapper RedisWrapper) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return wrapper.rawClient.EvalSha(ctx, sha1, keys, args...)
}

func (wrapper RedisWrapper) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return wrapper.rawClient.ScriptExists(ctx, hashes...)
}

func (wrapper RedisWrapper) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return wrapper.rawClient.ScriptLoad(ctx, script)
}

func (wrapper RedisWrapper) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return wrapper.rawClient.EvalRO(ctx, script, keys, args...)
}

func (wrapper RedisWrapper) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return wrapper.rawClient.EvalShaRO(ctx, sha1, keys, args...)
}
