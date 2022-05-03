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
| [**uGO**](https://github.com/ozanh/ugo) | `2.012461s` | `1.123ms` | uGO (VM) |
| [Tengo](https://github.com/d5/tengo) | `2.191303s` | `1.128ms` | Tengo (VM) |
| [go-lua](https://github.com/Shopify/go-lua) | `2.970045s` | `902µs` | Lua (VM) |
| [GopherLua](https://github.com/yuin/gopher-lua) | `3.236942s` | `1.424ms` | Lua (VM) |
| [goja](https://github.com/dop251/goja) | `3.359767s` | `1.412ms` | JavaScript (VM) |
| [starlark-go](https://github.com/google/starlark-go) | `4.714076s` | `1.37ms` | Starlark (Interpreter) |
| [Yaegi](https://github.com/traefik/yaegi) | `9.942251s` | `5.359ms` | Yaegi (Interpreter) |
| [gpython](https://github.com/go-python/gpython) | `9.183582s` | `1.6ms` | Python (Interpreter) |
| [otto](https://github.com/robertkrimen/otto) | `43.902086s` | `4.786ms` | JavaScript (Interpreter) |
| [Anko](https://github.com/mattn/anko) | `51.944854s` | `1.929ms` | Anko (Interpreter) |
| Go | `43.428ms` | `788µs` | Go (Native) |
| Lua | `995.895ms` | `448µs` | Lua (Native) |
| Python3 | `1.640003s` | `6.296ms` | Python3 (Native) |

## 2022-05-03 Results (Apple M1 Parallels VM, Ubuntu 20.04, Go1.17)

For uGO version 0.2.0 release candidate

| | fib(35) | fibt(35) |  Language (Type)  |
| :--- |    ---: |     ---: |  :---: |
| [**uGO**](https://github.com/ozanh/ugo) | `1.907682s` | `1.102ms` | uGO (VM) |
| [Tengo](https://github.com/d5/tengo) | `2.405873s` | `1.173ms` | Tengo (VM) |
| [go-lua](https://github.com/Shopify/go-lua) | `4.151206s` | `1.097ms` | Lua (VM) |
| [GopherLua](https://github.com/yuin/gopher-lua) | `4.261906s` | `1.327ms` | Lua (VM) |
| [goja](https://github.com/dop251/goja) | `4.628215s` | `1.571ms` | JavaScript (VM) |
| [starlark-go](https://github.com/google/starlark-go) | `5.324552s` | `1.363ms` | Starlark (Interpreter) |
| [Yaegi](https://github.com/traefik/yaegi) | `12.066798s` | `6.259ms` | Yaegi (Interpreter) |
| [gpython](https://github.com/go-python/gpython) | `11.700974s` | `1.593ms` | Python (Interpreter) |
| [otto](https://github.com/robertkrimen/otto) | `52.803458s` | `5.086ms` | JavaScript (Interpreter) |
| [Anko](https://github.com/mattn/anko) | `58.107654s` | `2.113ms` | Anko (Interpreter) |
| Go | `90.089ms` | `999µs` | Go (Native) |
| Lua | `1.036239s` | `559µs` | Lua (Native) |
| Python3 | `1.640863s` | `6.398ms` | Python3 (Native) |

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
