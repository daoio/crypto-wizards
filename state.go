package main

/*
	Simple game state

	0 - new game started
	{
		iter:
		1 - player 1 turn
		2 - player 2 turn
	}
	3 - player 1 wins
	4 - player 2 wins
*/
var State int

func init() {
	// i.e. game has been started
	State = 0
}
