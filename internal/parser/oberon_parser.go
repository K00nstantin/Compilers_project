// Code generated from /home/konstantin/go/Compilers_project/grammar/oberon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // oberon
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type oberonParser struct {
	*antlr.BaseParser
}

var OberonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func oberonParserInit() {
	staticData := &OberonParserStaticData
	staticData.LiteralNames = []string{
		"", "'.'", "'*'", "'H'", "'E'", "'+'", "'-'", "'='", "','", "'('", "')'",
		"';'", "':'", "'#'", "'<'", "'<='", "'>'", "'>='", "'/'", "'&'", "'~'",
		"'['", "']'", "'^'", "'{'", "'}'", "'..'", "':='", "'|'", "'ARRAY'",
		"'OF'", "'END'", "'POINTER'", "'TO'", "'RECORD'", "'PROCEDURE'", "'IN'",
		"'IS'", "'OR'", "'DIV'", "'MOD'", "'NIL'", "'TRUE'", "'FALSE'", "'IF'",
		"'THEN'", "'ELSIF'", "'ELSE'", "'CASE'", "'WHILE'", "'DO'", "'REPEAT'",
		"'UNTIL'", "'FOR'", "'BY'", "'BEGIN'", "'RETURN'", "'CONST'", "'TYPE'",
		"'VAR'", "'MODULE'", "'IMPORT'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "ARRAY", "OF", "END",
		"POINTER", "TO", "RECORD", "PROCEDURE", "IN", "IS", "OR", "DIV", "MOD",
		"NIL", "TRUE", "FALSE", "IF", "THEN", "ELSIF", "ELSE", "CASE", "WHILE",
		"DO", "REPEAT", "UNTIL", "FOR", "BY", "BEGIN", "RETURN", "CONST", "TYPE",
		"VAR", "MODULE", "IMPORT", "STRING", "HEXDIGIT", "IDENT", "LETTER",
		"DIGIT", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"ident", "qualident", "identdef", "integer", "real", "scaleFactor",
		"number", "constDeclaration", "constExpression", "typeDeclaration",
		"type_", "arrayType", "length", "recordType", "baseType", "fieldListSequence",
		"fieldList", "identList", "pointerType", "procedureType", "variableDeclaration",
		"expression", "relation", "simpleExpression", "addOperator", "term",
		"mulOperator", "factor", "designator", "selector", "set_", "element",
		"expList", "actualParameters", "statement", "assignment", "procedureCall",
		"statementSequence", "ifStatement", "caseStatement", "case_", "caseLabelList",
		"labelRange", "label", "whileStatement", "repeatStatement", "forStatement",
		"procedureDeclaration", "procedureHeading", "procedureBody", "declarationSequence",
		"formalParameters", "fPSection", "formalType", "module", "importList",
		"import_",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 68, 601, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 3, 1, 120, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 126, 8,
		2, 1, 3, 4, 3, 129, 8, 3, 11, 3, 12, 3, 130, 1, 3, 1, 3, 5, 3, 135, 8,
		3, 10, 3, 12, 3, 138, 9, 3, 1, 3, 3, 3, 141, 8, 3, 1, 4, 4, 4, 144, 8,
		4, 11, 4, 12, 4, 145, 1, 4, 1, 4, 5, 4, 150, 8, 4, 10, 4, 12, 4, 153, 9,
		4, 1, 4, 3, 4, 156, 8, 4, 1, 5, 1, 5, 3, 5, 160, 8, 5, 1, 5, 4, 5, 163,
		8, 5, 11, 5, 12, 5, 164, 1, 6, 1, 6, 3, 6, 169, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 186, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 192, 8, 11, 10,
		11, 12, 11, 195, 9, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 3, 13, 207, 8, 13, 1, 13, 3, 13, 210, 8, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 219, 8, 15, 10, 15, 12,
		15, 222, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17,
		231, 8, 17, 10, 17, 12, 17, 234, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 3, 19, 242, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 3, 21, 252, 8, 21, 1, 22, 1, 22, 1, 23, 3, 23, 257, 8, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 263, 8, 23, 10, 23, 12, 23, 266, 9,
		23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 274, 8, 25, 10, 25,
		12, 25, 277, 9, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 3, 27, 289, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 3, 27, 297, 8, 27, 1, 28, 1, 28, 5, 28, 301, 8, 28, 10, 28, 12,
		28, 304, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 3, 29, 317, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 5,
		30, 323, 8, 30, 10, 30, 12, 30, 326, 9, 30, 3, 30, 328, 8, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 31, 3, 31, 335, 8, 31, 1, 32, 1, 32, 1, 32, 5, 32,
		340, 8, 32, 10, 32, 12, 32, 343, 9, 32, 1, 33, 1, 33, 3, 33, 347, 8, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 358,
		8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 3, 36, 366, 8, 36, 1,
		37, 1, 37, 1, 37, 5, 37, 371, 8, 37, 10, 37, 12, 37, 374, 9, 37, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 385, 8,
		38, 10, 38, 12, 38, 388, 9, 38, 1, 38, 1, 38, 3, 38, 392, 8, 38, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 402, 8, 39, 10,
		39, 12, 39, 405, 9, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40,
		413, 8, 40, 1, 41, 1, 41, 1, 41, 5, 41, 418, 8, 41, 10, 41, 12, 41, 421,
		9, 41, 1, 42, 1, 42, 1, 42, 3, 42, 426, 8, 42, 1, 43, 1, 43, 1, 43, 3,
		43, 431, 8, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 5, 44, 442, 8, 44, 10, 44, 12, 44, 445, 9, 44, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 3, 46, 462, 8, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 3, 48, 476, 8, 48, 1, 49,
		1, 49, 1, 49, 3, 49, 481, 8, 49, 1, 49, 1, 49, 3, 49, 485, 8, 49, 1, 49,
		1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 493, 8, 50, 10, 50, 12, 50, 496,
		9, 50, 3, 50, 498, 8, 50, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 504, 8, 50,
		10, 50, 12, 50, 507, 9, 50, 3, 50, 509, 8, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 5, 50, 515, 8, 50, 10, 50, 12, 50, 518, 9, 50, 3, 50, 520, 8, 50, 1,
		50, 1, 50, 1, 50, 5, 50, 525, 8, 50, 10, 50, 12, 50, 528, 9, 50, 1, 51,
		1, 51, 1, 51, 1, 51, 5, 51, 534, 8, 51, 10, 51, 12, 51, 537, 9, 51, 3,
		51, 539, 8, 51, 1, 51, 1, 51, 1, 51, 3, 51, 544, 8, 51, 1, 52, 3, 52, 547,
		8, 52, 1, 52, 1, 52, 1, 52, 5, 52, 552, 8, 52, 10, 52, 12, 52, 555, 9,
		52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 5, 53, 562, 8, 53, 10, 53, 12, 53,
		565, 9, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 573, 8, 54,
		1, 54, 1, 54, 1, 54, 3, 54, 578, 8, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 55, 1, 55, 1, 55, 1, 55, 5, 55, 589, 8, 55, 10, 55, 12, 55, 592,
		9, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 3, 56, 599, 8, 56, 1, 56, 0,
		0, 57, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		106, 108, 110, 112, 0, 4, 1, 0, 5, 6, 3, 0, 7, 7, 13, 17, 36, 37, 2, 0,
		5, 6, 38, 38, 3, 0, 2, 2, 18, 19, 39, 40, 625, 0, 114, 1, 0, 0, 0, 2, 119,
		1, 0, 0, 0, 4, 123, 1, 0, 0, 0, 6, 140, 1, 0, 0, 0, 8, 143, 1, 0, 0, 0,
		10, 157, 1, 0, 0, 0, 12, 168, 1, 0, 0, 0, 14, 170, 1, 0, 0, 0, 16, 174,
		1, 0, 0, 0, 18, 176, 1, 0, 0, 0, 20, 185, 1, 0, 0, 0, 22, 187, 1, 0, 0,
		0, 24, 199, 1, 0, 0, 0, 26, 201, 1, 0, 0, 0, 28, 213, 1, 0, 0, 0, 30, 215,
		1, 0, 0, 0, 32, 223, 1, 0, 0, 0, 34, 227, 1, 0, 0, 0, 36, 235, 1, 0, 0,
		0, 38, 239, 1, 0, 0, 0, 40, 243, 1, 0, 0, 0, 42, 247, 1, 0, 0, 0, 44, 253,
		1, 0, 0, 0, 46, 256, 1, 0, 0, 0, 48, 267, 1, 0, 0, 0, 50, 269, 1, 0, 0,
		0, 52, 278, 1, 0, 0, 0, 54, 296, 1, 0, 0, 0, 56, 298, 1, 0, 0, 0, 58, 316,
		1, 0, 0, 0, 60, 318, 1, 0, 0, 0, 62, 331, 1, 0, 0, 0, 64, 336, 1, 0, 0,
		0, 66, 344, 1, 0, 0, 0, 68, 357, 1, 0, 0, 0, 70, 359, 1, 0, 0, 0, 72, 363,
		1, 0, 0, 0, 74, 367, 1, 0, 0, 0, 76, 375, 1, 0, 0, 0, 78, 395, 1, 0, 0,
		0, 80, 412, 1, 0, 0, 0, 82, 414, 1, 0, 0, 0, 84, 422, 1, 0, 0, 0, 86, 430,
		1, 0, 0, 0, 88, 432, 1, 0, 0, 0, 90, 448, 1, 0, 0, 0, 92, 453, 1, 0, 0,
		0, 94, 467, 1, 0, 0, 0, 96, 472, 1, 0, 0, 0, 98, 477, 1, 0, 0, 0, 100,
		497, 1, 0, 0, 0, 102, 529, 1, 0, 0, 0, 104, 546, 1, 0, 0, 0, 106, 563,
		1, 0, 0, 0, 108, 568, 1, 0, 0, 0, 110, 584, 1, 0, 0, 0, 112, 595, 1, 0,
		0, 0, 114, 115, 5, 64, 0, 0, 115, 1, 1, 0, 0, 0, 116, 117, 3, 0, 0, 0,
		117, 118, 5, 1, 0, 0, 118, 120, 1, 0, 0, 0, 119, 116, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 3, 0, 0, 0, 122, 3, 1,
		0, 0, 0, 123, 125, 3, 0, 0, 0, 124, 126, 5, 2, 0, 0, 125, 124, 1, 0, 0,
		0, 125, 126, 1, 0, 0, 0, 126, 5, 1, 0, 0, 0, 127, 129, 5, 66, 0, 0, 128,
		127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131,
		1, 0, 0, 0, 131, 141, 1, 0, 0, 0, 132, 136, 5, 66, 0, 0, 133, 135, 5, 63,
		0, 0, 134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139,
		141, 5, 3, 0, 0, 140, 128, 1, 0, 0, 0, 140, 132, 1, 0, 0, 0, 141, 7, 1,
		0, 0, 0, 142, 144, 5, 66, 0, 0, 143, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0,
		0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147,
		151, 5, 1, 0, 0, 148, 150, 5, 66, 0, 0, 149, 148, 1, 0, 0, 0, 150, 153,
		1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 155, 1, 0,
		0, 0, 153, 151, 1, 0, 0, 0, 154, 156, 3, 10, 5, 0, 155, 154, 1, 0, 0, 0,
		155, 156, 1, 0, 0, 0, 156, 9, 1, 0, 0, 0, 157, 159, 5, 4, 0, 0, 158, 160,
		7, 0, 0, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 162, 1, 0,
		0, 0, 161, 163, 5, 66, 0, 0, 162, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0,
		164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 11, 1, 0, 0, 0, 166, 169,
		3, 6, 3, 0, 167, 169, 3, 8, 4, 0, 168, 166, 1, 0, 0, 0, 168, 167, 1, 0,
		0, 0, 169, 13, 1, 0, 0, 0, 170, 171, 3, 4, 2, 0, 171, 172, 5, 7, 0, 0,
		172, 173, 3, 16, 8, 0, 173, 15, 1, 0, 0, 0, 174, 175, 3, 42, 21, 0, 175,
		17, 1, 0, 0, 0, 176, 177, 3, 4, 2, 0, 177, 178, 5, 7, 0, 0, 178, 179, 3,
		20, 10, 0, 179, 19, 1, 0, 0, 0, 180, 186, 3, 2, 1, 0, 181, 186, 3, 22,
		11, 0, 182, 186, 3, 26, 13, 0, 183, 186, 3, 36, 18, 0, 184, 186, 3, 38,
		19, 0, 185, 180, 1, 0, 0, 0, 185, 181, 1, 0, 0, 0, 185, 182, 1, 0, 0, 0,
		185, 183, 1, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186, 21, 1, 0, 0, 0, 187, 188,
		5, 29, 0, 0, 188, 193, 3, 24, 12, 0, 189, 190, 5, 8, 0, 0, 190, 192, 3,
		24, 12, 0, 191, 189, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0,
		0, 0, 193, 194, 1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0,
		196, 197, 5, 30, 0, 0, 197, 198, 3, 20, 10, 0, 198, 23, 1, 0, 0, 0, 199,
		200, 3, 16, 8, 0, 200, 25, 1, 0, 0, 0, 201, 206, 5, 34, 0, 0, 202, 203,
		5, 9, 0, 0, 203, 204, 3, 28, 14, 0, 204, 205, 5, 10, 0, 0, 205, 207, 1,
		0, 0, 0, 206, 202, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 209, 1, 0, 0,
		0, 208, 210, 3, 30, 15, 0, 209, 208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0,
		210, 211, 1, 0, 0, 0, 211, 212, 5, 31, 0, 0, 212, 27, 1, 0, 0, 0, 213,
		214, 3, 2, 1, 0, 214, 29, 1, 0, 0, 0, 215, 220, 3, 32, 16, 0, 216, 217,
		5, 11, 0, 0, 217, 219, 3, 32, 16, 0, 218, 216, 1, 0, 0, 0, 219, 222, 1,
		0, 0, 0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 31, 1, 0, 0,
		0, 222, 220, 1, 0, 0, 0, 223, 224, 3, 34, 17, 0, 224, 225, 5, 12, 0, 0,
		225, 226, 3, 20, 10, 0, 226, 33, 1, 0, 0, 0, 227, 232, 3, 4, 2, 0, 228,
		229, 5, 8, 0, 0, 229, 231, 3, 4, 2, 0, 230, 228, 1, 0, 0, 0, 231, 234,
		1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 35, 1, 0,
		0, 0, 234, 232, 1, 0, 0, 0, 235, 236, 5, 32, 0, 0, 236, 237, 5, 33, 0,
		0, 237, 238, 3, 20, 10, 0, 238, 37, 1, 0, 0, 0, 239, 241, 5, 35, 0, 0,
		240, 242, 3, 102, 51, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242,
		39, 1, 0, 0, 0, 243, 244, 3, 34, 17, 0, 244, 245, 5, 12, 0, 0, 245, 246,
		3, 20, 10, 0, 246, 41, 1, 0, 0, 0, 247, 251, 3, 46, 23, 0, 248, 249, 3,
		44, 22, 0, 249, 250, 3, 46, 23, 0, 250, 252, 1, 0, 0, 0, 251, 248, 1, 0,
		0, 0, 251, 252, 1, 0, 0, 0, 252, 43, 1, 0, 0, 0, 253, 254, 7, 1, 0, 0,
		254, 45, 1, 0, 0, 0, 255, 257, 7, 0, 0, 0, 256, 255, 1, 0, 0, 0, 256, 257,
		1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 264, 3, 50, 25, 0, 259, 260, 3,
		48, 24, 0, 260, 261, 3, 50, 25, 0, 261, 263, 1, 0, 0, 0, 262, 259, 1, 0,
		0, 0, 263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0,
		265, 47, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 268, 7, 2, 0, 0, 268, 49,
		1, 0, 0, 0, 269, 275, 3, 54, 27, 0, 270, 271, 3, 52, 26, 0, 271, 272, 3,
		54, 27, 0, 272, 274, 1, 0, 0, 0, 273, 270, 1, 0, 0, 0, 274, 277, 1, 0,
		0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 51, 1, 0, 0, 0,
		277, 275, 1, 0, 0, 0, 278, 279, 7, 3, 0, 0, 279, 53, 1, 0, 0, 0, 280, 297,
		3, 12, 6, 0, 281, 297, 5, 62, 0, 0, 282, 297, 5, 41, 0, 0, 283, 297, 5,
		42, 0, 0, 284, 297, 5, 43, 0, 0, 285, 297, 3, 60, 30, 0, 286, 288, 3, 56,
		28, 0, 287, 289, 3, 66, 33, 0, 288, 287, 1, 0, 0, 0, 288, 289, 1, 0, 0,
		0, 289, 297, 1, 0, 0, 0, 290, 291, 5, 9, 0, 0, 291, 292, 3, 42, 21, 0,
		292, 293, 5, 10, 0, 0, 293, 297, 1, 0, 0, 0, 294, 295, 5, 20, 0, 0, 295,
		297, 3, 54, 27, 0, 296, 280, 1, 0, 0, 0, 296, 281, 1, 0, 0, 0, 296, 282,
		1, 0, 0, 0, 296, 283, 1, 0, 0, 0, 296, 284, 1, 0, 0, 0, 296, 285, 1, 0,
		0, 0, 296, 286, 1, 0, 0, 0, 296, 290, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0,
		297, 55, 1, 0, 0, 0, 298, 302, 3, 2, 1, 0, 299, 301, 3, 58, 29, 0, 300,
		299, 1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303,
		1, 0, 0, 0, 303, 57, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 306, 5, 1,
		0, 0, 306, 317, 3, 0, 0, 0, 307, 308, 5, 21, 0, 0, 308, 309, 3, 64, 32,
		0, 309, 310, 5, 22, 0, 0, 310, 317, 1, 0, 0, 0, 311, 317, 5, 23, 0, 0,
		312, 313, 5, 9, 0, 0, 313, 314, 3, 2, 1, 0, 314, 315, 5, 10, 0, 0, 315,
		317, 1, 0, 0, 0, 316, 305, 1, 0, 0, 0, 316, 307, 1, 0, 0, 0, 316, 311,
		1, 0, 0, 0, 316, 312, 1, 0, 0, 0, 317, 59, 1, 0, 0, 0, 318, 327, 5, 24,
		0, 0, 319, 324, 3, 62, 31, 0, 320, 321, 5, 8, 0, 0, 321, 323, 3, 62, 31,
		0, 322, 320, 1, 0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324,
		325, 1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 319,
		1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 330, 5, 25,
		0, 0, 330, 61, 1, 0, 0, 0, 331, 334, 3, 42, 21, 0, 332, 333, 5, 26, 0,
		0, 333, 335, 3, 42, 21, 0, 334, 332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0,
		335, 63, 1, 0, 0, 0, 336, 341, 3, 42, 21, 0, 337, 338, 5, 8, 0, 0, 338,
		340, 3, 42, 21, 0, 339, 337, 1, 0, 0, 0, 340, 343, 1, 0, 0, 0, 341, 339,
		1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 65, 1, 0, 0, 0, 343, 341, 1, 0,
		0, 0, 344, 346, 5, 9, 0, 0, 345, 347, 3, 64, 32, 0, 346, 345, 1, 0, 0,
		0, 346, 347, 1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 349, 5, 10, 0, 0, 349,
		67, 1, 0, 0, 0, 350, 358, 3, 70, 35, 0, 351, 358, 3, 72, 36, 0, 352, 358,
		3, 76, 38, 0, 353, 358, 3, 78, 39, 0, 354, 358, 3, 88, 44, 0, 355, 358,
		3, 90, 45, 0, 356, 358, 3, 92, 46, 0, 357, 350, 1, 0, 0, 0, 357, 351, 1,
		0, 0, 0, 357, 352, 1, 0, 0, 0, 357, 353, 1, 0, 0, 0, 357, 354, 1, 0, 0,
		0, 357, 355, 1, 0, 0, 0, 357, 356, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358,
		69, 1, 0, 0, 0, 359, 360, 3, 56, 28, 0, 360, 361, 5, 27, 0, 0, 361, 362,
		3, 42, 21, 0, 362, 71, 1, 0, 0, 0, 363, 365, 3, 56, 28, 0, 364, 366, 3,
		66, 33, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 73, 1, 0, 0,
		0, 367, 372, 3, 68, 34, 0, 368, 369, 5, 11, 0, 0, 369, 371, 3, 68, 34,
		0, 370, 368, 1, 0, 0, 0, 371, 374, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 372,
		373, 1, 0, 0, 0, 373, 75, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 375, 376, 5,
		44, 0, 0, 376, 377, 3, 42, 21, 0, 377, 378, 5, 45, 0, 0, 378, 386, 3, 74,
		37, 0, 379, 380, 5, 46, 0, 0, 380, 381, 3, 42, 21, 0, 381, 382, 5, 45,
		0, 0, 382, 383, 3, 74, 37, 0, 383, 385, 1, 0, 0, 0, 384, 379, 1, 0, 0,
		0, 385, 388, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387,
		391, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 389, 390, 5, 47, 0, 0, 390, 392,
		3, 74, 37, 0, 391, 389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 393, 1,
		0, 0, 0, 393, 394, 5, 31, 0, 0, 394, 77, 1, 0, 0, 0, 395, 396, 5, 48, 0,
		0, 396, 397, 3, 42, 21, 0, 397, 398, 5, 30, 0, 0, 398, 403, 3, 80, 40,
		0, 399, 400, 5, 28, 0, 0, 400, 402, 3, 80, 40, 0, 401, 399, 1, 0, 0, 0,
		402, 405, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404,
		406, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 406, 407, 5, 31, 0, 0, 407, 79,
		1, 0, 0, 0, 408, 409, 3, 82, 41, 0, 409, 410, 5, 12, 0, 0, 410, 411, 3,
		74, 37, 0, 411, 413, 1, 0, 0, 0, 412, 408, 1, 0, 0, 0, 412, 413, 1, 0,
		0, 0, 413, 81, 1, 0, 0, 0, 414, 419, 3, 84, 42, 0, 415, 416, 5, 8, 0, 0,
		416, 418, 3, 84, 42, 0, 417, 415, 1, 0, 0, 0, 418, 421, 1, 0, 0, 0, 419,
		417, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 83, 1, 0, 0, 0, 421, 419, 1,
		0, 0, 0, 422, 425, 3, 86, 43, 0, 423, 424, 5, 26, 0, 0, 424, 426, 3, 86,
		43, 0, 425, 423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 85, 1, 0, 0, 0,
		427, 431, 3, 6, 3, 0, 428, 431, 5, 62, 0, 0, 429, 431, 3, 2, 1, 0, 430,
		427, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 429, 1, 0, 0, 0, 431, 87, 1,
		0, 0, 0, 432, 433, 5, 49, 0, 0, 433, 434, 3, 42, 21, 0, 434, 435, 5, 50,
		0, 0, 435, 443, 3, 74, 37, 0, 436, 437, 5, 46, 0, 0, 437, 438, 3, 42, 21,
		0, 438, 439, 5, 50, 0, 0, 439, 440, 3, 74, 37, 0, 440, 442, 1, 0, 0, 0,
		441, 436, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443,
		444, 1, 0, 0, 0, 444, 446, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 447,
		5, 31, 0, 0, 447, 89, 1, 0, 0, 0, 448, 449, 5, 51, 0, 0, 449, 450, 3, 74,
		37, 0, 450, 451, 5, 52, 0, 0, 451, 452, 3, 42, 21, 0, 452, 91, 1, 0, 0,
		0, 453, 454, 5, 53, 0, 0, 454, 455, 3, 0, 0, 0, 455, 456, 5, 27, 0, 0,
		456, 457, 3, 42, 21, 0, 457, 458, 5, 33, 0, 0, 458, 461, 3, 42, 21, 0,
		459, 460, 5, 54, 0, 0, 460, 462, 3, 16, 8, 0, 461, 459, 1, 0, 0, 0, 461,
		462, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 464, 5, 50, 0, 0, 464, 465,
		3, 74, 37, 0, 465, 466, 5, 31, 0, 0, 466, 93, 1, 0, 0, 0, 467, 468, 3,
		96, 48, 0, 468, 469, 5, 11, 0, 0, 469, 470, 3, 98, 49, 0, 470, 471, 3,
		0, 0, 0, 471, 95, 1, 0, 0, 0, 472, 473, 5, 35, 0, 0, 473, 475, 3, 4, 2,
		0, 474, 476, 3, 102, 51, 0, 475, 474, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0,
		476, 97, 1, 0, 0, 0, 477, 480, 3, 100, 50, 0, 478, 479, 5, 55, 0, 0, 479,
		481, 3, 74, 37, 0, 480, 478, 1, 0, 0, 0, 480, 481, 1, 0, 0, 0, 481, 484,
		1, 0, 0, 0, 482, 483, 5, 56, 0, 0, 483, 485, 3, 42, 21, 0, 484, 482, 1,
		0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 487, 5, 31, 0,
		0, 487, 99, 1, 0, 0, 0, 488, 494, 5, 57, 0, 0, 489, 490, 3, 14, 7, 0, 490,
		491, 5, 11, 0, 0, 491, 493, 1, 0, 0, 0, 492, 489, 1, 0, 0, 0, 493, 496,
		1, 0, 0, 0, 494, 492, 1, 0, 0, 0, 494, 495, 1, 0, 0, 0, 495, 498, 1, 0,
		0, 0, 496, 494, 1, 0, 0, 0, 497, 488, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0,
		498, 508, 1, 0, 0, 0, 499, 505, 5, 58, 0, 0, 500, 501, 3, 18, 9, 0, 501,
		502, 5, 11, 0, 0, 502, 504, 1, 0, 0, 0, 503, 500, 1, 0, 0, 0, 504, 507,
		1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 505, 506, 1, 0, 0, 0, 506, 509, 1, 0,
		0, 0, 507, 505, 1, 0, 0, 0, 508, 499, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0,
		509, 519, 1, 0, 0, 0, 510, 516, 5, 59, 0, 0, 511, 512, 3, 40, 20, 0, 512,
		513, 5, 11, 0, 0, 513, 515, 1, 0, 0, 0, 514, 511, 1, 0, 0, 0, 515, 518,
		1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 520, 1, 0,
		0, 0, 518, 516, 1, 0, 0, 0, 519, 510, 1, 0, 0, 0, 519, 520, 1, 0, 0, 0,
		520, 526, 1, 0, 0, 0, 521, 522, 3, 94, 47, 0, 522, 523, 5, 11, 0, 0, 523,
		525, 1, 0, 0, 0, 524, 521, 1, 0, 0, 0, 525, 528, 1, 0, 0, 0, 526, 524,
		1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 101, 1, 0, 0, 0, 528, 526, 1, 0,
		0, 0, 529, 538, 5, 9, 0, 0, 530, 535, 3, 104, 52, 0, 531, 532, 5, 11, 0,
		0, 532, 534, 3, 104, 52, 0, 533, 531, 1, 0, 0, 0, 534, 537, 1, 0, 0, 0,
		535, 533, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 539, 1, 0, 0, 0, 537,
		535, 1, 0, 0, 0, 538, 530, 1, 0, 0, 0, 538, 539, 1, 0, 0, 0, 539, 540,
		1, 0, 0, 0, 540, 543, 5, 10, 0, 0, 541, 542, 5, 12, 0, 0, 542, 544, 3,
		2, 1, 0, 543, 541, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 103, 1, 0, 0,
		0, 545, 547, 5, 59, 0, 0, 546, 545, 1, 0, 0, 0, 546, 547, 1, 0, 0, 0, 547,
		548, 1, 0, 0, 0, 548, 553, 3, 0, 0, 0, 549, 550, 5, 8, 0, 0, 550, 552,
		3, 0, 0, 0, 551, 549, 1, 0, 0, 0, 552, 555, 1, 0, 0, 0, 553, 551, 1, 0,
		0, 0, 553, 554, 1, 0, 0, 0, 554, 556, 1, 0, 0, 0, 555, 553, 1, 0, 0, 0,
		556, 557, 5, 12, 0, 0, 557, 558, 3, 106, 53, 0, 558, 105, 1, 0, 0, 0, 559,
		560, 5, 29, 0, 0, 560, 562, 5, 30, 0, 0, 561, 559, 1, 0, 0, 0, 562, 565,
		1, 0, 0, 0, 563, 561, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564, 566, 1, 0,
		0, 0, 565, 563, 1, 0, 0, 0, 566, 567, 3, 2, 1, 0, 567, 107, 1, 0, 0, 0,
		568, 569, 5, 60, 0, 0, 569, 570, 3, 0, 0, 0, 570, 572, 5, 11, 0, 0, 571,
		573, 3, 110, 55, 0, 572, 571, 1, 0, 0, 0, 572, 573, 1, 0, 0, 0, 573, 574,
		1, 0, 0, 0, 574, 577, 3, 100, 50, 0, 575, 576, 5, 55, 0, 0, 576, 578, 3,
		74, 37, 0, 577, 575, 1, 0, 0, 0, 577, 578, 1, 0, 0, 0, 578, 579, 1, 0,
		0, 0, 579, 580, 5, 31, 0, 0, 580, 581, 3, 0, 0, 0, 581, 582, 5, 1, 0, 0,
		582, 583, 5, 0, 0, 1, 583, 109, 1, 0, 0, 0, 584, 585, 5, 61, 0, 0, 585,
		590, 3, 112, 56, 0, 586, 587, 5, 8, 0, 0, 587, 589, 3, 112, 56, 0, 588,
		586, 1, 0, 0, 0, 589, 592, 1, 0, 0, 0, 590, 588, 1, 0, 0, 0, 590, 591,
		1, 0, 0, 0, 591, 593, 1, 0, 0, 0, 592, 590, 1, 0, 0, 0, 593, 594, 5, 11,
		0, 0, 594, 111, 1, 0, 0, 0, 595, 598, 3, 0, 0, 0, 596, 597, 5, 27, 0, 0,
		597, 599, 3, 0, 0, 0, 598, 596, 1, 0, 0, 0, 598, 599, 1, 0, 0, 0, 599,
		113, 1, 0, 0, 0, 63, 119, 125, 130, 136, 140, 145, 151, 155, 159, 164,
		168, 185, 193, 206, 209, 220, 232, 241, 251, 256, 264, 275, 288, 296, 302,
		316, 324, 327, 334, 341, 346, 357, 365, 372, 386, 391, 403, 412, 419, 425,
		430, 443, 461, 475, 480, 484, 494, 497, 505, 508, 516, 519, 526, 535, 538,
		543, 546, 553, 563, 572, 577, 590, 598,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// oberonParserInit initializes any static state used to implement oberonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewoberonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OberonParserInit() {
	staticData := &OberonParserStaticData
	staticData.once.Do(oberonParserInit)
}

// NewoberonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewoberonParser(input antlr.TokenStream) *oberonParser {
	OberonParserInit()
	this := new(oberonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &OberonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "oberon.g4"

	return this
}

// oberonParser tokens.
const (
	oberonParserEOF       = antlr.TokenEOF
	oberonParserT__0      = 1
	oberonParserT__1      = 2
	oberonParserT__2      = 3
	oberonParserT__3      = 4
	oberonParserT__4      = 5
	oberonParserT__5      = 6
	oberonParserT__6      = 7
	oberonParserT__7      = 8
	oberonParserT__8      = 9
	oberonParserT__9      = 10
	oberonParserT__10     = 11
	oberonParserT__11     = 12
	oberonParserT__12     = 13
	oberonParserT__13     = 14
	oberonParserT__14     = 15
	oberonParserT__15     = 16
	oberonParserT__16     = 17
	oberonParserT__17     = 18
	oberonParserT__18     = 19
	oberonParserT__19     = 20
	oberonParserT__20     = 21
	oberonParserT__21     = 22
	oberonParserT__22     = 23
	oberonParserT__23     = 24
	oberonParserT__24     = 25
	oberonParserT__25     = 26
	oberonParserT__26     = 27
	oberonParserT__27     = 28
	oberonParserARRAY     = 29
	oberonParserOF        = 30
	oberonParserEND       = 31
	oberonParserPOINTER   = 32
	oberonParserTO        = 33
	oberonParserRECORD    = 34
	oberonParserPROCEDURE = 35
	oberonParserIN        = 36
	oberonParserIS        = 37
	oberonParserOR        = 38
	oberonParserDIV       = 39
	oberonParserMOD       = 40
	oberonParserNIL       = 41
	oberonParserTRUE      = 42
	oberonParserFALSE     = 43
	oberonParserIF        = 44
	oberonParserTHEN      = 45
	oberonParserELSIF     = 46
	oberonParserELSE      = 47
	oberonParserCASE      = 48
	oberonParserWHILE     = 49
	oberonParserDO        = 50
	oberonParserREPEAT    = 51
	oberonParserUNTIL     = 52
	oberonParserFOR       = 53
	oberonParserBY        = 54
	oberonParserBEGIN     = 55
	oberonParserRETURN    = 56
	oberonParserCONST     = 57
	oberonParserTYPE      = 58
	oberonParserVAR       = 59
	oberonParserMODULE    = 60
	oberonParserIMPORT    = 61
	oberonParserSTRING    = 62
	oberonParserHEXDIGIT  = 63
	oberonParserIDENT     = 64
	oberonParserLETTER    = 65
	oberonParserDIGIT     = 66
	oberonParserCOMMENT   = 67
	oberonParserWS        = 68
)

// oberonParser rules.
const (
	oberonParserRULE_ident                = 0
	oberonParserRULE_qualident            = 1
	oberonParserRULE_identdef             = 2
	oberonParserRULE_integer              = 3
	oberonParserRULE_real                 = 4
	oberonParserRULE_scaleFactor          = 5
	oberonParserRULE_number               = 6
	oberonParserRULE_constDeclaration     = 7
	oberonParserRULE_constExpression      = 8
	oberonParserRULE_typeDeclaration      = 9
	oberonParserRULE_type_                = 10
	oberonParserRULE_arrayType            = 11
	oberonParserRULE_length               = 12
	oberonParserRULE_recordType           = 13
	oberonParserRULE_baseType             = 14
	oberonParserRULE_fieldListSequence    = 15
	oberonParserRULE_fieldList            = 16
	oberonParserRULE_identList            = 17
	oberonParserRULE_pointerType          = 18
	oberonParserRULE_procedureType        = 19
	oberonParserRULE_variableDeclaration  = 20
	oberonParserRULE_expression           = 21
	oberonParserRULE_relation             = 22
	oberonParserRULE_simpleExpression     = 23
	oberonParserRULE_addOperator          = 24
	oberonParserRULE_term                 = 25
	oberonParserRULE_mulOperator          = 26
	oberonParserRULE_factor               = 27
	oberonParserRULE_designator           = 28
	oberonParserRULE_selector             = 29
	oberonParserRULE_set_                 = 30
	oberonParserRULE_element              = 31
	oberonParserRULE_expList              = 32
	oberonParserRULE_actualParameters     = 33
	oberonParserRULE_statement            = 34
	oberonParserRULE_assignment           = 35
	oberonParserRULE_procedureCall        = 36
	oberonParserRULE_statementSequence    = 37
	oberonParserRULE_ifStatement          = 38
	oberonParserRULE_caseStatement        = 39
	oberonParserRULE_case_                = 40
	oberonParserRULE_caseLabelList        = 41
	oberonParserRULE_labelRange           = 42
	oberonParserRULE_label                = 43
	oberonParserRULE_whileStatement       = 44
	oberonParserRULE_repeatStatement      = 45
	oberonParserRULE_forStatement         = 46
	oberonParserRULE_procedureDeclaration = 47
	oberonParserRULE_procedureHeading     = 48
	oberonParserRULE_procedureBody        = 49
	oberonParserRULE_declarationSequence  = 50
	oberonParserRULE_formalParameters     = 51
	oberonParserRULE_fPSection            = 52
	oberonParserRULE_formalType           = 53
	oberonParserRULE_module               = 54
	oberonParserRULE_importList           = 55
	oberonParserRULE_import_              = 56
)

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(oberonParserIDENT, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, oberonParserRULE_ident)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(oberonParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQualidentContext is an interface to support dynamic dispatch.
type IQualidentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsQualidentContext differentiates from other interfaces.
	IsQualidentContext()
}

type QualidentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualidentContext() *QualidentContext {
	var p = new(QualidentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_qualident
	return p
}

func InitEmptyQualidentContext(p *QualidentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_qualident
}

func (*QualidentContext) IsQualidentContext() {}

func NewQualidentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualidentContext {
	var p = new(QualidentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_qualident

	return p
}

func (s *QualidentContext) GetParser() antlr.Parser { return s.parser }

func (s *QualidentContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *QualidentContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *QualidentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualidentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualidentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterQualident(s)
	}
}

func (s *QualidentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitQualident(s)
	}
}

func (s *QualidentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitQualident(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Qualident() (localctx IQualidentContext) {
	localctx = NewQualidentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, oberonParserRULE_qualident)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(116)
			p.Ident()
		}
		{
			p.SetState(117)
			p.Match(oberonParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(121)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentdefContext is an interface to support dynamic dispatch.
type IIdentdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsIdentdefContext differentiates from other interfaces.
	IsIdentdefContext()
}

type IdentdefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentdefContext() *IdentdefContext {
	var p = new(IdentdefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_identdef
	return p
}

func InitEmptyIdentdefContext(p *IdentdefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_identdef
}

func (*IdentdefContext) IsIdentdefContext() {}

func NewIdentdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentdefContext {
	var p = new(IdentdefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_identdef

	return p
}

func (s *IdentdefContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentdefContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *IdentdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterIdentdef(s)
	}
}

func (s *IdentdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitIdentdef(s)
	}
}

func (s *IdentdefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitIdentdef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Identdef() (localctx IIdentdefContext) {
	localctx = NewIdentdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, oberonParserRULE_identdef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Ident()
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__1 {
		{
			p.SetState(124)
			p.Match(oberonParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode
	AllHEXDIGIT() []antlr.TerminalNode
	HEXDIGIT(i int) antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(oberonParserDIGIT)
}

func (s *IntegerContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserDIGIT, i)
}

func (s *IntegerContext) AllHEXDIGIT() []antlr.TerminalNode {
	return s.GetTokens(oberonParserHEXDIGIT)
}

func (s *IntegerContext) HEXDIGIT(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserHEXDIGIT, i)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, oberonParserRULE_integer)
	var _la int

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == oberonParserDIGIT {
			{
				p.SetState(127)
				p.Match(oberonParserDIGIT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(oberonParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserHEXDIGIT {
			{
				p.SetState(133)
				p.Match(oberonParserHEXDIGIT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(139)
			p.Match(oberonParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRealContext is an interface to support dynamic dispatch.
type IRealContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode
	ScaleFactor() IScaleFactorContext

	// IsRealContext differentiates from other interfaces.
	IsRealContext()
}

type RealContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealContext() *RealContext {
	var p = new(RealContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_real
	return p
}

func InitEmptyRealContext(p *RealContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_real
}

func (*RealContext) IsRealContext() {}

func NewRealContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealContext {
	var p = new(RealContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_real

	return p
}

func (s *RealContext) GetParser() antlr.Parser { return s.parser }

func (s *RealContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(oberonParserDIGIT)
}

func (s *RealContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserDIGIT, i)
}

func (s *RealContext) ScaleFactor() IScaleFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScaleFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScaleFactorContext)
}

func (s *RealContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterReal(s)
	}
}

func (s *RealContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitReal(s)
	}
}

func (s *RealContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitReal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Real_() (localctx IRealContext) {
	localctx = NewRealContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, oberonParserRULE_real)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == oberonParserDIGIT {
		{
			p.SetState(142)
			p.Match(oberonParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(147)
		p.Match(oberonParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserDIGIT {
		{
			p.SetState(148)
			p.Match(oberonParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__3 {
		{
			p.SetState(154)
			p.ScaleFactor()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScaleFactorContext is an interface to support dynamic dispatch.
type IScaleFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode

	// IsScaleFactorContext differentiates from other interfaces.
	IsScaleFactorContext()
}

type ScaleFactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScaleFactorContext() *ScaleFactorContext {
	var p = new(ScaleFactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_scaleFactor
	return p
}

func InitEmptyScaleFactorContext(p *ScaleFactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_scaleFactor
}

func (*ScaleFactorContext) IsScaleFactorContext() {}

func NewScaleFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScaleFactorContext {
	var p = new(ScaleFactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_scaleFactor

	return p
}

func (s *ScaleFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *ScaleFactorContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(oberonParserDIGIT)
}

func (s *ScaleFactorContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserDIGIT, i)
}

func (s *ScaleFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScaleFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScaleFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterScaleFactor(s)
	}
}

func (s *ScaleFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitScaleFactor(s)
	}
}

func (s *ScaleFactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitScaleFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ScaleFactor() (localctx IScaleFactorContext) {
	localctx = NewScaleFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, oberonParserRULE_scaleFactor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(oberonParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__4 || _la == oberonParserT__5 {
		{
			p.SetState(158)
			_la = p.GetTokenStream().LA(1)

			if !(_la == oberonParserT__4 || _la == oberonParserT__5) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == oberonParserDIGIT {
		{
			p.SetState(161)
			p.Match(oberonParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext
	Real_() IRealContext

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *NumberContext) Real_() IRealContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, oberonParserRULE_number)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Integer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Real_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstDeclarationContext is an interface to support dynamic dispatch.
type IConstDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identdef() IIdentdefContext
	ConstExpression() IConstExpressionContext

	// IsConstDeclarationContext differentiates from other interfaces.
	IsConstDeclarationContext()
}

type ConstDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDeclarationContext() *ConstDeclarationContext {
	var p = new(ConstDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_constDeclaration
	return p
}

func InitEmptyConstDeclarationContext(p *ConstDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_constDeclaration
}

func (*ConstDeclarationContext) IsConstDeclarationContext() {}

func NewConstDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclarationContext {
	var p = new(ConstDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_constDeclaration

	return p
}

func (s *ConstDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclarationContext) Identdef() IIdentdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentdefContext)
}

func (s *ConstDeclarationContext) ConstExpression() IConstExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstExpressionContext)
}

func (s *ConstDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterConstDeclaration(s)
	}
}

func (s *ConstDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitConstDeclaration(s)
	}
}

func (s *ConstDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitConstDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ConstDeclaration() (localctx IConstDeclarationContext) {
	localctx = NewConstDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, oberonParserRULE_constDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Identdef()
	}
	{
		p.SetState(171)
		p.Match(oberonParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.ConstExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstExpressionContext is an interface to support dynamic dispatch.
type IConstExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsConstExpressionContext differentiates from other interfaces.
	IsConstExpressionContext()
}

type ConstExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstExpressionContext() *ConstExpressionContext {
	var p = new(ConstExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_constExpression
	return p
}

func InitEmptyConstExpressionContext(p *ConstExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_constExpression
}

func (*ConstExpressionContext) IsConstExpressionContext() {}

func NewConstExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstExpressionContext {
	var p = new(ConstExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_constExpression

	return p
}

func (s *ConstExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterConstExpression(s)
	}
}

func (s *ConstExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitConstExpression(s)
	}
}

func (s *ConstExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitConstExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ConstExpression() (localctx IConstExpressionContext) {
	localctx = NewConstExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, oberonParserRULE_constExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identdef() IIdentdefContext
	Type_() IType_Context

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_typeDeclaration
	return p
}

func InitEmptyTypeDeclarationContext(p *TypeDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_typeDeclaration
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) Identdef() IIdentdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentdefContext)
}

func (s *TypeDeclarationContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, oberonParserRULE_typeDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Identdef()
	}
	{
		p.SetState(177)
		p.Match(oberonParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Qualident() IQualidentContext
	ArrayType() IArrayTypeContext
	RecordType() IRecordTypeContext
	PointerType() IPointerTypeContext
	ProcedureType() IProcedureTypeContext

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) Qualident() IQualidentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualidentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualidentContext)
}

func (s *Type_Context) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *Type_Context) RecordType() IRecordTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecordTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecordTypeContext)
}

func (s *Type_Context) PointerType() IPointerTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointerTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointerTypeContext)
}

func (s *Type_Context) ProcedureType() IProcedureTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitType_(s)
	}
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, oberonParserRULE_type_)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case oberonParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Qualident()
		}

	case oberonParserARRAY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.ArrayType()
		}

	case oberonParserRECORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.RecordType()
		}

	case oberonParserPOINTER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.PointerType()
		}

	case oberonParserPROCEDURE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.ProcedureType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ARRAY() antlr.TerminalNode
	AllLength() []ILengthContext
	Length(i int) ILengthContext
	OF() antlr.TerminalNode
	Type_() IType_Context

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(oberonParserARRAY, 0)
}

func (s *ArrayTypeContext) AllLength() []ILengthContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILengthContext); ok {
			len++
		}
	}

	tst := make([]ILengthContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILengthContext); ok {
			tst[i] = t.(ILengthContext)
			i++
		}
	}

	return tst
}

func (s *ArrayTypeContext) Length(i int) ILengthContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthContext)
}

