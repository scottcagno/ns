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

// disk store
type Dsk struct {
	pth string
	sync.RWMutex
}

// init store
func InitDsk(pth string) Store {
	return &Dsk{
		pth: pth,
	}
}

// has method
func (self *Dsk) Has(dat *Dat) Dat {
	return Dat{}
}

// set method
func (self *Dsk) Set(dat *Dat) Dat {
	return Dat{}
}

// get method
func (self *Dsk) Get(dat *Dat) Dat {
	return Dat{}
}

// del method
func (self *Dsk) Del(dat *Dat) Dat {
	return Dat{}
}
