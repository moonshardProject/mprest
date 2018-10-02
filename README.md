# mprest
**mprest** which stands for MoonshardProject REST, is a RESTful API server. On startup it scans the devices on the network using ping sweep (using `fping` tool) and checks whether if the device runs a web server and its index.html, matches the IOT devices html defined in `/pages/` directory. After scanning the IOT devices, it runs an http server and serves the following jsons:

`/api/verification` it send back a status code with a name field with the value of 'moonshard'. This route enables the mobile apps and other external softwares to verify that this rest server belong the the moonshardProject.
`/api/devices` send the number and the IP addresses of the IOT devices that found via scanning.

Another method which we could use to verify that the found hosts are IOT devices, is to return a validation data (like a json file for example) from the device back to the RESTful api server, but this method needs reprogram of the IOT devices and makes their verification non-generic.

### Building
For building the project for Raspberry Pi, use the following command for building:
`env GOOS=linux GOARCH=arm GOARM=5 go build`