# Exercises
These are the exercises from the 10th chapter of the book Learning Go by Jon Bodner.
Source: https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/ch10.html#id351

* Create a module in your own public repository. This module has a single function named Add with two int parameters and one int return value. This function adds the two parameters together and returns them. Make this version v1.0.0.
* Add godoc comments to your module that describe the package and the Add function. Be sure to include a link to https://www.mathsisfun.com/numbers/addition.html in your Add function godoc. Make this version v1.0.1.
* Change Add to make it generic. Import the golang.org/x/exp/constraints package. Combine the Integer and Float types in that package to create an interface called Number. Rewrite Add to take in two parameters of type Number and return a value of type Number. Version your module again. Because this is a backward-breaking change, this should be v2.0.0 of your module.
