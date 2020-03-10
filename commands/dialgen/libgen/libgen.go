package libgen

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"regexp"
	"strings"
)

// XMLToFields : This function takes a string of a xml file location, and a string of the location of common.xml as input.
// It returns []*OutDefinition and a string of version number, which is required for go code generation - see gomavlib/commands/dialgen/gen
func XMLToFields(filePathXML string, filePathCommonXML string) ([]*OutDefinition, string) {
	outDefs, version, err := do("", filePathXML, filePathCommonXML)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return outDefs, version
}

/////////////// Code copied from original main.go in aler9/gomavlib/commands/dialgen below here... //////////////////////////////////////////

var reMsgName = regexp.MustCompile("^[A-Z0-9_]+$")
var reTypeIsArray = regexp.MustCompile("^(.+?)\\[([0-9]+)\\]$")

var dialectTypeToGo = map[string]string{
	"double":   "float64",
	"uint64_t": "uint64",
	"int64_t":  "int64",
	"float":    "float32",
	"uint32_t": "uint32",
	"int32_t":  "int32",
	"uint16_t": "uint16",
	"int16_t":  "int16",
	"uint8_t":  "uint8",
	"int8_t":   "int8",
	"char":     "string",
}

func dialectFieldGoToDef(in string) string {
	re := regexp.MustCompile("([A-Z])")
	in = re.ReplaceAllString(in, "_${1}")
	return strings.ToLower(in[1:])
}

func dialectFieldDefToGo(in string) string {
	return dialectMsgDefToGo(in)
}

func dialectMsgDefToGo(in string) string {
	re := regexp.MustCompile("_[a-z]")
	in = strings.ToLower(in)
	in = re.ReplaceAllStringFunc(in, func(match string) string {
		return strings.ToUpper(match[1:2])
	})
	return strings.ToUpper(in[:1]) + in[1:]
}

func filterDesc(in string) string {
	return strings.Replace(in, "\n", "", -1)
}

// OutEnumValue : Exported struct
type OutEnumValue struct {
	Value       string
	Name        string
	Description string
}

// OutEnum : Exported struct
type OutEnum struct {
	Name        string
	Description string
	Values      []*OutEnumValue
}

// OutField : Exported struct
type OutField struct {
	Description string
	Line        string
}

// OutMessage : Exported struct
type OutMessage struct {
	Name        string
	Description string
	Id          int
	Fields      []*OutField
}

// OutDefinition : Exported struct
type OutDefinition struct {
	Name     string
	Enums    []*OutEnum
	Messages []*OutMessage
}

func do(preamble string, mainDefAddr string, commonAddr string) ([]*OutDefinition, string, error) {
	version := ""
	defsProcessed := make(map[string]struct{})
	//isRemote := func() bool {
	//	_, err := url.ParseRequestURI(mainDefAddr)
	//	return err == nil
	//}()
	isRemote := false

	// parse all definitions recursively
	outDefs, err := definitionProcess(&version, defsProcessed, isRemote, mainDefAddr, commonAddr)
	if err != nil {
		return outDefs, version, err
	}

	return outDefs, version, nil
}

func definitionProcess(version *string, defsProcessed map[string]struct{}, isRemote bool, defAddr string, commonAddr string) ([]*OutDefinition, error) {
	// skip already processed
	if _, ok := defsProcessed[defAddr]; ok {
		return nil, nil
	}
	defsProcessed[defAddr] = struct{}{}

	fmt.Fprintf(os.Stderr, "definition %s\n", defAddr)

	content, err := definitionGet(isRemote, defAddr)
	if err != nil {
		return nil, err
	}

	def, err := DefinitionDecode(content)
	if err != nil {
		return nil, fmt.Errorf("unable to decode: %s", err)
	}

	addrPath, addrName := filepath.Split(defAddr)

	var outDefs []*OutDefinition

	// version
	if def.Version != "" {
		if *version != "" && *version != def.Version {
			return nil, fmt.Errorf("version defined twice (%s and %s)", def.Version, *version)
		}
		*version = def.Version
	}

	// includes
	for _, inc := range def.Includes {
		// prepend url to remote address
		if isRemote == true {
			inc = addrPath + inc
		}

		// Due to alternate location of some xml files, specify common.xml address if included in xml in question
		if (inc == "common.xml") && (commonAddr != "") {
			inc = commonAddr
		} else {
			// If common.xml (Or other included xml file) location not specified, then assume same directory as main xml file specified
			inc = filepath.Dir(defAddr) + "/" + inc
		}
		subDefs, err := definitionProcess(version, defsProcessed, isRemote, inc, commonAddr)
		if err != nil {
			return nil, err
		}
		outDefs = append(outDefs, subDefs...)
	}

	outDef := &OutDefinition{
		Name: addrName,
	}

	// enums
	for _, enum := range def.Enums {
		oute := &OutEnum{
			Name:        enum.Name,
			Description: filterDesc(enum.Description),
		}
		for _, val := range enum.Values {
			oute.Values = append(oute.Values, &OutEnumValue{
				Value:       val.Value,
				Name:        val.Name,
				Description: filterDesc(val.Description),
			})
		}
		outDef.Enums = append(outDef.Enums, oute)
	}

	// messages
	for _, msg := range def.Messages {
		outMsg, err := messageProcess(msg)
		if err != nil {
			return nil, err
		}
		outDef.Messages = append(outDef.Messages, outMsg)
	}

	outDefs = append(outDefs, outDef)
	return outDefs, nil
}

func definitionGet(isRemote bool, defAddr string) ([]byte, error) {
	if isRemote == true {
		byt, err := urlDownload(defAddr)
		if err != nil {
			return nil, fmt.Errorf("unable to download: %s", err)
		}
		return byt, nil
	}

	byt, err := ioutil.ReadFile(defAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to open: %s", err)
	}
	return byt, nil
}

