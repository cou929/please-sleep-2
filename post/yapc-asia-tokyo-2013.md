{"title":"YAPC::Asia Tokyo 2013 で発表したことなど","date":"2013-09-25T23:05:26+09:00","tags":["blog"]}

YAPC には去年初めて参加。今年はありがたいことに発表をひとつさせて頂くことができた。

<script async class="speakerdeck-embed" data-id="552f7d1003e60131954556f7ac4f018a" data-ratio="1.2994923857868" src="//speakerdeck.com/assets/embed.js"></script>
[BrowserStack を用いたクライアントサイドのテスト // Speaker Deck](https://speakerdeck.com/cou929/browserstack-woyong-itakuraiantosaidofalsetesuto)

内容は JavaScript をはじめとしたクライアントサイドのテスト自動化について。YAPC の懐の深さに甘えさせていただき、perl 直球ではない、自分の得意に近いテーマを選ばせてもらった。

内容としては [BrowserStack](http://www.browserstack.com/) の紹介だが、その前段としてクライアントサイドのテスト技法を紹介した。特に最近いろいろなフレームワークやライブラリが出てきている分野なので、自分でも流れを追いきれていないところだった。フロントエンドのテストにはどのような難しさがあり、いままではそれをどのように解決してきたのか。今回の発表資料をつくることで、自分でもあいまいだったところをはっきりさせ、理解を整理できたので、個人的にはとても有用だった。

今後は現実の運用にそった知見をためていきたい。たとえば今回の BrowserStack などを使うとかなり多くの環境を対象にできるが、実用上検証すべきはそのうちの一部だけだと思う。どの環境に対してテストをするのか、その対象をどういう手順でアップデートしていけばよいのか。こうした部分はまだよくわかっていない。またより UI 寄りの、エンドツーエンドなテストになっていくにつれ、テストコードの保守コストが高い。少しのコードの変更でもテストコードをたくさん直さないといけない。そもそもどこまでテストを書くのかという部分もまだ曖昧だ。なんでもかんでもテストを書けば良いというものではなく、重点的にやる部分とそうでない部分を分けて考える必要がある。このあたりが今気になっているところだ。

### 気になったトピック: Power Assert

もともと目当てにしていたトークが満席で見られないことも多かったが、たまたま見た中で興味をひかれたのが、gfx さんと tokuhirom さんが LT のなかで取り上げていた Power Assert だ。

[tokuhirom/Test-Power](https://github.com/tokuhirom/Test-Power)

たとえば次のように assert を書くとする (コードは擬似コード)

    a = 10
    b = 9
    assert 91 == a * b

当然この assertion は落ちるのだが、その時にエラーメッセージが次のように詳細に表示される。

    Assertion failed:
    
    assertion 91 == a * b
                 |  | | |
                 |  10| 9
                 |    90
                 false

つまり右辺の変数の内容と計算過程が表示されていて、assert 失敗の理由がとてもわかりやすくなっている。これはとてもうれしい。テストに落ちた際、printf デバッグをしかけながら原因を探る手間が大幅に減るのでテストコードのメンテナンス楽になる。またおそらくほとんどの assertion はこの単純な power assert に置き換えることができ、様々な、フレームワークによって異なる、assert メソッドをいちいち覚えなくてもよくなる。

もともとは Groovy に、しかもなんと 2009 年に!、取り入れられた機能らしい。

[Groovy 1.7 Power Assert \| Don't mind the language](http://dontmindthelanguage.wordpress.com/2009/12/11/groovy-1-7-power-assert/)

最近になって t_wada さんが js の実装を作ったそうだ。

[twada/power-assert](https://github.com/twada/power-assert)

### 感想

今年もいろいろと楽しませてもらった。追えていないトークがたくさんあるので、個人的 YAPC 復習会をやる必要がある。そして今年は現在の形での運営は最後になるらしい。主催の方々をはじめ今まで関わったたくさんの皆様お疲れ様でした。
