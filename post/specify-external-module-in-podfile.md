{"title":"Podfile で本家 Spec に無いライブラリを管理する","date":"2014-04-22T21:29:43+09:00","tags":["ios"]}

### github 上にある場合

github の url を `:git` で渡す。Podfile に以下のように記述する。

    pod 'ModuleName', :git => 'https://github.com/foo/ModuleNmae.git'

### ローカルにある場合

ローカルのパスを `:path` で渡す。Podfile に以下のように記述する。

    pod 'ModuleName', :path => '../ModuleName'

### ローカルの podspec を直接指定する場合

podspec ファイルへのパスを `:podspec` で渡す。 Podfile に以下のように記述する。

    pod 'ModuleName', :podspec => '~/path/to/ModuleName.podspec'

### 参考

[CocoaPods Guides - The Podfile](http://guides.cocoapods.org/using/the-podfile.html)
