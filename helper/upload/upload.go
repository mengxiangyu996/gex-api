package upload

import (
	"breeze-api/pkg/storage"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"time"
)

// 上传
type Upload struct{}

// 上传文件
func File(fileHeader *multipart.FileHeader) (string, error) {

	file, _ := fileHeader.Open()
	fileByte, _ := ioutil.ReadAll(file)

	// 保存文件
	url, err := storage.New(&storage.Config{
		SavePath: "./web/storage/upload/" + time.Now().Format("20060102") + "/",
	}).SetFile(&storage.File{
		FileName:    fileHeader.Filename,
		FileContent: fileByte,
	}).Save()

	return url, err
}

// 上传base64文件
func Base64(file, name string) (string, error) {

	fileByte, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return "", errors.New("base64文件转换失败")
	}

	// 保存文件
	url, err := storage.New(&storage.Config{
		SavePath: "./web/storage/upload/" + time.Now().Format("20060102") + "/",
	}).SetFile(&storage.File{
		FileName:    name,
		FileContent: fileByte,
	}).Save()

	return url, err
}
