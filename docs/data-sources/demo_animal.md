---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "demo_animal Data Source - demo-private-provider-csa-provider"
subcategory: ""
description: |-
  Animal Example Data Source.
---

# demo_animal (Data Source)

Animal Example Data Source.

## Example Usage

```terraform
data "demo_animal" "example" {
  animal_id = "78c3a1ba-8a77-47f9-846c-0e38ee902127"
  class     = "Bird"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `animal_id` (String) Id of animal.
- `class` (String) Class of animal.

### Read-Only

- `animal` (String) Example animal of the specified class.
- `id` (String) The ID of this resource.