func urlDownload(desturl string) ([]byte, error) {
	res, err := http.Get(desturl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad return code: %v", res.StatusCode)
	}

	byt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

func messageProcess(msg *DefinitionMessage) (*OutMessage, error) {
	if m := reMsgName.FindStringSubmatch(msg.Name); m == nil {
		return nil, fmt.Errorf("unsupported message name: %s", msg.Name)
	}

	outMsg := &OutMessage{
		Name:        dialectMsgDefToGo(msg.Name),
		Description: filterDesc(msg.Description),
		Id:          msg.Id,
	}

	for _, f := range msg.Fields {
		OutField, err := fieldProcess(f)
		if err != nil {
			return nil, err
		}
		outMsg.Fields = append(outMsg.Fields, OutField)
	}

	return outMsg, nil
}

func fieldProcess(field *DialectField) (*OutField, error) {
	outF := &OutField{
		Description: filterDesc(field.Description),
	}
	tags := make(map[string]string)

	newname := dialectFieldDefToGo(field.Name)

	// name conversion is not univoque: add tag
	if dialectFieldGoToDef(newname) != field.Name {
		tags["mavname"] = field.Name
	}

	outF.Line += newname

	typ := field.Type
	arrayLen := ""

	if typ == "uint8_t_mavlink_version" {
		typ = "uint8_t"
	}

	// string or array
	if matches := reTypeIsArray.FindStringSubmatch(typ); matches != nil {
		// string
		if matches[1] == "char" {
			tags["mavlen"] = matches[2]
			typ = "char"
			// array
		} else {
			arrayLen = matches[2]
			typ = matches[1]
		}
	}

	// extension
	if field.Extension == true {
		tags["mavext"] = "true"
	}

	typ = dialectTypeToGo[typ]
	if typ == "" {
		return nil, fmt.Errorf("unknown type: %s", typ)
	}

	outF.Line += " "
	if arrayLen != "" {
		outF.Line += "[" + arrayLen + "]"
	}
	if field.Enum != "" {
		outF.Line += field.Enum
		tags["mavenum"] = typ
	} else {
		outF.Line += typ
	}

	if len(tags) > 0 {
		var tmp []string
		for k, v := range tags {
			tmp = append(tmp, fmt.Sprintf("%s:\"%s\"", k, v))
		}
		sort.Strings(tmp)
		outF.Line += " `" + strings.Join(tmp, " ") + "`"
	}
	return outF, nil
}

////////////////////////////// Go Code Generation Functions... ///////////////////////////////////////

var tplDialect = template.Must(template.New("").Parse(
	`// Autogenerated with dialgen, do not edit.
{{- if .Preamble }}
//
// {{ .Preamble }}
//
{{- end }}
package {{ .PkgName }}

import (
	"github.com/team-rocos/gomavlib"
)

// Dialect contains the dialect object that can be passed to the library.
var Dialect = dialect

// dialect is not exposed directly such that it is not displayed in godoc.
var dialect = gomavlib.MustDialectCT({{.Version}}, []gomavlib.Message{
{{- range .Defs }}
    // {{ .Name }}
{{- range .Messages }}
    &Message{{ .Name }}{},
{{- end }}
{{- end }}
})

{{ range .Enums }}
// {{ .Description }}
type {{ .Name }} int

const (
{{- $pn := .Name }}
{{- range .Values }}
	// {{ .Description }}
	{{ .Name }} {{ $pn }} = {{ .Value }}
{{- end }}
)
{{ end }}

{{ range .Defs }}
// {{ .Name }}

{{ range .Messages }}
// {{ .Description }}
type Message{{ .Name }} struct {
{{- range .Fields }}
	// {{ .Description }}
    {{ .Line }}
{{- end }}
}

func (m *Message{{ .Name }}) GetId() uint32 {
    return {{ .Id }}
}

func (m *Message{{ .Name }}) SetField(field string, value interface{}) error {
	return gomavlib.SetMessageField(m, field, value)
}
{{ end }}
{{ end }}
`))

// GenerateGoCode : Exported Function
func GenerateGoCode(preamble string, mainDefAddr string, commonAddr string) error {

	outDefs, version := XMLToFields(mainDefAddr, commonAddr)

	// merge enums together
	enums := make(map[string]*OutEnum)
	for _, def := range outDefs {
		for _, defEnum := range def.Enums {
			if _, ok := enums[defEnum.Name]; !ok {
				enums[defEnum.Name] = &OutEnum{
					Name:        defEnum.Name,
					Description: defEnum.Description,
				}
			}
			enum := enums[defEnum.Name]

			for _, v := range defEnum.Values {
				enum.Values = append(enum.Values, v)
			}
		}
	}

	// fill enum missing values
	for _, enum := range enums {
		nextVal := 0
		for _, v := range enum.Values {
			if v.Value != "" {
				nextVal, _ = strconv.Atoi(v.Value)
				nextVal++
			} else {
				v.Value = strconv.Itoa(nextVal)
				nextVal++
			}
		}
	}

	// get package name
	// remove underscores since they can lead to errors
	// (for instance, when package name ends with _test)
	_, inFile := filepath.Split(mainDefAddr)
	pkgName := strings.TrimSuffix(inFile, ".xml")

	// dump
	return tplDialect.Execute(os.Stdout, map[string]interface{}{
		"PkgName":  pkgName,
		"Preamble": preamble,
		"Version": func() int {
			ret, _ := strconv.Atoi(version)
			return ret
		}(),
		"Defs":  outDefs,
		"Enums": enums,
	})
}
