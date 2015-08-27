# stash-go-import

> Go [import path](http://golang.org/cmd/go/#hdr-Remote_import_paths) support
> for [Atlassian Stash](https://www.atlassian.com/software/stash)

Makes it possible to use clean import paths such as
`git.mycompany.com/proj/repo` instead of the ugly
`git.mycompany.com:7999/proj/repo.git`.

# Installation

```
go get github.com/rubenv/stash-go-import
```

# Usage

```
Usage: stash-go-import

  -port=80: listen on port
  -sshPort=7999: use SSH port
```

# Docker image

There's a docker image available: [rubenv/stash-go-import](https://registry.hub.docker.com/u/rubenv/stash-go-import/).

Simply run the image and you're good to go:

```
docker run -d -p 7991:80 rubenv/stash-go-import
```

To pass extra arguments:

```
docker run -d -p 7991:80 rubenv/stash-go-import stash-go-import -sshPort 8001
```

# Configuring nginx

```
server {
    listen 443 ssl spdy;
    server_name git.mycompany.com;

    spdy_headers_comp 7;

    gzip             on;
    gzip_min_length  1000;
    gzip_types       text/plain application/xml text/css application/javascript application/json application/x-javascript;
    gzip_disable     "MSIE [1-6]\.";

    if ($args ~* "^go-get=1") {
        set $condition goget;
    }

    location / {
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        if ($condition = goget) {
            proxy_pass http://localhost:7991;
        }
        if ($condition != goget) {
            proxy_pass http://localhost:7990;
        }
    }
}
```

This will send any call with `?go-get=1` to `stash-go-import`, any other call
goes to Stash. Adjust to match your local setup.

## License

    (The MIT License)

    Copyright (C) 2015 by Ruben Vermeersch <ruben@rocketeer.be>

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
    THE SOFTWARE.
