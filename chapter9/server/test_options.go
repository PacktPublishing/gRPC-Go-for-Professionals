package main

type testOptions struct {
	isAvailable bool
}

var defaultTestOptions = testOptions{
	isAvailable: true,
}

type TestOption interface {
	apply(*testOptions)
}

type funcTestOption struct {
	f func(*testOptions)
}

func (fdo *funcTestOption) apply(do *testOptions) {
	fdo.f(do)
}

func newFuncTestOption(f func(*testOptions)) *funcTestOption {
	return &funcTestOption{
		f: f,
	}
}

func IsAvailable(a bool) TestOption {
	return newFuncTestOption(func(o *testOptions) {
		o.isAvailable = a
	})
}
