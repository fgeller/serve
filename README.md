# serve

Tiny app to serves files in current directory via HTTP. Attempts to open the
default browser at the served location for Linux and OSX.

Install via

    go get github.com/fgeller/serve

Make sure `$GOPATH/bin` is part of your `$PATH`.

Simply start via `serve` which will choose a random port:

    $ serve
    2015/10/26 09:26:05 Serving at http://0.0.0.0:60794.
    2015/10/26 09:26:05 Opening...

Or specify a port:

    $ serve 3022
    2015/10/26 09:43:36 Serving at http://0.0.0.0:3022
    2015/10/26 09:43:36 Opening...

Very similar to Python's `SimpleHTTPServer` module, which I had wrapped in the
following bash function:

    function serve {
        python -m SimpleHTTPServer "$@"
    }

Turns out it's not straight-forward to detect the port on which the server is
listening, so I couldn't simply add a step to open my browser at a newly served
location.
