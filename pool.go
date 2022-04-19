package syncx

import "sync"

type Pool[T any] struct {
	New   func() *T
	Reset func(*T)
	p     sync.Pool
}

func (p *Pool[T]) Get() *T {
	iface := p.p.Get()
	if iface == nil {
		if p.New == nil {
			return new(T)
		} else {
			return p.New()
		}
	}
	x := iface.(*T)
	if p.Reset != nil {
		p.Reset(x)
	}
	return x
}

func (p *Pool[T]) Put(x *T) {
	p.p.Put(x)
}
