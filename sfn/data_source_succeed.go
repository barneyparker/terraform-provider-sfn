package sfn

import (
	"context"
	"encoding/json"
	"strconv"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type SFNSucceed struct {
	Type       string                 `json:"Type"`
	Name       string                 `json:"Name"`
	Comment    string                 `json:"Comment,omitempty"`
}

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
	succeedStep := &SFNSucceed{}

	succeedStep.Type = "Succeed"
	
	if comment, ok := d.GetOk("comment"); ok {
		succeedStep.Comment = comment.(string)
	}
	
	if name, ok := d.GetOk("name"); ok {
		succeedStep.Name = name.(string)
	} else {
		succeedStep.Name = "Success"
	}

	jsonDoc, err := json.MarshalIndent(succeedStep, "", "  ")

	if err != nil {
		return diag.FromErr(err)
	}

	jsonString := string(jsonDoc)

	d.Set("step", jsonString)
	d.SetId(strconv.Itoa(StringHashcode(jsonString)))

	return nil
}