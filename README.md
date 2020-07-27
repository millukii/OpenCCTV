<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p>

<h3 align="center">Project Title</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> Few lines describing your project.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

Write about 1-2 paragraphs describing the purpose of your project.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them.

```
Give examples
```

### Installing

A step by step series of examples that tell you how to get a development env running.

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo.

## 🔧 Running the tests <a name = "tests"></a>

Explain how to run the automated tests for this system.

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## 🎈 Usage <a name="usage"></a>

Add notes about how to use the system.

## 🚀 Deployment <a name = "deployment"></a>

Add additional notes about how to deploy this on a live system.

## ⛏️ Built Using <a name = "built_using"></a>

- [MongoDB](https://www.mongodb.com/) - Database
- [Express](https://expressjs.com/) - Server Framework
- [VueJs](https://vuejs.org/) - Web Framework
- [NodeJs](https://nodejs.org/en/) - Server Environment

## ✍️ Authors <a name = "authors"></a>

- [@kylelobo](https://github.com/kylelobo) - Idea & Initial work

See also the list of [contributors](https://github.com/kylelobo/The-Documentation-Compendium/contributors) who participated in this project.

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
- Inspiration
- References

## ISAPI

The OPEN Intelligent Security Application Programming Interface (ISAPI) is a text protocol in
RESTful style based on HTTP for communicating between security devices/servers (e.g., cameras,
DVR, NVR, etc.) and client software/system. It defines the communication standard between
device/server and client software/system via the Internet Protocol (IP)

# HTTP Methods

- POST: Create resources. This method is only available for adding resource that does not exist before.
- GET: Retrieve resources. This method cannot change the system status, only return data as the response to the requester.
- PUT: Update resources. This method is usually for update the resource that already exists, but it can also be used to create the resource if the specific resource does not exist.
- DELETE: Delete resources.

# URL FORMAT

- <protocol>://<host>[:port][abs_path [?query]]

# Message Format

- For ISAPI protocol, the request and response messages generated among the interaction between
devices and platform are data in XML format or JSON format.
- For configuration Content-Type in the JSON format message is "application/json; charset='UTF-8'"
- For data Content-Type in the JSON format message is "application/octet-stream"

## Time Format
- The time format in the ISAPI protocol adopts ISO8601 standard (see details in http://
www.w3.org/TR/NOTE-datetime-970915 ), that is, YYY-MM-DDThh:mm:ss.sTZD (e.g.,
2017-08-16T20:17:06+08:00).

## Error Processing
- During the integration applications of ISAPI protocol, when the error of URL based on HTTP
occurred, the ResponseStatus message (in XML or JSON format) which contains error code will be
returned.

## Authentication 
- The authentication must based on HTTP Authentication: Basic and Digest Access Authentication,
see https://tools.ietf.org/html/rfc2617 for details.
- The request session must contain authentication information, otherwise, device will return 401
error code.

## Digest
- The message digest, which contains user name, password, specific nonce value, HTTP or RTSP
operation methods, and request URL, is generated by the MD5 algorithm, see the calculation rules
below.
qop=Undefined
Digest=MD5(MD5(A1):<nonce>:MD5(A2))
qop="auth:"
Digest=MD5(MD5(A1):<nonce>:<nc>:<cnonce>:<qop>:MD5(A2))
qop="auth-int:"
Digest=MD5(MD5(A1):<nonce>:<nc>:<cnonce>:<qop>:MD5(A2))