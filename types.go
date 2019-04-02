package msort

type LessFunc func(i, j int) bool

type SwapFunc func(i, j int)

type LenFunc func() int
