package davilib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

// ABECEDARIO ESPANOL
const AbcSpanish = "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ" // We start with Capital letter so that it isnt private to this package

// ABECEDARIO ESPANOL SIN J
const AbcSpanishNoJ = "ABCDEFGHIKLMNOPQRSTUVWXYZ"

// ABECEDARIO SIN Ñ
const AbcEnglish = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ABECEDARIO SIN Ñ SIN J
const AbcEnglishNoJ = "ABCDEFGHIKLMNOPQRSTUVWXYZ"

// Generalize Struct for letter and frequency data in Castillian Spanish
type char struct {
	letter rune
	freq   float32
	cons   bool // Is consonant?
}

// Spanish Data: https://es.wikipedia.org/wiki/Frecuencia_de_aparici%C3%B3n_de_letras
// 0.45 f = Vowels (a,e,i,o,u)
// 0.55 f = Cons. Las más frecuentes son S R N D L C con f = 0.37
// De Mayor a Menor Frecuencia: E A O S R N I D L C T U M P B G V Y Q H F Z J Ñ X K W
// Otros Datos Importantes:
// -  la frecuencia de las partículas "que", "el", "se", "me", etc. hace que la "e" sea más frecuente.
// El estilo narrativo. Si hay muchos verbos en infinitivo, habrá muchas "R".

// Analisis por Silabas
// Segun: http://www.lllf.uam.es/ESP/Publicaciones/LLI-UAM-4JTH.pdf
// La distribucion de frecuencias por tipos de silaba, segun la tabla 3, es:
// CV : f = 0.5135
// CVC: f = 0.1803
// V: f = 0.1075
// VC: f = 0.0860
// CVV: f = 0.0337
// CVVC: f = 0.0331
// CCV: f = 0.0296
// CCVC: f = 0.0088

var FreqTable = []char{
	char{
		letter: 'A',
		freq:   0.1253,
		cons:   false,
	},
	char{
		letter: 'B',
		freq:   0.0142,
		cons:   true,
	},
	char{
		letter: 'C',
		freq:   0.0468,
		cons:   true,
	},
	char{
		letter: 'D',
		freq:   0.0586,
		cons:   true,
	},
	char{
		letter: 'E',
		freq:   0.1368,
		cons:   false,
	},
	char{
		letter: 'F',
		freq:   0.0069,
		cons:   true,
	},
	char{
		letter: 'G',
		freq:   0.0101,
		cons:   true,
	},
	char{
		letter: 'H',
		freq:   0.0070,
		cons:   true,
	},
	char{
		letter: 'I',
		freq:   0.0625,
		cons:   false,
	},
	char{
		letter: 'J',
		freq:   0.0044,
		cons:   true,
	},
	char{
		letter: 'K',
		freq:   0.0002,
		cons:   true,
	},
	char{
		letter: 'L',
		freq:   0.0497,
		cons:   true,
	},
	char{
		letter: 'M',
		freq:   0.0315,
		cons:   true,
	},
	char{
		letter: 'N',
		freq:   0.0671,
		cons:   true,
	},
	char{
		letter: 'Ñ',
		freq:   0.0031,
		cons:   true,
	},
	char{
		letter: 'O',
		freq:   0.0868,
		cons:   false,
	},
	char{
		letter: 'P',
		freq:   0.0251,
		cons:   true,
	},
	char{
		letter: 'Q',
		freq:   0.0088,
		cons:   true,
	},
	char{
		letter: 'R',
		freq:   0.0687,
		cons:   true,
	},
	char{
		letter: 'S',
		freq:   0.0798,
		cons:   true,
	},
	char{
		letter: 'T',
		freq:   0.0463,
		cons:   true,
	},
	char{
		letter: 'U',
		freq:   0.0393,
		cons:   false,
	},
	char{
		letter: 'V',
		freq:   0.0090,
		cons:   true,
	},
	char{
		letter: 'W',
		freq:   0.0001,
		cons:   true,
	},
	char{
		letter: 'X',
		freq:   0.0022,
		cons:   true,
	},
	char{
		letter: 'Y',
		freq:   0.0090,
		cons:   true,
	},
	char{
		letter: 'Z',
		freq:   0.0052,
		cons:   true,
	},
}

