package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestHttpEpt(t *testing.T) {
	//err, res := done()
	//if err != nil {
	//	t.Log(err)
	//}
	//t.Log(string(res))

	g := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		g.Add(1)
		go func() {
			defer g.Done()
			err, res := done()
			if err != nil {
				t.Log(err)
			}
			t.Log(string(res))
		}()
	}
	g.Wait()
}

func done() (error, []byte) {
	cli := http.Client{}
	total := 268795
	s := time.Now().Format("20060102150405") + RandStringBytes(5)

	list := make([]map[string]interface{}, 0)

	for i := 0; i < total; i++ {
		list = append(list, map[string]interface{}{
			"number": "test", "name": "没问题", "other": "无形装逼", "age": 28,
		})
	}

	data := map[string]interface{}{
		"hash_mark": s,
		"total":     total,
		"header": []map[string]string{
			{"field": "number", "title": "编号"},
			{"field": "name", "title": "姓名"},
			{"field": "other", "title": "特长"},
			{"field": "age", "title": "年龄"},
		},
		"data": list,
	}

	jsonStr, _ := json.Marshal(data)

	resp, err := cli.Post("http://127.0.0.1:8081/ept", "application/json", bytes.NewBuffer(jsonStr))
	defer resp.Body.Close()
	if err != nil {
		return err, nil
	}

	res, _ := ioutil.ReadAll(resp.Body)

	return nil, res
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
