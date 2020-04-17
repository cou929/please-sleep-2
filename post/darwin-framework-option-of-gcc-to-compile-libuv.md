{"title":"mac で libuv を使ったプログラムのコンパイルするには -framework CoreServices というフラグが必要だった","date":"2012-11-25T21:33:39+09:00","tags":["mac"]}

[uvbook](https://github.com/nikhilm/uvbook) を読んでいて, 付属するサンプルコードの `Makefile` を見ていると mac の場合 `-framework CoreServices` というフラグをつけていた

    ifeq (Darwin, $(uname_S))
    CFLAGS+=-framework CoreServices
    SHARED_LIB_FLAGS=-bundle -undefined dynamic_lookup -o plugin/libhello.dylib
    endif

たしかに単純な hello world のサンプルでもこのオプションを付けないとコンパイルできなかった. hello world のサンプルは `uv_loop_new()` して `uv_run()` するだけの簡単なものだ.

    #include <stdio.h>
    #include <uv.h>
    
    int main() {
        uv_loop_t *loop = uv_loop_new();
    
        printf("Now quitting.\n");
        uv_run(loop);
    
        return 0;
    }

以下のようにするとコンパイルが通るが,

    $ gcc -g -Wall -I../../libuv/include ../../libuv/uv.a -framework CoreServices main.c

`framework` オプションがないと次のようにリンクする共有ライブラリが足りてない感じのエラーが出る.

    $ gcc -g -Wall -I../../libuv/include ../../libuv/uv.a main.c
    Undefined symbols for architecture x86_64:
      "_AbsoluteToNanoseconds", referenced from:
          _uv_hrtime in uv.a(darwin.o)
    ld: symbol(s) not found for architecture x86_64
    collect2: ld returned 1 exit status

たぶん libuv の darwin 周りのコードでプラットフォーム独自の何かを使っていて, それをリンクさせるには `-framework CoreServices` とさせないとだめなんだろう.
