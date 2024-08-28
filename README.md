# imgut - Image Utility Tool
[![GitHub tag (with filter)](https://img.shields.io/github/v/release/floholz/imgut?label=latest)](https://github.com/floholz/imgut/releases/latest)
[![GitHub License](https://img.shields.io/github/license/floholz/imgut)](./LICENSE)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/floholz/imgut?logo=go&labelColor=gray&label=%20)](https://go.dev/dl/)


> This is a simple tool to perform url pattern shenanigans. 


## Install

```shell
go install github.com/floholz/imgut@latest
```

## Usage

To download images matching an url pattern, use the `download` command.

```shell
imgut download "https://flagsapi.com/B{A-F}/flat/64.png"
```

To only check if a matching url is valid, use the `fuzz` command.

```shell
imgut fuzz "https://flagsapi.com/B{A-F}/flat/64.png"
```

See the full documentation with the `help` command or at [doc.md](doc.md).

```shell
imgut help
```

---

### 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](https://github.com/floholz/imgut/issues).


### 📝 License

This project was created by [floholz](https://github.com/floholz) under the [GPL-3.0](./LICENSE) license.

---

### Show your support

Give a ⭐ if this project helped you!