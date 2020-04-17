{"title":"emacs の grep モード","date":"2012-12-19T21:56:28+09:00","tags":["emacs"]}

なんで今まで使っていなかったのか.

grep してマッチしたファイルの該当行を開くということができる. もちろん syntax highlight も. 開発環境ならばもう ack はいらないかもしれない

### 使い方

- デフォルトでは `M-x grep` でミニバッファに grep コマンドがでるので入力, 実行すると新しいバッファに結果が出る.
  - grep コマンドの任意のオプションをその場でつけることもできる. grep コマンドそのものなのでわかりやすい
- 見たいファイルにカーソルを合わせて Return でそのファイルの該当行を別バッファで開く
- `M-g n` で次の候補, `M-g p` で前の候補に移動

### カスタマイズ

デフォルトでは `grep -nh -e` というオプションがついているが, `-r` もつけたい. またキーバインドも設定したい.

[Emacs実践入門 - おすすめEmacs設定2012 - ククログ(2012-03-20)](http://www.clear-code.com/blog/2012/3/20.html)

こちらのクリアコードさんの記事を参考に以下を設定しました.

<pre><code data-language="scheme">;;; grep
(define-key global-map (kbd "C-x g") 'grep)
(require 'grep)
(setq grep-command-before-query "grep -nH -r -e ")
(defun grep-default-command ()
  (if current-prefix-arg
      (let ((grep-command-before-target
             (concat grep-command-before-query
                     (shell-quote-argument (grep-tag-default)))))
        (cons (if buffer-file-name
                  (concat grep-command-before-target
                          " *."
                          (file-name-extension buffer-file-name))
                (concat grep-command-before-target " ."))
              (+ (length grep-command-before-target) 1)))
    (car grep-command)))
(setq grep-command (cons (concat grep-command-before-query " .")
                         (+ (length grep-command-before-query) 1)))
</code></pre>
