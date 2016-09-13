package trfile

import (
	"log"
	"os"
	"path"
	"testing"
	"time"
)

var testDir = "test"

func TestMakeDirFile(t *testing.T) {
	format := path.Join(testDir, "2006/01/02/Mon.log")
	w, err := NewFormat(format)

	if err != nil {
		t.Error(err)
	}

	log.SetOutput(w)
	log.Print("aaaaaaaaaaaaaaaaaaaa")

	if w.Name() != time.Now().Format(format) {
		t.Error(err)
	}

	info, err := os.Stat(w.Name())

	if err != nil {
		t.Error(err)
	}

	if info.Size() == 0 {
		t.Error(err)
	}

	if err := os.RemoveAll(testDir); err != nil {
		t.Error(err)
	}
}
