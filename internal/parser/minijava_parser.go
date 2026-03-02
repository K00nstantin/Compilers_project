// Code generated from MiniJava.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MiniJava
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)



    package mini_java.antlr;


// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type MiniJavaParser struct {
	*antlr.BaseParser
}

var MiniJavaParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func minijavaParserInit() {
  staticData := &MiniJavaParserStaticData
  staticData.LiteralNames = []string{
    "", "'class'", "'extends'", "'{'", "'}'", "'public'", "'static'", "'void'", 
    "'main'", "'('", "'String'", "'['", "']'", "')'", "';'", "','", "'if'", 
    "'else'", "'while'", "'System.out.println'", "'='", "'return'", "'recur'", 
    "'?'", "':'", "'.'", "'length'", "'-'", "'!'", "'new'", "'int'", "'+'", 
    "'*'", "'<'", "'&&'", "'this'", "'boolean'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "INT", "BOOL", "ID", "WS", "COMMENT", "LINE_COMMENT",
  }
  staticData.RuleNames = []string{
    "goal", "mainClassDeclaration", "classDeclaration", "mainClassBody", 
    "mainMethod", "mainMethodDeclaration", "classBody", "fieldDeclaration", 
    "varDeclaration", "methodDeclaration", "methodBody", "formalParameters", 
    "formalParameterList", "formalParameter", "type", "statement", "expression", 
    "methodArgumentList", "intArrayType", "booleanType", "intType",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 42, 306, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1, 
	0, 1, 0, 5, 0, 45, 8, 0, 10, 0, 12, 0, 48, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 2, 1, 2, 1, 3, 1, 
	3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 
	5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 86, 8, 6, 10, 
	6, 12, 6, 89, 9, 6, 1, 6, 5, 6, 92, 8, 6, 10, 6, 12, 6, 95, 9, 6, 1, 6, 
	1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 132, 
	8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 138, 8, 10, 10, 10, 12, 10, 141, 
	9, 10, 1, 10, 4, 10, 144, 8, 10, 11, 10, 12, 10, 145, 1, 10, 1, 10, 1, 
	11, 1, 11, 3, 11, 152, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 
	159, 8, 12, 10, 12, 12, 12, 162, 9, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 
	14, 1, 14, 1, 14, 3, 14, 171, 8, 14, 1, 15, 1, 15, 5, 15, 175, 8, 15, 10, 
	15, 12, 15, 178, 9, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 
	1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 
	15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 
	1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 
	15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 226, 8, 15, 
	1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 
	16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 
	1, 16, 1, 16, 3, 16, 251, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 
	16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 
	1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 
	16, 5, 16, 280, 8, 16, 10, 16, 12, 16, 283, 9, 16, 1, 17, 1, 17, 1, 17, 
	1, 17, 5, 17, 289, 8, 17, 10, 17, 12, 17, 292, 9, 17, 3, 17, 294, 8, 17, 
	1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 
	20, 0, 1, 32, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 
	30, 32, 34, 36, 38, 40, 0, 0, 325, 0, 42, 1, 0, 0, 0, 2, 51, 1, 0, 0, 0, 
	4, 55, 1, 0, 0, 0, 6, 63, 1, 0, 0, 0, 8, 67, 1, 0, 0, 0, 10, 72, 1, 0, 
	0, 0, 12, 83, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 102, 1, 0, 0, 0, 18, 
	131, 1, 0, 0, 0, 20, 135, 1, 0, 0, 0, 22, 149, 1, 0, 0, 0, 24, 155, 1, 
	0, 0, 0, 26, 163, 1, 0, 0, 0, 28, 170, 1, 0, 0, 0, 30, 225, 1, 0, 0, 0, 
	32, 250, 1, 0, 0, 0, 34, 284, 1, 0, 0, 0, 36, 297, 1, 0, 0, 0, 38, 301, 
	1, 0, 0, 0, 40, 303, 1, 0, 0, 0, 42, 46, 3, 2, 1, 0, 43, 45, 3, 4, 2, 0, 
	44, 43, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 
	0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49, 50, 5, 0, 0, 1, 50, 
	1, 1, 0, 0, 0, 51, 52, 5, 1, 0, 0, 52, 53, 5, 39, 0, 0, 53, 54, 3, 6, 3, 
	0, 54, 3, 1, 0, 0, 0, 55, 56, 5, 1, 0, 0, 56, 59, 5, 39, 0, 0, 57, 58, 
	5, 2, 0, 0, 58, 60, 3, 28, 14, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 
	0, 60, 61, 1, 0, 0, 0, 61, 62, 3, 12, 6, 0, 62, 5, 1, 0, 0, 0, 63, 64, 
	5, 3, 0, 0, 64, 65, 3, 8, 4, 0, 65, 66, 5, 4, 0, 0, 66, 7, 1, 0, 0, 0, 
	67, 68, 3, 10, 5, 0, 68, 69, 5, 3, 0, 0, 69, 70, 3, 30, 15, 0, 70, 71, 
	5, 4, 0, 0, 71, 9, 1, 0, 0, 0, 72, 73, 5, 5, 0, 0, 73, 74, 5, 6, 0, 0, 
	74, 75, 5, 7, 0, 0, 75, 76, 5, 8, 0, 0, 76, 77, 5, 9, 0, 0, 77, 78, 5, 
	10, 0, 0, 78, 79, 5, 11, 0, 0, 79, 80, 5, 12, 0, 0, 80, 81, 5, 39, 0, 0, 
	81, 82, 5, 13, 0, 0, 82, 11, 1, 0, 0, 0, 83, 87, 5, 3, 0, 0, 84, 86, 3, 
	14, 7, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 
	88, 1, 0, 0, 0, 88, 93, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 92, 3, 18, 
	9, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 
	1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 5, 4, 0, 0, 
	97, 13, 1, 0, 0, 0, 98, 99, 3, 28, 14, 0, 99, 100, 5, 39, 0, 0, 100, 101, 
	5, 14, 0, 0, 101, 15, 1, 0, 0, 0, 102, 103, 3, 28, 14, 0, 103, 104, 5, 
	39, 0, 0, 104, 105, 5, 14, 0, 0, 105, 17, 1, 0, 0, 0, 106, 107, 5, 5, 0, 
	0, 107, 108, 3, 28, 14, 0, 108, 109, 5, 39, 0, 0, 109, 110, 3, 22, 11, 
	0, 110, 132, 1, 0, 0, 0, 111, 112, 3, 28, 14, 0, 112, 113, 5, 39, 0, 0, 
	113, 114, 3, 22, 11, 0, 114, 115, 6, 9, -1, 0, 115, 132, 1, 0, 0, 0, 116, 
	117, 5, 5, 0, 0, 117, 118, 5, 39, 0, 0, 118, 119, 3, 22, 11, 0, 119, 120, 
	6, 9, -1, 0, 120, 132, 1, 0, 0, 0, 121, 122, 5, 5, 0, 0, 122, 123, 3, 28, 
	14, 0, 123, 124, 3, 22, 11, 0, 124, 125, 6, 9, -1, 0, 125, 132, 1, 0, 0, 
	0, 126, 127, 5, 5, 0, 0, 127, 128, 3, 28, 14, 0, 128, 129, 5, 39, 0, 0, 
	129, 130, 6, 9, -1, 0, 130, 132, 1, 0, 0, 0, 131, 106, 1, 0, 0, 0, 131, 
	111, 1, 0, 0, 0, 131, 116, 1, 0, 0, 0, 131, 121, 1, 0, 0, 0, 131, 126, 
	1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 3, 20, 10, 0, 134, 19, 1, 0, 
	0, 0, 135, 139, 5, 3, 0, 0, 136, 138, 3, 16, 8, 0, 137, 136, 1, 0, 0, 0, 
	138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 
	143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 144, 3, 30, 15, 0, 143, 142, 
	1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 
	0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 5, 4, 0, 0, 148, 21, 1, 0, 0, 0, 
	149, 151, 5, 9, 0, 0, 150, 152, 3, 24, 12, 0, 151, 150, 1, 0, 0, 0, 151, 
	152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 5, 13, 0, 0, 154, 23, 
	1, 0, 0, 0, 155, 160, 3, 26, 13, 0, 156, 157, 5, 15, 0, 0, 157, 159, 3, 
	26, 13, 0, 158, 156, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 
	0, 0, 160, 161, 1, 0, 0, 0, 161, 25, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 
	163, 164, 3, 28, 14, 0, 164, 165, 5, 39, 0, 0, 165, 27, 1, 0, 0, 0, 166, 
	171, 3, 36, 18, 0, 167, 171, 3, 38, 19, 0, 168, 171, 3, 40, 20, 0, 169, 
	171, 5, 39, 0, 0, 170, 166, 1, 0, 0, 0, 170, 167, 1, 0, 0, 0, 170, 168, 
	1, 0, 0, 0, 170, 169, 1, 0, 0, 0, 171, 29, 1, 0, 0, 0, 172, 176, 5, 3, 
	0, 0, 173, 175, 3, 30, 15, 0, 174, 173, 1, 0, 0, 0, 175, 178, 1, 0, 0, 
	0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178, 
	176, 1, 0, 0, 0, 179, 226, 5, 4, 0, 0, 180, 181, 5, 16, 0, 0, 181, 182, 
	5, 9, 0, 0, 182, 183, 3, 32, 16, 0, 183, 184, 5, 13, 0, 0, 184, 185, 3, 
	30, 15, 0, 185, 186, 5, 17, 0, 0, 186, 187, 3, 30, 15, 0, 187, 226, 1, 
	0, 0, 0, 188, 189, 5, 18, 0, 0, 189, 190, 5, 9, 0, 0, 190, 191, 3, 32, 
	16, 0, 191, 192, 5, 13, 0, 0, 192, 193, 3, 30, 15, 0, 193, 226, 1, 0, 0, 
	0, 194, 195, 5, 19, 0, 0, 195, 196, 5, 9, 0, 0, 196, 197, 3, 32, 16, 0, 
	197, 198, 5, 13, 0, 0, 198, 199, 5, 14, 0, 0, 199, 226, 1, 0, 0, 0, 200, 
	201, 5, 39, 0, 0, 201, 202, 5, 20, 0, 0, 202, 203, 3, 32, 16, 0, 203, 204, 
	5, 14, 0, 0, 204, 226, 1, 0, 0, 0, 205, 206, 5, 39, 0, 0, 206, 207, 5, 
	11, 0, 0, 207, 208, 3, 32, 16, 0, 208, 209, 5, 12, 0, 0, 209, 210, 5, 20, 
	0, 0, 210, 211, 3, 32, 16, 0, 211, 212, 5, 14, 0, 0, 212, 226, 1, 0, 0, 
	0, 213, 214, 5, 21, 0, 0, 214, 215, 3, 32, 16, 0, 215, 216, 5, 14, 0, 0, 
	216, 226, 1, 0, 0, 0, 217, 218, 5, 22, 0, 0, 218, 219, 3, 32, 16, 0, 219, 
	220, 5, 23, 0, 0, 220, 221, 3, 34, 17, 0, 221, 222, 5, 24, 0, 0, 222, 223, 
	3, 32, 16, 0, 223, 224, 5, 14, 0, 0, 224, 226, 1, 0, 0, 0, 225, 172, 1, 
	0, 0, 0, 225, 180, 1, 0, 0, 0, 225, 188, 1, 0, 0, 0, 225, 194, 1, 0, 0, 
	0, 225, 200, 1, 0, 0, 0, 225, 205, 1, 0, 0, 0, 225, 213, 1, 0, 0, 0, 225, 
	217, 1, 0, 0, 0, 226, 31, 1, 0, 0, 0, 227, 228, 6, 16, -1, 0, 228, 229, 
	5, 27, 0, 0, 229, 251, 3, 32, 16, 14, 230, 231, 5, 28, 0, 0, 231, 251, 
	3, 32, 16, 13, 232, 233, 5, 29, 0, 0, 233, 234, 5, 30, 0, 0, 234, 235, 
	5, 11, 0, 0, 235, 236, 3, 32, 16, 0, 236, 237, 5, 12, 0, 0, 237, 251, 1, 
	0, 0, 0, 238, 239, 5, 29, 0, 0, 239, 240, 5, 39, 0, 0, 240, 241, 5, 9, 
	0, 0, 241, 251, 5, 13, 0, 0, 242, 251, 5, 37, 0, 0, 243, 251, 5, 38, 0, 
	0, 244, 251, 5, 39, 0, 0, 245, 251, 5, 35, 0, 0, 246, 247, 5, 9, 0, 0, 
	247, 248, 3, 32, 16, 0, 248, 249, 5, 13, 0, 0, 249, 251, 1, 0, 0, 0, 250, 
	227, 1, 0, 0, 0, 250, 230, 1, 0, 0, 0, 250, 232, 1, 0, 0, 0, 250, 238, 
	1, 0, 0, 0, 250, 242, 1, 0, 0, 0, 250, 243, 1, 0, 0, 0, 250, 244, 1, 0, 
	0, 0, 250, 245, 1, 0, 0, 0, 250, 246, 1, 0, 0, 0, 251, 281, 1, 0, 0, 0, 
	252, 253, 10, 10, 0, 0, 253, 254, 5, 31, 0, 0, 254, 280, 3, 32, 16, 11, 
	255, 256, 10, 9, 0, 0, 256, 257, 5, 27, 0, 0, 257, 280, 3, 32, 16, 10, 
	258, 259, 10, 8, 0, 0, 259, 260, 5, 32, 0, 0, 260, 280, 3, 32, 16, 9, 261, 
	262, 10, 7, 0, 0, 262, 263, 5, 33, 0, 0, 263, 280, 3, 32, 16, 8, 264, 265, 
	10, 6, 0, 0, 265, 266, 5, 34, 0, 0, 266, 280, 3, 32, 16, 7, 267, 268, 10, 
	17, 0, 0, 268, 269, 5, 11, 0, 0, 269, 270, 3, 32, 16, 0, 270, 271, 5, 12, 
	0, 0, 271, 280, 1, 0, 0, 0, 272, 273, 10, 16, 0, 0, 273, 274, 5, 25, 0, 
	0, 274, 280, 5, 26, 0, 0, 275, 276, 10, 15, 0, 0, 276, 277, 5, 25, 0, 0, 
	277, 278, 5, 39, 0, 0, 278, 280, 3, 34, 17, 0, 279, 252, 1, 0, 0, 0, 279, 
	255, 1, 0, 0, 0, 279, 258, 1, 0, 0, 0, 279, 261, 1, 0, 0, 0, 279, 264, 
	1, 0, 0, 0, 279, 267, 1, 0, 0, 0, 279, 272, 1, 0, 0, 0, 279, 275, 1, 0, 
	0, 0, 280, 283, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 
	282, 33, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 284, 293, 5, 9, 0, 0, 285, 290, 
	3, 32, 16, 0, 286, 287, 5, 15, 0, 0, 287, 289, 3, 32, 16, 0, 288, 286, 
	1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 
	0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 285, 1, 0, 0, 0, 
	293, 294, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 5, 13, 0, 0, 296, 
	35, 1, 0, 0, 0, 297, 298, 5, 30, 0, 0, 298, 299, 5, 11, 0, 0, 299, 300, 
	5, 12, 0, 0, 300, 37, 1, 0, 0, 0, 301, 302, 5, 36, 0, 0, 302, 39, 1, 0, 
	0, 0, 303, 304, 5, 30, 0, 0, 304, 41, 1, 0, 0, 0, 17, 46, 59, 87, 93, 131, 
	139, 145, 151, 160, 170, 176, 225, 250, 279, 281, 290, 293,
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

// MiniJavaParserInit initializes any static state used to implement MiniJavaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMiniJavaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniJavaParserInit() {
  staticData := &MiniJavaParserStaticData
  staticData.once.Do(minijavaParserInit)
}

// NewMiniJavaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMiniJavaParser(input antlr.TokenStream) *MiniJavaParser {
	MiniJavaParserInit()
	this := new(MiniJavaParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &MiniJavaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MiniJava.g4"

	return this
}


// MiniJavaParser tokens.
const (
	MiniJavaParserEOF = antlr.TokenEOF
	MiniJavaParserT__0 = 1
	MiniJavaParserT__1 = 2
	MiniJavaParserT__2 = 3
	MiniJavaParserT__3 = 4
	MiniJavaParserT__4 = 5
	MiniJavaParserT__5 = 6
	MiniJavaParserT__6 = 7
	MiniJavaParserT__7 = 8
	MiniJavaParserT__8 = 9
	MiniJavaParserT__9 = 10
	MiniJavaParserT__10 = 11
	MiniJavaParserT__11 = 12
	MiniJavaParserT__12 = 13
	MiniJavaParserT__13 = 14
	MiniJavaParserT__14 = 15
	MiniJavaParserT__15 = 16
	MiniJavaParserT__16 = 17
	MiniJavaParserT__17 = 18
	MiniJavaParserT__18 = 19
	MiniJavaParserT__19 = 20
	MiniJavaParserT__20 = 21
	MiniJavaParserT__21 = 22
	MiniJavaParserT__22 = 23
	MiniJavaParserT__23 = 24
	MiniJavaParserT__24 = 25
	MiniJavaParserT__25 = 26
	MiniJavaParserT__26 = 27
	MiniJavaParserT__27 = 28
	MiniJavaParserT__28 = 29
	MiniJavaParserT__29 = 30
	MiniJavaParserT__30 = 31
	MiniJavaParserT__31 = 32
	MiniJavaParserT__32 = 33
	MiniJavaParserT__33 = 34
	MiniJavaParserT__34 = 35
	MiniJavaParserT__35 = 36
	MiniJavaParserINT = 37
	MiniJavaParserBOOL = 38
	MiniJavaParserID = 39
	MiniJavaParserWS = 40
	MiniJavaParserCOMMENT = 41
	MiniJavaParserLINE_COMMENT = 42
)

// MiniJavaParser rules.
const (
	MiniJavaParserRULE_goal = 0
	MiniJavaParserRULE_mainClassDeclaration = 1
	MiniJavaParserRULE_classDeclaration = 2
	MiniJavaParserRULE_mainClassBody = 3
	MiniJavaParserRULE_mainMethod = 4
	MiniJavaParserRULE_mainMethodDeclaration = 5
	MiniJavaParserRULE_classBody = 6
	MiniJavaParserRULE_fieldDeclaration = 7
	MiniJavaParserRULE_varDeclaration = 8
	MiniJavaParserRULE_methodDeclaration = 9
	MiniJavaParserRULE_methodBody = 10
	MiniJavaParserRULE_formalParameters = 11
	MiniJavaParserRULE_formalParameterList = 12
	MiniJavaParserRULE_formalParameter = 13
	MiniJavaParserRULE_type = 14
	MiniJavaParserRULE_statement = 15
	MiniJavaParserRULE_expression = 16
	MiniJavaParserRULE_methodArgumentList = 17
	MiniJavaParserRULE_intArrayType = 18
	MiniJavaParserRULE_booleanType = 19
	MiniJavaParserRULE_intType = 20
)

// IGoalContext is an interface to support dynamic dispatch.
type IGoalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MainClassDeclaration() IMainClassDeclarationContext
	EOF() antlr.TerminalNode
	AllClassDeclaration() []IClassDeclarationContext
	ClassDeclaration(i int) IClassDeclarationContext

	// IsGoalContext differentiates from other interfaces.
	IsGoalContext()
}

type GoalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoalContext() *GoalContext {
	var p = new(GoalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_goal
	return p
}

func InitEmptyGoalContext(p *GoalContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_goal
}

func (*GoalContext) IsGoalContext() {}

func NewGoalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoalContext {
	var p = new(GoalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_goal

	return p
}

func (s *GoalContext) GetParser() antlr.Parser { return s.parser }

func (s *GoalContext) MainClassDeclaration() IMainClassDeclarationContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainClassDeclarationContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainClassDeclarationContext)
}

func (s *GoalContext) EOF() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserEOF, 0)
}

func (s *GoalContext) AllClassDeclaration() []IClassDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClassDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IClassDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClassDeclarationContext); ok {
			tst[i] = t.(IClassDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *GoalContext) ClassDeclaration(i int) IClassDeclarationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *GoalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GoalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterGoal(s)
	}
}

func (s *GoalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitGoal(s)
	}
}

func (s *GoalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitGoal(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) Goal() (localctx IGoalContext) {
	localctx = NewGoalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MiniJavaParserRULE_goal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.MainClassDeclaration()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MiniJavaParserT__0 {
		{
			p.SetState(43)
			p.ClassDeclaration()
		}


		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(MiniJavaParserEOF)
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


// IMainClassDeclarationContext is an interface to support dynamic dispatch.
type IMainClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	MainClassBody() IMainClassBodyContext

	// IsMainClassDeclarationContext differentiates from other interfaces.
	IsMainClassDeclarationContext()
}

type MainClassDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainClassDeclarationContext() *MainClassDeclarationContext {
	var p = new(MainClassDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_mainClassDeclaration
	return p
}

func InitEmptyMainClassDeclarationContext(p *MainClassDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_mainClassDeclaration
}

func (*MainClassDeclarationContext) IsMainClassDeclarationContext() {}

func NewMainClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainClassDeclarationContext {
	var p = new(MainClassDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_mainClassDeclaration

	return p
}

func (s *MainClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MainClassDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *MainClassDeclarationContext) MainClassBody() IMainClassBodyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainClassBodyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainClassBodyContext)
}

func (s *MainClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MainClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMainClassDeclaration(s)
	}
}

func (s *MainClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMainClassDeclaration(s)
	}
}

