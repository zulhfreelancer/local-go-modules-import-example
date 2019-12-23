## How to import local Go module from another directory?

This project is located inside `$GOPATH/src/test-go-github` while the local Go module that I want to use (`zz.dev/go-github/v28/github`) is located inside `$GOPATH/src/go-github`.

After running `go build -i -v zz.dev/go-github/v28/github && go install` inside `$GOPATH/src/go-github`, I added the line below inside this project's _go.mod_ file...

```
replace zz.dev/go-github/v28 => /Users/my-name/go/src/go-github
```

...where `/Users/my-name/go` is my `$GOPATH`. Do note that I didn't put `/github` at the end of path on both sides. It worked.