// ProperlySpaceText
func ProperlySpaceText(text string, sz int, ln int) string {
	str := strings.ReplaceAll(text, " ", "")
	// fmt.Println(str)
	var b bytes.Buffer
	for i := 0; i < len(str); i++ {

		if ln != -1 && i%ln == 0 {
			b.WriteString("\n")
		} else if i%sz == 0 && i != 0 {
			b.WriteString(" ")
		}
		b.WriteString(string(str[i]))
	}
	newStr := b.String()
	return newStr
	// fmt.Println(newStr)
}

// RemoveRepeated
func RemoveRepeated(str string) string {
	arr := strings.Split(str, "")
	found := map[string]bool{}

	res := []string{}
	// Create a map of all unique elements.
	for v := range arr {
		if !found[arr[v]] {
			res = append(res, arr[v])
		}
		found[arr[v]] = true
	}

	// Place all keys from the map into a slice.

	return strings.Join(res, "")
}

// PrintMatrix
func PrintMatrix(mat [][]string) {

	for _, c := range mat {
		fmt.Println(c)
	}
}

// FindValueInMap
func FindValueInMap(m map[string]string, find string) bool {
	// fmt.Println("To Find:", find)
	for _, v := range m {
		// fmt.Println(k, v)
		if find == v {
			return true
		}
	}
	return false
}

// SeparateTextByNSpace
func SeparateTextByNSpace(str string, sep int) string {
	s := ""
	for _, c := range str {
		s += string(c)
		for i := 0; i < sep; i++ {
			s += " "
		}
	}
	return s
}

// PROBLEMA ACTUAL: Hay valores que ya estan asignados dentro del mapa
// 					sin embargo, como el ngrama correspondiente no lo esta,
// 					le asigna un codigo que ya existe igualmente
// FindPossibleExpInText:
func FindPossibleExpInText(txt string, ex string, ngramSize int) []string {
	strarr := strings.Split(ProperlySpaceText(txt, ngramSize, -1), " ") // Splitting
	// fmt.Println(ProperlySpaceText(txt, 2, 50))
	// count := 0
	l := make([]string, 0)
	for i := 0; i < len(strarr)-len(ex); i++ {
		m := make(map[string]string)
		str := ""
		// s := strarr[i]
		for j, c := range strings.Split(ex, "") {

			ngram := strarr[i+j]
			val, contains := m[ngram]
			if !contains {
				if !FindValueInMap(m, c) {

					// If map does not contain the ngram, and doesn't contain
					// the value, then any match is valid,
					//  so we just add to map and str and continue
					m[ngram] = c
					val = c
				} else {
					break
				}

				// We didn't find the key but we found the value
				// Since we don't want to allow multiple ngrams per symbol
				// we exit and try again
			}
			// Map contains ngram
			if val == c { // Is the ngram's associated value = value we need?
				str += ngram
			} else {
				break
			}
		}

		if len(str)/ngramSize >= len(ex) { // @info maybe just == would work?
			l = append(l, str)
		}
	}

	return l
	// for i := 0; i < len(strarr); i++ {
	// 	str := ""
	// 	if i+len(ex) >= len(strarr) {
	// 		break
	// 	}
	// 	m := make(map[string]string)
	// 	ji := i - 1 // Jizz dawg
	// 	for j, c := range ex {
	// 		fmt.Printf("(%d, %d) -> %d@%s := %s\n", i, j, count, m[string(c)], str)
	// 		if count == len(ex) {
	// 			l = append(l, str)
	// 			str = ""
	// 			count = 0
	// 			break
	// 		}
	// 		ch := string(c)
	// 		val, contains := m[ch]
	// 		if !contains {
	// 			m[ch] = string(strarr[j+i])
	// 			val = m[ch]
	// 		}
	// 		// Contains
	// 		ji++
	// 		fmt.Println(val, string(strarr[ji]), j+i, ji)
	// 		if val == string(strarr[j+i]) {
	// 			// fmt.Printf("%s, %s, %d\n", c, val, j%len(ex))
	// 			str += val
	// 			count++
	// 		} else {
	// 			str = ""
	// 			count = 0
	// 			break
	// 		}
	// 	}
	// 	// fmt.Println("I:=", i)
	// }
	// return l
}

