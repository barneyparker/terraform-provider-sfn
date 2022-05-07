package stepfunctions

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func dataSourceWait() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWaitRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringLenBetween(0, 64),
			},
			"comment": {
				Type: schema.TypeString,
				Optional: true,
				Default: "Pass Step",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"seconds": {
				Type: schema.TypeInt,
				Required: true,
			},
			"next": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringLenBetween(0, 64),
			},
			"inputpath": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"outputpath": {
				Type: schema.TypeString,
				Optional: true,
				Default: "",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"step": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceWaitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	step := ParseStep(d, "Wait")
	ParseParameters(d, step)

	if seconds, ok := d.GetOk("seconds"); ok {
		step["Seconds"] = seconds.(int)
	}

	return MarshallResource(d, step)
}