# Color

### Install Color package

```bash
go get github.com/fatih/color
```

### Printing the color output

```go
// Set the color
color.Set(color.FgBlue)

// Print the colored text
fmt.Println("This string is blue")

// Unset the color
color.Unset()

// Print the colored output without Printf
color.Red("We have red")
```

For more example check https://github.com/fatih/color