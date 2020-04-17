{"title":"各 shell での history コマンドで直近 n のコマンドだけを表示","date":"2012-12-09T18:49:55+09:00","tags":["nix"]}

- bash
  - ない. HISTTIMEFORMAT を考慮したパースが必要
  - `history 5` で件数絞込みだけ行う
- csh, tcsh
  - `history -h 5`
- ksh
  - `hist -ln -5 | tr -d '\t'`
  - 先頭にスペースが入るようだ
- zsh
  - `fc -ln -5`
  - 自分自身は含まれない
