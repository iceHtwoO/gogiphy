# GoGiphy
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/icehtwoo/gogiphy/blob/main/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/icehtwoo/gogiphy)

GoGiphy is a golang implementation of the giphy api.
### Installing

`go get` *will always pull the latest release.*

```sh
go get github.com/icehtwoo/gogiphy
```
or
```sh
go get github.com/icehtwoo/gogiphy@vx.y.z
```

### Usage

Import the package into your project.

```go
import "github.com/icehtwoo/gogiphy"
```

Create a giphy client to access the api. The key can be generated [here](https://developers.giphy.com/):
```go
giphy := gogiphy.NewClient(httpClient, "key", "en")
```

Examples can be found [here](https://github.com/iceHtwoO/gogiphy/tree/main/examples).

### TO-DO

- [X] Trending
- [X] Search
- [X] Translate
- [X] Random
- [ ] Action Register
- [ ] Random ID
- [ ] Get GIF by ID
- [ ] Get GIFs by ID
- [ ] Upload
- [ ] Categories
- [X] Autocomplete
- [X] Search suggestions
- [X] Trending Search Terms
- [ ] Documentation
- [ ] Tests
- [X] Examples
