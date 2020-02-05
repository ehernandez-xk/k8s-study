
`nslookup`

NsLookup queries the specified DNS server and retrieves the requested records that are associated with the domain name you provided. These records contain information like the domain name’s IP addresses.

The following types of DNS records are especially useful:

**A:** the IPv4 address of the domain.
**AAAA:** the domain’s IPv6 address.
**CNAME:** the canonical name — allowing one domain name to map on to another. This allows more than one website to refer to a single web server.
**MX:** the server that handles email for the domain.
**NS:** one or more authoritative name server records for the domain.
**TXT:** a record containing information for use outside the DNS server. The content takes the form name=value. This information is used for many things including authentication schemes such as SPF and DKIM.
