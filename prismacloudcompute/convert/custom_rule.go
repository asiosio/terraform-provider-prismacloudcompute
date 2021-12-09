package convert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/paloaltonetworks/prisma-cloud-compute-go/rule"
)

// Converts a custom rule schema to a custom rule object for SDK compatibility.
func SchemaToCustomRule(d *schema.ResourceData) rule.CustomRule {
	parsedRule := rule.CustomRule{}

	if val, ok := d.GetOk("description"); ok {
		parsedRule.Description = val.(string)
	}
	if val, ok := d.GetOk("message"); ok {
		parsedRule.Message = val.(string)
	}
	if val, ok := d.GetOk("name"); ok {
		parsedRule.Name = val.(string)
	}
	if val, ok := d.GetOk("script"); ok {
		parsedRule.Script = val.(string)
	}
	if val, ok := d.GetOk("type"); ok {
		parsedRule.Type = val.(string)
	}
	return parsedRule
}
