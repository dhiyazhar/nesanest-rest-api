package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func CurrentUnixNano() int64 {
	return time.Now().UnixNano()
}

func SaveUploadedFile(request *http.Request, fieldName string, uploadDir string) (string, error) {
	file, handler, err := request.FormFile(fieldName)
	if err != nil {
		return "", fmt.Errorf("failed to get file from form field '%s': %w", fieldName, err)
	}
	defer file.Close()

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create upload directory: %w", err)
		}
	}

	fileName := fmt.Sprintf("%d%s", CurrentUnixNano(), filepath.Ext(handler.Filename))
	fullPath := filepath.Join(uploadDir, fileName)

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to copy file data: %w", err)
	}

	return fullPath, nil
}
