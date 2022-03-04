package main

import (
	"flag"
	"fmt"
)

var (
	file  string
	table string
)

func init() {
	flag.StringVar(&file, "f", "", "The `file` to read from")
	flag.StringVar(&table, "t", "", "The table to read from")
}

func main() {
	flag.Parse()

	fmt.Printf("file: %v", file)
	fmt.Printf("table: %v", table)

	// fd, err := os.Open(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer fd.Close()

	// tokens := sqlparser.NewTokenizer(fd)
	// for {
	// 	stmt, err := sqlparser.ParseNext(tokens)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		if strings.Contains(err.Error(), "commit") {
	// 			fmt.Print("\ncommit;\n")
	// 		} else {
	// 			log.Fatal(err)
	// 		}
	// 		continue
	// 	}
	// 	switch stmt := stmt.(type) {
	// 	case *sqlparser.Insert:
	// 		// if stmt.Table.Name.String() == "produtosfornecedor" {
	// 		if stmt.Table.Name.String() == table {
	// 			log.Printf("Found insert for table %v", stmt.Table)
	// 			stmt.WalkSubtree(func(node sqlparser.SQLNode) (bool, error) {
	// 				switch node := node.(type) {
	// 				case sqlparser.ValTuple:
	// 					col := 0
	// 					node.WalkSubtree(func(node sqlparser.SQLNode) (bool, error) {
	// 						switch node := node.(type) {
	// 						case *sqlparser.SQLVal:
	// 							if col >= 15 && col <= 16 {
	// 								node.Val = []byte("")
	// 							}
	// 							col++
	// 						}
	// 						return true, nil
	// 					})
	// 					return false, nil
	// 				}
	// 				return true, nil
	// 			})
	// 		}
	// 		fmt.Printf("%s;\n", sqlparser.String(stmt))
	// 	default:
	// 		if stmt == nil {
	// 			continue
	// 		}
	// 		fmt.Printf("%s;\n", sqlparser.String(stmt))
	// 	}
	// }
}
