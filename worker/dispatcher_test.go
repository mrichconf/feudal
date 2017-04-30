// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package worker;

import (
  "testing"
//  "log"
  "github.com/mrichconf/feudal/message"
)

type testCtx struct {}

func (c *testCtx) Receive(m message.Envelope) {
  m.Sender().Send(m) 
}
func (c *testCtx) Send(m message.Envelope) {}
func (c *testCtx)  WorkerWhence(f Factory) Context { return nil }
func (c *testCtx)  Division() Division { return nil }
func (c *testCtx)  Manager() Ref { return nil }

func TestRunner(t *testing.T) {
  q := NewQueue()
  q.Start(DefaultDispatcher(), &testCtx{})

  i := NewInterrogator()
  q.receive <- message.New("test message", i)
  r := <- i
  if r.Body() != "test message" {
    t.Errorf("invalid message")
  }
}

func TestRunnerStartStop(t *testing.T) {
  q := NewQueue()
  q.Start(DefaultDispatcher(), &testCtx{})

  select {
  case _, ok := <- q.receive:
    if ! ok {
      t.Errorf("receive channel closed")
    }
  default:
  }

  select {
  case _, ok := <- q.quit:
    if ! ok {
      t.Errorf("quit channel closed")
    }
  default:
  }

  select {
  case _, ok := <- q.stopped:
    if ! ok {
      t.Errorf("stopped channel closed")
    }
  default:
  }

  q.Stop()

  select {
  case _, ok := <- q.receive:
    if ok {
      t.Errorf("receive channel not closed")
    }
  default:
    t.Errorf("receive channel not closed")
  }

  select {
  case _, ok := <- q.quit:
    if ok {
      t.Errorf("quit channel not closed")
    }
  default:
    t.Errorf("quit channel not closed")
  }

  select {
  case _, ok := <- q.stopped:
    if ok {
      t.Errorf("stopped channel not closed")
    }
  default:
    t.Errorf("stopped channel not closed")
  }
}
