# Contributing

> Contributing to fusion is welcome and encouraged! Got a terraform template that follows Secure Build practices and would be good for Fusion? 
> Contribute!

## Setup 🔧

Getting setup is pretty simple. Just make sure you've got Go installed and `$GOPATH/bin` is in your `PATH` and these next steps will
get you up and running.

```
# Install the dev tools
make tools

# Run the test suite
make test

# Install fusion and fusionctl to $GOPATH/bin
make install
```

## Adding a terraform resource to fusion 🧬

Terraform resources in fusion are very simple. A terraform resource is made up of a few things:

1. A [go text/template](https://pkg.go.dev/text/template) 
2. A struct for the template's configurable values (see [templates](./templates/))
3. A struct for the command to create the template that implements the `Run(ctx *commands.Context) error` method

That's it! This repository is designed intentionally to be simple. We even provide you with a dev tool to generate 99% of the code so all you need to bring is your terraform file.

## Using Fusionctl ⚡

1. Generate an example resource

```bash
# Let's generate an implementation of an AWS EC2 instance generator for fusion.

# This command outputs the following directly to Stdout:
# - generated Go code
# - generated template file
# - generated cli command

fusionctl new resource ec2_instance \
 --provider aws \
 --fields="name=string;description=string;vpc_id=string;ingress_from_port=int" \
 --verbose
```

2. Verify the output looks correct and has everything you want
3. Save the output

```bash
# Now that we've verified everything looks good, we're going to write the output directly
# into the project so you don't have to write any code!

# We do this with the `--save` flag.

# Navigate to your cloned instance of fusion
cd fusion/

# Generate the resource but with the --save flag
fusionctl new resource ec2_instance \
 --provider aws \
 --fields="name=string;description=string;vpc_id=string;ingress_from_port=int" \
 --verbose
 --save
```

4. Add your new command to it's associated provider struct (e.g. [aws commands](./internal/commands/awscmd/cmd_aws.go))
5. Try your command!

```bash
# Verify all the tests still pass
make test

# Install fusion again with your changes added
make install

# Try your new command
fusion new aws ec2_instance --name="example" --description="example" --vpc-id="1234" --ingress-from-port="8080"
```

6. Open a pull request with your new feature added! 🎉