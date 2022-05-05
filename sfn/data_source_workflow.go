package sfn

import (
	"context"
	"encoding/json"
	"strconv"
	"fmt"
	"reflect"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

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
	if start_step, ok := d.GetOk("start_step"); ok {
		workFlow := map[string]interface{}{
			"StartAt": start_step.(string),
		}

		states := make(map[string]interface{})
		
		if comment, ok := d.GetOk("comment"); ok {
			workFlow["Comment"] = comment.(string)
		}
	
		if steps, ok := d.GetOk("steps"); ok && len(steps.([]interface{})) > 0 {
			fmt.Println("Steps: ", reflect.TypeOf(steps), steps)
			for _, docStep := range steps.([]interface{}) {
				var myStep map[string]interface{}
				if err := json.Unmarshal([]byte(docStep.(string)), &myStep); err != nil {
					return diag.FromErr(err)
				}
				
				Name := myStep["Name"].(string)
				delete(myStep, "Name")
				states[Name] = myStep
			}
		}
		
		workFlow["States"] = states
		
		jsonDoc, err := json.MarshalIndent(workFlow, "", "  ")

		if err != nil {
			return diag.FromErr(err)
		}
	
		jsonString := string(jsonDoc)
	
		d.Set("json", jsonString)
		d.SetId(strconv.Itoa(StringHashcode(jsonString)))
	
		return nil
	} else {
		// TODO: Should probably do something here...
		return nil	
	}
}
