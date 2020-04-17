{"title":"xcode7 を入れると `node-gyp` でエラーが出る","date":"2015-09-24T18:00:00+09:00","tags":["nix"]}

xcode7 に更新すると、おそらく gcc まわりの更新の影響で、`node-gyp` まわりのエラーがでるようになった。その結果 atom の一部のプラグインが動かなくなった (rebuild に失敗する)。具体的には [ever-notedown](https://atom.io/packages/ever-notedown)。

結論としては node をビルドし直すと解決した。

- gcc まわりの更新によるものっぽいので、とりあえず node をソースからビルドし直すと大丈夫だった
	- `nvm` を使っているので、`nvm install -s <version>`
	- `nvm install` で落ちてくるバイナリは非対応だったので、ソースからビルドしなおした
- 関連ありそうな issue はこれ?
	- [XCode 7: library not found for -lgcc_s.10.5 · Issue #2933 · nodejs/node](https://github.com/nodejs/node/issues/2933)
	- 共有ライブラリのシンボリックリンクをいじるという、怪しげなワークアラウンドが紹介されている

あまり騒がれている話を聞かないんだけど、自分の環境だけだろうか… (そんなに特殊なつもりはないのだけれど)
