{"title":"最近読んだもの 17","date":"2021-09-26T22:30:00+09:00","tags":["readings"]}

## 記事

- [Taming Rails memory bloat \| Mike Perham](https://www.mikeperham.com/2018/04/25/taming-rails-memory-bloat/)
	- linux での glibc のメモリアロケーションと mri でのマルチスレッド環境は相性が悪いらしく、メモリのプラグメンテーションで使用量が上がる問題があるらしい
	- 緩和するオプションや、jemalloc での改善の紹介
- [Troubleshoot Clusters \| Kubernetes](https://kubernetes.io/docs/tasks/debug-application-cluster/debug-cluster/)
	- クラスタのデバッグをする際の overview
	- 一つ一つは当たり前のことなんだけど理解の整理に
- [Tammy Bryant Butow on SRE Apprentices](https://www.infoq.com/podcasts/sre-apprentices/)
	-  SRE 養成プログラム
	- プログラムの内容が見られるの別の記事だった
	- Apprenticeship は日本だと新卒ぽいニュアンスなんだろうか
	- これを進めることでチームのダイバーシティも改善したというのも興味深い
- [Thoughtfully Training SRE Apprentices: Establishing Padawan and Jedi Matches](https://www.infoq.com/articles/training-sre-apprentices/)
	- 上記の関連記事
	- より詳細に説明しているが、ほとんどが技術的な側面よりも人間的（社会的）側面の内容
	- 正しいし示唆的だが、今は技術的なものを読みたかった
	- 内容自体は良いもの
- [Section 230 and a Tragedy of the Commons \| October 2021 \| Communications of the ACM](https://cacm.acm.org/magazines/2021/10/255705-section-230-and-a-tragedy-of-the-commons/fulltext)
	- これは面白かった。アメリカの分断がソーシャルメディアによって強化されている構造の解説と、法によってどう対策していくか
	- ミスリードやデマ情報によって、アメリカは議会が占拠されたり、死亡者が出るまでになっている。その後トランプのアカウントはバンされたが、あまりに遅い
	- デマでもソーシャルメディアにとっては収益であること、デマの収益性をうわまわる罰則を今の法律では与えられないこと、ソーシャルメディアは原則中立でけんえつを避けたがることなど、現状に収束しているのには構造的な理由がある。事実ザッカーバーグもこの問題への対処は fb 一企業の努力では難しいと認めている
	- デマ情報の拡散に対して、プラットフォーム企業に責任を求め、より強い罰則で規制すること。プラットフォーム企業は AI と人力両面からミスリードやデマ情報の対策を進めることが必要と述べている
	- それに加えて、これらが成されたとしても本質的な問題は解消しないと述べているのか印象的だった。筆者はデマに踊らされる人を減らすには教育、特に科学、歴史、論理が必要だと考えていた。しかし現在高い教育を受けた人が反ワクチンを信じている例が多くあり、教育だけでは不十分だと認識を改めた。ではどうすればいいかは記載がない
	- 個人的には、これらの対策を進めるとして、どのようにミスリードやデマ情報を判定するのかがかなり難しい気がしたが、そこへの言及はなかった。もう研究が進んでいたりするのかな
- [Divide and Conquer \| October 2021 \| Communications of the ACM](https://cacm.acm.org/magazines/2021/10/255709-divide-and-conquer/fulltext)
	- タイミングバグ、heisenbug、分散システムのバグ、過去のコード変更に依存しないバグは git bisect で検出できないし、誤った結果によりミスリードにつながるという話
- [How to Reduce Memory Bloat in Ruby \| AppSignal Blog](https://blog.appsignal.com/2021/09/21/how-to-reduce-memory-bloat-in-ruby.html)
	- Ruby で Memory Bloat と呼ばれている現象の概説
	- メモリリークみたいなものだが、もっと急激なのでこう呼ばれているらしい
	- 原因は大きく二つで、メモリのフラグメンテーションとメモリ解放の遅れ
	- 前者は ruby に限らず一般的な話だが、後者は特有かも？
	- 対処法
		- フラグメンテーションに対しては、 jemalloc という別のアロケータを使うという物以外に、メモリのマネジメントをアプリ実装時に意識せよというものだった。Ruby にはそういう明示的に確保したり開放する api があるのかな？
		- 開放の遅れは、処理系のコードにパッチを当てる方法しか紹介されてなくて、なかなか受け入れがたく、謎だった
	- 全体として簡潔な説明で概要の理解には良い記事
- [New Google Cloud Deploy automates deploys to GKE \| Google Cloud Blog](https://cloud.google.com/blog/products/devops-sre/google-cloud-deploy-automates-deploys-to-gke)
	- GCP のマネージドの CD サービスらしい
- [Move over Rake, Thor is the new King](https://technology.doximity.com/articles/move-over-rake-thor-is-the-new-king)
	- rake の代替となる thor の紹介
	- まともなオプションパーサーが付いてるのは良いかも。そんなに複雑なことをすることがどれだけあるかは分からないが
