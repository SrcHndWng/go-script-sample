# About This

https://blog.cloudflare.com/using-go-as-a-scripting-language-in-linux/
を写経、改変したもの。

## Usage

### Before Usage

```
$ go get github.com/erning/gorun
$ sudo cp $GOPATH/bin/gorun /usr/local/bin/
```

gorunをgo getで取得した後に、どこからでも呼べるように/usr/local/bin/にgorunをコピーする。

### Run

```
$ gorun go_script.go 
your input is non!

$ gorun go_script.go sample
your input is sample!
$ echo $?
0

$ gorun go_script.go error
your input is error!
2019/10/31 21:24:28 error raise
$ echo $?
1
```
