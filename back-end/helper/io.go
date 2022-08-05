package helper

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn((len(letters)))]
	}
	return string(b)
}

func UploadImage(file multipart.File, header *multipart.FileHeader) string {
	dir := "assets/image"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Println(fmt.Sprintf("creating dir : %s", err.Error()))
		}
	}
	imageName := randSeq(20) + filepath.Ext(header.Filename)

	fileLocation := filepath.Join(dir, imageName)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		log.Println(err.Error())
		return ""
	}

	return fileLocation
}
