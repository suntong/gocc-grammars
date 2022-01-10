// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"fmt"

	"token"
)

type ActionTable [NumStates]ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (a ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", a.Accept, a.Ignore)
}

var ActTab = ActionTable{
	ActionRow{ // S0
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S1
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S2
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S3
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S4
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S5
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S6
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S7
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S8
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S9
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S10
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S11
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S12
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S13
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S14
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S15
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S16
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S17
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S18
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S19
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S20
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S21
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S22
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S23
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S24
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S25
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S26
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S27
		Accept: 2,
		Ignore: "",
	},
}