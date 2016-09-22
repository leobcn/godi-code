# [Dependency Injection and Testable Web Development in Go](http://blog.extremix.net/post/di)

This repository contains code accompanying the post above.

## Installation
- Download this repository with `go get -d https://github.org/kkrs/godi-code`
- Download and setup 
[Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).

## Running Tests
Running the e2e test requires that the appengine development server be running. It can be started with
`dev_appserver.py --skip_sdk_update_check=true --clear_datastore=true .`.

Test for this branch can be run with `go test -v`.
