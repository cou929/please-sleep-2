{"title":"meta 要素の `http-equiv=\"refresh\"` はページロードが完了した時点から指定した秒数待ってリダイレクトする","date":"2013-04-01T19:10:59+09:00","tags":["html"]}

[4 The elements of HTML — HTML Standard](http://www.whatwg.org/specs/web-apps/current-work/multipage/semantics.html#attr-meta-http-equiv-refresh)

> For the purposes of the previous paragraph, a refresh is said to have come due as soon as the later of the following two conditions occurs:

> * At least time seconds have elapsed since the document has completely loaded, adjusted to take into account user or user agent preferences.

ドキュメントが completely loaded されて, かつ指定した時間経ったあとに遷移するとある.

`completely loaded` の定義はここだけど

[12.2.6 The end — HTML Standard](http://www.whatwg.org/specs/web-apps/current-work/multipage/the-end.html#completely-loaded)

もろもろのステップの一番最後ということのようだ.

リダイレクト時にビーコンを飛ばしたり js を実行させるような場合, `http-equiv="refresh"` を使っても大丈夫そうだ. (もちろん onload を待つような書き方をしている js の実行は微妙そうだけど)
