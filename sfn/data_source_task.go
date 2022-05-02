package sfn

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func dataSourceTask() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRead,
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
			"resource": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringLenBetween(0, 64),
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
			"parameters": {
				Type: schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},
			"result": {
				Type: schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},
			"resultpath": {
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

func dataSourceTaskRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	step := ParseStep(d, "Task")
	ParseParameters(d, step)

	if resource, ok := d.GetOk("resource"); ok {
		step["Resource"] = resource.(string)
	}

	return MarshallResource(d, step)
}