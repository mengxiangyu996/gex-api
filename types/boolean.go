package types

import (
	"database/sql/driver"
	"errors"
)

type Boolean bool

// 用于将Boolean转换为数据库值
func (b Boolean) Value() (driver.Value, error) {
	if b {
		return 1, nil // 将true映射为1
	}
	return 0, nil // 将false映射为0
}

// 用于从数据库值转换Boolean
func (b *Boolean) Scan(value interface{}) error {
	// 确保传入的值是int或int64类型
	val, ok := value.(int)
	if !ok {
		valInt64, ok := value.(int64)
		if !ok {
			return errors.New("提供的值不是整数")
		}
		val = int(valInt64)
	}

	// 检查值是否为0或1
	if val == 0 || val == 1 {
		*b = Boolean(val == 1)
		return nil
	}

	return errors.New("提供的值不是有效的布尔值")
}
