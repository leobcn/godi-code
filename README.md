# [Dependency Injection and Testable Web Development in Go](http://blog.extremix.net/post/di)

This repository contains code for the blog post above.

## Installation
- Download and setup 
[Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).
- Download this repository with `go get -d github.com/kkrs/godi-code`

## Running Tests
Running the end-to-end test requires that the appengine development server be running. It can be
started by running `dev_appserver.py --skip_sdk_update_check=true --clear_datastore=true app`
from the repository root.

The test for this branch can then be run with `go test -v`.
