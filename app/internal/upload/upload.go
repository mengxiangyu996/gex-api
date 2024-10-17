package upload

import (
	"errors"
	"gex-api/config"
	"math/rand"
	"net/textproto"
	"os"
	"strings"
	"time"
)

// 上传文件
type Upload struct {
	Config *Config
	File   *File
}

var (
	UploadLocalDriver = "local"
	UploadOssDriver   = "oss"
)

type UploadOption func(*Config)

// 上传配置
type Config struct {
	Driver    string   // 上传驱动
	SavePath  string   // 保存路径
	UrlPath   string   // 访问地址路径
	LimitSize int      // 限制文件大小
	LimitType []string // 限制文件类型
}

// 文件信息
type File struct {
	FileName    string               // 文件名
	FileSize    int                  // 文件大小
	FileType    string               // 文件类型
	FileHeader  textproto.MIMEHeader // 文件头
	FileContent []byte               // 文件内容
}

// 返回结果
type Result struct {
	FileName   string `json:"fileName"`
	RandomName string `json:"randomName"`
	FileSize   int    `json:"fileSize"`
	FileType   string `json:"fileType"`
	SavePath   string `json:"savePath"`
	UrlPath    string `json:"urlPath"`
	Url        string `json:"url"`
}

// 初始化上传对象
func Init(options ...UploadOption) *Upload {

	todayPath := time.Now().Format("20060102") + "/"

	// 配置默认驱动
	config := &Config{
		Driver:    UploadLocalDriver,
		UrlPath:   "storage/" + todayPath,
		SavePath:  "web/storage/" + todayPath,
		LimitSize: 1024 * 1024 * 10,
		LimitType: []string{
			"image/jpeg",
			"image/png",
			"image/svg+xml",
			"audio/mpeg",
			"audio/x-m4a",
			"video/mp4",
			"video/x-flv",
			"video/x-m4v",
			"application/msword",
			"application/vnd.ms-excel",
			"application/vnd.ms-powerpoint",
			"application/pdf",
		},
	}

	for _, option := range options {
		option(config)
	}

	return &Upload{
		Config: config,
	}
}

// 设置上传驱动
func SetDriver(driver string) UploadOption {
	return func(config *Config) {
		config.Driver = driver
	}
}

// 设置保存路径
func SetSavePath(savePath string) UploadOption {
	return func(config *Config) {
		config.SavePath = savePath
	}
}

// 设置访问地址路径
func SetUrlPath(urlPath string) UploadOption {
	return func(config *Config) {
		config.UrlPath = urlPath
	}
}

// 设置限制文件大小
func SetLimitSize(limitSize int) UploadOption {
	return func(config *Config) {
		config.LimitSize = limitSize
	}
}

// 设置限制文件类型
func SetLimitType(limitType []string) UploadOption {
	return func(config *Config) {
		config.LimitType = limitType
	}
}

// 设置上传文件
func (t *Upload) SetFile(file *File) *Upload {

	t.File = file

	return t
}

// 保存文件
func (t *Upload) Save() (*Result, error) {

	var err error
	var domain string

	if config.App.Domain.Name == "" {
		return nil, errors.New("未找到域名，无法生成访问地址")
	}

	if config.App.Domain.SSL {
		domain = "https://" + config.App.Domain.Name
	} else {
		domain = "http://" + config.App.Domain.Name
	}

	if t.File == nil || len(t.File.FileContent) <= 0 {
		return nil, errors.New("上传文件数据不全，无法保存")
	}

	// 获取文件后缀并且生成hash文件名
	fileName := strings.Split(t.File.FileName, ".")
	if len(fileName) != 2 {
		return nil, errors.New("文件缺少后缀")
	}

	// 拼接随机文件名
	randomName := t.generateRandomName() + "." + fileName[1]

	if err = t.checkLimitSize(); err != nil {
		return nil, err
	}

	if err = t.checkLimitType(); err != nil {
		return nil, err
	}

	switch t.Config.Driver {
	case UploadLocalDriver:
		err = t.saveToLocal(randomName)
	case UploadOssDriver:
		err = t.saveToOss()
	default:
		err = t.saveToLocal(randomName)
	}

	if err != nil {
		return nil, err
	}

	return &Result{
		FileName:   t.File.FileName,
		RandomName: randomName,
		FileSize:   t.File.FileSize,
		FileType:   t.File.FileType,
		SavePath:   t.Config.SavePath,
		UrlPath:    t.Config.UrlPath,
		Url:        domain + "/" + t.Config.UrlPath + randomName,
	}, err
}

// 检查文件大小
func (t *Upload) checkLimitSize() error {

	if len(t.Config.LimitType) <= 0 || t.File.FileType == "" {
		return nil
	}

	for _, limitType := range t.Config.LimitType {
		if limitType == t.File.FileType {
			return nil
		}
	}

	return errors.New("文件格式不合法")
}

// 检查文件类型
func (t *Upload) checkLimitType() error {

	if t.Config.LimitSize > 0 && t.File.FileSize > 0 && t.Config.LimitSize < t.File.FileSize {
		return errors.New("文件大小超出限制")
	}

	return nil
}

// 生成随机字符串
func (t *Upload) generateRandomName() string {

	// 创建一个新的随机数生成器实例
	r := rand.New(rand.NewSource(int64(len(t.File.FileName))))

	// 定义可能的字符集，包括字母和数字
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 生成随机字符串
	var randomName string
	for i := 0; i < 64; i++ {
		// 从字符集中随机选择一个字符
		randomChar := chars[r.Intn(len(chars))]
		randomName = randomName + string(randomChar)
	}

	return randomName
}

// 保存到本地
func (t *Upload) saveToLocal(randomName string) error {

	if _, err := os.Stat(t.Config.SavePath); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(t.Config.SavePath, 0644); err != nil {
				return err
			}
		}
	}

	return os.WriteFile(t.Config.SavePath+randomName, t.File.FileContent, 0644)
}

// 保存到Oss
func (t *Upload) saveToOss() error {

	// TODO

	return nil
}