// PerformTabularTransformation
func PerformTabularTransformation(str string, oKey string, sz int) (string, [][]string) {
	s := strings.Split(ProperlySpaceText(str, sz, -1), " ")
	key := RemoveRepeated(oKey)
	m := make([][]string, int(len(str)/len(key)+1))
	m1 := make([][]string, int(len(str)/len(key)+1))

	for i := 0; i < len(m); i++ {
		m[i] = make([]string, len(key))
		m1[i] = make([]string, len(key))
	}
	r := -1
	for i, c := range s {
		// fmt.Println(i, c, i%len(key))

		if i%len(key) == 0 {
			r++
		}
		m[r][i%len(key)] = c
	}
	ma := MapTabularKeyPriorities(key)
	// fmt.Println("[G  R  A  M  N  E  T  I  V]")
	// PrintMatrix(m)
	// Perform Column Swaps
	for i, c := range key {

		for j := 0; j < len(m); j++ {
			m1[j][ma[string(c)]] = m[j][i]
			m1[j][i] = m[j][ma[string(c)]]
		}
	}
	// fmt.Println("[A  E  G  I  M  N  R  T  V]")
	// PrintMatrix(m1)
	// fmt.Println(m)
	// Convert Matrix to String
	var newStr string = ""
	for i := 0; i < len(m[i]); i++ {
		for j := 0; j < len(m); j++ {
			newStr += m1[j][i]
		}
	}
	return newStr, m1
}

// MapTabularKeyPriorities
func MapTabularKeyPriorities(key string) map[string]int {
	arr := strings.Split(key, "")
	sort.Strings(arr)
	m := make(map[string]int)
	v := 0
	for _, c := range arr {
		if _, contains := m[c]; !contains {
			m[c] = v
			v++
		}
	}
	return m
}

// GetDifferentSymbols
func GetDifferentSymbols(str *string, sz int) ([]string, int, int) {
	arr := make([]string, len(*str))
	var val int = 0
	var q int = 0
	m := GenerateSymbolFreqMap(*str, sz)
	for k, v := range m {
		if v == 0 {
			continue
		}
		arr[val] = k
		val++
		q += v
	}

	arr2 := make([]string, val)
	for i := 0; i < val; i++ {
		arr2[i] = arr[i]
	}
	return arr2, val, q
}

// GenerateSymbolFreqMap: NOTE: IT DOES NOT IGNORE SPACES
func GenerateSymbolFreqMap(text string, sz int) map[string]int {
	m := make(map[string]int)
	str := ProperlySpaceText(text, sz, -1)
	arr := strings.Split(str, " ")
	for _, s := range arr {
		m[s]++
	}

	return m
	// if sz == 1 {
	// 	for _, c := range AbcSpanish {
	// 		C := string(c)
	// 		m[C] = strings.Count(str, C)
	// 	}
	// } else {
	// 	for i := 0; i < len(str); i++ {
	// 		var b bytes.Buffer
	// 		if str[i] == ' ' {
	// 			i++
	// 			continue
	// 		}
	// 		b.WriteString(string(str[i]))
	// 		for j := 1; j < sz && i+j < len(str); j++ {
	// 			if str[i+j] == ' ' {
	// 				b.WriteString(string(str[i+j+1]))
	// 				j++
	// 				i++
	// 			} else {
	// 				b.WriteString(string(str[i+j]))
	// 			}
	// 			i++
	// 		}
	// 		C := b.String()
	// 		m[C] = strings.Count(str, C)
	// 	}
	// }
}

// FindInArr
func FindInArr(arr []rune, find rune) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == find {
			return i
		}
	}
	return -1
}

// CaesarRotation
func CaesarRotation(text string, rot int, alphabet string) string {
	if rot == 0 {
		return text
	}
	if rot < 0 {
		rot = len(alphabet) - 1 + rot
	}
	// str := make([]rune, len(text)) // Strings are inmutable so we make an array to convert it at a later point
	// length := len(alphabet) - 1
	// fmt.Println(length)
	// ABC := []rune(alphabet)
	// for idx, c := range text {
	// 	if unicode.IsSpace(c) {
	// 		str[idx] = ' '
	// 		continue
	// 	}
	// 	i := (findInStr(alphabet, c) + rot) % length
	// 	// fmt.Printf("(%s at %d)->%d + %d->(%s at %d)\n", string(c), idx, findInStr(alphabet, c), rot, string(ABC[(findInStr(alphabet, c)+rot)%length]), (findInStr(alphabet, c)+rot)%length)
	// 	str[idx] = ABC[i]
	// }
	// return string(str)

	// Make the inverse alphabet
	inverseAbc := make([]rune, len(alphabet))
	abc := []rune(alphabet)

	for i := 0; i < len(abc); i++ {
		inverseAbc[i] = abc[(i+rot)%len(alphabet)]
	}
	str := make([]rune, len(text))
	cyphertext := []rune(text)
	for i := 0; i < len(text); i++ {
		idx := FindInArr(abc, cyphertext[i])
		if idx == -1 {
			str[i] = ' '
			continue
		}
		// fmt.Println(idx)
		str[i] = inverseAbc[idx]
	}
	return string(str)
}

