package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/fatih/color"
)

func findLangBlock(body []byte) ([]byte, error) {
	reg := regexp.MustCompile(`name="programTypeId".+?</select>`)
	tmp := reg.Find(body)
	if tmp == nil {
		return nil, errors.New("Cannot find language selection")
	}
	return tmp, nil
}

func findLang(body []byte) (map[string]string, error) {
	reg := regexp.MustCompile(`value="(.+?)"[\s\S]*?>([\s\S]+?)<`)
	tmp := reg.FindAllSubmatch(body, -1)
	if tmp == nil {
		return nil, errors.New("Cannot find any language")
	}
	ret := make(map[string]string)
	for i := 0; i < len(tmp); i++ {
		ret[string(tmp[i][1])] = string(tmp[i][2])
	}
	return ret, nil
}

// GetLangList get language list from url (require login)
func (c *Client) GetLangList(url string) (langs map[string]string, err error) {
	color.Cyan("Getting language list...")
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	block, err := findLangBlock(body)
	if err != nil {
		return
	}

	return findLang(block)
}

// Langs generated by
// ^[\s\S]*?value="(.+?)"[\s\S]*?>([\s\S]+?)<[\s\S]*?$
//     "\1": "\2",
var Langs = map[string]string{
	"43": "GNU GCC C11 5.1.0",
	"52": "Clang++17 Diagnostics",
	"42": "GNU G++11 5.1.0",
	"50": "GNU G++14 6.4.0",
	"54": "GNU G++17 7.3.0",
	"2":  "Microsoft Visual C++ 2010",
	"59": "Microsoft Visual C++ 2017",
	"9":  "C# Mono 5.18",
	"28": "D DMD32 v2.083.1",
	"32": "Go 1.11.4",
	"12": "Haskell GHC 8.6.3",
	"36": "Java 1.8.0_162",
	"48": "Kotlin 1.3.10",
	"19": "OCaml 4.02.1",
	"3":  "Delphi 7",
	"4":  "Free Pascal 3.0.2",
	"51": "PascalABC.NET 3.4.2",
	"13": "Perl 5.20.1",
	"6":  "PHP 7.2.13",
	"7":  "Python 2.7.15",
	"31": "Python 3.7.2",
	"40": "PyPy 2.7 (6.0.0)",
	"41": "PyPy 3.5 (6.0.0)",
	"8":  "Ruby 2.0.0p645",
	"49": "Rust 1.31.1",
	"20": "Scala 2.12.8",
	"34": "JavaScript V8 4.8.0",
	"55": "Node.js 9.4.0",
}

// LangsExt language's ext
var LangsExt = map[string]string{
	"GNU C11":               "c",
	"Clang++17 Diagnostics": "cpp",
	"GNU C++11":             "cpp",
	"GNU C++14":             "cpp",
	"GNU C++17":             "cpp",
	"MS C++":                "cpp",
	"MS C++ 2017":           "cpp",
	"Mono C#":               "cs",
	"D":                     "d",
	"Go":                    "go",
	"Haskell":               "hs",
	"Kotlin":                "kt",
	"Ocaml":                 "ml",
	"Delphi":                "pas",
	"FPC":                   "pas",
	"PascalABC.NET":         "pas",
	"Perl":                  "pl",
	"PHP":                   "php",
	"Python 2":              "py",
	"Python 3":              "py",
	"PyPy 2":                "py",
	"PyPy 3":                "py",
	"Ruby":                  "rb",
	"Rust":                  "rs",
	"JavaScript":            "js",
	"Node.js":               "js",
}
