package stepfunctions

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func dataSourceFail() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFailRead,
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
			"error": {
				Type: schema.TypeString,
				Optional: true,
				Default: "Pass Step",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"cause": {
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

func dataSourceFailRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	step := ParseStep(d, "Fail")

	if error, ok := d.GetOk("error"); ok {
		step["Error"] = error.(string)
	}

	if cause, ok := d.GetOk("cause"); ok {
		step["Cause"] = cause.(string)
	}

	return MarshallResource(d, step)
}