{"title":"CKAD を取得した","date":"2021-07-10T16:00:00+09:00","tags":["k8s"]}

<div data-iframe-width="150" data-iframe-height="270" data-share-badge-id="e3862d34-4b5b-4eb8-972e-c11b56343f84" data-share-badge-host="https://www.credly.com"></div><script type="text/javascript" async src="//cdn.credly.com/assets/utilities/embed.js"></script>

<div></div>

[CKAD: Certified Kubernetes Application Developer \- Credly](https://www.credly.com/badges/e3862d34-4b5b-4eb8-972e-c11b56343f84/public_url)

## 試験対策

Udemy の講義が安くなっていたタイミングで買ってあったので、こちらを一通り見た。

[Kubernetes Certified Application Developer \(CKAD\) Training \| Udemy](https://www.udemy.com/course/certified-kubernetes-application-developer/)

あとは本家 Linux Fundation のサイトで受験料を支払うと、受験 Tips (`alias k=kubectl` するといいよ、みたいな細かいところまであった) とか、練習問題も 2 セッションついてきたので、前日に一通りやった。

[Certified Kubernetes Application Developer \(CKAD\) \- Linux Foundation \- Training](https://training.linuxfoundation.org/certification/certified-kubernetes-application-developer-ckad/)

Udemy にも練習問題はついているんだけど、試験環境 (kakakoda ベースの古い方?) の遅延がひどくてかなりストレスだった。本家の方 ([killer.sh](https://killer.sh/) ベース) は動作も軽快だし UI も本番と近いのでこちらだけで良かった。

一応 [dgkanatsios/CKAD\-exercises: A set of exercises to prepare for Certified Kubernetes Application Developer exam by Cloud Native Computing Foundation](https://github.com/dgkanatsios/CKAD-exercises) も流し見したけど Udemy と本家の練習問題だけやれば十分だったと思う。

基本的には Web 上のリソースだけを使っていたが、ピンポイントでこの本も参照した。

<div class="amazlet-box" style="margin-bottom:0px;"><div class="amazlet-image" style="float:left;margin:0px 12px 1px 0px;"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4295009792/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank"><img src="https://images-na.ssl-images-amazon.com/images/I/510to2Qh1FL._SX390_BO1,204,203,200_.jpg" alt="Kubernetes完全ガイド 第2版 (Top Gear)" style="border: none; width: 113px;" /></a></div><div class="amazlet-info" style="line-height:120%; margin-bottom: 10px"><div class="amazlet-name" style="margin-bottom:10px;line-height:120%"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4295009792/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Kubernetes完全ガイド 第2版 (Top Gear)</a></div><div class="amazlet-detail">青山 真也  (著)<br/></div><div class="amazlet-sub-info" style="float: left;"><div class="amazlet-link" style="margin-top: 5px"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4295009792/pleasesleep-22/ref=nosim/" name="amazletlink" target="_blank">Amazon.co.jpで詳細を見る</a></div></div></div><div class="amazlet-footer" style="clear: left"></div></div>

## 各種料金

Udemy の講義は定価 2 万円ほどだけど、頻繁に 9 割以上割り引きしているので、個人で買うならタイミングを見計らったほうが良い。

試験の料金は 2021 年の 7 月から値上げしたらしく、Exam only で $375 だった。自分はちょうど値上げ直後でタイミングが悪かった。

[Certification Exam Prices Increase July 1 \- Lock in Current Pricing \- Linux Foundation \- Training](https://training.linuxfoundation.org/announcements/certification-exam-prices-increase-july-1-lock-in-current-pricing/)

例年だと Cyber monday や Black friday にセールをしていたり、kubecon などのカンファレンスに参加すると割り引きプロモーションコードがもらえたりするらしいので、急いでない人はこちらもタイミングを見計らったほうが良いと思う。

あとは毎年 Q3 に試験内容を更新しているようなので、そういう意味でもこの 7 月の受験はちょっとタイミングが悪かったかもしれない。

[CKAD Program Changes: 2021 \- Linux Foundation \- Training](https://training.linuxfoundation.org/ckad-program-change-2021/)

## 雑感

4 月から業務で k8s を使うようになって、必要なところだけを落下傘的に勉強していた。多少業務で触り始めたこのタイミングで体系的に学べたのはとても良かった。あまり資格をとったりした経験がない (初めてかもしれない) が、Udemy での講義形式での勉強と試験対策の練習問題をたくさんやることが、思っていたよりも理解の定着につながった感覚がある。効率よく学習できて良かったと思う。

一方で CKAD は問題数が多いので [Imperative commands](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/imperative-command/) をかなり駆使することになる。実業務でそんなに使わないような気がするのはちょっと気になった。

次は CKA, CKS もやってみようと思っている (Udemy の割り引きキャンペーンや受験料割引キャンペーンを待ちつつ)。
