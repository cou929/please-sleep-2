{"title":"CentOS5 に python 開発環境構築 (がうまくいかなかった)","date":"2012-12-18T12:57:54+09:00","tags":["nix, python"]}

目標は virtualenvwrapper まで準備すること

よく覚えていないけど python2.6 は入っていたので (ほんとは 2.7 がよかったが...) これをメインの python 処理系として使うことにする. たぶん epel から入れたんだと思う.

    $ ls -l /usr/bin/python*
    lrwxrwxrwx 1 root root   18 12月  4 21:20 /usr/bin/python -> /usr/bin/python2.4
    -rwxr-xr-x 1 root root 8304  9月 22  2011 /usr/bin/python2.4
    -rwxr-xr-x 2 root root 4736 11月  7 23:48 /usr/bin/python2.6
    -rwxr-xr-x 2 root root 4736 11月  7 23:48 /usr/bin/python26

setuptools がなかったのでいれる

    # こうしないと 2.6 の site-packages に入らない
    $ sudo sh -c 'curl http://peak.telecommunity.com/dist/ez_setup.py | python26'

virtualenv と virtualenvwrapper を入れる

    # 2.6 の site-packages に入れて欲しいので easy_install はフルパス指定する
    $ sudo /usr/bin/easy_install-2.6 virtualenv
    $ sudo /usr/bin/easy_install-2.6 virtualenvwrapper

`.zshrc` に virtualenvwrapper の設定を追加

    PYTHON_VER=2.6
    export VIRTUALENV_BIN=/usr/bin/virtualenv-$PYTHON_VER
    export PYTHONPATH=/usr/lib/python$PYTHON_VER/:$PYTHONPATH
    export WORKON_HOME=$HOME/.virtualenvs
    . /usr/bin/virtualenvwrapper.sh
    
    mkvenv () {
        base_python=`which python$1`
        mkvirtualenv --distribute --python=$base_python $2
    }

ただどうしても python2.4 がデフォルトの python のままだとうまく行かず, 綺麗な解決法も思いつかなかったので, `/usr/bin/python` へ `/usr/bin/python26` からリンクを貼ることにしました. 敗北...

    $ suro rm /usr/bin/python
    $ suro ln -s /usr/bin/python26 /usr/bin/python

こうすると困るのが yum が動かなくなることで, こちらは別途しらべよう...

    $ yum
    There was a problem importing one of the Python modules
    required to run yum. The error leading to this problem was:
    
       No module named yum
    
    Please install a package which provides this module, or
    verify that the module is installed correctly.
    
    It's possible that the above module doesn't match the
    current version of Python, which is:
    2.6.8 (unknown, Nov  7 2012, 14:47:45)
    [GCC 4.1.2 20080704 (Red Hat 4.1.2-52)]
    
    If you cannot solve this problem yourself, please go to
    the yum faq at:
      http://wiki.linux.duke.edu/YumFaq

ともあれこれで準備はできたので, 仮想環境を作ってみます

    $ mkvenv 2.6 dev
    (dev)$ python -V
    Python 2.6.8

### 参考
[2012.05版 Python開発のお気に入り構成（ポロリもあるよ） - YAMAGUCHI::weblog](http://ymotongpoo.hatenablog.com/entry/20120516/1337123564)
