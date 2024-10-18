package badge

var materials = map[string]string{
	"алюминий": "aluminum",
	"пластик":  "plastic",
	"жесть":    "tin",
	"стекло":   "glass",
	"серебро":  "silver",
	"золото":   "gold",
	"латунь":   "brass",
	"медь":     "copper",
}

var variants = map[string]string{
	"тяжелый": "heavy",
	"легкий":  "light",
}

var manufacturingTypes = map[string]string{
	"литый/штамп": "cast/stamped",
	"накладной":   "overlay",
	"закатной":    "embossed",
	"фототипия":   "photogravure",
	"ситалл":      "sitall",
}

var fastenings = map[string]string{
	"булавка": "pin",
	"игла":    "needle",
	"цанга":   "clamp",
	"винт":    "screw",
}

var coverages = map[string]string{
	"холодная эмаль": "cold enamel",
	"горячая эмаль":  "hot enamel",
	"краска":         "paint",
}
