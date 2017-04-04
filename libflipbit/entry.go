package libflipbit

type Entry struct {
	Name string
	Namespace string
	Ports Ports
	Hosts []string
	LoadBalancers []string
	Remained bool
	Changed bool
}

type Entries []Entry
