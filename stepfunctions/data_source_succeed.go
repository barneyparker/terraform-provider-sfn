package stepfunctions

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func dataSourceSucceed() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSucceedRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringLenBetween(0, 64),
			},"comment": {
				Type: schema.TypeString,
				Optional: true,
				Default: "Pass Step",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"step": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSucceedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	step := ParseStep(d, "Succeed")
	return MarshallResource(d, step)
}