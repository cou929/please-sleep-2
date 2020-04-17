{"title":"iOS7 で 3rd party domain での localStorage の挙動","date":"2013-12-04T22:07:18+09:00","tags":["html"]}

iOS7 の Safari にて。サードパーティドメインの iframe 内で、そのドメインの localStorage にデータを保存しても、Safari の再起動でそのデータが削除されているようにみえる。Safari の再起動とは、ホームボタンをダブルタップしてアプリのプロセスを殺して再起動するという作業だ。タブを開閉したり、プロセスは継続したままたホームボタンを一回タップして終了・起動を繰り返した場合は当然データが保持されている。よってセッションストレージよりはデータの寿命が長いが、しかしプロセスの終了とともにデータが揮発するという、よくわからない現象だ。

サンプルが少ないので確かではないが、以下の条件で発生するようだった。

- iOS 7.0.2、7.0.4 で発生した
  - iOS7 未満では発生しなかった
- ハードのバージョンは iPhone4S、iPhone5 で発生した
  - <del>iPhone5S では発生しなかった</del> (正確でない可能性あり。追記参照)
- いずれも Safari のプライバシー設定はデフォルト。Cookie はサードパーティのみを拒否、DNT はオフ。
- プライベートブラウズモードではなく、通常モード。

端末がなく検証できないため確実ではないが、iOS6 以前のデフォルト設定 (サードパーティクッキー拒否) の Safari では、今回のようにサードパーティドメインの iframe 内でも、そのドメインの localStorage にデータを保存しできるし、そのデータが永続化されていたように思う。サードパーティクッキーの書き込みを拒否しているので、今回の挙動は本来あるべき姿としては、ある程度納得のいくものではなる。とはいえ、機種によって挙動が違うように見えたし、またどうせやるならクッキー同様サードパーティでの書き込みは拒否してしまうか、あるいは個別にポリシーを設定できるようにすべき? と思ってしまう。

ポリシー変更なのか、バグなのかは定かではないが、なによりこの現象に関する情報が Web 上で見当たらなかったため自信がない。iOS7 & localStorage といえば、当初はバグでプライベートブラウズがデフォルトでオンになってしまっていたというものがあるようだが、今回のケースには該当しない。

[iOS7搭載iPhoneでアプリやブラウザの自動ログインができない症状 バックグランド更新のバグが原因？｜スマートフォン不具合速報](http://blog.livedoor.jp/sumahoreview/archives/33460892.html)

#### 追記 (2013-12-03)

PC の Safari 7.0 (9537.71) でも再現することを確認。よって Safari のあるバージョンからサードパーティの WebStorage に対するポリシーが変更されたと考えてよさそうだ。またこのことから、前述の "iPhone5S では発生しなかった" は間違いかもしれない。機種を問わず Safari のバージョンで挙動が変わっていると考えたほうが自然。

さらに、こちらで WebStorage の仕様をみたところ以下の記述があった。

[Web Storage の仕様を読む - Please Sleep](http://please-sleep.cou929.nu/reading-webstorage-w3c-spec.html)

> Expiring stored data
>
> User agents may, if so configured by the user, automatically delete stored data after a period of time.
>
> For example, a user agent could be configured to treat third-party local storage areas as session-only storage, deleting the data once the user had closed all the browsing contexts that could access it.

あくまで "MAY" での規定のなかでの実装例としてあげられているものでしかないが、"ユーザーエージェントはサードパーティの localStorage をセッションオンリーのストレージとして扱う" とあり、今回の挙動とぴったり一致している。

ちなみに workaround として Web SQL Database や (Safari は対応していないが) Indexed Database も考えられるが、いずれも仕様の Privacy のセクションに WebStorage と同様の記述があった。普通に考えてこれらのプライバシーポリシーをストレージごとに別にするとは思えないので、このへんの別ストレージ技術を使っての workaround は無理そうだ。

- [Web SQL Database](http://www.w3.org/TR/webdatabase/#user-tracking)
- [Indexed Database API](http://www.w3.org/TR/IndexedDB/#privacy)

 (追記ここまで)

#### 追記 (2013-12-04)

webkit への以下のコミットから、このような挙動に変更されたようだ。

[Bug 115004 – Change approach to third-party blocking for LocalStorage](https://bugs.webkit.org/show_bug.cgi?id=115004)

詳細は [webkint ではサードパーティドメインの localStorage が sessionStorage になる - Please Sleep](http://please-sleep.cou929.nu/webkit-storage-blocking-policy.html)

 (追記ここまで)

### 検証コード

localStorage から get、なかったら現在時刻と乱数を set し、もう一度 get、その結果を表示するスクリプト。これをファーストパーティの script タグ内と、同様の動作をする html を別ドメインの iframe に読み込む。iframe に読み込む html は dropbox から配信することにする。ファーストパーティのドメインは "please-sleep.cou929.nu"、そこから呼び出されるサードパーティの iframe のドメインは "dl.dropboxusercontent.com" となる。通常の PC のブラウザなどでは、初回は get できず set を行い、以降のアクセスではそれぞれのドメインの localStorage に保存してあるデータが表示される。iOS7 の iPhone Mobile Safari でこのページを開き、何度かリロードし 1st party / 3rd party それぞれの localStorage にデータが保存されていることを確認。その後 Safari のプロセスを殺し、再度立ち上げるとサードパーティの (iframe 内の) スクリプトは localStorage からデータを get できないという現象が起こる。

<div id="result"></div>
<script>
  (function() {
      var key = 'localstorage_test_item',
          result_field = document.getElementById('result');

      if (!window.localStorage) {
          result_field.innerHTML = 'localStorage is not supported';
          return;
      }

      var got_first = window.localStorage.getItem(key);
      if (!got_first) {
          value = [(new Date()).toString(), Math.floor(Math.random() * 10e12)].join('\t');
          window.localStorage.setItem(key, value);
      }
      var got_second = window.localStorage.getItem(key);

      result_field.innerHTML = [
          'first get:  ' + (got_first || ''),
          'second get: ' + (got_second || ''),
      ].join('<br/>')
  })();
</script>
<iframe src="https://dl.dropboxusercontent.com/u/151946/localstoragesample/sample_iframe.html"></iframe>

コードの内容はこちら。

<script src="https://gist.github.com/cou929/7734384.js"></script>

