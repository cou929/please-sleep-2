{"title":"32bit or 64bit の見分け方","date":"2012-12-03T01:29:48+09:00","tags":["nix"]}

`uname -a` して x86_64 とか書いてあれば 64bit 対応カーネルと思えばいい. cpu は `/proc/cpuinfo` みて, flags に `lm` という値が入っていれば 64bit と思えばいい (lm はロングモードの略)
