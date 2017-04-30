// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package worker;

import (
  "github.com/mrichconf/feudal/message"
)

type Dispatcher interface {
  Dispatch(m message.Envelope, r message.Receiver)
}

type DispatchReceive struct {}

func (d *DispatchReceive) Dispatch(m message.Envelope, r message.Receiver) {
  r.Receive(m)
}

func DefaultDispatcher() Dispatcher {
  return &DispatchReceive{}
}
