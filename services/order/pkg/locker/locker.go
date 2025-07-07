package locker

import (
	"errors"
	"time"

	"github.com/go-redsync/redsync/v4"
)

type Locker struct {
	mutex *redsync.Mutex
}

func NewLocker(
	Name string,
	Expiry time.Duration,
	Tries int,
	RetryDelay time.Duration,
) *Locker {
	return &Locker{
		mutex: GetRedsync().NewMutex(
			Name,
			redsync.WithExpiry(Expiry),
			redsync.WithTries(Tries),
			redsync.WithRetryDelay(RetryDelay),
		),
	}
}

func (l *Locker) Acquire() error {
	if l.mutex == nil {
		return errors.New("mutex is not initialized")
	}
	return l.mutex.Lock()
}

func (l *Locker) Release() (bool, error) {
	if l.mutex == nil {
		return false, errors.New("mutex is not initialized")
	}
	return l.mutex.Unlock()
}
