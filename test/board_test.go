package test

import (
	"github.com/eliias/peterpawner/chess"
	"testing"
)

func TestFen(t *testing.T) {
	var game chess.Game
	var fen string
	var res string

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	game = chess.Load(fen)
	res = chess.Save(game)

	if res != fen {
		t.Error(
			"For", "FEN 1",
			"expected", fen,
			"got", res)
	}

	fen = "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
	game = chess.Load(fen)
	res = chess.Save(game)
	if res != fen {
		t.Error(
			"For", "FEN 2",
			"expected", fen,
			"got", res)
	}

	fen = "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"
	game = chess.Load(fen)
	res = chess.Save(game)
	if res != fen {
		t.Error(
			"For", "FEN 3",
			"expected", fen,
			"got", res)
	}

	fen = "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"
	game = chess.Load(fen)
	res = chess.Save(game)
	if res != fen {
		t.Error(
			"For", "FEN 4",
			"expected", fen,
			"got", res)
	}
}
