// goherokuname generates Heroku-like random names. Such "block-black-0231".
//
// It supports both decimal and hexadecimal suffixes. It supports customized
// rendering, in which the delimiter, the length of suffix and the acceptable
// characters for suffix are tweakable.
package goherokuname

import "math/rand"

var nouns = [...]string{
	"art", "atom", "band", "bar", "base", "bird", "block", "boat", "bonus",
	"bread", "breeze", "brook", "bush", "butterfly", "cake", "cell",
	"cherry", "cloud", "coke", "credit", "darkness", "dawn", "dew", "disk",
	"dream", "dust", "fashion", "feather", "field", "fire", "firefly",
	"flower", "fog", "forest", "frog", "frost", "glade", "glitter", "grass",
	"hall", "hat", "haze", "heart", "hill", "hola", "king", "kiwi", "lab",
	"lake", "leaf", "limit", "lion", "math", "meadow", "mode", "moon",
	"morning", "mountain", "mouse", "mud", "night", "paper", "penguin",
	"pine", "poetry", "pond", "queen", "rain", "recipe", "resonance",
	"rice", "river", "salad", "scene", "sea", "shadow", "shape", "silence",
	"sky", "smoke", "snow", "snowflake", "sound", "star", "sun", "sun",
	"sunset", "surf", "term", "thunder", "tiger", "toast", "tooth", "tree",
	"truth", "union", "unit", "violet", "voice", "water", "water",
	"waterfall", "wave", "wildflower", "wind", "wood",
}

var adjectives = []string{
	"aged", "ancient", "autumn", "billowing", "bitter", "black", "blue",
	"bold", "broad", "broken", "calm", "cold", "cool", "crimson", "curly",
	"damp", "dark", "dawn", "delicate", "divine", "dry", "dry", "empty",
	"falling", "fancy", "flat", "floral", "fragrant", "free", "frosty",
	"gentle", "green", "hidden", "holy", "hushy", "icy", "jolly", "late",
	"lingering", "little", "lively", "long", "lucky", "misty", "morning",
	"muddy", "mute", "nameless", "noisy", "odd", "old", "orange", "patient",
	"plain", "polished", "proud", "purple", "quiet", "rapid", "raspy",
	"red", "restless", "rough", "round", "royal", "shinny", "shrill", "shy",
	"silent", "small", "snowy", "soft", "solitary", "sparkling", "spring",
	"square", "steep", "still", "summer", "super", "sweet", "throbbing",
	"tight", "tiny", "twilight", "wandering", "weathered", "white", "wild",
	"winter", "wispy", "withered", "yellow", "young",
}

// HaikunateCustom generates a Heroku-like random name in which the delimiter,
// length and acceptable characters for suffix are tweakable.
func HaikunateCustom(delimiter string, toklen int, tokchars string) string {
	noun := nouns[rand.Intn(len(nouns))]
	adjective := adjectives[rand.Intn(len(adjectives))]

	token := ""
	for i := 0; i < toklen; i++ {
		token += string(tokchars[rand.Intn(len(tokchars))])
	}

	return adjective + delimiter + noun + delimiter + token
}

// Haikunate generate standard Heroku-like random name, with "-" as delimiter,
// decimal with 4 digits.
func Haikunate() string {
	return HaikunateCustom("-", 4, "0123456789")
}

// HaikunateHex generate standard Heroku-like random name, with "-" as
// delimiter, hexadecimal with 4 digits.
func HaikunateHex() string {
	return HaikunateCustom("-", 4, "0123456789abcdef")
}
