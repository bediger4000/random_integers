# Probability of Two Random Integers Relatively Prime

Theoretically,
the probability of two random integers being relatively prime
is 6/(π * π) = 0.6079271020...

# Build and run

[Code](ri.go)

```sh
$ go build ri.go
$ ./ri 100000  # 100,000 pairs of random (63 bit) integers
0.604240
$
```

This actually works.
The more pairs you examine,
the closer you get to the theoretical value.

I wonder if different pseudo-random number generators make different answers?
