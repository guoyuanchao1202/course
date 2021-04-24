package data

import (
	"course/config"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
)

var DataPath string

func InitDataDir(path config.PathConfig) error {
	DataPath = path.PathName
	// 如果指定路径是一个文件夹，直接返回
	if isDir(path.PathName) {
		return nil
	}
	return os.Mkdir(path.PathName, os.ModePerm)
}

// 判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func Save(fileHeader *multipart.FileHeader) (string, error) {
	fileName := fileHeader.Filename
	upLoadFile, err := fileHeader.Open()
	if err != nil {
		log.Println("open upLoadFile failed: ", err.Error())
		return "", err
	}
	defer upLoadFile.Close()
	localFile, err := os.Create(fmt.Sprintf("%s/%s", DataPath, fileName))
	if err != nil {
		log.Println("create local file failed: ", err.Error())
		return "", err
	}
	_, err = io.Copy(localFile, upLoadFile)
	return fmt.Sprintf("%s/%s", DataPath, fileName), err
}

func GetData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func Remove(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		log.Println(path, " is not exist")
		return nil
	}
	return os.Remove(path)
}
