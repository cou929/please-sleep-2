{"title":"JavaScript Timer を使ったパフォーマンス改善","date":"2013-04-11T00:00:49+09:00","tags":["javascript"]}

[Yielding with JavaScript Timers - O'Reilly Answers](http://answers.oreilly.com/topic/1506-yielding-with-javascript-timers/) よ読んだメモ。タイマーを使って処理を小分けにし、定期的に処理を UI スレッドに戻してあげる事で、UI にラグが出るというパフォーマンスの低下を防ぐというもの。

### 導入

Web アプリの UI にラグが出てスムーズでない場合は JavaScript の timer を使ったパターンを適用することを検討する価値がある. この記事は Nicholas C. Zakas の High Performance JavaScript からの妙録だ.

処理時間が長くなってしまう js の処理はどうしても存在する. その場合適切なタイミングで js の実行を止め, UI スレッドに処理を戻すようにするとスムーズな UI を実現できる. JavaScript の timer を使ってこれを実装する.

### タイマーの基本

Timer は `setTimeout()`, `setIntervl()` 関数で作成される。この2つは実行する関数と待機時間という共通の引数をとる。それぞれ一度だけ実行するか、定期的に実行するかの違いがある。

タイマーと UI スレッドとのやりとりの仕方は、実行時間が長いスクリプトを短いセグメントに分けることに役立つ。タイマーの関数は JavaScript エンジンにある時間待ったあとに JavaScript のタスクを UI キューに追加する。

<pre><code data-language="javascript">function greeting(){
    alert("Hello world!");
}
setTimeout(greeting, 250);</code></pre>

このコードは 250 ms 後に `greeting()` 関数を UI キューに追加する。その間他の UI アップデートや JavaScript の処理が実行される。250 ms 後に処理が実行されるのではないことに注意。UI キューに追加されたあとはキューの中の他のタスクが実行されたあとに当該タスクの番になる。

### タイマーの精度

タイマーの時間設定は正確ではない。よって時間の計測にはタイマーは使えない。例えば Windows の時間解像度は 15ms だ。

### タイマーによる配列処理

ループは実行時間がながい処理の典型的なケースだ。この場合ループのそれぞれのステップをタイマーにするというパターンを検討する価値がある。次の例を見てみよう。

<pre><code data-language="javascript">for (var i=0, len=items.length; i < len; i++){
    process(items[i]);
}</code></pre>

ループ内のプロセスがとても重いか、あるいは `item` の長さが大きければ処理時間が長くなる。この場合、

- 処理を同期的に行う必要があるか
- データを順に連続して処理する必要があるか

この2つの質問の答えがいずれも No ならばタイマーでこの処理を分けることができる。

<pre><code data-language="javascript">var todo = items.concat(); //create a clone of the original
setTimeout(function(){
    //get next item in the array and process it
    process(todo.shift());
    
    //if there's more items to process, create another timer
    if(todo.length > 0){
        setTimeout(arguments.callee, 25);
    } else {
        callback(items);
    }
}, 25);</code></pre>

基本的なアイデアは、もとの配列をキューとして扱い、ひとつひとつをタイマーで処理するというものだ。

待機時間はユースケースによるが、一般的には最低 25ms にするとよい。それ以下だと UI のアップデートには短すぎる。

このパターンを次のように抽象化できる。

<pre><code data-language="javascript">function processArray(items, process, callback){
    var todo = items.concat(); //create a clone of the original
    setTimeout(function(){
        process(todo.shift());
        if (todo.length > 0){
            setTimeout(arguments.callee, 25);
        } else {
            callback(items);
        }
    }, 25);
}</code></pre>

一般的にこのパターンを適用すると、途中で UI スレッドの処理を行う分総実行時間は伸びる。これはスムーズな UI を実現するためのトレードオフだ。

### タスクのセットアップ

一つの大きな長い処理を細かいサブタスクに分割し、それぞれをタイマー処理する例を考える。

<pre><code data-language="javascript">function saveDocument(id){
        //save the document
        openDocument(id)
        writeText(id);
        closeDocument(id);
        //update the UI to indicate success
        updateUI(id);
}</code></pre>

この関数の中身をそれぞれの処理ごとにタイマーにする。

<pre><code data-language="javascript">function saveDocument(id){
    var tasks = [openDocument, writeText, closeDocument, updateUI];
    setTimeout(function(){

        //execute the next task
        var task = tasks.shift();
        task(id);

        //determine if there's more
        if (tasks.length > 0){
            setTimeout(arguments.callee, 25);
        }
    }, 25);
}</code></pre>

サブタスクを配列にまとめ順にタイマーで実行している。この処理を抽象化すると次になる。

<pre><code data-language="javascript">function multistep(steps, args, callback){
    var tasks = steps.concat(); //clone the array
        setTimeout(function(){
        
        //execute the next task
        var task = tasks.shift();
        task.apply(null, args || []);
        
        //determine if there's more
        if (tasks.length > 0){
            setTimeout(arguments.callee, 25);
        } else {
            callback();
        }
    }, 25);
}</code></pre>

### 時間別のコード

一度にひとつの処理しかしないのはしばしば非効率だ。例えば 1 処理に 1ms かかるものをループで 1000 回実行する場合、これを 25ms のタイマーで分割すると約 26sec かかってしまう。たとえば 50 回連続でしょりしてからタイマーをはさむと総実行時間は 1.5sec に抑えられる。

たとえば連続して実行できる JavaScript の時間は 100ms 以内とポリシーを決めて、 100ms ごとにタイマーをはさむようにこれまでの処理を改良してみる。

<pre><code data-language="javascript">function timedProcessArray(items, process, callback){
    var todo = items.concat(); //create a clone of the original

    setTimeout(function(){
        var start = +new Date();

        do {
            process(todo.shift());
        } while (todo.length > 0 && (+new Date() - start < 50));

        if (todo.length > 0){
            setTimeout(arguments.callee, 25);
        } else {
            callback(items);
        }
    }, 25);
}</code></pre>

### タイマーとパフォーマンス

タイマーを使うことでパフォーマンスを改善できるが、使い過ぎると別の問題になることもある。今回の例はある処理が完了したあとに別のタイマーを仕込むものだが、定期定期にタイマーを設定するタイプの処理はパフォーマンスが低下しやすい。Google Mobile の Neil Thomas の調査によると、

[Gmail for Mobile HTML5 Series: Using Timers Effectively - The official Google Code blog](http://googlecode.blogspot.jp/2009/07/gmail-for-mobile-html5-series-using.html)

1秒以上の長い粒度でのタイマー設定は大丈夫だが 100ms 程度の短い粒度でタイマーをセットするようなコードではパフォーマンスの劣化が見られたという。
