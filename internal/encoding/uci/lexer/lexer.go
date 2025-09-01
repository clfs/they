// Package lexer implements a lexer for the UCI protocol.
package lexer

type TokenType int

const (
	// Literals:
	TokenType_LiteralString TokenType = iota
	TokenType_LiteralNumber

	// Keywords:
	TokenType_BestMove
	TokenType_BestMove_Ponder
	TokenType_CopyProtection
	TokenType_CopyProtection_Checking
	TokenType_CopyProtection_Error
	TokenType_CopyProtection_OK
	TokenType_Debug
	TokenType_Debug_Off
	TokenType_Debug_On
	TokenType_Go
	TokenType_Go_BInc
	TokenType_Go_BTime
	TokenType_Go_Depth
	TokenType_Go_Infinite
	TokenType_Go_Mate
	TokenType_Go_MovesToGo
	TokenType_Go_MoveTime
	TokenType_Go_Nodes
	TokenType_Go_Ponder
	TokenType_Go_SearchMoves
	TokenType_Go_WInc
	TokenType_Go_WTime
	TokenType_ID
	TokenType_IsReady
	TokenType_PonderHit
	TokenType_Position
	TokenType_Position_FEN
	TokenType_Position_Moves
	TokenType_Position_Startpos
	TokenType_Quit
	TokenType_ReadyOK
	TokenType_Register
	TokenType_Register_Code
	TokenType_Register_Later
	TokenType_Register_Name
	TokenType_Registration
	TokenType_SetOption
	TokenType_SetOption_Name
	TokenType_SetOption_Value
	TokenType_Stop
	TokenType_UCI
	TokenType_UCINewGame
	TokenType_UCIOK

	// Special:
	TokenType_EOL
	TokenType_EOF
)
