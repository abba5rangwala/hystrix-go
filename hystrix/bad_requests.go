package hystrix

import "fmt"

type BadRequestsErr struct {
	err error
}

func New(s string) *BadRequestsErr {
	return &BadRequestsErr{
		err: CircuitError{"bad request error: " + s},
	}
}

func Wrap(err error) *BadRequestsErr {
	return &BadRequestsErr{
		err: fmt.Errorf("hystrix: bad request error: %w", err),
	}
}

func (b *BadRequestsErr) Error() string {
	return b.err.Error()
}

func isBadRequestError(err error) bool {
	_, ok := err.(*BadRequestsErr)
	return ok
}
