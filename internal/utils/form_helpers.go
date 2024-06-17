package utils

import (
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"time"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}
	return i
}

func StringToTime(s string) time.Time {
	t, err := time.Parse(time.DateTime, s)

	if err != nil {
		return time.Now().Add(24 * time.Hour)
	}
	return t
}

func UploadFile(file *multipart.FileHeader) (string, error) {
	fileHeader := file
	f, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	// f, err := imageProcessing(file, 20)

	uploaderURL, err := SaveFile(f, fileHeader)
	if err != nil {
		return "", err
	}
	return uploaderURL, nil

}

func FileHeaderToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	// Open the file from the multipart form
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the content of the file into a byte slice
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}