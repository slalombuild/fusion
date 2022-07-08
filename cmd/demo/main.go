package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

type Demo struct {
	Run         *demo.Run
	Name        string
	Description string
}

func main() {
	Start("fusion", demoFusion(), demoFusionctl())
}

func setup(ctx *cli.Context) error {
	// Ensure can be used for easy sequential command execution
	err := demo.Ensure(
		"echo 'Checking GO environment...'",
		"go env",
		"echo 'Installing fusion CLI'",
		"make install",
	)
	if err != nil {
		return errors.Wrap(err, "failed to validate GO environment - please ensure GO is installed and setup on your system")
	}

	content := `
	resource "foo" "bar" {
		name = "baz"
	}
	`
	err = os.WriteFile("test.tf", []byte(content), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to create demo terraform file 'test.tf'")
	}

	return nil
}

func cleanup(ctx *cli.Context) error {
	err := os.Remove("test.tf")
	if err != nil {
		return errors.Wrap(err, "failed to delete demo terraform file 'test.tf'")
	}
	return nil
}

func Start(name string, demos ...Demo) {
	demoRunner := demo.New()
	demoRunner.Name = name
	demoRunner.HideVersion = true

	// Register the demo run
	for _, d := range demos {
		demoRunner.Add(d.Run, d.Name, d.Description)
	}

	// Run the application, which registers all signal handlers and waits for
	// the app to exit
	demoRunner.Setup(setup)
	demoRunner.Cleanup(cleanup)
	demoRunner.Run()
}

// demoFusion is an example usage of the fusion CLI
// application
func demoFusion() Demo {
	// A new run contains a title and an optional description
	r := demo.NewRun(
		"Fusion 🧬",
		"Generate secure by default terraform templates",
	)

	// Show help command
	r.Step(demo.S(
		"See usage",
	), demo.S(
		"fusion --help",
	))

	// Generate a new terraform resource
	r.Step(demo.S(
		"Generate a terraform resource",
	), demo.S(
		"fusion new resource aws lambda",
	))

	// Generate a new terraform resource
	r.Step(demo.S(
		"See available options for resource",
	), demo.S(
		"fusion new resource aws lambda --help",
	))

	r.Step(
		demo.S("Customize generated resource"),
		demo.S("fusion new resource aws lambda --filename 'foo.zip' --handler 'index.handler' --runtime='go1.x' --verbose --no-color"),
	)

	// It is also not needed at all to provide a command
	r.Step(demo.S(
		"Notice how the fields in the above terraform configuration",
		"updated automatically based on our provided flags",
	), nil)

	r.Step(demo.S(
		"More functionality will be coming soon such as stacks and multi-cloud provider support.",
		"Stay tuned! 🧬",
	), nil)

	return Demo{r, "fusion", "Generate secure by default terraform templates"}
}

// demoFusionctl is an example usage of the fusionctl CLI
// application
func demoFusionctl() Demo {
	// A new run contains a title and an optional description
	r := demo.NewRun(
		"Fusionctl 🔨",
		"Generate code and import terraform resources into fusion",
	)

	// Show help command
	r.Step(demo.S(
		"See usage",
	), demo.S(
		"fusionctl --help",
	))

	// Generate a new terraform resource
	r.Step(demo.S(
		"Generate a new fusion command for your resource",
		"INFO: To auto-save the generated files in their correct locations, use the --save flag",
	), demo.S(
		"fusionctl new resource eks_cluster --provider=gcp --import='./test.tf' -v",
	))

	r.Step(demo.S("Notice how the fusionctl command generated a Go template, Go source, and a fusion CLI implementation for your new resource"), nil)
	r.Step(demo.S("fusionctl is the fastest way to get involved contributing to the fusion project ⚡"), nil)

	return Demo{r, "fusionctl", "Generate code and import terraform resources into fusion"}
}
