package collectors

type Collector interface {
	Name() string
	Collect() (float64, error)
}
