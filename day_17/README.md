# Day 17: Chronospatial Computer

I wondered about this day.... The reddit comments were all "woot, intcode computer", but I suspected something more nefarious was going to happen (looking at you, [arithmetic logic unit](https://adventofcode.com/2021/day/24)). And, sadly, I was right.

I despise these kinds of problems. I spent a bit of time trying to reverse engineer what was happening, but this is not something I'm even remotely good at.

Read a comment that mentioned using octals for everything and with that, I noticed that there was one octal digit for every item of output. My program is 16 characters, so to replicate, I'll need 16 octal digits in my answer. From there, I noticed that A was ultimately dividing by 8 every time and that the digits-to-output were happening in reverse (least-significant digit produced first output and so on, with the most-significant digit producing the final output item). Then I saw this - https://www.reddit.com/r/adventofcode/comments/1hgo81r/2024_day_17_genuinely_enjoyed_this/. That really gave me something to go with.

I tried a variety of things that involved running through numbers 0-7 for each item, and recursing when a match was found. I tried forwards and backwards, but it never worked.

Back to reddit. Eventually found this - https://www.reddit.com/r/adventofcode/comments/1hg38ah/comment/m2pyn7q/. And that showed me what I was doing wrong. My assumption was that for each position, I was looking for that many values in the output, but that's not completely correct. I'm not entirely sure why yet... Ultimately, the check I needed was this: `slicesEqual(output, program[len(program)-pos:]` (unfortunately I don't remember what I was originally using... I think I was using something more like `len(program)-1-pos`). Anyway, that simple change fixed my code.

UPDATE: Saw something on Reddit way after the fact that mentioned that the xdv commands (divide by 2^x) are ultimately just bit-shifting to the right by the operand value. I verified that this value doesn't get too large and modified my code to do the bit-shift instead of calculating 2^x and then doing the divide. Interestingly, part 1 took 1 extra microsecond, but the execution of part two was cut in half. I'll call that a win.
