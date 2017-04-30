// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package worker;

type Division interface {
  Join(c Context)
  Disolve()
}

type Members struct {
  workers []Context
}

func NewDivision() Division {
  return &Members{ workers: make([]Context, 0) }
}

func (m *Members) Join(c Context) {
  m.workers = append(m.workers, c)
}

func (m *Members) Disolve() {
  for _,w := range m.workers {
    w.(Runner).Stop()
  }
  m.workers = nil
}
