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
		"", "'.'", "'*'", "'='", "','", "'('", "')'", "';'", "':'", "'#'", "'<'",
		"'<='", "'>'", "'>='", "'+'", "'-'", "'/'", "'&'", "'~'", "'['", "']'",
		"'^'", "'{'", "'}'", "'..'", "':='", "'|'", "'ARRAY'", "'OF'", "'END'",
		"'POINTER'", "'TO'", "'RECORD'", "'PROCEDURE'", "'IN'", "'IS'", "'OR'",
		"'DIV'", "'MOD'", "'NIL'", "'TRUE'", "'FALSE'", "'IF'", "'THEN'", "'ELSIF'",
		"'ELSE'", "'CASE'", "'WHILE'", "'DO'", "'REPEAT'", "'UNTIL'", "'FOR'",
		"'BY'", "'BEGIN'", "'RETURN'", "'CONST'", "'TYPE'", "'VAR'", "'MODULE'",
		"'IMPORT'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "ARRAY", "OF", "END", "POINTER",
		"TO", "RECORD", "PROCEDURE", "IN", "IS", "OR", "DIV", "MOD", "NIL",
		"TRUE", "FALSE", "IF", "THEN", "ELSIF", "ELSE", "CASE", "WHILE", "DO",
		"REPEAT", "UNTIL", "FOR", "BY", "BEGIN", "RETURN", "CONST", "TYPE",
		"VAR", "MODULE", "IMPORT", "STRING", "INTEGER", "HEXNUMBER", "REAL",
		"IDENT", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"ident", "qualident", "identdef", "number", "constDeclaration", "constExpression",
		"typeDeclaration", "type_", "arrayType", "length", "recordType", "baseType",
		"fieldListSequence", "fieldList", "identList", "pointerType", "procedureType",
		"variableDeclaration", "expression", "relation", "simpleExpression",
		"addOperator", "term", "mulOperator", "factor", "designator", "selector",
		"set_", "element", "expList", "actualParameters", "statement", "assignment",
		"procedureCall", "statementSequence", "ifStatement", "caseStatement",
		"case_", "caseLabelList", "labelRange", "label", "whileStatement", "repeatStatement",
		"forStatement", "procedureDeclaration", "procedureHeading", "procedureBody",
		"declarationSequence", "formalParameters", "fPSection", "formalType",
		"module", "importList", "import_",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 550, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 114, 8, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 3, 2, 120, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		139, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 145, 8, 8, 10, 8, 12, 8, 148,
		9, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 160, 8, 10, 1, 10, 3, 10, 163, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 5, 12, 172, 8, 12, 10, 12, 12, 12, 175, 9, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14, 184, 8, 14, 10, 14,
		12, 14, 187, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 195,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 205,
		8, 18, 1, 19, 1, 19, 1, 20, 3, 20, 210, 8, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 5, 20, 216, 8, 20, 10, 20, 12, 20, 219, 9, 20, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 5, 22, 227, 8, 22, 10, 22, 12, 22, 230, 9, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24,
		242, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 250, 8, 24,
		1, 25, 1, 25, 5, 25, 254, 8, 25, 10, 25, 12, 25, 257, 9, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 266, 8, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 5, 27, 272, 8, 27, 10, 27, 12, 27, 275, 9, 27, 3, 27, 277,
		8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 3, 28, 284, 8, 28, 1, 29, 1,
		29, 1, 29, 5, 29, 289, 8, 29, 10, 29, 12, 29, 292, 9, 29, 1, 30, 1, 30,
		3, 30, 296, 8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 3, 31, 307, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		3, 33, 315, 8, 33, 1, 34, 1, 34, 1, 34, 5, 34, 320, 8, 34, 10, 34, 12,
		34, 323, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 5, 35, 334, 8, 35, 10, 35, 12, 35, 337, 9, 35, 1, 35, 1, 35, 3,
		35, 341, 8, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		5, 36, 351, 8, 36, 10, 36, 12, 36, 354, 9, 36, 1, 36, 1, 36, 1, 37, 1,
		37, 1, 37, 1, 37, 3, 37, 362, 8, 37, 1, 38, 1, 38, 1, 38, 5, 38, 367, 8,
		38, 10, 38, 12, 38, 370, 9, 38, 1, 39, 1, 39, 1, 39, 3, 39, 375, 8, 39,
		1, 40, 1, 40, 1, 40, 3, 40, 380, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 391, 8, 41, 10, 41, 12, 41, 394,
		9, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 411, 8, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 3,
		45, 425, 8, 45, 1, 46, 1, 46, 1, 46, 3, 46, 430, 8, 46, 1, 46, 1, 46, 3,
		46, 434, 8, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 442, 8,
		47, 10, 47, 12, 47, 445, 9, 47, 3, 47, 447, 8, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 5, 47, 453, 8, 47, 10, 47, 12, 47, 456, 9, 47, 3, 47, 458, 8, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 464, 8, 47, 10, 47, 12, 47, 467, 9,
		47, 3, 47, 469, 8, 47, 1, 47, 1, 47, 1, 47, 5, 47, 474, 8, 47, 10, 47,
		12, 47, 477, 9, 47, 1, 48, 1, 48, 1, 48, 1, 48, 5, 48, 483, 8, 48, 10,
		48, 12, 48, 486, 9, 48, 3, 48, 488, 8, 48, 1, 48, 1, 48, 1, 48, 3, 48,
		493, 8, 48, 1, 49, 3, 49, 496, 8, 49, 1, 49, 1, 49, 1, 49, 5, 49, 501,
		8, 49, 10, 49, 12, 49, 504, 9, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 5,
		50, 511, 8, 50, 10, 50, 12, 50, 514, 9, 50, 1, 50, 1, 50, 1, 51, 1, 51,
		1, 51, 1, 51, 3, 51, 522, 8, 51, 1, 51, 1, 51, 1, 51, 3, 51, 527, 8, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 538,
		8, 52, 10, 52, 12, 52, 541, 9, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 3,
		53, 548, 8, 53, 1, 53, 0, 0, 54, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92,
		94, 96, 98, 100, 102, 104, 106, 0, 5, 1, 0, 61, 63, 3, 0, 3, 3, 9, 13,
		34, 35, 1, 0, 14, 15, 2, 0, 14, 15, 36, 36, 3, 0, 2, 2, 16, 17, 37, 38,
		567, 0, 108, 1, 0, 0, 0, 2, 113, 1, 0, 0, 0, 4, 117, 1, 0, 0, 0, 6, 121,
		1, 0, 0, 0, 8, 123, 1, 0, 0, 0, 10, 127, 1, 0, 0, 0, 12, 129, 1, 0, 0,
		0, 14, 138, 1, 0, 0, 0, 16, 140, 1, 0, 0, 0, 18, 152, 1, 0, 0, 0, 20, 154,
		1, 0, 0, 0, 22, 166, 1, 0, 0, 0, 24, 168, 1, 0, 0, 0, 26, 176, 1, 0, 0,
		0, 28, 180, 1, 0, 0, 0, 30, 188, 1, 0, 0, 0, 32, 192, 1, 0, 0, 0, 34, 196,
		1, 0, 0, 0, 36, 200, 1, 0, 0, 0, 38, 206, 1, 0, 0, 0, 40, 209, 1, 0, 0,
		0, 42, 220, 1, 0, 0, 0, 44, 222, 1, 0, 0, 0, 46, 231, 1, 0, 0, 0, 48, 249,
		1, 0, 0, 0, 50, 251, 1, 0, 0, 0, 52, 265, 1, 0, 0, 0, 54, 267, 1, 0, 0,
		0, 56, 280, 1, 0, 0, 0, 58, 285, 1, 0, 0, 0, 60, 293, 1, 0, 0, 0, 62, 306,
		1, 0, 0, 0, 64, 308, 1, 0, 0, 0, 66, 312, 1, 0, 0, 0, 68, 316, 1, 0, 0,
		0, 70, 324, 1, 0, 0, 0, 72, 344, 1, 0, 0, 0, 74, 361, 1, 0, 0, 0, 76, 363,
		1, 0, 0, 0, 78, 371, 1, 0, 0, 0, 80, 379, 1, 0, 0, 0, 82, 381, 1, 0, 0,
		0, 84, 397, 1, 0, 0, 0, 86, 402, 1, 0, 0, 0, 88, 416, 1, 0, 0, 0, 90, 421,
		1, 0, 0, 0, 92, 426, 1, 0, 0, 0, 94, 446, 1, 0, 0, 0, 96, 478, 1, 0, 0,
		0, 98, 495, 1, 0, 0, 0, 100, 512, 1, 0, 0, 0, 102, 517, 1, 0, 0, 0, 104,
		533, 1, 0, 0, 0, 106, 544, 1, 0, 0, 0, 108, 109, 5, 64, 0, 0, 109, 1, 1,
		0, 0, 0, 110, 111, 3, 0, 0, 0, 111, 112, 5, 1, 0, 0, 112, 114, 1, 0, 0,
		0, 113, 110, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		116, 3, 0, 0, 0, 116, 3, 1, 0, 0, 0, 117, 119, 3, 0, 0, 0, 118, 120, 5,
		2, 0, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 5, 1, 0, 0, 0,
		121, 122, 7, 0, 0, 0, 122, 7, 1, 0, 0, 0, 123, 124, 3, 4, 2, 0, 124, 125,
		5, 3, 0, 0, 125, 126, 3, 10, 5, 0, 126, 9, 1, 0, 0, 0, 127, 128, 3, 36,
		18, 0, 128, 11, 1, 0, 0, 0, 129, 130, 3, 4, 2, 0, 130, 131, 5, 3, 0, 0,
		131, 132, 3, 14, 7, 0, 132, 13, 1, 0, 0, 0, 133, 139, 3, 2, 1, 0, 134,
		139, 3, 16, 8, 0, 135, 139, 3, 20, 10, 0, 136, 139, 3, 30, 15, 0, 137,
		139, 3, 32, 16, 0, 138, 133, 1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 138, 135,
		1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 137, 1, 0, 0, 0, 139, 15, 1, 0,
		0, 0, 140, 141, 5, 27, 0, 0, 141, 146, 3, 18, 9, 0, 142, 143, 5, 4, 0,
		0, 143, 145, 3, 18, 9, 0, 144, 142, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 146,
		1, 0, 0, 0, 149, 150, 5, 28, 0, 0, 150, 151, 3, 14, 7, 0, 151, 17, 1, 0,
		0, 0, 152, 153, 3, 10, 5, 0, 153, 19, 1, 0, 0, 0, 154, 159, 5, 32, 0, 0,
		155, 156, 5, 5, 0, 0, 156, 157, 3, 22, 11, 0, 157, 158, 5, 6, 0, 0, 158,
		160, 1, 0, 0, 0, 159, 155, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 162,
		1, 0, 0, 0, 161, 163, 3, 24, 12, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1,
		0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 5, 29, 0, 0, 165, 21, 1, 0, 0,
		0, 166, 167, 3, 2, 1, 0, 167, 23, 1, 0, 0, 0, 168, 173, 3, 26, 13, 0, 169,
		170, 5, 7, 0, 0, 170, 172, 3, 26, 13, 0, 171, 169, 1, 0, 0, 0, 172, 175,
		1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 25, 1, 0,
		0, 0, 175, 173, 1, 0, 0, 0, 176, 177, 3, 28, 14, 0, 177, 178, 5, 8, 0,
		0, 178, 179, 3, 14, 7, 0, 179, 27, 1, 0, 0, 0, 180, 185, 3, 4, 2, 0, 181,
		182, 5, 4, 0, 0, 182, 184, 3, 4, 2, 0, 183, 181, 1, 0, 0, 0, 184, 187,
		1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 29, 1, 0,
		0, 0, 187, 185, 1, 0, 0, 0, 188, 189, 5, 30, 0, 0, 189, 190, 5, 31, 0,
		0, 190, 191, 3, 14, 7, 0, 191, 31, 1, 0, 0, 0, 192, 194, 5, 33, 0, 0, 193,
		195, 3, 96, 48, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 33,
		1, 0, 0, 0, 196, 197, 3, 28, 14, 0, 197, 198, 5, 8, 0, 0, 198, 199, 3,
		14, 7, 0, 199, 35, 1, 0, 0, 0, 200, 204, 3, 40, 20, 0, 201, 202, 3, 38,
		19, 0, 202, 203, 3, 40, 20, 0, 203, 205, 1, 0, 0, 0, 204, 201, 1, 0, 0,
		0, 204, 205, 1, 0, 0, 0, 205, 37, 1, 0, 0, 0, 206, 207, 7, 1, 0, 0, 207,
		39, 1, 0, 0, 0, 208, 210, 7, 2, 0, 0, 209, 208, 1, 0, 0, 0, 209, 210, 1,
		0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 217, 3, 44, 22, 0, 212, 213, 3, 42,
		21, 0, 213, 214, 3, 44, 22, 0, 214, 216, 1, 0, 0, 0, 215, 212, 1, 0, 0,
		0, 216, 219, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218,
		41, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220, 221, 7, 3, 0, 0, 221, 43, 1,
		0, 0, 0, 222, 228, 3, 48, 24, 0, 223, 224, 3, 46, 23, 0, 224, 225, 3, 48,
		24, 0, 225, 227, 1, 0, 0, 0, 226, 223, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0,
		228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 45, 1, 0, 0, 0, 230, 228,
		1, 0, 0, 0, 231, 232, 7, 4, 0, 0, 232, 47, 1, 0, 0, 0, 233, 250, 3, 6,
		3, 0, 234, 250, 5, 60, 0, 0, 235, 250, 5, 39, 0, 0, 236, 250, 5, 40, 0,
		0, 237, 250, 5, 41, 0, 0, 238, 250, 3, 54, 27, 0, 239, 241, 3, 50, 25,
		0, 240, 242, 3, 60, 30, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0,
		242, 250, 1, 0, 0, 0, 243, 244, 5, 5, 0, 0, 244, 245, 3, 36, 18, 0, 245,
		246, 5, 6, 0, 0, 246, 250, 1, 0, 0, 0, 247, 248, 5, 18, 0, 0, 248, 250,
		3, 48, 24, 0, 249, 233, 1, 0, 0, 0, 249, 234, 1, 0, 0, 0, 249, 235, 1,
		0, 0, 0, 249, 236, 1, 0, 0, 0, 249, 237, 1, 0, 0, 0, 249, 238, 1, 0, 0,
		0, 249, 239, 1, 0, 0, 0, 249, 243, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250,
		49, 1, 0, 0, 0, 251, 255, 3, 2, 1, 0, 252, 254, 3, 52, 26, 0, 253, 252,
		1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0,
		0, 0, 256, 51, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 259, 5, 1, 0, 0,
		259, 266, 3, 0, 0, 0, 260, 261, 5, 19, 0, 0, 261, 262, 3, 58, 29, 0, 262,
		263, 5, 20, 0, 0, 263, 266, 1, 0, 0, 0, 264, 266, 5, 21, 0, 0, 265, 258,
		1, 0, 0, 0, 265, 260, 1, 0, 0, 0, 265, 264, 1, 0, 0, 0, 266, 53, 1, 0,
		0, 0, 267, 276, 5, 22, 0, 0, 268, 273, 3, 56, 28, 0, 269, 270, 5, 4, 0,
		0, 270, 272, 3, 56, 28, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0,
		273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275,
		273, 1, 0, 0, 0, 276, 268, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278,
		1, 0, 0, 0, 278, 279, 5, 23, 0, 0, 279, 55, 1, 0, 0, 0, 280, 283, 3, 36,
		18, 0, 281, 282, 5, 24, 0, 0, 282, 284, 3, 36, 18, 0, 283, 281, 1, 0, 0,
		0, 283, 284, 1, 0, 0, 0, 284, 57, 1, 0, 0, 0, 285, 290, 3, 36, 18, 0, 286,
		287, 5, 4, 0, 0, 287, 289, 3, 36, 18, 0, 288, 286, 1, 0, 0, 0, 289, 292,
		1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 59, 1, 0,
		0, 0, 292, 290, 1, 0, 0, 0, 293, 295, 5, 5, 0, 0, 294, 296, 3, 58, 29,
		0, 295, 294, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297,
		298, 5, 6, 0, 0, 298, 61, 1, 0, 0, 0, 299, 307, 3, 64, 32, 0, 300, 307,
		3, 66, 33, 0, 301, 307, 3, 70, 35, 0, 302, 307, 3, 72, 36, 0, 303, 307,
		3, 82, 41, 0, 304, 307, 3, 84, 42, 0, 305, 307, 3, 86, 43, 0, 306, 299,
		1, 0, 0, 0, 306, 300, 1, 0, 0, 0, 306, 301, 1, 0, 0, 0, 306, 302, 1, 0,
		0, 0, 306, 303, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 305, 1, 0, 0, 0,
		306, 307, 1, 0, 0, 0, 307, 63, 1, 0, 0, 0, 308, 309, 3, 50, 25, 0, 309,
		310, 5, 25, 0, 0, 310, 311, 3, 36, 18, 0, 311, 65, 1, 0, 0, 0, 312, 314,
		3, 50, 25, 0, 313, 315, 3, 60, 30, 0, 314, 313, 1, 0, 0, 0, 314, 315, 1,
		0, 0, 0, 315, 67, 1, 0, 0, 0, 316, 321, 3, 62, 31, 0, 317, 318, 5, 7, 0,
		0, 318, 320, 3, 62, 31, 0, 319, 317, 1, 0, 0, 0, 320, 323, 1, 0, 0, 0,
		321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 69, 1, 0, 0, 0, 323, 321,
		1, 0, 0, 0, 324, 325, 5, 42, 0, 0, 325, 326, 3, 36, 18, 0, 326, 327, 5,
		43, 0, 0, 327, 335, 3, 68, 34, 0, 328, 329, 5, 44, 0, 0, 329, 330, 3, 36,
		18, 0, 330, 331, 5, 43, 0, 0, 331, 332, 3, 68, 34, 0, 332, 334, 1, 0, 0,
		0, 333, 328, 1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 335,
		336, 1, 0, 0, 0, 336, 340, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 338, 339,
		5, 45, 0, 0, 339, 341, 3, 68, 34, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1,
		0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 343, 5, 29, 0, 0, 343, 71, 1, 0, 0,
		0, 344, 345, 5, 46, 0, 0, 345, 346, 3, 36, 18, 0, 346, 347, 5, 28, 0, 0,
		347, 352, 3, 74, 37, 0, 348, 349, 5, 26, 0, 0, 349, 351, 3, 74, 37, 0,
		350, 348, 1, 0, 0, 0, 351, 354, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 352,
		353, 1, 0, 0, 0, 353, 355, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 355, 356,
		5, 29, 0, 0, 356, 73, 1, 0, 0, 0, 357, 358, 3, 76, 38, 0, 358, 359, 5,
		8, 0, 0, 359, 360, 3, 68, 34, 0, 360, 362, 1, 0, 0, 0, 361, 357, 1, 0,
		0, 0, 361, 362, 1, 0, 0, 0, 362, 75, 1, 0, 0, 0, 363, 368, 3, 78, 39, 0,
		364, 365, 5, 4, 0, 0, 365, 367, 3, 78, 39, 0, 366, 364, 1, 0, 0, 0, 367,
		370, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 77, 1,
		0, 0, 0, 370, 368, 1, 0, 0, 0, 371, 374, 3, 80, 40, 0, 372, 373, 5, 24,
		0, 0, 373, 375, 3, 80, 40, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0,
		0, 375, 79, 1, 0, 0, 0, 376, 380, 5, 61, 0, 0, 377, 380, 5, 60, 0, 0, 378,
		380, 3, 2, 1, 0, 379, 376, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 378,
		1, 0, 0, 0, 380, 81, 1, 0, 0, 0, 381, 382, 5, 47, 0, 0, 382, 383, 3, 36,
		18, 0, 383, 384, 5, 48, 0, 0, 384, 392, 3, 68, 34, 0, 385, 386, 5, 44,
		0, 0, 386, 387, 3, 36, 18, 0, 387, 388, 5, 48, 0, 0, 388, 389, 3, 68, 34,
		0, 389, 391, 1, 0, 0, 0, 390, 385, 1, 0, 0, 0, 391, 394, 1, 0, 0, 0, 392,
		390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 395, 1, 0, 0, 0, 394, 392,
		1, 0, 0, 0, 395, 396, 5, 29, 0, 0, 396, 83, 1, 0, 0, 0, 397, 398, 5, 49,
		0, 0, 398, 399, 3, 68, 34, 0, 399, 400, 5, 50, 0, 0, 400, 401, 3, 36, 18,
		0, 401, 85, 1, 0, 0, 0, 402, 403, 5, 51, 0, 0, 403, 404, 3, 0, 0, 0, 404,
		405, 5, 25, 0, 0, 405, 406, 3, 36, 18, 0, 406, 407, 5, 31, 0, 0, 407, 410,
		3, 36, 18, 0, 408, 409, 5, 52, 0, 0, 409, 411, 3, 10, 5, 0, 410, 408, 1,
		0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 413, 5, 48, 0,
		0, 413, 414, 3, 68, 34, 0, 414, 415, 5, 29, 0, 0, 415, 87, 1, 0, 0, 0,
		416, 417, 3, 90, 45, 0, 417, 418, 5, 7, 0, 0, 418, 419, 3, 92, 46, 0, 419,
		420, 3, 0, 0, 0, 420, 89, 1, 0, 0, 0, 421, 422, 5, 33, 0, 0, 422, 424,
		3, 4, 2, 0, 423, 425, 3, 96, 48, 0, 424, 423, 1, 0, 0, 0, 424, 425, 1,
		0, 0, 0, 425, 91, 1, 0, 0, 0, 426, 429, 3, 94, 47, 0, 427, 428, 5, 53,
		0, 0, 428, 430, 3, 68, 34, 0, 429, 427, 1, 0, 0, 0, 429, 430, 1, 0, 0,
		0, 430, 433, 1, 0, 0, 0, 431, 432, 5, 54, 0, 0, 432, 434, 3, 36, 18, 0,
		433, 431, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435,
		436, 5, 29, 0, 0, 436, 93, 1, 0, 0, 0, 437, 443, 5, 55, 0, 0, 438, 439,
		3, 8, 4, 0, 439, 440, 5, 7, 0, 0, 440, 442, 1, 0, 0, 0, 441, 438, 1, 0,
		0, 0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0,
		444, 447, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 437, 1, 0, 0, 0, 446,
		447, 1, 0, 0, 0, 447, 457, 1, 0, 0, 0, 448, 454, 5, 56, 0, 0, 449, 450,
		3, 12, 6, 0, 450, 451, 5, 7, 0, 0, 451, 453, 1, 0, 0, 0, 452, 449, 1, 0,
		0, 0, 453, 456, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0,
		455, 458, 1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 457, 448, 1, 0, 0, 0, 457,
		458, 1, 0, 0, 0, 458, 468, 1, 0, 0, 0, 459, 465, 5, 57, 0, 0, 460, 461,
		3, 34, 17, 0, 461, 462, 5, 7, 0, 0, 462, 464, 1, 0, 0, 0, 463, 460, 1,
		0, 0, 0, 464, 467, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0, 465, 466, 1, 0, 0,
		0, 466, 469, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 468, 459, 1, 0, 0, 0, 468,
		469, 1, 0, 0, 0, 469, 475, 1, 0, 0, 0, 470, 471, 3, 88, 44, 0, 471, 472,
		5, 7, 0, 0, 472, 474, 1, 0, 0, 0, 473, 470, 1, 0, 0, 0, 474, 477, 1, 0,
		0, 0, 475, 473, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 95, 1, 0, 0, 0,
		477, 475, 1, 0, 0, 0, 478, 487, 5, 5, 0, 0, 479, 484, 3, 98, 49, 0, 480,
		481, 5, 7, 0, 0, 481, 483, 3, 98, 49, 0, 482, 480, 1, 0, 0, 0, 483, 486,
		1, 0, 0, 0, 484, 482, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 488, 1, 0,
		0, 0, 486, 484, 1, 0, 0, 0, 487, 479, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0,
		488, 489, 1, 0, 0, 0, 489, 492, 5, 6, 0, 0, 490, 491, 5, 8, 0, 0, 491,
		493, 3, 2, 1, 0, 492, 490, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 97, 1,
		0, 0, 0, 494, 496, 5, 57, 0, 0, 495, 494, 1, 0, 0, 0, 495, 496, 1, 0, 0,
		0, 496, 497, 1, 0, 0, 0, 497, 502, 3, 0, 0, 0, 498, 499, 5, 4, 0, 0, 499,
		501, 3, 0, 0, 0, 500, 498, 1, 0, 0, 0, 501, 504, 1, 0, 0, 0, 502, 500,
		1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 505, 1, 0, 0, 0, 504, 502, 1, 0,
		0, 0, 505, 506, 5, 8, 0, 0, 506, 507, 3, 100, 50, 0, 507, 99, 1, 0, 0,
		0, 508, 509, 5, 27, 0, 0, 509, 511, 5, 28, 0, 0, 510, 508, 1, 0, 0, 0,
		511, 514, 1, 0, 0, 0, 512, 510, 1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513,
		515, 1, 0, 0, 0, 514, 512, 1, 0, 0, 0, 515, 516, 3, 2, 1, 0, 516, 101,
		1, 0, 0, 0, 517, 518, 5, 58, 0, 0, 518, 519, 3, 0, 0, 0, 519, 521, 5, 7,
		0, 0, 520, 522, 3, 104, 52, 0, 521, 520, 1, 0, 0, 0, 521, 522, 1, 0, 0,
		0, 522, 523, 1, 0, 0, 0, 523, 526, 3, 94, 47, 0, 524, 525, 5, 53, 0, 0,
		525, 527, 3, 68, 34, 0, 526, 524, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527,
		528, 1, 0, 0, 0, 528, 529, 5, 29, 0, 0, 529, 530, 3, 0, 0, 0, 530, 531,
		5, 1, 0, 0, 531, 532, 5, 0, 0, 1, 532, 103, 1, 0, 0, 0, 533, 534, 5, 59,
		0, 0, 534, 539, 3, 106, 53, 0, 535, 536, 5, 4, 0, 0, 536, 538, 3, 106,
		53, 0, 537, 535, 1, 0, 0, 0, 538, 541, 1, 0, 0, 0, 539, 537, 1, 0, 0, 0,
		539, 540, 1, 0, 0, 0, 540, 542, 1, 0, 0, 0, 541, 539, 1, 0, 0, 0, 542,
		543, 5, 7, 0, 0, 543, 105, 1, 0, 0, 0, 544, 547, 3, 0, 0, 0, 545, 546,
		5, 25, 0, 0, 546, 548, 3, 0, 0, 0, 547, 545, 1, 0, 0, 0, 547, 548, 1, 0,
		0, 0, 548, 107, 1, 0, 0, 0, 54, 113, 119, 138, 146, 159, 162, 173, 185,
		194, 204, 209, 217, 228, 241, 249, 255, 265, 273, 276, 283, 290, 295, 306,
		314, 321, 335, 340, 352, 361, 368, 374, 379, 392, 410, 424, 429, 433, 443,
		446, 454, 457, 465, 468, 475, 484, 487, 492, 495, 502, 512, 521, 526, 539,
		547,
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
	oberonParserARRAY     = 27
	oberonParserOF        = 28
	oberonParserEND       = 29
	oberonParserPOINTER   = 30
	oberonParserTO        = 31
	oberonParserRECORD    = 32
	oberonParserPROCEDURE = 33
	oberonParserIN        = 34
	oberonParserIS        = 35
	oberonParserOR        = 36
	oberonParserDIV       = 37
	oberonParserMOD       = 38
	oberonParserNIL       = 39
	oberonParserTRUE      = 40
	oberonParserFALSE     = 41
	oberonParserIF        = 42
	oberonParserTHEN      = 43
	oberonParserELSIF     = 44
	oberonParserELSE      = 45
	oberonParserCASE      = 46
	oberonParserWHILE     = 47
	oberonParserDO        = 48
	oberonParserREPEAT    = 49
	oberonParserUNTIL     = 50
	oberonParserFOR       = 51
	oberonParserBY        = 52
	oberonParserBEGIN     = 53
	oberonParserRETURN    = 54
	oberonParserCONST     = 55
	oberonParserTYPE      = 56
	oberonParserVAR       = 57
	oberonParserMODULE    = 58
	oberonParserIMPORT    = 59
	oberonParserSTRING    = 60
	oberonParserINTEGER   = 61
	oberonParserHEXNUMBER = 62
	oberonParserREAL      = 63
	oberonParserIDENT     = 64
	oberonParserCOMMENT   = 65
	oberonParserWS        = 66
)

