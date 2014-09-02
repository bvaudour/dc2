package lib

import (
	u "util"
)

type Memory map[string]Stack

func MInit() Memory {
	return make(map[string]Stack)
}

func (m *Memory) Exists(k string) bool {
	_, ok := (*m)[k]
	return ok
}

func (m *Memory) Get(k string) *Stack {
	out := (*m)[k]
	return &out
}

func (m *Memory) Put(k string, aa ...Any) {
	(*m)[k] = aa
}

func (m *Memory) Add(k string, aa ...Any) {
	if !m.Exists(k) {
		m.Put(k, aa...)
	} else {
		s := m.Get(k)
		s.Push(aa...)
	}
}

func (m *Memory) Last(k string) Any {
	return m.Get(k).Last()
}

func (m *Memory) Copy(k string) Stack {
	return m.Get(k).Copy()
}

func (m *Memory) Remove(k string) {
	delete(*m, k)
}

func (m *Memory) Str(k string) string {
	s := m.Get(k)
	return u.Format("%s → %s", k, s.Str())
}

func (m *Memory) Sci(k string) string {
	s := m.Get(k)
	return u.Format("%s → %s", k, s.Sci())
}

func (m *Memory) Str0() string {
	out := ""
	for k, _ := range *m {
		u.Concat(&out, "\n", m.Str(k))
	}
	return out
}

func (m *Memory) Sci0() string {
	out := ""
	for k, _ := range *m {
		u.Concat(&out, `\n`, m.Sci(k))
	}
	return out
}
