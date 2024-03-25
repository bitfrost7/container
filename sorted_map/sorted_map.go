package sorted_map

import "container/list"

type SortedMap[k comparable] struct {
	l *list.List
	m map[k]*list.Element
}

func NewSortedMap[k comparable]() *SortedMap[k] {
	return &SortedMap[k]{
		l: list.New(),
		m: make(map[k]*list.Element),
	}
}

func (m *SortedMap[k]) Insert(key k, v any) {
	e := m.l.PushBack(v)
	m.m[key] = e
}

func (m *SortedMap[k]) Delete(key k) {
	e, ok := m.m[key]
	if !ok {
		return
	}
	delete(m.m, key)
	m.l.Remove(e)
}

func (m *SortedMap[k]) Search(key k) (v any, exist bool) {
	e, ok := m.m[key]
	if !ok {
		return nil, false
	}
	return e.Value, true
}

func (m *SortedMap[k]) Len() int {
	return m.l.Len()
}

func (m *SortedMap[k]) Range(f func(v any)) {
	for e := m.l.Front(); e != nil; e = e.Next() {
		f(e.Value)
	}
}
