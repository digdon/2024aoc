# Day 24: Crossed Wires
Part 1 was easy enough, of course. There are a couple of ways I could have done this. First I thought about using a topological sort to get the proper order for all of the gates. From a stack perspective, this might have been the better choice. Instead, I went with essentially a DFS approach. I sorted all of the 'z' gates from LSB to MSB anbd then processed each one, working my way to the starting point for each.

Part 2 was yet another reverse engineering problem. As I've said before, I'm not good at these, so I went right to reddit for ideas.

Essentially, we're looking at a ripple carry adder. A pretty easy circuit, but I just didn't want to bother breaking it down. After a bunch of searching, I found this [link](https://www.bytesizego.com/blog/aoc-day24-golang), which eventually led me to [here](https://www.reddit.com/r/adventofcode/comments/1hla5ql/2024_day_24_part_2_a_guide_on_the_idea_behind_the/). I had a few bugs along the way, but eventually got the 4 rules implemented correctly and got the star.

Perhaps one day I'll embrace this sort of puzzle.... doubt it though.
