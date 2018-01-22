# NMAP bulk

Runs a bulk of nmap commands based on JSON input with hosts.

### Usage:

`./bulk_nmap -config=./config.json`

where the contents of `config.json` are:

```
{
	"input": "./input.json",
	"output": "./output.txt",
	"command": "nmap"
}
```

and the `input.json`

```
{
	"gpestana.com": "127.0.0.1",
	"google.com": "198.168.22.10"
}
```

### Output:

If successful, the output will be written to `output.json`:

```
[
  {
    "Host": "gpestana.com",
    "IP": "127.0.0.1",
    "Output": "\nStarting Nmap 7.60 ( https://nmap.org ) at 2018-01-22 21:54 EET\nNmap scan report for gpestana.com (104.27.161.169)\nHost is up (0.020s latency).\nOther addresses for gpestana.com (not scanned): 104.27.160.169 2400:cb00:2048:1::681b:a1a9 2400:cb00:2048:1::681b:a0a9\nNot shown: 996 filtered ports\nPORT     STATE SERVICE\n80/tcp   open  http\n443/tcp  open  https\n8080/tcp open  http-proxy\n8443/tcp open  https-alt\n\nNmap done: 1 IP address (1 host up) scanned in 4.41 seconds\n"
  },
  {
    "Host": "google.com",
    "IP": "198.168.22.10",
    "Output": "\nStarting Nmap 7.60 ( https://nmap.org ) at 2018-01-22 21:54 EET\nNmap scan report for google.com (172.217.20.46)\nHost is up (0.033s latency).\nOther addresses for google.com (not scanned): 2a00:1450:400f:80b::200e\nrDNS record for 172.217.20.46: arn11s01-in-f14.1e100.net\nNot shown: 998 filtered ports\nPORT    STATE SERVICE\n80/tcp  open  http\n443/tcp open  https\n\nNmap done: 1 IP address (1 host up) scanned in 5.20 seconds\n"
  }
]
```
