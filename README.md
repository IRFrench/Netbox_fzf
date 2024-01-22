# Netbox_fzf
This is a tool which creates a self updating ssh config file with data from Netbox. This was designed to be used with fzf to find ssh hostnames easily. This was an awesome idea that didn't come from me, I just wanted to make it open source. You can thanks Laurence for that ;)

## Why does this exist

When working with a lot of different servers or virtual machines, you either have to remember host names or ip addresses in order to ssh onto them. Most people don't do this, and often a database is used to store and record these devices, like Netbox

This tool gets these devices from Netbox and adds them to an ssh configuration file. This can then be used by fzf to search and find any host with ease

### fzf

People much smarter than me made a [fuzzy finder for terminals](https://github.com/junegunn/fzf). This tool is very smart and removes a lot of hassle when trying to write commands. One thing it will do, it look at the ssh config file when trying to ssh onto a device.

This means, you can run `ssh **` and be given a list of every host you have in your known hosts or config. Therefore, if you fill your config file of all the hosts you will ssh onto, it will be findable with fzf!

## How to use

This project has been setup with github workflows to create binaries on tags. So simply grab the binary for your system and rename it to whatver you would like it to be, like `netbox_fzf` for example.

Mrk.2!

The orginal tool used environment variables for setup. While easy to use, this has a limiting factor of 'I don't like that so much', so it has changed.

The tool now uses flags!

- `-t <netbox_api_token>` provides the token to the service. Since this is a manditory field, you can also provide this via the environment variable `TOKEN`.
- `-n <netbox_url_file_path>` provides the path to the netbox url file.
- `-s <ssh_config_template_path>` provides the path to the ssh config template.
- `-o <output_file_path>` provides the path for the created config file. I figured this would be easier than requiring any jobs to also move the file.
- `-v` enables verbose mode, which adds some extra debug log messages
- `-h` provides a help menu with __most__ of this information on it!

Want more flags? Me too! They were fun to make!

### Binaries

In the most recent release there should be multiple files with different OS and architectures as their file names.

Simply, download the file for your OS and CPU architecture:

```
amd64 (x86) = intel, amd
arm64 = apple silicon

Debian = MacOS
Linux = Linux (Ubuntu)
Windows = Windows
```

### Setup Netbox urls

This is just adding Netbox API urls to a text file. (Almost?) Every page on the Netbox UI is APIable by just adding a prefix of `api/`. This means you can perform a search for any devices or VMS you would like to have in your ssh config and add that URL to the txt file with the api prefix.

For an example of a Netbox urls file, please look at the `example_netbox.txt` file in the examples directory.

### Setup the ssh template

This uses golangs templating, with documentation available [here](https://pkg.go.dev/text/template).

The following structs are passed across to the templating service:


```
type templateData struct {
	NetboxConfigs []NetboxConfigLists
}

type NetboxConfigLists struct {
	Url    string
	Config []SshDeviceSettings
}

type SshDeviceSettings struct {
	Name string
	Ip   string
}
```

This means you can loop through all of the netbox configs using `{.NetboxConfigs}`.

For an example of a template, please look at the `example_ssh_template` file in the examples directory.

## Improvements and bugs

I think I like this tool, so I'll keep up to date with changes and issues raised. If you want something changed or added, just raise an issue or pull request.