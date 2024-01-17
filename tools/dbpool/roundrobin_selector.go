package dbpool

import (
	"database/sql"
	"sync/atomic"
)

type RoundRobin struct {
	next uint32
}

func (r *RoundRobin) Next(dBs []*sql.DB) *sql.DB {
	n := atomic.AddUint32(&r.next, 1)
	return dBs[(int(n)-1)%len(dBs)]
}

func NewRoundRobinSelector() *RoundRobin {
	return &RoundRobin{}
}
