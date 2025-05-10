package fileops

type FileOperator interface {
	Copy(src, dst string) error
	ReadAll(path string) ([]byte, error)
	Exists(path string) bool
}

type OSFileOperator struct {
}

func (o *OSFileOperator) Copy(src, dst string) error {
	return nil
}

func (o *OSFileOperator) ReadAll(path string) ([]byte, error) {
	return nil, nil
}

func (o *OSFileOperator) Exists(path string) bool {
	return false
}
