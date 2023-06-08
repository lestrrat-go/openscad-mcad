package mcad_test

import (
	"context"
	"os"

	mcad "github.com/lestrrat-go/openscad-mcad"
)

func ExampleScrew() {
	mcad.Screw(context.Background(), os.Stdout)
}
