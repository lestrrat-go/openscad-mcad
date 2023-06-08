package mcad

import (
	"context"
	"fmt"
	"io"

	"github.com/lestrrat-go/openscad/dsl"
)

func Screw(ctx context.Context, w io.Writer) error {
	if err := Constants(ctx, w); err != nil {
		return err
	}
	fmt.Fprint(w, "\n")
	if err := Deg(ctx, w); err != nil {
		return err
	}
	fmt.Fprint(w, "\n")
	if err := Helix(ctx, w); err != nil {
		return err
	}
	fmt.Fprint(w, "\n")
	if err := Auger(ctx, w); err != nil {
		return err
	}
	fmt.Fprint(w, "\n")
	if err := BallGroove(ctx, w); err != nil {
		return err
	}
	fmt.Fprint(w, "\n")
	if err := BallGroove2(ctx, w); err != nil {
		return err
	}
	return nil
}

func Helix(ctx context.Context, w io.Writer) error {
	pitch := dsl.Variable("pitch")
	length := dsl.Variable("length")
	slices := dsl.Variable("slices").Value(500)
	rotations := dsl.Variable("rotations").Value(
		dsl.Div(length, pitch),
	)
	return dsl.Module("helix").
		Parameters(pitch, length, slices).
		Actions(
			rotations,
			dsl.LinearExtrude(length, false, 10, dsl.Mul(360, rotations), slices).
				Fn(100).
				Add(dsl.Children()),
		).EmitStmt(ctx, w)
}

func Auger(ctx context.Context, w io.Writer) error {
	pitch := dsl.Variable("pitch")
	length := dsl.Variable("length")
	outsideRadius := dsl.Variable("outside_radius")
	innerRadius := dsl.Variable("inner_radius")
	taperRatio := dsl.Variable("taper_ratio").Value(0.25)

	return dsl.Module("auger").
		Parameters(pitch, length, outsideRadius, innerRadius, taperRatio).
		Actions(
			dsl.Union(
				dsl.Call("helix", pitch, length),
				dsl.Polygon(
					dsl.List(
						dsl.List(0, innerRadius),
						dsl.List(outsideRadius, dsl.Mul(innerRadius, taperRatio)),
						dsl.List(outsideRadius, dsl.Mul(dsl.Mul(innerRadius, -1), taperRatio)),
						dsl.List(0, dsl.Mul(-1, innerRadius)),
					),
					dsl.List(0, 1, 2, 3),
				),
				dsl.Cylinder(length, innerRadius, nil),
			),
		).EmitStmt(ctx, w)
}

func BallGroove(ctx context.Context, w io.Writer) error {
	pitch := dsl.Variable("pitch")
	length := dsl.Variable("length")
	diameter := dsl.Variable("diameter")
	ballRadius := dsl.Variable("ball_radius")

	return dsl.Module("ball_groove").
		Parameters(pitch, length, diameter, ballRadius).
		Actions(
			dsl.Call("helix", pitch, length, dsl.Variable("slices").Value(100)).
				Add(dsl.Translate(dsl.List(diameter, 0, 0)).
					Add(dsl.Circle(ballRadius))),
		).EmitStmt(ctx, w)
}

func BallGroove2(ctx context.Context, w io.Writer) error {
	pitch := dsl.Variable("pitch")
	length := dsl.Variable("length")
	diameter := dsl.Variable("diameter")
	ballRadius := dsl.Variable("ball_radius")
	slices := dsl.Variable("slices").Value(200)

	rotations := dsl.Variable("rotations").Value(dsl.Div(length, pitch))
	radius := dsl.Variable("radius").Value(dsl.Div(diameter, 2))
	offset := dsl.Variable("offset").Value(dsl.Div(length, slices))

	i := dsl.Variable("i")
	z := dsl.Variable("z").Value(dsl.Mul(i, offset))
	return dsl.Module("ball_groove2").
		Parameters(pitch, length, diameter, ballRadius, slices).
		Actions(
			rotations,
			radius,
			offset,
			dsl.Union(
				dsl.For(dsl.LoopVar(i, dsl.ForRange(0, slices))).
					Add(
						dsl.Let(z).
							Add(dsl.Translate(dsl.Call("helix_curve", pitch, radius, z)).
								Add(dsl.Sphere(ballRadius).Fa(5).Fs(1)),
							),
					),
			),
		).EmitStmt(ctx, w)

}
