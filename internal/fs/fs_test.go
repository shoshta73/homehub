package fs

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	t.Run("File exists", func(t *testing.T) {
		exists := FileExists("fs_test.go")
		if !exists {
			t.Error("File does not exist")
		}
	})

	t.Run("File does not exist", func(t *testing.T) {
		exists := FileExists("fs_test.go.not")
		if exists {
			t.Error("File exists")
		}
	})
}

func TestDirExists(t *testing.T) {
	t.Run("Directory exists", func(t *testing.T) {
		exists := DirExists("data")
		if !exists {
			t.Error("Directory does not exist")
		}
	})

	t.Run("Directory does not exist", func(t *testing.T) {
		exists := DirExists("data.not")
		if exists {
			t.Error("Directory exists")
		}
	})
}

func TestCreateFile(t *testing.T) {
	err := CreateFile("data/test.txt")
	if err != nil {
		t.Error("Failed to create file")
	}
}
