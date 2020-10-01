# Serve the files in the "starting-files" folder

## To get your images to serve, use:

```
func StripPrefix(prefix string, h Handler) Handler
func FileServer(root FileSystem) Handler
```