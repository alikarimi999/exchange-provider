package try

import "errors"

// MaxRetries is the maximum number of retries before bailing.

var errMaxRetriesReached = errors.New("exceeded retry limit")

// Func represents functions that can be retried.
type Func func(attempt uint64) (retry bool, err error)

// Do keeps trying the function until the second argument
// returns false, or no error is returned.
func Do(maxRetries uint64, fn Func) error {
	var err error
	var cont bool
	attempt := uint64(1)
	for {
		cont, err = fn(attempt)
		if !cont || err == nil {
			break
		}
		attempt++
		if attempt > maxRetries {
			return errMaxRetriesReached
		}
	}
	return err
}

// IsMaxRetries checks whether the error is due to hitting the
// maximum number of retries or not.
func IsMaxRetries(err error) bool {
	return err == errMaxRetriesReached
}
