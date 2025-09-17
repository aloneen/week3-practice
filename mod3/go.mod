module github.com/aloneen/week3-practice/mod3

go 1.23

replace github.com/aloneen/week3-practice/mod1 => ../mod1

replace github.com/aloneen/week3-practice/mod2 => ../mod2

require (
	github.com/aloneen/week3-practice/mod1 v0.0.0-20250917133300-a9011872321e
	github.com/aloneen/week3-practice/mod2 v0.0.0-20250917133300-a9011872321e
)
