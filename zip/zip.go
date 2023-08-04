package zip

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func ZipFile(sourcePath string) (string, error) {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return "", err
	}
	defer sourceFile.Close()

	zipFilePath := sourcePath + ".gz"
	gzipFile, err := os.Create(zipFilePath)
	if err != nil {
		return "", err
	}
	defer gzipFile.Close()

	writer := gzip.NewWriter(gzipFile)
	_, err = io.Copy(writer, sourceFile)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}
	log.Printf("zip finished, sourcePath: %s, zipFilePath: %s", sourcePath, zipFilePath)
	return zipFilePath, nil
}
