 # clispin
 Wrap your functions in elegant unicode spinners.

 ## Usage
```go
spinner := clispin.New(nil)
spinner.Start(func() {
    x := 100
    for i := 0; i <= x; i++ {
        spinner.Printf("Downloading %d Gb", i)
        time.Sleep(time.Millisecond * 100)
    }
})
```

## Built in spinners
TODO