/**
 *@Author:sario
 *Date:2022/7/16
 *@Desc:
 */
package dao

import (
	"github.com/go-redis/redis"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

//测试前需打开redis
func Test_initRedis(t *testing.T) {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err := Redisdb.Ping().Result()
	if err != nil {
		t.Errorf("link redis err %v", err)
	}
}

func TestRUNDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RUNDB()
			if Db == nil {
				t.Errorf("link mysql err")
			}
		})
	}
}
