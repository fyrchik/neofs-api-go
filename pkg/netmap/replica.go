package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

// Replica represents v2-compatible object replica descriptor.
type Replica netmap.Replica

// NewReplica creates and returns new Replica instance.
func NewReplica() *Replica {
	return NewReplicaFromV2(new(netmap.Replica))
}

// NewReplicaFromV2 converts v2 Replica to Replica.
func NewReplicaFromV2(f *netmap.Replica) *Replica {
	return (*Replica)(f)
}

// ToV2 converts Replica to v2 Replica.
func (r *Replica) ToV2() *netmap.Replica {
	return (*netmap.Replica)(r)
}

// Count returns number of object replicas.
func (r *Replica) Count() uint32 {
	return (*netmap.Replica)(r).
		GetCount()
}

// SetCount sets number of object replicas.
func (r *Replica) SetCount(c uint32) {
	(*netmap.Replica)(r).
		SetCount(c)
}

// Selector returns name of selector bucket to put replicas.
func (r *Replica) Selector() string {
	return (*netmap.Replica)(r).
		GetSelector()
}

// SetSelector sets name of selector bucket to put replicas.
func (r *Replica) SetSelector(s string) {
	(*netmap.Replica)(r).
		SetSelector(s)
}
