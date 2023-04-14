# Go Color Package

[![codecov](https://codecov.io/gh/alan890104/color/branch/main/graph/badge.svg?token=Q5B4MEAJS9)](https://codecov.io/gh/alan890104/color)

Package color provides functions to colorize text in the terminal.
It is simple, lightweight, and easy to use.

## Install

```go
import "github.com/alan890104/color"
```

## Usage

This is a simple example of using the package.

```go
package main

import (
    "fmt"
    "github.com/alan890104/color"
)

func main() {
    // Print colorized text
    fmt.Println(color.Red("This is red"))
    fmt.Println(color.Green("This is green"))
    fmt.Println(color.Yellow("This is yellow"))
    fmt.Println(color.Blue("This is blue"))
    fmt.Println(color.Black("This is black"))

    // Print colorized text with format
    fmt.Println(color.Magenta("This is %s", "magenta"))
    fmt.Println(color.Cyan("This is %s", "cyan"))
    fmt.Println(color.White("This is %s", "white"))
}
```

Also, you can use `Info`, `Debug`, `Fatal`, `Warning`, `Error`, `Success` to print colorized text.

```go
package main

import (
    "fmt"

    "github.com/alan890104/color"
)

func main() {
    fmt.Println(color.Infof("This is %s", "info"))
    fmt.Println(color.Debugf("This is %s", "debug"))
    fmt.Println(color.Fatalf("This is %s", "fatal"))
    fmt.Println(color.Warningf("This is %s", "warning"))
    fmt.Println(color.Errorf("This is %s", "error"))
    fmt.Println(color.Successf("This is %s", "success"))
}
```

## Supported Colors

<table>
  <tr>
    <th>Color Name</th>
    <th>Color</th>
  </tr>
  <tr>
    <td>Black</td>
    <td style="color:#000000">Black</td>
  </tr>
  <tr>
    <td>Red</td>
    <td style="color:#ff0000">Red</td>
  </tr>
  <tr>
    <td>Green</td>
    <td style="color:#00ff00">Green</td>
  </tr>
  <tr>
    <td>Yellow</td>
    <td style="color:#ffff00">Yellow</td>
  </tr>
  <tr>
    <td>Blue</td>
    <td style="color:#0000ff">Blue</td>
  </tr>
  <tr>
    <td>Magenta</td>
    <td style="color:#ff00ff">Magenta</td>
  </tr>
  <tr>
    <td>Cyan</td>
    <td style="color:#00ffff">Cyan</td>
  </tr>
  <tr>
    <td>White</td>
    <td style="color:#ffffff">White</td>
  </tr>
</table>
