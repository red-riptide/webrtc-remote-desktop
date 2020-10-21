package config

var	(
	// RobotGoJSKeyMap return keyboard mapping between JS codes and Robotgo keys
	RobotGoJSKeyMap map[float64]string
)

func init() {
	RobotGoJSKeyMap = make(map[float64]string)

	RobotGoJSKeyMap[1] = ""
	RobotGoJSKeyMap[2] = ""
	RobotGoJSKeyMap[3] = ""
	RobotGoJSKeyMap[4] = ""
	RobotGoJSKeyMap[5] = ""
	RobotGoJSKeyMap[6] = ""
	RobotGoJSKeyMap[7] = ""
	RobotGoJSKeyMap[8] = "backspace"
	RobotGoJSKeyMap[9] = "tab"
	RobotGoJSKeyMap[10] = ""
	RobotGoJSKeyMap[11] = ""
	RobotGoJSKeyMap[12] = ""
	RobotGoJSKeyMap[13] = "enter"
	RobotGoJSKeyMap[14] = ""
	RobotGoJSKeyMap[15] = ""
	RobotGoJSKeyMap[16] = "shift"
	RobotGoJSKeyMap[17] = "ctrl"
	RobotGoJSKeyMap[18] = "alt"
	RobotGoJSKeyMap[19] = ""
	RobotGoJSKeyMap[20] = "capslock"
	RobotGoJSKeyMap[21] = ""
	RobotGoJSKeyMap[22] = ""
	RobotGoJSKeyMap[23] = ""
	RobotGoJSKeyMap[24] = ""
	RobotGoJSKeyMap[25] = ""
	RobotGoJSKeyMap[26] = ""
	RobotGoJSKeyMap[27] = "escape"
	RobotGoJSKeyMap[28] = ""
	RobotGoJSKeyMap[29] = ""
	RobotGoJSKeyMap[30] = ""
	RobotGoJSKeyMap[31] = ""
	RobotGoJSKeyMap[32] = "space"
	RobotGoJSKeyMap[33] = "pageup"
	RobotGoJSKeyMap[34] = "pagedown"
	RobotGoJSKeyMap[35] = "end"
	RobotGoJSKeyMap[36] = "home"
	RobotGoJSKeyMap[37] = "left"
	RobotGoJSKeyMap[38] = "up"
	RobotGoJSKeyMap[39] = "right"
	RobotGoJSKeyMap[40] = "down"
	RobotGoJSKeyMap[41] = ""
	RobotGoJSKeyMap[42] = ""
	RobotGoJSKeyMap[43] = ""
	RobotGoJSKeyMap[44] = "printscreen"
	RobotGoJSKeyMap[45] = "insert"
	RobotGoJSKeyMap[46] = "delete"
	RobotGoJSKeyMap[47] = ""
	RobotGoJSKeyMap[48] = "0"
	RobotGoJSKeyMap[49] = "1"
	RobotGoJSKeyMap[50] = "2"
	RobotGoJSKeyMap[51] = "3"
	RobotGoJSKeyMap[52] = "4"
	RobotGoJSKeyMap[53] = "5"
	RobotGoJSKeyMap[54] = "6"
	RobotGoJSKeyMap[55] = "7"
	RobotGoJSKeyMap[56] = "8"
	RobotGoJSKeyMap[57] = "9"
	RobotGoJSKeyMap[58] = ""
	RobotGoJSKeyMap[59] = ""
	RobotGoJSKeyMap[60] = ""
	RobotGoJSKeyMap[61] = ""
	RobotGoJSKeyMap[62] = ""
	RobotGoJSKeyMap[63] = ""
	RobotGoJSKeyMap[64] = ""
	RobotGoJSKeyMap[65] = "a"
	RobotGoJSKeyMap[66] = "b"
	RobotGoJSKeyMap[67] = "c"
	RobotGoJSKeyMap[68] = "d"
	RobotGoJSKeyMap[69] = "e"
	RobotGoJSKeyMap[70] = "f"
	RobotGoJSKeyMap[71] = "g"
	RobotGoJSKeyMap[72] = "h"
	RobotGoJSKeyMap[73] = "i"
	RobotGoJSKeyMap[74] = "j"
	RobotGoJSKeyMap[75] = "k"
	RobotGoJSKeyMap[76] = "l"
	RobotGoJSKeyMap[77] = "m"
	RobotGoJSKeyMap[78] = "n"
	RobotGoJSKeyMap[79] = "o"
	RobotGoJSKeyMap[80] = "p"
	RobotGoJSKeyMap[81] = "q"
	RobotGoJSKeyMap[82] = "r"
	RobotGoJSKeyMap[83] = "s"
	RobotGoJSKeyMap[84] = "t"
	RobotGoJSKeyMap[85] = "u"
	RobotGoJSKeyMap[86] = "v"
	RobotGoJSKeyMap[87] = "w"
	RobotGoJSKeyMap[88] = "x"
	RobotGoJSKeyMap[89] = "y"
	RobotGoJSKeyMap[90] = "z"
	RobotGoJSKeyMap[91] = "lcmd"
	RobotGoJSKeyMap[92] = "rcmd"
	RobotGoJSKeyMap[93] = ""
	RobotGoJSKeyMap[94] = ""
	RobotGoJSKeyMap[95] = ""
	RobotGoJSKeyMap[96] = "num0"
	RobotGoJSKeyMap[97] = "num1"
	RobotGoJSKeyMap[98] = "num2"
	RobotGoJSKeyMap[99] = "num3"
	RobotGoJSKeyMap[100] = "num4"
	RobotGoJSKeyMap[101] = "num5"
	RobotGoJSKeyMap[102] = "num6"
	RobotGoJSKeyMap[103] = "num7"
	RobotGoJSKeyMap[104] = "num8"
	RobotGoJSKeyMap[105] = "num9"
	RobotGoJSKeyMap[106] = "num*"
	RobotGoJSKeyMap[107] = "num+"
	RobotGoJSKeyMap[108] = ""
	RobotGoJSKeyMap[109] = "num-"
	RobotGoJSKeyMap[110] = "num."
	RobotGoJSKeyMap[111] = "num/"
	RobotGoJSKeyMap[112] = "f1"
	RobotGoJSKeyMap[113] = "f2"
	RobotGoJSKeyMap[114] = "f3"
	RobotGoJSKeyMap[115] = "f4"
	RobotGoJSKeyMap[116] = "f5"
	RobotGoJSKeyMap[117] = "f6"
	RobotGoJSKeyMap[118] = "f7"
	RobotGoJSKeyMap[119] = "f8"
	RobotGoJSKeyMap[120] = "f9"
	RobotGoJSKeyMap[121] = "f10"
	RobotGoJSKeyMap[122] = "f11"
	RobotGoJSKeyMap[123] = "f12"
	RobotGoJSKeyMap[124] = ""
	RobotGoJSKeyMap[125] = ""
	RobotGoJSKeyMap[126] = ""
	RobotGoJSKeyMap[127] = ""
	RobotGoJSKeyMap[128] = ""
	RobotGoJSKeyMap[129] = ""
	RobotGoJSKeyMap[130] = ""
	RobotGoJSKeyMap[131] = ""
	RobotGoJSKeyMap[132] = ""
	RobotGoJSKeyMap[133] = ""
	RobotGoJSKeyMap[134] = ""
	RobotGoJSKeyMap[135] = ""
	RobotGoJSKeyMap[136] = ""
	RobotGoJSKeyMap[137] = ""
	RobotGoJSKeyMap[138] = ""
	RobotGoJSKeyMap[139] = ""
	RobotGoJSKeyMap[140] = ""
	RobotGoJSKeyMap[141] = ""
	RobotGoJSKeyMap[142] = ""
	RobotGoJSKeyMap[143] = ""
	RobotGoJSKeyMap[144] = "num_lock"
	RobotGoJSKeyMap[145] = ""
	RobotGoJSKeyMap[146] = ""
	RobotGoJSKeyMap[147] = ""
	RobotGoJSKeyMap[148] = ""
	RobotGoJSKeyMap[149] = ""
	RobotGoJSKeyMap[150] = ""
	RobotGoJSKeyMap[151] = ""
	RobotGoJSKeyMap[152] = ""
	RobotGoJSKeyMap[153] = ""
	RobotGoJSKeyMap[154] = ""
	RobotGoJSKeyMap[155] = ""
	RobotGoJSKeyMap[156] = ""
	RobotGoJSKeyMap[157] = ""
	RobotGoJSKeyMap[158] = ""
	RobotGoJSKeyMap[159] = ""
	RobotGoJSKeyMap[160] = ""
	RobotGoJSKeyMap[161] = ""
	RobotGoJSKeyMap[162] = ""
	RobotGoJSKeyMap[163] = ""
	RobotGoJSKeyMap[164] = ""
	RobotGoJSKeyMap[165] = ""
	RobotGoJSKeyMap[166] = ""
	RobotGoJSKeyMap[167] = ""
	RobotGoJSKeyMap[168] = ""
	RobotGoJSKeyMap[169] = ""
	RobotGoJSKeyMap[170] = ""
	RobotGoJSKeyMap[171] = ""
	RobotGoJSKeyMap[172] = ""
	RobotGoJSKeyMap[173] = ""
	RobotGoJSKeyMap[174] = ""
	RobotGoJSKeyMap[175] = ""
	RobotGoJSKeyMap[176] = ""
	RobotGoJSKeyMap[177] = ""
	RobotGoJSKeyMap[178] = ""
	RobotGoJSKeyMap[179] = ""
	RobotGoJSKeyMap[180] = ""
	RobotGoJSKeyMap[181] = ""
	RobotGoJSKeyMap[182] = ""
	RobotGoJSKeyMap[183] = ""
	RobotGoJSKeyMap[184] = ""
	RobotGoJSKeyMap[185] = ""
	RobotGoJSKeyMap[186] = ""
	RobotGoJSKeyMap[187] = ""
	RobotGoJSKeyMap[188] = ""
	RobotGoJSKeyMap[189] = ""
	RobotGoJSKeyMap[190] = ""
	RobotGoJSKeyMap[191] = ""
	RobotGoJSKeyMap[192] = ""
	RobotGoJSKeyMap[193] = ""
	RobotGoJSKeyMap[194] = ""
	RobotGoJSKeyMap[195] = ""
	RobotGoJSKeyMap[196] = ""
	RobotGoJSKeyMap[197] = ""
	RobotGoJSKeyMap[198] = ""
	RobotGoJSKeyMap[199] = ""
	RobotGoJSKeyMap[200] = ""

	// Random Keys
	RobotGoJSKeyMap[201] = ""
	RobotGoJSKeyMap[202] = ""
	RobotGoJSKeyMap[203] = ""
	RobotGoJSKeyMap[204] = ""
	RobotGoJSKeyMap[205] = ""
	RobotGoJSKeyMap[206] = ""
	RobotGoJSKeyMap[207] = ""
	RobotGoJSKeyMap[208] = ""
	RobotGoJSKeyMap[209] = ""
	RobotGoJSKeyMap[210] = ""

}