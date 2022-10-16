resource "aws_lambda_function" "lambda-emailService" {
  function_name = var.functionName
  role          = aws_iam_role.iam_lambda.arn
  s3_bucket     = var.artifactBucket
  s3_key        = var.artifactName
  memory_size   = 128
  handler       = var.handler
  runtime       = var.runtime
  timeout       = 30


  environment {
    variables = local.env
  }

  ephemeral_storage {
    size = 512
  }

  depends_on = [
    aws_cloudwatch_log_group.lambda,
    aws_iam_policy_attachment.lambda_logging
  ]

  tags = {
    Env   = "Production"
    Name  = "Email Service"
    About = "Send Emails"
  }
}

resource "aws_lambda_alias" "production" {
  function_name    = aws_lambda_function.lambda-emailService.function_name
  function_version = "1"
  name             = "production"
}

resource "aws_lambda_alias" "dev" {
  function_name    = aws_lambda_function.lambda-emailService.function_name
  function_version = "1"
  name             = "dev"
}

locals {
  env = {
    SMTP       = var.smtp_host
    PORT       = var.smtp_port
    AWS_REGION = var.aws_region
  }
}

