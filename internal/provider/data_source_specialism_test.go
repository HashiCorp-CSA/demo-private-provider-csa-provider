package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAnimal(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAnimal,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.csa_animal.foo", "class", regexp.MustCompile("^Bird")),
					resource.TestMatchResourceAttr("data.csa_animal.foo", "animal", regexp.MustCompile("^Peregrine Falcon")),
				),
			},
		},
	})
}

const testAccDataSourceAnimal = `
data "csa_animal" "foo" {
  class = "Bird"
}
`
