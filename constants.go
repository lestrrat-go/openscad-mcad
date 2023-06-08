package mcad

import (
	"context"
	"io"

	"github.com/lestrrat-go/openscad/dsl"
)

func Constants(ctx context.Context, w io.Writer) error {
	tau := dsl.Variable("TAU").Value(6.2831853071)
	return dsl.Stmts(
		tau,
		dsl.Variable("PI").Value(dsl.Div(tau, 2)),
		dsl.Variable("mm_per_inch").Value(25.4),
	).EmitStmt(ctx, w)
}
