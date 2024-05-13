# Deez Distance

An implementation of the wagner-fischer algorithm, 
for detecing the "distance" between two strings, in Golang.

## Sources:
- https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm
- https://en.wikipedia.org/wiki/Levenshtein_distance

## Usage:

```bash
$ go run main.go -w hello
Distance from Deez: 4.00

$ go run main.go -w deez
Distance from Deez: 0.00
```

