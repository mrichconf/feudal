// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package message;

type Envelope interface {
  Body() interface{}
  Sender() Sender
}

type Receiver interface {
  Receive(m Envelope)
}
type Sender interface {
  Send(m Envelope)
}

type AbstractEnvelope struct {
  body interface{}
  sender Sender
}

func (e *AbstractEnvelope) Body() interface{} { return e.body }
func (e *AbstractEnvelope) Sender() Sender { return e.sender }

func New(b interface{}, s Sender) Envelope {
  return &AbstractEnvelope{ body: b, sender: s }
}
