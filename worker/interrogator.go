// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package worker;

import (
  "github.com/mrichconf/feudal/message"
)

type Interrogator chan message.Envelope

func (i Interrogator) Send(m message.Envelope) {
  i <- m
}

func NewInterrogator() Interrogator {
  return make(chan message.Envelope)
}
