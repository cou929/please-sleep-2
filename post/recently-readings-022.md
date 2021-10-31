{"title":"最近読んだもの 22","date":"2021-11-01T00:30:00+09:00","tags":["readings"]}

## 記事

- [Cache \(computing\) \- Wikipedia](https://en.wikipedia.org/wiki/Cache_(computing))
	- かなりいい記事だった。キャッシュについて網羅的に、よく一般化され整理された説明
	- 特に書き込みのポリシーはこれまで経験が少なく面白かった
	- Write-through or Write-back x Write allocate or No-write allocate という基本的な要素に分解され、その組み合わせとそれぞれの説明
- [Fuzzing vs property testing](https://www.tedinski.com/2018/12/11/fuzzing-and-property-testing.html)
	- Fuzzing と Property testing (Property based testing) の違いについて
	- 前者のほうがよりブラックボックス・入出力に関する事前知識を不要とする・時間がかかるもの
	- 対して pbt はある程度入出力の条件を人間が指定することで、fuzzing より高速に、example based testing よりは網羅的にテストできる
		- (ここでは通常のテストケースを `Example based` と呼ぶ)
	- という理解であっているかな?
	- あとは不変条件の表明が PBT と相性が良い (相互補完的に良くしていける) というような記述もなるほどだった
- [Our experience with Hypothesis testing \(and why do we love it so much \!\) \- Parsec](https://parsec.cloud/our_experience_with_hypothesis_testing/)
	- statefull な PBT というもがあって、これが本命という感がある
	- いろいろな副作用のある操作を登録しておくと、それらをいろいろな呼び出し方で呼び出してくれる
	- あとは shrink という仕組みがかっこいい
		- fail するケースを見つけたら、条件を削っていって最小の再現条件を探しに行く
	- 試してみたい
- [Pull Request Merge Queue Limited Beta \| GitHub Changelog](https://github.blog/changelog/2021-10-27-pull-request-merge-queue-limited-beta/)
	- 参加人数・チーム数が多いリポジトリの場合、ベースブランチの進みが早いので、プリリクエストを作成・レビュー・CI しているうちにすぐコンフリクトする
	- プルリクエストが ready になったらマージキューに入れて、CI が落ちなければマージされるという機能
	- そのような大人数のプロジェクトでは生産性にかなり貢献しそう
- [Command palette beta \| GitHub Changelog](https://github.blog/changelog/2021-10-27-command-palette-beta/)
	- cmd + k でコマンドパレットが出るようになった。これは便利
- [How Pokémon GO scales to millions of requests? \| Google Cloud Blog](https://cloud.google.com/blog/topics/developers-practitioners/how-pok%C3%A9mon-go-scales-millions-requests)
	- 規模がすごいことはわかったけど、それ以外は何もわからなかった
- [To Learn a New Language, Read Its Standard Library \- Pat Shaughnessy](http://patshaughnessy.net/2021/10/23/to-learn-a-new-language-read-its-standard-library)
	- 新しい言語を学ぶには標準ライブラリを読むといいよとのこと
	- この人が最近読んだ Crystal の Array の実装が読みやすいというだけの話な気がした。`like reading a fairy tale in a children’s book` に読めるものはそんなに多くはないのでは
- [Jaana Dogan ヤナ ドガン on Twitter: "When designing \#golang APIs, I always turn off my syntax highlighting\. In order for me to be comfortable with design, it must be readable without syntax highlighting\." / Twitter](https://mobile.twitter.com/rakyll/status/1453818247501934603)
	- Go で API 設計をする際にあえてシンタックスハイライトを切っているらしい
	- ハイライトなしでも読みやすくあるべきなので、とのこと
	- この発想はなかったが、Go なら確かにハイライトなしでも読みやすそう
- [Chris's Wiki :: blog/programming/GoVersionOfYourSource](https://utcc.utoronto.ca/~cks/space/blog/programming/GoVersionOfYourSource)
	- Go1.8 からビルド時のリビジョン（と 変更や untracked file があるかどうか）をバイナリに埋め込んでくれるらしい
	- こういう運用的なベストプラクティスをコアが提供するのが Go ぽい
	- Ruby や Rails が開発時の体験や効率を大事にしているのに対して、Go は運用時にフォーカスしている感がある。個人的には共感することが多い
