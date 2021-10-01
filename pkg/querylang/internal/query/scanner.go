//line pkg/querylang/internal/query/scanner.rl:1
package query

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/i582/CodeQuery/pkg/querylang/token"
)

//line pkg/querylang/internal/query/scanner.go:15
const lexer_start int = 23
const lexer_first_final int = 23
const lexer_error int = 0

const lexer_en_main int = 23
const lexer_en_php int = 24
const lexer_en_property int = 86
const lexer_en_template_string int = 91

//line pkg/querylang/internal/query/scanner.rl:17

func initLexer(lex *Lexer) {

//line pkg/querylang/internal/query/scanner.go:31
	{
		lex.cs = lexer_start
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line pkg/querylang/internal/query/scanner.rl:21
}

func (lex *Lexer) Lex() *token.Token {
	eof := lex.pe
	var tok token.ID

	tkn := lex.tokenPool.Get()

	lblStart := 0
	lblEnd := 0

	_, _ = lblStart, lblEnd

//line pkg/querylang/internal/query/scanner.go:54
	{
		var _widec int16
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		goto _resume

	_again:
		switch lex.cs {
		case 23:
			goto st23
		case 24:
			goto st24
		case 25:
			goto st25
		case 26:
			goto st26
		case 1:
			goto st1
		case 27:
			goto st27
		case 28:
			goto st28
		case 2:
			goto st2
		case 3:
			goto st3
		case 4:
			goto st4
		case 5:
			goto st5
		case 29:
			goto st29
		case 6:
			goto st6
		case 7:
			goto st7
		case 30:
			goto st30
		case 31:
			goto st31
		case 32:
			goto st32
		case 33:
			goto st33
		case 34:
			goto st34
		case 8:
			goto st8
		case 9:
			goto st9
		case 35:
			goto st35
		case 10:
			goto st10
		case 36:
			goto st36
		case 11:
			goto st11
		case 12:
			goto st12
		case 13:
			goto st13
		case 37:
			goto st37
		case 38:
			goto st38
		case 39:
			goto st39
		case 14:
			goto st14
		case 15:
			goto st15
		case 40:
			goto st40
		case 16:
			goto st16
		case 41:
			goto st41
		case 17:
			goto st17
		case 42:
			goto st42
		case 43:
			goto st43
		case 44:
			goto st44
		case 45:
			goto st45
		case 46:
			goto st46
		case 47:
			goto st47
		case 48:
			goto st48
		case 49:
			goto st49
		case 50:
			goto st50
		case 51:
			goto st51
		case 52:
			goto st52
		case 53:
			goto st53
		case 54:
			goto st54
		case 55:
			goto st55
		case 56:
			goto st56
		case 57:
			goto st57
		case 58:
			goto st58
		case 59:
			goto st59
		case 60:
			goto st60
		case 61:
			goto st61
		case 62:
			goto st62
		case 63:
			goto st63
		case 64:
			goto st64
		case 65:
			goto st65
		case 66:
			goto st66
		case 67:
			goto st67
		case 68:
			goto st68
		case 69:
			goto st69
		case 70:
			goto st70
		case 71:
			goto st71
		case 72:
			goto st72
		case 73:
			goto st73
		case 74:
			goto st74
		case 18:
			goto st18
		case 19:
			goto st19
		case 20:
			goto st20
		case 21:
			goto st21
		case 75:
			goto st75
		case 76:
			goto st76
		case 77:
			goto st77
		case 78:
			goto st78
		case 79:
			goto st79
		case 80:
			goto st80
		case 81:
			goto st81
		case 82:
			goto st82
		case 83:
			goto st83
		case 84:
			goto st84
		case 85:
			goto st85
		case 86:
			goto st86
		case 87:
			goto st87
		case 88:
			goto st88
		case 22:
			goto st22
		case 89:
			goto st89
		case 90:
			goto st90
		case 91:
			goto st91
		case 0:
			goto st0
		}

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof
		}
	_resume:
		switch lex.cs {
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 1:
			goto st_case_1
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 29:
			goto st_case_29
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 35:
			goto st_case_35
		case 10:
			goto st_case_10
		case 36:
			goto st_case_36
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 40:
			goto st_case_40
		case 16:
			goto st_case_16
		case 41:
			goto st_case_41
		case 17:
			goto st_case_17
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 22:
			goto st_case_22
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 0:
			goto st_case_0
		}
		goto st_out
	tr44:
		lex.cs = 23
//line pkg/querylang/internal/query/scanner.rl:120
		lex.te = (lex.p) + 1
		{
			lex.cs = 24
			lex.ungetCnt(1)
		}
		goto _again
	st23:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof23
		}
	st_case_23:
//line NONE:1
		lex.ts = (lex.p)

//line pkg/querylang/internal/query/scanner.go:461
		goto tr44
	tr0:
//line pkg/querylang/internal/query/scanner.rl:127
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st24
	tr2:
//line NONE:1
		switch lex.act {
		case 3:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 4:
			{
				(lex.p) = (lex.te) - 1

				s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
				_, err := strconv.ParseInt(s, 2, 0)

				if err == nil {
					lex.setTokenPos(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						lex.cs = 24
						goto _out
					}
				}

				lex.setTokenPos(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 5:
			{
				(lex.p) = (lex.te) - 1

				base := 10
				if lex.data[lex.ts] == '0' {
					base = 8
				}

				s := strings.ReplaceAll(string(lex.data[lex.ts:lex.te]), "_", "")
				_, err := strconv.ParseInt(s, base, 0)

				if err == nil {
					lex.setTokenPos(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						lex.cs = 24
						goto _out
					}
				}

				lex.setTokenPos(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 6:
			{
				(lex.p) = (lex.te) - 1

				s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
				_, err := strconv.ParseInt(s, 16, 0)

				if err == nil {
					lex.setTokenPos(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						lex.cs = 24
						goto _out
					}
				}

				lex.setTokenPos(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 7:
			{
				(lex.p) = (lex.te) - 1

				s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
				_, err := strconv.ParseInt(s, 8, 0)

				if err == nil {
					lex.setTokenPos(tkn)
					tok = token.T_LNUMBER
					{
						(lex.p)++
						lex.cs = 24
						goto _out
					}
				}

				lex.setTokenPos(tkn)
				tok = token.T_DNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 14:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_SELECT
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 15:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_FROM
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 16:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_WHERE
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 17:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_WITH
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 18:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_AND
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 20:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_NOT
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 21:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_COUNT
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 22:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_LIMIT
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 23:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_OFFSET
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 24:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_DESC
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 25:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_ASC
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 32:
			{
				(lex.p) = (lex.te) - 1
				lex.setTokenPos(tkn)
				tok = token.T_STRING
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}
		case 33:
			{
				(lex.p) = (lex.te) - 1

				c := lex.data[lex.p]
				lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
			}
		}

		goto st24
	tr5:
//line pkg/querylang/internal/query/scanner.rl:200
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_CONSTANT_STRING
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr9:
//line pkg/querylang/internal/query/scanner.rl:231
		(lex.p) = (lex.te) - 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st24
	tr15:
//line pkg/querylang/internal/query/scanner.rl:129
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr17:
//line pkg/querylang/internal/query/scanner.rl:223
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPos(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr26:
//line pkg/querylang/internal/query/scanner.rl:210
		lex.te = (lex.p) + 1
		{
			isDocComment := false
			if lex.te-lex.ts > 4 && string(lex.data[lex.ts:lex.ts+3]) == "/**" {
				isDocComment = true
			}

			if isDocComment {
				lex.addFreeFloatingToken(tkn, token.T_DOC_COMMENT, lex.ts, lex.te)
			} else {
				lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
			}
		}
		goto st24
	tr27:
//line pkg/querylang/internal/query/scanner.rl:140
		(lex.p) = (lex.te) - 1
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}

			s := strings.ReplaceAll(string(lex.data[lex.ts:lex.te]), "_", "")
			_, err := strconv.ParseInt(s, base, 0)

			if err == nil {
				lex.setTokenPos(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}

			lex.setTokenPos(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr32:
//line pkg/querylang/internal/query/scanner.rl:229
		(lex.p) = (lex.te) - 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr41:
//line pkg/querylang/internal/query/scanner.rl:196
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_ORDER_BY
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr45:
//line pkg/querylang/internal/query/scanner.rl:231
		lex.te = (lex.p) + 1
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st24
	tr48:
//line pkg/querylang/internal/query/scanner.rl:223
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr56:
//line pkg/querylang/internal/query/scanner.rl:176
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_EQUAL
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr69:
//line pkg/querylang/internal/query/scanner.rl:127
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st24
	tr71:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:127
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st24
	tr75:
//line pkg/querylang/internal/query/scanner.rl:231
		lex.te = (lex.p)
		(lex.p)--
		{
			c := lex.data[lex.p]
			lex.error(fmt.Sprintf("WARNING: Unexpected character in input: '%c' (ASCII=%d)", c, c))
		}
		goto st24
	tr76:
//line pkg/querylang/internal/query/scanner.rl:223
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPos(tkn)
			tok = token.ID(int(lex.data[lex.ts]))
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr78:
//line pkg/querylang/internal/query/scanner.rl:206
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st24
	tr81:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:206
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_COMMENT, lex.ts, lex.te)
		}
		goto st24
	tr85:
		lex.cs = 24
//line pkg/querylang/internal/query/scanner.rl:198
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPos(tkn)
			tok = token.T_OBJECT_OPERATOR
			lex.cs = 86
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	tr86:
//line pkg/querylang/internal/query/scanner.rl:129
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPos(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr89:
//line pkg/querylang/internal/query/scanner.rl:140
		lex.te = (lex.p)
		(lex.p)--
		{
			base := 10
			if lex.data[lex.ts] == '0' {
				base = 8
			}

			s := strings.ReplaceAll(string(lex.data[lex.ts:lex.te]), "_", "")
			_, err := strconv.ParseInt(s, base, 0)

			if err == nil {
				lex.setTokenPos(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}

			lex.setTokenPos(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr95:
//line pkg/querylang/internal/query/scanner.rl:130
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
			_, err := strconv.ParseInt(s, 2, 0)

			if err == nil {
				lex.setTokenPos(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}

			lex.setTokenPos(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr96:
//line pkg/querylang/internal/query/scanner.rl:165
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
			_, err := strconv.ParseInt(s, 8, 0)

			if err == nil {
				lex.setTokenPos(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}

			lex.setTokenPos(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr97:
//line pkg/querylang/internal/query/scanner.rl:155
		lex.te = (lex.p)
		(lex.p)--
		{
			s := strings.ReplaceAll(string(lex.data[lex.ts+2:lex.te]), "_", "")
			_, err := strconv.ParseInt(s, 16, 0)

			if err == nil {
				lex.setTokenPos(tkn)
				tok = token.T_LNUMBER
				{
					(lex.p)++
					lex.cs = 24
					goto _out
				}
			}

			lex.setTokenPos(tkn)
			tok = token.T_DNUMBER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr98:
//line pkg/querylang/internal/query/scanner.rl:178
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPos(tkn)
			tok = token.T_SMALLER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr99:
//line pkg/querylang/internal/query/scanner.rl:180
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_SMALLER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr100:
//line pkg/querylang/internal/query/scanner.rl:177
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_NOT_EQUAL
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr101:
//line pkg/querylang/internal/query/scanner.rl:179
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPos(tkn)
			tok = token.T_GREATER
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr102:
//line pkg/querylang/internal/query/scanner.rl:181
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_GREATER_OR_EQUAL
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr103:
//line pkg/querylang/internal/query/scanner.rl:229
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPos(tkn)
			tok = token.T_STRING
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	tr130:
//line pkg/querylang/internal/query/scanner.rl:188
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPos(tkn)
			tok = token.T_OR
			{
				(lex.p)++
				lex.cs = 24
				goto _out
			}
		}
		goto st24
	st24:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof24
		}
	st_case_24:
//line NONE:1
		lex.ts = (lex.p)

//line pkg/querylang/internal/query/scanner.go:849
		switch lex.data[(lex.p)] {
		case 10:
			goto tr1
		case 13:
			goto tr47
		case 32:
			goto tr46
		case 33:
			goto tr48
		case 34:
			goto tr49
		case 39:
			goto tr50
		case 45:
			goto st30
		case 46:
			goto st33
		case 47:
			goto tr53
		case 48:
			goto tr54
		case 60:
			goto st43
		case 61:
			goto tr56
		case 62:
			goto st44
		case 65:
			goto st45
		case 66:
			goto tr59
		case 67:
			goto st50
		case 68:
			goto st54
		case 70:
			goto st57
		case 76:
			goto st60
		case 78:
			goto st64
		case 79:
			goto st66
		case 83:
			goto st75
		case 87:
			goto st80
		case 92:
			goto tr45
		case 96:
			goto tr45
		case 97:
			goto st45
		case 98:
			goto tr59
		case 99:
			goto st50
		case 100:
			goto st54
		case 102:
			goto st57
		case 108:
			goto st60
		case 110:
			goto st64
		case 111:
			goto st66
		case 115:
			goto st75
		case 119:
			goto st80
		case 124:
			goto tr48
		case 126:
			goto tr48
		}
		switch {
		case lex.data[(lex.p)] < 36:
			switch {
			case lex.data[(lex.p)] < 9:
				if lex.data[(lex.p)] <= 8 {
					goto tr45
				}
			case lex.data[(lex.p)] > 12:
				if 14 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 35 {
					goto tr45
				}
			default:
				goto tr46
			}
		case lex.data[(lex.p)] > 44:
			switch {
			case lex.data[(lex.p)] < 58:
				if 49 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
					goto tr28
				}
			case lex.data[(lex.p)] > 64:
				switch {
				case lex.data[(lex.p)] > 94:
					if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
						goto tr45
					}
				case lex.data[(lex.p)] >= 91:
					goto tr48
				}
			default:
				goto tr48
			}
		default:
			goto tr48
		}
		goto tr62
	tr46:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st25
	tr72:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:47

		goto st25
	st25:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof25
		}
	st_case_25:
//line pkg/querylang/internal/query/scanner.go:979
		switch lex.data[(lex.p)] {
		case 10:
			goto tr1
		case 13:
			goto tr70
		case 32:
			goto tr46
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr46
		}
		goto tr69
	tr1:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st26
	tr73:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st26
	st26:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof26
		}
	st_case_26:
//line pkg/querylang/internal/query/scanner.go:1029
		switch lex.data[(lex.p)] {
		case 10:
			goto tr73
		case 13:
			goto tr74
		case 32:
			goto tr72
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr72
		}
		goto tr71
	tr70:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st1
	tr74:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st1
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
//line pkg/querylang/internal/query/scanner.go:1073
		if lex.data[(lex.p)] == 10 {
			goto tr1
		}
		goto tr0
	tr47:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st27
	st27:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof27
		}
	st_case_27:
//line pkg/querylang/internal/query/scanner.go:1095
		if lex.data[(lex.p)] == 10 {
			goto tr1
		}
		goto tr75
	tr49:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:231
		lex.act = 33
		goto st28
	st28:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof28
		}
	st_case_28:
//line pkg/querylang/internal/query/scanner.go:1112
		switch lex.data[(lex.p)] {
		case 10:
			goto tr4
		case 13:
			goto tr4
		case 34:
			goto tr5
		case 36:
			goto st3
		case 92:
			goto st4
		case 123:
			goto st5
		}
		goto st2
	tr4:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st2
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
//line pkg/querylang/internal/query/scanner.go:1145
		switch lex.data[(lex.p)] {
		case 10:
			goto tr4
		case 13:
			goto tr4
		case 34:
			goto tr5
		case 36:
			goto st3
		case 92:
			goto st4
		case 123:
			goto st5
		}
		goto st2
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr4
		case 13:
			goto tr4
		case 34:
			goto tr5
		case 36:
			goto st3
		case 92:
			goto st4
		case 96:
			goto st2
		}
		switch {
		case lex.data[(lex.p)] < 91:
			if lex.data[(lex.p)] <= 64 {
				goto st2
			}
		case lex.data[(lex.p)] > 94:
			if 124 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
				goto st2
			}
		default:
			goto st2
		}
		goto tr2
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr4
		case 13:
			goto tr4
		}
		goto st2
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr4
		case 13:
			goto tr4
		case 34:
			goto tr5
		case 36:
			goto tr2
		case 92:
			goto st4
		}
		goto st2
	tr50:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st29
	st29:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof29
		}
	st_case_29:
//line pkg/querylang/internal/query/scanner.go:1233
		switch lex.data[(lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr11
		case 39:
			goto tr5
		case 92:
			goto st7
		}
		goto st6
	tr11:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st6
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
//line pkg/querylang/internal/query/scanner.go:1262
		switch lex.data[(lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr11
		case 39:
			goto tr5
		case 92:
			goto st7
		}
		goto st6
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch lex.data[(lex.p)] {
		case 10:
			goto tr11
		case 13:
			goto tr11
		}
		goto st6
	st30:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		if lex.data[(lex.p)] == 45 {
			goto st31
		}
		goto tr76
	tr80:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st31
	tr82:
//line pkg/querylang/internal/query/scanner.rl:47

		goto st31
	tr84:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st31
	st31:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof31
		}
	st_case_31:
//line pkg/querylang/internal/query/scanner.go:1330
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr79
		case 525:
			goto tr80
		}
		if 512 <= _widec && _widec <= 767 {
			goto st31
		}
		goto tr78
	tr79:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st32
	tr83:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st32
	st32:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof32
		}
	st_case_32:
//line pkg/querylang/internal/query/scanner.go:1410
		_widec = int16(lex.data[(lex.p)])
		switch {
		case lex.data[(lex.p)] < 11:
			switch {
			case lex.data[(lex.p)] > 9:
				if 10 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotNewLine() {
						_widec += 256
					}
				}
			default:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotNewLine() {
					_widec += 256
				}
			}
		case lex.data[(lex.p)] > 12:
			switch {
			case lex.data[(lex.p)] > 13:
				if 14 <= lex.data[(lex.p)] {
					_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
					if lex.isNotNewLine() {
						_widec += 256
					}
				}
			case lex.data[(lex.p)] >= 13:
				_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
				if lex.isNotNewLine() {
					_widec += 256
				}
			}
		default:
			_widec = 256 + (int16(lex.data[(lex.p)]) - 0)
			if lex.isNotNewLine() {
				_widec += 256
			}
		}
		switch _widec {
		case 522:
			goto tr83
		case 525:
			goto tr84
		}
		if 512 <= _widec && _widec <= 767 {
			goto tr82
		}
		goto tr81
	st33:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr16
		}
		goto tr85
	tr16:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:129
		lex.act = 3
		goto st34
	st34:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof34
		}
	st_case_34:
//line pkg/querylang/internal/query/scanner.go:1480
		switch lex.data[(lex.p)] {
		case 69:
			goto st8
		case 95:
			goto st10
		case 101:
			goto st8
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr16
		}
		goto tr86
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 43:
			goto st9
		case 45:
			goto st9
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr14
		}
		goto tr2
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr14
		}
		goto tr2
	tr14:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:129
		lex.act = 3
		goto st35
	st35:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof35
		}
	st_case_35:
//line pkg/querylang/internal/query/scanner.go:1529
		if lex.data[(lex.p)] == 95 {
			goto st9
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr14
		}
		goto tr86
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr16
		}
		goto tr15
	tr53:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st36
	st36:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof36
		}
	st_case_36:
//line pkg/querylang/internal/query/scanner.go:1556
		if lex.data[(lex.p)] == 42 {
			goto st11
		}
		goto tr76
	tr20:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st11
	tr22:
//line pkg/querylang/internal/query/scanner.rl:47

		goto st11
	tr24:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st11
	st11:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof11
		}
	st_case_11:
//line pkg/querylang/internal/query/scanner.go:1596
		switch lex.data[(lex.p)] {
		case 10:
			goto tr19
		case 13:
			goto tr20
		case 42:
			goto st13
		}
		goto st11
	tr19:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st12
	tr23:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st12
	st12:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof12
		}
	st_case_12:
