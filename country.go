package strategy1

import (
	"strategy1/government"
	"strategy1/territory"
)

type Country struct {
	Government *government.Parliament
	States     []*territory.State
}
