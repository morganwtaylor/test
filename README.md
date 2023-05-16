
# Swift Control Testing Automation

This program will enable automated, repeated control testing for SWIFT. The script includes all control items that are able to be automated.


## Prerequisites

#### All Platforms
- Go installed on the machine running the script.
- A list of hostnames that you want to test against in a file names "hostnames.txt".
- SSH credentials with an environment variables named "SSH_PASSWORD"

#### Windows
- The PuTTY SSH client, Plink, installed and added to your system's PATH (https://www.putty.org)

#### Linux

- The Expect scripting language installed ("sudo apt-get install expect -y" for Ubunty/Debian or "sudo yum install expect -y" for CentOS/RHEL).



## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`SSH_PASSWORD`

`SSH_USERNAME`

