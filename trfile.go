package trfile

import (
	"os"
	"time"
)

type TimeRotateFile struct {
	fileNameFormat string
	file           *os.File
	FileMode       os.FileMode
}

func New(fileNameFormat string) (*TimeRotateFile, error) {
	trf := &TimeRotateFile{
		fileNameFormat: fileNameFormat,
		FileMode:       0666,
	}

	if err := trf.fileSet(); err != nil {
		return nil, err
	}

	return trf, nil
}

func (trf *TimeRotateFile) fileNameNow() string {
	return time.Now().Format(trf.fileNameFormat)
}

func (trf *TimeRotateFile) fileSet() error {
	file, err := os.OpenFile(trf.fileNameNow(), os.O_RDWR|os.O_CREATE|os.O_APPEND, trf.FileMode)

	if err != nil {
		return err
	}

	trf.file = file

	return nil
}

func (trf *TimeRotateFile) Write(b []byte) (int, error) {
	if fileNameNow := trf.fileNameNow(); fileNameNow != trf.file.Name() {
		if err := trf.file.Close(); err != nil {
			return 0, err
		}

		if err := trf.fileSet(); err != nil {
			return 0, err
		}
	}

	return trf.file.Write(b)
}

func (trf *TimeRotateFile) Name() string {
	return trf.file.Name()
}

func (trf *TimeRotateFile) Close() error {
	return trf.file.Close()
}
