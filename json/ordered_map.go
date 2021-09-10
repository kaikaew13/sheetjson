package json

import "sort"

type OrderedMap struct {
	M      map[string][]interface{}
	KeyLen int
	ValLen int
	It     Iterator
}

type Iterator struct {
	cur      int
	iterable []string
}

func newIterator(iterable []string) *Iterator {
	it := &Iterator{
		iterable: iterable,
		cur:      0,
	}

	return it
}

func (it *Iterator) Next() (cur string, ok bool) {
	if it.cur == len(it.iterable) {
		return
	}

	tmp := it.iterable[it.cur]
	it.cur++
	return tmp, true
}

func newOrderedMap(uJSON UnmarshalledJSON) *OrderedMap {
	m := make(map[string][]interface{})

	for _, u := range uJSON {
		for k, v := range u {
			m[k] = append(m[k], v)
		}
	}

	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return &OrderedMap{
		M:      m,
		KeyLen: len(m),
		ValLen: len(uJSON) + 1,
		It:     *newIterator(keys),
	}
}
