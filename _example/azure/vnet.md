# Example - Create a vnet

> Create a new azure virtual network


## Command

```shell
# Create a new default vnet
fusion new azure vnet
```

## Output

```json
resource "azurerm_resource_group" "this" {
  name     = "resource_group_name"
  location = "centralus"
}

resource "azurerm_virtual_network" "this" {
  name                = "vnet_name"
  resource_group_name = azurerm_resource_group.this.name
  location            = azurerm_resource_group.this.location
  address_space       = ["0.0.0.0/0"]
}

resource "azurerm_subnet" "this" {
  for_each             = { for subnet in var.virtual_network_settings.subnets : subnet.name => subnet }
  name                 = each.value["name"]
  resource_group_name  = azurerm_resource_group.this.name
  virtual_network_name = azurerm_virtual_network.this.name
  address_prefixes     = [each.value["cidr_block"]]
  dynamic "delegation" {
    for_each = each.value.delegation_name != null ? [1] : []
    content {
      name = each.value.delegation_name
      service_delegation {
        name = each.value.service_delegation_name
      }
    }
  }
}

## The following variable should be separated out into a different file (ex. variables.tf)
variable "virtual_network_settings" {
  description = "An object map of lists that contains the network CIDR block, and subnets"
  type = object({
    address_space = list(string)
    subnets = list(
      object({
        name                    = string
        cidr_block              = string
        security_group          = string
        delegation_name         = string
        service_delegation_name = string
      })
    )
  })
}
```
