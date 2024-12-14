# Day 14: Restroom Redoubt

Part 1 was pretty easy and could easily be done via brute force. Of course, it's logical to assume part 2 would simply be, "ok, now what about 100,000,000,000 seconds?" Given that, it was pretty easy to work out the position of any robot at any second using simple math. Done.

Part 2..... but of course. Reddit was full of a lot of displeasure about this. I think the biggest issue is that we need to find an iteration that displays a tree, without any desciption of the tree at all.

I started part 2 by iterating through every second from 0 to 10,000. In order to form some sort of tree, I assumed that there'd have to be a column with a large number of robots in it. Whenever I find a column with more than a certain number of robots, I print out the grid to see if a tree formed. I fiddled with this value a bunch of times until I got to 35, at which point I saw the tree.

Once I knew what the target looked like, I was able to refine my end conditions. The tree itself is contained in a box, so rather than looking just for a column with 35 robots in it, I modified the check to be:
1. look for a column with 30 or more robots in it
1. check that column to see if there are 30 or more consecutive robots

If both conditions are satisfied, we very likely found the iteration.

Saw a comment on reddit that suggested that at the target iteration, each robot would be in a unique position. I tried this with my own result and can confirm that this is the case.
