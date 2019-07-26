package main 

import (
	"pack"
)


func main() {

	pack.ShowVersion()
	pack.Prepare()
	pack.DoPack()
	pack.CleanUp()

}
