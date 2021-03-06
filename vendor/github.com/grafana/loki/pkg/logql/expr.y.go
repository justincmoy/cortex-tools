// Code generated by goyacc -p expr -o pkg/logql/expr.y.go pkg/logql/expr.y. DO NOT EDIT.

package logql

import __yyfmt__ "fmt"


import (
	"github.com/grafana/loki/pkg/logql/log"
	"github.com/prometheus/prometheus/pkg/labels"
	"time"
)

type exprSymType struct {
	yys                   int
	Expr                  Expr
	Filter                labels.MatchType
	Grouping              *grouping
	Labels                []string
	LogExpr               LogSelectorExpr
	LogRangeExpr          *logRange
	Matcher               *labels.Matcher
	Matchers              []*labels.Matcher
	RangeAggregationExpr  SampleExpr
	RangeOp               string
	ConvOp                string
	Selector              []*labels.Matcher
	VectorAggregationExpr SampleExpr
	MetricExpr            SampleExpr
	VectorOp              string
	BinOpExpr             SampleExpr
	binOp                 string
	bytes                 uint64
	str                   string
	duration              time.Duration
	LiteralExpr           *literalExpr
	BinOpModifier         BinOpOptions
	LabelParser           *labelParserExpr
	LineFilters           *lineFilterExpr
	PipelineExpr          MultiStageExpr
	PipelineStage         StageExpr
	BytesFilter           log.LabelFilterer
	NumberFilter          log.LabelFilterer
	DurationFilter        log.LabelFilterer
	LabelFilter           log.LabelFilterer
	UnitFilter            log.LabelFilterer
	LineFormatExpr        *lineFmtExpr
	LabelFormatExpr       *labelFmtExpr
	LabelFormat           log.LabelFmt
	LabelsFormat          []log.LabelFmt
	UnwrapExpr            *unwrapExpr
}

const BYTES = 57346
const IDENTIFIER = 57347
const STRING = 57348
const NUMBER = 57349
const DURATION = 57350
const RANGE = 57351
const MATCHERS = 57352
const LABELS = 57353
const EQ = 57354
const RE = 57355
const NRE = 57356
const OPEN_BRACE = 57357
const CLOSE_BRACE = 57358
const OPEN_BRACKET = 57359
const CLOSE_BRACKET = 57360
const COMMA = 57361
const DOT = 57362
const PIPE_MATCH = 57363
const PIPE_EXACT = 57364
const OPEN_PARENTHESIS = 57365
const CLOSE_PARENTHESIS = 57366
const BY = 57367
const WITHOUT = 57368
const COUNT_OVER_TIME = 57369
const RATE = 57370
const SUM = 57371
const AVG = 57372
const MAX = 57373
const MIN = 57374
const COUNT = 57375
const STDDEV = 57376
const STDVAR = 57377
const BOTTOMK = 57378
const TOPK = 57379
const BYTES_OVER_TIME = 57380
const BYTES_RATE = 57381
const BOOL = 57382
const JSON = 57383
const REGEXP = 57384
const LOGFMT = 57385
const PIPE = 57386
const LINE_FMT = 57387
const LABEL_FMT = 57388
const UNWRAP = 57389
const AVG_OVER_TIME = 57390
const SUM_OVER_TIME = 57391
const MIN_OVER_TIME = 57392
const MAX_OVER_TIME = 57393
const STDVAR_OVER_TIME = 57394
const STDDEV_OVER_TIME = 57395
const QUANTILE_OVER_TIME = 57396
const BYTES_CONV = 57397
const DURATION_CONV = 57398
const DURATION_SECONDS_CONV = 57399
const OR = 57400
const AND = 57401
const UNLESS = 57402
const CMP_EQ = 57403
const NEQ = 57404
const LT = 57405
const LTE = 57406
const GT = 57407
const GTE = 57408
const ADD = 57409
const SUB = 57410
const MUL = 57411
const DIV = 57412
const MOD = 57413
const POW = 57414

var exprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BYTES",
	"IDENTIFIER",
	"STRING",
	"NUMBER",
	"DURATION",
	"RANGE",
	"MATCHERS",
	"LABELS",
	"EQ",
	"RE",
	"NRE",
	"OPEN_BRACE",
	"CLOSE_BRACE",
	"OPEN_BRACKET",
	"CLOSE_BRACKET",
	"COMMA",
	"DOT",
	"PIPE_MATCH",
	"PIPE_EXACT",
	"OPEN_PARENTHESIS",
	"CLOSE_PARENTHESIS",
	"BY",
	"WITHOUT",
	"COUNT_OVER_TIME",
	"RATE",
	"SUM",
	"AVG",
	"MAX",
	"MIN",
	"COUNT",
	"STDDEV",
	"STDVAR",
	"BOTTOMK",
	"TOPK",
	"BYTES_OVER_TIME",
	"BYTES_RATE",
	"BOOL",
	"JSON",
	"REGEXP",
	"LOGFMT",
	"PIPE",
	"LINE_FMT",
	"LABEL_FMT",
	"UNWRAP",
	"AVG_OVER_TIME",
	"SUM_OVER_TIME",
	"MIN_OVER_TIME",
	"MAX_OVER_TIME",
	"STDVAR_OVER_TIME",
	"STDDEV_OVER_TIME",
	"QUANTILE_OVER_TIME",
	"BYTES_CONV",
	"DURATION_CONV",
	"DURATION_SECONDS_CONV",
	"OR",
	"AND",
	"UNLESS",
	"CMP_EQ",
	"NEQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"POW",
}
var exprStatenames = [...]string{}

const exprEofCode = 1
const exprErrCode = 2
const exprInitialStackSize = 16


var exprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const exprPrivate = 57344

const exprLast = 397

