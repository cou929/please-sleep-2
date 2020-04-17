{"title":"blog をいろいろなおした","date":"2012-12-24T00:23:11+09:00","tags":["memo"]}

### RSS をつけた

[http://please-sleep.cou929.nu/rss.xml](http://please-sleep.cou929.nu/rss.xml)

[PyRSS2Gen](http://www.dalkescientific.com/Python/PyRSS2Gen.html) というモジュールが簡単で, 何も考えずにできました. こんな感じ.

<pre><code data-language="python">for article in articles:
    markdown_string = ''
    with open(article['path'], 'r') as f:
        f.readline()  # remove header
        markdown_string = f.read().decode('utf-8')
    html = markdown.markdown(markdown_string)
    url = host + article['filename'] + '.html'
    items.append(PyRSS2Gen.RSSItem(
            title=article['title'],
            link=url,
            description=html,
            guid=PyRSS2Gen.Guid(url),
            pubDate=datetime.datetime.fromtimestamp(article['issued'])
            ))

rss = PyRSS2Gen.RSS2(
    title='Please Sleep',
    link=host,
    description=u'From notes on my laptop. ブログ未満な作業ログとか',
    lastBuildDate=datetime.datetime.now(),
    items=items
    )

dest_path = os.path.join(self.output_path, self.rss_filename)
rss.write_xml(open(dest_path, 'w'))</code></pre>

### syntax highlight
上のサンプルコードもそうだけど, シンタックスハイライトをつけました. Rainbow というクライアントサイドでやってくれるモジュールを利用.

[Rainbow - Javascript Code Syntax Highlighting](http://craig.is/making/rainbows)

このブログは, 今風に言うと static website generator (ありていに言うと html を吐き出す自前のスクリプト) で作っているので, クライアント側でできることはクライアント側に寄せていきたい. disqus とかこの rainbow のおかげで静的ファイルしか置かなくてもブログっぽいことができて便利だなと思います.

このほかにもスタイルや構成も細かく修正しました.
