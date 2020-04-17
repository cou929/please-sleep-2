{"title":"新しい mac を買った時にやったこと","date":"2013-03-24T21:56:33+09:00","tags":["mac"]}

教訓: Time Machine を導入しよう

### System Preference

- tap to click
- 3 finger drag
- caps lock to control
- spotlight shortcut to ^ + command + i
- remove icons from menubar
- move dock to left and hide automatically
- remove icons from dock
- make fastest key repeat speed and delay until repeat - mission control > hot corners > set desktop to left down
- shorten computer name in sharing > computer name

### apps

- chrome canary
  - setting as default browser (from safari)
- chrome
- firefox nightly
- firefox
- opera
- google japanese input
- iterm2
  - set font menlo, 14px
  - set window transparent
- Skype
- dropbox

### installs

- Xcode command line tools (https://developer.apple.com/downloads/index.action#)
- install Xcode from mac app store
- homebrew

        ruby -e "$(curl -fsSL https://raw.github.com/mxcl/homebrew/go)"
        brew install git
        brew update

- brew install emacs
- dev screen

        brew install autoconf
        brew install automake
        git clone git://git.savannah.gnu.org/screen.git
        cd screen/src
        ./configure --enable-pam --enable-colors256 --enable-rxvt_osc --enable-use-locale --enable-telnet
        make
        sudo make install

- sh -x Dropbox/resource/dotfiles/bootsstrap.sh
- set zsh as login shell (chsh)
- brew install tig
- brew install ack
- zsh-syntax-higlighting

        mkdir ~/src
        cd src
        git clone https://github.com/zsh-users/zsh-syntax-highlighting.git

- python

        sudo easy_install virtualenv
        sudo easy_install virtualenvwrapper
        mkvenv 2.7 dev

- perl

        curl -kL http://install.perlbrew.pl | bash
        source ~/perl5/perlbrew/etc/bashrc
        perlbrew install -n perl-5.16.2
        perlbrew use 5.16.2
        perlbrew install-cpanm
        cpanm MIYAGAWA/carton-v0.9_7.tar.gz

- node

        curl -L git.io/nodebrew | perl - setup
        export PATH=$HOME/.nodebrew/current/bin:$PATH  # also set this line to rc file
        nodebrew install v0.10.1
