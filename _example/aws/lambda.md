# Example - Create a lambda

> Create a new aws lambda function


## Command

```shell
# Create a new default lambda
fusion new aws lambda
```

## Output

```json
resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "my_lambda" {
  filename         = "foo.zip"
  function_name    = "MyLambda"
  role             = aws_iam_role.iam_for_lambda.arn
  handler          = "handler.index"
  source_code_hash = filebase64sha256("foo.zip")

  runtime = "NODE_14.x"

  environment {
    variables = {}
  }
}
```