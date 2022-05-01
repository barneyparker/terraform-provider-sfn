package sfn

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"sfn_workflow": dataSourceWorkflow(),
			"sfn_success": dataSourceSucceed(),
			"sfn_fail": dataSourceFail(),
			"sfn_pass": dataSourcePass(),
			"sfn_wait": dataSourceWait(),
		},
	}
}
