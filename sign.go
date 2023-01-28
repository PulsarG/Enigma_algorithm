package enigma

var SignsArray = [162]string{
	"а",
	"б",
	"в",
	"г",
	"д",
	"е",
	"ё",
	"ж",
	"з",
	"и",
	"й",
	"к",
	"л",
	"м",
	"н",
	"о",
	"п",
	"р",
	"с",
	"т",
	"у",
	"ф",
	"х",
	"ц",
	"ч",
	"ш",
	"щ",
	"ъ",
	"ы",
	"ь",
	"э",
	"ю",
	"я",
	"А",
	"Б",
	"В",
	"Г",
	"Д",
	"Е",
	"Ё",
	"Ж",
	"З",
	"И",
	"Й",
	"К",
	"Л",
	"М",
	"Н",
	"О",
	"П",
	"Р",
	"С",
	"Т",
	"У",
	"Ф",
	"Х",
	"Ц",
	"Ч",
	"Ш",
	"Щ",
	"Ъ",
	"Ы",
	"Ь",
	"Э",
	"Ю",
	"Я",
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
	"g",
	"h",
	"i",
	"j",
	"k",
	"l",
	"m",
	"n",
	"o",
	"p",
	"q",
	"r",
	"s",
	"t",
	"u",
	"v",
	"w",
	"x",
	"y",
	"z",
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"0",
	"@",
	"#",
	"$",
	"%",
	"^",
	"&",
	"*",
	"_",
	"+",
	"=",
	",",
	".",
	"!",
	"?",
	"-",
	"(",
	")",
	"№",
	";",
	":",
	"/",
	"|",
	"`",
	"~",
	"[",
	"]",
	"{",
	"}",
	"<",
	">",
	"'",
	" ",
	"—",
	"\n",
}

var Mirror = [162]int{
	161,
	70,
	14,
	92,
	156,
	79,
	118,
	160,
	110,
	152,
	74,
	136,
	159,
	72,
	2,
	114,
	54,
	41,
	87,
	51,
	132,
	82,
	127,
	128,
	138,
	153,
	75,
	99,
	145,
	64,
	154,
	44,
	77,
	117,
	104,
	143,
	129,
	135,
	122,
	158,
	101,
	17,
	80,
	108,
	31,
	115,
	131,
	94,
	85,
	126,
	148,
	19,
	86,
	96,
	16,
	112,
	66,
	106,
	142,
	134,
	130,
	93,
	103,
	140,
	29,
	137,
	56,
	133,
	109,
	151,
	1,
	113,
	13,
	146,
	10,
	26,
	111,
	32,
	141,
	5,
	42,
	119,
	21,
	107,
	157,
	48,
	52,
	18,
	98,
	125,
	102,
	139,
	3,
	61,
	47,
	155,
	53,
	124,
	88,
	27,
	149,
	40,
	90,
	62,
	34,
	150,
	57,
	83,
	43,
	68,
	8,
	76,
	55,
	71,
	15,
	45,
	144,
	33,
	6,
	81,
	147,
	123,
	38,
	121,
	97,
	89,
	49,
	22,
	23,
	36,
	60,
	46,
	20,
	67,
	59,
	37,
	11,
	65,
	24,
	91,
	63,
	78,
	58,
	35,
	116,
	28,
	73,
	120,
	50,
	100,
	105,
	69,
	9,
	25,
	30,
	95,
	4,
	84,
	39,
	12,
	7,
	0,
}