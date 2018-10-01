package adapter

type applyFunc func(changeName string) error
type createFunc func(map[string]string) (Adapter, error)
type markFunc func() error
type saveFunc func(changeName string) error
type statusFunc func() error

var adapterMap = map[string]createFunc{
	"k8s": k8sCreate,
}

// Adapter defines the standard type for adapters
type Adapter interface {
	Apply(changeName string) error
	Mark() error
	Save(changeName string) error
	Status() error
}

// Create an adapter
func Create(name string, options map[string]string) (Adapter, error) {
	return adapterMap[name](options)
}
