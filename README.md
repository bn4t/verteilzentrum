# Verteilzentrum

*A minimalistic mailing list*

Verteilzentrum is a minimalistic mailing list following the [KISS](https://en.wikipedia.org/wiki/KISS_principle) philosophy. 

**Features:**
- Single config file
- Multiple lists
- Blacklisting
- Whitelisting
- Configurable publishing rights

# Config
## General options

##### hostname
The hostname of the list server. If you specify a TLS certificate it has to be valid for this hostname.

##### read_timeout
SMTP read timeout in milliseconds.

##### write_timeout
SMTP write timeout in milliseconds.

##### max_message_bytes
Maximum incoming message size in bytes.

##### tls_cert_file
Path to the TLS certificate file.

##### tls_key_file
Path to the corresponding private key to the TLS certificate. 

To disable inbound TLS just comment out both TLS settings.

#### Example
````toml
[verteilzentrum]
hostname = "lists.example.com"
read_timeout = 100000
write_timeout = 100000
max_message_bytes = 1048576 # 1024 * 1024
tls_cert_file = "/some/path/cert.pem"
tls_key_file = "/some/path/key.pem"
````

## Lists
Lists are represented toml tables in an array.

#### Table elements
##### name 
The name of the list which also serves as the list address.
##### whitelist
Array of whitelisted email addresses which are allowed to subscribe to the list. Supports wildcards.

If empty the whitelist is disabled.
##### blacklist
Array of blacklisted email addresses. 

Blacklisted addresses are not allowed to interact in any way with the list. Can be empty. Supports wildcards. 

Important: The Blacklist has a higher priority than the whitelist.
##### can_publish
Array of email addresses which are allowed to publish messages to the list. Supports wildcards.

#### Example
````toml
[[list]]
name = "news@lists.example.com"
whitelist = ["*"]
blacklist = ["baduser@gmail.com"]
can_publish = ["admin@example.com"]

[[list]]
name = "private-list@lists.example.com"
whitelist = ["postmaster@example.com","admin@example.com"]
blacklist = []
can_publish = ["admin@example.com", "postmaster@example.com"]
````

## Command line flags
- `-config <config file>` - The location of the config file to use. Defaults to `config.toml` in the working directory.
- `-datadir <data directory>` - The location where all persistent data is stored. Defaults to the working directory.
