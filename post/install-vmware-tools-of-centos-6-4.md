{"title":"VMWare Fusion 上の CentOS6.4 に VMWare Tools をインストール","date":"2013-11-01T19:47:32+09:00","tags":["nix"]}

VMWare Tools というツールがある。VMWare のゲスト OS に直接インストールするツール群で、パフォーマンスの向上やホスト OS とのもろもろの同期など、いろいろと便利な機能を提供してくれる。そのなかのひとつにホスト OS との時刻の同期がある。ゲスト OS の時刻は自然とずれて、いろろと対策するもなかなかうまくいってなかったが、自分の場合 VMWare Tools での同期でうまくいった。

### 手順

- 対象のゲスト OS を起動し、VMWare Fusion のメニューから "仮想マシン" -> "VMware Toolsのインストール" を洗濯する
  - もし VM 上の仮想 CD/DVD ドライブがオフになっていた場合はここで警告がでる。その場合はゲスト OS を一度シャットダウンし、VMWare 上の設定項目から CD/DVD ドライブを有効化すればよい
- ダイアログがでるのでインストールを選択すると、ダウンロードがはじまる
- ダウンロードが完了したあとはゲスト OS でのプロンプト作業
- CD ドライブをマウントする

        $ mkdir /mnt/cdrom
        $ mount /dev/cdrom /mnt/cdrom

- VMWare Tools の tar ball があるので、コピー。ドライブはアンマウントしておく

        $ cp /mnt/cdrom/VMwareTools-XXXX.tar.gz /tmp/VMwareTools.tar.gz
        $ umount /mnt/cdrom/

- 展開してインストール

        $ cd /tmp/
        $ tar -zxvf VMwareTools.tar.gz
        $ cd vmware-tools-distrib/
        $ ./vmware-install.pl

  - vmware-install.pl からプロンプトでいろいろと尋ねられるが、すべてデフォルトでよいはず
- インストールの確認とリブート

        $ /usr/bin/vmware-toolbox-cmd -v
        $ status vmware-tools
        $ vmware-tools start/running
        $ shutdown -r now

### 参考

- [CentOS6.2にVMware Toolsをインストールする \| mawatari.jp](http://mawatari.jp/archives/centos-6-2-vmware-tools-install-log)
