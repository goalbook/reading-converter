package main

import (
    "log"
    "encoding/csv"
    "strings"
    "strconv"
    //"fmt"
    //"io"
)

type LexileConverter struct {
    Records [][]string
}

func main() {
    log.Printf("hello world! %v", conversionTableHeader)
}

func NewConversionMap(v []string) map[string]string {
    r := map[string]string{}
    for i := 0; i < len(conversionTableHeader); i++ {
        r[conversionTableHeader[i]] = v[i]
    }
    return r
}

func NewLexileConverter() (LexileConverter, error) {
    conv := LexileConverter{}

    r := csv.NewReader(strings.NewReader(conversionTableData))

    records, err := r.ReadAll()
    if err != nil {
        log.Printf("Error parsing CSV: %v", err)
        return conv, err
    }

    conv.Records = records

    return conv, nil
}

func (lconv *LexileConverter) Convert(s string) (map[string]string, error) {
    r := map[string]string{}
    v := strings.TrimSpace(s)

    if strings.HasPrefix(v, "BR") {
        v = "0" // We treat any BR lexile values as if they were score 0 which is not exactly right.
    }
    n, err := strconv.Atoi(v)
    if err != nil {
        return r, err
    }

    for i := 0; i < len(lconv.Records); i++ {
        lstart, err := strconv.Atoi(lconv.Records[i][0])
        if err != nil {
            return r, err
        }
        lend, err := strconv.Atoi(lconv.Records[i][1])
        if err != nil {
            return r, err
        }
        if n >= lstart && n <= lend {
            r = NewConversionMap(lconv.Records[i])
            break
        }
    }
    return r, nil
}


// struct {
    
//     primary_value: [ 430 ]
//     primary_type: [ Lexile ]

//     converted_scores: {

//         [type]: [ approximate, low, high ]



//     }
// }

// ConvertScore(Score, Type)


// Columns := [ 'Grade Level' ],

// ConversionTable := 
// `sdf,sdfsdf,sdfsdf,
// sdfsdf,sdfsdf,sdfsdf,`


// lexile_start    lexile_end  lexile  grade   grade_ext   ar  dra fp  gr



// package main

// import (
//     "encoding/csv"
//     "fmt"
//     "log"
//     "strings"
// )

// func main() {
//     in := `first_name,last_name,username
// "Rob","Pike",rob
// Ken,Thompson,ken
// "Robert","Griesemer","gri"
// `
//     r := csv.NewReader(strings.NewReader(in))

//     records, err := r.ReadAll()
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Print(records)
// }


