package timeoutwrapper

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var (
	wg WorkerGroup

	errorWrapperChan chan error
	errorWrapperResp error

	responseChan chan []reflect.Value
	response     []reflect.Value
)

func timelimit(d time.Duration) {
	start := time.Now()
	defer func() {
		recover()
	}()
	time.Sleep(d)
	errorWrapperChan <- errors.New(fmt.Sprint("error timeout tooks ", time.Since(start)))
}

func operation() {
	defer wg.FinishAllWorkers()
	for {
		select {
		case errorWrapperResp = <-errorWrapperChan:
			return
		case response = <-responseChan:
			return
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func Call(timeout time.Duration, f interface{}, args ...interface{}) (resp interface{}, err error) {
	responseChan = make(chan []reflect.Value)
	errorWrapperChan = make(chan error)

	defer func() {
		close(responseChan)
		close(errorWrapperChan)
		errorWrapperResp = nil
		response = nil
	}()

	wg.Add(3)
	go operation()
	go publisher(f, args...)
	go timelimit(timeout)

	wg.Wait()

	if errorWrapperResp != nil {
		return nil, errorWrapperResp
	}

	// Iterate through the results to check their types
	for _, r := range response {
		if r.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			// The value is of type error
			if !r.IsNil() {
				err = r.Interface().(error)
			} else {
				err = nil
			}
		} else {
			// The value is not an error
			resp = r.Interface()
		}
	}

	return
}
