# shinshutsu

gongfu tea timer (for lazy brewing)

## Installations

I couldn't figure out how to do this less laboriously.

This assumes `$GOPATH` and `$GOBIN` to be setup properly.

For `Go` version `1.16` and above.

```sh
go env -w GO111MODULE=off
```

Dependencies

```sh
go get github.com/mitchellh/go-homedir
go get gopkg.in/yaml.v2
go get github.com/onsi/ginkgo
go get github.com/onsi/gomega
```

Installations. One by one.

```sh
go get github.com/EmilRehnberg/shinshutsu
cd $GOPATH/src/github.com/EmilRehnberg/shinshutsu
cd alarmclock
go install
cd ../brewci
go install
cd ../executables/countdown
go install
cd ../shinshutsu
go install
```
