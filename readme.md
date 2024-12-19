# Advent Of Code, 2024

This project contains my Advent of Code puzzle input and solutions for 2024! 😎

I was _going_ to write these solutions in Clojure but my friend and co-worker [skershaw](https://github.com/skershaw/kotlin-AoC-2024) did his in Kotlin as a way to learn the language. I've been somewhat interested in learning Go for a while, I decided to take a page out of his book and do the puzzles in Go this year. I also need to brush up on my Java, I'm going to write solutions in Java as well. 😅

Beware, I'm learning Go as I go and many of these solutions are probably going to be clunky and not-so-idiomatic. But they will be correct and hopefully get more idiomatic and less clunky as I code up more of them. 🤞

## Running the Solutions

Go and Java have their own way of doing things, of course.

### Go

Change directories to the particular day and then the "go" directory. From there you can ask the `go` command to compile the solution. A new executable will be created, named after the day (i.e. `day01_solution`). You may then execute it to see the output.

```shell
cd day01/go
go build
./day01_solution
```

### Java

Change directories to the particular day and then the "java" directory. You may invoke the `java` command directly and pass is the source file, it will then compile and run the class file and you can see the output.

```shell
cd day01/java
java solution.java
```