// oberonParser rules.
const (
	oberonParserRULE_ident                = 0
	oberonParserRULE_qualident            = 1
	oberonParserRULE_identdef             = 2
	oberonParserRULE_number               = 3
	oberonParserRULE_constDeclaration     = 4
	oberonParserRULE_constExpression      = 5
	oberonParserRULE_typeDeclaration      = 6
	oberonParserRULE_type_                = 7
	oberonParserRULE_arrayType            = 8
	oberonParserRULE_length               = 9
	oberonParserRULE_recordType           = 10
	oberonParserRULE_baseType             = 11
	oberonParserRULE_fieldListSequence    = 12
	oberonParserRULE_fieldList            = 13
	oberonParserRULE_identList            = 14
	oberonParserRULE_pointerType          = 15
	oberonParserRULE_procedureType        = 16
	oberonParserRULE_variableDeclaration  = 17
	oberonParserRULE_expression           = 18
	oberonParserRULE_relation             = 19
	oberonParserRULE_simpleExpression     = 20
	oberonParserRULE_addOperator          = 21
	oberonParserRULE_term                 = 22
	oberonParserRULE_mulOperator          = 23
	oberonParserRULE_factor               = 24
	oberonParserRULE_designator           = 25
	oberonParserRULE_selector             = 26
	oberonParserRULE_set_                 = 27
	oberonParserRULE_element              = 28
	oberonParserRULE_expList              = 29
	oberonParserRULE_actualParameters     = 30
	oberonParserRULE_statement            = 31
	oberonParserRULE_assignment           = 32
	oberonParserRULE_procedureCall        = 33
	oberonParserRULE_statementSequence    = 34
	oberonParserRULE_ifStatement          = 35
	oberonParserRULE_caseStatement        = 36
	oberonParserRULE_case_                = 37
	oberonParserRULE_caseLabelList        = 38
	oberonParserRULE_labelRange           = 39
	oberonParserRULE_label                = 40
	oberonParserRULE_whileStatement       = 41
	oberonParserRULE_repeatStatement      = 42
	oberonParserRULE_forStatement         = 43
	oberonParserRULE_procedureDeclaration = 44
	oberonParserRULE_procedureHeading     = 45
	oberonParserRULE_procedureBody        = 46
	oberonParserRULE_declarationSequence  = 47
	oberonParserRULE_formalParameters     = 48
	oberonParserRULE_fPSection            = 49
	oberonParserRULE_formalType           = 50
	oberonParserRULE_module               = 51
	oberonParserRULE_importList           = 52
	oberonParserRULE_import_              = 53
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
		p.SetState(108)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)
			p.Ident()
		}
		{
			p.SetState(111)
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
		p.SetState(115)
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
		p.SetState(117)
		p.Ident()
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__1 {
		{
			p.SetState(118)
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	REAL() antlr.TerminalNode
	HEXNUMBER() antlr.TerminalNode

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

func (s *NumberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(oberonParserINTEGER, 0)
}

func (s *NumberContext) REAL() antlr.TerminalNode {
	return s.GetToken(oberonParserREAL, 0)
}

func (s *NumberContext) HEXNUMBER() antlr.TerminalNode {
	return s.GetToken(oberonParserHEXNUMBER, 0)
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
	p.EnterRule(localctx, 6, oberonParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2305843009213693952) != 0) {
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
	p.EnterRule(localctx, 8, oberonParserRULE_constDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Identdef()
	}
	{
		p.SetState(124)
		p.Match(oberonParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
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
	p.EnterRule(localctx, 10, oberonParserRULE_constExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
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
	p.EnterRule(localctx, 12, oberonParserRULE_typeDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Identdef()
	}
	{
		p.SetState(130)
		p.Match(oberonParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
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
	p.EnterRule(localctx, 14, oberonParserRULE_type_)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case oberonParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.Qualident()
		}

	case oberonParserARRAY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
			p.ArrayType()
		}

	case oberonParserRECORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(135)
			p.RecordType()
		}

	case oberonParserPOINTER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.PointerType()
		}

	case oberonParserPROCEDURE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(137)
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
	p.EnterRule(localctx, 16, oberonParserRULE_arrayType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(oberonParserARRAY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Length()
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__3 {
		{
			p.SetState(142)
			p.Match(oberonParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Length()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Match(oberonParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
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
	p.EnterRule(localctx, 18, oberonParserRULE_length)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
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
	p.EnterRule(localctx, 20, oberonParserRULE_recordType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(oberonParserRECORD)
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

	if _la == oberonParserT__4 {
		{
			p.SetState(155)
			p.Match(oberonParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.BaseType()
		}
		{
			p.SetState(157)
			p.Match(oberonParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserIDENT {
		{
			p.SetState(161)
			p.FieldListSequence()
		}

	}
	{
		p.SetState(164)
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
	p.EnterRule(localctx, 22, oberonParserRULE_baseType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
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
	p.EnterRule(localctx, 24, oberonParserRULE_fieldListSequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.FieldList()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__6 {
		{
			p.SetState(169)
			p.Match(oberonParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.FieldList()
		}

		p.SetState(175)
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
	p.EnterRule(localctx, 26, oberonParserRULE_fieldList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.IdentList()
	}
	{
		p.SetState(177)
		p.Match(oberonParserT__7)
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
	p.EnterRule(localctx, 28, oberonParserRULE_identList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Identdef()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__3 {
		{
			p.SetState(181)
			p.Match(oberonParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Identdef()
		}

		p.SetState(187)
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
	p.EnterRule(localctx, 30, oberonParserRULE_pointerType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(oberonParserPOINTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Match(oberonParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
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
	p.EnterRule(localctx, 32, oberonParserRULE_procedureType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(oberonParserPROCEDURE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__4 {
		{
			p.SetState(193)
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
	p.EnterRule(localctx, 34, oberonParserRULE_variableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.IdentList()
	}
	{
		p.SetState(197)
		p.Match(oberonParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
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
	p.EnterRule(localctx, 36, oberonParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.SimpleExpression()
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&51539623432) != 0 {
		{
			p.SetState(201)
			p.Relation()
		}
		{
			p.SetState(202)
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
	p.EnterRule(localctx, 38, oberonParserRULE_relation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&51539623432) != 0) {
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
	p.EnterRule(localctx, 40, oberonParserRULE_simpleExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__13 || _la == oberonParserT__14 {
		{
			p.SetState(208)
			_la = p.GetTokenStream().LA(1)

			if !(_la == oberonParserT__13 || _la == oberonParserT__14) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(211)
		p.Term()
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719525888) != 0 {
		{
			p.SetState(212)
			p.AddOperator()
		}
		{
			p.SetState(213)
			p.Term()
		}

		p.SetState(219)
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
	p.EnterRule(localctx, 42, oberonParserRULE_addOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68719525888) != 0) {
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
	p.EnterRule(localctx, 44, oberonParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Factor()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&412317057028) != 0 {
		{
			p.SetState(223)
			p.MulOperator()
		}
		{
			p.SetState(224)
			p.Factor()
		}

		p.SetState(230)
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
	p.EnterRule(localctx, 46, oberonParserRULE_mulOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&412317057028) != 0) {
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
	p.EnterRule(localctx, 48, oberonParserRULE_factor)
	var _la int

	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case oberonParserINTEGER, oberonParserHEXNUMBER, oberonParserREAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Number()
		}

	case oberonParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(oberonParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserNIL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(235)
			p.Match(oberonParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(236)
			p.Match(oberonParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserFALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(237)
			p.Match(oberonParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserT__21:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(238)
			p.Set_()
		}

	case oberonParserIDENT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(239)
			p.Designator()
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == oberonParserT__4 {
			{
				p.SetState(240)
				p.ActualParameters()
			}

		}

	case oberonParserT__4:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(243)
			p.Match(oberonParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Expression()
		}
		{
			p.SetState(245)
			p.Match(oberonParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserT__17:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(247)
			p.Match(oberonParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
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
	p.EnterRule(localctx, 50, oberonParserRULE_designator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Qualident()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2621442) != 0 {
		{
			p.SetState(252)
			p.Selector()
		}

		p.SetState(257)
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

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	ExpList() IExpListContext

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
	p.EnterRule(localctx, 52, oberonParserRULE_selector)
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case oberonParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Match(oberonParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Ident()
		}

	case oberonParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Match(oberonParserT__18)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.ExpList()
		}
		{
			p.SetState(262)
			p.Match(oberonParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.Match(oberonParserT__20)
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
	p.EnterRule(localctx, 54, oberonParserRULE_set_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(oberonParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&1116892827847108097) != 0 {
		{
			p.SetState(268)
			p.Element()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserT__3 {
			{
				p.SetState(269)
				p.Match(oberonParserT__3)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(270)
				p.Element()
			}

			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(278)
		p.Match(oberonParserT__22)
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
	p.EnterRule(localctx, 56, oberonParserRULE_element)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Expression()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__23 {
		{
			p.SetState(281)
			p.Match(oberonParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
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
	p.EnterRule(localctx, 58, oberonParserRULE_expList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Expression()
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__3 {
		{
			p.SetState(286)
			p.Match(oberonParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.Expression()
		}

		p.SetState(292)
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
	p.EnterRule(localctx, 60, oberonParserRULE_actualParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(oberonParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-5)) & ^0x3f) == 0 && ((int64(1)<<(_la-5))&1116892827847108097) != 0 {
		{
			p.SetState(294)
			p.ExpList()
		}

	}
	{
		p.SetState(297)
		p.Match(oberonParserT__5)
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
	p.EnterRule(localctx, 62, oberonParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(299)
			p.Assignment()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(300)
			p.ProcedureCall()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(301)
			p.IfStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 4 {
		{
			p.SetState(302)
			p.CaseStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 5 {
		{
			p.SetState(303)
			p.WhileStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 6 {
		{
			p.SetState(304)
			p.RepeatStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) == 7 {
		{
			p.SetState(305)
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
	p.EnterRule(localctx, 64, oberonParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Designator()
	}
	{
		p.SetState(309)
		p.Match(oberonParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
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
	p.EnterRule(localctx, 66, oberonParserRULE_procedureCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Designator()
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__4 {
		{
			p.SetState(313)
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
	p.EnterRule(localctx, 68, oberonParserRULE_statementSequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Statement()
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__6 {
		{
			p.SetState(317)
			p.Match(oberonParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Statement()
		}

		p.SetState(323)
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
	p.EnterRule(localctx, 70, oberonParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(oberonParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.Expression()
	}
	{
		p.SetState(326)
		p.Match(oberonParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.StatementSequence()
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserELSIF {
		{
			p.SetState(328)
			p.Match(oberonParserELSIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Expression()
		}
		{
			p.SetState(330)
			p.Match(oberonParserTHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.StatementSequence()
		}

		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserELSE {
		{
			p.SetState(338)
			p.Match(oberonParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.StatementSequence()
		}

	}
	{
		p.SetState(342)
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
	p.EnterRule(localctx, 72, oberonParserRULE_caseStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(oberonParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Expression()
	}
	{
		p.SetState(346)
		p.Match(oberonParserOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Case_()
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__25 {
		{
			p.SetState(348)
			p.Match(oberonParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Case_()
		}

		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(355)
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
	p.EnterRule(localctx, 74, oberonParserRULE_case_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-60)) & ^0x3f) == 0 && ((int64(1)<<(_la-60))&19) != 0 {
		{
			p.SetState(357)
			p.CaseLabelList()
		}
		{
			p.SetState(358)
			p.Match(oberonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)
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
	p.EnterRule(localctx, 76, oberonParserRULE_caseLabelList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.LabelRange()
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__3 {
		{
			p.SetState(364)
			p.Match(oberonParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.LabelRange()
		}

		p.SetState(370)
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
	p.EnterRule(localctx, 78, oberonParserRULE_labelRange)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Label()
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__23 {
		{
			p.SetState(372)
			p.Match(oberonParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
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
	INTEGER() antlr.TerminalNode
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

func (s *LabelContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(oberonParserINTEGER, 0)
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
	p.EnterRule(localctx, 80, oberonParserRULE_label)
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case oberonParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(376)
			p.Match(oberonParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
			p.Match(oberonParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case oberonParserIDENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(378)
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
	p.EnterRule(localctx, 82, oberonParserRULE_whileStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)
		p.Match(oberonParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.Expression()
	}
	{
		p.SetState(383)
		p.Match(oberonParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.StatementSequence()
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserELSIF {
		{
			p.SetState(385)
			p.Match(oberonParserELSIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Expression()
		}
		{
			p.SetState(387)
			p.Match(oberonParserDO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(388)
			p.StatementSequence()
		}

		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(395)
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
	p.EnterRule(localctx, 84, oberonParserRULE_repeatStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(oberonParserREPEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.StatementSequence()
	}
	{
		p.SetState(399)
		p.Match(oberonParserUNTIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
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
	p.EnterRule(localctx, 86, oberonParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Match(oberonParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Ident()
	}
	{
		p.SetState(404)
		p.Match(oberonParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)
		p.Expression()
	}
	{
		p.SetState(406)
		p.Match(oberonParserTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Expression()
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserBY {
		{
			p.SetState(408)
			p.Match(oberonParserBY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.ConstExpression()
		}

	}
	{
		p.SetState(412)
		p.Match(oberonParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.StatementSequence()
	}
	{
		p.SetState(414)
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
	p.EnterRule(localctx, 88, oberonParserRULE_procedureDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(416)
		p.ProcedureHeading()
	}
	{
		p.SetState(417)
		p.Match(oberonParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(418)
		p.ProcedureBody()
	}
	{
		p.SetState(419)
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
	p.EnterRule(localctx, 90, oberonParserRULE_procedureHeading)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(oberonParserPROCEDURE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(422)
		p.Identdef()
	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__4 {
		{
			p.SetState(423)
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
	p.EnterRule(localctx, 92, oberonParserRULE_procedureBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.DeclarationSequence()
	}
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserBEGIN {
		{
			p.SetState(427)
			p.Match(oberonParserBEGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.StatementSequence()
		}

	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserRETURN {
		{
			p.SetState(431)
			p.Match(oberonParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Expression()
		}

	}
	{
		p.SetState(435)
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
	p.EnterRule(localctx, 94, oberonParserRULE_declarationSequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserCONST {
		{
			p.SetState(437)
			p.Match(oberonParserCONST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserIDENT {
			{
				p.SetState(438)
				p.ConstDeclaration()
			}
			{
				p.SetState(439)
				p.Match(oberonParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(445)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserTYPE {
		{
			p.SetState(448)
			p.Match(oberonParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserIDENT {
			{
				p.SetState(449)
				p.TypeDeclaration()
			}
			{
				p.SetState(450)
				p.Match(oberonParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(456)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserVAR {
		{
			p.SetState(459)
			p.Match(oberonParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserIDENT {
			{
				p.SetState(460)
				p.VariableDeclaration()
			}
			{
				p.SetState(461)
				p.Match(oberonParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(467)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserPROCEDURE {
		{
			p.SetState(470)
			p.ProcedureDeclaration()
		}
		{
			p.SetState(471)
			p.Match(oberonParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(477)
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
	p.EnterRule(localctx, 96, oberonParserRULE_formalParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(478)
		p.Match(oberonParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserVAR || _la == oberonParserIDENT {
		{
			p.SetState(479)
			p.FPSection()
		}
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == oberonParserT__6 {
			{
				p.SetState(480)
				p.Match(oberonParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(481)
				p.FPSection()
			}

			p.SetState(486)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(489)
		p.Match(oberonParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__7 {
		{
			p.SetState(490)
			p.Match(oberonParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
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
	p.EnterRule(localctx, 98, oberonParserRULE_fPSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserVAR {
		{
			p.SetState(494)
			p.Match(oberonParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(497)
		p.Ident()
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__3 {
		{
			p.SetState(498)
			p.Match(oberonParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(499)
			p.Ident()
		}

		p.SetState(504)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(505)
		p.Match(oberonParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(506)
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
	p.EnterRule(localctx, 100, oberonParserRULE_formalType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(512)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserARRAY {
		{
			p.SetState(508)
			p.Match(oberonParserARRAY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(509)
			p.Match(oberonParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(515)
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
	p.EnterRule(localctx, 102, oberonParserRULE_module)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.Match(oberonParserMODULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(518)
		p.Ident()
	}
	{
		p.SetState(519)
		p.Match(oberonParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserIMPORT {
		{
			p.SetState(520)
			p.ImportList()
		}

	}
	{
		p.SetState(523)
		p.DeclarationSequence()
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserBEGIN {
		{
			p.SetState(524)
			p.Match(oberonParserBEGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(525)
			p.StatementSequence()
		}

	}
	{
		p.SetState(528)
		p.Match(oberonParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(529)
		p.Ident()
	}
	{
		p.SetState(530)
		p.Match(oberonParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(531)
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
	p.EnterRule(localctx, 104, oberonParserRULE_importList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(533)
		p.Match(oberonParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(534)
		p.Import_()
	}
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == oberonParserT__3 {
		{
			p.SetState(535)
			p.Match(oberonParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(536)
			p.Import_()
		}

		p.SetState(541)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(542)
		p.Match(oberonParserT__6)
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
	p.EnterRule(localctx, 106, oberonParserRULE_import_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(544)
		p.Ident()
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == oberonParserT__24 {
		{
			p.SetState(545)
			p.Match(oberonParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(546)
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
