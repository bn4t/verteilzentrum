# Verteilzentrum

*A minimalistic mailing list*

Verteilzentrum is a minimalistic mailing list following the [KISS](https://en.wikipedia.org/wiki/KISS_principle) philosophy. 

**Features:**
- Single config file
- Multiple lists
- Blacklisting
- Whitelisting
- Configurable publishing rights

# Installation
0. Install golang (>=1.14), gcc, make and build-essential if you don't have them already
1. Clone the repository: `git clone https://github.com/bn4t/verteilzentrum.git`
2. Checkout the latest stable tag 
3. Make sure `go` is in your `$PATH` and run `make build` to build the verteilzentrum binary
4. Run `sudo make install` to install verteilzentrum on your system. This will create the directories `/etc/verteilzentrum` (config directory) and `/var/lib/verteilzentrum` (data directory). Additionally the user `verteilzentrum` will be created.
5. If you have systemd installed you can run `sudo make install-systemd` to install the systemd service. Run `service verteilzentrum start` to start the verteilzentrum service. Verteilzentrum will automatically run as the `verteilzentrum` user.

You can make certificates and private key files accessible to the `verteilzentrum` user with the following command:
````shell script
setfacl -m u:verteilzentrum:rx /etc/letsencrypt/ 
````

#### Increasing deliverability
To increase deliverability it is recommended to set up an [SPF](https://en.wikipedia.org/wiki/Sender_Policy_Framework) and [DMARC](https://en.wikipedia.org/wiki/DMARC) record. 


# How to use

#### Subscribing to a list
Send an email (content doesn't matter) to `subscribe+$list_name`. E.g. `subscribe+news@lists.example.com`.

You will receive a confirmation email that subscribing was successful.

#### Unsubscribing from a list
Send an email (content doesn't matter) to `unsubscribe+$list_name`. E.g. `unsubscribe+news@lists.example.com`.

You will receive a confirmation email that unsubscribing was successful.

# Config

By default the config is located at `/etc/verteilzentrum/config.toml`.

## General options

##### bind_to
The address including port on which the server should listen for non tls connections.
Listens by default on `0.0.0.0:25`.

##### bind_to_tls
The address including port on which the server should listen for tls connections.
Can be left empty if no tls certificates are configured.

Listens by default on `0.0.0.0:465`.

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

##### data_dir
The location where all persistent data is stored.
 
#### mta_address
The address of the mta used to send mailing list messages.

This mta is used to send messages for all configured mailing lists.

#### mta_auth_method
The auth method used for authentication to the mta.

Can be either `PLAIN` or `ANONYMOUS`.

#### mta_username
The username used for authentication to the mta.

#### mta_password
The password used for authentication to the mta.


#### Example
````toml
[verteilzentrum]
bind_to = "0.0.0.0:25"
bind_to_tls = "0.0.0.0:465"
hostname = "lists.example.com"
read_timeout = 100000
write_timeout = 100000
max_message_bytes = 1048576 # 1024 * 1024
tls_cert_file = "/some/path/cert.pem"
tls_key_file = "/some/path/key.pem"
data_dir = "/var/lib/verteilzentrum"
mta_address = "smtp.example.com"
mta_auth_method = "PLAIN"
mta_username = "lists@example.com"
mta_password = "secret"
````

## Lists
Lists are represented as toml tables in an array.

#### Table elements
##### name 
The name of the list which also serves as the list address.
##### whitelist
Array of whitelisted email addresses which are allowed to interact with the list. Supports wildcards.

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


For further examples take a look at the [example config](configs/config.example.toml).

## Command line flags
- `-config <config file>` - The location of the config file to use. Defaults to `config.toml` in the working directory.

# Deinstallation
Run `sudo make uninstall` to uninstall verteilzentrum. 
This will remove the verteilzentrum binary and the directories `/etc/verteilzentrum` and `/var/lib/verteilzentrum` if they are empty.

To remove the systemd service run `sudo make uninstall-systemd`.

# Contributing
Feel free to send patches to me@bn4t.me or to open pull requests on Github. 

# License
This project is licensed under the GPL version 3 and later. See the [LICENSE](LICENSE) file. 