// GetPolybiusKey
func GetPolybiusKey(alphabet string, key string) ([][]string, string) {

	k := make([]rune, len(alphabet))
	mat := make([][]string, int(math.Sqrt(float64(len(alphabet)))))
	for i := 0; i < len(mat); i++ {
		mat[i] = make([]string, int(math.Sqrt(float64(len(alphabet)))))
	}
	val := 0
	// Generamos la clave k
	for i, c := range key {
		if FindInArr(k, c) == -1 {
			k[i] = c
			val++
		}
	}

	for _, c := range alphabet {
		if FindInArr(k, c) == -1 {
			k[val] = c
			val++
		}
	}

	// Intentamos obtener la matriz de clave
	for i, v := 0, 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			// fmt.Println(i, j, string(k[v]), v)
			// fmt.Println(i, j)
			mat[i][j] = string(k[v])
			v++
		}
	}
	return mat, string(k)
}

// LogBase:
func LogBase(a float64, b int) float64 {
	return math.Log(a) / math.Log(float64(b))
}

// CheckEntropy:
func CheckEntropyStr(data string, b int) float64 {
	// Para hacer que se trate de una cadena de bytes,
	// b = 8;
	// Un byte puede tener 8 valores distintos
	// Por lo demas esto es una aplicacion directa de la formula de Shannon
	// vista en  https://en.wikipedia.org/wiki/Entropy_%28information_theory%29
	var ent float64 = -1
	var pT float64 = 0
	// Conteo de Frecuencias
	m := make(map[rune]int)
	for _, r := range data {
		m[r]++
	}
	// Calculamos la distribucion de probabilidad y el logaritmo
	// Sobre el total de caracteres distints
	for _, v := range m {
		var pI float64 = float64(v) * LogBase(float64(v), b)
		pT += pI
	}
	// Sin embargo, como no estamos basandonos en el string data
	// sino en el map, debemos realizar un ultimo calculo
	ent *= pT/float64(len(data)) - LogBase(float64(len(data)), b)
	return ent
}

type pixl struct {
	r, g, b, a int
}

func ConvertRGBToFloat(rgba *pixl) float64 {
	return float64(rgba.r + rgba.g*256 + rgba.b*256*256)
}

func ConvertFloatToRGB(val float64, rgba *pixl) {

	rgba.b = int(val / 256 * 256)
	rgba.g = int((val - float64(rgba.b)*math.Pow(256, 2)) / 256)
	rgba.r = int(int(val) % 256)
}

func CalculateFileEntropy(filename string) float64 {
	dat, _ := ioutil.ReadFile(filename)
	return CheckEntropyStr(string(dat), 8)
}

func GenerateHeatmap(window int, filename string) []pixl {

	dat, _ := ioutil.ReadFile(filename)
	numPixls := len(dat)/window + 1
	heatmap := make([]pixl, numPixls)
	data := make([]string, numPixls)
	// Hacemos un split del string en chunks de bytes del sz de la ventan de bytes
	i := 0
	for idx, b := range dat {
		if idx%window == 0 {
			i++
		} else {
			data[i] += string(b)
		}
	}
	for i = 0; i < len(data); i++ {
		ent := CheckEntropyStr(data[i], 8) // 8 bc bytes are 8 bits
		ent *= 10000                       // Para hacer el cambio mas obvio y visible en el mapa
		pixel := pixl{0, 0, 0, 255}        // Init a RGBA(0,0,0,255)
		ConvertFloatToRGB(ent, &pixel)
		heatmap[i] = pixel
	}

	return heatmap
}

func ApproximateEntropyFromHeatmap(heatmap []pixl) float64 {

	var val float64 = 0
	for _, pix := range heatmap {
		val += ConvertRGBToFloat(&pix)
	}
	return val / 10000 / float64(len(heatmap))
}

func ApproximateEntropyOfFile(filename string, window int) float64 {
	heatmap := GenerateHeatmap(window, filename)
	return ApproximateEntropyFromHeatmap(heatmap)
}
