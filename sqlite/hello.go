package main

import "fmt"
import "rsc.io/quote/v4"

func main() {
    fmt.Println("Run quote.Go()...")
    fmt.Printf("Glass(): \"%s\"\n",quote.Glass())
    fmt.Printf("Go()   : \"%s\"\n",quote.Go())
    fmt.Printf("Hello(): \"%s\"\n",quote.Hello())
    fmt.Printf("Opt()  : \"%s\"\n",quote.Opt())
    fmt.Println("Completed.")
}
