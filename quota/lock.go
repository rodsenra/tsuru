// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quota

import "sync"

type multiLocker struct {
	m   map[string]*sync.Mutex
	mut sync.Mutex
}

func (l *multiLocker) Lock(name string) {
	l.mut.Lock()
	defer l.mut.Unlock()
	_, ok := l.m[name]
	if !ok {
		l.m[name] = new(sync.Mutex)
	}
	l.m[name].Lock()
}

func (l *multiLocker) Unlock(name string) {
	l.mut.Lock()
	defer l.mut.Unlock()
	l.m[name].Unlock()
}