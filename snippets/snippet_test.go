package snippets_test

import (
	"log"
	"os"

	"github.com/SlalomBuild/fusion/snippets"
	"github.com/SlalomBuild/fusion/templates/aws"
)

func ExampleVSCode_NewItem() {
	// Generate a new snippet item from an existing
	// Go template
	v := snippets.VSCode{}
	v.NewItem("IAM Policy", aws.TEMPLATE_AWS_IAM_POLICY)
}

func ExampleVSCode_NewSnippetMap() {
	v := snippets.VSCode{}
	// Create a new snippet map
	sm := v.NewSnippetMap()

	// Build a list of snippets to add to the map
	s := []*snippets.VSCodeSnippet{
		v.NewItem("IAM Policy", aws.TEMPLATE_AWS_IAM_POLICY),
		v.NewItem("Lambda function", aws.TEMPLATE_AWS_LAMBDA_FUNCTION),
	}

	// Append all the snippets
	for _, snippet := range s {
		v.AddItem(&sm, snippet)
	}

	// Render the snippet file as json
	err := snippets.JSON(os.Stdout, s)
	if err != nil {
		log.Fatal(err)
	}

	// OR write the snippet json to a file
	f, err := os.Create("fusion-snippets.json")
	if err != nil {
		log.Fatal(err)
	}

	// The JSON method can write to stdout,stderr, a file
	// or anything else that implements io.Writer
	err = snippets.JSON(f, s)
	if err != nil {
		log.Fatal(err)
	}
}
