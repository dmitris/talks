package ioplus

import (
	"context"
	"io/ioutil"
	"time"
)

func ReadFile(ctx context.Context, filename string) (data []byte, err error) { // HL
	// if there is no deadline, set to the default
	var cancelFn context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancelFn = context.WithTimeout(ctx, 1*time.Second) // HL
		defer cancelFn()
	}
	type retType struct {
		data []byte
		err  error
	}

	done := make(chan retType, 1)

	go func() {
		data, err := ioutil.ReadFile(filename)
		done <- retType{data, err}
	}() // define and call

	// Now, wait for job to finish, or for a deadline to expire.
	select {
	case <-ctx.Done(): // HL
		return []byte{}, ctx.Err()
	case ret := <-done:
		return ret.data, ret.err
	}
}
