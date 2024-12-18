# Difference between := and = operators in Golang

In golang, `:=` is for declaration, assignment and also for reassignment,  while `=` is for assignment only.

## Rules

1.  You cannot use `:=` outside of funcs. It is becasuse, outside a func, a statement should start with a keyword:

```
// no keywords below, illega.
illegal := 42

// `var` keyword makes this statement legal.
var legal = 42

func foo() {
  alsoLegal := 42
  // reason: it's in a func scope.
}
```

2. You can't use them twice (in the same scope):

```
legal := 42
legal := 42 // <-- error
```

Because, `:=` introduces "a new variable", hence using it twice does not redeclare a second variable, so it's illegal.

3. You can use them for multi-variable declarations and assignments:

```
foo, bar   := 42, 314
jazz, bazz := 22, 7
```

4. You can use them twice in "multi-variable" declarations, if one of the variables is new:

```
foo, bar  := someFunc()
foo, jazz := someFunc()  // <-- jazz is new
baz, foo  := someFunc()  // <-- baz is new
```

This is legal, because, you're not declaring all the variables, you're just reassigning new values to the existing variables, and declaring new variables at the same time. This is called redeclaration.

5. You can use the short declaration to declare a variable in a newer scope even if that variable is already declared with the same name before:

```
var foo int = 34

func some() {
  // because foo here is scoped to some func
  foo := 42  // <-- legal
  foo = 314  // <-- legal
}
```

Here, `foo := 42` is legal, because, it declares `foo` in `some()` func's scope. `foo = 314` is legal, because, it just assigns a new value to `foo`.

6. You can declare the same name in short statement blocks like: if, for, switch:

```
foo := 42
if foo := someFunc(); foo == 314 {
  // foo is scoped to 314 here
  // ...
}
// foo is still 42 here
```

Because, foo in if foo := ..., only belongs to that if clause and it's in a different scope.

## Reference
- [Short Variable Declaration Rules](https://blog.learngoprogramming.com/golang-short-variable-declaration-rules-6df88c881ee)
- [A Visual Guide to GO Variables](https://blog.learngoprogramming.com/learn-go-lang-variables-visual-tutorial-and-ebook-9a061d29babe)