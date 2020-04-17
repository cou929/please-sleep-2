{"title":"RDS の mysql でスレッド・クエリを kill する","date":"2015-02-14T23:09:29+09:00","tags":["nix"]}

kill ではなく、RDS 固有のコマンド `mysql.rds_kill` を使わないといけない。マネージドのサービスなので、すべての権限を渡していませんということらしい。

    mysql> SELECT * FROM information_schema.PROCESSLIST;
    ...
    mysql> CALL mysql.rds_kill(thread-ID)
    mysql> CALL mysql.rds_kill_query(thread-ID)

その他にも、レプリケーションのエラーをスキップする `CALL mysql.rds_skip_repl_error;` など、いくつかのコマンドは独自コマンドを使う必要がある。

[付録: MySQL に関連する一般的な DBA タスク - Amazon Relational Database Service](http://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/Appendix.MySQL.CommonDBATasks.html)

ちなみにプロセスリストを見ると rdsadmin というユーザーがいる。すべての権限を持った自動で作られるユーザーで、こいつがバックアップをとったりしてくれているらしい。

    mysql> select * from information_schema.PROCESSLIST WHERE USER = 'rdsadmin';
    +----+----------+-----------------+-------+---------+------+-------+------+
    | ID | USER     | HOST            | DB    | COMMAND | TIME | STATE | INFO |
    +----+----------+-----------------+-------+---------+------+-------+------+
    |  8 | rdsadmin | localhost:27138 | mysql | Sleep   |    9 |       | NULL |
    +----+----------+-----------------+-------+---------+------+-------+------+
    1 row in set (0.02 sec)
    mysql> SELECT * FROM mysql.user\G
    *************************** 1. row ***************************
                      Host: localhost
                      User: rdsadmin
                  Password: *901DE9BC618FBCE207050F268074B590BEDEC78A
               Select_priv: Y
               Insert_priv: Y
               Update_priv: Y
               Delete_priv: Y
               Create_priv: Y
                 Drop_priv: Y
               Reload_priv: Y
             Shutdown_priv: Y
              Process_priv: Y
                 File_priv: Y
                Grant_priv: Y
           References_priv: Y
                Index_priv: Y
                Alter_priv: Y
              Show_db_priv: Y
                Super_priv: Y
     Create_tmp_table_priv: Y
          Lock_tables_priv: Y
              Execute_priv: Y
           Repl_slave_priv: Y
          Repl_client_priv: Y
          Create_view_priv: Y
            Show_view_priv: Y
       Create_routine_priv: Y
        Alter_routine_priv: Y
          Create_user_priv: Y
                Event_priv: Y
              Trigger_priv: Y
    Create_tablespace_priv: Y
                  ssl_type:
                ssl_cipher:
               x509_issuer:
              x509_subject:
             max_questions: 0
               max_updates: 0
           max_connections: 0
      max_user_connections: 0
                    plugin:
     authentication_string:
          password_expired: N


### 参考

- [付録: MySQL に関連する一般的な DBA タスク - Amazon Relational Database Service](http://docs.aws.amazon.com/ja_jp/AmazonRDS/latest/UserGuide/Appendix.MySQL.CommonDBATasks.html)
- [Amazon RDS な MySQL で 不要 process を kill する - Garbage in, gospel out](http://libitte.hatenablog.jp/entry/20141202/1417452564)
