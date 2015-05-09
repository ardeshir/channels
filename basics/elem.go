package main 

import "fmt"

func main() {

	elem := make(map[string]string)

	elem["H"] = "Hydrogen"
	elem["He"] = "Helium"
	elem["Li"] = "Lithium"
	elem["Be"] = "Beryllium"
	elem["B"] = "Boron"
	elem["C"] = "Carbon"
	elem["N"] = "Nitrogen"
	elem["O"] = "Oxygen"
	elem["F"] = "Fluorine"
	elem["Ne"] = "Neon"

  for k, v := range elem {
  	fmt.Printf("key: %s,  value: %s\n", k, v)
  }

elements := map[string]map[string]string{
"H": map[string]string{
		"name":"Hydrogen",
		"state":"gas",
},
"He": map[string]string{
		"name":"Helium",
		"state":"gas",
},
"Li": map[string]string{
		"name":"Lithium",
		"state":"solid",
	},
"Be": map[string]string{
		"name":"Beryllium",
		"state":"solid",
	},
"B": map[string]string{
		"name":"Boron",
		"state":"solid",
	},
"C": map[string]string{
		"name":"Carbon",
		"state":"solid",
	},
"N": map[string]string{
		"name":"Nitrogen",
		"state":"gas",
	},
"O": map[string]string{
		"name":"Oxygen",
		"state":"gas",
	},
"F": map[string]string{
		"name":"Fluorine",
		"state":"gas",
	},
"Ne": map[string]string{
		"name":"Neon",
		"state":"gas",
	},
}
if _, ok := elements["Li"]; ok {

   for k, _ := range elements {
	fmt.Printf( "%s - %s : %s\n", k, elements[k]["name"], elements[k]["state"])
   }

}


}