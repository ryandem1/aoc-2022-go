package day9

import (
	"github.com/ryandem1/aoc_2022_go/common"
	"strconv"
	"strings"
)

// readMotions will convert the raw string line input into ropeMotion objects and send them through a channel for use
// in the puzzle
func readMotions() chan *ropeMotion {
	motions := make(chan *ropeMotion)

	go func() {
		for line := range common.ReadLinesFromFile("day9") {
			fields := strings.Fields(line)

			direction := map[string]motionDirection{
				"U": up,
				"D": down,
				"L": left,
				"R": right,
			}[fields[0]]
			amount, err := strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}

			motion := &ropeMotion{
				direction: direction,
				amount:    amount,
			}
			motions <- motion
		}
		close(motions)
	}()
	return motions
}

// moveHeadPosition will apply a motion on a headPos according to the logic in the puzzle. Will output the new position
// of the head with the applied motion
func moveHeadPosition(headPos common.Coords2D, motion *ropeMotion) (movedHeadPos common.Coords2D) {
	switch motion.direction {
	case up:
		movedHeadPos.X = headPos.X
		movedHeadPos.Y = headPos.Y + motion.amount
	case down:
		movedHeadPos.X = headPos.X
		movedHeadPos.Y = headPos.Y - motion.amount
	case left:
		movedHeadPos.X = headPos.X - motion.amount
		movedHeadPos.Y = headPos.Y
	case right:
		movedHeadPos.X = headPos.X + motion.amount
		movedHeadPos.Y = headPos.Y
	}
	return movedHeadPos
}

// moveTailPosition will apply the corresponding tail move according to the headPos. While the headPos moves
// according to the motion, the tail position will follow the head according to the puzzle logic.
func moveTailPosition(tailPos common.Coords2D, headPos common.Coords2D) (movedTailPos common.Coords2D) {
	return movedTailPos
}

// applyMotion will update both the head and tail of a rope according to a motion. Will keep track of the positions
// visited
func (rope *bridgeRope) applyMotion(motion *ropeMotion) {
	rope.headPos = moveHeadPosition(rope.headPos, motion)
	rope.tailPos = moveTailPosition(rope.tailPos, rope.headPos)
}
