# git-fast-import-go

This is a package to easily create streams for `git fast-import`.

## How to get

```
$ go get https://github.com/steeve/git-fast-import-go
```

## How to use

```go
    committer := &gitfastimport.Signature{
        Name:  "Foo Bar",
        Email: "foo@bar.com",
        When: time.Now()
    }

    gitfastimport.WriteCommit(os.Stdout, "refs/heads/master", "git message", Committer, Committer)
    gitfastimport.WriteFileModify(os.Stdout, 644, "myfile.txt")
    gitfastimport.WriteDataBegin(os.Stdout)
    os.Stdout.Write([]byte("File content"))
    gitfastimport.WriteDataEnd(os.Stdout)
```

Then run:
```
$ go run program.go | git fast-import
```
