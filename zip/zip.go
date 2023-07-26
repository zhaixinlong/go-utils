package zip

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func ZipFile(sourcePath string) string {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	zipFileName := sourcePath + ".gz"
	gzipFile, err := os.Create(zipFileName)
	if err != nil {
		panic(err)
	}
	defer gzipFile.Close()

	writer := gzip.NewWriter(gzipFile)
	_, err = io.Copy(writer, sourceFile)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	log.Printf("zip finished, sourcePath: %s", sourcePath)
	return fmt.Sprintf("./%s", zipFileName)
}
