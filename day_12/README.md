# Day 12: Garden Groups
Well, that was an interesting one...

Part 1 was pretty easy. I go through the entire grid, performing a "flood fill" every time I find a grid point I haven't yet visited. The flood fill adds up the number of matching points and also calculates the edge/perimeter for the region as it goes.

Part 2 certainly looked much more tricky. I saw a hint that suggested counting corners as a way to determine the number of sides. Indeed, as I worked through a few examples, this seemed to work out correctly. The tricky part was working out how to determine that a particular point is a corner. In the case of `AAAA`, both the left-most and right-most points each have two corners, so we have to work around that as well.

I found this really great hint - https://www.reddit.com/r/adventofcode/comments/1hcpgcp/comment/m1pxhb8/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button. It explained very succinctly how to calculate a corner. Coded it up and it worked first time. Woot!
