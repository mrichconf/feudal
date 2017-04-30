// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package worker;

import (
  "testing"
  "github.com/mrichconf/feudal/message"
)

func TestInterrogator(t *testing.T) {
  i := NewInterrogator()
  m := message.New("foo", i)
  go i.Send(m)
  r := <- i
  if r.Body() != "foo" {
    t.Errorf("Invalid message")
  }
}