//line pkg/querylang/internal/query/scanner.go:1637
		switch lex.data[(lex.p)] {
		case 10:
			goto tr23
		case 13:
			goto tr24
		case 42:
			goto tr25
		}
		goto tr22
	tr25:
//line pkg/querylang/internal/query/scanner.rl:47

		goto st13
	st13:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof13
		}
	st_case_13:
//line pkg/querylang/internal/query/scanner.go:1656
		switch lex.data[(lex.p)] {
		case 10:
			goto tr19
		case 13:
			goto tr20
		case 42:
			goto st13
		case 47:
			goto tr26
		}
		goto st11
	tr54:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:140
		lex.act = 5
		goto st37
	st37:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof37
		}
	st_case_37:
//line pkg/querylang/internal/query/scanner.go:1680
		switch lex.data[(lex.p)] {
		case 46:
			goto tr90
		case 66:
			goto st15
		case 69:
			goto st8
		case 79:
			goto st16
		case 88:
			goto st17
		case 95:
			goto st14
		case 98:
			goto st15
		case 101:
			goto st8
		case 111:
			goto st16
		case 120:
			goto st17
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr28
		}
		goto tr89
	tr90:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:129
		lex.act = 3
		goto st38
	st38:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof38
		}
	st_case_38:
//line pkg/querylang/internal/query/scanner.go:1719
		switch lex.data[(lex.p)] {
		case 69:
			goto st8
		case 101:
			goto st8
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr16
		}
		goto tr86
	tr28:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:140
		lex.act = 5
		goto st39
	st39:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof39
		}
	st_case_39:
