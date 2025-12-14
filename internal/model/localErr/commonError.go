package localErr

type CommonError struct {
	Code   int
	ErrMsg string
}

func (c *CommonError) Error() string {
	return c.ErrMsg
}
