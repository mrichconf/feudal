// Copyright 2017 Matthew Rich <matthewrich.conf@gmail.com>. All rights reserved.

package message;

type Empty struct {}

type Register struct { Handler interface{} }

type Terminated struct {}

type Serialized struct { Content []byte }

type Unserialize struct { Size int }

type Error struct { e error }
