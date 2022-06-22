package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAnimal() *schema.Resource {
	return &schema.Resource{
		Description: "Animal Example Data Source.",

		ReadContext: dataSourceAnimalsRead,

		Schema: map[string]*schema.Schema{
			"class": {
				Description: "Class of animal.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"animal": {
				Description: "Example animal of the specified class.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceAnimalsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	architect := d.Get("class").(string)
	client := animal_client{id: architect}

	d.SetId(architect)

	d.Set("animal", client.GetAnimalFromClass())

	return diags
}