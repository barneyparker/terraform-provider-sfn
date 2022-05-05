package sfn

import (
	"hash/crc32"
	"encoding/json"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func StringHashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func ParseStep(d *schema.ResourceData, Type string) map[string]interface{} {
	step := map[string]interface{}{
		"Type": Type,
	}
	
	if name, ok := d.GetOk("name"); ok {
		step["Name"] = name.(string)
	} else {
		step["Name"] = Type
	}

	if comment, ok := d.GetOk("comment"); ok {
		step["Comment"] = comment.(string)
	}

	if next, ok := d.GetOk("next"); ok {
		step["Next"] = next.(string)
	} else {
		step["End"] = true
	}

	return step
}

func ParseParameters(d *schema.ResourceData, step map[string]interface{}) {
	if inputpath, ok := d.GetOk("inputpath"); ok {
		step["InputPath"] = inputpath.(string)
	}

	if parameters, ok := d.GetOk("parameters"); ok {
		step["Parameters"] = parameters.(map[string]interface{})
	} else {
		step["Parameters"] = make(map[string]interface{})
	}

	if result, ok := d.GetOk("result"); ok {
		step["Result"] = result.(map[string]interface{})
	}

	if resultpath, ok := d.GetOk("resultpath"); ok {
		step["ResultPath"] = resultpath.(string)
	}

	if outputpath, ok := d.GetOk("outputpath"); ok {
		step["OutputPath"] = outputpath.(string)
	}
}

func MarshallResource(d *schema.ResourceData, step map[string]interface{}) diag.Diagnostics {
	jsonDoc, err := json.MarshalIndent(step, "", "  ")

	if err != nil {
		return diag.FromErr(err)
	}

	jsonString := string(jsonDoc)

	//return diag.Errorf("Resource: %s", jsonString)

	d.Set("step", jsonString)
	d.SetId(strconv.Itoa(StringHashcode(jsonString)))

	return nil
}