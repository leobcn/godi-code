# [Dependency Injection and Testable Web Development in Go](http://blog.extremix.net/post/di)

This repository contains code for the blog post above.

## Installation
- Download and setup 
[Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).
- Download this repository with `go get -d github.com/kkrs/godi-code`

The repository has two branches: `first` and `di` after the sections 'The First Attempt' and
'Depdendency Injection'.

## Running Tests
Running the end-to-end test requires that the appengine development server be running. It can be
started by running `dev_appserver.py --skip_sdk_update_check=true --clear_datastore=true app`
from the repository root. Since the development server stores state, re-running the end-to-end test
requires that the development server be restarted.

### first
The test for this branch can then be run with `go test -v`.

### di
The branch `di` has two tests, e2e, int and can be run with
`go test -v -tags=e2e` # end-to-end, requires development server to be started
`go test -v -tags=int` # no need for development server
