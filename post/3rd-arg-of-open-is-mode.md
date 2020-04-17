{"title":"open(2) の第三引数は新規作成時のモードを指定する","date":"2014-02-24T01:22:34+09:00","tags":["nix"]}

open(2) の第三引数に渡すのってなんだっけと思ったが, `O_CREAT` 指定で存在しないファイルを open したケースで新規作成されるファイルのモードを指定するんだった. 以下 man より:

    The oflag argument may indicate that the file is to be created if it does not exist (by specifying the O_CREAT flag).  In this case, open requires a third argument mode_t mode; the file is cre-
    ated with mode mode as described in chmod(2) and modified by the process' umask value (see umask(2)).
