package set

type Set struct {
	element map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		element: make(map[string]struct{}),
	}
}

func (set *Set) Add(element string) {
	set.element[element] = struct{}{}
}

func (set *Set) Remove(element string) {
	delete(set.element, element)
}

func (set *Set) Size() int {
	return len(set.element)
}
