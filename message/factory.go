// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package message;

import (
  "github.com/golang/protobuf/proto"
)

type Factory interface {
  NewMessage() proto.Message
}
