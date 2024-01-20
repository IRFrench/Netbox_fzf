# Netbox_fzf
This is a tool which creates a self updating ssh config file with data from Netbox. This was designed to be used with fzf to find ssh hostnames easily


## How to use:

The tool currently only supports Device and Virual Machine API types. The idea for this being that servers and virtual machines will be the main things you would need to ssh onto

Since I wanted the tool to be used by anyone for any number of things, I wanted to make it so you could add any netbox api requests with config.
For this, simply add any netbox API urls (like `https://<netbox_url>/api/dcim/devices/?tenant_id=2`) to a file called `netbox.txt` in the netbox directory. These will need to return a list of results in the API.

Then add your template to the 