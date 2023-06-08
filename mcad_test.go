package mcad_test

import (
	"context"
	"os"

	"github.com/lestrrat-go/openscad"
)

func ExampleScrew() {
	stmts, ok := openscad.Lookup("screw.scad")
	if ok {
		stmts.EmitStmt(context.Background(), os.Stdout)
	}
	// OUTPUT:
	// include <curves.scad>
	//
	// module helix(pitch, length, slices=500)
	// {
	//   rotations=length/pitch;
	//   linear_extrude(height=length, center=false, convexity=10, twist=360*rotations, scale=slices, $fn=100)
	//     children();
	// }
	//
	// module auger(pitch, length, outside_radius, inner_radius, taper_ratio=0.25)
	// {
	//   union()
	//   {
	//     helix(pitch, length);
	//     polygon(points=[[0, inner_radius], [outside_radius, inner_radius*taper_ratio], [outside_radius, inner_radius*-1*taper_ratio], [0, -1*inner_radius]], paths=[0, 1, 2, 3]);
	//     cylinder(h=length, r=inner_radius);
	//   }
	// }
	//
	// module ball_groove(pitch, length, diameter, ball_radius)
	// {
	//   helix(pitch, length, slices)
	//     translate([diameter, 0, 0])
	//       circle(r=ball_radius);
	// }
	//
	// module ball_groove2(pitch, length, diameter, ball_radius, slices=200)
	// {
	//   rotations=length/pitch;
	//   radius=diameter/2;
	//   offset=length/slices;
	//   union()
	//     for(i=[0:slices]) {
	//       let(z=i*offset)
	//         translate(helix_curve(pitch, radius, z))
	//           sphere(r=ball_radius, $fa=5, $fs=1);
	//     }
	// }
	//
}
