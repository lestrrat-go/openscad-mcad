package mcad

import (
	"github.com/lestrrat-go/openscad"
	"github.com/lestrrat-go/openscad/dsl"
)

func Deg() openscad.Stmt {
	angle := dsl.Variable("angle")
	return dsl.Function("deg").
		Parameters(angle).
		Body(dsl.Mul(360, dsl.Div(angle, dsl.Variable("TAU"))))
}
