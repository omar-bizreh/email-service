resource "aws_cloudwatch_log_group" "lambda" {
  name              = "/aws/lambda/${var.functionName}"
  retention_in_days = 3
}
