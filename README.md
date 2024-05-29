<h1 align="center">
  favirecon
  <br>
</h1>

<h4 align="center">Use favicon.ico to improve your target recon phase</h4>

<h6 align="center"> Coded with 💙 by edoardottt </h6>

<p align="center">

  <a href="https://github.com/edoardottt/favirecon/actions">
      <img src="https://github.com/edoardottt/favirecon/actions/workflows/go.yml/badge.svg" alt="go action">
  </a>

  <a href="https://goreportcard.com/report/github.com/edoardottt/favirecon">
      <img src="https://goreportcard.com/badge/github.com/edoardottt/favirecon" alt="go report card">
  </a>

<br>
  <!--Tweet button-->
  <a href="https://twitter.com/intent/tweet?text=favirecon%20-%20Use%20favicon.ico%20to%20improve%20your%20target%20recon%20phase.%20Quickly%20detect%20technologies,%20WAF,%20exposed%20panels,%20known%20services.%20https%3A%2F%2Fgithub.com%2Fedoardottt%2Ffavirecon%20%23golang%20%23github%20%23linux%20%23infosec%20%23bugbounty" target="_blank">Share on Twitter!
  </a>
</p>

<p align="center">
  <a href="#install-">Install</a> •
  <a href="#get-started-">Get Started</a> •
  <a href="#examples-bulb">Examples</a> •
  <a href="#changelog-">Changelog</a> •
  <a href="#contributing-">Contributing</a> •
  <a href="#license-">License</a>
</p>

<p align="center">
  <img src="https://github.com/edoardottt/images/blob/main/favirecon/favirecon.gif">
</p>
  
> **Note**
> This tool heavily relies on its favicon hash database. If you think you have a new favicon hash that's worth adding or you think there is a wrong hash-service association please open an issue.  

Install 📡
----------

### Homebrew

```console
brew install favirecon
```

### Snap

```console
sudo snap install favirecon
```

### Go

```console
go install github.com/edoardottt/favirecon/cmd/favirecon@latest
```

Get Started 🎉
----------

```console
Usage:
  favirecon [flags]

Flags:
INPUT:
   -u, -url string   Input domain
   -l, -list string  File containing input domains
   -cidr             Interpret input as CIDR

CONFIGURATIONS:
   -hash string[]        Filter results having these favicon hashes (comma separated)
   -c, -concurrency int  Concurrency level (default 50)
   -t, -timeout int      Connection timeout in seconds (default 10)
   -rl, -rate-limit int  Set a rate limit (per second)
   -px, -proxy string    Set a proxy server (URL)

OUTPUT:
   -o, -output string  File to write output results
   -v, -verbose        Verbose output
   -s, -silent         Silent output. Print only results
   -j, -json           JSON output
```

Examples 💡
----------

Identify a single domain

```console
favirecon -u https://www.github.com
```

```console
echo https://www.github.com | favirecon
```

Grab all possible results from a list of domains (protocols needed!)

```console
favirecon -l targets.txt
```

```console
cat targets.txt | favirecon
```

Grab all possible results belonging to a specific target(s) (protocols needed!)

```console
cat targets.txt | favirecon -hash 708578229
```

Grab all possible results from single CIDR

```console
favirecon -u 192.168.1.0/24 -cidr
```

Use a Proxy

```console
favirecon -u https://www.github.com -px http://127.0.0.1:8080
```

JSON Output

```console
favirecon -u https://www.github.com -j
```

Changelog 📌
-------

Detailed changes for each release are documented in the [release notes](https://github.com/edoardottt/favirecon/releases).

Contributing 🛠
-------

Just open an [issue](https://github.com/edoardottt/favirecon/issues) / [pull request](https://github.com/edoardottt/favirecon/pulls).

Before opening a pull request, download [golangci-lint](https://golangci-lint.run/usage/install/) and run

```console
golangci-lint run
```

If there aren't errors, go ahead :)

In the news 📰
-------

- [Hive Five Newsletter by Securibee](https://securib.ee/newsletter/)
  
License 📝
-------

This repository is under [MIT License](https://github.com/edoardottt/favirecon/blob/main/LICENSE).  
[edoardoottavianelli.it](https://www.edoardoottavianelli.it) to contact me.
