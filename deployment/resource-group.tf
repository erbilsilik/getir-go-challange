# Creates a resource group for getirgochallange in Azure account.

resource "azurerm_resource_group" "getirgochallange" {
  name     = var.app_name
  location = var.location
}
