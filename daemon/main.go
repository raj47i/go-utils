package daemon

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

var (
	// ErrAlreadyRunning -- an instance is already running
	ErrAlreadyRunning = errors.New("the program is already running")
)

func lockFile(name string) string {
	return fmt.Sprintf("/var/lock/%s.lock", name)
}

// Run Exclusively
func Run(appName string) (*os.File, error) {
	// open/create lock file
	f, err := os.OpenFile(lockFile(appName), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	// set the lock type to F_WRLCK, therefor the file has to be opened writable
	flock := syscall.Flock_t{
		Type: syscall.F_WRLCK,
		Pid:  int32(os.Getpid()),
	}
	// try to obtain an exclusive lock - FcntlFlock seems to be the portable *ix way
	if err := syscall.FcntlFlock(f.Fd(), syscall.F_SETLK, &flock); err != nil {
		return nil, ErrAlreadyRunning
	}
	return f, nil
}

// Exit gracefully
func Exit(f *os.File) error {
	// set the lock type to F_UNLCK
	flock := syscall.Flock_t{
		Type: syscall.F_UNLCK,
		Pid:  int32(os.Getpid()),
	}
	if err := syscall.FcntlFlock(f.Fd(), syscall.F_SETLK, &flock); err != nil {
		return fmt.Errorf("failed to unlock the lock file: %v", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close the lock file: %v", err)
	}
	if err := os.Remove(lockFile(f.Name())); err != nil {
		return fmt.Errorf("failed to remove the lock file: %v", err)
	}
	return nil
}
