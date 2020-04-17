{"title":"sphinx で ValueError: unknown locale: UTF-8 というエラーが出た","date":"2012-11-24T13:04:31+09:00","tags":["python"]}

[uvbook](https://github.com/nikhilm/uvbook) を読もうと思って make html したら `ValueError: unknown locale: UTF-8` という例外がでてビルドできなかった

    (dev)[12:34 kosei@mba uvbook]% make html
    sphinx-build -b html -d build/doctrees   source build/html
    Traceback (most recent call last):
      File "/Users/kosei/.virtualenvs/dev/bin/sphinx-build", line 9, in <module>
        load_entry_point('Sphinx==1.1.3', 'console_scripts', 'sphinx-build')()
      File "/Users/kosei/.virtualenvs/dev/lib/python2.7/site-packages/sphinx/__init__.py", line 47, in main
        from sphinx import cmdline
      File "/Users/kosei/.virtualenvs/dev/lib/python2.7/site-packages/sphinx/cmdline.py", line 18, in <module>
        from docutils.utils import SystemMessage
      File "/Users/kosei/.virtualenvs/dev/lib/python2.7/site-packages/docutils/utils/__init__.py", line 19, in <module>
        from docutils.io import FileOutput
      File "/Users/kosei/.virtualenvs/dev/lib/python2.7/site-packages/docutils/io.py", line 18, in <module>
        from docutils.error_reporting import locale_encoding, ErrorString, ErrorOutput
      File "/Users/kosei/.virtualenvs/dev/lib/python2.7/site-packages/docutils/error_reporting.py", line 47, in <module>
        locale_encoding = locale.getlocale()[1] or locale.getdefaultlocale()[1]
      File "/Users/kosei/.virtualenvs/dev/lib/python2.7/locale.py", line 496, in getdefaultlocale
        return _parse_localename(localename)
      File "/Users/kosei/.virtualenvs/dev/lib/python2.7/locale.py", line 428, in _parse_localename
        raise ValueError, 'unknown locale: %s' % localename
    ValueError: unknown locale: UTF-8
    make: *** [html] Error 1

どうも Mac でよく起こる現象のようだ.

### 対応方法

mac では LC_ALL が空になっていることがあるらしく, ここが空だとうまくいかない. その場合は `export LC_ALL='ja_JP.UTF-8'` してあげればいい.

    (dev)[12:41 kosei@mba uvbook]% locale
    LANG="ja_JP.UTF-8"
    LC_COLLATE="ja_JP.UTF-8"
    LC_CTYPE="UTF-8"
    LC_MESSAGES="ja_JP.UTF-8"
    LC_MONETARY="ja_JP.UTF-8"
    LC_NUMERIC="ja_JP.UTF-8"
    LC_TIME="ja_JP.UTF-8"
    LC_ALL=
    (dev)[12:41 kosei@mba uvbook]% export LC_ALL='ja_JP.UTF-8'
    (dev)[12:41 kosei@mba uvbook]% make html
    sphinx-build -b html -d build/doctrees   source build/html
    Making output directory...
    Running Sphinx v1.1.3
    loading pickled environment... done
    building [html]: targets for 10 source files that are out of date
    updating environment: 0 added, 0 changed, 0 removed
    looking for now-outdated files... none found
    preparing documents... done
    writing output... [ 10%] about
    writing output... [ 20%] basics
    writing output... [ 30%] filesystem
    writing output... [ 40%] index
    writing output... [ 50%] introduction
    writing output... [ 60%] multiple
    writing output... [ 70%] networking
    writing output... [ 80%] processes
    writing output... [ 90%] threads
    writing output... [100%] utilities
    
    writing additional files... genindex search
    copying static files... WARNING: html_static_path entry '/Users/kosei/src/uvbook/source/_static' does not exist
    done
    dumping search index... done
    dumping object inventory... done
    build succeeded, 1 warning.
    
    Build finished. The HTML pages are in build/html.

あるいは .bashrc などに `export LC_ALL='ja_JP.UTF-8'` を書いておいて永続化しておけばよい.

### 原因

python ビルトインの locale モジュール (`/usr/lib/python2.7/locale.py` とか, virtualenv だったら `~/.virtualenvs/dev/lib/python2.7/locale.py` とか) の `getdefaultlocale()` 関数で, こんな感じで環境変数からロケール文字列を取得している処理がある.

<script src="https://gist.github.com/4138273.js?file=locale.py"></script>

環境変数を

- LC_ALL
- LC_CTYPE
- LANG
- LANGUAGE

の順にみていって見つかり次第ループを抜けているが, mac の場合 (少なくても自分の環境では) LC_ALL がからっぽで, 次の LC_CTYPE には 'UTF-8' という文字列が入っていた

    >>> import os
    >>> os.environ.get('LC_ALL');
    >>> os.environ.get('LC_CTYPE');
    'UTF-8'
    >>> os.environ.get('LANG');
    'ja_JP.UTF-8'
    >>> os.environ.get('LANGUAGE');

環境変数から値をとったあとの `_parse_localename()` という関数は `ja_jp.UTF-8` のようなフォーマットの文字列を期待しているので, `UTF-8` が渡されるて例外を吐いてしまっていたようだ.

今回の場合は `LC_ALL` に値を入れてあげれば `LC_TYPE` の前に見つかってループからぬけるので良い.

    $ export LC_ALL='ja_JP.UTF-8'

### よくわからない / あとで調べる

- 自分の環境が特殊だったのか, locale.py の方を直したほうがよいのか.
  - よくあるエラーだったとしたら python のコアモジュールだしすぐなおりそうなものだし, 自分の環境が悪いのかもしれない
  - [Fresh installation of sphinx-quickstart fails - Stack Overflow](http://stackoverflow.com/questions/10921430/fresh-installation-of-sphinx-quickstart-fails) を見る限り snow leopard で起こりやすい現象みたいだ
- LC_ALL, LC_CTYPE などなど, ロケール系の環境変数の意味
  - LC_CTYPE に 'UTF-8' が入っているのは問題ないのか. そうだとしたら `_parse_localename` が悪いのかもしれないし
