# Math Problems
This repostiory contains bunch of math problems and their solutions in plain English. I will try to include the source for each problem. Sources might be erroneous or not original.

PRs are welcome.

## Probability and Statistics

1. Find the probability of obtaining four of a kind in an ordinary five-card poker hand. [Ash, Basic Probability Theory]

2. Three balls are dropped into three boxes. Find the probability that exactly one box will be empty. [Ash, Basic Probability Theory]

## Algorithms and Data Structures


# Solutions

## Probability and Statistics

### 1. (13)(48) / (52 choose 5)

There are (52 choose 5) distinct poker hands (without regard to order). To obtain four of a kind,

(a) choose the face value to appear four times (13 choices: A, K, Q, ..., 2)

(b) choose the fifth card (52 - 4 = 48 remaing cards to choose from)

Thus *p* = (13)(48) / (52 choose 5)

### 2. 2/3

There are (3 choose 2) ways to choose two non-empty boxes. Now select which box of those two to contain just one (rather than two) balls. Thus *p* = (2) / (3 choose 2) = 2/3

## Algorithms and Data Structures