var exprAct = [...]int{

	70, 171, 53, 153, 145, 4, 179, 100, 63, 2,
	52, 45, 61, 56, 5, 217, 120, 214, 235, 66,
	14, 40, 41, 42, 43, 44, 45, 250, 11, 42,
	43, 44, 45, 253, 235, 213, 6, 76, 257, 213,
	17, 18, 28, 29, 31, 32, 30, 33, 34, 35,
	36, 19, 20, 214, 242, 91, 116, 118, 119, 245,
	94, 21, 22, 23, 24, 25, 26, 27, 92, 214,
	214, 71, 72, 224, 214, 124, 155, 118, 119, 176,
	15, 16, 167, 122, 129, 167, 130, 131, 132, 133,
	134, 135, 136, 137, 138, 139, 140, 141, 142, 143,
	69, 167, 71, 72, 232, 111, 117, 221, 150, 46,
	47, 50, 51, 48, 49, 40, 41, 42, 43, 44,
	45, 110, 162, 168, 248, 161, 156, 159, 160, 157,
	158, 128, 178, 172, 127, 181, 215, 225, 174, 59,
	175, 59, 227, 126, 68, 225, 57, 58, 57, 58,
	226, 182, 183, 184, 37, 38, 39, 46, 47, 50,
	51, 48, 49, 40, 41, 42, 43, 44, 45, 209,
	106, 173, 211, 186, 216, 91, 219, 222, 94, 177,
	169, 212, 115, 223, 122, 220, 210, 60, 103, 60,
	228, 38, 39, 46, 47, 50, 51, 48, 49, 40,
	41, 42, 43, 44, 45, 121, 97, 99, 98, 187,
	104, 105, 215, 11, 233, 91, 236, 59, 113, 234,
	74, 123, 244, 91, 57, 58, 170, 243, 125, 11,
	73, 59, 112, 247, 256, 114, 11, 123, 57, 58,
	252, 218, 251, 249, 6, 106, 254, 173, 17, 18,
	28, 29, 31, 32, 30, 33, 34, 35, 36, 19,
	20, 173, 106, 103, 241, 60, 238, 239, 240, 21,
	22, 23, 24, 25, 26, 27, 147, 59, 170, 60,
	103, 230, 231, 59, 57, 58, 59, 166, 15, 16,
	57, 58, 165, 57, 58, 207, 75, 106, 208, 206,
	106, 204, 164, 106, 205, 203, 192, 173, 164, 193,
	191, 147, 163, 173, 147, 103, 55, 147, 103, 185,
	189, 103, 163, 190, 188, 60, 201, 151, 255, 202,
	200, 60, 149, 144, 60, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 106,
	148, 146, 198, 148, 146, 199, 197, 146, 195, 3,
	229, 196, 194, 154, 101, 109, 62, 103, 65, 246,
	180, 67, 67, 154, 152, 96, 95, 54, 107, 102,
	108, 93, 10, 9, 13, 97, 99, 98, 8, 104,
	105, 217, 237, 12, 7, 64, 1,
}
var exprPact = [...]int{

	13, -1000, 96, -1000, -1000, 272, 13, -1000, -1000, -1000,
	-1000, 366, 121, 77, -1000, 223, 213, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, 272, -1000, 125, 165, 359, -1000, -1000, -1000,
	-1000, 97, 81, 96, 216, 166, -1000, 44, 198, 221,
	120, 111, 108, -1000, -1000, 13, -1000, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, -1000, 327, -1000, 292, -1000, -1000, -1000, -1000, 326,
	-1000, -1000, -1000, 240, 321, 368, 64, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 367, -1000, 306, 296, 286, 281,
	99, 161, 269, 214, 55, 160, 13, 365, 365, 132,
	48, 48, -40, -40, -61, -61, -61, -61, -46, -46,
	-46, -46, -46, -46, -1000, 292, 240, 240, 240, -1000,
	295, -1000, 154, -1000, 197, 316, 302, 354, 348, 322,
	297, 291, -1000, -1000, -1000, -1000, -1000, -1000, 46, 214,
	263, 26, 127, 344, 217, 83, 46, 13, 49, 126,
	-1000, 118, 257, 292, 298, -1000, 358, 276, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	80, -27, 263, -1000, 240, -1000, 25, 211, 255, 30,
	203, -1000, -1000, 35, -1000, 364, -1000, -1000, -1000, -1000,
	-1000, -1000, 46, -27, 292, -1000, -1000, 101, -1000, -1000,
	-1000, -17, 233, 231, 9, 46, -1000, -1000, 323, -27,
	-32, -1000, -1000, 225, -1000, 14, -1000, -1000,
}
var exprPgo = [...]int{

	0, 396, 8, 13, 0, 6, 359, 5, 16, 7,
	395, 394, 393, 392, 14, 388, 384, 383, 382, 296,
	381, 10, 2, 380, 379, 378, 4, 377, 376, 375,
	3, 374, 1, 364,
}
var exprR1 = [...]int{

	0, 1, 2, 2, 7, 7, 7, 7, 7, 6,
	6, 6, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 32, 32, 32, 13,
	13, 13, 11, 11, 11, 11, 15, 15, 15, 15,
	15, 3, 3, 3, 3, 14, 14, 14, 10, 10,
	9, 9, 9, 9, 21, 21, 22, 22, 22, 22,
	22, 27, 27, 20, 20, 20, 28, 30, 30, 31,
	31, 31, 29, 26, 26, 26, 26, 26, 26, 26,
	26, 33, 33, 25, 25, 25, 25, 25, 25, 25,
	23, 23, 23, 23, 23, 23, 23, 24, 24, 24,
	24, 24, 24, 24, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 19,
	19, 18, 18, 18, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 5, 5, 4, 4,
}
var exprR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	2, 3, 2, 4, 3, 5, 3, 5, 3, 5,
	4, 6, 3, 4, 3, 2, 3, 6, 3, 1,
	1, 1, 4, 6, 5, 7, 4, 5, 5, 6,
	7, 1, 1, 1, 1, 3, 3, 3, 1, 3,
	3, 3, 3, 3, 1, 2, 1, 2, 2, 2,
	2, 2, 3, 1, 1, 2, 2, 3, 3, 1,
	3, 3, 2, 1, 1, 1, 3, 2, 3, 3,
	3, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 0,
	1, 1, 2, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 4, 4,
}
var exprChk = [...]int{

	-1000, -1, -2, -6, -7, -14, 23, -11, -15, -17,
	-18, 15, -12, -16, 7, 67, 68, 27, 28, 38,
	39, 48, 49, 50, 51, 52, 53, 54, 29, 30,
	33, 31, 32, 34, 35, 36, 37, 58, 59, 60,
	67, 68, 69, 70, 71, 72, 61, 62, 65, 66,
	63, 64, -21, -22, -27, 44, -3, 21, 22, 14,
	62, -7, -6, -2, -10, 2, -9, 5, 23, 23,
	-4, 25, 26, 7, 7, -19, 40, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -22, -3, -20, -26, -28, -29, 41, 43, 42,
	-9, -33, -24, 23, 45, 46, 5, -25, -23, 6,
	24, 24, 16, 2, 19, 16, 12, 62, 13, 14,
	-8, 7, -14, 23, -7, 7, 23, 23, 23, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, -2, -2, 6, -26, 59, 19, 58, 6,
	-26, 6, -31, -30, 5, 12, 62, 65, 66, 63,
	64, 61, -9, 6, 6, 6, 6, 2, 24, 19,
	9, -32, -21, 44, -14, -8, 24, 19, -7, -5,
	5, -5, -26, -26, -26, 24, 19, 12, 8, 4,
	7, 8, 4, 7, 8, 4, 7, 8, 4, 7,
	8, 4, 7, 8, 4, 7, 8, 4, 7, -4,
	-8, -32, -21, 9, 44, 9, -32, 47, 24, -32,
	-21, 24, -4, -7, 24, 19, 24, 24, -30, 2,
	5, 6, 24, -32, -26, 9, 5, -13, 55, 56,
	57, 9, 24, 24, -32, 24, 5, -4, 23, -32,
	44, 9, 9, 24, -4, 5, 9, 24,
}
var exprDef = [...]int{

	0, -2, 1, 2, 3, 9, 0, 4, 5, 6,
	7, 0, 0, 0, 121, 0, 0, 133, 134, 135,
	136, 137, 138, 139, 140, 141, 142, 143, 124, 125,
	126, 127, 128, 129, 130, 131, 132, 119, 119, 119,
	119, 119, 119, 119, 119, 119, 119, 119, 119, 119,
	119, 119, 10, 54, 56, 0, 0, 41, 42, 43,
	44, 3, 2, 0, 0, 0, 48, 0, 0, 0,
	0, 0, 0, 122, 123, 0, 120, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 55, 0, 57, 58, 59, 60, 63, 64, 0,
	73, 74, 75, 0, 0, 0, 0, 81, 82, 61,
	8, 11, 45, 46, 0, 47, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 121, 0, 0, 0, 104,
	105, 106, 107, 108, 109, 110, 111, 112, 113, 114,
	115, 116, 117, 118, 62, 77, 0, 0, 0, 65,
	0, 66, 72, 69, 0, 0, 0, 0, 0, 0,
	0, 0, 49, 50, 51, 52, 53, 25, 32, 0,
	12, 0, 0, 0, 0, 0, 36, 0, 3, 0,
	144, 0, 78, 79, 80, 76, 0, 0, 88, 95,
	102, 87, 94, 101, 83, 90, 97, 84, 91, 98,
	85, 92, 99, 86, 93, 100, 89, 96, 103, 34,
	0, 14, 22, 16, 0, 18, 0, 0, 0, 0,
	0, 24, 38, 3, 37, 0, 146, 147, 70, 71,
	67, 68, 33, 23, 28, 20, 26, 0, 29, 30,
	31, 13, 0, 0, 0, 39, 145, 35, 0, 15,
	0, 17, 19, 0, 40, 0, 21, 27,
}
var exprTok1 = [...]int{

	1,
}
var exprTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72,
}
var exprTok3 = [...]int{
	0,
}

var exprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}


/*	parser for yacc output	*/

var (
	exprDebug        = 0
	exprErrorVerbose = false
)

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

type exprParser interface {
	Parse(exprLexer) int
	Lookahead() int
}

type exprParserImpl struct {
	lval  exprSymType
	stack [exprInitialStackSize]exprSymType
	char  int
}

func (p *exprParserImpl) Lookahead() int {
	return p.char
}

func exprNewParser() exprParser {
	return &exprParserImpl{}
}

const exprFlag = -1000

func exprTokname(c int) string {
	if c >= 1 && c-1 < len(exprToknames) {
		if exprToknames[c-1] != "" {
			return exprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !exprErrorVerbose {
		return "syntax error"
	}

	for _, e := range exprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + exprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := exprPact[state]
	for tok := TOKSTART; tok-1 < len(exprToknames); tok++ {
		if n := base + tok; n >= 0 && n < exprLast && exprChk[exprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if exprDef[state] == -2 {
		i := 0
		for exprExca[i] != -1 || exprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; exprExca[i] >= 0; i += 2 {
			tok := exprExca[i]
			if tok < TOKSTART || exprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if exprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += exprTokname(tok)
	}
	return res
}

func exprlex1(lex exprLexer, lval *exprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = exprTok1[0]
		goto out
	}
	if char < len(exprTok1) {
		token = exprTok1[char]
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			token = exprTok2[char-exprPrivate]
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		token = exprTok3[i+0]
		if token == char {
			token = exprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = exprTok2[1] /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(token), uint(char))
	}
	return char, token
}

func exprParse(exprlex exprLexer) int {
	return exprNewParser().Parse(exprlex)
}

func (exprrcvr *exprParserImpl) Parse(exprlex exprLexer) int {
	var exprn int
	var exprVAL exprSymType
	var exprDollar []exprSymType
	_ = exprDollar // silence set and not used
	exprS := exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprrcvr.char = -1
	exprtoken := -1 // exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		exprstate = -1
		exprrcvr.char = -1
		exprtoken = -1
	}()
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprtoken), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = exprPact[exprstate]
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprrcvr.char < 0 {
		exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
	}
	exprn += exprtoken
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = exprAct[exprn]
	if exprChk[exprn] == exprtoken { /* valid shift */
		exprrcvr.char = -1
		exprtoken = -1
		exprVAL = exprrcvr.lval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = exprDef[exprstate]
	if exprn == -2 {
		if exprrcvr.char < 0 {
			exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && exprExca[xi+1] == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = exprExca[xi+0]
			if exprn < 0 || exprn == exprtoken {
				break
			}
		}
		exprn = exprExca[xi+1]
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error(exprErrorMessage(exprstate, exprtoken))
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = exprPact[exprS[exprp].yys] + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = exprAct[exprn] /* simulate a shift of "error" */
					if exprChk[exprstate] == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprtoken))
			}
			if exprtoken == exprEofCode {
				goto ret1
			}
			exprrcvr.char = -1
			exprtoken = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= exprR2[exprn]
	// exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if exprp+1 >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = exprR1[exprn]
	exprg := exprPgo[exprn]
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = exprAct[exprg]
	} else {
		exprstate = exprAct[exprj]
		if exprChk[exprstate] != -exprn {
			exprstate = exprAct[exprg]
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprlex.(*parser).expr = exprDollar[1].Expr
		}
	case 2:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Expr = exprDollar[1].LogExpr
		}
	case 3:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Expr = exprDollar[1].MetricExpr
		}
	case 4:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricExpr = exprDollar[1].RangeAggregationExpr
		}
	case 5:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricExpr = exprDollar[1].VectorAggregationExpr
		}
	case 6:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricExpr = exprDollar[1].BinOpExpr
		}
	case 7:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.MetricExpr = exprDollar[1].LiteralExpr
		}
	case 8:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.MetricExpr = exprDollar[2].MetricExpr
		}
	case 9:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LogExpr = newMatcherExpr(exprDollar[1].Selector)
		}
	case 10:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogExpr = newPipelineExpr(newMatcherExpr(exprDollar[1].Selector), exprDollar[2].PipelineExpr)
		}
	case 11:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogExpr = exprDollar[2].LogExpr
		}
	case 12:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newMatcherExpr(exprDollar[1].Selector), exprDollar[2].duration, nil)
		}
	case 13:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newMatcherExpr(exprDollar[2].Selector), exprDollar[4].duration, nil)
		}
	case 14:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newMatcherExpr(exprDollar[1].Selector), exprDollar[2].duration, exprDollar[3].UnwrapExpr)
		}
	case 15:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newMatcherExpr(exprDollar[2].Selector), exprDollar[4].duration, exprDollar[5].UnwrapExpr)
		}
	case 16:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newMatcherExpr(exprDollar[1].Selector), exprDollar[3].duration, exprDollar[2].UnwrapExpr)
		}
	case 17:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newMatcherExpr(exprDollar[2].Selector), exprDollar[5].duration, exprDollar[3].UnwrapExpr)
		}
	case 18:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newPipelineExpr(newMatcherExpr(exprDollar[1].Selector), exprDollar[2].PipelineExpr), exprDollar[3].duration, nil)
		}
	case 19:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newPipelineExpr(newMatcherExpr(exprDollar[2].Selector), exprDollar[3].PipelineExpr), exprDollar[5].duration, nil)
		}
	case 20:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newPipelineExpr(newMatcherExpr(exprDollar[1].Selector), exprDollar[2].PipelineExpr), exprDollar[4].duration, exprDollar[3].UnwrapExpr)
		}
	case 21:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newPipelineExpr(newMatcherExpr(exprDollar[2].Selector), exprDollar[3].PipelineExpr), exprDollar[6].duration, exprDollar[4].UnwrapExpr)
		}
	case 22:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newPipelineExpr(newMatcherExpr(exprDollar[1].Selector), exprDollar[3].PipelineExpr), exprDollar[2].duration, nil)
		}
	case 23:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.LogRangeExpr = newLogRange(newPipelineExpr(newMatcherExpr(exprDollar[1].Selector), exprDollar[3].PipelineExpr), exprDollar[2].duration, exprDollar[4].UnwrapExpr)
		}
	case 24:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LogRangeExpr = exprDollar[2].LogRangeExpr
		}
	case 26:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.UnwrapExpr = newUnwrapExpr(exprDollar[3].str, "")
		}
	case 27:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.UnwrapExpr = newUnwrapExpr(exprDollar[5].str, exprDollar[3].ConvOp)
		}
	case 28:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.UnwrapExpr = exprDollar[1].UnwrapExpr.addPostFilter(exprDollar[3].LabelFilter)
		}
	case 29:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ConvOp = OpConvBytes
		}
	case 30:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ConvOp = OpConvDuration
		}
	case 31:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.ConvOp = OpConvDurationSeconds
		}
	case 32:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.RangeAggregationExpr = newRangeAggregationExpr(exprDollar[3].LogRangeExpr, exprDollar[1].RangeOp, nil, nil)
		}
	case 33:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.RangeAggregationExpr = newRangeAggregationExpr(exprDollar[5].LogRangeExpr, exprDollar[1].RangeOp, nil, &exprDollar[3].str)
		}
	case 34:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.RangeAggregationExpr = newRangeAggregationExpr(exprDollar[3].LogRangeExpr, exprDollar[1].RangeOp, exprDollar[5].Grouping, nil)
		}
	case 35:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.RangeAggregationExpr = newRangeAggregationExpr(exprDollar[5].LogRangeExpr, exprDollar[1].RangeOp, exprDollar[7].Grouping, &exprDollar[3].str)
		}
	case 36:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[3].MetricExpr, exprDollar[1].VectorOp, nil, nil)
		}
	case 37:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[4].MetricExpr, exprDollar[1].VectorOp, exprDollar[2].Grouping, nil)
		}
	case 38:
		exprDollar = exprS[exprpt-5 : exprpt+1]
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[3].MetricExpr, exprDollar[1].VectorOp, exprDollar[5].Grouping, nil)
		}
	case 39:
		exprDollar = exprS[exprpt-6 : exprpt+1]
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[5].MetricExpr, exprDollar[1].VectorOp, nil, &exprDollar[3].str)
		}
	case 40:
		exprDollar = exprS[exprpt-7 : exprpt+1]
		{
			exprVAL.VectorAggregationExpr = mustNewVectorAggregationExpr(exprDollar[5].MetricExpr, exprDollar[1].VectorOp, exprDollar[7].Grouping, &exprDollar[3].str)
		}
	case 41:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Filter = labels.MatchRegexp
		}
	case 42:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Filter = labels.MatchEqual
		}
	case 43:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Filter = labels.MatchNotRegexp
		}
	case 44:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Filter = labels.MatchNotEqual
		}
	case 45:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Selector = exprDollar[2].Matchers
		}
	case 46:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Selector = exprDollar[2].Matchers
		}
	case 47:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
		}
	case 48:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Matchers = []*labels.Matcher{exprDollar[1].Matcher}
		}
	case 49:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matchers = append(exprDollar[1].Matchers, exprDollar[3].Matcher)
		}
	case 50:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchEqual, exprDollar[1].str, exprDollar[3].str)
		}
	case 51:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchNotEqual, exprDollar[1].str, exprDollar[3].str)
		}
	case 52:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchRegexp, exprDollar[1].str, exprDollar[3].str)
		}
	case 53:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Matcher = mustNewMatcher(labels.MatchNotRegexp, exprDollar[1].str, exprDollar[3].str)
		}
	case 54:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.PipelineExpr = MultiStageExpr{exprDollar[1].PipelineStage}
		}
	case 55:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.PipelineExpr = append(exprDollar[1].PipelineExpr, exprDollar[2].PipelineStage)
		}
	case 56:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.PipelineStage = exprDollar[1].LineFilters
		}
	case 57:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.PipelineStage = exprDollar[2].LabelParser
		}
	case 58:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.PipelineStage = &labelFilterExpr{LabelFilterer: exprDollar[2].LabelFilter}
		}
	case 59:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.PipelineStage = exprDollar[2].LineFormatExpr
		}
	case 60:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.PipelineStage = exprDollar[2].LabelFormatExpr
		}
	case 61:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LineFilters = newLineFilterExpr(nil, exprDollar[1].Filter, exprDollar[2].str)
		}
	case 62:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LineFilters = newLineFilterExpr(exprDollar[1].LineFilters, exprDollar[2].Filter, exprDollar[3].str)
		}
	case 63:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LabelParser = newLabelParserExpr(OpParserTypeJSON, "")
		}
	case 64:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LabelParser = newLabelParserExpr(OpParserTypeLogfmt, "")
		}
	case 65:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LabelParser = newLabelParserExpr(OpParserTypeRegexp, exprDollar[2].str)
		}
	case 66:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LineFormatExpr = newLineFmtExpr(exprDollar[2].str)
		}
	case 67:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LabelFormat = log.NewRenameLabelFmt(exprDollar[1].str, exprDollar[3].str)
		}
	case 68:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LabelFormat = log.NewTemplateLabelFmt(exprDollar[1].str, exprDollar[3].str)
		}
	case 69:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LabelsFormat = []log.LabelFmt{exprDollar[1].LabelFormat}
		}
	case 70:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LabelsFormat = append(exprDollar[1].LabelsFormat, exprDollar[3].LabelFormat)
		}
	case 72:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LabelFormatExpr = newLabelFmtExpr(exprDollar[2].LabelsFormat)
		}
	case 73:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LabelFilter = log.NewStringLabelFilter(exprDollar[1].Matcher)
		}
	case 74:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LabelFilter = exprDollar[1].UnitFilter
		}
	case 75:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LabelFilter = exprDollar[1].NumberFilter
		}
	case 76:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LabelFilter = exprDollar[2].LabelFilter
		}
	case 77:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LabelFilter = log.NewAndLabelFilter(exprDollar[1].LabelFilter, exprDollar[2].LabelFilter)
		}
	case 78:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LabelFilter = log.NewAndLabelFilter(exprDollar[1].LabelFilter, exprDollar[3].LabelFilter)
		}
	case 79:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LabelFilter = log.NewAndLabelFilter(exprDollar[1].LabelFilter, exprDollar[3].LabelFilter)
		}
	case 80:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.LabelFilter = log.NewOrLabelFilter(exprDollar[1].LabelFilter, exprDollar[3].LabelFilter)
		}
	case 81:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.UnitFilter = exprDollar[1].DurationFilter
		}
	case 82:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.UnitFilter = exprDollar[1].BytesFilter
		}
	case 83:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.DurationFilter = log.NewDurationLabelFilter(log.LabelFilterGreaterThan, exprDollar[1].str, exprDollar[3].duration)
		}
	case 84:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.DurationFilter = log.NewDurationLabelFilter(log.LabelFilterGreaterThanOrEqual, exprDollar[1].str, exprDollar[3].duration)
		}
	case 85:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.DurationFilter = log.NewDurationLabelFilter(log.LabelFilterLesserThan, exprDollar[1].str, exprDollar[3].duration)
		}
	case 86:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.DurationFilter = log.NewDurationLabelFilter(log.LabelFilterLesserThanOrEqual, exprDollar[1].str, exprDollar[3].duration)
		}
	case 87:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.DurationFilter = log.NewDurationLabelFilter(log.LabelFilterNotEqual, exprDollar[1].str, exprDollar[3].duration)
		}
	case 88:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.DurationFilter = log.NewDurationLabelFilter(log.LabelFilterEqual, exprDollar[1].str, exprDollar[3].duration)
		}
	case 89:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.DurationFilter = log.NewDurationLabelFilter(log.LabelFilterEqual, exprDollar[1].str, exprDollar[3].duration)
		}
	case 90:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.BytesFilter = log.NewBytesLabelFilter(log.LabelFilterGreaterThan, exprDollar[1].str, exprDollar[3].bytes)
		}
	case 91:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.BytesFilter = log.NewBytesLabelFilter(log.LabelFilterGreaterThanOrEqual, exprDollar[1].str, exprDollar[3].bytes)
		}
	case 92:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.BytesFilter = log.NewBytesLabelFilter(log.LabelFilterLesserThan, exprDollar[1].str, exprDollar[3].bytes)
		}
	case 93:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.BytesFilter = log.NewBytesLabelFilter(log.LabelFilterLesserThanOrEqual, exprDollar[1].str, exprDollar[3].bytes)
		}
	case 94:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.BytesFilter = log.NewBytesLabelFilter(log.LabelFilterNotEqual, exprDollar[1].str, exprDollar[3].bytes)
		}
	case 95:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.BytesFilter = log.NewBytesLabelFilter(log.LabelFilterEqual, exprDollar[1].str, exprDollar[3].bytes)
		}
	case 96:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.BytesFilter = log.NewBytesLabelFilter(log.LabelFilterEqual, exprDollar[1].str, exprDollar[3].bytes)
		}
	case 97:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.NumberFilter = log.NewNumericLabelFilter(log.LabelFilterGreaterThan, exprDollar[1].str, mustNewFloat(exprDollar[3].str))
		}
	case 98:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.NumberFilter = log.NewNumericLabelFilter(log.LabelFilterGreaterThanOrEqual, exprDollar[1].str, mustNewFloat(exprDollar[3].str))
		}
	case 99:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.NumberFilter = log.NewNumericLabelFilter(log.LabelFilterLesserThan, exprDollar[1].str, mustNewFloat(exprDollar[3].str))
		}
	case 100:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.NumberFilter = log.NewNumericLabelFilter(log.LabelFilterLesserThanOrEqual, exprDollar[1].str, mustNewFloat(exprDollar[3].str))
		}
	case 101:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.NumberFilter = log.NewNumericLabelFilter(log.LabelFilterNotEqual, exprDollar[1].str, mustNewFloat(exprDollar[3].str))
		}
	case 102:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.NumberFilter = log.NewNumericLabelFilter(log.LabelFilterEqual, exprDollar[1].str, mustNewFloat(exprDollar[3].str))
		}
	case 103:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.NumberFilter = log.NewNumericLabelFilter(log.LabelFilterEqual, exprDollar[1].str, mustNewFloat(exprDollar[3].str))
		}
	case 104:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("or", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 105:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("and", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 106:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("unless", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 107:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("+", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 108:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("-", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 109:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("*", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 110:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("/", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 111:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("%", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 112:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("^", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 113:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("==", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 114:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("!=", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 115:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr(">", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 116:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr(">=", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 117:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("<", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 118:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.BinOpExpr = mustNewBinOpExpr("<=", exprDollar[3].BinOpModifier, exprDollar[1].Expr, exprDollar[4].Expr)
		}
	case 119:
		exprDollar = exprS[exprpt-0 : exprpt+1]
		{
			exprVAL.BinOpModifier = BinOpOptions{}
		}
	case 120:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.BinOpModifier = BinOpOptions{ReturnBool: true}
		}
	case 121:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.LiteralExpr = mustNewLiteralExpr(exprDollar[1].str, false)
		}
	case 122:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LiteralExpr = mustNewLiteralExpr(exprDollar[2].str, false)
		}
	case 123:
		exprDollar = exprS[exprpt-2 : exprpt+1]
		{
			exprVAL.LiteralExpr = mustNewLiteralExpr(exprDollar[2].str, true)
		}
	case 124:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeSum
		}
	case 125:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeAvg
		}
	case 126:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeCount
		}
	case 127:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeMax
		}
	case 128:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeMin
		}
	case 129:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeStddev
		}
	case 130:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeStdvar
		}
	case 131:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeBottomK
		}
	case 132:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.VectorOp = OpTypeTopK
		}
	case 133:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeCount
		}
	case 134:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeRate
		}
	case 135:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeBytes
		}
	case 136:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeBytesRate
		}
	case 137:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeAvg
		}
	case 138:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeSum
		}
	case 139:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeMin
		}
	case 140:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeMax
		}
	case 141:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeStdvar
		}
	case 142:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeStddev
		}
	case 143:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.RangeOp = OpRangeTypeQuantile
		}
	case 144:
		exprDollar = exprS[exprpt-1 : exprpt+1]
		{
			exprVAL.Labels = []string{exprDollar[1].str}
		}
	case 145:
		exprDollar = exprS[exprpt-3 : exprpt+1]
		{
			exprVAL.Labels = append(exprDollar[1].Labels, exprDollar[3].str)
		}
	case 146:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.Grouping = &grouping{without: false, groups: exprDollar[3].Labels}
		}
	case 147:
		exprDollar = exprS[exprpt-4 : exprpt+1]
		{
			exprVAL.Grouping = &grouping{without: true, groups: exprDollar[3].Labels}
		}
	}
	goto exprstack /* stack new state and value */
}
