package storage

import (
	"encoding/json"
	"go_pass/file"

	"github.com/fatih/color"
)

func SaveToJson(data *[]byte) {
	file.WriteFile(*data, "storage.json")
}

func ReadFromJson() []byte {
	data, err := file.ReadFile("storage.json")
	if err != nil {
		color.Red(err.Error())
	}
	var binData []byte
	err = json.Unmarshal(data, &binData)
	if err != nil {
		color.Red(err.Error())
	}
	return binData
}
