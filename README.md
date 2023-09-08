# rainbowgo
Print rainbows to the terminal in Go very easily. 

Want to print a rainbow to your terminal in Go? It's very easy. This package handles all the work of calculating gradients for you. 

## Inititalise
A variable of type Rainbow can be made and then initalised with the desired length of the rainbow. 
```go
var r rainbowgo.Rainbow
r.Init(int)
```
## Functions:
Print and Println will print a string in a uniform colour.
```go
r.Print(string)
r.Println(string)
```
Next will make set the rainbow variable, r, to the next colour.
```go
r.Next()
```
Print2d and Print2dln will print a string with a secondary rainbow going sideways. These functions also take an interger to set the length of the second rainbow. The second rainbow will start with the same colour as the main rainbow. This can be used to make diagonal or even curved rainbows. 
```go
r.Print2d(string, int)
r.Print2dln(string, int)
```

## Full Example
```go
package main

import "github.com/Sophuwu300/rainbowgo"

func main() {
    var r rainbowgo.Rainbow
    r.Init(20)
    for i := 0; i < 20; i++ {
        r.Print("Hello Github! I am a rainbow. ")
        r.Print2dln("I am a rainbow with more colours!", 35) 
        r.Next() 
    }
}
```
## Result:
![image](https://github.com/Sophuwu300/rainbowgo/assets/78772568/c0a9fe5f-279a-46d2-945b-dc798026843e)
