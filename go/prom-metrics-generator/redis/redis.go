package redis

import (
	"github.com/go-redis/redis/v8"
	"tmomon/logger"
	"tmomon/metrics"
	"context"
	"strconv"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func RedisConn(
	url string,
){
	rdb = redis.NewClient(&redis.Options{
		Addr: url,
	})
}

func RedisIfExists(
	stream string,
) int64 {
	// Check if key exists or not
	exists, err := rdb.Exists(ctx, stream).Result()
	if err != nil {
		clog.Error.Println(err)
	}
	return exists
}

func RedisStreamLength(
	url string,
	stream string,
){
	COMMON_FIELD := "{ redis_uri=\""+url+"\", stream=\""+stream+"\"}"
	
	// Length of stream
	len, err := rdb.XLen(ctx, stream).Result()
	if err != nil {
		clog.Error.Println(err)
	}else{
		REDIS_STREAM_LENGTH := strconv.FormatInt(len, 10) // base 10 (decimal)
		// Write a metrics
		metrics.WriteMetrics("redis_stream_length", "Redis stream length", "gauge", COMMON_FIELD, COMMON_FIELD, REDIS_STREAM_LENGTH)
	}
}

func RedisStreamPending(
	url string,
	stream string,
	group string,
){
	COMMON_FIELD := "{ redis_uri=\""+url+"\", stream=\""+stream+"\", consumer_group=\""+group+"\"}"

	// Pending
	pen, err := rdb.XPending(ctx, stream, group).Result()
	if err != nil {
		clog.Error.Println(err)
		metrics.WriteMetrics("redis_consumer_group_stream_pending", "Redis pending stream at consumer_group level", "gauge", COMMON_FIELD, COMMON_FIELD, "-1")
	}else{
		REDIS_CONSUMER_GROUP_STREAM_PENDING := strconv.FormatInt(pen.Count, 10) // base 10 (decimal)
		// Write a metrics
		metrics.WriteMetrics("redis_consumer_group_stream_pending", "Redis pending stream at consumer_group level", "gauge", COMMON_FIELD, COMMON_FIELD, REDIS_CONSUMER_GROUP_STREAM_PENDING)
	}
}
