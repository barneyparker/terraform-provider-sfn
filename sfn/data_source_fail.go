package sfn

import (
	"context"
	"encoding/json"
	"strconv"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type SFNFail struct {
	Type       string                 `json:"Type"`
	Name       string                 `json:"Name"`
	Comment    string                 `json:"Comment,omitempty"`
	Error      string                 `json:"Error, ommitempty"`
	Cause      string                 `json:"Cause,omitempty"`
}

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
	failStep := &SFNFail{}

	failStep.Type = "Fail"

	if comment, ok := d.GetOk("comment"); ok {
		failStep.Comment = comment.(string)
	}

	if name, ok := d.GetOk("name"); ok {
		failStep.Name = name.(string)
	} else {
		failStep.Name = "Fail"
	}

	if error, ok := d.GetOk("error"); ok {
		failStep.Error = error.(string)
	}

	if cause, ok := d.GetOk("cause"); ok {
		failStep.Cause = cause.(string)
	}

	jsonDoc, err := json.MarshalIndent(failStep, "", "  ")

	if err != nil {
		return diag.FromErr(err)
	}

	jsonString := string(jsonDoc)

	d.Set("step", jsonString)
	d.SetId(strconv.Itoa(StringHashcode(jsonString)))

	return nil
}