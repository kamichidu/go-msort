package compare

import "reflect"

type NilOrder interface {
	compare(interface{}, interface{}) int
}

var (
	NilFirst NilOrder = &nilFirst{}
	NilLast  NilOrder = &nilLast{}
)

type nilFirst struct{}

func (*nilFirst) compare(a, b interface{}) int {
	arv := reflect.ValueOf(a)
	brv := reflect.ValueOf(b)
	switch {
	case !arv.IsNil() && !brv.IsNil():
		return 0
	case arv.IsNil():
		return -1
	default:
		return 1
	}
}

type nilLast struct{}

func (*nilLast) compare(a, b interface{}) int {
	arv := reflect.ValueOf(a)
	brv := reflect.ValueOf(b)
	switch {
	case !arv.IsNil() && !brv.IsNil():
		return 0
	case arv.IsNil():
		return 1
	default:
		return -1
	}
}

//go:generate gobst -o ./numbers.go -var-file ./numbers.json ./numbers.go.tmpl
//go:generate gofmt -w ./numbers.go
