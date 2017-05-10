package ctx_range

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

func Range(ctx context.Context, in interface{}, out interface{}) {
	rIn := reflect.ValueOf(in)
	if rIn.Kind() != reflect.Chan {
		panic(fmt.Errorf("expected Range input to be a channel, got %s", rIn.Kind()))
	}

	rOut := reflect.ValueOf(out)
	if rOut.Kind() != reflect.Chan {
		panic(fmt.Errorf("expected Range output to be a channel, got %s", rOut.Kind()))
	}

	inOf := rIn.Type().Elem()
	outOf := rOut.Type().Elem()
	if inOf != outOf {
		panic(fmt.Errorf("inOf (%s) and outOf (%s) must channels of the same element type", inOf, outOf))
	}

	if ctx.Done() == nil {
		panic("context should be cancelable")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			index, val, isOpen := reflect.Select([]reflect.SelectCase{
				reflect.SelectCase{
					Dir:  reflect.SelectRecv,
					Chan: reflect.ValueOf(ctx.Done()),
				},
				reflect.SelectCase{
					Dir:  reflect.SelectRecv,
					Chan: rIn,
				},
			})

			if !isOpen || index == 0 {
				return
			}

			if index == 1 {
				rOut.Send(val)
			}
		}
	}()

	go func() {
		wg.Wait()
		for {
			if rOut.Len() == 0 {
				rOut.Close()
				return
			}
		}
	}()
}
