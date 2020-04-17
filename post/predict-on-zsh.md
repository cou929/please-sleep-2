{"title":"zsh の前方予測補完","date":"2012-12-09T18:50:39+09:00","tags":["nix"]}

zsh に predict-on というオプションがあって, インクリメンタルサーチのように入力した文字にマッチする候補を保管してくれるようになる. 便利そうなのでしばらく使ってみよう.

    autoload predict-on
    predict-on
