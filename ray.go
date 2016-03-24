package raktracer

import (
	"fmt"
)

type ray struct {
	pos, dir vector
}

func (r ray) String() string {
	return fmt.Sprintf("ray{pos:%s dir:%s}", r.pos, r.dir)
}
