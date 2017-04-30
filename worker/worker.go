// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package worker

import(
  "github.com/golang/protobuf/proto"
)

// Common worker interface

type Task interface {
  Handler(worker Worker, msg proto.Message) (proto.Message, error)
}

// On Start() a worker should (optionally) spin up a go routine to handle requests.
type Worker interface {
  Start(task Task) error
  GetResponseChannel() chan proto.Message
}

// wa interface
// onReceive
// unhandled
// getSender().tell() or getContext().reply()
// send message

type WorkerFactory interface {
  NewWorker(connection interface{}) (Worker, error)
}
