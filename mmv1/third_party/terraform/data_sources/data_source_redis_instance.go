package google

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func DataSourceGoogleRedisInstance() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(ResourceRedisInstance().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project", "region")

	return &schema.Resource{
		Read:   dataSourceGoogleRedisInstanceRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleRedisInstanceRead(d *schema.ResourceData, meta interface{}) error {
	id, err := ReplaceVars(d, meta.(*Config), "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return err
	}
	d.SetId(id)

	return resourceRedisInstanceRead(d, meta)
}
