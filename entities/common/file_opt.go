package common

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func CopyAndDelete(src, dst string) error {
	// Copy file
	err := CopyFile(src, dst)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	// Hapus file lama
	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
