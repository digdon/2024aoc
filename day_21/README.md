# Day 21: Keypad Conundrum

The day that seemingly broke AoC. A ridiculously hard puzzle by any measure. It required so much study and working stuff out, piece by tiny piece, that I think a lot of people just went to reddit for solutions. An interesting puzzle, sure, but, as I've said before, I have no interest in these.

The two big things that helped me - [this great post](https://www.reddit.com/r/adventofcode/comments/1hjgyps/2024_day_21_part_2_i_got_greedyish/) and this [very tidy Java implementation](https://github.com/ash42/adventofcode/blob/main/adventofcode2024/src/nl/michielgraat/adventofcode2024/day21/Day21.java).

I copied a lot of ideas from the Java implementation, obviously. I ran into some interesting issues along the way, especially around what `strings.Split` returns. Once I figured out I was getting an extra 'A' (an empty split string), it was smooth sailing... mostly.

The memoization was interesting. Usually I see 10s or 100s of thousands of cache hits, but for part 2, there were only ~1900. Yet that made the difference between never finishing and taking a mere 2ms. Fascinating stuff and I'm always amazed and how well it works.
