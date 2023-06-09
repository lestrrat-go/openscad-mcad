package mcad

import (
	"github.com/lestrrat-go/openscad"
	"github.com/lestrrat-go/openscad/dsl"
)

func init() {
	openscad.Register("constants.scad", Constants())
}

func Constants() openscad.Stmt {
	tau := dsl.Variable("TAU").Value(6.2831853071)
	return dsl.Stmts(
		tau,
		dsl.Variable("PI").Value(dsl.Div(tau, 2)),
		dsl.Variable("mm_per_inch").Value(25.4),
	)
}
