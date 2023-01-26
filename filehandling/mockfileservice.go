package filehandling

import (
	"fmt"
	"os"
)

type MockFileService struct{}

func NewMockFileService() *MockFileService {
	return &MockFileService{}
}

func (mfs *MockFileService) ReadFile(name string) ([]byte, error) {
	//TODO implement me
	return []byte(fmt.Sprintf("{%q:[%q,%q,%q]}", "Cards", "a", "b", "c")), nil
}

func (mfs *MockFileService) WriteFile(name string, data []byte, perm os.FileMode) error {
	//TODO implement me
	return nil
}
