package errutil

import "errors"

func AsType[E error](err error) (E, bool) {
	return errors.AsType[E](err)
}
