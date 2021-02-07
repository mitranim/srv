## Overview

Extremely simple Go tool that serves files out of a given folder, using a file resolution algorithm similar to GitHub Pages, Netlify, or the default Nginx config. Provides a Go "library" (less than 100 LoC) and an optional CLI tool.

## Why

Useful when making a site that will be served by GitHub Pages or similar. Usable from Go code. Usable from a shell. Not overweight.

## Usage

An almost drop-in replacement for `http.FileServer`:

```go
import "github.com/mitranim/srv"

http.ListenAndServe(":<some-port>", srv.FileServer("."))
```

For CLI usage, first install Go: https://golang.org. Then run this:

```sh
go get -u github.com/mitranim/srv/srv

# Get help
srv -h

# Actually run
srv
```

## File Resolution

* For directories including `/`, use `index.html` when available.
* For "clean" links without an extension, try appending `.html`.
* For "not found", try using `404.html` if one exists, falling back on a simple hardcoded message.

## License

https://unlicense.org

## Misc

I'm receptive to suggestions. If this library _almost_ satisfies you but needs changes, open an issue or chat me up. Contacts: https://mitranim.com/#contacts
