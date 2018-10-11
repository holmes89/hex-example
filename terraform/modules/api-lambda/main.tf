resource "aws_lambda_function" "lambda_func" {
  function_name = "${var.name}-${var.stage}"

  # The bucket name as created earlier with "aws s3api create-bucket"
  s3_bucket = "${var.bucket}"
  s3_key    = "${var.app_version}/main.zip"

  # "main" is the filename within the zip file (main.js) and "handler"
  # is the name of the property under which the handler function was
  # exported in that file.
  handler = "main"
  runtime = "go1.x"

  role = "${aws_iam_role.lambda_exec.arn}"

  environment={
    variables = "${var.env_vars}"
  }
}

# IAM role which dictates what other AWS services the Lambda function
# may access.
resource "aws_iam_role" "lambda_exec" {
  name = "${var.name}-${var.stage}-role"

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

resource "aws_lambda_permission" "allow_api_gateway" {
    function_name = "${aws_lambda_function.lambda_func.function_name}"
    statement_id = "AllowExecutionFromApiGateway"
    action = "lambda:InvokeFunction"
    principal = "apigateway.amazonaws.com"
    source_arn = "${aws_iam_role.lambda_exec.arn}"
}


resource "aws_api_gateway_rest_api" "lambda_func" {
  name        = "${var.name}-${var.stage}-api"
  description = "REST API for ${var.display_name}"
}

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = "${aws_api_gateway_rest_api.lambda_func.id}"
  parent_id   = "${aws_api_gateway_rest_api.lambda_func.root_resource_id}"
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "proxy" {
  rest_api_id   = "${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id   = "${aws_api_gateway_resource.proxy.id}"
  http_method   = "ANY"
  authorization = "NONE"
}


resource "aws_api_gateway_integration" "lambda" {
  rest_api_id = "${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id = "${aws_api_gateway_method.proxy.resource_id}"
  http_method = "${aws_api_gateway_method.proxy.http_method}"

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri = "${aws_lambda_function.lambda_func.invoke_arn}"
}

resource "aws_api_gateway_method" "proxy_root" {
  rest_api_id   = "${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id   = "${aws_api_gateway_rest_api.lambda_func.root_resource_id}"
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda_root" {
  rest_api_id = "${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id = "${aws_api_gateway_method.proxy_root.resource_id}"
  http_method = "${aws_api_gateway_method.proxy_root.http_method}"

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri = "${aws_lambda_function.lambda_func.invoke_arn}"
}

resource "aws_api_gateway_deployment" "lambda_func" {
  depends_on = [
    "aws_api_gateway_integration.lambda",
    "aws_api_gateway_integration.lambda_root",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.lambda_func.id}"
  stage_name  = "${var.stage}"
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.lambda_func.arn}"
  principal     = "apigateway.amazonaws.com"

  # The /*/* portion grants access from any method on any resource
  # within the API Gateway "REST API".
   source_arn = "${aws_api_gateway_deployment.lambda_func.execution_arn}/*/*"
}

resource "aws_iam_policy" "lambda_logging" {
  name = "${var.name}-${var.stage}-logging"
  path = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role = "${aws_iam_role.lambda_exec.name}"
  policy_arn = "${aws_iam_policy.lambda_logging.arn}"
}

# aws_api_gateway_method._
resource "aws_api_gateway_method" "proxy_method" {
  rest_api_id   ="${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id   = "${aws_api_gateway_method.proxy.resource_id}"
  http_method   = "OPTIONS"
  authorization = "NONE"
}
resource "aws_api_gateway_method" "proxy_root_method" {
  rest_api_id   ="${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id   = "${aws_api_gateway_method.proxy_root.resource_id}"
  http_method   = "OPTIONS"
  authorization = "NONE"
}

# aws_api_gateway_integration._
resource "aws_api_gateway_integration" "proxy_options_integration" {
  rest_api_id ="${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id = "${aws_api_gateway_method.proxy.resource_id}"
  http_method = "${aws_api_gateway_method.proxy_method.http_method}"

  type = "MOCK"

  request_templates {
    "application/json" = "{ \"statusCode\": 200 }"
  }
}

# aws_api_gateway_integration._
resource "aws_api_gateway_integration" "proxy_options_root_integration" {
  rest_api_id ="${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id = "${aws_api_gateway_method.proxy_root.resource_id}"
  http_method = "${aws_api_gateway_method.proxy_root_method.http_method}"

  type = "MOCK"

  request_templates {
    "application/json" = "{ \"statusCode\": 200 }"
  }
}


# aws_api_gateway_integration_response._
resource "aws_api_gateway_integration_response" "proxy_options_response" {
  rest_api_id ="${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id = "${aws_api_gateway_method.proxy.resource_id}"
  http_method = "${aws_api_gateway_method.proxy_method.http_method}"
  status_code = 200

  depends_on = [
    "aws_api_gateway_integration.proxy_options_integration",
  ]
}

# aws_api_gateway_integration_response._
resource "aws_api_gateway_integration_response" "proxy_options_root_response" {
  rest_api_id ="${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id = "${aws_api_gateway_method.proxy_root.resource_id}"
  http_method = "${aws_api_gateway_method.proxy_root_method.http_method}"
  status_code = 200

  depends_on = [
    "aws_api_gateway_integration.proxy_options_integration",
  ]
}


# aws_api_gateway_method_response._
resource "aws_api_gateway_method_response" "proxy_options" {
  rest_api_id ="${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id = "${aws_api_gateway_method.proxy.resource_id}"
  http_method = "${aws_api_gateway_method.proxy_method.http_method}"
  status_code = 200

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
    "method.response.header.Access-Control-Max-Age"       = true
  }

  response_models = {
    "application/json" = "Empty"
  }

  depends_on = [
    "aws_api_gateway_method.proxy_method",
  ]
}

resource "aws_api_gateway_method_response" "proxy_root_options" {
  rest_api_id ="${aws_api_gateway_rest_api.lambda_func.id}"
  resource_id = "${aws_api_gateway_method.proxy_root.resource_id}"
  http_method = "${aws_api_gateway_method.proxy_root_method.http_method}"
  status_code = 200

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
    "method.response.header.Access-Control-Max-Age"       = true
  }

  response_models = {
    "application/json" = "Empty"
  }

  depends_on = [
    "aws_api_gateway_method.proxy_root_method",
  ]
}
