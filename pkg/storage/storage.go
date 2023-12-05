package storage

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Storage struct {
	Config *Config
}

// 配置
type Config struct {
	FileName    string              // 文件名称
	FileSize    int64               // 文件大小
	FileType    string              // 文件类型
	FileHeader  map[string][]string // 文件头
	FileContent []byte              // 文件内容
	LimitSize   int64               // 限制文件大小
	LimitType   []string            // 限制文件类型
	SavePath    string              // 保存路径
	Driver      string              // 储存驱动
}

// 存储驱动
var (
	LocalDriver = "local"
	OSSDriver   = "oss"
)

// 初始化默认配置
func Default() *Storage {
	return &Storage{
		Config: &Config{
			FileName: fmt.Sprint(md5.Sum([]byte(time.Now().String()))),
			SavePath: "./web/storage/",
			Driver:   LocalDriver,
		},
	}
}

// 初始化
func New(config *Config) *Storage {
	return &Storage{
		Config: config,
	}
}

// 设置文件名
func (t *Storage) SetFileName(fileName string) *Storage {

	t.Config.FileName = fileName

	return t
}

// 设置文件大小
func (t *Storage) SetFileSize(fileSize int64) *Storage {

	t.Config.FileSize = fileSize

	return t
}

// 设置文件类型
func (t *Storage) SetFileType(fileType string) *Storage {

	t.Config.FileType = fileType

	return t
}

// 设置文件头部
func (t *Storage) SetFileHeader(fileHeader map[string][]string) *Storage {

	t.Config.FileHeader = fileHeader

	return t
}

// 设置文件内容
func (t *Storage) SetFileContent(fileContent []byte) *Storage {

	t.Config.FileContent = fileContent

	return t
}

// 设置限制文件大小
func (t *Storage) SetLimitFileSize(limitSize int64) *Storage {

	t.Config.LimitSize = limitSize

	return t
}

// 设置限制文件类型
func (t *Storage) SetLimitFileType(limitType []string) *Storage {

	t.Config.LimitType = limitType

	return t
}

// 设置文件保存路径
func (t *Storage) SetSavePath(savePath string) *Storage {

	t.Config.SavePath = savePath

	return t
}

// 设置储存驱动
func (t *Storage) SetDriver(driver string) *Storage {

	t.Config.Driver = driver

	return t
}

// 保存文件
func (t *Storage) Save() (string, error) {

	var (
		url string
		err error
	)

	// 检查文件限制
	if err := t.checkFileLimit(); err != nil {
		return url, err
	}

	if t.Config.FileName == "" {
		return url, errors.New("缺少文件名")
	}

	if len(t.Config.FileContent) <= 0 {
		return url, errors.New("缺少文件内容")
	}

	// 选择储存驱动
	switch t.Config.Driver {
	case LocalDriver:
		// 保存到本地
		err = t.saveToLocal()
		// 设置文件地址
		url = t.Config.SavePath + t.Config.FileName
	case OSSDriver:
		// 保存到OSS
		err = t.saveToOSS()
	default:
		err = t.saveToLocal()
		// 设置文件地址
		url = t.Config.SavePath + t.Config.FileName
	}

	return url, err
}

// 检查文件限制
func (t *Storage) checkFileLimit() error {

	// 检查文件类型
	if err := t.checkLimitType(); err != nil {
		return err
	}

	// 检查文件大小
	if err := t.checkLimitSize(); err != nil {
		return err
	}

	return nil
}

// 检查文件类型
func (t *Storage) checkLimitType() error {

	if len(t.Config.LimitType) <= 0 || t.Config.FileType == "" {
		return nil
	}

	for _, limitType := range t.Config.LimitType {
		// 类型匹配合法
		if limitType == t.Config.FileType {
			return nil
		}
	}

	// 类型匹配不合法
	return errors.New("文件格式不合法")
}

// 检查文件大小
func (t *Storage) checkLimitSize() error {

	if t.Config.LimitSize > 0 && t.Config.FileSize > 0 && t.Config.LimitSize < t.Config.FileSize {
		return errors.New("文件大小超出限制")
	}

	return nil
}

// 保存至本地
func (t *Storage) saveToLocal() error {

	// 检查文件保存路径
	if _, err := os.Stat(t.Config.SavePath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(t.Config.SavePath, 0666)
		}
	}

	return ioutil.WriteFile(t.Config.SavePath+t.Config.FileName, t.Config.FileContent, 0666)
}

// 保存至OSS
func (t *Storage) saveToOSS() error {

	// TODO

	return nil
}
