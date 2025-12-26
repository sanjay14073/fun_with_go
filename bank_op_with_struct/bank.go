package main

import "fmt"

type BankAccount struct {
	name    string
	balance int
}

type Bank struct {
	branchName     string
	listOfAccounts []BankAccount
	ifscCode       string
}

func CreateAc(name string, balance int) *BankAccount {
	var x = &BankAccount{name: name, balance: balance}
	return x
}

func CreateBranch(branchName string, ifsc string) *Bank {
	var branch = &Bank{}
	branch.branchName = branchName
	branch.ifscCode = ifsc
	var emptyList = []BankAccount{}
	branch.listOfAccounts = emptyList
	return branch
}

func (br *Bank) PrintAcOfBranch() {
	fmt.Println("*****Printing All Accounts******")
	for i := range len(br.listOfAccounts) {
		fmt.Println(br.listOfAccounts[i]);
	}
}

func (br* Bank) AddAcToBank(bankAc BankAccount) *Bank {
	br.listOfAccounts = append(br.listOfAccounts, bankAc)
	return br
}

func main() {
	var bankBranch=CreateBranch("Mysuru","ABC123");
	var bankAc=CreateAc("san",200);
	bankBranch.AddAcToBank(*bankAc);
	bankBranch.PrintAcOfBranch();
	bankAc=CreateAc("maha",500);
	bankBranch.AddAcToBank(*bankAc);
	bankBranch.PrintAcOfBranch();
}