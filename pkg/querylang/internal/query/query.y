%{
package query

import (
    "github.com/i582/CodeQuery/pkg/querylang/ast"
    "github.com/i582/CodeQuery/pkg/querylang/token"
)

%}

%union{
    node             ast.Node
    token            *token.Token
    list             []ast.Node
}

%token <token> T_SELECT
%token <token> T_FROM
%token <token> T_WHERE
%token <token> T_WITH
%token <token> T_COUNT
%token <token> T_LIMIT
%token <token> T_OFFSET
%token <token> T_ORDER_BY
%token <token> T_DESC
%token <token> T_ASC

%token <token> T_LNUMBER
%token <token> T_DNUMBER
%token <token> T_CONSTANT_STRING
%token <token> T_STRING
%token <token> T_VARIABLE
%token <token> T_COMMENT
%token <token> T_DOC_COMMENT
%token <token> T_WHITESPACE
%token <token> T_OR
%token <token> T_XOR
%token <token> T_AND
%token <token> T_NOT
%token <token> T_EQUAL
%token <token> T_NOT_EQUAL
%token <token> T_SMALLER
%token <token> T_GREATER
%token <token> T_SMALLER_OR_EQUAL
%token <token> T_GREATER_OR_EQUAL
%token <token> T_OBJECT_OPERATOR

%token <token> '"'
%token <token> '{'
%token <token> '}'
%token <token> ';'
%token <token> ':'
%token <token> '('
%token <token> ')'
%token <token> '['
%token <token> ']'
%token <token> '?'
%token <token> '&'
%token <token> '-'
%token <token> '+'
%token <token> '!'
%token <token> '~'
%token <token> '@'
%token <token> '$'
%token <token> ','
%token <token> '|'
%token <token> '='
%token <token> '^'
%token <token> '*'
%token <token> '/'
%token <token> '%'
%token <token> '<'
%token <token> '>'
%token <token> '.'

%left ','
%left T_OR
%left T_XOR
%left T_AND
%left '|'
%left '^'
%left '&'
%nonassoc T_EQUAL T_NOT_EQUAL
%nonassoc T_SMALLER T_SMALLER_OR_EQUAL T_GREATER T_GREATER_OR_EQUAL
%left '.'
%left '+' '-'
%left '*' '/' '%'
%right T_NOT
%right '['

%type <list> statement_list
%type <node> statement

%type <token> identifier

%type <node> select select_subject non_empty_selection_list selection
%type <node> from
%type <node> limit
%type <node> order
%type <token> order_dir
%type <node> from_subject
%type <node> where
%type <node> expr expr_or_select
%type <node> variable
%type <node> scalar
%type <node> deferencable
%type <node> argument argument_list non_empty_argument_list
%type <node> with
%type <node> with_subject with_subject_list non_empty_with_subject_list

%%

start:
        statement_list { yylex.(*Parser).rootNode = @b.NewRoot($1) }
;

statement_list:
        statement   { $$ = []ast.Node{$1} }
    |   /* empty */ { $$ = []ast.Node{} }
;

statement:
        error  { $$ = nil }
    |   select { $$ = $1 }
;

select:
        T_SELECT select_subject from where with limit order
            { $$ = @b.NewSelectExpr($1, $2, $3, $4, $5, $6, $7) }
;

select_subject:
        non_empty_selection_list { $$ = @b.NewSelectSubjectExpr($1, nil, nil) }
    |   T_COUNT '(' '*' ')'      { $$ = @b.NewSelectSubjectExpr(nil, $1, $3) }
    |   '*'                      { $$ = @b.NewSelectSubjectExpr(nil, nil, $1) }
;

order:
	T_ORDER_BY expr order_dir { $$ = @b.NewOrderByExpr($1, $2, $3) }
    |   /* empty */               { $$ = nil }
;

order_dir:
	T_ASC       { $$ = $1 }
    |   T_DESC      { $$ = $1 }
    |   /* empty */ { $$ = nil }
;

limit:
	T_LIMIT expr             { $$ = @b.NewLimitExpr($1, $2) }
    |   /* empty */              { $$ = nil }
;

from:
        T_FROM from_subject { $$ = @b.NewFromExpr($1, $2) }
;

from_subject:
        expr      { $$ = $1 }
;

where:
        T_WHERE expr { $$ = @b.NewWhereExpr($1, $2) }
    |   /* empty */  { $$ = nil }
;

with:
        T_WITH with_subject_list { $$ = @b.NewWithExpr($1, $2) }
    |   /* empty */              { $$ = nil }
;

with_subject_list:
        /* empty */                 { $$ = nil }
    |   non_empty_with_subject_list { $$ = $1 }
;

with_subject:
        expr { $$ = $1 }
;

non_empty_with_subject_list:
        non_empty_with_subject_list ',' with_subject
            { $$ = @b.AppendToSeparatedList($1, $2, $3) }
    |   with_subject
            { $$ = @b.NewSeparatedList($1) }
;

variable:
        identifier { $$ = @b.NewVariable($1) }
;

argument:
        expr_or_select { $$ = $1 }
;

argument_list:
        /* empty */             { $$ = nil }
    |   non_empty_argument_list { $$ = $1 }
;

non_empty_argument_list:
        non_empty_argument_list ',' argument
            { $$ = @b.AppendToSeparatedList($1, $2, $3) }
    |   argument
            { $$ = @b.NewSeparatedList($1) }
;

deferencable:
	deferencable T_OBJECT_OPERATOR T_STRING '(' argument_list ')'
	    { $$ = @b.NewMethodCallExpr($1, $2, $3, $4, $5, $6) }
    |   variable
            { $$ = $1 }
;

expr_or_select:
        expr   { $$ = $1 }
    |   select { $$ = $1 }
;

expr:
        deferencable                         { $$ = $1 }
    |   deferencable T_EQUAL expr            { $$ = @b.NewComparisonExpr($1, $2, $3) }
    |   deferencable T_NOT_EQUAL expr        { $$ = @b.NewComparisonExpr($1, $2, $3) }
    |   deferencable T_GREATER expr          { $$ = @b.NewComparisonExpr($1, $2, $3) }
    |   deferencable T_SMALLER expr          { $$ = @b.NewComparisonExpr($1, $2, $3) }
    |   deferencable T_SMALLER_OR_EQUAL expr { $$ = @b.NewComparisonExpr($1, $2, $3) }
    |   deferencable T_GREATER_OR_EQUAL expr { $$ = @b.NewComparisonExpr($1, $2, $3) }

    |   '(' expr ')'              { $$ = $2 }
    |   expr T_OR expr            { $$ = @b.NewBinaryExpr($1, $2, $3)}
    |   expr T_AND expr           { $$ = @b.NewBinaryExpr($1, $2, $3)}
    |   T_NOT expr                { $$ = @b.NewNotExpr($1, $2) }
    |   scalar                    { $$ = $1 }

    |   select                    { $$ = $1 }
;

scalar:
        T_LNUMBER         { $$ = @b.NewBasicLit($1) }
    |   T_DNUMBER         { $$ = @b.NewBasicLit($1) }
    |   T_CONSTANT_STRING { $$ = @b.NewBasicLit($1) }
;

identifier:
        T_STRING   { $$ = $1 }
;

selection:
        identifier { $$ = @b.NewIdentifier($1) }
;

non_empty_selection_list:
        non_empty_selection_list ',' selection
            { $$ = @b.AppendToSeparatedList($1, $2, $3) }
    |   selection
            { $$ = @b.NewSeparatedList($1) }
;

%%
