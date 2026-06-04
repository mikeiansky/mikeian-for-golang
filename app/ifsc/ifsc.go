package main

import (
	"fmt"

	ifsc "github.com/razorpay/ifsc/v2/src/go"
)

// todo: change funcs not required to lower case.

func main() {

	fmt.Println(ifsc.Validate("IOBA0003123")) // Returns true
	fmt.Println(ifsc.Validate("ZSBL0000331")) // Returns false

	ifsc.ValidateBankCode("PUNB") // Returns true
	ifsc.ValidateBankCode("ABCD") // Returns false

	code, _ := ifsc.GetBankName("PUNB") // Returns "Punjab National Bank", nil
	fmt.Println(code)
	fmt.Println(ifsc.GetBankName("ZSBL")) // Returns "", errors.New(invalid bank code)
	ifsc.GetBankName(ifsc.HDFC)           // Returns "HDFC Bank", nil

	ifsc.GetBankDetails("PUNB")
	// or
	ifsc.GetBankDetails(ifsc.PUNB)

	/* Returns
		(*ifsc.Bank){
		Name	  : "Punjab National Bank",
		BankCode  : "024",
		Code	  : "PUNB",
		Type	  : "PSB",
		IFSC	  : "PUNB0244200",
		MICR      : "110024001",
		IIN       : "508568",
		APBS      : true,
		AchCredit : true,
		AchDebit  : true,
		NachDebit : true,
		Upi       : true
	}), nil
	*/

	ifsc.LookUP("KKBK0000261")

	/*
		Returns
		(*ifsc.IFSCResponse)({
		 Bank	  :  "Kotak Mahindra Bank",
		 Branch	  :  "GURGAON",
		 Address  :  "KOTAK MAHINDRA BANK LTD. UNIT NO. 8&9, SEWA CORPORATE PARK, MG ROAD, REVENUE STATE OF SARHAUL TEHSIL, DISTT,- GURGAON- 122001",
		 Contact  :  "4131000",
		 City	  :  "GURGAON",
		 District :  "GURGAON",
		 State	  :  "HARYANA",
		 IFSC	  :  "KKBK0000261",
		 BankCode :  "KKBK"
		}), nil
	*/
}