//line pkg/querylang/internal/query/scanner.go:1742
		switch lex.data[(lex.p)] {
		case 46:
			goto tr90
		case 69:
			goto st8
		case 95:
			goto st14
		case 101:
			goto st8
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr28
		}
		goto tr89
	st14:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
			goto tr28
		}
		goto tr27
	st15:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr29
		}
		goto tr2
	tr29:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:130
		lex.act = 4
		goto st40
	st40:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof40
		}
	st_case_40:
//line pkg/querylang/internal/query/scanner.go:1787
		if lex.data[(lex.p)] == 95 {
			goto st15
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 49 {
			goto tr29
		}
		goto tr95
	st16:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 55 {
			goto tr30
		}
		goto tr2
	tr30:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:165
		lex.act = 7
		goto st41
	st41:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof41
		}
	st_case_41:
//line pkg/querylang/internal/query/scanner.go:1816
		if lex.data[(lex.p)] == 95 {
			goto st16
		}
		if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 55 {
			goto tr30
		}
		goto tr96
	st17:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr2
	tr31:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:155
		lex.act = 6
		goto st42
	st42:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof42
		}
	st_case_42:
//line pkg/querylang/internal/query/scanner.go:1854
		if lex.data[(lex.p)] == 95 {
			goto st17
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto tr31
			}
		case lex.data[(lex.p)] > 70:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 102 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto tr97
	st43:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		switch lex.data[(lex.p)] {
		case 61:
			goto tr99
		case 62:
			goto tr100
		}
		goto tr98
	st44:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		if lex.data[(lex.p)] == 61 {
			goto tr102
		}
		goto tr101
	st45:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		switch lex.data[(lex.p)] {
		case 78:
			goto st47
		case 83:
			goto st48
		case 96:
			goto tr103
		case 110:
			goto st47
		case 115:
			goto st48
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	tr62:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:229
		lex.act = 32
		goto st46
	tr106:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:187
		lex.act = 18
		goto st46
	tr107:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:194
		lex.act = 25
		goto st46
	tr111:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:190
		lex.act = 21
		goto st46
	tr114:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:193
		lex.act = 24
		goto st46
	tr117:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:184
		lex.act = 15
		goto st46
	tr121:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:191
		lex.act = 22
		goto st46
	tr123:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:189
		lex.act = 20
		goto st46
	tr129:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:192
		lex.act = 23
		goto st46
	tr138:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:183
		lex.act = 14
		goto st46
	tr143:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:185
		lex.act = 16
		goto st46
	tr145:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:186
		lex.act = 17
		goto st46
	st46:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof46
		}
	st_case_46:
