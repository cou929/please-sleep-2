{"title":"LWP::UserAgent で SSL 接続すると warn が出る","date":"2013-05-11T16:57:06+09:00","tags":["perl"]}

具体的にはこんな warn メッセージが出る。

    *******************************************************************
     Using the default of SSL_verify_mode of SSL_VERIFY_NONE for client
     is deprecated! Please set SSL_verify_mode to SSL_VERIFY_PEER
     together with SSL_ca_file|SSL_ca_path for verification.
     If you really don't want to verify the certificate and keep the
     connection open to Man-In-The-Middle attacks please set
     SSL_verify_mode explicitly to SSL_VERIFY_NONE in your application.
    *******************************************************************

このメッセージを出しているのは IO::Socket::SSL で、`SSL_verify_mode` オプションを明示的に指定せよという内容。このコミットでこのインタフェース変更が行われたようだ。

[1.79 - start migration to more secure default of SSL_verify_mode by issu... · e388825 · noxxi/p5-io-socket-ssl](https://github.com/noxxi/p5-io-socket-ssl/commit/e3888257eda1ad9ba69a2334c2a415576876c6b8)

ちなみに LWP::UserAgent は SSL 接続のバックエンドとして Net::SSL も使うことができる。その場合はこの warn は出ないはずだ。また他の IO::Socket::SSL を使っているモジュールでも同様のはずである。

解消法としては `ssl_opts` オプションで `SSL_verify_mode` を明示的に指定してあげれば良い。たとえば証明書のチェックをちゃんと行わないようにしたいときは:

<pre><code data-language="perl">use LWP::UserAgent;
use IO::Socket::SSL qw/SSL_VERIFY_NONE/;

my $ua = LWP::UserAgent->new(
    ssl_opts => {
        verify_hostname => 0,
        SSL_verify_mode => SSL_VERIFY_NONE,
    }
)</code></pre>

Furl は pod のドキュメントでこのことに言及していて親切だった。

[Furl - Lightning-fast URL fetcher - metacpan.org](https://metacpan.org/module/TOKUHIROM/Furl-2.15/lib/Furl.pm)

ちなみに、 LWP::UserAgent は昔は LWP::Protocol::https をバンドルしていたが、次のコミットからは分離するようになった。LWP::UserAgent で SSL 接続する場合は LWP::Protocol::https も別途入れてあげる必要がある。

[Unbundle LWP::Protocol::https [RT#66838] · 31278a5 · libwww-perl/libwww-perl](https://github.com/libwww-perl/libwww-perl/commit/31278a53c9dfb66deb8d68e3926343b675b57edc)
