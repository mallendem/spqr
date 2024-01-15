// Code generated by goyacc -o gram.go -p yy gram.y. DO NOT EDIT.

//line gram.y:3
package spqrparser

import __yyfmt__ "fmt"

//line gram.y:3

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"strings"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

//line gram.y:24
type yySymType struct {
	yys      int
	str      string
	byte     byte
	bytes    []byte
	integer  int
	uinteger uint
	bool     bool
	empty    struct{}

	set       *Set
	statement Statement
	show      *Show

	drop   *Drop
	create *Create

	kill   *Kill
	lock   *Lock
	unlock *Unlock

	ds            *DataspaceDefinition
	kr            *KeyRangeDefinition
	shard         *ShardDefinition
	sharding_rule *ShardingRuleDefinition

	register_router   *RegisterRouter
	unregister_router *UnregisterRouter

	split *SplitKeyRange
	move  *MoveKeyRange
	unite *UniteKeyRange

	shutdown *Shutdown
	listen   *Listen

	trace     *TraceStmt
	stoptrace *StopTraceStmt

	attach *AttachTable

	entrieslist []ShardingRuleEntry
	shruleEntry ShardingRuleEntry

	sharding_rule_selector *ShardingRuleSelector
	key_range_selector     *KeyRangeSelector
	dataspace_selector     *DataspaceSelector

	colref ColumnRef
	where  WhereClauseNode
}

const IDENT = 57346
const COMMAND = 57347
const SHOW = 57348
const KILL = 57349
const WHERE = 57350
const OR = 57351
const AND = 57352
const TEQ = 57353
const SCONST = 57354
const ICONST = 57355
const TSEMICOLON = 57356
const TOPENBR = 57357
const TCLOSEBR = 57358
const SHUTDOWN = 57359
const LISTEN = 57360
const REGISTER = 57361
const UNREGISTER = 57362
const ROUTER = 57363
const ROUTE = 57364
const CREATE = 57365
const ADD = 57366
const DROP = 57367
const LOCK = 57368
const UNLOCK = 57369
const SPLIT = 57370
const MOVE = 57371
const COMPOSE = 57372
const SET = 57373
const CASCADE = 57374
const ATTACH = 57375
const SHARDING = 57376
const COLUMN = 57377
const TABLE = 57378
const HASH = 57379
const FUNCTION = 57380
const KEY = 57381
const RANGE = 57382
const DATASPACE = 57383
const SHARDS = 57384
const KEY_RANGES = 57385
const ROUTERS = 57386
const SHARD = 57387
const HOST = 57388
const SHARDING_RULES = 57389
const RULE = 57390
const COLUMNS = 57391
const VERSION = 57392
const BY = 57393
const FROM = 57394
const TO = 57395
const WITH = 57396
const UNITE = 57397
const ALL = 57398
const ADDRESS = 57399
const FOR = 57400
const CLIENT = 57401
const IDENTITY = 57402
const MURMUR = 57403
const CITY = 57404
const START = 57405
const STOP = 57406
const TRACE = 57407
const MESSAGES = 57408
const OP = 57409

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"COMMAND",
	"SHOW",
	"KILL",
	"WHERE",
	"OR",
	"AND",
	"TEQ",
	"SCONST",
	"ICONST",
	"TSEMICOLON",
	"TOPENBR",
	"TCLOSEBR",
	"SHUTDOWN",
	"LISTEN",
	"REGISTER",
	"UNREGISTER",
	"ROUTER",
	"ROUTE",
	"CREATE",
	"ADD",
	"DROP",
	"LOCK",
	"UNLOCK",
	"SPLIT",
	"MOVE",
	"COMPOSE",
	"SET",
	"CASCADE",
	"ATTACH",
	"SHARDING",
	"COLUMN",
	"TABLE",
	"HASH",
	"FUNCTION",
	"KEY",
	"RANGE",
	"DATASPACE",
	"SHARDS",
	"KEY_RANGES",
	"ROUTERS",
	"SHARD",
	"HOST",
	"SHARDING_RULES",
	"RULE",
	"COLUMNS",
	"VERSION",
	"BY",
	"FROM",
	"TO",
	"WITH",
	"UNITE",
	"ALL",
	"ADDRESS",
	"FOR",
	"CLIENT",
	"IDENTITY",
	"MURMUR",
	"CITY",
	"START",
	"STOP",
	"TRACE",
	"MESSAGES",
	"OP",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line gram.y:693

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 224

