# govpn-connect

*govpn-connect* is a simple command-line tool designed to streamline the process of starting an OpenVPN session using the openvpn3 command. By leveraging environment variables for username and password storage, this project eliminates the need for users to repeatedly input their credentials, enhancing both convenience and security.

## Usage

```sh
# After download your ovpn config, import the config
openvpn3 config-import --config myvpn.ovpn --name myvpn --persistent

# (optional) config to allow compression
openvpn3 config-manage --config myvpn --allow-compression yes

# Set you auth name and pass. (You can set this env variables on ~/.bashrc or ~/.zshrc)
export OPENVPN_USERNAME=myuser
export OPENVPN_PASSWORD=mypassword

# connect using the config imported
govpn-connect --config myvpn

# check the session connected
openvpn3 sessions-list
```

## Installation

Requirements:
- [jq](https://stedolan.github.io/jq/download/)
- [curl](https://curl.se/download.html)

```sh
curl -fsSL https://raw.githubusercontent.com/leonardobiffi/govpn-connect/main/scripts/install.sh | sh
```

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License
This project is licensed under the MIT License.
