package trfile

import (
	"os"
	"path"
	"time"
)

type TimeRotateFile struct {
	fileNameFormat string
	file           *os.File
}

var (
	FileMode    os.FileMode = 0644
	DirFileMode os.FileMode = 0755
)

func NewFormat(fileNameFormat string) (*TimeRotateFile, error) {
	trf := &TimeRotateFile{
		fileNameFormat: fileNameFormat,
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
	fileNameNow := trf.fileNameNow()

	if dir, _ := path.Split(fileNameNow); len(dir) > 0 {
		if err := os.MkdirAll(dir, DirFileMode); err != nil {
			return err
		}
	}

	file, err := os.OpenFile(fileNameNow, os.O_RDWR|os.O_CREATE|os.O_APPEND, FileMode)

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