var yyAct = [...]uint8{
	124, 156, 70, 121, 60, 97, 132, 108, 131, 29,
	30, 114, 88, 54, 53, 181, 182, 183, 69, 134,
	32, 31, 36, 37, 158, 128, 22, 21, 26, 27,
	28, 33, 34, 135, 25, 86, 38, 81, 87, 112,
	81, 102, 158, 195, 81, 81, 80, 194, 81, 84,
	193, 192, 176, 175, 166, 165, 81, 129, 35, 101,
	137, 100, 93, 151, 90, 82, 23, 24, 141, 113,
	173, 154, 134, 68, 99, 45, 189, 63, 103, 104,
	106, 174, 94, 107, 110, 89, 135, 83, 188, 118,
	117, 119, 116, 115, 56, 117, 105, 92, 85, 160,
	46, 125, 126, 127, 111, 47, 109, 45, 79, 81,
	136, 48, 120, 59, 138, 130, 142, 139, 57, 91,
	61, 191, 62, 64, 190, 44, 147, 41, 74, 75,
	76, 152, 187, 43, 153, 161, 162, 42, 157, 155,
	163, 109, 186, 164, 167, 78, 77, 168, 52, 40,
	49, 170, 98, 55, 171, 93, 51, 172, 144, 96,
	50, 144, 157, 146, 145, 81, 146, 145, 177, 81,
	169, 72, 178, 66, 179, 72, 122, 39, 184, 71,
	73, 149, 185, 71, 140, 1, 19, 18, 150, 17,
	16, 15, 13, 196, 197, 198, 199, 14, 200, 201,
	202, 203, 9, 10, 180, 159, 133, 20, 6, 5,
	4, 3, 8, 12, 11, 7, 67, 65, 58, 2,
	123, 148, 143, 95,
}

var yyPact = [...]int16{
	3, -1000, 135, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 66, 66, -51, -52, 34, 79, 38, 38, 169,
	14, 167, -1000, 38, 38, 38, 125, 124, 72, -1000,
	-1000, -1000, -1000, -1000, -1000, 165, 17, 47, 44, -1000,
	-1000, -1000, -1000, -21, -54, -1000, -1000, 45, -1000, 16,
	87, 41, -1000, 42, -1000, 151, -1000, 139, 139, -1000,
	-1000, -1000, -1000, -1000, 9, 6, -13, 165, 40, 165,
	-1000, -1000, 105, 52, -15, 23, -55, 139, -1000, 36,
	33, -1000, 80, -1000, 165, -1000, 161, -1000, -1000, -1000,
	165, 165, 165, -32, -1000, -1000, 4, 70, 37, 165,
	8, 171, 22, 167, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 157, 161, 177, -1000, 12, -1000, -1000, 167, 30,
	37, -16, -1000, 62, 165, 165, -1000, 171, 2, 1,
	-1000, 167, -1000, 161, -1000, -1000, -1000, 154, 167, -1000,
	-1000, 167, -1000, -1000, 165, -16, -1000, -1000, 29, -1000,
	43, -1000, -1000, 0, -1, 167, 139, -1000, 157, -1000,
	-1000, -1000, -1000, 165, -45, 167, 139, 120, 110, -1000,
	-1000, -1000, 51, 39, 102, 99, -2, -3, -1000, -1000,
	-6, -10, 165, 165, 165, 165, -34, -34, -34, -34,
	-1000, -1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 223, 3, 222, 221, 220, 2, 0, 5, 219,
	218, 94, 4, 217, 216, 215, 214, 213, 212, 211,
	210, 209, 208, 207, 127, 137, 133, 125, 8, 6,
	7, 206, 205, 204, 1, 203, 202, 197, 192, 191,
	190, 189, 187, 186, 185, 177,
}

var yyR1 = [...]int8{
	0, 44, 45, 45, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 8, 6, 6, 6, 7, 3, 3, 3,
	4, 4, 5, 2, 2, 2, 1, 1, 13, 14,
	15, 18, 18, 18, 18, 18, 18, 18, 18, 19,
	19, 19, 19, 21, 21, 22, 23, 20, 20, 20,
	20, 16, 36, 24, 25, 25, 28, 28, 29, 30,
	30, 31, 31, 33, 33, 33, 32, 32, 34, 34,
	26, 26, 26, 26, 27, 27, 35, 10, 11, 12,
	39, 17, 17, 40, 41, 38, 37, 42, 43, 43,
}

