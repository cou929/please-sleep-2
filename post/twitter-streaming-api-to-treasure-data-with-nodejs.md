{"title":"Node.js で Twitter Streaming API のデータを Treasure Data に流しこむ","date":"2012-11-22T02:16:56+09:00","tags":["javascript, data"]}

![Treasure Data Logo](http://gyazo.com/bc4423bdc3d659bf2e76da5543668ae6.png?1353518187)

とても簡単で驚いた. CentOS で環境構築. node は入っているものとする.

1. twitter にアプリを登録して api をたたくのに必要なキー類を発行する
   - [Sign in with your Twitter account \| Twitter Developers](https://dev.twitter.com/apps/new)
2. Treasure Data のアカウントも Sign up
   - [Sign Up \| Treasure Data](https://www.treasure-data.com/signup/)
3. td コマンドの準備 ([Installing td command on Redhat and CentOS / Installing td command / Knowledge Base - Treasure Data Platform Support](http://help.treasure-data.com/kb/installation/installing-td-command-on-redhat-and-centos))

        $ curl -L https://get.rvm.io | bash -s stable --ruby
        $ source ~/.rvm/scripts/rvm
        $ rvm install 1.9.2 # takes time
        $ rvm use 1.9.2 --default
        $ gem install td
        $ td
        usage: td [options] COMMAND [args]

4. アカウントの authorize

        $ td account -f
        プロンプトが出るのでメアドとパスワードを入れる

5. td-agent のインストール
   1. treasure data のリポジトリを追加. 以下を `/etc/yum.repos.d/td.repo` に保存

            [treasuredata]
            name=TreasureData
            baseurl=http://packages.treasure-data.com/redhat/$basearch
            gpgcheck=0

    2. インストール

            $ sudo yum update
            $ sudo yum install -y td-agent

6. api key を td-agent.conf に設定
   1. api key を確認

            $ td apikey:show
   2. `/etc/td-agent/td-agent.conf` を編集

            <match td.*.*>
              apikey に上記のキーを設定

            <source>
              type forward
              port 24224 を追加

7. `sudo /etc/init.d/td-agent start`
8. package.json を準備

        {
          "name": "sample-app",
          "version": "0.0.1",
          "private": true,
          "dependencies": {
            "ntwitter": "~0.5.0",
            "fluent-logger": "~0.1.0"
          }
        }

9. `npm install`

準備はこれで OK

コードはこんな感じ

<script src="https://gist.github.com/4126075.js?file=twit_streaming_to_td_sample.js"></script>
[Twitter Streaming API を td へ流し込む — Gist](https://gist.github.com/4126075)
(Twitter の各種キーは自分で取得したものに変更)

これだけで, javascript でのtwitter 検索結果を treasure data に突っ込んでくれる. 結果を見るにはこんな感じ

    % td tables
    +----------+------------+------+-------+--------+
    | Database | Table      | Type | Count | Schema |
    +----------+------------+------+-------+--------+
    | test_db  | javascript | log  | 72    |        |
    | test_db  | test       | log  | 4     |        |
    +----------+------------+------+-------+--------+
    
    % td query -w -d test_db 'select v["text"] from javascript'
    Job 1138223 is queued.
    Use 'td job:show 1138223' to show the status.
    queued...
      started at 2012-11-21T17:10:17Z
      Hive history file=/tmp/1179/hive_job_log__22870500.txt
      Total MapReduce jobs = 1
      Launching Job 1 out of 1
      Number of reduce tasks is set to 0 since there's no reduce operator
      Starting Job = job_201209262127_100367, Tracking URL = http://ip-10-8-189-47.ec2.internal:50030/jobdetails.jsp?jobid=job_201209262127_100367
      Kill Command = /usr/lib/hadoop/bin/hadoop job  -Dmapred.job.tracker=10.8.189.47:8021 -kill job_201209262127_100367
      2012-11-21 17:10:39,495 Stage-1 map = 0%,  reduce = 0%
      finished at 2012-11-21T17:10:45Z
      2012-11-21 17:10:43,542 Stage-1 map = 100%,  reduce = 0%
      2012-11-21 17:10:44,555 Stage-1 map = 100%,  reduce = 100%
      Ended Job = job_201209262127_100367
      OK
      MapReduce time taken: 12.389 seconds
      Time taken: 12.529 seconds
    Status     : success
    Result     :
    +----------------------------------------------------------------------------------------------------------------------------------------------+
    | _c0                                                                                                                                          |
    +----------------------------------------------------------------------------------------------------------------------------------------------+
    | All the same, it's fun learning about JavaScript objects and prototypes and getting to grips with PHP namespaces                             |
    | Avanzamos de Sql Inyection a SSJS(server Side javascript-inyection) http://t.co/3vdW5ht8                                                     |
    | #recruiting ¦ Junior PHP Developer - £21,000 - Cheltenham ¦ #PHP #HTML #CSS #Javascript  ¦ http://t.co/AdV9mzbA                              |
    | Idea: tyson.js, the only plugin that punches you in the face for writing shitty JavaScript.                                                  |
    | javascript にコンパイル時に関空のお土産郵送って、やっぱり boost コンパイルしないと死ぬ。けどまぁ今回の用途ならいいんですよ遅刻はするけれど                                                               |
    | RT @clarkewoodnews: #recruiting ¦ Junior PHP Developer - £21,000 - Cheltenham ¦ #PHP #HTML #CSS #Javascript  ¦ http://t.co/AdV9mzbA          |
    | RT @SpringSource: #Javascript Dependency Analysis in the #Scripted Editor http://t.co/bKL55Git \n#springsource                               |
    | javascript:void(0);                                                                                                                          |
    | Apuntada! Iniciación a la programación con Javascript #cursosonline en @EscuelaIT  Empieza hoy! http://t.co/9EUp4GbR                         |
    
    ...

簡単すぎてほんとすごい.
