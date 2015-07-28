// +build linux darwin freebsd

package shm

import (
	"os"
	"syscall"
	"testing"

	"github.com/pborman/uuid"
)

func TestCanCreateAndUseSharedRegion(t *testing.T) {
	var (
		expected = ([]byte)("a test")
		name     = "shm-test-" + uuid.New()
	)
	file, err := Open(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Error(err)
		}
		if err := Unlink(file.Name()); err != nil {
			t.Error(err)
		}
	}()
	if err := syscall.Ftruncate(
		int(file.Fd()), int64(len(expected)),
	); err != nil {
		t.Fatal(err)
	}

	if _, err := file.Write(expected); err != nil {
		t.Fatal(err)
	}
	if err := file.Sync(); err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, len(expected))
	if _, err := file.ReadAt(buf, 0); err != nil {
		t.Fatal(err)
	}
	if string(buf) != string(expected) {
		t.Fatalf("Expected %q. Got: %q", string(expected), string(buf))
	}
}