var yyR2 = [...]int8{
	0, 2, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 3, 3, 3, 0, 2, 1, 1,
	2, 2, 4, 2, 4, 2, 3, 3, 4, 2,
	2, 2, 2, 4, 4, 3, 5, 2, 2, 2,
	2, 3, 2, 2, 6, 5, 1, 2, 2, 2,
	0, 2, 2, 1, 2, 2, 3, 0, 3, 0,
	11, 11, 10, 10, 5, 4, 2, 3, 3, 2,
	6, 3, 3, 4, 4, 2, 1, 5, 3, 3,
}

var yyChk = [...]int16{
	-1000, -44, -9, -19, -20, -21, -22, -15, -18, -36,
	-35, -16, -17, -38, -37, -39, -40, -41, -42, -43,
	-23, 24, 23, 63, 64, 31, 25, 26, 27, 6,
	7, 18, 17, 28, 29, 55, 19, 20, 33, -45,
	14, -24, -25, -26, -27, 41, 34, 39, 45, -24,
	-25, -26, -27, 65, 65, -24, -11, 39, -10, 34,
	-12, 41, -11, 39, -11, -13, 4, -14, 59, 4,
	-6, 12, 4, 13, -11, -11, -11, 21, 21, 36,
	-7, 4, 48, 40, -7, 54, 56, 59, 66, 40,
	48, 32, 56, -7, 40, -1, 8, -8, 13, -8,
	52, 53, 54, -7, -7, 56, -7, -7, -30, 36,
	-7, 52, 54, 46, 66, -8, 56, -7, 56, -7,
	32, -2, 15, -5, -7, -7, -7, -7, 57, 53,
	-30, -28, -29, -31, 35, 49, -7, 52, -6, -8,
	13, 46, -6, -3, 4, 10, 9, -2, -4, 4,
	11, 51, -6, -12, 41, -28, -34, -29, 58, -32,
	37, -7, -7, -6, -8, 53, 53, -6, -2, 16,
	-6, -6, -34, 41, 38, 53, 53, -6, -8, -7,
	-33, 60, 61, 62, -6, -8, 22, 22, 37, 37,
	22, 22, 53, 53, 53, 53, -7, -7, -7, -7,
	-34, -34, -34, -34,
}

