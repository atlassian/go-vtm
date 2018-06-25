# go-vtm

[![GoDoc](https://godoc.org/github.com/atlassian/go-stingray?status.svg)](https://godoc.org/github.com/whitepages/go-stingray)

go-vtm is a client library for accessing the Brocade Virtual Traffic Manager 
(formerly Riverbed Stingray, and Zeus ZXTM) REST API.

For documentation of the vTM API, see the "Brocade Virtual Traffic Manager: REST API Guide" version 17.2, May 2017 from the
[Brocade support site](http://community.brocade.com/t5/vADC-Docs/Brocade-Virtual-Application-Delivery-Controller-vADC-Product/ta-p/73714).

## Usage

```go
import "github.com/atlassian/go-vtm"
```

Create a new Stingray client.

```go
client := stingray.NewClient(nil, "https://localhost:9070", "username", "password")
```

Manage an extra file.

```go
fmt.Println("Writing...")
r := stingray.NewExtraFile("name")
r.Content = []byte("Test")
resp, err := client.Set(r)
fmt.Println("Status:", resp.Status)
if err != nil {
	log.Fatal(err)
}

fmt.Println("Reading...")
r, resp, err = client.GetExtraFile("name")
if err != nil {
	log.Fatal(err)
}
fmt.Println("Status:", resp.Status)
fmt.Println("Content:", string(r.Content))

fmt.Println("Deleting...")
r = stingray.NewExtraFile("name")
resp, err = client.Delete(r)
if err != nil {
	log.Fatal(err)
}
fmt.Println("Status:", resp.Status)
```

```sh
Writing...
Status: 201 Created
Reading...
Status: 200 OK
Content: Test
Deleting...
Status: 204 No Content
```

## Supported Configuration Resources

Support for resources is being added as needed.

- [x] Action Program
- [ ] Alerting Action
- [ ] Aptimizer Application Scope
- [ ] Aptimizer Profile
- [ ] Bandwidth Class
- [ ] Cloud Credentials
- [ ] Custom configuration set
- [ ] Event Type
- [x] Extra File
- [ ] GLB Service
- [ ] Global Settings
- [x] License
- [ ] Location
- [x] Monitor
- [x] Monitor Program
- [ ] NAT Configuration
- [x] Pool
- [x] Protection Class
- [x] Rate Shaping Class
- [x] Rule
- [x] SLM Class
- [ ] SSL Client Key Pair
- [x] SSL Key Pair
- [x] SSL Trusted Certificate
- [ ] Security Settings
- [ ] Session Persistence Class
- [x] Traffic IP Group
- [x] Traffic Manager (incomplete)
- [ ] TrafficScript Authenticator
- [ ] User Authenticator
- [ ] User Group
- [x] Virtual Server

## Supported Statistics Resources
- [ ] Actions
- [ ] Bandwidth
- [ ] Cache
- [ ] Cloud API Credentials
- [ ] Connection Rate Limit
- [ ] Events
- [ ] Extras
- [ ] GLB Services
- [ ] Globals
- [ ] Listen IPs
- [ ] Locations
- [ ] Network Interface
- [ ] Nodes
- [ ] Per-location Service
- [ ] Per Node SLM
- [x] Pools
- [ ] Rule Authenticators
- [ ] Rules
- [ ] Service Level Monitors
- [ ] Service Protection
- [ ] SSL OCSP Stapling
- [ ] Traffic IPs
- [ ] Virtual Servers
