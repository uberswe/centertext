# centertext

Still a work in progress

A package which will take an image.Image, truetype.Font and text string and then draw the text centered on the provided image.

Here is how it could be used

```go
out, err := centertext.OnImage(img, font, "Hello world")
```