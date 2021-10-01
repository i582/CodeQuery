package query

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/i582/CodeQuery/pkg/querylang/token"
)

%%{
    machine lexer;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

func initLexer(lex *Lexer)  {
    %% write init;
}

func (lex *Lexer) Lex() *token.Token {
    eof := lex.pe
    var tok token.ID

    tkn := lex.tokenPool.Get()

    lblStart := 0
    lblEnd   := 0

    _, _ = lblStart, lblEnd

    %%{
        action new_line {
            if lex.data[lex.p] == '\n' {
                lex.newLines.Append(lex.p+1)
            }

            if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
                lex.newLines.Append(lex.p+1)
            }
        }

        action is_not_comment_end { lex.isNotNewLine()  }

        newline = ('\r\n' >(nl, 1) | '\r' >(nl, 0) | '\n' >(nl, 0)) $new_line %{};
        any_line = any | newline;
        whitespace = [\t\v\f ];
        whitespace_line = [\t\v\f ] | newline;

        lnum = [0-9]+('_'[0-9]+)*;
        dnum = (lnum?"." lnum)|(lnum"."lnum?);
        hnum = '0x'i[0-9a-fA-F]+('_'[0-9a-fA-F]+)*;
        bnum = '0b'i[01]+('_'[01]+)*;
        onum = '0o'i[0-7]+('_'[0-7]+)*;

        exponent_dnum  = (lnum | dnum) ('e'|'E') ('+'|'-')? lnum;
        varname_first  = [a-zA-Z_] | (0x0080..0x00FF);
        varname_second = varname_first | [0-9];
        varname        = varname_first (varname_second)*;
        operators      = ';'|':'|','|'.'|'['|']'|'('|')'|'|'|'/'|'^'|'&'|'+'|'-'|'*'|'='|'%'|'!'|'~'|'$'|'<'|'>'|'?'|'@';

        prepush { lex.growCallStack(); }

        constant_string =
            start: (
                "'"         -> qoute
                | "b"i? '"' -> double_qoute
            ),

            # single qoute string

            qoute: (
                (any - [\\'\r\n]) -> qoute
                | "\r" @new_line  -> qoute
                | "\n" @new_line  -> qoute
                | "\\"            -> qoute_any
                | "'"             -> final
            ),
            qoute_any: (
                (any - [\r\n])   -> qoute
                | "\r" @new_line -> qoute
                | "\n" @new_line -> qoute
            ),

            # double qoute string

            double_qoute: (
                (any - [\\"${\r\n]) -> double_qoute
                | "\r" @new_line    -> double_qoute
                | "\n" @new_line    -> double_qoute
                | "\\"              -> double_qoute_any
                | '"'               -> final
                | '$'               -> double_qoute_nonvarname
                | '{'               -> double_qoute_nondollar
            ),
            double_qoute_any: (
                (any - [\r\n])     -> double_qoute
                | "\r" @new_line   -> double_qoute
                | "\n" @new_line   -> double_qoute
            ),
            double_qoute_nondollar: (
                (any - [\\$"\r\n]) -> double_qoute
                | "\r" @new_line   -> double_qoute
                | "\n" @new_line   -> double_qoute
                | "\\"             -> double_qoute_any
                | '"'              -> final
            ),
            double_qoute_nonvarname: (
                (any - [\\${"\r\n] - varname_first) -> double_qoute
                | "\r" @new_line                    -> double_qoute
                | "\n" @new_line                    -> double_qoute
                | "\\"                              -> double_qoute_any
                | '$'                               -> double_qoute_nonvarname
                | '"'                               -> final
            );

        main := |*
            any => {
                fnext php;
                lex.ungetCnt(1)
            };
        *|;

        php := |*
            whitespace_line*                => {lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)};

            (dnum | exponent_dnum)          => {lex.setTokenPos(tkn); tok = token.T_DNUMBER; fbreak;};
            bnum => {
                s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
                _, err := strconv.ParseInt(s, 2, 0)

                if err == nil {
                    lex.setTokenPos(tkn); tok = token.T_LNUMBER; fbreak;
                }

                lex.setTokenPos(tkn); tok = token.T_DNUMBER; fbreak;
            };
            lnum => {
                base := 10
                if lex.data[lex.ts] == '0' {
                    base = 8
                }

                s := strings.ReplaceAll(string(lex.data[lex.ts:lex.te]), "_", "")
                _, err := strconv.ParseInt(s, base, 0)

                if err == nil {
                    lex.setTokenPos(tkn); tok = token.T_LNUMBER; fbreak;
                }

                lex.setTokenPos(tkn); tok = token.T_DNUMBER; fbreak;
            };
            hnum => {
                s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
                _, err := strconv.ParseInt(s, 16, 0)

                if err == nil {
                    lex.setTokenPos(tkn); tok = token.T_LNUMBER; fbreak;
                }

                lex.setTokenPos(tkn); tok = token.T_DNUMBER; fbreak;
            };
            onum => {
                s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
                _, err := strconv.ParseInt(s, 8, 0)

                if err == nil {
                    lex.setTokenPos(tkn); tok = token.T_LNUMBER; fbreak;
                }

                lex.setTokenPos(tkn); tok = token.T_DNUMBER; fbreak;
            };

            '='   => {lex.setTokenPos(tkn); tok = token.T_EQUAL; fbreak;};
            '<>'  => {lex.setTokenPos(tkn); tok = token.T_NOT_EQUAL; fbreak;};
            '<'   => {lex.setTokenPos(tkn); tok = token.T_SMALLER; fbreak;};
            '>'   => {lex.setTokenPos(tkn); tok = token.T_GREATER; fbreak;};
            '<='  => {lex.setTokenPos(tkn); tok = token.T_SMALLER_OR_EQUAL; fbreak;};
            '>='  => {lex.setTokenPos(tkn); tok = token.T_GREATER_OR_EQUAL; fbreak;};

            'select'i  => {lex.setTokenPos(tkn); tok = token.T_SELECT; fbreak;};
            'from'i    => {lex.setTokenPos(tkn); tok = token.T_FROM; fbreak;};
            'where'i   => {lex.setTokenPos(tkn); tok = token.T_WHERE; fbreak;};
            'with'i    => {lex.setTokenPos(tkn); tok = token.T_WITH; fbreak;};
            'and'i     => {lex.setTokenPos(tkn); tok = token.T_AND; fbreak;};
            'or'i      => {lex.setTokenPos(tkn); tok = token.T_OR; fbreak;};
            'not'i     => {lex.setTokenPos(tkn); tok = token.T_NOT; fbreak;};
            'count'i   => {lex.setTokenPos(tkn); tok = token.T_COUNT; fbreak;};
            'limit'i   => {lex.setTokenPos(tkn); tok = token.T_LIMIT; fbreak;};
            'offset'i  => {lex.setTokenPos(tkn); tok = token.T_OFFSET; fbreak;};
            'desc'i    => {lex.setTokenPos(tkn); tok = token.T_DESC; fbreak;};
            'asc'i     => {lex.setTokenPos(tkn); tok = token.T_ASC; fbreak;};

            'order'i whitespace_line+ 'by'i => {lex.setTokenPos(tkn); tok = token.T_ORDER_BY; fbreak;};

            "."        => { lex.setTokenPos(tkn); tok = token.T_OBJECT_OPERATOR; fnext property; fbreak; };

            constant_string => {
                lex.setTokenPos(tkn);
                tok = token.T_CONSTANT_STRING;
                fbreak;
            };

            '--' any_line* when is_not_comment_end => {
                lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
            };

            '/*' any_line* :>> '*/' {
                isDocComment := false;
                if lex.te - lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
                    isDocComment = true;
                }

                if isDocComment {
                    lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
                } else {
                    lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
                }
            };

            operators => {
                lex.setTokenPos(tkn);
                tok = token.ID(int(lex.data[lex.ts]));
                fbreak;
            };

            varname  => { lex.setTokenPos(tkn); tok = token.T_STRING;   fbreak; };

            any_line => {
                c := lex.data[lex.p]
                lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c));
            };
        *|;

        property := |*
            whitespace_line* => {lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)};
            "."              => {lex.setTokenPos(tkn); tok = token.T_OBJECT_OPERATOR; fbreak;};
            varname          => {lex.setTokenPos(tkn); tok = token.T_STRING; fnext php; fbreak;};
            any              => {lex.ungetCnt(1); fgoto php;};
        *|;

        template_string := |*
            '"'                => {lex.setTokenPos(tkn); tok = token.ID(int('"')); fnext php; fbreak;};

        *|;


        write exec;
    }%%

    tkn.Value = lex.data[lex.ts:lex.te]
    tkn.ID = token.ID(tok)

    return tkn
}