{"title":"perl モジュールのアップデート方法","date":"2013-10-20T17:40:18+09:00","tags":["perl"]}

- cpanm 自身は `cpanm --self-upgrade`
  - perlbrew の場合は `perlbrew install-cpanm` など
- cpanm でインストールしているモジュールは `cpanm <module>` とするとその時点の最新版をインストール
  - `cpanm Carton`
- carton で管理しているモジュールは
  - `carton update <module>`
  - cpanfile に必要なバージョン定義を書いて `carton install`
