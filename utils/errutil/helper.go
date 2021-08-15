package errutil

func InternalError(err error) error {
	return ErrInternal.Msg(err.Error())
}
