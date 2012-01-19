webtestutil - Google Go (golang) utilities for web tests
====================================

## DESCRIPTION

This package provides some helpers and useful types for web testing. Including, but not limited to the [Gorilla web toolkit](http://gorilla-web.appspot.com/).

## INSTALLATION

Install using goinstall:

    goinstall http://github.com/chlu/webtestutil

Clone and build using gomake:

    git clone http://github.com/chlu/webtestutil
    cd webtestutil
    gomake install

## USAGE

### Testing

This package provides some tools to write functional tests for web
applications in Google Go. They  should not be a replacement for proper
unit tests, but are a simple (almost) end-to-end test for a web application.

### Sessions

If you use the (quite awesome) [sessions package](http://gorilla-web.appspot.com/pkg/gorilla/sessions/) of the Gorilla web toolkit,
you need some glue code to simulate sessions inside a test.

Import the webutils package in your test:

    package my_test

    import (
        "testing"

        "webtestutil"
    )

Initialize the testing session store before the tests:

    func init() {
        webtestutil.RegisterTestingStore()
    }

Make sure to reset the session after each request:

    func TestMyRequest() {
        r, _ := http.NewRequest("GET", "/path", nil)
    	w := httptest.NewRecorder()

        defer webtestutil.ResetSession()

        myHandler.ServeHTTP(w, r)
    }
    
Put some values in the session for your tests:

    func TestMyRequest() {
        r, _ := http.NewRequest("GET", "/path", nil)
    	w := httptest.NewRecorder()

        webtestutil.SessionData()["username"] = "j.doe"
        defer webtestutil.ResetSession()

        myHandler.ServeHTTP(w, r)
    }

## LICENSE

This package is licensed under an MIT license (see LICENSE).