# beeceej.com

- [A Script written in Go](cmd/main.go)
- [A Makefile](Makefile)
- [Some Templates](./templates)


## Building

```
make clean all
```

## Developing

```
make clean all
cd output
python -m http.server &
cd ../
# make changes...
make
# view changes in browser...
# repeat
```
