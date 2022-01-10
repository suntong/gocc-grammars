// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 11: // ['\v','\v']
			return 1
		case r == 12: // ['\f','\f']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 60: // ['<','<']
			return 2
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case r == 95: // ['_','_']
			return 3
		case 97 <= r && r <= 122: // ['a','z']
			return 3
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 4
		case r == 62: // ['>','>']
			return 5
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case r == 95: // ['_','_']
			return 3
		case 97 <= r && r <= 122: // ['a','z']
			return 3
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 4
		case r == 62: // ['>','>']
			return 5
		case 65 <= r && r <= 90: // ['A','Z']
			return 3
		case r == 95: // ['_','_']
			return 3
		case 97 <= r && r <= 122: // ['a','z']
			return 3
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case r == 58: // [':',':']
			return 6
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 58: // [':',':']
			return 7
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 8
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case r == 40: // ['(','(']
			return 9
		case r == 60: // ['<','<']
			return 10
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 91: // ['[','[']
			return 12
		case r == 95: // ['_','_']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 123: // ['{','{']
			return 13
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 14
		case r == 95: // ['_','_']
			return 14
		case 97 <= r && r <= 122: // ['a','z']
			return 14
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 15
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 95: // ['_','_']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 124: // ['|','|']
			return 16
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 62: // ['>','>']
			return 18
		case 65 <= r && r <= 90: // ['A','Z']
			return 14
		case r == 95: // ['_','_']
			return 14
		case 97 <= r && r <= 122: // ['a','z']
			return 14
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 15
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 95: // ['_','_']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 124: // ['|','|']
			return 16
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 40: // ['(','(']
			return 19
		case r == 60: // ['<','<']
			return 20
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 91: // ['[','[']
			return 22
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		case r == 123: // ['{','{']
			return 23
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 17
		case r == 62: // ['>','>']
			return 18
		case 65 <= r && r <= 90: // ['A','Z']
			return 14
		case r == 95: // ['_','_']
			return 14
		case 97 <= r && r <= 122: // ['a','z']
			return 14
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case r == 40: // ['(','(']
			return 9
		case r == 60: // ['<','<']
			return 10
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 91: // ['[','[']
			return 12
		case r == 95: // ['_','_']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		case r == 123: // ['{','{']
			return 13
		case r == 124: // ['|','|']
			return 16
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 65 <= r && r <= 90: // ['A','Z']
			return 24
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 24
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		case r == 124: // ['|','|']
			return 16
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case r == 62: // ['>','>']
			return 27
		case 65 <= r && r <= 90: // ['A','Z']
			return 24
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 24
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		case r == 124: // ['|','|']
			return 16
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 26
		case r == 62: // ['>','>']
			return 27
		case 65 <= r && r <= 90: // ['A','Z']
			return 24
		case r == 95: // ['_','_']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 24
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 40: // ['(','(']
			return 19
		case r == 60: // ['<','<']
			return 20
		case 65 <= r && r <= 90: // ['A','Z']
			return 21
		case r == 91: // ['[','[']
			return 22
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 122: // ['a','z']
			return 21
		case r == 123: // ['{','{']
			return 23
		case r == 124: // ['|','|']
			return 16
		}
		return NoState
	},
}
