[![made-with-Go](https://img.shields.io/badge/made%20with-Go-brightgreen.svg)](http://golang.org)
<h1 align="center">Scopein</h1> <br>

<p align="center">
  <a href="#--usage--explanation">Usage</a> •
  <a href="#--installation--requirements">Installation</a>
</p>

<h3 align="center">Scopein only brings to you urls in scope, facilitating you recon process and avoiding scanning out of scope websites</h3>

<img src="https://cdn.discordapp.com/attachments/876919540682989609/977918443720417280/unknown.png">


## Contents:

- [Installation](#--installation--requirements)
- [Usage](#--usage--explanation)
  - [Usage with In Scope Targets](#printing-only-in-scope-targets)
  - [Usage with Out of Scope Targets](#printing-only-in-scope-targets-by-setting-out-of-scope-urls)
  - [Reading from config file](#reading-from-config-file)
  - [Using with other tools](#chaining-with-other-tools)

## - Installation & Requirements:

Installing the tool ->

Using go
```bash
▶ go install github.com/ferreiraklet/scopein@latest
```

Using git clone
```bash
▶ git clone https://github.com/ferreiraklet/scopein.git
▶ cd scopein
▶ go build scopein.go
▶ chmod +x scopein
▶ ./scopein -h
```
<br>


## - Usage & Explanation:

In Your recon process, when doing subdomain recon & url recon, you may get urls that is not in the scope, such as: "bit.ly", "twitter.com", random urls or subdomains. Here comes scopein, he only shows in terminal in scope urls

scopein is very easy to use, follow the steps =>

<br>
  
### Printing only In scope targets

Pay attention to the syntax!
```bash
cat urls | scopein -s "targetscope.com"
cat subdomains | scopein -s "targetscope.com|targetscope2.com"
```

### Printing only In scope targets by setting Out of Scope urls
 
```bash
cat targets | scopein -b "outscope.com"
cat targets | scopein -b "outscope.com|outscope2.com"
```

### Reading from config file

```bash
cat targets
---output---
https://google.com
https://redacted.com
https://example.com
https://twitter.com

cat scope
---output---
google.com
twitter.com


Print urls from in scope config file ->
cat targets | scopein -f scope

Print permitted urls from out of scope file
cat targets | scopein -bf scope
```


### Chaining with other tools

```bash
echo "http://testphp.vulnweb.com" | waybackurls | scopein -s "testphp.vulnweb.com"

echo "http://testphp.vulnweb.com" | waybackurls | scopein -b "twitter.com"

echo "http://testphp.vulnweb.com" | hakrawler | scopein -s "testphp.vulnweb.com"

echo "http://testphp.vulnweb.com" | gauplus -b svg,jpg,png,gif,pdf,js,css | scopein -f scopes

echo "http://testphp.vulnweb.com" | gau | scopein -bf scopes
```
    

## Check out some of my other programs <br>

> [Nilo](https://github.com/ferreiraklet/nilo) - Checks if URL has status 200

> [Jeeves](https://github.com/ferreiraklet/jeeves) - Time based blind Injection Scanner

> [Airixss](https://github.com/ferreiraklet/airixss) - XSS Reflected recon


If any error in the program, talk to me immediatly.

Contact:

Discord: ferreira#9313
Twitter: ferreiraklet

## This project is for educational and bug bounty porposes only! I do not support any illegal activities!.
