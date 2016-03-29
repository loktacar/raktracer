package raktracer

import (
	"fmt"
)

type Ray struct {
	pos, dir Vector
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray{pos:%s dir:%s}", r.pos, r.dir)
}
