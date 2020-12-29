# Go Salesforce

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

This package provides a number of Salesforce helpers in Go:

* `apex` performs Apex codegen, specifically for sending email.
* `fsdb` provides a Golang SDK and filesystem cache for Salesforce using the Salesforce REST API.
* `sobjects` provides basic structs for Salesforce.
* `workbench` provies a basic XML reader for https://workbench.developerforce.com

For OAuth 2.0 helpers for the Salesforce API, see [`oauth2more/salesforce`](https://github.com/grokify/oauth2more/tree/master/salesforce).

## Reference Files

### Entity Prefixes

The `entityprefixes.csv` file contains the prefixes from the Salesforce Standard Field Record ID Prefix Decoder, Knowledge Article Number: 000005995. This webpage is located here:

[https://help.salesforce.com/apex/HTViewSolution?urlname=Standard-Field-Record-ID-Prefix-Decoder&language=en_US](https://help.salesforce.com/apex/HTViewSolution?urlname=Standard-Field-Record-ID-Prefix-Decoder&language=en_US)

For more see Daniel Ballinger's website as mentioned by Salesforce:

* [Obscure Salesforce object key prefixes](http://www.fishofprey.com/2011/09/obscure-salesforce-object-key-prefixes.html)

## Tools

### Salesforce CLI:

* https://developer.salesforce.com/tools/sfdxcli
* https://developer.salesforce.com/docs/atlas.en-us.sfdx_setup.meta/sfdx_setup/sfdx_setup_install_cli.htm

### VS Code Extension

https://salesforcedx-vscode.netlify.app/tools/vscode/en/getting-started/install

## Contributing

Features, Issues, and Pull Requests are always welcome.

 [build-status-svg]: https://github.com/grokify/go-salesforce/workflows/build/badge.svg
 [build-status-link]: https://github.com/grokify/go-salesforce/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-salesforce
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-salesforce
 [codeclimate-status-svg]: https://codeclimate.com/github/grokify/go-salesforce/badges/gpa.svg
 [codeclimate-status-link]: https://codeclimate.com/github/grokify/go-salesforce
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-salesforce
 [docs-godoc-link]: https://pkg.go.dev/github.com/grokify/go-salesforce
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-salesforce/blob/master/LICENSE
