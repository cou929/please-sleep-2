{"title":"mac の tips 3 つ","date":"2014-03-18T23:24:32+09:00","tags":["mac"]}

最近知った mac の便利ワザ

### sips コマンドで画像編集

`sips` というコマンドがデフォルトで入っていて簡単な画像処理ができる。例えば画像ファイルを 50x50 にリサイズするのはこんなワンライナーで可能。

    for file in *.png; do sips --resampleWidth 50 $file --out ${file%.png}-converted.png; done

こういうちょっとした処理ならいちいち imagemagick をいれなくても、デフォルトのコマンドだけでできる。

### DigitalColor Mater

画面上のあるピクセルの色情報がほしい時に使う。`/Applications/Utilities/` 以下にデフォルトで入っている。

![](/images/digitalcolor.png)

昔は Firefox アドオンの [ColorZilla](https://addons.mozilla.org/en-US/firefox/addon/colorzilla/) をよく使っていたけども、同じようなものがデフォルトで提供されていて驚いた。

### Audio Midi Setup

これも `/Applications/Utilities/` に入っているツールで、音声入出力の調整ができる。自分の場合イヤホンを MacBook Air のイヤホンジャックに直接挿したときにノイズがひどくて困っていたのだが、ここで出力形式を操作してイヤホンの仕様に合うものにしてあげるとノイズが低減できていい感じだった。調整項目の選択肢は多くないので、ノイズに困っているけどオーディオインターフェイスなどは導入しないで手軽に解決したいなーというひとは、てきとうにここをいじってみるといいかもしれない。

![](/images/audio_midi_setup.png)

