package main

import (
	"fmt"
)

func main() {
	kinds := []string{
		"MacroRecharge",
		"ItemDetail",
		"LogRecord",
		"SyncID",
		"ImportRecord",
		"FileError",
		"InstallmentSimulation",
		"Employee",
		"Order",
		"_mail",
		"MacroError",
		"Checkout",
		"Items",
		"Address",
		"MacroOrderDate",
		"Organization",
		"OrderReportRecord",
		"Import",
		"Department",
		"User",
		"TowerDataResponse",
		"Fee",
		"RoutingDetail",
		"OrderCreditToken",
		"_Unique_OrgCNPJ",
		"_Unique_UserLogin",
		"_AE_Backup_Information_Kind_Type_Info",
		"News",
		"ImportClean",
		"ACL",
		"CEP",
		"ProductRef",
		"CNABOrder",
		"Itinerary",
		"_GAE_MR_TaskPayload",
		"SdxRefV3",
		"CNABFilename",
		"MacroOrder",
		"Routing",
		"AddressInfo",
		"MacroDate",
		"_AE_Backup_Information_Kind_Files",
		"IssuerLogin",
		"AddressRef",
		"CachedNumber",
		"_AE_Backup_Information",
		"Config",
		"ConfirmationFilename",
		"_AE_DatastoreAdmin_Operation",
		"__BlobInfo__",
		"PaymentMethod",
		"Site",
		"Macro",
		"_Counters",
	}
	for i, _ := range sort(kinds) {
		fmt.Printf("\"%v\",\n", kinds[i])
	}
}

func sort(pool []string) []string {
	for i := len(pool); i > 0; i-- {
		for j := 1; j < i; j++ {
			if pool[j-1] > pool[j] {
				tmp := pool[j]
				pool[j] = pool[j-1]
				pool[j-1] = tmp
			}
		}
	}
	return pool
}
