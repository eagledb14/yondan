# Yondan
## Background
Yondan is a clone of the internet scanning site [shodan.io](https://www.shodan.io/). The purpose is to host a local version of the site that can be used to training purposes. It can scan local networks given the cidr ranges in resources/ranges.txt and be given custom easter egg ips in resources/flags.txt.

## Installation
To install, make sure you have golang and nmap installed on your machine. Then clone and cd into the repo.

## Usage

build.sh allows for an easy wat to compile and zip all the required files into a tar file. It currently only supports linux, but can be changed for windows by changed the go build target to 
```
GOOS=windows GOARCH=amd64
```

Yondan has 2 args, the first is the port that it serves from, the default is port 8080. 

The second arg is the interval that it scans the network in minutes. More aggressive intervals will create more noise for a SIM tool to find, though has no other purpose if you are not updating the network regularyl. The default value is 5 minutes.

```
//runs default configurations
./yondan

//runs on port 80, with default scan interval
./yondan 80 

//runs on port 80 with 10 minue scan interval
./yondan 80 10
```

Resouces folder is required to be in the same directory as the yondan binary to run.


## Example ranges.txt
```
google.com
facebook.com
amazon.com
127.0.0.0/28
```

## Example flags.txt
```
[
	{
		"ip": "36.45.10.107",
		"hostname": "example-flags.com",
		"ports": [
			{
				"id": 8080,
				"service": "http"
			},
			{
				"id": 92,
				"service": "traceroute"
			}
		]
	},
	{
		"ip": "8.8.8.8",
		"hostname": "google.com",
		"ports": [
			{
				"id": 443,
				"service": "https"
			},
			{
				"id": 8010,
				"service": "fortinet"
			}
		]
	}

```
