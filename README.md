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
