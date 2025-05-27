package update

type Middler interface {
	Middle(binpath string) error
}

type MiddleFn func(string) error
