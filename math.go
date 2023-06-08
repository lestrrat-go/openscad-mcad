package mcad

import (
	"context"
	"io"

	"github.com/lestrrat-go/openscad/dsl"
)

func Deg(ctx context.Context, w io.Writer) error {
	angle := dsl.Variable("angle")
	return dsl.Function("deg").
		Parameters(angle).
		Body(dsl.Mul(360, dsl.Div(angle, dsl.Variable("TAU")))).
		EmitStmt(ctx, w)
}
