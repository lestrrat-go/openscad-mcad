package mcad

import (
	"github.com/lestrrat-go/openscad"
	"github.com/lestrrat-go/openscad/dsl"
)

func init() {
	openscad.Register("curves.scad", Curves())
}

func Curves() openscad.Stmts {
	pitch := dsl.Variable("pitch")
	z := dsl.Variable("z")
	radius := dsl.Variable("radius")
	return dsl.Stmts(
		dsl.Use("math.scad"),
		dsl.Include("constants.scad"),

		dsl.Function("b").
			Parameters(pitch).
			Body(dsl.Div(pitch, dsl.Variable("TAU"))),

		dsl.Function("t").
			Parameters(pitch, z).
			Body(dsl.Div(z, dsl.Call("b", pitch))),

		dsl.Function("helix_curve").
			Parameters(pitch, radius, z).
			Body(
				dsl.List(
					dsl.Mul(radius, dsl.Cos(dsl.Call("deg", (dsl.Call("t", pitch, z))))),
					dsl.Mul(radius, dsl.Sin(dsl.Call("deg", (dsl.Call("t", pitch, z))))),
					z,
				),
			),
	)
}