func (s *MainClassDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMainClassDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) MainClassDeclaration() (localctx IMainClassDeclarationContext) {
	localctx = NewMainClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MiniJavaParserRULE_mainClassDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(MiniJavaParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(MiniJavaParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(53)
		p.MainClassBody()
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


// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ClassBody() IClassBodyContext
	Type_() ITypeContext

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_classDeclaration
	return p
}

func InitEmptyClassDeclarationContext(p *ClassDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_classDeclaration
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *ClassDeclarationContext) ClassBody() IClassBodyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassBodyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *ClassDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitClassDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MiniJavaParserRULE_classDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(MiniJavaParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(MiniJavaParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == MiniJavaParserT__1 {
		{
			p.SetState(57)
			p.Match(MiniJavaParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Type_()
		}

	}
	{
		p.SetState(61)
		p.ClassBody()
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


// IMainClassBodyContext is an interface to support dynamic dispatch.
type IMainClassBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MainMethod() IMainMethodContext

	// IsMainClassBodyContext differentiates from other interfaces.
	IsMainClassBodyContext()
}

type MainClassBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainClassBodyContext() *MainClassBodyContext {
	var p = new(MainClassBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_mainClassBody
	return p
}

func InitEmptyMainClassBodyContext(p *MainClassBodyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_mainClassBody
}

func (*MainClassBodyContext) IsMainClassBodyContext() {}

func NewMainClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainClassBodyContext {
	var p = new(MainClassBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_mainClassBody

	return p
}

func (s *MainClassBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MainClassBodyContext) MainMethod() IMainMethodContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainMethodContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainMethodContext)
}

func (s *MainClassBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainClassBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MainClassBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMainClassBody(s)
	}
}

func (s *MainClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMainClassBody(s)
	}
}

func (s *MainClassBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMainClassBody(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) MainClassBody() (localctx IMainClassBodyContext) {
	localctx = NewMainClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MiniJavaParserRULE_mainClassBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(MiniJavaParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(64)
		p.MainMethod()
	}
	{
		p.SetState(65)
		p.Match(MiniJavaParserT__3)
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


// IMainMethodContext is an interface to support dynamic dispatch.
type IMainMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MainMethodDeclaration() IMainMethodDeclarationContext
	Statement() IStatementContext

	// IsMainMethodContext differentiates from other interfaces.
	IsMainMethodContext()
}

type MainMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainMethodContext() *MainMethodContext {
	var p = new(MainMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_mainMethod
	return p
}

func InitEmptyMainMethodContext(p *MainMethodContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_mainMethod
}

func (*MainMethodContext) IsMainMethodContext() {}

func NewMainMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainMethodContext {
	var p = new(MainMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_mainMethod

	return p
}

func (s *MainMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MainMethodContext) MainMethodDeclaration() IMainMethodDeclarationContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainMethodDeclarationContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainMethodDeclarationContext)
}

func (s *MainMethodContext) Statement() IStatementContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *MainMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MainMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMainMethod(s)
	}
}

func (s *MainMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMainMethod(s)
	}
}

func (s *MainMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMainMethod(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) MainMethod() (localctx IMainMethodContext) {
	localctx = NewMainMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MiniJavaParserRULE_mainMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.MainMethodDeclaration()
	}
	{
		p.SetState(68)
		p.Match(MiniJavaParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Statement()
	}
	{
		p.SetState(70)
		p.Match(MiniJavaParserT__3)
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


// IMainMethodDeclarationContext is an interface to support dynamic dispatch.
type IMainMethodDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsMainMethodDeclarationContext differentiates from other interfaces.
	IsMainMethodDeclarationContext()
}

type MainMethodDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainMethodDeclarationContext() *MainMethodDeclarationContext {
	var p = new(MainMethodDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_mainMethodDeclaration
	return p
}

func InitEmptyMainMethodDeclarationContext(p *MainMethodDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_mainMethodDeclaration
}

func (*MainMethodDeclarationContext) IsMainMethodDeclarationContext() {}

func NewMainMethodDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainMethodDeclarationContext {
	var p = new(MainMethodDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_mainMethodDeclaration

	return p
}

func (s *MainMethodDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MainMethodDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *MainMethodDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainMethodDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MainMethodDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMainMethodDeclaration(s)
	}
}

func (s *MainMethodDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMainMethodDeclaration(s)
	}
}

func (s *MainMethodDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMainMethodDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) MainMethodDeclaration() (localctx IMainMethodDeclarationContext) {
	localctx = NewMainMethodDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MiniJavaParserRULE_mainMethodDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(MiniJavaParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(MiniJavaParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(MiniJavaParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(MiniJavaParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(MiniJavaParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(MiniJavaParserT__9)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(MiniJavaParserT__10)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(MiniJavaParserT__11)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(MiniJavaParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(MiniJavaParserT__12)
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


// IClassBodyContext is an interface to support dynamic dispatch.
type IClassBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldDeclaration() []IFieldDeclarationContext
	FieldDeclaration(i int) IFieldDeclarationContext
	AllMethodDeclaration() []IMethodDeclarationContext
	MethodDeclaration(i int) IMethodDeclarationContext

	// IsClassBodyContext differentiates from other interfaces.
	IsClassBodyContext()
}

type ClassBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyContext() *ClassBodyContext {
	var p = new(ClassBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_classBody
	return p
}

func InitEmptyClassBodyContext(p *ClassBodyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_classBody
}

func (*ClassBodyContext) IsClassBodyContext() {}

func NewClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyContext {
	var p = new(ClassBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_classBody

	return p
}

func (s *ClassBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyContext) AllFieldDeclaration() []IFieldDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFieldDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDeclarationContext); ok {
			tst[i] = t.(IFieldDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ClassBodyContext) FieldDeclaration(i int) IFieldDeclarationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDeclarationContext)
}

func (s *ClassBodyContext) AllMethodDeclaration() []IMethodDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMethodDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IMethodDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMethodDeclarationContext); ok {
			tst[i] = t.(IMethodDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ClassBodyContext) MethodDeclaration(i int) IMethodDeclarationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodDeclarationContext)
}

func (s *ClassBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ClassBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterClassBody(s)
	}
}

func (s *ClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitClassBody(s)
	}
}

func (s *ClassBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitClassBody(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) ClassBody() (localctx IClassBodyContext) {
	localctx = NewClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MiniJavaParserRULE_classBody)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(MiniJavaParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(84)
				p.FieldDeclaration()
			}


		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 619549032480) != 0) {
		{
			p.SetState(90)
			p.MethodDeclaration()
		}


		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(MiniJavaParserT__3)
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


// IFieldDeclarationContext is an interface to support dynamic dispatch.
type IFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

	// IsFieldDeclarationContext differentiates from other interfaces.
	IsFieldDeclarationContext()
}

type FieldDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclarationContext() *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_fieldDeclaration
	return p
}

func InitEmptyFieldDeclarationContext(p *FieldDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_fieldDeclaration
}

func (*FieldDeclarationContext) IsFieldDeclarationContext() {}

func NewFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_fieldDeclaration

	return p
}

func (s *FieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FieldDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *FieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitFieldDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) FieldDeclaration() (localctx IFieldDeclarationContext) {
	localctx = NewFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MiniJavaParserRULE_fieldDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Type_()
	}
	{
		p.SetState(99)
		p.Match(MiniJavaParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(MiniJavaParserT__13)
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


// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_varDeclaration
	return p
}

func InitEmptyVarDeclarationContext(p *VarDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_varDeclaration
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitVarDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) VarDeclaration() (localctx IVarDeclarationContext) {
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MiniJavaParserRULE_varDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Type_()
	}
	{
		p.SetState(103)
		p.Match(MiniJavaParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(MiniJavaParserT__13)
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


// IMethodDeclarationContext is an interface to support dynamic dispatch.
type IMethodDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MethodBody() IMethodBodyContext
	Type_() ITypeContext
	ID() antlr.TerminalNode
	FormalParameters() IFormalParametersContext

	// IsMethodDeclarationContext differentiates from other interfaces.
	IsMethodDeclarationContext()
}

type MethodDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodDeclarationContext() *MethodDeclarationContext {
	var p = new(MethodDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_methodDeclaration
	return p
}

func InitEmptyMethodDeclarationContext(p *MethodDeclarationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_methodDeclaration
}

func (*MethodDeclarationContext) IsMethodDeclarationContext() {}

func NewMethodDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodDeclarationContext {
	var p = new(MethodDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_methodDeclaration

	return p
}

func (s *MethodDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodDeclarationContext) MethodBody() IMethodBodyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodBodyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodBodyContext)
}

func (s *MethodDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MethodDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *MethodDeclarationContext) FormalParameters() IFormalParametersContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParametersContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *MethodDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MethodDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMethodDeclaration(s)
	}
}

func (s *MethodDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMethodDeclaration(s)
	}
}

func (s *MethodDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMethodDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) MethodDeclaration() (localctx IMethodDeclarationContext) {
	localctx = NewMethodDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MiniJavaParserRULE_methodDeclaration)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(106)
			p.Match(MiniJavaParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Type_()
		}
		{
			p.SetState(108)
			p.Match(MiniJavaParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(109)
			p.FormalParameters()
		}


	case 2:
		{
			p.SetState(111)
			p.Type_()
		}
		{
			p.SetState(112)
			p.Match(MiniJavaParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(113)
			p.FormalParameters()
		}
		notifyErrorListeners("method declaration without public");


	case 3:
		{
			p.SetState(116)
			p.Match(MiniJavaParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(MiniJavaParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(118)
			p.FormalParameters()
		}
		notifyErrorListeners("method declaration without return type");


	case 4:
		{
			p.SetState(121)
			p.Match(MiniJavaParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Type_()
		}
		{
			p.SetState(123)
			p.FormalParameters()
		}
		notifyErrorListeners("method declaration without method name");


	case 5:
		{
			p.SetState(126)
			p.Match(MiniJavaParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Type_()
		}
		{
			p.SetState(128)
			p.Match(MiniJavaParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		notifyErrorListeners("method declaration without argument list");

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(133)
		p.MethodBody()
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


// IMethodBodyContext is an interface to support dynamic dispatch.
type IMethodBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVarDeclaration() []IVarDeclarationContext
	VarDeclaration(i int) IVarDeclarationContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsMethodBodyContext differentiates from other interfaces.
	IsMethodBodyContext()
}

type MethodBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodBodyContext() *MethodBodyContext {
	var p = new(MethodBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_methodBody
	return p
}

func InitEmptyMethodBodyContext(p *MethodBodyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_methodBody
}

func (*MethodBodyContext) IsMethodBodyContext() {}

func NewMethodBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodBodyContext {
	var p = new(MethodBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_methodBody

	return p
}

func (s *MethodBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodBodyContext) AllVarDeclaration() []IVarDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclarationContext); ok {
			tst[i] = t.(IVarDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *MethodBodyContext) VarDeclaration(i int) IVarDeclarationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *MethodBodyContext) AllStatement() []IStatementContext {
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

func (s *MethodBodyContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *MethodBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MethodBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMethodBody(s)
	}
}

func (s *MethodBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMethodBody(s)
	}
}

func (s *MethodBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMethodBody(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) MethodBody() (localctx IMethodBodyContext) {
	localctx = NewMethodBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MiniJavaParserRULE_methodBody)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(MiniJavaParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(136)
				p.VarDeclaration()
			}


		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 549762957320) != 0) {
		{
			p.SetState(142)
			p.Statement()
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
		p.Match(MiniJavaParserT__3)
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


// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FormalParameterList() IFormalParameterListContext

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
	p.RuleIndex = MiniJavaParserRULE_formalParameters
	return p
}

func InitEmptyFormalParametersContext(p *FormalParametersContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_formalParameters
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) FormalParameterList() IFormalParameterListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParameterListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitFormalParameters(s)
	}
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) FormalParameters() (localctx IFormalParametersContext) {
	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MiniJavaParserRULE_formalParameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(MiniJavaParserT__8)
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


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 619549032448) != 0) {
		{
			p.SetState(150)
			p.FormalParameterList()
		}

	}
	{
		p.SetState(153)
		p.Match(MiniJavaParserT__12)
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


// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFormalParameter() []IFormalParameterContext
	FormalParameter(i int) IFormalParameterContext

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_formalParameterList
	return p
}

func InitEmptyFormalParameterListContext(p *FormalParameterListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_formalParameterList
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllFormalParameter() []IFormalParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFormalParameterContext); ok {
			len++
		}
	}

	tst := make([]IFormalParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFormalParameterContext); ok {
			tst[i] = t.(IFormalParameterContext)
			i++
		}
	}

	return tst
}

func (s *FormalParameterListContext) FormalParameter(i int) IFormalParameterContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFormalParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFormalParameterContext)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitFormalParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) FormalParameterList() (localctx IFormalParameterListContext) {
	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MiniJavaParserRULE_formalParameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.FormalParameter()
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == MiniJavaParserT__14 {
		{
			p.SetState(156)
			p.Match(MiniJavaParserT__14)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(157)
			p.FormalParameter()
		}


		p.SetState(162)
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


// IFormalParameterContext is an interface to support dynamic dispatch.
type IFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	ID() antlr.TerminalNode

	// IsFormalParameterContext differentiates from other interfaces.
	IsFormalParameterContext()
}

type FormalParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterContext() *FormalParameterContext {
	var p = new(FormalParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_formalParameter
	return p
}

func InitEmptyFormalParameterContext(p *FormalParameterContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_formalParameter
}

func (*FormalParameterContext) IsFormalParameterContext() {}

func NewFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterContext {
	var p = new(FormalParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_formalParameter

	return p
}

func (s *FormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FormalParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *FormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterFormalParameter(s)
	}
}

func (s *FormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitFormalParameter(s)
	}
}

func (s *FormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) FormalParameter() (localctx IFormalParameterContext) {
	localctx = NewFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MiniJavaParserRULE_formalParameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Type_()
	}
	{
		p.SetState(164)
		p.Match(MiniJavaParserID)
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


// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntArrayType() IIntArrayTypeContext
	BooleanType() IBooleanTypeContext
	IntType() IIntTypeContext
	ID() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) IntArrayType() IIntArrayTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntArrayTypeContext)
}

