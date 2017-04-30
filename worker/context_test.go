// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package worker;

import (
  "testing"
  "log"
  "github.com/mrichconf/feudal/message"
)

type testWorker struct {
  AbstractContext
}

func (w *testWorker) Receive(m message.Envelope) {
  switch m.Body().(type) {
  case string:
    m.Sender().Send(m)
  default:
    log.Fatal("unexpected message")
  }
}

func TestNew(t *testing.T) {
  nw := &testWorker{}
  nw.Init(nw)
  nw.Start(DefaultDispatcher(), nw)
  i := NewInterrogator()
  m := message.New("test message", i)
  nw.Send(m)
  r := <- i
  if r.Body() != "test message" {
    t.Errorf("invalid message")
  }
}