func (s *ArrayTypeContext) OF() antlr.TerminalNode {
	return s.GetToken(oberonParserOF, 0)
}

func (s *ArrayTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, oberonParserRULE_arrayType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(oberonParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Length()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__7 {
		{
			p.SetState(189)
			p.Match(oberonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Length()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.Match(oberonParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILengthContext is an interface to support dynamic dispatch.
type ILengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConstExpression() IConstExpressionContext

	// IsLengthContext differentiates from other interfaces.
	IsLengthContext()
}

type LengthContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthContext() *LengthContext {
	var p = new(LengthContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_length
	return p
}

func InitEmptyLengthContext(p *LengthContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_length
}

func (*LengthContext) IsLengthContext() {}

func NewLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthContext {
	var p = new(LengthContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_length

	return p
}

func (s *LengthContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthContext) ConstExpression() IConstExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstExpressionContext)
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitLength(s)
	}
}

func (s *LengthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitLength(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Length() (localctx ILengthContext) {
	localctx = NewLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, oberonParserRULE_length)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.ConstExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecordTypeContext is an interface to support dynamic dispatch.
type IRecordTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RECORD() antlr.TerminalNode
	END() antlr.TerminalNode
	BaseType() IBaseTypeContext
	FieldListSequence() IFieldListSequenceContext

	// IsRecordTypeContext differentiates from other interfaces.
	IsRecordTypeContext()
}

type RecordTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecordTypeContext() *RecordTypeContext {
	var p = new(RecordTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_recordType
	return p
}

func InitEmptyRecordTypeContext(p *RecordTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_recordType
}

func (*RecordTypeContext) IsRecordTypeContext() {}

func NewRecordTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecordTypeContext {
	var p = new(RecordTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_recordType

	return p
}

func (s *RecordTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RecordTypeContext) RECORD() antlr.TerminalNode {
	return s.GetToken(oberonParserRECORD, 0)
}

func (s *RecordTypeContext) END() antlr.TerminalNode {
	return s.GetToken(oberonParserEND, 0)
}

func (s *RecordTypeContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *RecordTypeContext) FieldListSequence() IFieldListSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListSequenceContext)
}

func (s *RecordTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecordTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterRecordType(s)
	}
}

func (s *RecordTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitRecordType(s)
	}
}

func (s *RecordTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitRecordType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) RecordType() (localctx IRecordTypeContext) {
	localctx = NewRecordTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, oberonParserRULE_recordType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(oberonParserRECORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__8 {
		{
			p.SetState(202)
			p.Match(oberonParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.BaseType()
		}
		{
			p.SetState(204)
			p.Match(oberonParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserIDENT {
		{
			p.SetState(208)
			p.FieldListSequence()
		}

	}
	{
		p.SetState(211)
		p.Match(oberonParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Qualident() IQualidentContext

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_baseType
	return p
}

func InitEmptyBaseTypeContext(p *BaseTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_baseType
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) Qualident() IQualidentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualidentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualidentContext)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, oberonParserRULE_baseType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Qualident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldListSequenceContext is an interface to support dynamic dispatch.
type IFieldListSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldList() []IFieldListContext
	FieldList(i int) IFieldListContext

	// IsFieldListSequenceContext differentiates from other interfaces.
	IsFieldListSequenceContext()
}

type FieldListSequenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListSequenceContext() *FieldListSequenceContext {
	var p = new(FieldListSequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_fieldListSequence
	return p
}

func InitEmptyFieldListSequenceContext(p *FieldListSequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_fieldListSequence
}

func (*FieldListSequenceContext) IsFieldListSequenceContext() {}

func NewFieldListSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListSequenceContext {
	var p = new(FieldListSequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_fieldListSequence

	return p
}

func (s *FieldListSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListSequenceContext) AllFieldList() []IFieldListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldListContext); ok {
			len++
		}
	}

	tst := make([]IFieldListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldListContext); ok {
			tst[i] = t.(IFieldListContext)
			i++
		}
	}

	return tst
}

func (s *FieldListSequenceContext) FieldList(i int) IFieldListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *FieldListSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterFieldListSequence(s)
	}
}

func (s *FieldListSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitFieldListSequence(s)
	}
}

func (s *FieldListSequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitFieldListSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) FieldListSequence() (localctx IFieldListSequenceContext) {
	localctx = NewFieldListSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, oberonParserRULE_fieldListSequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.FieldList()
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__10 {
		{
			p.SetState(216)
			p.Match(oberonParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.FieldList()
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldListContext is an interface to support dynamic dispatch.
type IFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentList() IIdentListContext
	Type_() IType_Context

	// IsFieldListContext differentiates from other interfaces.
	IsFieldListContext()
}

type FieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListContext() *FieldListContext {
	var p = new(FieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_fieldList
	return p
}

func InitEmptyFieldListContext(p *FieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_fieldList
}

func (*FieldListContext) IsFieldListContext() {}

func NewFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListContext {
	var p = new(FieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_fieldList

	return p
}

func (s *FieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListContext) IdentList() IIdentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *FieldListContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterFieldList(s)
	}
}

func (s *FieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitFieldList(s)
	}
}

func (s *FieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) FieldList() (localctx IFieldListContext) {
	localctx = NewFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, oberonParserRULE_fieldList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.IdentList()
	}
	{
		p.SetState(224)
		p.Match(oberonParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentListContext is an interface to support dynamic dispatch.
type IIdentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentdef() []IIdentdefContext
	Identdef(i int) IIdentdefContext

	// IsIdentListContext differentiates from other interfaces.
	IsIdentListContext()
}

type IdentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentListContext() *IdentListContext {
	var p = new(IdentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_identList
	return p
}

func InitEmptyIdentListContext(p *IdentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_identList
}

func (*IdentListContext) IsIdentListContext() {}

func NewIdentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentListContext {
	var p = new(IdentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_identList

	return p
}

func (s *IdentListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentListContext) AllIdentdef() []IIdentdefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentdefContext); ok {
			len++
		}
	}

	tst := make([]IIdentdefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentdefContext); ok {
			tst[i] = t.(IIdentdefContext)
			i++
		}
	}

	return tst
}

func (s *IdentListContext) Identdef(i int) IIdentdefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentdefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentdefContext)
}

func (s *IdentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterIdentList(s)
	}
}

func (s *IdentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitIdentList(s)
	}
}

