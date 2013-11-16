// *
// * (c) Copyright 2013, Scott Cagno. All rights reserved.
// * This source code is free under the BSD style license.
// *
// * ns :: node store - a distributed data store
// *

package ns

import (
	"sync"
)

// memory store
type Mem struct {
	st map[string]interface{}
	sync.RWMutex
}

// init store
func InitMem() Store {
	return &Mem{
		st: make(map[string]interface{}),
	}
}

// has method
func (self *Mem) Has(dat *Dat) Dat {
	return Dat{}
}

// set method
func (self *Mem) Set(dat *Dat) Dat {
	return Dat{}
}

// get method
func (self *Mem) Get(dat *Dat) Dat {
	return Dat{}
}

// del method
func (self *Mem) Del(dat *Dat) Dat {
	return Dat{}
}
