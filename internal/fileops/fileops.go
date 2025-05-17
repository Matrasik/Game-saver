package fileops

import (
	"fmt"
	"log"
	"os"
	"path"
)

type FileOperator interface {
	Copy(src, dst string) error
	ReadAll(path string) (map[string][]byte, error)
	Exists(path string) (bool, error)
}

type OSFileOperator struct {
	//Buffer map[string][]byte
}

//func NewOSFileOperator() *OSFileOperator{
//	return &OSFileOperator{
//		Buffer: make(map[string][]byte)}
//}
//
//type DirContent struct {
//	files map[string][]byte
//}

func (o *OSFileOperator) Copy(src, dst string) error {
	m, err := o.ReadAll(src)
	if err != nil {
		log.Printf("Ошибка чтения папки: %s", err)
	}
	//Цикл по мапе
	for elem := range m {
		folderPath := path.Join(dst, elem)
		err = os.WriteFile(folderPath, m[elem], 0644)
		if err != nil {
			log.Printf("Ошибка создания файла на основе мапы: %s", err)
		}
	}
	return nil
}

func (o *OSFileOperator) ReadAll(path string) (map[string][]byte, error) {
	m := make(map[string][]byte)
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения папки %#v", err)
	}
	for _, elem := range dir {
		if !elem.IsDir() {
			file, err := os.ReadFile(path + string(os.PathSeparator) + elem.Name())
			if err != nil {
				return nil, fmt.Errorf("ошибка чтения файла %s, : %#v", elem.Name(), err)
			}
			m[elem.Name()] = file
		}
	}
	return m, nil
	//TODO Сделать еще и углубление в уровне папки. Сохранять в ту же мапу.
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
