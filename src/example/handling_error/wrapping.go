package handling_error

import (
	"errors"
	"fmt"
	"os"
)

func (se StatusErr) Unwrap() error {
	return se.InternalError
}

func Wrapping() {
	fileChecker := func(name string) error {
		f, err := os.Open(name)
		if err != nil {
			return fmt.Errorf("[fileChecker] %w", err)
		}
		f.Close()
		return nil
	}

	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		if err = errors.Unwrap(err); err != nil {
			fmt.Println(err)
		}
	}
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	login := func(uid, pwd string) error {
		return nil
	}

	getData := func(file string) ([]byte, error) {
		return []byte{}, nil
	}

	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:        InvalidLogin,
			Message:       fmt.Sprintf("invalid credentials for user %s", uid),
			InternalError: err,
		}
	}

	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:        NotFound,
			Message:       fmt.Sprintf("file %s not found", file),
			InternalError: err,
		}
	}

	return data, nil
}
