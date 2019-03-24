package berr

import "fmt"

func Errorf(format string, a ...interface{}) *BizError {
	err := new(BizError)
	if len(a) > 0 {
		err.Message = fmt.Sprintf(format, a)
	} else {
		err.Message = format
	}
	return err
}

type BizError struct {
	Message string
}

func (err *BizError) Error() string {
	return err.Message
}
