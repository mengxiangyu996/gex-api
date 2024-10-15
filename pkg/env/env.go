package env

import (
	"encoding/json"
	"os"
	"strings"
)

// 读取配置文件
func Get(key string, defaultValue interface{}) interface{} {

	data, err := os.ReadFile("env.json")
	if err != nil {
		return defaultValue
	}

	var result map[string]interface{}
	if err = json.Unmarshal(data, &result); err != nil {
		return defaultValue
	}

	// 逐级访问下级字段
	for _, field := range strings.Split(key, ".") {
		value, ok := result[field]
		if !ok {
			return defaultValue
		}

		result, ok = value.(map[string]interface{})
		if !ok {
			return value
		}
	}

	return result
}