[![Go Reference](https://pkg.go.dev/badge/github.com/toqueteos/lumberjack.svg)](https://pkg.go.dev/github.com/toqueteos/lumberjack)

# Lumberjack (forked)

Forked from: [natefinch/lumberjack](https://github.com/natefinch/lumberjack)

Lumberjack is a Go package for writing logs to rolling files.

Lumberjack is intended to be one part of a logging infrastructure.
It is not an all-in-one solution, but instead is a pluggable
component at the bottom of the logging stack that simply controls the files
to which logs are written.

Lumberjack plays well with any logging package that can write to an
[io.Writer](https://pkg.go.dev/io#Writer), including the standard library's log package.

Lumberjack assumes that only one process is writing to the output files.
Using the same lumberjack configuration from multiple processes on the same
machine will result in improper behavior.

## Install

```bash
go get -u github.com/toqueteos/lumberjack
```

## Example

To use lumberjack with the standard library's log package, just pass it into the SetOutput function when your application starts.

Code:

```go
log.SetOutput(&lumberjack.Logger{
    Filename:   "/var/log/myapp/foo.log",
    MaxSize:    500, // megabytes
    MaxBackups: 3,
    MaxAge:     28, //days
    Compress:   true, // disabled by default
})
```
