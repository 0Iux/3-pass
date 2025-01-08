package file

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile(path string) ([]byte, error) {
	if !IsJson(path) {
		color.Red("Данный путь не ялвяется .json")
		return nil, errors.New("не ялвяется .json")
	}

	data, err := os.ReadFile(path)

	if err != nil {
		color.Red("Ошибка при чтении файла")
		return nil, err
	}
	return data, nil
}

func IsJson(path string) bool {
	return path[len(path)-5:] == ".json"
}

func ByteWriter(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
