// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package worker;

import (
  "github.com/mrichconf/feudal/message"
//  "log"
)

type Context interface {
  message.Receiver
  message.Sender

  WorkerWhence(f Factory) Context
  Division() Division
  Manager() Ref
}

type Factory func() Context

type AbstractContext struct {
  Context
  Queue
  division Division
  manager Ref
}

func (c *AbstractContext) Init(i Context) {
  c.Context = i
  c.receive = make(chan message.Envelope)
  c.quit = make(chan bool)
  c.stopped = make(chan bool)
  c.division = NewDivision()
}

func (c *AbstractContext) WorkerWhence(f Factory) Context {
  worker := f()
  worker.(*AbstractContext).manager = c
  worker.(Runner).Start(DefaultDispatcher(), worker)
  c.division.Join(worker)
  return worker
}

func (c *AbstractContext) Division() Division {
  return c.division
}

func (c *AbstractContext) Send(m message.Envelope) {
//  c.receiveChannel <- m
  c.receive <- m
}
