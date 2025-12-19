# BUILD

## BUILD IN WINDOWS

```shell
# FOR LINUX
$env:GOOS = "linux"; $env:GOARCH="amd64"; go build -a -v -tags netgo -ldflags '-w' -o bin\router-gostashlg-template

# FOR WINDOWS
$env:GOOS = "windows"; go build -a -tags netgo -ldflags '-w' -o .\bin\router-gostashlg-template.exe
```

## BUILD IN LINUX OR MAC

```shell
# FOR LINUX
GOOS=linux GOARCH=amd64 go build -a -v -tags netgo -ldflags '-w' -o bin/router-gostashlg-template

#FOR WINDOWS
GOOS=windows GOARCH=amd64 go build -a -v -tags netgo -ldflags '-w' -o bin/router-gostashlg-template.exe
```
