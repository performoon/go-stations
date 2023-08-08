package model

type ErrNotFound struct {
}

func (e *ErrNotFound) Error() string {
	errMessage := "errMessage"
	return errMessage
}
