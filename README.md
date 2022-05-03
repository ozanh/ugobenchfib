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

## 2022-05-03 Results (Apple M1 Parallels VM, Ubuntu 20.04, Go1.18)

For uGO version 0.2.0 release candidate

| | fib(35) | fibt(35) |  Language (Type)  |
| :--- |    ---: |     ---: |  :---: |
| [**uGO**](https://github.com/ozanh/ugo) | `1.984686s` | `1.261ms` | uGO (VM) |
| [Tengo](https://github.com/d5/tengo) | `2.185415s` | `1.135ms` | Tengo (VM) |
| [go-lua](https://github.com/Shopify/go-lua) | `2.898563s` | `1.006ms` | Lua (VM) |
| [GopherLua](https://github.com/yuin/gopher-lua) | `3.211634s` | `1.74ms` | Lua (VM) |
| [goja](https://github.com/dop251/goja) | `3.326141s` | `1.414ms` | JavaScript (VM) |
| [starlark-go](https://github.com/google/starlark-go) | `4.726492s` | `1.369ms` | Starlark (Interpreter) |
| [Yaegi](https://github.com/traefik/yaegi) | `9.918823s` | `6.171ms` | Yaegi (Interpreter) |
| [gpython](https://github.com/go-python/gpython) | `9.293992s` | `1.551ms` | Python (Interpreter) |
| [otto](https://github.com/robertkrimen/otto) | `43.896718s` | `4.697ms` | JavaScript (Interpreter) |
| [Anko](https://github.com/mattn/anko) | `52.176917s` | `1.902ms` | Anko (Interpreter) |
| Go | `43.266ms` | `892µs` | Go (Native) |
| Lua | `1.000147s` | `463µs` | Lua (Native) |
| Python3 | `1.639605s` | `6.316ms` | Python3 (Native) |

## 2022-05-03 Results (Apple M1 Parallels VM, Ubuntu 20.04, Go1.17)

For uGO version 0.2.0 release candidate

| | fib(35) | fibt(35) |  Language (Type)  |
| :--- |    ---: |     ---: |  :---: |
| [**uGO**](https://github.com/ozanh/ugo) | `2.126575s` | `1.193ms` | uGO (VM) |
| [Tengo](https://github.com/d5/tengo) | `2.406659s` | `1.215ms` | Tengo (VM) |
| [go-lua](https://github.com/Shopify/go-lua) | `4.136014s` | `933µs` | Lua (VM) |
| [GopherLua](https://github.com/yuin/gopher-lua) | `4.266191s` | `1.471ms` | Lua (VM) |
| [goja](https://github.com/dop251/goja) | `4.647449s` | `1.748ms` | JavaScript (VM) |
| [starlark-go](https://github.com/google/starlark-go) | `5.340413s` | `1.333ms` | Starlark (Interpreter) |
| [Yaegi](https://github.com/traefik/yaegi) | `11.9866s` | `6.287ms` | Yaegi (Interpreter) |
| [gpython](https://github.com/go-python/gpython) | `11.683367s` | `1.613ms` | Python (Interpreter) |
| [otto](https://github.com/robertkrimen/otto) | `52.584071s` | `5.265ms` | JavaScript (Interpreter) |
| [Anko](https://github.com/mattn/anko) | `58.179963s` | `1.948ms` | Anko (Interpreter) |
| Go | `91.376ms` | `862µs` | Go (Native) |
| Lua | `1.055571s` | `623µs` | Lua (Native) |
| Python3 | `1.645764s` | `6.404ms` | Python3 (Native) |

## 2022-05-03 Results (Apple M1 Parallels VM, Ubuntu 20.04, Go1.17)

For uGO version 0.1.2

| | fib(35) | fibt(35) |  Language (Type)  |
| :--- |    ---: |     ---: |  :---: |
| [**uGO**](https://github.com/ozanh/ugo) | `2.128649s` | `1.19ms` | uGO (VM) |
| [Tengo](https://github.com/d5/tengo) | `2.414746s` | `1.152ms` | Tengo (VM) |
| [go-lua](https://github.com/Shopify/go-lua) | `4.137094s` | `923µs` | Lua (VM) |
| [GopherLua](https://github.com/yuin/gopher-lua) | `4.284386s` | `1.891ms` | Lua (VM) |
| [goja](https://github.com/dop251/goja) | `4.657529s` | `1.731ms` | JavaScript (VM) |
| [starlark-go](https://github.com/google/starlark-go) | `5.36188s` | `1.386ms` | Starlark (Interpreter) |
| [Yaegi](https://github.com/traefik/yaegi) | `11.98428s` | `5.795ms` | Yaegi (Interpreter) |
| [gpython](https://github.com/go-python/gpython) | `11.696546s` | `1.531ms` | Python (Interpreter) |
| [otto](https://github.com/robertkrimen/otto) | `52.983575s` | `5.462ms` | JavaScript (Interpreter) |
| [Anko](https://github.com/mattn/anko) | `58.27169s` | `1.902ms` | Anko (Interpreter) |
| Go | `89.783ms` | `784µs` | Go (Native) |
| Lua | `995.561ms` | `464µs` | Lua (Native) |
| Python3 | `1.648357s` | `6.348ms` | Python3 (Native) |

## 2020-12-04 Results (Intel i5, Ubuntu 18.04, Go1.15)

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
