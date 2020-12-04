package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	dir := os.Args[1]

	fmt.Println("| | fib(35) | fibt(35) |  Language (Type)  |")
	fmt.Println("| :--- |    ---: |     ---: |  :---: |")

	// uGO
	fmt.Printf("| [**uGO**](https://github.com/ozanh/ugo) | `%s` | `%s` | uGO (VM) |\n",
		formatDuration(execCommand("ugo", dir+"/fib.ugo")),
		formatDuration(execCommand("ugo", dir+"/fibtc.ugo")))

	// Tengo
	fmt.Printf("| [Tengo](https://github.com/d5/tengo) | `%s` | `%s` | Tengo (VM) |\n",
		formatDuration(execCommand("tengo", dir+"/fib.tengo")),
		formatDuration(execCommand("tengo", dir+"/fibtc.tengo")))

	// go-lua
	fmt.Printf("| [go-lua](https://github.com/Shopify/go-lua) | `%s` | `%s` | Lua (VM) |\n",
		formatDuration(execCommand("go-lua", dir+"/fib.lua")),
		formatDuration(execCommand("go-lua", dir+"/fibtc.lua")))

	// GopherLua (glua)
	fmt.Printf("| [GopherLua](https://github.com/yuin/gopher-lua) | `%s` | `%s` | Lua (VM) |\n",
		formatDuration(execCommand("glua", dir+"/fib.lua")),
		formatDuration(execCommand("glua", dir+"/fibtc.lua")))

	// goja
	fmt.Printf("| [goja](https://github.com/dop251/goja) | `%s` | `%s` | JavaScript (VM) |\n",
		formatDuration(execCommand("goja", dir+"/fib.js")),
		formatDuration(execCommand("goja", dir+"/fibtc.js")))

	// Starlark
	fmt.Printf("| [starlark-go](https://github.com/google/starlark-go) | `%s` | `%s` | Starlark (Interpreter) |\n",
		formatDuration(execCommand("starlark", "-recursion", dir+"/fib.star")),
		formatDuration(execCommand("starlark", "-recursion", dir+"/fibtc.star")))

	// Yaegi
	fmt.Printf("| [Yaegi](https://github.com/containous/yaegi) | `%s` | `%s` | Yaegi (Interpreter) |\n",
		formatDuration(execCommand("yaegi", dir+"/fib.yaegi")),
		formatDuration(execCommand("yaegi", dir+"/fibtc.yaegi")))

	// Gpython
	fmt.Printf("| [gpython](https://github.com/go-python/gpython) | `%s` | `%s` | Python (Interpreter) |\n",
		formatDuration(execCommand("gpython", dir+"/fib.py")),
		formatDuration(execCommand("gpython", dir+"/fibtc.py")))

	// otto
	fmt.Printf("| [otto](https://github.com/robertkrimen/otto) | `%s` | `%s` | JavaScript (Interpreter) |\n",
		formatDuration(execCommand("otto", dir+"/fib.js")),
		formatDuration(execCommand("otto", dir+"/fibtc.js")))

	// Anko
	fmt.Printf("| [Anko](https://github.com/mattn/anko) | `%s` | `%s` | Anko (Interpreter) |\n",
		formatDuration(execCommand("anko", dir+"/fib.ank")),
		formatDuration(execCommand("anko", dir+"/fibtc.ank")))

	// Go
	fmt.Printf("| Go | `%s` | `%s` | Go (Native) |\n",
		formatDuration(execCommand("gofib")),
		formatDuration(execCommand("gofibtc")))

	// Lua
	fmt.Printf("| Lua | `%s` | `%s` | Lua (Native) |\n",
		formatDuration(execCommand("lua", dir+"/fib.lua")),
		formatDuration(execCommand("lua", dir+"/fibtc.lua")))

	// Python
	fmt.Printf("| Python3 | `%s` | `%s` | Python3 (Native) |\n",
		formatDuration(execCommand("python3", dir+"/fib.py")),
		formatDuration(execCommand("python3", dir+"/fibtc.py")))
}

func execCommand(name string, args ...string) time.Duration {
	cmd := exec.Command(name, args...)

	start := time.Now()

	_ = cmd.Run()

	return time.Since(start)
}

func formatDuration(d time.Duration) string {
	return d.Round(time.Microsecond).String()
}
