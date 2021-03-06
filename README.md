# Bencode

[![License](https://img.shields.io/github/license/ArunMurugan78/bencode)](https://github.com/ArunMurugan78/bencode/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/ArunMurugan78/bencode.svg)](https://pkg.go.dev/github.com/ArunMurugan78/bencode)
[![Test](https://github.com/ArunMurugan78/bencode/actions/workflows/test.yml/badge.svg)](https://github.com/ArunMurugan78/bencode/actions/workflows/test.yml)

bencode is a golang package for bencoding and bdecoding data from and from to equivalents. 

>Bencode (pronounced like Ben-code) is the encoding used by the peer-to-peer file sharing system BitTorrent for storing and transmitting loosely structured data. Wikipedia

## Installation
```bash
go get github.com/ArunMurugan78/bencode
```

## Quick Start

```go
package main

import (
	"bytes"
	"fmt"

	"github.com/ArunMurugan78/bencode"
)

func main() {
	data := map[string]interface{}{"name": "rose", "list": []interface{}{"pickachu", "anakin", "NYC", 78}}

	encodedData, _ := bencode.Encode(data)

	fmt.Println(string(encodedData)) // d4:listl8:pickachu6:anakin3:NYCi78ee4:name4:rosee

	decodedData, _ := bencode.Parse(bytes.NewReader(encodedData))

	fmt.Println(decodedData) // map[list:[pickachu anakin NYC 78] name:rose]
}

```

## Invalid Data?
To be able to encode a given data into bencode, it's data type has to be one of the below.

| Type | Description|
| :--  | :--- |
| `int`| Can be any one of the `int` and `uint` variants.|
| `string` | UTF-8 encoded strings.|
| `[]interface{}`| The slice can hold any one of the valid data type.|
| `map[string]interface{}` | The key has to be a string and the value can be any one of the valid data type.


## License
The project is licensed under the [MIT](https://github.com/ArunMurugan78/bencode/blob/master/LICENSE) license.