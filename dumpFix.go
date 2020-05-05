package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	insert := ""
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648259,90821,'CI*****VA','42797163826','04*****77','2017-10-16 17:24:12',19,788043,'06/05/94',0,'','ED*****VA','','','','','','','','005006459','F','4658887550566400','SSP','2012-04-19 12:00:00','SP','Solteiro(a)','','sodexo',0.00,NULL,NULL,0);\n"
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648260,90821,'DA*****OR','01291141111','00*****14','2017-10-16 17:25:57',27,788045,'18/05/1984',0,'','AN*****VA','','','','','','','','001452632','M','4676808301608960','SESP','2015-04-10 12:00:00','MT','Solteiro(a)','','sodexo',0.00,'NULL',NULL,0);\n"
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648261,90821,'GL*****VA','37307778823','04*****86','2017-10-16 17:24:13',19,788043,'01/01/90',0,'','NO*****\\','','','','','','','','005006472','F','5095870945034240','SSP','2010-01-27 12:00:00','SP','Casado(a)','','sodexo',0.00,NULL,NULL,0);\n"
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648259,90821,'CI*****VA','42797163826','04*****77','2017-10-16 17:24:12',19,788043,'06/05/94',0,'','ED*****VA','','','','','','','','005006459','F','4658887550566400','SSP','2012-04-19 12:00:00','SP','Solteiro(a)','','sodexo',0.00,NULL,NULL,0);\n"
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648260,90821,'DA*****OR','01291141111','00*****14','2017-10-16 17:25:57',27,788045,'18/05/1984',0,'','AN*****VA','','','','','','','','001452632','M','4676808301608960','SESP','2015-04-10 12:00:00','MT','Solteiro(a)','','sodexo',0.00,'NULL',NULL,0);\n"
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648261,90821,'GL*****VA','37307778823','04*****86','2017-10-16 17:24:13',19,788043,'01/01/90',0,'','NO*****\\','','','','','','','','005006472','F','5095870945034240','SSP','2010-01-27 12:00:00','SP','Casado(a)','','sodexo',0.00,NULL,NULL,0);\n"
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648259,90821,'CI*****VA','42797163826','04*****77','2017-10-16 17:24:12',19,788043,'05/04/94',0,'','ED*****VA','','','','','','','','005006459','F','4658887550566400','SSP','2012-04-19 12:00:00','SP','Solteiro(a)','','sodexo',0.00,NULL,NULL,0);\n"
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648260,90821,'DA*****OR','01291141111','00*****14','2017-10-16 17:25:57',27,788045,'17/11/1984',0,'','AN*****VA','','','','','','','','001452632','M','4676808301608960','SESP','2015-04-10 12:00:00','MT','Solteiro(a)','','sodexo',0.00,'NULL',NULL,0);\n"
	insert += "INSERT  IGNORE INTO `funcionarios` VALUES (6648261,90821,'GL*****VA','37307778823','04*****86','2017-10-16 17:24:13',19,788043,'22/06/90',0,'','NO*****\\','','','','','','','','005006472','F','5095870945034240','SSP','2010-01-27 12:00:00','SP','Casado(a)','','sodexo',0.00,NULL,NULL,0);\n"
	dumpFix(insert)
}

func dumpFix(input string) string {
	rows := strings.Split(input, "\n")
	for i := range rows {
		if strings.Contains(rows[i], "INSERT  IGNORE INTO `") {
			data := strings.Split(myregexp("INSERT\\s+IGNORE INTO `.*` VALUES \\((.*)\\)", rows[i]), ",")
			for j := range data {
				if strings.HasSuffix(data[j], "\\'") {
					fmt.Printf("Before fix: %s\n", rows[i])
					data[j] = fmt.Sprintf("%v%v", data[j][:len(data[j])-2], data[j][len(data[j])-1:])
					fmt.Printf("After  fix: %s\n", rows[i])
				}
			}
			rows[i] = fmt.Sprintf("%s VALUES %s", strings.Split(rows[i], "VALUES")[0], strings.Join(data, ","))
			fmt.Println()
		}
	}
	return strings.Join(rows, "\n")
}

func myregexp(exp, text string) string {
	re := regexp.MustCompile(exp)
	match := re.FindStringSubmatch(text)
	if len(match) < 1 {
		log.Println("Unable to find match for regexp")
		return ""
	}
	if strings.Contains(exp, "|") && len(match[1]) == 0 && len(match) > 2 {
		log.Println("Found secondary match")
		return match[2]
	}
	return match[1]
}
