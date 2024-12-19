# Day 19: Linen Layout

This one certainly looked relatively easy at the start - a simple DFS, trimming down the matching pattern off the front of the design string and apply some recursion. Worked great for the test data, but, for my input at least, was never finishing.

My work group mentioned DP/memoization for part 2, so I knew that was going to be required at some point, and based on what I was seeing in my part 1, it seemed obvious that I'd had to employ the cache for part 1 as well. Memoization is always a struggle for me because I never know what to use as the key. I started by trying to be clever (re: stupid) and cache based on the current design string PLUS which pattern was being analyzed. Made sense at the time, but, of course, it never finished. However, all I _really_ needed to use for the key was the design portion itself. Which makes sense now... for a particular design string, is there a solution - yes/no. Pattern has nothing to do with it. If a pattern matches, there's a solution and I don't need to know which pattern(s) matched, end of story.

As soon as I figured that out (thanks, Reddit), part 1 started working for me.

Part 2 was incredibly easy and quick change to the code, from merely finding a single solution (return true if a solution is found) to counting all possible solutions and caching that instead.