func (s *IdentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitIdentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) IdentList() (localctx IIdentListContext) {
	localctx = NewIdentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, oberonParserRULE_identList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Identdef()
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__7 {
		{
			p.SetState(228)
			p.Match(oberonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Identdef()
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPointerTypeContext is an interface to support dynamic dispatch.
type IPointerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POINTER() antlr.TerminalNode
	TO() antlr.TerminalNode
	Type_() IType_Context

	// IsPointerTypeContext differentiates from other interfaces.
	IsPointerTypeContext()
}

type PointerTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointerTypeContext() *PointerTypeContext {
	var p = new(PointerTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_pointerType
	return p
}

func InitEmptyPointerTypeContext(p *PointerTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_pointerType
}

func (*PointerTypeContext) IsPointerTypeContext() {}

func NewPointerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerTypeContext {
	var p = new(PointerTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_pointerType

	return p
}

func (s *PointerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerTypeContext) POINTER() antlr.TerminalNode {
	return s.GetToken(oberonParserPOINTER, 0)
}

func (s *PointerTypeContext) TO() antlr.TerminalNode {
	return s.GetToken(oberonParserTO, 0)
}

func (s *PointerTypeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterPointerType(s)
	}
}

func (s *PointerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitPointerType(s)
	}
}

func (s *PointerTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitPointerType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) PointerType() (localctx IPointerTypeContext) {
	localctx = NewPointerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, oberonParserRULE_pointerType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(oberonParserPOINTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Match(oberonParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProcedureTypeContext is an interface to support dynamic dispatch.
type IProcedureTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROCEDURE() antlr.TerminalNode
	FormalParameters() IFormalParametersContext

	// IsProcedureTypeContext differentiates from other interfaces.
	IsProcedureTypeContext()
}

type ProcedureTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureTypeContext() *ProcedureTypeContext {
	var p = new(ProcedureTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureType
	return p
}

func InitEmptyProcedureTypeContext(p *ProcedureTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureType
}

func (*ProcedureTypeContext) IsProcedureTypeContext() {}

func NewProcedureTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureTypeContext {
	var p = new(ProcedureTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_procedureType

	return p
}

func (s *ProcedureTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureTypeContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(oberonParserPROCEDURE, 0)
}

func (s *ProcedureTypeContext) FormalParameters() IFormalParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *ProcedureTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterProcedureType(s)
	}
}

func (s *ProcedureTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitProcedureType(s)
	}
}

func (s *ProcedureTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitProcedureType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ProcedureType() (localctx IProcedureTypeContext) {
	localctx = NewProcedureTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, oberonParserRULE_procedureType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(oberonParserPROCEDURE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__8 {
		{
			p.SetState(240)
			p.FormalParameters()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentList() IIdentListContext
	Type_() IType_Context

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) IdentList() IIdentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *VariableDeclarationContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, oberonParserRULE_variableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.IdentList()
	}
	{
		p.SetState(244)
		p.Match(oberonParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimpleExpression() []ISimpleExpressionContext
	SimpleExpression(i int) ISimpleExpressionContext
	Relation() IRelationContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllSimpleExpression() []ISimpleExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleExpressionContext); ok {
			len++
		}
	}

	tst := make([]ISimpleExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleExpressionContext); ok {
			tst[i] = t.(ISimpleExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) SimpleExpression(i int) ISimpleExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleExpressionContext)
}

func (s *ExpressionContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, oberonParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.SimpleExpression()
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&206158684288) != 0 {
		{
			p.SetState(248)
			p.Relation()
		}
		{
			p.SetState(249)
			p.SimpleExpression()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IN() antlr.TerminalNode
	IS() antlr.TerminalNode

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_relation
	return p
}

func InitEmptyRelationContext(p *RelationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_relation
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) IN() antlr.TerminalNode {
	return s.GetToken(oberonParserIN, 0)
}

func (s *RelationContext) IS() antlr.TerminalNode {
	return s.GetToken(oberonParserIS, 0)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (s *RelationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitRelation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Relation() (localctx IRelationContext) {
	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, oberonParserRULE_relation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&206158684288) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleExpressionContext is an interface to support dynamic dispatch.
type ISimpleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllAddOperator() []IAddOperatorContext
	AddOperator(i int) IAddOperatorContext

	// IsSimpleExpressionContext differentiates from other interfaces.
	IsSimpleExpressionContext()
}

type SimpleExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleExpressionContext() *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_simpleExpression
	return p
}

func InitEmptySimpleExpressionContext(p *SimpleExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_simpleExpression
}

func (*SimpleExpressionContext) IsSimpleExpressionContext() {}

func NewSimpleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleExpressionContext {
	var p = new(SimpleExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_simpleExpression

	return p
}

func (s *SimpleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleExpressionContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *SimpleExpressionContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpleExpressionContext) AllAddOperator() []IAddOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAddOperatorContext); ok {
			len++
		}
	}

	tst := make([]IAddOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAddOperatorContext); ok {
			tst[i] = t.(IAddOperatorContext)
			i++
		}
	}

	return tst
}

func (s *SimpleExpressionContext) AddOperator(i int) IAddOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOperatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddOperatorContext)
}

func (s *SimpleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitSimpleExpression(s)
	}
}

func (s *SimpleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitSimpleExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) SimpleExpression() (localctx ISimpleExpressionContext) {
	localctx = NewSimpleExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, oberonParserRULE_simpleExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__4 || _la == oberonParserT__5 {
		{
			p.SetState(255)
			_la = p.GetTokenStream().LA(1)

			if !(_la == oberonParserT__4 || _la == oberonParserT__5) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(258)
		p.Term()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274877907040) != 0 {
		{
			p.SetState(259)
			p.AddOperator()
		}
		{
			p.SetState(260)
			p.Term()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddOperatorContext is an interface to support dynamic dispatch.
type IAddOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR() antlr.TerminalNode

	// IsAddOperatorContext differentiates from other interfaces.
	IsAddOperatorContext()
}

type AddOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOperatorContext() *AddOperatorContext {
	var p = new(AddOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_addOperator
	return p
}

func InitEmptyAddOperatorContext(p *AddOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_addOperator
}

func (*AddOperatorContext) IsAddOperatorContext() {}

func NewAddOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOperatorContext {
	var p = new(AddOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_addOperator

	return p
}

func (s *AddOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(oberonParserOR, 0)
}

func (s *AddOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterAddOperator(s)
	}
}

func (s *AddOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitAddOperator(s)
	}
}

func (s *AddOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitAddOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) AddOperator() (localctx IAddOperatorContext) {
	localctx = NewAddOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, oberonParserRULE_addOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274877907040) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllMulOperator() []IMulOperatorContext
	MulOperator(i int) IMulOperatorContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMulOperator() []IMulOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMulOperatorContext); ok {
			len++
		}
	}

	tst := make([]IMulOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMulOperatorContext); ok {
			tst[i] = t.(IMulOperatorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) MulOperator(i int) IMulOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulOperatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulOperatorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, oberonParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Factor()
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1649268228100) != 0 {
		{
			p.SetState(270)
			p.MulOperator()
		}
		{
			p.SetState(271)
			p.Factor()
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMulOperatorContext is an interface to support dynamic dispatch.
type IMulOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsMulOperatorContext differentiates from other interfaces.
	IsMulOperatorContext()
}

type MulOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulOperatorContext() *MulOperatorContext {
	var p = new(MulOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_mulOperator
	return p
}

func InitEmptyMulOperatorContext(p *MulOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_mulOperator
}

func (*MulOperatorContext) IsMulOperatorContext() {}

func NewMulOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulOperatorContext {
	var p = new(MulOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_mulOperator

	return p
}

func (s *MulOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MulOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(oberonParserDIV, 0)
}

func (s *MulOperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(oberonParserMOD, 0)
}

func (s *MulOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterMulOperator(s)
	}
}

func (s *MulOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitMulOperator(s)
	}
}

func (s *MulOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitMulOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) MulOperator() (localctx IMulOperatorContext) {
	localctx = NewMulOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, oberonParserRULE_mulOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1649268228100) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	STRING() antlr.TerminalNode
	NIL() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	Set_() ISet_Context
	Designator() IDesignatorContext
	ActualParameters() IActualParametersContext
	Expression() IExpressionContext
	Factor() IFactorContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *FactorContext) STRING() antlr.TerminalNode {
	return s.GetToken(oberonParserSTRING, 0)
}

func (s *FactorContext) NIL() antlr.TerminalNode {
	return s.GetToken(oberonParserNIL, 0)
}

func (s *FactorContext) TRUE() antlr.TerminalNode {
	return s.GetToken(oberonParserTRUE, 0)
}

func (s *FactorContext) FALSE() antlr.TerminalNode {
	return s.GetToken(oberonParserFALSE, 0)
}

func (s *FactorContext) Set_() ISet_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_Context)
}

