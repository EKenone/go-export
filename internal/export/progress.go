package export

import (
	"context"
	"encoding/json"
	"go-export/pkg/redis"
	"log"
	"time"
)

type ProgressData struct {
	Total   int
	Current int
	Status  int
	Mark    string
	Url     string
}

const (
	RedisPrefix = "go-export:progress:"
	Expired     = 24 * time.Hour
)

const (
	StatusWait = iota
	StatusSuccess
)

func InitProgress(mark string, total int) {
	data := ProgressData{
		Total:  total,
		Mark:   mark,
		Status: StatusWait,
	}

	setProgressData(mark, data)
}

func Stepping(mark string, current int) {
	data := CurrentProgress(mark)

	data.Current = current

	setProgressData(mark, data)
}

func Finish(mark string, url string) {
	data := CurrentProgress(mark)

	data.Url = url
	data.Current = data.Total
	data.Status = StatusSuccess

	setProgressData(mark, data)
}

func CurrentProgress(mark string) ProgressData {
	var data ProgressData

	by, _ := redis.GetClient().Get(context.Background(), redisKey(mark)).Bytes()

	_ = json.Unmarshal(by, &data)

	return data
}

func setProgressData(mark string, data ProgressData) {
	by, _ := json.Marshal(data)

	err := redis.GetClient().Set(context.Background(), redisKey(mark), by, Expired)

	if err != nil {
		log.Println(err)
	}
}

func redisKey(key string) string {
	return RedisPrefix + key
}
