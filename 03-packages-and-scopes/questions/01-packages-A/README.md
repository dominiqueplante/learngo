## Where to store the source code files that belong to a package?
1. Each file should go into a different directory
2. In a single directory

Correct Answer: 2

## Why a package clause is used in a Go source-code file?
1. It's used for importing a package
2. It's used for letting Go know that the file belongs to a package
3. It's used for declaring a new function

Correct Answer: 2

> **1:** `import` statement does that.
>
> **3:** `func` statement does that.

## Where you should put the `package clause` in a Go source-code file?
1. As the first code in a Go source code file
2. As the last code in a Go source code file
3. You can put it anywhere

Correct Answer: 1

## How many times you can use `package clause` for a single source code file?
1. Once
2. None
3. Multiple times

Correct Answer: 1

## What is the correct way to use the `package clause`?
1. `my package`
2. `package main`
3. `pkg main`

Correct Answer: 2

## Which one is correct?
1. All files belong to the same package cannot call each others' functions
2. All files belong to the same package can call each others' functions

Correct Answer: 2

## How to run multiple Go files?
1. `go run *.*go`
2. `go build *go`
3. `go run go`
4. `go run *.go`

Correct Answer: 4

> **4:** You can also call it like (assuming there are `file1.go` `file2.go` and `file3.go` in the same directory): `go run file1.go file2.go file3.go`