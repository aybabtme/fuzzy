# package fuzzy

Fuzzy search of strings.

There is a lazy and an eager matcher.

* The lazy matcher is preferable if matches are queried only once the whole
  fuzzy query is entered.
* The eager matcher is slightly faster if matches are queried everytime the fuzzy
  query is updated.

# Performance

I have not researched proper algorithms for doing this, but the fuzzy match itself
is at most `O(n)` where `n` is the total number of runes in the lines. I also
haven't profiled the code.

Likely it could be faster if working on `[]rune` instead of `string` to save
allocs.
