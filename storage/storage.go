package storage

import (
	"encoding/json"
	"go_pass/file"

	"github.com/fatih/color"
)

func ByteWriter(data *[]byte) {
	file.ByteWriter(*data, "storage.json")
}

func ReadFromJson(path string) ([]byte, error) {
	data, err := file.ReadFile(path)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	var binData []byte
	err = json.Unmarshal(data, &binData)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	return binData, nil
}