func (s *FactorContext) Designator() IDesignatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *FactorContext) ActualParameters() IActualParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActualParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActualParametersContext)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, oberonParserRULE_factor)
	var _la int

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case oberonParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Number()
		}

	case oberonParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.Match(oberonParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserNIL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(282)
			p.Match(oberonParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(283)
			p.Match(oberonParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserFALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(284)
			p.Match(oberonParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserT__23:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(285)
			p.Set_()
		}

	case oberonParserIDENT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(286)
			p.Designator()
		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == oberonParserT__8 {
			{
				p.SetState(287)
				p.ActualParameters()
			}

		}

	case oberonParserT__8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(290)
			p.Match(oberonParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.Expression()
		}
		{
			p.SetState(292)
			p.Match(oberonParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserT__19:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(294)
			p.Match(oberonParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.Factor()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDesignatorContext is an interface to support dynamic dispatch.
type IDesignatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Qualident() IQualidentContext
	AllSelector() []ISelectorContext
	Selector(i int) ISelectorContext

	// IsDesignatorContext differentiates from other interfaces.
	IsDesignatorContext()
}

type DesignatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesignatorContext() *DesignatorContext {
	var p = new(DesignatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_designator
	return p
}

func InitEmptyDesignatorContext(p *DesignatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_designator
}

func (*DesignatorContext) IsDesignatorContext() {}

func NewDesignatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesignatorContext {
	var p = new(DesignatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_designator

	return p
}

func (s *DesignatorContext) GetParser() antlr.Parser { return s.parser }

func (s *DesignatorContext) Qualident() IQualidentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualidentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualidentContext)
}

func (s *DesignatorContext) AllSelector() []ISelectorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectorContext); ok {
			len++
		}
	}

	tst := make([]ISelectorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectorContext); ok {
			tst[i] = t.(ISelectorContext)
			i++
		}
	}

	return tst
}

func (s *DesignatorContext) Selector(i int) ISelectorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *DesignatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesignatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesignatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterDesignator(s)
	}
}

