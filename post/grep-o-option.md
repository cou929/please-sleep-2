{"title":"grep -o -l","date":"2017-08-06T17:44:38+09:00","tags":["nix"]}

`-o` はマッチした部分だけを出力する。もちろん `ag` などにもあるオプション。

<pre><code>-o, --only-matching
        Prints only the matching part of the lines.
</code></pre>

正規表現と組み合わせて、特定のワードが何度出現するか数えたり、特定のパターンにマッチする単語が何種類あるかを調べたりするのに便利。

以下ははてなからエクスポートした movable type 形式のブログ記事ファイルから、`TITLE:` のようなタグの全種類を調べる例。

<pre><code>ag -o '^([A-Z][A-Z ]+):' /path/to/exported_data.txt | sort -u
</code></pre>

`-l` はマッチしたファイル名だけを出力する。該当するファイルの一覧を取得したい場合に便利

<pre><code># 例えば、本日分のログが記録されているファイルをリストアップする
$ sudo ag -l '2017-08-06' /var/log/
/var/log/corecaptured.log
/var/log/system.log
</code></pre>

また反対に `-L` でマッチしないファイル名を出力できる。
