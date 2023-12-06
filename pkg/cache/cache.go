package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

var path = "./web/storage/cache/"

// 存储数据
// expire 过期时间，为0时不会过期
func Set(key string, value interface{}, expire time.Duration) error {

	cache := new(&cache{
		key:   key,
		value: value,
	})

	if expire > 0 {
		cache.expire = time.Now().Add(expire).Unix()
	}

	return cache.writeToFile()
}

// 获取数据
func Get(key string) interface{} {

	cache := new(&cache{
		key: key,
	})

	cache = cache.readFormFile()
	if cache == nil {
		return nil
	}

	if cache.expire > 0 && time.Now().Unix() > cache.expire {
		cache.removeFile()
		return nil
	}

	return cache.value
}

// 删除数据
func Del(key string) error {

	cache := new(&cache{
		key: key,
	})

	return cache.removeFile()
}

type cache struct {
	key    string
	expire int64
	value  interface{}
}

func new(c *cache) *cache {
	return c
}

// 缓存文件中读取数据
func (t *cache) readFormFile() *cache {

	contentByte, err := ioutil.ReadFile(path + t.key)
	if err != nil {
		return nil
	}

	var content map[string]interface{}

	if err := json.Unmarshal(contentByte, &content); err != nil {
		return nil
	}

	return &cache{
		key:    t.key,
		expire: int64(content["expire"].(float64)),
		value:  content["value"],
	}
}

// 数据写入缓存文件
func (t *cache) writeToFile() error {

	content, _ := json.Marshal(map[string]interface{}{
		"expire": t.expire,
		"value":  t.value,
	})

	// 检查文件保存路径
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, 0644)
		}
	}

	return ioutil.WriteFile(path+t.key, content, 0644)
}

// 删除缓存文件
func (t *cache) removeFile() error {
	return os.Remove(path + t.key)
}
