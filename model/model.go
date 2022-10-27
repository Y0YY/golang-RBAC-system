package model

import (
	"encoding/json"
	"errors"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func ReadData(filename string) ([]byte, error) {
	path := "./data/" + filename + ".txt"
	if !Exists(path) {
		_, err := os.Create(path)
		if err != nil {
			return nil, errors.New("创建数据文件失败")
		}
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("读取数据文件失败")
	}
	return data, nil
}

func SaveData(m map[string]interface{}, filename string) error {
	data, _ := json.Marshal(m)
	path := "./data/" + filename + ".txt"
	err := os.WriteFile(path, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
