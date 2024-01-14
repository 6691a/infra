resource "azurerm_dns_zone" "playhub" {
    name               = "playhub.kra"
    resource_group_name = var.resource_group_name
}

resource "azurerm_dns_mx_record" "dooray" {
    name                = "@"
    zone_name           = azurerm_dns_zone.playhub.name
    resource_group_name = var.resource_group_name
    ttl                 = 600

    record {
        exchange = "aspmx1.dooray.com."
        preference = 1
    }

    record {
        exchange = var.custom_domain_mx_record
        preference = 32767
    }
}

resource "azurerm_dns_txt_record" "dooray" {
    name                = "@"
    zone_name           = azurerm_dns_zone.playhub.name
    resource_group_name = var.resource_group_name
    ttl                 = 600
    record {
        value = "v=spf1 include:_spf.dooray.com ~all"
    }
}
