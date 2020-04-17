{"title":"dropbox を複数インスタンス起動する","date":"2013-09-30T22:46:40+09:00","tags":["mac"]}

`$HOME` を変えるだけで、あとはふつうにきどうすればよい。つぎのように nohup であげておくといいかもしれない。

    HOME=/path/to/dropbox/dir nohup /Applications/Dropbox.app/Contents/MacOS/Dropbox &

### 参考

[1台のPCで複数のDropboxアカウントを使う方法 : ライフハッカー［日本版］](http://www.lifehacker.jp/2013/01/130110multipledropbox.html)
