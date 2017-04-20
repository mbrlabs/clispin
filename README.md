 # clispin
 Wrap your functions in elegant unicode spinners.

 ## Usage
```go
spinner := clispin.New(nil)
spinner.Start(func() {
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