// +build !linux,!darwin,!freebsd

package shm

import (
	"os"
	"testing"
)

func TestNotSupported(t *testing.T) {
	file, err := Open(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != ErrPlatformNotSupported {
		t.Fatalf("Expected %q to be %q", err, ErrPlatformNotSupported)
	}
}
