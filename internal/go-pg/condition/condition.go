package condition

import (
	"github.com/go-pg/pg/v10"
	"go/types"
)

type Condition interface {
	Apply(query *pg.Query)
}

func Apply(query *pg.Query, conditions ...Condition) {
	for _, c := range conditions {
		c.Apply(query)
	}
}

type Limit uint

func (limit Limit) Apply(query *pg.Query) {
	query.Limit(int(limit))
}

type Offset uint

func (offset Offset) Apply(query *pg.Query) {
	query.Offset(int(offset))
}

type None types.Nil

func (none None) Apply(_ *pg.Query) {}
