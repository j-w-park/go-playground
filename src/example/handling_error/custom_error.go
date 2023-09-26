package handling_error

import (
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func throwIf1(flag bool) error {
	var genErr StatusErr // genErr is zero value of the struct StatusErr, not nil.
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func throwIf2(flag bool) error {
	if flag {
		return StatusErr{Status: NotFound}
	}
	return nil
}

func throwIf3(flag bool) error {
	var genErr error // genErr is zero value of the interface error which is nil.
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func CustomError() {
	err := throwIf1(true)
	fmt.Println(err != nil) // true

	err = throwIf1(false)
	fmt.Println(err != nil) // true

	err = throwIf2(false)
	fmt.Println(err != nil) // false

	err = throwIf3(false)
	fmt.Println(err != nil) // false
}
