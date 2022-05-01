terraform {
  required_providers {
    sfn = {
      version = "~> 0.0.1"
      source  = "barneyparker/aws/sfn"
    }
  }
}

data "sfn_workflow" "example" {
  comment    = "Example Workflow"
  start_step = data.sfn_pass.pass.name
  steps = [
    data.sfn_pass.pass.step,
  ]
}

data "sfn_pass" "pass" {
  name      = "myPass"
  comment   = "Example Pass"
  inputpath = "$.inputpath"
  parameters = {
    "parameter" = "$.parameter"
  }

  result = {
    "result" = "$.result"
  }
  resultpath = "$.resultpath"
  outputpath = "$.outputpath"
  next      = "end"
}

output "workflow" {
  value = data.sfn_workflow.example.json
}
