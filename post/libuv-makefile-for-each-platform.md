{"title":"libuv ビルド時のプラットフォームごとの対処","date":"2012-11-25T22:26:22+09:00","tags":["nix, libuv"]}

全部 Makefile で吸収されている. `uname -s` の値に応じてプラットフォームごとに適切なオブジェクトファイルのコンパイル, リンクが行われ, 最終的に `uv.a` ができあがる.

まず Makefile の中では MINGW かどうかを判定してそうでなければ `config-unix.mk` をインクルードする.

    ifneq (,$(findstring MINGW,$(uname_S)))
    include config-mingw.mk
    else
    include config-unix.mk
    endif

`config-unix.mk` では以下のようにプラットフォームごとに対象のオブジェクトファイルやフラグをセットする

    ifeq (Darwin,$(uname_S))
    EV_CONFIG=config_darwin.h
    EIO_CONFIG=config_darwin.h
    CPPFLAGS += -D_DARWIN_USE_64_BIT_INODE=1 -Isrc/ares/config_darwin
    LINKFLAGS+=-framework CoreServices
    OBJS += src/unix/darwin.o
    OBJS += src/unix/kqueue.o
    endif

最終的には `ar rcs uv.a ...` が実行されて `uv.a` ができあがる

    uv.a: $(OBJS) src/cares.o src/fs-poll.o src/uv-common.o src/unix/ev/ev.o src/unix/uv-eio.o src/unix/eio/eio.o $(CARES_OBJS)
        $(AR) rcs uv.a $^

実際に mac で make するとこんな感じ

    ar rcs uv.a src/unix/async.o src/unix/core.o src/unix/dl.o src/unix/error.o src/unix/fs.o src/unix/loop.o src/unix/loop-watcher.o src/unix/pipe.o src/unix/poll.o src/unix/process.o src/unix/stream.o src/unix/tcp.o src/unix/thread.o src/unix/timer.o src/unix/tty.o src/unix/udp.o src/unix/darwin.o src/unix/kqueue.o src/cares.o src/fs-poll.o src/uv-common.o src/unix/ev/ev.o src/unix/uv-eio.o src/unix/eio/eio.o src/ares/ares__close_sockets.o src/ares/ares__get_hostent.o src/ares/ares__read_line.o src/ares/ares__timeval.o src/ares/ares_cancel.o src/ares/ares_data.o src/ares/ares_destroy.o src/ares/ares_expand_name.o src/ares/ares_expand_string.o src/ares/ares_fds.o src/ares/ares_free_hostent.o src/ares/ares_free_string.o src/ares/ares_gethostbyaddr.o src/ares/ares_gethostbyname.o src/ares/ares_getnameinfo.o src/ares/ares_getopt.o src/ares/ares_getsock.o src/ares/ares_init.o src/ares/ares_library_init.o src/ares/ares_llist.o src/ares/ares_mkquery.o src/ares/ares_nowarn.o src/ares/ares_options.o src/ares/ares_parse_a_reply.o src/ares/ares_parse_aaaa_reply.o src/ares/ares_parse_mx_reply.o src/ares/ares_parse_ns_reply.o src/ares/ares_parse_ptr_reply.o src/ares/ares_parse_srv_reply.o src/ares/ares_parse_txt_reply.o src/ares/ares_process.o src/ares/ares_query.o src/ares/ares_search.o src/ares/ares_send.o src/ares/ares_strcasecmp.o src/ares/ares_strdup.o src/ares/ares_strerror.o src/ares/ares_timeout.o src/ares/ares_version.o src/ares/ares_writev.o src/ares/bitncmp.o src/ares/inet_net_pton.o src/ares/inet_ntop.o

というわけでコードを読むときは `src/unix` の下を読んでいくので大丈夫そうだ.
