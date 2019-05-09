# rem-word

命令行里背单词：终端显示单词，键盘操作。

`word_freq.csv`摘选了经济学人中出现10次以上的单词，总计35275个。

### run

```
go get ./...
cp conf.toml.example conf.toml
// add your youdao id/secret
go build
./rem-word -csv_file=word_freq_sample.csv
```



### deps

https://github.com/nsf/termbox-go