var yyDef = [...]int8{
	0, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 96, 0, 0, 0, 0, 0, 0, 1,
	3, 49, 50, 51, 52, 0, 0, 0, 0, 57,
	58, 59, 60, 0, 0, 40, 41, 0, 43, 0,
	45, 0, 62, 0, 86, 36, 38, 0, 0, 39,
	95, 23, 24, 25, 0, 0, 0, 0, 0, 0,
	63, 26, 70, 0, 0, 0, 0, 0, 55, 0,
	0, 47, 46, 89, 0, 61, 0, 91, 22, 92,
	0, 0, 0, 0, 98, 99, 0, 70, 0, 0,
	0, 0, 0, 0, 53, 54, 42, 88, 44, 87,
	48, 37, 0, 0, 32, 0, 93, 94, 0, 0,
	0, 79, 66, 77, 0, 0, 69, 0, 0, 0,
	22, 0, 85, 0, 27, 28, 29, 0, 0, 30,
	31, 0, 97, 56, 0, 79, 65, 67, 0, 68,
	0, 71, 72, 0, 0, 0, 0, 84, 35, 33,
	34, 90, 64, 0, 0, 0, 0, 0, 0, 78,
	76, 73, 0, 0, 0, 0, 0, 0, 74, 75,
	0, 0, 0, 0, 0, 0, 79, 79, 79, 79,
	82, 83, 80, 81,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:183
		{
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:184
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:189
		{
			setParseTree(yylex, yyDollar[1].create)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:193
		{
			setParseTree(yylex, yyDollar[1].create)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:197
		{
			setParseTree(yylex, yyDollar[1].trace)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:201
		{
			setParseTree(yylex, yyDollar[1].stoptrace)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:205
		{
			setParseTree(yylex, yyDollar[1].set)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:209
		{
			setParseTree(yylex, yyDollar[1].drop)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:213
		{
			setParseTree(yylex, yyDollar[1].lock)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:217
		{
			setParseTree(yylex, yyDollar[1].unlock)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:221
		{
			setParseTree(yylex, yyDollar[1].show)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:225
		{
			setParseTree(yylex, yyDollar[1].kill)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:229
		{
			setParseTree(yylex, yyDollar[1].listen)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:233
		{
			setParseTree(yylex, yyDollar[1].shutdown)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:237
		{
			setParseTree(yylex, yyDollar[1].split)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:241
		{
			setParseTree(yylex, yyDollar[1].move)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:245
		{
			setParseTree(yylex, yyDollar[1].unite)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:249
		{
			setParseTree(yylex, yyDollar[1].register_router)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:253
		{
			setParseTree(yylex, yyDollar[1].unregister_router)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:257
		{
			setParseTree(yylex, yyDollar[1].attach)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:262
		{
			yyVAL.uinteger = uint(yyDollar[1].uinteger)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:267
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:271
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:273
		{
			yyVAL.str = strconv.Itoa(int(yyDollar[1].uinteger))
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:278
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:284
		{
			yyVAL.str = yyDollar[1].str
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:286
		{
			yyVAL.str = "AND"
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:288
		{
			yyVAL.str = "OR"
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:293
		{
			yyVAL.str = yyDollar[1].str
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:295
		{
			yyVAL.str = "="
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:301
		{
			yyVAL.colref = ColumnRef{
				ColName: yyDollar[1].str,
			}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:309
		{
			yyVAL.where = yyDollar[2].where
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:312
		{
			yyVAL.where = WhereClauseLeaf{
				ColRef: yyDollar[1].colref,
				Op:     yyDollar[2].str,
				Value:  yyDollar[3].str,
			}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:320
		{
			yyVAL.where = WhereClauseOp{
				Op:    yyDollar[2].str,
				Left:  yyDollar[1].where,
				Right: yyDollar[3].where,
			}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:330
		{
			yyVAL.where = WhereClauseEmpty{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:334
		{
			yyVAL.where = yyDollar[2].where
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:341
		{
			switch v := strings.ToLower(string(yyDollar[1].str)); v {
			case DatabasesStr, RoutersStr, PoolsStr, ShardsStr, BackendConnectionsStr, KeyRangesStr, ShardingRules, ClientsStr, StatusStr, DataspacesStr, VersionStr:
				yyVAL.str = v
			default:
				yyVAL.str = UnsupportedStr
			}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:352
		{
			switch v := string(yyDollar[1].str); v {
			case ClientStr:
				yyVAL.str = v
			default:
				yyVAL.str = "unsupp"
			}
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:363
		{
			yyVAL.set = &Set{Element: yyDollar[2].ds}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:370
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].key_range_selector}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:374
		{
			yyVAL.drop = &Drop{Element: &KeyRangeSelector{KeyRangeID: `*`}}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:378
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].sharding_rule_selector}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:382
		{
			yyVAL.drop = &Drop{Element: &ShardingRuleSelector{ID: `*`}}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:386
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].dataspace_selector, CascadeDelete: false}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:390
		{
			yyVAL.drop = &Drop{Element: &DataspaceSelector{ID: `*`}, CascadeDelete: false}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:394
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].dataspace_selector, CascadeDelete: true}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:398
		{
			yyVAL.drop = &Drop{Element: &DataspaceSelector{ID: `*`}, CascadeDelete: true}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:404
		{
			yyVAL.create = &Create{Element: yyDollar[2].ds}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:409
		{
			yyVAL.create = &Create{Element: yyDollar[2].sharding_rule}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:414
		{
			yyVAL.create = &Create{Element: yyDollar[2].kr}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:418
		{
			yyVAL.create = &Create{Element: yyDollar[2].shard}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:425
		{
			yyVAL.trace = &TraceStmt{All: true}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:428
		{
			yyVAL.trace = &TraceStmt{
				Client: yyDollar[4].uinteger,
			}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:436
		{
			yyVAL.stoptrace = &StopTraceStmt{}
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:443
		{
			yyVAL.attach = &AttachTable{
				Table:     yyDollar[3].str,
				Dataspace: yyDollar[5].dataspace_selector,
			}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:453
		{
			yyVAL.create = &Create{Element: yyDollar[2].ds}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:458
		{
			yyVAL.create = &Create{Element: yyDollar[2].sharding_rule}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:463
		{
			yyVAL.create = &Create{Element: yyDollar[2].kr}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:467
		{
			yyVAL.create = &Create{Element: yyDollar[2].shard}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:474
		{
			yyVAL.show = &Show{Cmd: yyDollar[2].str, Where: yyDollar[3].where}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:480
		{
			yyVAL.lock = &Lock{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:488
		{
			yyVAL.ds = &DataspaceDefinition{ID: yyDollar[2].str}
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:494
		{
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: yyDollar[3].str, TableName: yyDollar[4].str, Entries: yyDollar[5].entrieslist, Dataspace: yyDollar[6].str}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:499
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: "shrule" + str, TableName: yyDollar[3].str, Entries: yyDollar[4].entrieslist, Dataspace: yyDollar[5].str}
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:508
		{
			yyVAL.entrieslist = make([]ShardingRuleEntry, 0)
			yyVAL.entrieslist = append(yyVAL.entrieslist, yyDollar[1].shruleEntry)
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:514
		{
			yyVAL.entrieslist = append(yyDollar[1].entrieslist, yyDollar[2].shruleEntry)
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:520
		{
			yyVAL.shruleEntry = ShardingRuleEntry{
				Column:       yyDollar[1].str,
				HashFunction: yyDollar[2].str,
			}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:529
		{
			yyVAL.str = yyDollar[2].str
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:532
		{
			yyVAL.str = ""
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:536
		{
			yyVAL.str = yyDollar[2].str
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:541
		{
			yyVAL.str = yyDollar[2].str
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:547
		{
			yyVAL.str = "identity"
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:549
		{
			yyVAL.str = "murmur"
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:551
		{
			yyVAL.str = "city"
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:557
		{
			yyVAL.str = yyDollar[3].str
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:560
		{
			yyVAL.str = ""
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:563
		{
			yyVAL.str = yyDollar[3].str
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:566
		{
			yyVAL.str = "default"
		}
	case 80:
		yyDollar = yyS[yypt-11 : yypt+1]
//line gram.y:571
		{
			yyVAL.kr = &KeyRangeDefinition{LowerBound: []byte(yyDollar[5].str), UpperBound: []byte(yyDollar[7].str), ShardID: yyDollar[10].str, KeyRangeID: yyDollar[3].str, Dataspace: yyDollar[11].str}
		}
	case 81:
		yyDollar = yyS[yypt-11 : yypt+1]
//line gram.y:575
		{
			yyVAL.kr = &KeyRangeDefinition{LowerBound: []byte(strconv.FormatUint(uint64(yyDollar[5].uinteger), 10)), UpperBound: []byte(strconv.FormatUint(uint64(yyDollar[7].uinteger), 10)), ShardID: yyDollar[10].str, KeyRangeID: yyDollar[3].str, Dataspace: yyDollar[11].str}
		}
	case 82:
		yyDollar = yyS[yypt-10 : yypt+1]
//line gram.y:579
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.kr = &KeyRangeDefinition{LowerBound: []byte(yyDollar[4].str), UpperBound: []byte(yyDollar[6].str), ShardID: yyDollar[9].str, KeyRangeID: "kr" + str, Dataspace: yyDollar[10].str}
		}
	case 83:
		yyDollar = yyS[yypt-10 : yypt+1]
//line gram.y:587
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.kr = &KeyRangeDefinition{LowerBound: []byte(strconv.FormatUint(uint64(yyDollar[4].uinteger), 10)), UpperBound: []byte(strconv.FormatUint(uint64(yyDollar[6].uinteger), 10)), ShardID: yyDollar[9].str, KeyRangeID: "kr" + str, Dataspace: yyDollar[10].str}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:598
		{
			yyVAL.shard = &ShardDefinition{Id: yyDollar[2].str, Hosts: []string{yyDollar[5].str}}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:603
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.shard = &ShardDefinition{Id: "shard" + str, Hosts: []string{yyDollar[4].str}}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:614
		{
			yyVAL.unlock = &Unlock{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:620
		{
			yyVAL.sharding_rule_selector = &ShardingRuleSelector{ID: yyDollar[3].str}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:626
		{
			yyVAL.key_range_selector = &KeyRangeSelector{KeyRangeID: yyDollar[3].str}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:632
		{
			yyVAL.dataspace_selector = &DataspaceSelector{ID: yyDollar[2].str}
		}
	case 90:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:638
		{
			yyVAL.split = &SplitKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, KeyRangeFromID: yyDollar[4].str, Border: []byte(yyDollar[6].str)}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:644
		{
			yyVAL.kill = &Kill{Cmd: yyDollar[2].str, Target: yyDollar[3].uinteger}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:647
		{
			yyVAL.kill = &Kill{Cmd: "client", Target: yyDollar[3].uinteger}
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:653
		{
			yyVAL.move = &MoveKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, DestShardID: yyDollar[4].str}
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:659
		{
			yyVAL.unite = &UniteKeyRange{KeyRangeIDL: yyDollar[2].key_range_selector.KeyRangeID, KeyRangeIDR: yyDollar[4].str}
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:665
		{
			yyVAL.listen = &Listen{addr: yyDollar[2].str}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:671
		{
			yyVAL.shutdown = &Shutdown{}
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:679
		{
			yyVAL.register_router = &RegisterRouter{ID: yyDollar[3].str, Addr: yyDollar[5].str}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:685
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: yyDollar[3].str}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:690
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: `*`}
		}
	}
	goto yystack /* stack new state and value */
}
