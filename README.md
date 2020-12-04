# uGO Language Benchmark

For initial setup, run:

```bash
make setup
```

To start the benchmark, run:

```bash
make start
```

You can see all the source codes used in this benchmark in [code](https://github.com/ozanh/ugobenchfib/tree/master/code)
directory.

## 2020-12-04 Results (Intel i5, Ubuntu 18, Go1.15)

| | fib(35) | fibt(35) |  Language (Type)  |
| :--- |    ---: |     ---: |  :---: |
| [**uGO**](https://github.com/ozanh/ugo) | `2.383596s` | `1.566ms` | uGO (VM) |
| [Tengo](https://github.com/d5/tengo) | `2.585655s` | `1.427ms` | Tengo (VM) |
| [go-lua](https://github.com/Shopify/go-lua) | `4.522862s` | `1.094ms` | Lua (VM) |
| [GopherLua](https://github.com/yuin/gopher-lua) | `4.871662s` | `1.674ms` | Lua (VM) |
| [goja](https://github.com/dop251/goja) | `5.300007s` | `1.781ms` | JavaScript (VM) |
| [starlark-go](https://github.com/google/starlark-go) | `11.032018s` | `1.536ms` | Starlark (Interpreter) |
| [Yaegi](https://github.com/traefik/yaegi) | `15.033684s` | `4.343ms` | Yaegi (Interpreter) |
| [gpython](https://github.com/go-python/gpython) | `16.531675s` | `2.136ms` | Python (Interpreter) |
| [otto](https://github.com/robertkrimen/otto) | `1m8.47862s` | `7.017ms` | JavaScript (Interpreter) |
| [Anko](https://github.com/mattn/anko) | `1m16.396712s` | `2.479ms` | Anko (Interpreter) |
| Go | `60.671ms` | `704µs` | Go (Native) |
| Python3 | `2.565661s` | `15.361ms` | Python3 (Native) |

**Note:** This benchmark repository is fork of <https://github.com/d5/tengobench> and its LICENSE is placed in
[NOTICE](NOTICE.txt) file.
