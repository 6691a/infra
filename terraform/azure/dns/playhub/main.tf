resource "azurerm_dns_zone" "playhub" {
    name               = "playhub.kra"
    resource_group_name = local.resource_group_name
}

resource "azurerm_dns_mx_record" "mx" {
    name                = "@"
    zone_name           = azurerm_dns_zone.playhub.name
    resource_group_name = local.resource_group_name
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

resource "azurerm_dns_txt_record" "txt" {
    name                = "@"
    zone_name           = azurerm_dns_zone.playhub.name
    resource_group_name = local.resource_group_name
    ttl                 = 600
    record {
        value = "v=spf1 include:_spf.dooray.com ~all"
    }

    record {
        value = "v=spf1 include:spf.protection.outlook.com -all"
    }

    record {
        value = var.smtp_recode
    }
}

resource "azurerm_dns_cname_record" "dkim1" {
    name                = "selector1-azurecomm-prod-net._domainkey"
    zone_name           = azurerm_dns_zone.playhub.name
    resource_group_name = local.resource_group_name
    ttl                 = 600
    record {
        cname = "selector1-azurecomm-prod-net._domainkey.azurecomm.net"
    }
}

resource "azurerm_dns_cname_record" "dkim2" {
    name                = "selector2-azurecomm-prod-net._domainkey"
    zone_name           = azurerm_dns_zone.playhub.name
    resource_group_name = local.resource_group_name
    ttl                 = 600
    record {
        cname = "selector2-azurecomm-prod-net._domainkey.azurecomm.net"
    }
}