<p align="center">
  <img src="demo.gif">
</p>

<p align="center">
  <a href="https://travis-ci.org/mbrlabs/clispin">
      <img src="https://travis-ci.org/mbrlabs/clispin.svg?branch=master" alt="Travis status">
  </a>
  <a href="https://godoc.org/github.com/mbrlabs/clispin">
      <img src="https://godoc.org/github.com/mbrlabs/clispin?status.svg" alt="Godoc">
  </a>
</p>

# clispin
Clispin is a Go library, that makes it ridiculously easy to integrate fancy unicode spinners in your app.

 ## Usage
```go
spinner := clispin.New(nil)
spinner.Start(func() {
    // Do your work here. Update spinner text with Print or Printf
    spinner.Printf("Downloading file %d/2", 1)
    time.Sleep(time.Second)
    spinner.Printf("Downloading file %d/2", 2)
    time.Sleep(time.Second)

    spinner.Print("Processing files")
    time.Sleep(time.Second)
    spinner.Print("Done")
})
```

```clispin.New()``` creates a new spinner. You can pass in a ```Sprite``` struct to customize the spinner
animation. If no sprite is provided, a default one is chosen.      

There are also a couple of [sprites](https://github.com/mbrlabs/clispin/blob/master/sprite.go) included in this library. 