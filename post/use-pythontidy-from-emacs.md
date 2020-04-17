{"title":"emacs で PythonTidy の設定","date":"2012-10-05T23:08:06+09:00","tags":["python"]}

[EmacsWiki: Python Programming In Emacs](http://emacswiki.org/emacs/PythonProgrammingInEmacs#toc21)

ここの言うとおりにしただけなんだけど,

1. どこからかスクリプトを持ってくる
  - `git clone https://github.com/witsch/PythonTidy.git` とか
2. パスの通っているところにスクリプトを設置
3. `.emacs` の設定

        ;;; pythontidy
        (defun pytidy-whole-buffer ()
          (interactive)
          (let ((a (point)))
            (shell-command-on-region (point-min) (point-max) "PythonTidy.py" t)
            (goto-char a)))
        (add-hook 'python-mode-hook '(lambda ()
                                       (define-key python-mode-map "\C-ct" 'pytidy-whole-buffer)))

これで python-mode && C-ct で, バッファ全体を PythonTidy にかける. しばらくはこれで運用してみる.

 (このへんの設定はしている前提です)

[emacs の python 開発環境を整える - フリーフォーム フリークアウト](http://d.hatena.ne.jp/cou929_la/20110525/1306321857)
