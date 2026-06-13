/*
BSD License

Copyright (c) 2020, Tom Everett
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

1. Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
3. Neither the name of Tom Everett nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

/*
* https://people.inf.ethz.ch/wirth/Oberon/Oberon07.Report.pdf
*/

// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

grammar oberon;

// Parser rules
ident
    : IDENT
    ;

qualident
    : (ident '.')? ident
    ;

identdef
    : ident '*'?
    ;

number
    : INTEGER
    | REAL
    | HEXNUMBER
    ;

constDeclaration
    : identdef '=' constExpression
    ;

constExpression
    : expression
    ;

typeDeclaration
    : identdef '=' type_
    ;

type_
    : qualident
    | arrayType
    | recordType
    | pointerType
    | procedureType
    ;

arrayType
    : ARRAY length (',' length)* OF type_
    ;

length
    : constExpression
    ;

recordType
    : RECORD ('(' baseType ')')? fieldListSequence? END
    ;

baseType
    : qualident
    ;

fieldListSequence
    : fieldList (';' fieldList)*
    ;

fieldList
    : identList ':' type_
    ;

identList
    : identdef (',' identdef)*
    ;

pointerType
    : POINTER TO type_
    ;

procedureType
    : PROCEDURE formalParameters?
    ;

variableDeclaration
    : identList ':' type_
    ;

expression
    : simpleExpression (relation simpleExpression)?
    ;

relation
    : '='
    | '#'
    | '<'
    | '<='
    | '>'
    | '>='
    | IN
    | IS
    ;

simpleExpression
    : ('+' | '-')? term (addOperator term)*
    ;

addOperator
    : '+'
    | '-'
    | OR
    ;

term
    : factor (mulOperator factor)*
    ;

mulOperator
    : '*'
    | '/'
    | DIV
    | MOD
    | '&'
    ;

factor
    : number
    | STRING
    | NIL
    | TRUE
    | FALSE
    | set_
    | designator (actualParameters)?
    | '(' expression ')'
    | '~' factor
    ;

designator
    : qualident selector*
    ;

selector
    : '.' ident
    | '[' expList ']'
    | '^'
    ;

set_
    : '{' (element (',' element)*)? '}'
    ;

element
    : expression ('..' expression)?
    ;

expList
    : expression (',' expression)*
    ;

actualParameters
    : '(' expList? ')'
    ;

statement
    : (
        assignment
        | procedureCall
        | ifStatement
        | caseStatement
        | whileStatement
        | repeatStatement
        | forStatement
    )?
    ;

assignment
    : designator ':=' expression
    ;

procedureCall
    : designator actualParameters?
    ;

statementSequence
    : statement (';' statement)*
    ;

ifStatement
    : IF expression THEN statementSequence (ELSIF expression THEN statementSequence)* (
        ELSE statementSequence
    )? END
    ;

caseStatement
    : CASE expression OF case_ ('|' case_)* END
    ;

case_
    : (caseLabelList ':' statementSequence)?
    ;

caseLabelList
    : labelRange (',' labelRange)*
    ;

labelRange
    : label ('..' label)?
    ;

label
    : INTEGER
    | STRING
    | qualident
    ;

whileStatement
    : WHILE expression DO statementSequence (ELSIF expression DO statementSequence)* END
    ;

repeatStatement
    : REPEAT statementSequence UNTIL expression
    ;

forStatement
    : FOR ident ':=' expression TO expression (BY constExpression)? DO statementSequence END
    ;

procedureDeclaration
    : procedureHeading ';' procedureBody ident
    ;

procedureHeading
    : PROCEDURE identdef formalParameters?
    ;

procedureBody
    : declarationSequence (BEGIN statementSequence)? (RETURN expression)? END
    ;

declarationSequence
    : (CONST (constDeclaration ';')*)? (TYPE (typeDeclaration ';')*)? (
        VAR (variableDeclaration ';')*
    )? (procedureDeclaration ';')*
    ;

formalParameters
    : '(' (fPSection (';' fPSection)*)? ')' (':' qualident)?
    ;

fPSection
    : VAR? ident (',' ident)* ':' formalType
    ;

formalType
    : (ARRAY OF)* qualident
    ;

module
    : MODULE ident ';' importList? declarationSequence (BEGIN statementSequence)? END ident '.' EOF
    ;

importList
    : IMPORT import_ (',' import_)* ';'
    ;

import_
    : ident (':=' ident)?
    ;

// Lexer rules (tokens)
ARRAY   : 'ARRAY';
OF      : 'OF';
END     : 'END';
POINTER : 'POINTER';
TO      : 'TO';
RECORD  : 'RECORD';
PROCEDURE : 'PROCEDURE';
IN      : 'IN';
IS      : 'IS';
OR      : 'OR';
DIV     : 'DIV';
MOD     : 'MOD';
NIL     : 'NIL';
TRUE    : 'TRUE';
FALSE   : 'FALSE';
IF      : 'IF';
THEN    : 'THEN';
ELSIF   : 'ELSIF';
ELSE    : 'ELSE';
CASE    : 'CASE';
WHILE   : 'WHILE';
DO      : 'DO';
REPEAT  : 'REPEAT';
UNTIL   : 'UNTIL';
FOR     : 'FOR';
BY      : 'BY';
BEGIN   : 'BEGIN';
RETURN  : 'RETURN';MODULE TestSet6;
  VAR a, b, c: SET;
      ok: BOOLEAN;
      result: INTEGER;
BEGIN
  a := {1,2,3};
  b := {2,3,4};
  c := (a + b) * {2,3,4,5};
  ok := (c = {2,3,4}) & (a - b = {1}) & (a / b = {1,4});
  IF ok THEN
    result := 0
  ELSE
    result := 1
  END
END TestSet6.
CONST   : 'CONST';
TYPE    : 'TYPE';
VAR     : 'VAR';
MODULE  : 'MODULE';
IMPORT  : 'IMPORT';

STRING
    : ('"' .*? '"')
    | (DIGIT HEXDIGIT* 'X')
    ;

INTEGER : DIGIT+;
HEXNUMBER : DIGIT HEXDIGIT* 'H';
REAL    : DIGIT+ '.' DIGIT* ('E' ('+'|'-')? DIGIT+)?;

fragment HEXDIGIT : DIGIT | [A-F];
fragment DIGIT    : [0-9];

IDENT   : LETTER (LETTER | DIGIT)*;
fragment LETTER : [a-zA-Z];

COMMENT : '(*' .*? '*)' -> skip;
WS      : [ \t\r\n] -> skip;