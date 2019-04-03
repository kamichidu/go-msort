package compare

type NilOrder interface {
	compare(bool, bool) int
}

var (
	NilFirst NilOrder = &nilFirst{}
	NilLast  NilOrder = &nilLast{}
)

type nilFirst struct{}

func (*nilFirst) compare(aIsNil, bIsNil bool) int {
	switch {
	case !aIsNil && bIsNil:
		return 1
	case aIsNil && !bIsNil:
		return -1
	default:
		return 0
	}
}

type nilLast struct{}

func (*nilLast) compare(aIsNil, bIsNil bool) int {
	switch {
	case !aIsNil && bIsNil:
		return -1
	case aIsNil && !bIsNil:
		return 1
	default:
		return 0
	}
}

//go:generate gobst -o ./numbers.go -var-file ./numbers.json ./numbers.go.tmpl
//go:generate gofmt -w ./numbers.go
//go:generate gobst -o ./numbers_test.go -var-file ./numbers.json ./numbers_test.go.tmpl
//go:generate gofmt -w ./numbers_test.go
