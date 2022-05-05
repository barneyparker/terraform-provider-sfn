resource "aws_sfn_state_machine" "sfn_state_machine" {
  name     = "my-state-machine"
  role_arn = aws_iam_role.iam_for_sfn.arn
  #type = "EXPRESS"

  definition = data.sfn_workflow.example.json
}

resource "aws_iam_role" "iam_for_sfn" {
  name               = "iam_for_sfn"
  assume_role_policy = data.aws_iam_policy_document.iam_for_sfn_assume.json
}

data "aws_iam_policy_document" "iam_for_sfn_assume" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["states.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy" "iam_for_sfn" {
  name   = "permissions"
  role   = aws_iam_role.iam_for_sfn.id
  policy = data.aws_iam_policy_document.iam_for_sfn.json
}

data "aws_iam_policy_document" "iam_for_sfn" {
  statement {
    actions   = ["iam:ListRoles"]
    resources = ["*"]
  }
}