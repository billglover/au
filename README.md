# Algorithms Unlocked

This repository contains Go implementations of the algorithms discussed in Algorithms Unlocked. They don't claim to be the best implementations, but a collection of implementations that helped me learn and understand the algorithms discussed in the book.

Although I'm undertaking this as a personal exercise, I welcome issues, discussion and contribution from others. The intent is to facilitate discussion and learning, so all innovative contributions are welcome.

## Guidelines

Because there are many ways to tackle the implementation of these algorithms, I've set out some basic guidelines to help facilitate learning. These are guidelines and not rules. They can be ignored, but if so, reasons should be documented along with the algorithm in question.

1. Use the Go standard library only
2. Include benchmarks for each algorithm
3. Include tests that validate the correctness of the result
4. Include the pseudocode implementation in the `README.md` for each algorithm
5. Follow the folder structure outlined below

## Structure

An example directory structure is shown below. Where alternative implementations are available these should be implemented as separate functions within the main file.

```plain
linear-search/
├── README.md
├── linear-search.go
└── linear-search_test.go
```
