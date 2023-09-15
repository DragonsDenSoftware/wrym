package errs

import (
	"errors"
	"fmt"
)

var FlagRequired = func(flag string) error {
	return fmt.Errorf("%s flag is required", flag)
}

var FlagRequiredForOtherFlag = func(flag, other string) error {
	return fmt.Errorf("%s flag required when using %s flag", flag, other)
}

var FlagRequiredForCommand = func(flag, command string) error {
	return fmt.Errorf("%s flag required when using %s command", flag, command)
}

type Iterable interface {
	[]any | map[any]any | *string
}

func Iterate[i Iterable](v i) (itr *iterator[i], err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()

	// TODO: check if iterable is map and
	// TODO: loop through keys and make
	// TODO: sure they're comparable
	// TODO: if not, panic and recover
	// TODO: error stating keys not
	// TODO: comparable

	switch any(v).(type) {
	case []any:
		itr.inType = Slice
	case map[any]any:
		itr.inType = Map
	case *string:
		itr.inType = String
	}

	itr.in = v

	return itr, nil
}

type iterator[i Iterable] struct {
	in i
	inType
	out    i
	filter func(e Elem) error
	mapFn  func(e Elem)
}

type inType int64

const (
	Slice inType = iota
	Map
	String
)

type Elem any

func (itr *iterator[i]) Filter(fn func(e Elem) error) *iterator[i] {
	var out i

	switch itr.inType {
	case Slice:
		inSlice := any(itr.in).([]any)

		outSlice := make([]any, 0, len(inSlice))

		for idx := range inSlice {
			if err := fn(inSlice[idx].(Elem)); err == nil {
				outSlice = append(outSlice, inSlice[idx])
			}
		}

		out = any(outSlice).(i)
	case Map:
		inMap := any(itr.in).(map[any]any)

		outMap := make(map[any]any)

		for k, v := range inMap {
			if err := fn(v.(Elem)); err == nil {
				outMap[k] = v
			}
		}

		out = any(outMap).(i)
	case String:
		inStringSlice := []byte(*any(itr.in).(*string))

		outStringSlice := make([]byte, 0, len(inStringSlice))

		for idx := range inStringSlice {
			if err := fn(any(inStringSlice[idx]).(Elem)); err == nil {
				outStringSlice = append(outStringSlice, inStringSlice[idx])
			}
		}

		outString := string(outStringSlice)

		out = any(&outString).(i)
	}

	itr.out = out

	return itr
}

func (itr *iterator[i]) ResSlice() []any {

	switch itr.inType {
	case Slice, String:

	}

	return any(itr.out).([]any)
}

func (itr *iterator[i]) ResMap() map[any]any {
	return any(itr.out).(map[any]any)
}

func (itr *iterator[i]) ResString() *string {
	return any(itr.out).(*string)
}
