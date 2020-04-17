{"title":"\"Crash Course: Design for Startups\" Memo","date":"2012-04-21T12:06:48+09:00","tags":["design"]}

[Crash Course: Design for Startups — PaulStamatiou.com](http://paulstamatiou.com/startup-web-design-ux-crash-course) を読んだメモ. 翻訳ではなく自分に必要なところだけメモっただけ.

[Paul Stamatiou](https://twitter.com/Stammy) は [Notifo](http://paulstamatiou.com/startup-update-notifo-iphone-app-v2-more), や [Pic A Fight](http://paulstamatiou.com/pic-a-fight-launch-viral-facemash-instagram) のデザイナ. Web デザインについてのスタートアップガイド.

- Subtle is key! Except when it's not.
  - 覚えたてのテクニックを使いたくなりがち
    - text-shadow や border-radius やグラディエントや
  - 不必要に細部に凝りすぎるのは危険

- Get Inspired and Stay Thirsty My Friends Organized
  - いいデザインはクリップしておく
  - [LittleSnapper](http://www.realmacsoftware.com/littlesnapper/) を使っている

- Process
  - ターゲットオーディエンスの定義
    - [The Five W's of UX - 52 Weeks of UX](http://52weeksofux.com/post/890288783/the-five-ws-of-ux)
  - 上記で定義したターゲットから, どのようにレイアウトするか考える
    - 近いサイトを幾つかピックアップ
  - レイアウトをスケッチする
    - 手書き
      - [Stencils, Sketch Pads & Accessories for User Interface Design — UI Stencils](http://www.uistencils.com/)
    - または photothop など
    - とにかく変更しやすい方法で
  - スケッチを html にする
    - 20 分くらいでさっと
    - sass, chrome dev tool を活用
  - chrome dev tool でいろいろパラメータを変えて, スクリーンショットをとる
  - グレイスケールでレイアウトして, カラーを考える
    - [Review: Iconfactory xScope — PaulStamatiou.com](http://paulstamatiou.com/review-iconfactory-xscope)
  - 色調整, タイポグラフィの調整はいつも最後まで行う
  - ここまではだいたい "ホームページ" のための作業
    - なんとなく頭の中にあるコーディング規約に従って, サイドバーやセカンダリーナビゲーション, コンテンツページ, などの調整をする
  - フィードバックと改善

- Required Reading
  - [Universal Principles of Design](http://www.amazon.com/gp/product/1592535879/ref=as_li_ss_tl?ie=UTF8&linkCode=as2&camp=1789&creative=390957&creativeASIN=1592535879&tag=paulstamatiou-20) や [The Design of Everyday Things](http://www.amazon.com/gp/product/0465067107/ref=as_li_ss_tl?ie=UTF8&linkCode=as2&camp=1789&creative=390957&creativeASIN=0465067107&tag=paulstamatiou-20) のデザインの名著でもいいけども, ここでは web グラフィックデザインにフォーカスする
  - [Non-Designer's Design Book](http://www.amazon.com/gp/product/0321534042/ref=as_li_ss_tl?ie=UTF8&linkCode=as2&camp=1789&creative=390957&creativeASIN=0321534042&tag=paulstamatiou-20)
    - 色の理論の基礎と C.R.A.P (コントラスト, 繰り返し, アラインメント, 近接) を学ぶ
  - [Five Simple Steps - A Practical Guide to Designing for the Web](http://www.fivesimplesteps.com/products/a-practical-guide-to-designing-for-the-web)
    - だいたい網羅されていておすすめ
    - 自分は作者を応援するために買ったけど, 無料でも読める
  - [The Elements of Typographic Style Applied to the Web](http://webtypography.net/toc/)
    - Web デザインにおけるタイポグラフィならこれ

- More Homework
  - [Typekit](https://typekit.com/) のアカウントをとって font をブラウズ, 構造を見たり, font stack について学ぶ
  - 実際にサイトで使ってみる. [フォントのフォールバック](http://blog.typekit.com/2011/03/24/type-study-choosing-fallback-fonts/) を忘れずに
  - [Lettering.js](http://letteringjs.com/) を学ぶ
  - Web デザイナはタイポグラフィを見落としがち. 良いデザイナはテキストを UI として扱う
  - [Principles of grouping - Wikipedia, the free encyclopedia](http://en.wikipedia.org/wiki/Principles_of_grouping)
  - [Particletree » Visualizing Fitts’s Law](http://particletree.com/features/visualizing-fittss-law/)
  - [The Meaning of User Experience \| User Intelligence](http://www.userintelligence.com/ideas/blog/2011/04/meaning-user-experience)

- But Stammy, we launch in one week!!?!111
  - 良いアイコンセットを買う
    - [Picons \| Vector Icons and Pictograms](http://picons.me/)
    - [Pictos](http://pictos.cc/)
    - [Helveticons](http://helveticons.ch/)
  - クリックできるものはすべて hover や action の素材を準備する
    - フィードバック重要
    - `L`o`V`e `HA`te: link, visited, hover, active
  - とにかくアラインさせよ
  - 色
    - [Color Trends + Palettes :: COLOURlovers](http://www.colourlovers.com/) や [0to255](http://0to255.com/) でピックアップ
  - テクスチャ
    - [Tileables - Never Ending Patterns](http://tileabl.es/)
  - プレインなエッジではなくて, 1px のラインを入れる. `rgba(255,255,255, [value from 0-1])`
  - 個別の要素のスタイリングは [UI-Patterns.com](http://ui-patterns.com/) へ
  - より空白を
    - 要素間の空白は, だいたいフォントサイズくらいとればいい
  - たくさんのリストの要素はゼブラカラーに
  - form input には padding を. スタンダードでないフィールドには tooltip で説明を. `:focus` もスタイリングする
  - 一番大きな要素は, 一番重要な情報になってる?
  - スペースになんでもつめこもうとしないこと
  - 余裕があったら A/B テスト
    - [Optimizely: A/B testing software you'll actually use](http://www.optimizely.com/)
  - その他いろいろ

- Try New Things

- UX
  - [52 Weeks of UX](http://52weeksofux.com/)

- More books..
  - [Envisioning Information](http://www.amazon.com/Envisioning-Information-Edward-R-Tufte/dp/0961392118/ref=sr_1_1?s=books&ie=UTF8&qid=1302222716&sr=1-1&tag=paulstamatiou-20)
  - [The Elements of User Experience](http://www.amazon.com/gp/product/0321683684/ref=as_li_ss_tl?ie=UTF8&linkCode=as2&camp=1789&creative=390957&creativeASIN=0321683684&tag=paulstamatiou-20)
  - [Don't Make Me Think](http://www.amazon.com/Dont-Make-Me-Think-Usability/dp/0321344758/ref=sr_1_3?ie=UTF8&s=books&qid=1302222422&sr=8-3&tag=paulstamatiou-20)
  - [Rocket Surgery Made Easy](http://www.amazon.com/Rocket-Surgery-Made-Easy-Yourself/dp/0321657292/ref=pd_bxgy_b_img_b&tag=paulstamatiou-20)
  - [User Interface Design for Programmers](http://www.amazon.com/User-Interface-Design-Programmers-Spolsky/dp/1893115941/ref=sr_1_1?s=books&ie=UTF8&qid=1302222505&sr=1-1&tag=paulstamatiou-20)
  - [Designing Interfaces: Patterns for Effective Interaction Design](http://www.amazon.com/Designing-Interfaces-Patterns-Effective-Interaction/dp/0596008031/ref=sr_1_1?s=books&ie=UTF8&qid=1302222526&sr=1-1&tag=paulstamatiou-20)
  - [The Nature & Aesthetics of Design](http://www.amazon.com/Nature-Aesthetics-Design-David-Pye/dp/0713652861/ref=pd_bxgy_b_text_b&tag=paulstamatiou-20)
  - [Designing for the Digital Age: How to Create Human-Centered Products and Services](http://www.amazon.com/Designing-Digital-Age-Human-Centered-Products/dp/0470229101/ref=sr_1_1?s=books&ie=UTF8&qid=1302240617&sr=1-1&tag=paulstamatiou-20)
  - [The Humane Interface: New Directions for Designing Interactive Systems](http://www.amazon.com/Humane-Interface-Directions-Designing-Interactive/dp/0201379376/ref=sr_1_1?s=books&ie=UTF8&qid=1302222607&sr=1-1&tag=paulstamatiou-20)
  - [The Inmates Are Running the Asylum: Why High Tech Products Drive Us Crazy and How to Restore the Sanity](http://www.amazon.com/Inmates-Are-Running-Asylum-Products/dp/0672326140/ref=sr_1_1?s=books&ie=UTF8&qid=1302222641&sr=1-1&tag=paulstamatiou-20)
  - [About Face 3: The Essentials of Interaction Design](http://www.amazon.com/About-Face-Essentials-Interaction-Design/dp/0470084111/ref=pd_sim_b_1&tag=paulstamatiou-20)
  - [The Smashing Book](https://shop.smashingmagazine.com/smashing-book-2-intl.html)
  - [The Web Designer's Idea Book #2](http://www.amazon.com/Web-Designers-Idea-Book-Vol/dp/160061972X/ref=sr_1_1?ie=UTF8&s=books&qid=1302222562&sr=1-1&tag=paulstamatiou-20)
  - [Designing Interactions](http://www.amazon.com/Designing-Interactions-Bill-Moggridge/dp/0262134748/ref=sr_1_1?s=books&ie=UTF8&qid=1302222693&sr=1-1&tag=paulstamatiou-20)
  - [The Information Design Handbook](http://www.amazon.com/Information-Design-Handbook-Visocky-OGrady/dp/160061048X/ref=sr_1_1?s=books&ie=UTF8&qid=1302222743&sr=1-1&tag=paulstamatiou-20)
  - [Change by Design: How Design Thinking Transforms Organizations and Inspires Innovation](http://www.amazon.com/Change-Design-Transforms-Organizations-Innovation/dp/0061766089/ref=sr_1_1?s=books&ie=UTF8&qid=1302222757&sr=1-1&tag=paulstamatiou-20)
  - [The Art of Innovation Lessons in Creativity from IDEO, America's Leading Design Firm](http://www.amazon.com/Art-Innovation-Lessons-Creativity-Americas/dp/0385499841/ref=sr_1_1?s=books&ie=UTF8&qid=1302222783&sr=1-1&tag=paulstamatiou-20)
  - [Thoughtless Acts?: Observations on Intuitive Design](http://www.amazon.com/Thoughtless-Acts-Observations-Intuitive-Design/dp/0811847756/ref=sr_1_1?ie=UTF8&s=books&qid=1302222862&sr=1-1&tag=paulstamatiou-20)