//line pkg/querylang/internal/query/scanner.go:2016
		if lex.data[(lex.p)] == 96 {
			goto tr2
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr2
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr2
				}
			case lex.data[(lex.p)] >= 91:
				goto tr2
			}
		default:
			goto tr2
		}
		goto tr62
	st47:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		switch lex.data[(lex.p)] {
		case 68:
			goto tr106
		case 96:
			goto tr103
		case 100:
			goto tr106
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st48:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		switch lex.data[(lex.p)] {
		case 67:
			goto tr107
		case 96:
			goto tr103
		case 99:
			goto tr107
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	tr59:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:229
		lex.act = 32
		goto st49
	st49:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof49
		}
	st_case_49:
//line pkg/querylang/internal/query/scanner.go:2112
		switch lex.data[(lex.p)] {
		case 34:
			goto st2
		case 96:
			goto tr103
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st50:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		switch lex.data[(lex.p)] {
		case 79:
			goto st51
		case 96:
			goto tr103
		case 111:
			goto st51
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st51:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		switch lex.data[(lex.p)] {
		case 85:
			goto st52
		case 96:
			goto tr103
		case 117:
			goto st52
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st52:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		switch lex.data[(lex.p)] {
		case 78:
			goto st53
		case 96:
			goto tr103
		case 110:
			goto st53
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st53:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr111
		case 96:
			goto tr103
		case 116:
			goto tr111
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st54:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		switch lex.data[(lex.p)] {
		case 69:
			goto st55
		case 96:
			goto tr103
		case 101:
			goto st55
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st55:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		switch lex.data[(lex.p)] {
		case 83:
			goto st56
		case 96:
			goto tr103
		case 115:
			goto st56
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st56:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		switch lex.data[(lex.p)] {
		case 67:
			goto tr114
		case 96:
			goto tr103
		case 99:
			goto tr114
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st57:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		switch lex.data[(lex.p)] {
		case 82:
			goto st58
		case 96:
			goto tr103
		case 114:
			goto st58
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st58:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		switch lex.data[(lex.p)] {
		case 79:
			goto st59
		case 96:
			goto tr103
		case 111:
			goto st59
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st59:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		switch lex.data[(lex.p)] {
		case 77:
			goto tr117
		case 96:
			goto tr103
		case 109:
			goto tr117
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st60:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		switch lex.data[(lex.p)] {
		case 73:
			goto st61
		case 96:
			goto tr103
		case 105:
			goto st61
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st61:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		switch lex.data[(lex.p)] {
		case 77:
			goto st62
		case 96:
			goto tr103
		case 109:
			goto st62
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st62:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		switch lex.data[(lex.p)] {
		case 73:
			goto st63
		case 96:
			goto tr103
		case 105:
			goto st63
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st63:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr121
		case 96:
			goto tr103
		case 116:
			goto tr121
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st64:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		switch lex.data[(lex.p)] {
		case 79:
			goto st65
		case 96:
			goto tr103
		case 111:
			goto st65
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st65:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr123
		case 96:
			goto tr103
		case 116:
			goto tr123
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st66:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		switch lex.data[(lex.p)] {
		case 70:
			goto st67
		case 82:
			goto st71
		case 96:
			goto tr103
		case 102:
			goto st67
		case 114:
			goto st71
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st67:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		switch lex.data[(lex.p)] {
		case 70:
			goto st68
		case 96:
			goto tr103
		case 102:
			goto st68
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st68:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		switch lex.data[(lex.p)] {
		case 83:
			goto st69
		case 96:
			goto tr103
		case 115:
			goto st69
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st69:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		switch lex.data[(lex.p)] {
		case 69:
			goto st70
		case 96:
			goto tr103
		case 101:
			goto st70
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st70:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr129
		case 96:
			goto tr103
		case 116:
			goto tr129
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st71:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch lex.data[(lex.p)] {
		case 68:
			goto st72
		case 96:
			goto tr130
		case 100:
			goto st72
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr130
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr130
				}
			case lex.data[(lex.p)] >= 91:
				goto tr130
			}
		default:
			goto tr130
		}
		goto tr62
	st72:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch lex.data[(lex.p)] {
		case 69:
			goto st73
		case 96:
			goto tr103
		case 101:
			goto st73
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st73:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		switch lex.data[(lex.p)] {
		case 82:
			goto tr133
		case 96:
			goto tr103
		case 114:
			goto tr133
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	tr133:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st74
	st74:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof74
		}
	st_case_74:
//line pkg/querylang/internal/query/scanner.go:2895
		switch lex.data[(lex.p)] {
		case 10:
			goto tr34
		case 13:
			goto tr35
		case 32:
			goto st18
		case 96:
			goto tr103
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto st18
				}
			default:
				goto tr103
			}
		case lex.data[(lex.p)] > 47:
			switch {
			case lex.data[(lex.p)] < 91:
				if 58 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 64 {
					goto tr103
				}
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			default:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	tr37:
//line pkg/querylang/internal/query/scanner.rl:47

		goto st18
	st18:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof18
		}
	st_case_18:
//line pkg/querylang/internal/query/scanner.go:2942
		switch lex.data[(lex.p)] {
		case 10:
			goto tr34
		case 13:
			goto tr35
		case 32:
			goto st18
		case 66:
			goto st21
		case 98:
			goto st21
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto st18
		}
		goto tr32
	tr34:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st19
	tr38:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st19
	st19:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof19
		}
	st_case_19:
//line pkg/querylang/internal/query/scanner.go:2990
		switch lex.data[(lex.p)] {
		case 10:
			goto tr38
		case 13:
			goto tr39
		case 32:
			goto tr37
		case 66:
			goto tr40
		case 98:
			goto tr40
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr37
		}
		goto tr32
	tr35:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st20
	tr39:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st20
	st20:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof20
		}
	st_case_20:
//line pkg/querylang/internal/query/scanner.go:3038
		if lex.data[(lex.p)] == 10 {
			goto tr34
		}
		goto tr32
	tr40:
//line pkg/querylang/internal/query/scanner.rl:47

		goto st21
	st21:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof21
		}
	st_case_21:
//line pkg/querylang/internal/query/scanner.go:3052
		switch lex.data[(lex.p)] {
		case 89:
			goto tr41
		case 121:
			goto tr41
		}
		goto tr32
	st75:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		switch lex.data[(lex.p)] {
		case 69:
			goto st76
		case 96:
			goto tr103
		case 101:
			goto st76
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st76:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		switch lex.data[(lex.p)] {
		case 76:
			goto st77
		case 96:
			goto tr103
		case 108:
			goto st77
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st77:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch lex.data[(lex.p)] {
		case 69:
			goto st78
		case 96:
			goto tr103
		case 101:
			goto st78
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st78:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		switch lex.data[(lex.p)] {
		case 67:
			goto st79
		case 96:
			goto tr103
		case 99:
			goto st79
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st79:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		switch lex.data[(lex.p)] {
		case 84:
			goto tr138
		case 96:
			goto tr103
		case 116:
			goto tr138
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st80:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		switch lex.data[(lex.p)] {
		case 72:
			goto st81
		case 73:
			goto st84
		case 96:
			goto tr103
		case 104:
			goto st81
		case 105:
			goto st84
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st81:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		switch lex.data[(lex.p)] {
		case 69:
			goto st82
		case 96:
			goto tr103
		case 101:
			goto st82
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st82:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		switch lex.data[(lex.p)] {
		case 82:
			goto st83
		case 96:
			goto tr103
		case 114:
			goto st83
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st83:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		switch lex.data[(lex.p)] {
		case 69:
			goto tr143
		case 96:
			goto tr103
		case 101:
			goto tr143
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st84:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		switch lex.data[(lex.p)] {
		case 84:
			goto st85
		case 96:
			goto tr103
		case 116:
			goto st85
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	st85:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		switch lex.data[(lex.p)] {
		case 72:
			goto tr145
		case 96:
			goto tr103
		case 104:
			goto tr145
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr103
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr103
				}
			case lex.data[(lex.p)] >= 91:
				goto tr103
			}
		default:
			goto tr103
		}
		goto tr62
	tr42:
//line pkg/querylang/internal/query/scanner.rl:238
		(lex.p) = (lex.te) - 1
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st86
	tr146:
//line pkg/querylang/internal/query/scanner.rl:241
		lex.te = (lex.p) + 1
		{
			lex.ungetCnt(1)
			{
				goto st24
			}
		}
		goto st86
	tr149:
//line pkg/querylang/internal/query/scanner.rl:239
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.T_OBJECT_OPERATOR
			{
				(lex.p)++
				lex.cs = 86
				goto _out
			}
		}
		goto st86
	tr151:
//line pkg/querylang/internal/query/scanner.rl:238
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st86
	tr153:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:238
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.addFreeFloatingToken(tkn, token.T_WHITESPACE, lex.ts, lex.te)
		}
		goto st86
	tr157:
//line pkg/querylang/internal/query/scanner.rl:241
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.ungetCnt(1)
			{
				goto st24
			}
		}
		goto st86
	tr158:
		lex.cs = 86
//line pkg/querylang/internal/query/scanner.rl:240
		lex.te = (lex.p)
		(lex.p)--
		{
			lex.setTokenPos(tkn)
			tok = token.T_STRING
			lex.cs = 24
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st86:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof86
		}
	st_case_86:
//line NONE:1
		lex.ts = (lex.p)

//line pkg/querylang/internal/query/scanner.go:3458
		switch lex.data[(lex.p)] {
		case 10:
			goto tr43
		case 13:
			goto tr148
		case 32:
			goto tr147
		case 46:
			goto tr149
		case 96:
			goto tr146
		}
		switch {
		case lex.data[(lex.p)] < 14:
			switch {
			case lex.data[(lex.p)] > 8:
				if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
					goto tr147
				}
			default:
				goto tr146
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr146
				}
			case lex.data[(lex.p)] >= 91:
				goto tr146
			}
		default:
			goto tr146
		}
		goto st90
	tr147:
//line NONE:1
		lex.te = (lex.p) + 1

		goto st87
	tr154:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:47

		goto st87
	st87:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof87
		}
	st_case_87:
//line pkg/querylang/internal/query/scanner.go:3511
		switch lex.data[(lex.p)] {
		case 10:
			goto tr43
		case 13:
			goto tr152
		case 32:
			goto tr147
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr147
		}
		goto tr151
	tr43:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st88
	tr155:
//line NONE:1
		lex.te = (lex.p) + 1

//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st88
	st88:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof88
		}
	st_case_88:
//line pkg/querylang/internal/query/scanner.go:3561
		switch lex.data[(lex.p)] {
		case 10:
			goto tr155
		case 13:
			goto tr156
		case 32:
			goto tr154
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 12 {
			goto tr154
		}
		goto tr153
	tr152:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st22
	tr156:
//line pkg/querylang/internal/query/scanner.rl:47

//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st22
	st22:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof22
		}
	st_case_22:
//line pkg/querylang/internal/query/scanner.go:3605
		if lex.data[(lex.p)] == 10 {
			goto tr43
		}
		goto tr42
	tr148:
//line pkg/querylang/internal/query/scanner.rl:35

		if lex.data[lex.p] == '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		if lex.data[lex.p] == '\r' && lex.data[lex.p+1] != '\n' {
			lex.newLines.Append(lex.p + 1)
		}

		goto st89
	st89:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof89
		}
	st_case_89:
//line pkg/querylang/internal/query/scanner.go:3627
		if lex.data[(lex.p)] == 10 {
			goto tr43
		}
		goto tr157
	st90:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		if lex.data[(lex.p)] == 96 {
			goto tr158
		}
		switch {
		case lex.data[(lex.p)] < 58:
			if lex.data[(lex.p)] <= 47 {
				goto tr158
			}
		case lex.data[(lex.p)] > 64:
			switch {
			case lex.data[(lex.p)] > 94:
				if 123 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 127 {
					goto tr158
				}
			case lex.data[(lex.p)] >= 91:
				goto tr158
			}
		default:
			goto tr158
		}
		goto st90
	tr159:
		lex.cs = 91
//line pkg/querylang/internal/query/scanner.rl:245
		lex.te = (lex.p) + 1
		{
			lex.setTokenPos(tkn)
			tok = token.ID(int('"'))
			lex.cs = 24
			{
				(lex.p)++
				goto _out
			}
		}
		goto _again
	st91:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof91
		}
	st_case_91:
//line NONE:1
		lex.ts = (lex.p)

//line pkg/querylang/internal/query/scanner.go:3675
		if lex.data[(lex.p)] == 34 {
			goto tr159
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	st_out:
	_test_eof23:
		lex.cs = 23
		goto _test_eof
	_test_eof24:
		lex.cs = 24
		goto _test_eof
	_test_eof25:
		lex.cs = 25
		goto _test_eof
	_test_eof26:
		lex.cs = 26
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof27:
		lex.cs = 27
		goto _test_eof
	_test_eof28:
		lex.cs = 28
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof4:
		lex.cs = 4
		goto _test_eof
	_test_eof5:
		lex.cs = 5
		goto _test_eof
	_test_eof29:
		lex.cs = 29
		goto _test_eof
	_test_eof6:
		lex.cs = 6
		goto _test_eof
	_test_eof7:
		lex.cs = 7
		goto _test_eof
	_test_eof30:
		lex.cs = 30
		goto _test_eof
	_test_eof31:
		lex.cs = 31
		goto _test_eof
	_test_eof32:
		lex.cs = 32
		goto _test_eof
	_test_eof33:
		lex.cs = 33
		goto _test_eof
	_test_eof34:
		lex.cs = 34
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof35:
		lex.cs = 35
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof
	_test_eof36:
		lex.cs = 36
		goto _test_eof
	_test_eof11:
		lex.cs = 11
		goto _test_eof
	_test_eof12:
		lex.cs = 12
		goto _test_eof
	_test_eof13:
		lex.cs = 13
		goto _test_eof
	_test_eof37:
		lex.cs = 37
		goto _test_eof
	_test_eof38:
		lex.cs = 38
		goto _test_eof
	_test_eof39:
		lex.cs = 39
		goto _test_eof
	_test_eof14:
		lex.cs = 14
		goto _test_eof
	_test_eof15:
		lex.cs = 15
		goto _test_eof
	_test_eof40:
		lex.cs = 40
		goto _test_eof
	_test_eof16:
		lex.cs = 16
		goto _test_eof
	_test_eof41:
		lex.cs = 41
		goto _test_eof
	_test_eof17:
		lex.cs = 17
		goto _test_eof
	_test_eof42:
		lex.cs = 42
		goto _test_eof
	_test_eof43:
		lex.cs = 43
		goto _test_eof
	_test_eof44:
		lex.cs = 44
		goto _test_eof
	_test_eof45:
		lex.cs = 45
		goto _test_eof
	_test_eof46:
		lex.cs = 46
		goto _test_eof
	_test_eof47:
		lex.cs = 47
		goto _test_eof
	_test_eof48:
		lex.cs = 48
		goto _test_eof
	_test_eof49:
		lex.cs = 49
		goto _test_eof
	_test_eof50:
		lex.cs = 50
		goto _test_eof
	_test_eof51:
		lex.cs = 51
		goto _test_eof
	_test_eof52:
		lex.cs = 52
		goto _test_eof
	_test_eof53:
		lex.cs = 53
		goto _test_eof
	_test_eof54:
		lex.cs = 54
		goto _test_eof
	_test_eof55:
		lex.cs = 55
		goto _test_eof
	_test_eof56:
		lex.cs = 56
		goto _test_eof
	_test_eof57:
		lex.cs = 57
		goto _test_eof
	_test_eof58:
		lex.cs = 58
		goto _test_eof
	_test_eof59:
		lex.cs = 59
		goto _test_eof
	_test_eof60:
		lex.cs = 60
		goto _test_eof
	_test_eof61:
		lex.cs = 61
		goto _test_eof
	_test_eof62:
		lex.cs = 62
		goto _test_eof
	_test_eof63:
		lex.cs = 63
		goto _test_eof
	_test_eof64:
		lex.cs = 64
		goto _test_eof
	_test_eof65:
		lex.cs = 65
		goto _test_eof
	_test_eof66:
		lex.cs = 66
		goto _test_eof
	_test_eof67:
		lex.cs = 67
		goto _test_eof
	_test_eof68:
		lex.cs = 68
		goto _test_eof
	_test_eof69:
		lex.cs = 69
		goto _test_eof
	_test_eof70:
		lex.cs = 70
		goto _test_eof
	_test_eof71:
		lex.cs = 71
		goto _test_eof
	_test_eof72:
		lex.cs = 72
		goto _test_eof
	_test_eof73:
		lex.cs = 73
		goto _test_eof
	_test_eof74:
		lex.cs = 74
		goto _test_eof
	_test_eof18:
		lex.cs = 18
		goto _test_eof
	_test_eof19:
		lex.cs = 19
		goto _test_eof
	_test_eof20:
		lex.cs = 20
		goto _test_eof
	_test_eof21:
		lex.cs = 21
		goto _test_eof
	_test_eof75:
		lex.cs = 75
		goto _test_eof
	_test_eof76:
		lex.cs = 76
		goto _test_eof
	_test_eof77:
		lex.cs = 77
		goto _test_eof
	_test_eof78:
		lex.cs = 78
		goto _test_eof
	_test_eof79:
		lex.cs = 79
		goto _test_eof
	_test_eof80:
		lex.cs = 80
		goto _test_eof
	_test_eof81:
		lex.cs = 81
		goto _test_eof
	_test_eof82:
		lex.cs = 82
		goto _test_eof
	_test_eof83:
		lex.cs = 83
		goto _test_eof
	_test_eof84:
		lex.cs = 84
		goto _test_eof
	_test_eof85:
		lex.cs = 85
		goto _test_eof
	_test_eof86:
		lex.cs = 86
		goto _test_eof
	_test_eof87:
		lex.cs = 87
		goto _test_eof
	_test_eof88:
		lex.cs = 88
		goto _test_eof
	_test_eof22:
		lex.cs = 22
		goto _test_eof
	_test_eof89:
		lex.cs = 89
		goto _test_eof
	_test_eof90:
		lex.cs = 90
		goto _test_eof
	_test_eof91:
		lex.cs = 91
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 25:
				goto tr69
			case 26:
				goto tr71
			case 1:
				goto tr0
			case 27:
				goto tr75
			case 28:
				goto tr75
			case 2:
				goto tr2
			case 3:
				goto tr2
			case 4:
				goto tr2
			case 5:
				goto tr2
			case 29:
				goto tr75
			case 6:
				goto tr9
			case 7:
				goto tr9
			case 30:
				goto tr76
			case 31:
				goto tr78
			case 32:
				goto tr81
			case 33:
				goto tr85
			case 34:
				goto tr86
			case 8:
				goto tr2
			case 9:
				goto tr2
			case 35:
				goto tr86
			case 10:
				goto tr15
			case 36:
				goto tr76
			case 11:
				goto tr17
			case 12:
				goto tr17
			case 13:
				goto tr17
			case 37:
				goto tr89
			case 38:
				goto tr86
			case 39:
				goto tr89
			case 14:
				goto tr27
			case 15:
				goto tr2
			case 40:
				goto tr95
			case 16:
				goto tr2
			case 41:
				goto tr96
			case 17:
				goto tr2
			case 42:
				goto tr97
			case 43:
				goto tr98
			case 44:
				goto tr101
			case 45:
				goto tr103
			case 46:
				goto tr2
			case 47:
				goto tr103
			case 48:
				goto tr103
			case 49:
				goto tr103
			case 50:
				goto tr103
			case 51:
				goto tr103
			case 52:
				goto tr103
			case 53:
				goto tr103
			case 54:
				goto tr103
			case 55:
				goto tr103
			case 56:
				goto tr103
			case 57:
				goto tr103
			case 58:
				goto tr103
			case 59:
				goto tr103
			case 60:
				goto tr103
			case 61:
				goto tr103
			case 62:
				goto tr103
			case 63:
				goto tr103
			case 64:
				goto tr103
			case 65:
				goto tr103
			case 66:
				goto tr103
			case 67:
				goto tr103
			case 68:
				goto tr103
			case 69:
				goto tr103
			case 70:
				goto tr103
			case 71:
				goto tr130
			case 72:
				goto tr103
			case 73:
				goto tr103
			case 74:
				goto tr103
			case 18:
				goto tr32
			case 19:
				goto tr32
			case 20:
				goto tr32
			case 21:
				goto tr32
			case 75:
				goto tr103
			case 76:
				goto tr103
			case 77:
				goto tr103
			case 78:
				goto tr103
			case 79:
				goto tr103
			case 80:
				goto tr103
			case 81:
				goto tr103
			case 82:
				goto tr103
			case 83:
				goto tr103
			case 84:
				goto tr103
			case 85:
				goto tr103
			case 87:
				goto tr151
			case 88:
				goto tr153
			case 22:
				goto tr42
			case 89:
				goto tr157
			case 90:
				goto tr158
			}
		}

	_out:
		{
		}
	}

//line pkg/querylang/internal/query/scanner.rl:251

	tkn.Value = lex.data[lex.ts:lex.te]
	tkn.ID = token.ID(tok)

	return tkn
}
