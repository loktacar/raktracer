package raktracer

import (
	"fmt"
)

type Ray struct {
	Pos, Dir Vector
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray{Pos:%s Dir:%s}", r.Pos, r.Dir)
}
