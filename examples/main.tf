terraform {
  required_providers {
    stepfunctions = {
      version = "~> 0.0.1"
      source  = "barneyparker/aws/stepfunctions"
    }
  }
}

data "stepfunctions_workflow" "example" {
  comment    = "Example Workflow x"
  start_step = data.stepfunctions_pass.pass.name
  steps = [
    data.stepfunctions_pass.pass.step,
    #data.stepfunctions_wait.wait.step,
    #data.stepfunctions_success.success.step,
    #data.stepfunctions_fail.fail.step,
    data.stepfunctions_task.list_roles.step,
  ]
}

data "stepfunctions_success" "success" {
  comment = "Success!"
}

data "stepfunctions_fail" "fail" {
  comment = "Failed"
  error   = "this was the error"
  cause   = "shit happens :("
}

data "stepfunctions_pass" "pass" {
  name    = "myPass"
  comment = "Example Pass"
  /*inputpath = "$.inputpath"
  parameters = {
    "parameter" = "$.parameter"
  }

  result = {
    "result" = "$.result"
  }
  resultpath = "$.resultpath"
  outputpath = "$.outputpath"*/
  next = data.stepfunctions_task.list_roles.name
}

data "stepfunctions_wait" "wait" {
  name       = "myWait"
  comment    = "a comment...."
  seconds    = 5
  inputpath  = "$.inputpath"
  outputpath = "$.outputpath"
  next       = data.stepfunctions_task.list_roles.name
}

data "stepfunctions_task" "list_roles" {
  name       = "list_roles"
  comment    = "List IAM Roles"
  resource   = "arn:aws:states:::aws-sdk:iam:listRoles"
  resultpath = "$"
}

output "workflow" {
  value = data.stepfunctions_workflow.example.json
}
