package msort

import (
	"reflect"
	"sort"
)

type Builder struct {
	SwapFn SwapFunc

	LenFn LenFunc

	lessFns []LessFunc
}

func NewBuilderFrom(slice interface{}) Builder {
	// panic if slice is not a slice
	swapFn := reflect.Swapper(slice)
	rv := reflect.ValueOf(slice)
	return Builder{
		SwapFn: swapFn,
		LenFn:  func() int { return rv.Len() },
	}
}

func (b Builder) ByFunc(fn LessFunc) Builder {
	b.lessFns = append(b.lessFns, fn)
	return b
}

func (b Builder) ToSorter() sort.Interface {
	if len(b.lessFns) == 0 {
		panic("msort: no LessFunc registered")
	}
	return &multiSort{
		swapFn:  b.SwapFn,
		lenFn:   b.LenFn,
		lessFns: b.lessFns,
	}
}

func Less(i, j int, l ...LessFunc) bool {
	for _, fn := range l {
		if fn(i, j) {
			return true
		} else if fn(j, i) {
			return false
		}
	}
	return false
}

type multiSort struct {
	swapFn SwapFunc

	lenFn LenFunc

	lessFns []LessFunc
}

func (ms *multiSort) Less(i, j int) bool {
	return Less(i, j, ms.lessFns...)
}

func (ms *multiSort) Swap(i, j int) {
	ms.swapFn(i, j)
}

func (ms *multiSort) Len() int {
	return ms.lenFn()
}

var _ sort.Interface = (*multiSort)(nil)
