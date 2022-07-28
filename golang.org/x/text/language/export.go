// export by github.com/goplus/igop/cmd/qexp

package language

import (
	q "golang.org/x/text/language"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "language",
		Path: "golang.org/x/text/language",
		Deps: map[string]string{
			"errors":                              "errors",
			"fmt":                                 "fmt",
			"golang.org/x/text/internal/language": "language",
			"golang.org/x/text/internal/language/compact": "compact",
			"sort":    "sort",
			"strconv": "strconv",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{
			"Coverage":   reflect.TypeOf((*q.Coverage)(nil)).Elem(),
			"Matcher":    reflect.TypeOf((*q.Matcher)(nil)).Elem(),
			"ValueError": reflect.TypeOf((*q.ValueError)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Base":        reflect.TypeOf((*q.Base)(nil)).Elem(),
			"CanonType":   reflect.TypeOf((*q.CanonType)(nil)).Elem(),
			"Confidence":  reflect.TypeOf((*q.Confidence)(nil)).Elem(),
			"Extension":   reflect.TypeOf((*q.Extension)(nil)).Elem(),
			"MatchOption": reflect.TypeOf((*q.MatchOption)(nil)).Elem(),
			"Region":      reflect.TypeOf((*q.Region)(nil)).Elem(),
			"Script":      reflect.TypeOf((*q.Script)(nil)).Elem(),
			"Tag":         reflect.TypeOf((*q.Tag)(nil)).Elem(),
			"Variant":     reflect.TypeOf((*q.Variant)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Afrikaans":                reflect.ValueOf(&q.Afrikaans),
			"Albanian":                 reflect.ValueOf(&q.Albanian),
			"AmericanEnglish":          reflect.ValueOf(&q.AmericanEnglish),
			"Amharic":                  reflect.ValueOf(&q.Amharic),
			"Arabic":                   reflect.ValueOf(&q.Arabic),
			"Armenian":                 reflect.ValueOf(&q.Armenian),
			"Azerbaijani":              reflect.ValueOf(&q.Azerbaijani),
			"Bengali":                  reflect.ValueOf(&q.Bengali),
			"BrazilianPortuguese":      reflect.ValueOf(&q.BrazilianPortuguese),
			"BritishEnglish":           reflect.ValueOf(&q.BritishEnglish),
			"Bulgarian":                reflect.ValueOf(&q.Bulgarian),
			"Burmese":                  reflect.ValueOf(&q.Burmese),
			"CanadianFrench":           reflect.ValueOf(&q.CanadianFrench),
			"Catalan":                  reflect.ValueOf(&q.Catalan),
			"Chinese":                  reflect.ValueOf(&q.Chinese),
			"Croatian":                 reflect.ValueOf(&q.Croatian),
			"Czech":                    reflect.ValueOf(&q.Czech),
			"Danish":                   reflect.ValueOf(&q.Danish),
			"Dutch":                    reflect.ValueOf(&q.Dutch),
			"English":                  reflect.ValueOf(&q.English),
			"ErrMissingLikelyTagsData": reflect.ValueOf(&q.ErrMissingLikelyTagsData),
			"Estonian":                 reflect.ValueOf(&q.Estonian),
			"EuropeanPortuguese":       reflect.ValueOf(&q.EuropeanPortuguese),
			"EuropeanSpanish":          reflect.ValueOf(&q.EuropeanSpanish),
			"Filipino":                 reflect.ValueOf(&q.Filipino),
			"Finnish":                  reflect.ValueOf(&q.Finnish),
			"French":                   reflect.ValueOf(&q.French),
			"Georgian":                 reflect.ValueOf(&q.Georgian),
			"German":                   reflect.ValueOf(&q.German),
			"Greek":                    reflect.ValueOf(&q.Greek),
			"Gujarati":                 reflect.ValueOf(&q.Gujarati),
			"Hebrew":                   reflect.ValueOf(&q.Hebrew),
			"Hindi":                    reflect.ValueOf(&q.Hindi),
			"Hungarian":                reflect.ValueOf(&q.Hungarian),
			"Icelandic":                reflect.ValueOf(&q.Icelandic),
			"Indonesian":               reflect.ValueOf(&q.Indonesian),
			"Italian":                  reflect.ValueOf(&q.Italian),
			"Japanese":                 reflect.ValueOf(&q.Japanese),
			"Kannada":                  reflect.ValueOf(&q.Kannada),
			"Kazakh":                   reflect.ValueOf(&q.Kazakh),
			"Khmer":                    reflect.ValueOf(&q.Khmer),
			"Kirghiz":                  reflect.ValueOf(&q.Kirghiz),
			"Korean":                   reflect.ValueOf(&q.Korean),
			"Lao":                      reflect.ValueOf(&q.Lao),
			"LatinAmericanSpanish":     reflect.ValueOf(&q.LatinAmericanSpanish),
			"Latvian":                  reflect.ValueOf(&q.Latvian),
			"Lithuanian":               reflect.ValueOf(&q.Lithuanian),
			"Macedonian":               reflect.ValueOf(&q.Macedonian),
			"Malay":                    reflect.ValueOf(&q.Malay),
			"Malayalam":                reflect.ValueOf(&q.Malayalam),
			"Marathi":                  reflect.ValueOf(&q.Marathi),
			"ModernStandardArabic":     reflect.ValueOf(&q.ModernStandardArabic),
			"Mongolian":                reflect.ValueOf(&q.Mongolian),
			"Nepali":                   reflect.ValueOf(&q.Nepali),
			"Norwegian":                reflect.ValueOf(&q.Norwegian),
			"Persian":                  reflect.ValueOf(&q.Persian),
			"Polish":                   reflect.ValueOf(&q.Polish),
			"Portuguese":               reflect.ValueOf(&q.Portuguese),
			"Punjabi":                  reflect.ValueOf(&q.Punjabi),
			"Romanian":                 reflect.ValueOf(&q.Romanian),
			"Russian":                  reflect.ValueOf(&q.Russian),
			"Serbian":                  reflect.ValueOf(&q.Serbian),
			"SerbianLatin":             reflect.ValueOf(&q.SerbianLatin),
			"SimplifiedChinese":        reflect.ValueOf(&q.SimplifiedChinese),
			"Sinhala":                  reflect.ValueOf(&q.Sinhala),
			"Slovak":                   reflect.ValueOf(&q.Slovak),
			"Slovenian":                reflect.ValueOf(&q.Slovenian),
			"Spanish":                  reflect.ValueOf(&q.Spanish),
			"Supported":                reflect.ValueOf(&q.Supported),
			"Swahili":                  reflect.ValueOf(&q.Swahili),
			"Swedish":                  reflect.ValueOf(&q.Swedish),
			"Tamil":                    reflect.ValueOf(&q.Tamil),
			"Telugu":                   reflect.ValueOf(&q.Telugu),
			"Thai":                     reflect.ValueOf(&q.Thai),
			"TraditionalChinese":       reflect.ValueOf(&q.TraditionalChinese),
			"Turkish":                  reflect.ValueOf(&q.Turkish),
			"Ukrainian":                reflect.ValueOf(&q.Ukrainian),
			"Und":                      reflect.ValueOf(&q.Und),
			"Urdu":                     reflect.ValueOf(&q.Urdu),
			"Uzbek":                    reflect.ValueOf(&q.Uzbek),
			"Vietnamese":               reflect.ValueOf(&q.Vietnamese),
			"Zulu":                     reflect.ValueOf(&q.Zulu),
		},
		Funcs: map[string]reflect.Value{
			"CompactIndex":        reflect.ValueOf(q.CompactIndex),
			"Compose":             reflect.ValueOf(q.Compose),
			"Comprehends":         reflect.ValueOf(q.Comprehends),
			"EncodeM49":           reflect.ValueOf(q.EncodeM49),
			"Make":                reflect.ValueOf(q.Make),
			"MatchStrings":        reflect.ValueOf(q.MatchStrings),
			"MustParse":           reflect.ValueOf(q.MustParse),
			"MustParseBase":       reflect.ValueOf(q.MustParseBase),
			"MustParseRegion":     reflect.ValueOf(q.MustParseRegion),
			"MustParseScript":     reflect.ValueOf(q.MustParseScript),
			"NewCoverage":         reflect.ValueOf(q.NewCoverage),
			"NewMatcher":          reflect.ValueOf(q.NewMatcher),
			"Parse":               reflect.ValueOf(q.Parse),
			"ParseAcceptLanguage": reflect.ValueOf(q.ParseAcceptLanguage),
			"ParseBase":           reflect.ValueOf(q.ParseBase),
			"ParseExtension":      reflect.ValueOf(q.ParseExtension),
			"ParseRegion":         reflect.ValueOf(q.ParseRegion),
			"ParseScript":         reflect.ValueOf(q.ParseScript),
			"ParseVariant":        reflect.ValueOf(q.ParseVariant),
			"PreferSameScript":    reflect.ValueOf(q.PreferSameScript),
		},
		TypedConsts: map[string]igop.TypedConst{
			"All":              {reflect.TypeOf(q.All), constant.MakeInt64(int64(q.All))},
			"BCP47":            {reflect.TypeOf(q.BCP47), constant.MakeInt64(int64(q.BCP47))},
			"CLDR":             {reflect.TypeOf(q.CLDR), constant.MakeInt64(int64(q.CLDR))},
			"Default":          {reflect.TypeOf(q.Default), constant.MakeInt64(int64(q.Default))},
			"Deprecated":       {reflect.TypeOf(q.Deprecated), constant.MakeInt64(int64(q.Deprecated))},
			"DeprecatedBase":   {reflect.TypeOf(q.DeprecatedBase), constant.MakeInt64(int64(q.DeprecatedBase))},
			"DeprecatedRegion": {reflect.TypeOf(q.DeprecatedRegion), constant.MakeInt64(int64(q.DeprecatedRegion))},
			"DeprecatedScript": {reflect.TypeOf(q.DeprecatedScript), constant.MakeInt64(int64(q.DeprecatedScript))},
			"Exact":            {reflect.TypeOf(q.Exact), constant.MakeInt64(int64(q.Exact))},
			"High":             {reflect.TypeOf(q.High), constant.MakeInt64(int64(q.High))},
			"Legacy":           {reflect.TypeOf(q.Legacy), constant.MakeInt64(int64(q.Legacy))},
			"Low":              {reflect.TypeOf(q.Low), constant.MakeInt64(int64(q.Low))},
			"Macro":            {reflect.TypeOf(q.Macro), constant.MakeInt64(int64(q.Macro))},
			"No":               {reflect.TypeOf(q.No), constant.MakeInt64(int64(q.No))},
			"Raw":              {reflect.TypeOf(q.Raw), constant.MakeInt64(int64(q.Raw))},
			"SuppressScript":   {reflect.TypeOf(q.SuppressScript), constant.MakeInt64(int64(q.SuppressScript))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"CLDRVersion":    {"untyped string", constant.MakeString(string(q.CLDRVersion))},
			"NumCompactTags": {"untyped int", constant.MakeInt64(int64(q.NumCompactTags))},
		},
	})
}
