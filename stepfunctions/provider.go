package stepfunctions

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"stepfunctions_workflow": dataSourceWorkflow(),
			"stepfunctions_success": dataSourceSucceed(),
			"stepfunctions_fail": dataSourceFail(),
			"stepfunctions_pass": dataSourcePass(),
			"stepfunctions_wait": dataSourceWait(),
			"stepfunctions_task": dataSourceTask(),
		},
	}
}
