package sfn

import (
	"context"
	"encoding/json"
	"strconv"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type SFNWorkflow struct {
	Comment    string        `json:"Comment,omitempty"`
	StartsAt   string        `json:"StartsAt"`
	Steps      []interface{} `json:"States"`
}

func dataSourceWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkflowRead,
		Schema: map[string]*schema.Schema{
			"comment": {
				Type: schema.TypeString,
				Optional: true,
				Default: "Step-Function Workflow",
				ValidateFunc: validation.StringLenBetween(0, 512),
			},
			"start_step": {
				Type: schema.TypeString,
				Required: true,
			},
			"steps": {
				Type: schema.TypeList,
				Required: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceWorkflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	workFlow := &SFNWorkflow{}

	if comment, ok := d.GetOk("comment"); ok {
		workFlow.Comment = comment.(string)
	}

	if start_step, ok := d.GetOk("start_step"); ok {
		workFlow.StartsAt = start_step.(string)
	}

	if steps, ok := d.GetOk("steps"); ok {
		workFlow.Steps = steps.([]interface{})
	}

	jsonDoc, err := json.MarshalIndent(workFlow, "", "  ")

	if err != nil {
		// should never happen if the above code is correct
		return diag.FromErr(err)
	}

	jsonString := string(jsonDoc)

	d.Set("json", jsonString)
	d.SetId(strconv.Itoa(StringHashcode(jsonString)))

	return nil
}
