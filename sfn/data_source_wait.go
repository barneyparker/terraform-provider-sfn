package sfn

import (
	"context"
	"encoding/json"
	"strconv"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type SFNWait struct {
	Type       string                 `json:"Type"`
	Name       string                 `json:"Name"`
	Comment    string                 `json:"Comment,omitempty"`
	Seconds    int                    `json:"Seconds"`
	Next       string                 `json:"Next"`
	InputPath  string                 `json:"InputPath,omitempty"`
	OutputPath string                 `json:"OutputPath,omitempty"`
}

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
	waitStep := &SFNWait{}

	waitStep.Type = "Wait"

	if comment, ok := d.GetOk("comment"); ok {
		waitStep.Comment = comment.(string)
	}

	if name, ok := d.GetOk("name"); ok {
		waitStep.Name = name.(string)
	}

	if seconds, ok := d.GetOk("seconds"); ok {
		waitStep.Seconds = seconds.(int)
	}

	if next, ok := d.GetOk("next"); ok {
		waitStep.Next = next.(string)
	}

	if inputpath, ok := d.GetOk("inputpath"); ok {
		waitStep.InputPath = inputpath.(string)
	}

	if outputpath, ok := d.GetOk("outputpath"); ok {
		waitStep.OutputPath = outputpath.(string)
	}

	jsonDoc, err := json.MarshalIndent(waitStep, "", "  ")

	if err != nil {
		return diag.FromErr(err)
	}

	jsonString := string(jsonDoc)

	d.Set("step", jsonString)
	d.SetId(strconv.Itoa(StringHashcode(jsonString)))

	return nil
}