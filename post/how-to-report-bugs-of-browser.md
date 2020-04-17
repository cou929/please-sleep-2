{"title":"ブラウザのバグを見つけたときにやること","date":"2012-12-17T12:02:50+09:00","tags":["memo"]}

<img src="https://raw.github.com/paulirish/browser-logos/master/all-desktop.png" alt="five browsers logo" style="max-width:100%;">
[paulirish/browser-logos · GitHub](https://github.com/paulirish/browser-logos/)

1. バグを見つける
   - 表示系のものならスクリーンショットをとっておく
2. すでに報告されていないか調べる
   - ブラウザごとのバグトラッカーで検索してみる
   - サポートフォーラムなどコミュニティに相談してみる
3. ブラウザの種類とバージョンの絞込み
   - 他のブラウザでは発生するか. 他のバージョンでも発生するか (stable を使っているなら beta や nightly でも見てみる等)
4. テストケースの最小化
   - バグを再現させる最小のテストケースを作ります
   - 今回出会ったのは表示系のバグだったので, html / css を削りながら現象を再現させていき, 同じ現象が起こる最小の html を作りました
5. バグの報告先とフォーマットの確認
   - ブラウザごとにバグ報告のガイドラインがあるはずなので, それを読めばどのように報告すればよいかわかるはずです
   - chromium は選択肢を選ぶといい感じに報告してくれるウィザードを用意していたりします
   - chrome / safari のバグの場合, どちらか特有のものなのか or いずれも再現するかに応じて, 報告先が chromium / webkit と変わります.
6. レポートを書く
   - バグを発生させる操作手順
   - 発生する事象
   - 期待している動作
   - 上記を再現する最小のテストケース
   - 必要に応じてスクリーンショット
   - その他ガイドラインに応じて
7. 報告する

### 主要ブラウザの報告ガイドラインと報告先

- chromium (chrome)
  - [Bug Life Cycle and Reporting Guidelines - The Chromium Projects](http://www.chromium.org/for-testers/bug-reporting-guidelines)
  - [Bug Reporting Guidlines for the Mac & Linux builds - The Chromium Projects](http://www.chromium.org/for-testers/bug-reporting-guidlines-for-the-mac-linux-builds)
  - [Issues - chromium - An open-source browser project to help move the web forward. - Google Project Hosting](https://code.google.com/p/chromium/issues/list)
- webkit (safari)
  - [The WebKit Open Source Project - Reporting Bugs](http://www.webkit.org/quality/reporting.html)
  - [The WebKit Open Source Project - Bug Reporting Guidelines](http://www.webkit.org/quality/bugwriting.html)
  - [WebKit Bugzilla](https://bugs.webkit.org/)
- firefox
  - [Bug writing guidelines \| MDN](https://developer.mozilla.org/en/docs/Bug_writing_guidelines)
  - [Bugzilla@Mozilla – Main Page](https://bugzilla.mozilla.org/)
- opera
  - [Guidelines for Filing Good Bug Reports](http://www.opera.com/support/bugs/guidelines/)
  - [Opera: Support - Reporting bugs](http://www.opera.com/support/bugs/)
  - [Opera Bug Report Wizard](https://bugs.opera.com/wizard/)
  - [System Dashboard - Opera Bug Tracking System](https://bugs.opera.com/secure/Dashboard.jspa)
- ie
  - [Announcing Internet Explorer Feedback - IEBlog - Site Home - MSDN Blogs](http://blogs.msdn.com/b/ie/archive/2006/03/24/announcing-internet-explorer-feedback.aspx)

### 各ブラウザのリリースチャンネル

- chrome
  - [Google Chrome - Get a fast new browser. For PC, Mac, and Linux](https://tools.google.com/dlpage/chromesxs/)
  - [Chrome Release Channels - The Chromium Projects](http://www.chromium.org/getting-involved/dev-channel)
- safari
  - [WebKit Nightly Builds](http://nightly.webkit.org/)
- firefox
  - [Firefox Nightly Builds](http://nightly.mozilla.org/)
  - [Mozilla Firefox Web Browser — Download Firefox Aurora or Beta & Help Determine the Next Firefox — mozilla.org](http://www.mozilla.org/en-US/firefox/channel/)
- opera
  - [Opera web browser \| Opera Next](http://www.opera.com/browser/next/)
- ie
  - [Download IE10 Platform Preview](http://ie.microsoft.com/testdrive/Info/Downloads/Default.html)

普段から開発チャンネルを使ってデバッグに協力するのはよい習慣だと思います. 異なるバージョンの共存はこちらが参考になります (Mac の場合)

[MacをJavaScriptの開発環境にするメモ - os0x.blog](http://os0x.hatenablog.com/entry/20110101/1293831128)

### 今回報告したバグ

[Bug 102402 – REGRESSION (Subpixel layout): Gray vertical lines appear when moving insertion point to right on Mac](https://bugs.webkit.org/show_bug.cgi?id=102402)
