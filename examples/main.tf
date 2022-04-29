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

output "workflow" {
  value = data.sfn_workflow.example.json
}