func (s *DesignatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitDesignator(s)
	}
}

func (s *DesignatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitDesignator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Designator() (localctx IDesignatorContext) {
	localctx = NewDesignatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, oberonParserRULE_designator)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Qualident()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(299)
				p.Selector()
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	ExpList() IExpListContext
	Qualident() IQualidentContext

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *SelectorContext) ExpList() IExpListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpListContext)
}

func (s *SelectorContext) Qualident() IQualidentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualidentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualidentContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, oberonParserRULE_selector)
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case oberonParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.Match(oberonParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Ident()
		}

	case oberonParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			p.Match(oberonParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.ExpList()
		}
		{
			p.SetState(309)
			p.Match(oberonParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(311)
			p.Match(oberonParserT__22)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(312)
			p.Match(oberonParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Qualident()
		}
		{
			p.SetState(314)
			p.Match(oberonParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISet_Context is an interface to support dynamic dispatch.
type ISet_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsSet_Context differentiates from other interfaces.
	IsSet_Context()
}

type Set_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_Context() *Set_Context {
	var p = new(Set_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_set_
	return p
}

func InitEmptySet_Context(p *Set_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_set_
}

func (*Set_Context) IsSet_Context() {}

func NewSet_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_Context {
	var p = new(Set_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_set_

	return p
}

func (s *Set_Context) GetParser() antlr.Parser { return s.parser }

func (s *Set_Context) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *Set_Context) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *Set_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterSet_(s)
	}
}

func (s *Set_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitSet_(s)
	}
}

func (s *Set_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitSet_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Set_() (localctx ISet_Context) {
	localctx = NewSet_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, oberonParserRULE_set_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(oberonParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&3026419430629867539) != 0 {
		{
			p.SetState(319)
			p.Element()
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserT__7 {
			{
				p.SetState(320)
				p.Match(oberonParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(321)
				p.Element()
			}

			p.SetState(326)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(329)
		p.Match(oberonParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_element
	return p
}

func InitEmptyElementContext(p *ElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_element
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ElementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitElement(s)
	}
}

func (s *ElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, oberonParserRULE_element)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(331)
		p.Expression()
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__25 {
		{
			p.SetState(332)
			p.Match(oberonParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Expression()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpListContext is an interface to support dynamic dispatch.
type IExpListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpListContext differentiates from other interfaces.
	IsExpListContext()
}

type ExpListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpListContext() *ExpListContext {
	var p = new(ExpListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_expList
	return p
}

func InitEmptyExpListContext(p *ExpListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_expList
}

func (*ExpListContext) IsExpListContext() {}

func NewExpListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpListContext {
	var p = new(ExpListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_expList

	return p
}

func (s *ExpListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterExpList(s)
	}
}

func (s *ExpListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitExpList(s)
	}
}

func (s *ExpListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitExpList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ExpList() (localctx IExpListContext) {
	localctx = NewExpListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, oberonParserRULE_expList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		p.Expression()
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__7 {
		{
			p.SetState(337)
			p.Match(oberonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Expression()
		}

		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActualParametersContext is an interface to support dynamic dispatch.
type IActualParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpList() IExpListContext

	// IsActualParametersContext differentiates from other interfaces.
	IsActualParametersContext()
}

type ActualParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActualParametersContext() *ActualParametersContext {
	var p = new(ActualParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_actualParameters
	return p
}

func InitEmptyActualParametersContext(p *ActualParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_actualParameters
}

func (*ActualParametersContext) IsActualParametersContext() {}

func NewActualParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActualParametersContext {
	var p = new(ActualParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_actualParameters

	return p
}

func (s *ActualParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ActualParametersContext) ExpList() IExpListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpListContext)
}

func (s *ActualParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActualParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActualParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterActualParameters(s)
	}
}

func (s *ActualParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitActualParameters(s)
	}
}

func (s *ActualParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitActualParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ActualParameters() (localctx IActualParametersContext) {
	localctx = NewActualParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, oberonParserRULE_actualParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(oberonParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&3026419430629867539) != 0 {
		{
			p.SetState(345)
			p.ExpList()
		}

	}
	{
		p.SetState(348)
		p.Match(oberonParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	ProcedureCall() IProcedureCallContext
	IfStatement() IIfStatementContext
	CaseStatement() ICaseStatementContext
	WhileStatement() IWhileStatementContext
	RepeatStatement() IRepeatStatementContext
	ForStatement() IForStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ProcedureCall() IProcedureCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureCallContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) CaseStatement() ICaseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) RepeatStatement() IRepeatStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, oberonParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(357)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(350)
			p.Assignment()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(351)
			p.ProcedureCall()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(352)
			p.IfStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 4 {
		{
			p.SetState(353)
			p.CaseStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 5 {
		{
			p.SetState(354)
			p.WhileStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 6 {
		{
			p.SetState(355)
			p.RepeatStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 7 {
		{
			p.SetState(356)
			p.ForStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Designator() IDesignatorContext
	Expression() IExpressionContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Designator() IDesignatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, oberonParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Designator()
	}
	{
		p.SetState(360)
		p.Match(oberonParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProcedureCallContext is an interface to support dynamic dispatch.
type IProcedureCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Designator() IDesignatorContext
	ActualParameters() IActualParametersContext

	// IsProcedureCallContext differentiates from other interfaces.
	IsProcedureCallContext()
}

type ProcedureCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureCallContext() *ProcedureCallContext {
	var p = new(ProcedureCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureCall
	return p
}

func InitEmptyProcedureCallContext(p *ProcedureCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureCall
}

func (*ProcedureCallContext) IsProcedureCallContext() {}

func NewProcedureCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureCallContext {
	var p = new(ProcedureCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_procedureCall

	return p
}

func (s *ProcedureCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureCallContext) Designator() IDesignatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesignatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesignatorContext)
}

func (s *ProcedureCallContext) ActualParameters() IActualParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActualParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActualParametersContext)
}

func (s *ProcedureCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterProcedureCall(s)
	}
}

func (s *ProcedureCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitProcedureCall(s)
	}
}

func (s *ProcedureCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitProcedureCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ProcedureCall() (localctx IProcedureCallContext) {
	localctx = NewProcedureCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, oberonParserRULE_procedureCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Designator()
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__8 {
		{
			p.SetState(364)
			p.ActualParameters()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementSequenceContext is an interface to support dynamic dispatch.
type IStatementSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementSequenceContext differentiates from other interfaces.
	IsStatementSequenceContext()
}

type StatementSequenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementSequenceContext() *StatementSequenceContext {
	var p = new(StatementSequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_statementSequence
	return p
}

func InitEmptyStatementSequenceContext(p *StatementSequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_statementSequence
}

func (*StatementSequenceContext) IsStatementSequenceContext() {}

func NewStatementSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementSequenceContext {
	var p = new(StatementSequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_statementSequence

	return p
}

func (s *StatementSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementSequenceContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementSequenceContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterStatementSequence(s)
	}
}

func (s *StatementSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitStatementSequence(s)
	}
}

func (s *StatementSequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitStatementSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) StatementSequence() (localctx IStatementSequenceContext) {
	localctx = NewStatementSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, oberonParserRULE_statementSequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.Statement()
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__10 {
		{
			p.SetState(368)
			p.Match(oberonParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Statement()
		}

		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllTHEN() []antlr.TerminalNode
	THEN(i int) antlr.TerminalNode
	AllStatementSequence() []IStatementSequenceContext
	StatementSequence(i int) IStatementSequenceContext
	END() antlr.TerminalNode
	AllELSIF() []antlr.TerminalNode
	ELSIF(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(oberonParserIF, 0)
}

func (s *IfStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllTHEN() []antlr.TerminalNode {
	return s.GetTokens(oberonParserTHEN)
}

func (s *IfStatementContext) THEN(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserTHEN, i)
}

func (s *IfStatementContext) AllStatementSequence() []IStatementSequenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			len++
		}
	}

	tst := make([]IStatementSequenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementSequenceContext); ok {
			tst[i] = t.(IStatementSequenceContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) StatementSequence(i int) IStatementSequenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *IfStatementContext) END() antlr.TerminalNode {
	return s.GetToken(oberonParserEND, 0)
}

func (s *IfStatementContext) AllELSIF() []antlr.TerminalNode {
	return s.GetTokens(oberonParserELSIF)
}

func (s *IfStatementContext) ELSIF(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserELSIF, i)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(oberonParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, oberonParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.Match(oberonParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(376)
		p.Expression()
	}
	{
		p.SetState(377)
		p.Match(oberonParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.StatementSequence()
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserELSIF {
		{
			p.SetState(379)
			p.Match(oberonParserELSIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Expression()
		}
		{
			p.SetState(381)
			p.Match(oberonParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)
			p.StatementSequence()
		}

		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserELSE {
		{
			p.SetState(389)
			p.Match(oberonParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.StatementSequence()
		}

	}
	{
		p.SetState(393)
		p.Match(oberonParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseStatementContext is an interface to support dynamic dispatch.
type ICaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	OF() antlr.TerminalNode
	AllCase_() []ICase_Context
	Case_(i int) ICase_Context
	END() antlr.TerminalNode

	// IsCaseStatementContext differentiates from other interfaces.
	IsCaseStatementContext()
}

type CaseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseStatementContext() *CaseStatementContext {
	var p = new(CaseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_caseStatement
	return p
}

func InitEmptyCaseStatementContext(p *CaseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_caseStatement
}

func (*CaseStatementContext) IsCaseStatementContext() {}

func NewCaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStatementContext {
	var p = new(CaseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_caseStatement

	return p
}

func (s *CaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStatementContext) CASE() antlr.TerminalNode {
	return s.GetToken(oberonParserCASE, 0)
}

func (s *CaseStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseStatementContext) OF() antlr.TerminalNode {
	return s.GetToken(oberonParserOF, 0)
}

func (s *CaseStatementContext) AllCase_() []ICase_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICase_Context); ok {
			len++
		}
	}

	tst := make([]ICase_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICase_Context); ok {
			tst[i] = t.(ICase_Context)
			i++
		}
	}

	return tst
}

func (s *CaseStatementContext) Case_(i int) ICase_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICase_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICase_Context)
}

func (s *CaseStatementContext) END() antlr.TerminalNode {
	return s.GetToken(oberonParserEND, 0)
}

func (s *CaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterCaseStatement(s)
	}
}

func (s *CaseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitCaseStatement(s)
	}
}

func (s *CaseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitCaseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) CaseStatement() (localctx ICaseStatementContext) {
	localctx = NewCaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, oberonParserRULE_caseStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(oberonParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Expression()
	}
	{
		p.SetState(397)
		p.Match(oberonParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Case_()
	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__27 {
		{
			p.SetState(399)
			p.Match(oberonParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Case_()
		}

		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(406)
		p.Match(oberonParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICase_Context is an interface to support dynamic dispatch.
type ICase_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CaseLabelList() ICaseLabelListContext
	StatementSequence() IStatementSequenceContext

	// IsCase_Context differentiates from other interfaces.
	IsCase_Context()
}

type Case_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_Context() *Case_Context {
	var p = new(Case_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_case_
	return p
}

func InitEmptyCase_Context(p *Case_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_case_
}

func (*Case_Context) IsCase_Context() {}

func NewCase_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_Context {
	var p = new(Case_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_case_

	return p
}

func (s *Case_Context) GetParser() antlr.Parser { return s.parser }

func (s *Case_Context) CaseLabelList() ICaseLabelListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseLabelListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseLabelListContext)
}

func (s *Case_Context) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *Case_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterCase_(s)
	}
}

func (s *Case_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitCase_(s)
	}
}

func (s *Case_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitCase_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Case_() (localctx ICase_Context) {
	localctx = NewCase_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, oberonParserRULE_case_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-62)) & ^0x3f) == 0 && ((int64(1)<<(_la-62))&21) != 0 {
		{
			p.SetState(408)
			p.CaseLabelList()
		}
		{
			p.SetState(409)
			p.Match(oberonParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.StatementSequence()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseLabelListContext is an interface to support dynamic dispatch.
type ICaseLabelListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLabelRange() []ILabelRangeContext
	LabelRange(i int) ILabelRangeContext

	// IsCaseLabelListContext differentiates from other interfaces.
	IsCaseLabelListContext()
}

type CaseLabelListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseLabelListContext() *CaseLabelListContext {
	var p = new(CaseLabelListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_caseLabelList
	return p
}

func InitEmptyCaseLabelListContext(p *CaseLabelListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_caseLabelList
}

func (*CaseLabelListContext) IsCaseLabelListContext() {}

func NewCaseLabelListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseLabelListContext {
	var p = new(CaseLabelListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_caseLabelList

	return p
}

func (s *CaseLabelListContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseLabelListContext) AllLabelRange() []ILabelRangeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelRangeContext); ok {
			len++
		}
	}

	tst := make([]ILabelRangeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelRangeContext); ok {
			tst[i] = t.(ILabelRangeContext)
			i++
		}
	}

	return tst
}

func (s *CaseLabelListContext) LabelRange(i int) ILabelRangeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelRangeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelRangeContext)
}

func (s *CaseLabelListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseLabelListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseLabelListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterCaseLabelList(s)
	}
}

func (s *CaseLabelListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitCaseLabelList(s)
	}
}

func (s *CaseLabelListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitCaseLabelList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) CaseLabelList() (localctx ICaseLabelListContext) {
	localctx = NewCaseLabelListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, oberonParserRULE_caseLabelList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.LabelRange()
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__7 {
		{
			p.SetState(415)
			p.Match(oberonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.LabelRange()
		}

		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelRangeContext is an interface to support dynamic dispatch.
type ILabelRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLabel() []ILabelContext
	Label(i int) ILabelContext

	// IsLabelRangeContext differentiates from other interfaces.
	IsLabelRangeContext()
}

type LabelRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelRangeContext() *LabelRangeContext {
	var p = new(LabelRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_labelRange
	return p
}

func InitEmptyLabelRangeContext(p *LabelRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_labelRange
}

func (*LabelRangeContext) IsLabelRangeContext() {}

func NewLabelRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelRangeContext {
	var p = new(LabelRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_labelRange

	return p
}

func (s *LabelRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelRangeContext) AllLabel() []ILabelContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelContext); ok {
			len++
		}
	}

	tst := make([]ILabelContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelContext); ok {
			tst[i] = t.(ILabelContext)
			i++
		}
	}

	return tst
}

func (s *LabelRangeContext) Label(i int) ILabelContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LabelRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterLabelRange(s)
	}
}

func (s *LabelRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitLabelRange(s)
	}
}

func (s *LabelRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitLabelRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) LabelRange() (localctx ILabelRangeContext) {
	localctx = NewLabelRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, oberonParserRULE_labelRange)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Label()
	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__25 {
		{
			p.SetState(423)
			p.Match(oberonParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Label()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer() IIntegerContext
	STRING() antlr.TerminalNode
	Qualident() IQualidentContext

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_label
	return p
}

func InitEmptyLabelContext(p *LabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_label
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *LabelContext) STRING() antlr.TerminalNode {
	return s.GetToken(oberonParserSTRING, 0)
}

func (s *LabelContext) Qualident() IQualidentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualidentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualidentContext)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, oberonParserRULE_label)
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case oberonParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(427)
			p.Integer()
		}

	case oberonParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Match(oberonParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(429)
			p.Qualident()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllDO() []antlr.TerminalNode
	DO(i int) antlr.TerminalNode
	AllStatementSequence() []IStatementSequenceContext
	StatementSequence(i int) IStatementSequenceContext
	END() antlr.TerminalNode
	AllELSIF() []antlr.TerminalNode
	ELSIF(i int) antlr.TerminalNode

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(oberonParserWHILE, 0)
}

func (s *WhileStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *WhileStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) AllDO() []antlr.TerminalNode {
	return s.GetTokens(oberonParserDO)
}

func (s *WhileStatementContext) DO(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserDO, i)
}

func (s *WhileStatementContext) AllStatementSequence() []IStatementSequenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			len++
		}
	}

	tst := make([]IStatementSequenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementSequenceContext); ok {
			tst[i] = t.(IStatementSequenceContext)
			i++
		}
	}

	return tst
}

func (s *WhileStatementContext) StatementSequence(i int) IStatementSequenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *WhileStatementContext) END() antlr.TerminalNode {
	return s.GetToken(oberonParserEND, 0)
}

func (s *WhileStatementContext) AllELSIF() []antlr.TerminalNode {
	return s.GetTokens(oberonParserELSIF)
}

func (s *WhileStatementContext) ELSIF(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserELSIF, i)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, oberonParserRULE_whileStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(oberonParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Expression()
	}
	{
		p.SetState(434)
		p.Match(oberonParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(435)
		p.StatementSequence()
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserELSIF {
		{
			p.SetState(436)
			p.Match(oberonParserELSIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.Expression()
		}
		{
			p.SetState(438)
			p.Match(oberonParserDO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.StatementSequence()
		}

		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(446)
		p.Match(oberonParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepeatStatementContext is an interface to support dynamic dispatch.
type IRepeatStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REPEAT() antlr.TerminalNode
	StatementSequence() IStatementSequenceContext
	UNTIL() antlr.TerminalNode
	Expression() IExpressionContext

	// IsRepeatStatementContext differentiates from other interfaces.
	IsRepeatStatementContext()
}

type RepeatStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatStatementContext() *RepeatStatementContext {
	var p = new(RepeatStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_repeatStatement
	return p
}

func InitEmptyRepeatStatementContext(p *RepeatStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_repeatStatement
}

func (*RepeatStatementContext) IsRepeatStatementContext() {}

func NewRepeatStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatStatementContext {
	var p = new(RepeatStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_repeatStatement

	return p
}

func (s *RepeatStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatStatementContext) REPEAT() antlr.TerminalNode {
	return s.GetToken(oberonParserREPEAT, 0)
}

func (s *RepeatStatementContext) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *RepeatStatementContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(oberonParserUNTIL, 0)
}

func (s *RepeatStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RepeatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterRepeatStatement(s)
	}
}

func (s *RepeatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitRepeatStatement(s)
	}
}

func (s *RepeatStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitRepeatStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) RepeatStatement() (localctx IRepeatStatementContext) {
	localctx = NewRepeatStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, oberonParserRULE_repeatStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(oberonParserREPEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.StatementSequence()
	}
	{
		p.SetState(450)
		p.Match(oberonParserUNTIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Ident() IIdentContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	TO() antlr.TerminalNode
	DO() antlr.TerminalNode
	StatementSequence() IStatementSequenceContext
	END() antlr.TerminalNode
	BY() antlr.TerminalNode
	ConstExpression() IConstExpressionContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(oberonParserFOR, 0)
}

func (s *ForStatementContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ForStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) TO() antlr.TerminalNode {
	return s.GetToken(oberonParserTO, 0)
}

func (s *ForStatementContext) DO() antlr.TerminalNode {
	return s.GetToken(oberonParserDO, 0)
}

func (s *ForStatementContext) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *ForStatementContext) END() antlr.TerminalNode {
	return s.GetToken(oberonParserEND, 0)
}

func (s *ForStatementContext) BY() antlr.TerminalNode {
	return s.GetToken(oberonParserBY, 0)
}

func (s *ForStatementContext) ConstExpression() IConstExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstExpressionContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, oberonParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.Match(oberonParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(454)
		p.Ident()
	}
	{
		p.SetState(455)
		p.Match(oberonParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.Expression()
	}
	{
		p.SetState(457)
		p.Match(oberonParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
		p.Expression()
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserBY {
		{
			p.SetState(459)
			p.Match(oberonParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)
			p.ConstExpression()
		}

	}
	{
		p.SetState(463)
		p.Match(oberonParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(464)
		p.StatementSequence()
	}
	{
		p.SetState(465)
		p.Match(oberonParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProcedureDeclarationContext is an interface to support dynamic dispatch.
type IProcedureDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ProcedureHeading() IProcedureHeadingContext
	ProcedureBody() IProcedureBodyContext
	Ident() IIdentContext

	// IsProcedureDeclarationContext differentiates from other interfaces.
	IsProcedureDeclarationContext()
}

type ProcedureDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureDeclarationContext() *ProcedureDeclarationContext {
	var p = new(ProcedureDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureDeclaration
	return p
}

func InitEmptyProcedureDeclarationContext(p *ProcedureDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureDeclaration
}

func (*ProcedureDeclarationContext) IsProcedureDeclarationContext() {}

func NewProcedureDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureDeclarationContext {
	var p = new(ProcedureDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_procedureDeclaration

	return p
}

func (s *ProcedureDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureDeclarationContext) ProcedureHeading() IProcedureHeadingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureHeadingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureHeadingContext)
}

func (s *ProcedureDeclarationContext) ProcedureBody() IProcedureBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureBodyContext)
}

func (s *ProcedureDeclarationContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ProcedureDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterProcedureDeclaration(s)
	}
}

func (s *ProcedureDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitProcedureDeclaration(s)
	}
}

func (s *ProcedureDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitProcedureDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ProcedureDeclaration() (localctx IProcedureDeclarationContext) {
	localctx = NewProcedureDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, oberonParserRULE_procedureDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(467)
		p.ProcedureHeading()
	}
	{
		p.SetState(468)
		p.Match(oberonParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
		p.ProcedureBody()
	}
	{
		p.SetState(470)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProcedureHeadingContext is an interface to support dynamic dispatch.
type IProcedureHeadingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROCEDURE() antlr.TerminalNode
	Identdef() IIdentdefContext
	FormalParameters() IFormalParametersContext

	// IsProcedureHeadingContext differentiates from other interfaces.
	IsProcedureHeadingContext()
}

type ProcedureHeadingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureHeadingContext() *ProcedureHeadingContext {
	var p = new(ProcedureHeadingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureHeading
	return p
}

func InitEmptyProcedureHeadingContext(p *ProcedureHeadingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureHeading
}

func (*ProcedureHeadingContext) IsProcedureHeadingContext() {}

func NewProcedureHeadingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureHeadingContext {
	var p = new(ProcedureHeadingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_procedureHeading

	return p
}

func (s *ProcedureHeadingContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureHeadingContext) PROCEDURE() antlr.TerminalNode {
	return s.GetToken(oberonParserPROCEDURE, 0)
}

func (s *ProcedureHeadingContext) Identdef() IIdentdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentdefContext)
}

func (s *ProcedureHeadingContext) FormalParameters() IFormalParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *ProcedureHeadingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureHeadingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureHeadingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterProcedureHeading(s)
	}
}

func (s *ProcedureHeadingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitProcedureHeading(s)
	}
}

func (s *ProcedureHeadingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitProcedureHeading(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ProcedureHeading() (localctx IProcedureHeadingContext) {
	localctx = NewProcedureHeadingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, oberonParserRULE_procedureHeading)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.Match(oberonParserPROCEDURE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(473)
		p.Identdef()
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__8 {
		{
			p.SetState(474)
			p.FormalParameters()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProcedureBodyContext is an interface to support dynamic dispatch.
type IProcedureBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclarationSequence() IDeclarationSequenceContext
	END() antlr.TerminalNode
	BEGIN() antlr.TerminalNode
	StatementSequence() IStatementSequenceContext
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsProcedureBodyContext differentiates from other interfaces.
	IsProcedureBodyContext()
}

type ProcedureBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureBodyContext() *ProcedureBodyContext {
	var p = new(ProcedureBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureBody
	return p
}

func InitEmptyProcedureBodyContext(p *ProcedureBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_procedureBody
}

func (*ProcedureBodyContext) IsProcedureBodyContext() {}

func NewProcedureBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureBodyContext {
	var p = new(ProcedureBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_procedureBody

	return p
}

func (s *ProcedureBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureBodyContext) DeclarationSequence() IDeclarationSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSequenceContext)
}

func (s *ProcedureBodyContext) END() antlr.TerminalNode {
	return s.GetToken(oberonParserEND, 0)
}

func (s *ProcedureBodyContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(oberonParserBEGIN, 0)
}

func (s *ProcedureBodyContext) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *ProcedureBodyContext) RETURN() antlr.TerminalNode {
	return s.GetToken(oberonParserRETURN, 0)
}

func (s *ProcedureBodyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProcedureBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterProcedureBody(s)
	}
}

func (s *ProcedureBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitProcedureBody(s)
	}
}

func (s *ProcedureBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitProcedureBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ProcedureBody() (localctx IProcedureBodyContext) {
	localctx = NewProcedureBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, oberonParserRULE_procedureBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.DeclarationSequence()
	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserBEGIN {
		{
			p.SetState(478)
			p.Match(oberonParserBEGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.StatementSequence()
		}

	}
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserRETURN {
		{
			p.SetState(482)
			p.Match(oberonParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.Expression()
		}

	}
	{
		p.SetState(486)
		p.Match(oberonParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationSequenceContext is an interface to support dynamic dispatch.
type IDeclarationSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	VAR() antlr.TerminalNode
	AllProcedureDeclaration() []IProcedureDeclarationContext
	ProcedureDeclaration(i int) IProcedureDeclarationContext
	AllConstDeclaration() []IConstDeclarationContext
	ConstDeclaration(i int) IConstDeclarationContext
	AllTypeDeclaration() []ITypeDeclarationContext
	TypeDeclaration(i int) ITypeDeclarationContext
	AllVariableDeclaration() []IVariableDeclarationContext
	VariableDeclaration(i int) IVariableDeclarationContext

	// IsDeclarationSequenceContext differentiates from other interfaces.
	IsDeclarationSequenceContext()
}

type DeclarationSequenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationSequenceContext() *DeclarationSequenceContext {
	var p = new(DeclarationSequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_declarationSequence
	return p
}

func InitEmptyDeclarationSequenceContext(p *DeclarationSequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_declarationSequence
}

func (*DeclarationSequenceContext) IsDeclarationSequenceContext() {}

func NewDeclarationSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationSequenceContext {
	var p = new(DeclarationSequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_declarationSequence

	return p
}

func (s *DeclarationSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationSequenceContext) CONST() antlr.TerminalNode {
	return s.GetToken(oberonParserCONST, 0)
}

func (s *DeclarationSequenceContext) TYPE() antlr.TerminalNode {
	return s.GetToken(oberonParserTYPE, 0)
}

func (s *DeclarationSequenceContext) VAR() antlr.TerminalNode {
	return s.GetToken(oberonParserVAR, 0)
}

func (s *DeclarationSequenceContext) AllProcedureDeclaration() []IProcedureDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProcedureDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IProcedureDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProcedureDeclarationContext); ok {
			tst[i] = t.(IProcedureDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationSequenceContext) ProcedureDeclaration(i int) IProcedureDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProcedureDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProcedureDeclarationContext)
}

func (s *DeclarationSequenceContext) AllConstDeclaration() []IConstDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IConstDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstDeclarationContext); ok {
			tst[i] = t.(IConstDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationSequenceContext) ConstDeclaration(i int) IConstDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclarationContext)
}

func (s *DeclarationSequenceContext) AllTypeDeclaration() []ITypeDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclarationContext); ok {
			tst[i] = t.(ITypeDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationSequenceContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *DeclarationSequenceContext) AllVariableDeclaration() []IVariableDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclarationContext); ok {
			tst[i] = t.(IVariableDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationSequenceContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *DeclarationSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterDeclarationSequence(s)
	}
}

func (s *DeclarationSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitDeclarationSequence(s)
	}
}

func (s *DeclarationSequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitDeclarationSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) DeclarationSequence() (localctx IDeclarationSequenceContext) {
	localctx = NewDeclarationSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, oberonParserRULE_declarationSequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserCONST {
		{
			p.SetState(488)
			p.Match(oberonParserCONST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserIDENT {
			{
				p.SetState(489)
				p.ConstDeclaration()
			}
			{
				p.SetState(490)
				p.Match(oberonParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(496)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserTYPE {
		{
			p.SetState(499)
			p.Match(oberonParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserIDENT {
			{
				p.SetState(500)
				p.TypeDeclaration()
			}
			{
				p.SetState(501)
				p.Match(oberonParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(507)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserVAR {
		{
			p.SetState(510)
			p.Match(oberonParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserIDENT {
			{
				p.SetState(511)
				p.VariableDeclaration()
			}
			{
				p.SetState(512)
				p.Match(oberonParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(518)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserPROCEDURE {
		{
			p.SetState(521)
			p.ProcedureDeclaration()
		}
		{
			p.SetState(522)
			p.Match(oberonParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFPSection() []IFPSectionContext
	FPSection(i int) IFPSectionContext
	Qualident() IQualidentContext

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_formalParameters
	return p
}

func InitEmptyFormalParametersContext(p *FormalParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_formalParameters
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) AllFPSection() []IFPSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFPSectionContext); ok {
			len++
		}
	}

	tst := make([]IFPSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFPSectionContext); ok {
			tst[i] = t.(IFPSectionContext)
			i++
		}
	}

	return tst
}

func (s *FormalParametersContext) FPSection(i int) IFPSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFPSectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFPSectionContext)
}

func (s *FormalParametersContext) Qualident() IQualidentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualidentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualidentContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitFormalParameters(s)
	}
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) FormalParameters() (localctx IFormalParametersContext) {
	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, oberonParserRULE_formalParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
		p.Match(oberonParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserVAR || _la == oberonParserIDENT {
		{
			p.SetState(530)
			p.FPSection()
		}
		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserT__10 {
			{
				p.SetState(531)
				p.Match(oberonParserT__10)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(532)
				p.FPSection()
			}

			p.SetState(537)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(540)
		p.Match(oberonParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__11 {
		{
			p.SetState(541)
			p.Match(oberonParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)
			p.Qualident()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFPSectionContext is an interface to support dynamic dispatch.
type IFPSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	FormalType() IFormalTypeContext
	VAR() antlr.TerminalNode

	// IsFPSectionContext differentiates from other interfaces.
	IsFPSectionContext()
}

type FPSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFPSectionContext() *FPSectionContext {
	var p = new(FPSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_fPSection
	return p
}

func InitEmptyFPSectionContext(p *FPSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_fPSection
}

func (*FPSectionContext) IsFPSectionContext() {}

func NewFPSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FPSectionContext {
	var p = new(FPSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_fPSection

	return p
}

func (s *FPSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *FPSectionContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *FPSectionContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FPSectionContext) FormalType() IFormalTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalTypeContext)
}

func (s *FPSectionContext) VAR() antlr.TerminalNode {
	return s.GetToken(oberonParserVAR, 0)
}

func (s *FPSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FPSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FPSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterFPSection(s)
	}
}

func (s *FPSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitFPSection(s)
	}
}

func (s *FPSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitFPSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) FPSection() (localctx IFPSectionContext) {
	localctx = NewFPSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, oberonParserRULE_fPSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserVAR {
		{
			p.SetState(545)
			p.Match(oberonParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(548)
		p.Ident()
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__7 {
		{
			p.SetState(549)
			p.Match(oberonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(550)
			p.Ident()
		}

		p.SetState(555)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(556)
		p.Match(oberonParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(557)
		p.FormalType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFormalTypeContext is an interface to support dynamic dispatch.
type IFormalTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Qualident() IQualidentContext
	AllARRAY() []antlr.TerminalNode
	ARRAY(i int) antlr.TerminalNode
	AllOF() []antlr.TerminalNode
	OF(i int) antlr.TerminalNode

	// IsFormalTypeContext differentiates from other interfaces.
	IsFormalTypeContext()
}

type FormalTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalTypeContext() *FormalTypeContext {
	var p = new(FormalTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_formalType
	return p
}

func InitEmptyFormalTypeContext(p *FormalTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_formalType
}

func (*FormalTypeContext) IsFormalTypeContext() {}

func NewFormalTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalTypeContext {
	var p = new(FormalTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_formalType

	return p
}

func (s *FormalTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalTypeContext) Qualident() IQualidentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualidentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualidentContext)
}

func (s *FormalTypeContext) AllARRAY() []antlr.TerminalNode {
	return s.GetTokens(oberonParserARRAY)
}

func (s *FormalTypeContext) ARRAY(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserARRAY, i)
}

func (s *FormalTypeContext) AllOF() []antlr.TerminalNode {
	return s.GetTokens(oberonParserOF)
}

func (s *FormalTypeContext) OF(i int) antlr.TerminalNode {
	return s.GetToken(oberonParserOF, i)
}

func (s *FormalTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterFormalType(s)
	}
}

func (s *FormalTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitFormalType(s)
	}
}

func (s *FormalTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitFormalType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) FormalType() (localctx IFormalTypeContext) {
	localctx = NewFormalTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, oberonParserRULE_formalType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserARRAY {
		{
			p.SetState(559)
			p.Match(oberonParserARRAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)
			p.Match(oberonParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(565)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(566)
		p.Qualident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODULE() antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	DeclarationSequence() IDeclarationSequenceContext
	END() antlr.TerminalNode
	EOF() antlr.TerminalNode
	ImportList() IImportListContext
	BEGIN() antlr.TerminalNode
	StatementSequence() IStatementSequenceContext

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_module
	return p
}

func InitEmptyModuleContext(p *ModuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_module
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) MODULE() antlr.TerminalNode {
	return s.GetToken(oberonParserMODULE, 0)
}

func (s *ModuleContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *ModuleContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ModuleContext) DeclarationSequence() IDeclarationSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationSequenceContext)
}

func (s *ModuleContext) END() antlr.TerminalNode {
	return s.GetToken(oberonParserEND, 0)
}

func (s *ModuleContext) EOF() antlr.TerminalNode {
	return s.GetToken(oberonParserEOF, 0)
}

func (s *ModuleContext) ImportList() IImportListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportListContext)
}

func (s *ModuleContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(oberonParserBEGIN, 0)
}

func (s *ModuleContext) StatementSequence() IStatementSequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementSequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementSequenceContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitModule(s)
	}
}

func (s *ModuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitModule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, oberonParserRULE_module)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(568)
		p.Match(oberonParserMODULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(569)
		p.Ident()
	}
	{
		p.SetState(570)
		p.Match(oberonParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserIMPORT {
		{
			p.SetState(571)
			p.ImportList()
		}

	}
	{
		p.SetState(574)
		p.DeclarationSequence()
	}
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserBEGIN {
		{
			p.SetState(575)
			p.Match(oberonParserBEGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(576)
			p.StatementSequence()
		}

	}
	{
		p.SetState(579)
		p.Match(oberonParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(580)
		p.Ident()
	}
	{
		p.SetState(581)
		p.Match(oberonParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(582)
		p.Match(oberonParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportListContext is an interface to support dynamic dispatch.
type IImportListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	AllImport_() []IImport_Context
	Import_(i int) IImport_Context

	// IsImportListContext differentiates from other interfaces.
	IsImportListContext()
}

type ImportListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportListContext() *ImportListContext {
	var p = new(ImportListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_importList
	return p
}

func InitEmptyImportListContext(p *ImportListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_importList
}

func (*ImportListContext) IsImportListContext() {}

func NewImportListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportListContext {
	var p = new(ImportListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_importList

	return p
}

func (s *ImportListContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportListContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(oberonParserIMPORT, 0)
}

func (s *ImportListContext) AllImport_() []IImport_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImport_Context); ok {
			len++
		}
	}

	tst := make([]IImport_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImport_Context); ok {
			tst[i] = t.(IImport_Context)
			i++
		}
	}

	return tst
}

func (s *ImportListContext) Import_(i int) IImport_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImport_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImport_Context)
}

func (s *ImportListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterImportList(s)
	}
}

func (s *ImportListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitImportList(s)
	}
}

func (s *ImportListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitImportList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) ImportList() (localctx IImportListContext) {
	localctx = NewImportListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, oberonParserRULE_importList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
		p.Match(oberonParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(585)
		p.Import_()
	}
	p.SetState(590)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__7 {
		{
			p.SetState(586)
			p.Match(oberonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(587)
			p.Import_()
		}

		p.SetState(592)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(593)
		p.Match(oberonParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImport_Context is an interface to support dynamic dispatch.
type IImport_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsImport_Context differentiates from other interfaces.
	IsImport_Context()
}

type Import_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_Context() *Import_Context {
	var p = new(Import_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_import_
	return p
}

func InitEmptyImport_Context(p *Import_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = oberonParserRULE_import_
}

func (*Import_Context) IsImport_Context() {}

func NewImport_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_Context {
	var p = new(Import_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = oberonParserRULE_import_

	return p
}

func (s *Import_Context) GetParser() antlr.Parser { return s.parser }

func (s *Import_Context) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *Import_Context) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *Import_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.EnterImport_(s)
	}
}

func (s *Import_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(oberonListener); ok {
		listenerT.ExitImport_(s)
	}
}

func (s *Import_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case oberonVisitor:
		return t.VisitImport_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *oberonParser) Import_() (localctx IImport_Context) {
	localctx = NewImport_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, oberonParserRULE_import_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(595)
		p.Ident()
	}
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__26 {
		{
			p.SetState(596)
			p.Match(oberonParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(597)
			p.Ident()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
