package timeoutwrapper

import (
	"errors"
	"reflect"
)

func publisher(f interface{}, args ...interface{}) {
	fnValue := reflect.ValueOf(f)
	fnType := fnValue.Type()

	defer func() {
		recover()
	}()

	// Validate that the function is a function type
	if fnType.Kind() != reflect.Func {
		errorWrapperChan <- errors.New("provided argument is not a function")
		return
	}

	// Validate the number of arguments
	if len(args) != fnType.NumIn() {
		errorWrapperChan <- errors.New("the number of arguments does not match the function signature")
		return
	}

	// Prepare the arguments
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	// Call the function with arguments
	results := fnValue.Call(in)

	if len(results) > 2 {
		errorWrapperChan <- errors.New("return function limited only 2 return value, response and error")
		return
	}

	// If the function has return values, capture them
	responseChan <- results
}
