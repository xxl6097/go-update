package update

type Middler interface {
	Middle(binpath string) error
}
