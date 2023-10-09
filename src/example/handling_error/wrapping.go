package handling_error

import (
	"errors"
	"fmt"
	"os"
)

func (se StatusErr) Unwrap() error {
	return se.InternalError
}

func checkFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		// The %w verb is used to wrap an error. the error returned by
		// fmt.Errorf will have an Unwrap method returning the argument of %w,
		// which must be an error. In all other ways, %w is identical to %v.
		return fmt.Errorf("[fileChecker] %w", err)
	}
	f.Close()
	return nil
}

func login(uid, pwd, filePath string) ([]byte, error) {
	login := func(uid, pwd string) error {
		return nil // Pretending as success
		// return os.ErrPermission
	}
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:        InvalidLogin,
			Message:       fmt.Sprintf("invalid credentials for user %s", uid),
			InternalError: err,
		}
	}

	getData := func(file string) ([]byte, error) {
		// return []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}, nil // Pretending as success
		return nil, os.ErrNotExist
	}
	data, err := getData(filePath)
	if err != nil {
		return nil, StatusErr{
			Status:        NotFound,
			Message:       filePath,
			InternalError: err,
		}
	}

	return data, nil
}

func ExplicitUnwrap() {
	err := checkFile("not_here.txt")
	if err != nil {
		fmt.Printf("[Wrapping] %v\n", err)
		if err = errors.Unwrap(err); err != nil {
			fmt.Printf("[Internal] %v\n", err)
		}
	}
}

func IsAndAs() {
	// The errors.Is behaves like a comparison(==) to a sentinel error, and the
	// errors.As behaves like a type assertion. When operating on wrapped
	// errors, however, these functions consider all the errors in a chain.

	err := checkFile("not_here.txt")
	if errors.Is(err, os.ErrNotExist) { // The errors.Is compares an error to a value.
		fmt.Println("os.ErrNotExist caught")
	} else {
		panic("Unknown error: " + err.Error())
	}

	data, err := login("joe", "1234", "data.txt")
	if err != nil {
		var se StatusErr

		if errors.As(err, &se) { // The errors.As tests whether an error is a specific type.
			fmt.Printf("%v: %s\n", se.InternalError, se.Message)
		} else {
			panic("Unknown error: " + err.Error())
		}
		return
	}
	fmt.Println(string(data))
}
