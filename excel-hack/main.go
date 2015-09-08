package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

type StringSet struct {
	set map[string]bool
}

func NewStringSet() *StringSet {
	return &StringSet{make(map[string]bool)}
}

func (set *StringSet) Add(s string) bool {
	_, found := set.set[s]
	set.set[s] = true
	return !found //False if it existed already
}

func (set *StringSet) Get(s string) bool {
	_, found := set.set[s]
	return found //true if it existed already
}

func (set *StringSet) GetAll() []string {
	var all []string
	for key, val := range set.set {
		if val {
			all = append(all, key)
		}
	}

	return all
}

func (set *StringSet) Remove(s string) {
	delete(set.set, s)
}

func (set *StringSet) isEmpty() bool {
	for _, val := range set.set {
		if val == true {
			return false
		}
	}
	return true
}

func main() {
	subjectSet := NewStringSet()
	excelFileName := "/Users/5iM/workarea/PiC/KS5_Subject_Level_Data_2012-England_2.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, sheet := range xlFile.Sheets {
		//fmt.Println("sheet is:", sheet.Name)
		if sheet.Name == "A level" {
			fmt.Println("fount sheet A level")
			for _, headCol := range sheet.Rows[2].Cells {
				fmt.Println("Head col name is: ", headCol.Value)
			}

			//for i := sheet.R
			for ind, row := range sheet.Rows {
				if ind >= 3 {
					subjectSet.Add(row.Cells[5].Value)
					//fmt.Println("subjectSet is:", row.Cells[5].Value)
					//fmt.Println("subjectSet is:", subjectSet.Get(row.Cells[5].Value))
				} /*
					if ind == 10 {
						break
					}
				*/
			}
			allSubjects := subjectSet.GetAll()
			for _, val := range allSubjects {
				fmt.Println("the subject are:", val)
			}

			fmt.Println("Max subjects are: ", len(allSubjects))
		}
		/*
		   for _, row := range sheet.Rows {
		       for _, cell := range row.Cells {
		           fmt.Printf("%s\n", cell.String())
		       }
		   }
		*/
	}
}
