package fileops

import (
	"os"
)

type FileOperator interface {
	Copy(src, dst string) error
	ReadAll(path string) (map[string][]byte, error)
	Exists(path string) (bool, error)
}

type OSFileOperator struct {
}

func (o *OSFileOperator) Copy(src, dst string) error {

	return nil
}

func (o *OSFileOperator) ReadAll(path string) (map[string][]byte, error) {
	return nil, nil
}

func (o *OSFileOperator) Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

//TODO Проверять на наличие файлов, даже если папка есть
