package utils

const (
	ESC			string = "\033["
	// text color
	BLACK		string = ESC+"30m"
	RED			string = ESC+"31m"
	GREEN		string = ESC+"32m"
	BROWN		string = ESC+"33m"
	BLUE		string = ESC+"34m"
	PURPLE		string = ESC+"35m"
	CYAN		string = ESC+"36m"
	GRAY		string = ESC+"37m"

	// background color
	BKG_BLACK	string = ESC+"40m"
	BKG_RED		string = ESC+"41m"
	BKG_GREEN	string = ESC+"42m"
	BKG_BROWN	string = ESC+"43m"
	BKG_BLUE	string = ESC+"44m"
	BKG_PURPLE	string = ESC+"45m"
	BKG_CYAN	string = ESC+"46m"
	BKG_GRAY	string = ESC+"47m"

	// font
	BOLD		string = ESC+"1m"
	UNDERLINE	string = ESC+"4m"
	BLINK		string = ESC+"5m"
	INVERSION	string = ESC+"7m"

	// end
	RESET		string = ESC+"0m"

	// other
	QUOTES		string = ESC+"4b"
	PUSH_POS	string = ESC+"s"
	POP_POS		string = ESC+"u"
	CLR_BGN		string = ESC+"1K"
	CLR_STR		string = ESC+"2K"
	CLR_SCR		string = ESC+"2J"

	SET_HOME	string = ESC+"H"
	SET_POS		string = ESC+"%d;%dH"
	UP_POS		string = ESC+"%dA"
	DWN_POS		string = ESC+"%dB"
	FRWRD_POS	string = ESC+"%dC"
	RVRS_POS	string = ESC+"%dD"
	CURS_DIS	string = ESC+"?25l"
	CURS_EN		string = ESC+"?25h"
)

/*
ANSI Escape Sequences:
https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
*/