func (s *TypeContext) BooleanType() IBooleanTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanTypeContext)
}

func (s *TypeContext) IntType() IIntTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntTypeContext)
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MiniJavaParserRULE_type)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.IntArrayType()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.BooleanType()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.IntType()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(169)
			p.Match(MiniJavaParserID)
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


// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = MiniJavaParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type WhileStatementContext struct {
	StatementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}


func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type PrintStatementContext struct {
	StatementContext
}

func NewPrintStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStatementContext {
	var p = new(PrintStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayAssignStatementContext struct {
	StatementContext
}

func NewArrayAssignStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAssignStatementContext {
	var p = new(ArrayAssignStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ArrayAssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAssignStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *ArrayAssignStatementContext) AllExpression() []IExpressionContext {
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

func (s *ArrayAssignStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *ArrayAssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterArrayAssignStatement(s)
	}
}

func (s *ArrayAssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitArrayAssignStatement(s)
	}
}

func (s *ArrayAssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitArrayAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type IfElseStatementContext struct {
	StatementContext
}

func NewIfElseStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseStatementContext {
	var p = new(IfElseStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfElseStatementContext) AllStatement() []IStatementContext {
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

func (s *IfElseStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *IfElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterIfElseStatement(s)
	}
}

func (s *IfElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitIfElseStatement(s)
	}
}

func (s *IfElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitIfElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type AssignStatementContext struct {
	StatementContext
}

func NewAssignStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStatementContext {
	var p = new(AssignStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *AssignStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *AssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type RecurStatementContext struct {
	StatementContext
}

func NewRecurStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecurStatementContext {
	var p = new(RecurStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *RecurStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecurStatementContext) AllExpression() []IExpressionContext {
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

func (s *RecurStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *RecurStatementContext) MethodArgumentList() IMethodArgumentListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodArgumentListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodArgumentListContext)
}


func (s *RecurStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterRecurStatement(s)
	}
}

func (s *RecurStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitRecurStatement(s)
	}
}

func (s *RecurStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitRecurStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type ReturnStatementContext struct {
	StatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}


type NestedStatementContext struct {
	StatementContext
}

func NewNestedStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedStatementContext {
	var p = new(NestedStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *NestedStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedStatementContext) AllStatement() []IStatementContext {
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

func (s *NestedStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *NestedStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterNestedStatement(s)
	}
}

func (s *NestedStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitNestedStatement(s)
	}
}

func (s *NestedStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitNestedStatement(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MiniJavaParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MiniJavaParserRULE_statement)
	var _la int

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNestedStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(MiniJavaParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 549762957320) != 0) {
			{
				p.SetState(173)
				p.Statement()
			}


			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(179)
			p.Match(MiniJavaParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		localctx = NewIfElseStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.Match(MiniJavaParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(MiniJavaParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(182)
			p.expression(0)
		}
		{
			p.SetState(183)
			p.Match(MiniJavaParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Statement()
		}
		{
			p.SetState(185)
			p.Match(MiniJavaParserT__16)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Statement()
		}


	case 3:
		localctx = NewWhileStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Match(MiniJavaParserT__17)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Match(MiniJavaParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(190)
			p.expression(0)
		}
		{
			p.SetState(191)
			p.Match(MiniJavaParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Statement()
		}


	case 4:
		localctx = NewPrintStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.Match(MiniJavaParserT__18)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Match(MiniJavaParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(196)
			p.expression(0)
		}
		{
			p.SetState(197)
			p.Match(MiniJavaParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(MiniJavaParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 5:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(200)
			p.Match(MiniJavaParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(MiniJavaParserT__19)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(202)
			p.expression(0)
		}
		{
			p.SetState(203)
			p.Match(MiniJavaParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 6:
		localctx = NewArrayAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(205)
			p.Match(MiniJavaParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(MiniJavaParserT__10)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(207)
			p.expression(0)
		}
		{
			p.SetState(208)
			p.Match(MiniJavaParserT__11)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(MiniJavaParserT__19)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(210)
			p.expression(0)
		}
		{
			p.SetState(211)
			p.Match(MiniJavaParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 7:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(213)
			p.Match(MiniJavaParserT__20)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(214)
			p.expression(0)
		}
		{
			p.SetState(215)
			p.Match(MiniJavaParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 8:
		localctx = NewRecurStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(217)
			p.Match(MiniJavaParserT__21)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(218)
			p.expression(0)
		}
		{
			p.SetState(219)
			p.Match(MiniJavaParserT__22)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(220)
			p.MethodArgumentList()
		}
		{
			p.SetState(221)
			p.Match(MiniJavaParserT__23)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(222)
			p.expression(0)
		}
		{
			p.SetState(223)
			p.Match(MiniJavaParserT__13)
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


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = MiniJavaParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type LtExpressionContext struct {
	ExpressionContext
}

func NewLtExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LtExpressionContext {
	var p = new(LtExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LtExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LtExpressionContext) AllExpression() []IExpressionContext {
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

func (s *LtExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *LtExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterLtExpression(s)
	}
}

func (s *LtExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitLtExpression(s)
	}
}

func (s *LtExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitLtExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ObjectInstantiationExpressionContext struct {
	ExpressionContext
}

func NewObjectInstantiationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectInstantiationExpressionContext {
	var p = new(ObjectInstantiationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ObjectInstantiationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectInstantiationExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}


func (s *ObjectInstantiationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterObjectInstantiationExpression(s)
	}
}

func (s *ObjectInstantiationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitObjectInstantiationExpression(s)
	}
}

func (s *ObjectInstantiationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitObjectInstantiationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayInstantiationExpressionContext struct {
	ExpressionContext
}

func NewArrayInstantiationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayInstantiationExpressionContext {
	var p = new(ArrayInstantiationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayInstantiationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInstantiationExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *ArrayInstantiationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterArrayInstantiationExpression(s)
	}
}

func (s *ArrayInstantiationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitArrayInstantiationExpression(s)
	}
}

func (s *ArrayInstantiationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitArrayInstantiationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type IdentifierExpressionContext struct {
	ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}


func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type MethodCallExpressionContext struct {
	ExpressionContext
}

func NewMethodCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallExpressionContext {
	var p = new(MethodCallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MethodCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MethodCallExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserID, 0)
}

func (s *MethodCallExpressionContext) MethodArgumentList() IMethodArgumentListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodArgumentListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodArgumentListContext)
}


func (s *MethodCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMethodCallExpression(s)
	}
}

func (s *MethodCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMethodCallExpression(s)
	}
}

func (s *MethodCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMethodCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type NotExpressionContext struct {
	ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type BooleanLitExpressionContext struct {
	ExpressionContext
}

func NewBooleanLitExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLitExpressionContext {
	var p = new(BooleanLitExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanLitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLitExpressionContext) BOOL() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserBOOL, 0)
}


func (s *BooleanLitExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterBooleanLitExpression(s)
	}
}

func (s *BooleanLitExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitBooleanLitExpression(s)
	}
}

func (s *BooleanLitExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitBooleanLitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ParenExpressionContext struct {
	ExpressionContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

func (s *ParenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitParenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type IntLitExpressionContext struct {
	ExpressionContext
}

func NewIntLitExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLitExpressionContext {
	var p = new(IntLitExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IntLitExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitExpressionContext) INT() antlr.TerminalNode {
	return s.GetToken(MiniJavaParserINT, 0)
}


func (s *IntLitExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterIntLitExpression(s)
	}
}

func (s *IntLitExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitIntLitExpression(s)
	}
}

func (s *IntLitExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitIntLitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type AndExpressionContext struct {
	ExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
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

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

func (s *AndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayAccessExpressionContext struct {
	ExpressionContext
}

func NewArrayAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAccessExpressionContext {
	var p = new(ArrayAccessExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ArrayAccessExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *ArrayAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterArrayAccessExpression(s)
	}
}

func (s *ArrayAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitArrayAccessExpression(s)
	}
}

func (s *ArrayAccessExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitArrayAccessExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type AddExpressionContext struct {
	ExpressionContext
}

func NewAddExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExpressionContext {
	var p = new(AddExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExpressionContext) AllExpression() []IExpressionContext {
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

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *AddExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterAddExpression(s)
	}
}

func (s *AddExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitAddExpression(s)
	}
}

func (s *AddExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitAddExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ThisExpressionContext struct {
	ExpressionContext
}

func NewThisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThisExpressionContext {
	var p = new(ThisExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ThisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterThisExpression(s)
	}
}

func (s *ThisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitThisExpression(s)
	}
}

func (s *ThisExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitThisExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayLengthExpressionContext struct {
	ExpressionContext
}

func NewArrayLengthExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLengthExpressionContext {
	var p = new(ArrayLengthExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayLengthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLengthExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *ArrayLengthExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterArrayLengthExpression(s)
	}
}

func (s *ArrayLengthExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitArrayLengthExpression(s)
	}
}

func (s *ArrayLengthExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitArrayLengthExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type NegExpressionContext struct {
	ExpressionContext
}

func NewNegExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegExpressionContext {
	var p = new(NegExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NegExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *NegExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterNegExpression(s)
	}
}

func (s *NegExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitNegExpression(s)
	}
}

func (s *NegExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitNegExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type SubExpressionContext struct {
	ExpressionContext
}

func NewSubExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExpressionContext {
	var p = new(SubExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SubExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExpressionContext) AllExpression() []IExpressionContext {
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

func (s *SubExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *SubExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterSubExpression(s)
	}
}

func (s *SubExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitSubExpression(s)
	}
}

func (s *SubExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitSubExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type MulExpressionContext struct {
	ExpressionContext
}

func NewMulExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExpressionContext {
	var p = new(MulExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExpressionContext) AllExpression() []IExpressionContext {
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

func (s *MulExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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


func (s *MulExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMulExpression(s)
	}
}

func (s *MulExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMulExpression(s)
	}
}

func (s *MulExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMulExpression(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MiniJavaParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MiniJavaParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, MiniJavaParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(228)
			p.Match(MiniJavaParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(229)
			p.expression(14)
		}


	case 2:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(230)
			p.Match(MiniJavaParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(231)
			p.expression(13)
		}


	case 3:
		localctx = NewArrayInstantiationExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(232)
			p.Match(MiniJavaParserT__28)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(MiniJavaParserT__29)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(MiniJavaParserT__10)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(235)
			p.expression(0)
		}
		{
			p.SetState(236)
			p.Match(MiniJavaParserT__11)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 4:
		localctx = NewObjectInstantiationExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(MiniJavaParserT__28)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(MiniJavaParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(MiniJavaParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(MiniJavaParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 5:
		localctx = NewIntLitExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(242)
			p.Match(MiniJavaParserINT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 6:
		localctx = NewBooleanLitExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(243)
			p.Match(MiniJavaParserBOOL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 7:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(244)
			p.Match(MiniJavaParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 8:
		localctx = NewThisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(245)
			p.Match(MiniJavaParserT__34)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 9:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(246)
			p.Match(MiniJavaParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(247)
			p.expression(0)
		}
		{
			p.SetState(248)
			p.Match(MiniJavaParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(279)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniJavaParserRULE_expression)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(253)
					p.Match(MiniJavaParserT__30)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(254)
					p.expression(11)
				}


			case 2:
				localctx = NewSubExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniJavaParserRULE_expression)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(256)
					p.Match(MiniJavaParserT__26)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(257)
					p.expression(10)
				}


			case 3:
				localctx = NewMulExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniJavaParserRULE_expression)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(259)
					p.Match(MiniJavaParserT__31)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(260)
					p.expression(9)
				}


			case 4:
				localctx = NewLtExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniJavaParserRULE_expression)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(262)
					p.Match(MiniJavaParserT__32)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(263)
					p.expression(8)
				}


			case 5:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniJavaParserRULE_expression)
				p.SetState(264)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(265)
					p.Match(MiniJavaParserT__33)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(266)
					p.expression(7)
				}


			case 6:
				localctx = NewArrayAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniJavaParserRULE_expression)
				p.SetState(267)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(268)
					p.Match(MiniJavaParserT__10)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(269)
					p.expression(0)
				}
				{
					p.SetState(270)
					p.Match(MiniJavaParserT__11)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 7:
				localctx = NewArrayLengthExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniJavaParserRULE_expression)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(273)
					p.Match(MiniJavaParserT__24)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(274)
					p.Match(MiniJavaParserT__25)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 8:
				localctx = NewMethodCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MiniJavaParserRULE_expression)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(276)
					p.Match(MiniJavaParserT__24)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(277)
					p.Match(MiniJavaParserID)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(278)
					p.MethodArgumentList()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMethodArgumentListContext is an interface to support dynamic dispatch.
type IMethodArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsMethodArgumentListContext differentiates from other interfaces.
	IsMethodArgumentListContext()
}

type MethodArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodArgumentListContext() *MethodArgumentListContext {
	var p = new(MethodArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_methodArgumentList
	return p
}

func InitEmptyMethodArgumentListContext(p *MethodArgumentListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_methodArgumentList
}

func (*MethodArgumentListContext) IsMethodArgumentListContext() {}

func NewMethodArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodArgumentListContext {
	var p = new(MethodArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_methodArgumentList

	return p
}

func (s *MethodArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodArgumentListContext) AllExpression() []IExpressionContext {
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

func (s *MethodArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *MethodArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MethodArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterMethodArgumentList(s)
	}
}

func (s *MethodArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitMethodArgumentList(s)
	}
}

func (s *MethodArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitMethodArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) MethodArgumentList() (localctx IMethodArgumentListContext) {
	localctx = NewMethodArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MiniJavaParserRULE_methodArgumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(MiniJavaParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 997371937280) != 0) {
		{
			p.SetState(285)
			p.expression(0)
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == MiniJavaParserT__14 {
			{
				p.SetState(286)
				p.Match(MiniJavaParserT__14)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(287)
				p.expression(0)
			}


			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(295)
		p.Match(MiniJavaParserT__12)
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


// IIntArrayTypeContext is an interface to support dynamic dispatch.
type IIntArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIntArrayTypeContext differentiates from other interfaces.
	IsIntArrayTypeContext()
}

type IntArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntArrayTypeContext() *IntArrayTypeContext {
	var p = new(IntArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_intArrayType
	return p
}

func InitEmptyIntArrayTypeContext(p *IntArrayTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_intArrayType
}

func (*IntArrayTypeContext) IsIntArrayTypeContext() {}

func NewIntArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntArrayTypeContext {
	var p = new(IntArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_intArrayType

	return p
}

func (s *IntArrayTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *IntArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IntArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterIntArrayType(s)
	}
}

func (s *IntArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitIntArrayType(s)
	}
}

func (s *IntArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitIntArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) IntArrayType() (localctx IIntArrayTypeContext) {
	localctx = NewIntArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MiniJavaParserRULE_intArrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(MiniJavaParserT__29)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Match(MiniJavaParserT__10)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(299)
		p.Match(MiniJavaParserT__11)
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


// IBooleanTypeContext is an interface to support dynamic dispatch.
type IBooleanTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBooleanTypeContext differentiates from other interfaces.
	IsBooleanTypeContext()
}

type BooleanTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanTypeContext() *BooleanTypeContext {
	var p = new(BooleanTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_booleanType
	return p
}

func InitEmptyBooleanTypeContext(p *BooleanTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_booleanType
}

func (*BooleanTypeContext) IsBooleanTypeContext() {}

func NewBooleanTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanTypeContext {
	var p = new(BooleanTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_booleanType

	return p
}

func (s *BooleanTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *BooleanTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BooleanTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterBooleanType(s)
	}
}

func (s *BooleanTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitBooleanType(s)
	}
}

func (s *BooleanTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitBooleanType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) BooleanType() (localctx IBooleanTypeContext) {
	localctx = NewBooleanTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MiniJavaParserRULE_booleanType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(MiniJavaParserT__35)
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


// IIntTypeContext is an interface to support dynamic dispatch.
type IIntTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIntTypeContext differentiates from other interfaces.
	IsIntTypeContext()
}

type IntTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntTypeContext() *IntTypeContext {
	var p = new(IntTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_intType
	return p
}

func InitEmptyIntTypeContext(p *IntTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MiniJavaParserRULE_intType
}

func (*IntTypeContext) IsIntTypeContext() {}

func NewIntTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntTypeContext {
	var p = new(IntTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MiniJavaParserRULE_intType

	return p
}

func (s *IntTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MiniJavaListener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MiniJavaVisitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MiniJavaParser) IntType() (localctx IIntTypeContext) {
	localctx = NewIntTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MiniJavaParserRULE_intType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(MiniJavaParserT__29)
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


func (p *MiniJavaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
			var t *ExpressionContext = nil
			if localctx != nil { t = localctx.(*ExpressionContext) }
			return p.Expression_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MiniJavaParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 17)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 16)

	case 7:
			return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

