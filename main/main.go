package main

import (
	"fmt"
	"strings"
)

type Person struct {
	phone   string
	name    string
	address string
}

var phoneMap = map[string]Person{}
var tooManyMap = map[string]int{}

func CleanAddress(addr string) string {
	// Remove "wrong" characters
	addr = strings.ReplaceAll(addr, "/", "")
	addr = strings.ReplaceAll(addr, "+", "")
	addr = strings.ReplaceAll(addr, "?", "")
	addr = strings.ReplaceAll(addr, "$", "")
	addr = strings.ReplaceAll(addr, ";", "")
	addr = strings.ReplaceAll(addr, "!", "")
	addr = strings.ReplaceAll(addr, ":", "")
	addr = strings.ReplaceAll(addr, "*", "")
	addr = strings.ReplaceAll(addr, ",", "")

	// Trim all the extra whitespaces
	addrComponents := strings.Split(addr, " ")
	newAddrComponents := make([]string, 0)
	addr = addr[:0]

	for _, component := range addrComponents {
		if len(component) > 0 {
			newAddrComponents = append(newAddrComponents, component)
		}
	}

	addr = strings.Join(newAddrComponents, " ")

	// Replace _ with " "
	addr = strings.ReplaceAll(addr, "_", " ")

	return addr
}

// Returns phone number, name, address
func ProcessPhoneString(str string) (string, string, string) {
	// Find the phone
	var plusIndex = strings.Index(str, "+")
	var dashIndex = plusIndex + strings.Index(str[plusIndex:], "-")
	var limit = plusIndex + (dashIndex - plusIndex) + 13
	var phoneNumber = str[plusIndex+1 : limit]

	// Find the name
	var openBracketIndex = strings.Index(str, "<")
	var closeBracketIndex = strings.Index(str, ">")
	var name = str[openBracketIndex+1 : closeBracketIndex]

	// Find the address, remove the phone number and name from the string
	str = strings.ReplaceAll(str, "+"+phoneNumber, "")
	str = strings.ReplaceAll(str, "<"+name+">", "")
	var address = CleanAddress(str)

	return phoneNumber, name, address
}

func ProcessPhoneMap(dir string) {
	for _, phoneStr := range strings.Split(dir, "\n") {
		if len(strings.TrimSpace(phoneStr)) == 0 {
			continue
		}

		phoneStr = strings.TrimSpace(phoneStr)
		number, name, address := ProcessPhoneString(phoneStr)

		_, exist := phoneMap[number]

		if exist {
			tooManyMap[number] = 1
			delete(phoneMap, number)
		} else {
			var person Person
			person.phone = number
			person.name = name
			person.address = address

			phoneMap[number] = person
		}

		//fmt.Println(number, name, address)
	}
}

func Phone(dir, num string) string {
	if len(phoneMap) == 0 {
		ProcessPhoneMap(dir)
	}

	_, tooMany := tooManyMap[num]
	person, exist := phoneMap[num]

	if tooMany {
		return "Error => Too many people: " + num
	} else if ! exist {
		return "Error => Not found: " + num
	}

	return fmt.Sprintf("Phone => %v, Name => %v, Address => %v", person.phone, person.name, person.address)
}

func main() {
	var dr = "/+1-541-754-3010 156 Alphand_St. <J Steeve>\n 133, Green, Rd. <E Kustur> NY-56423 ;+1-541-914-3010\n" + "+1-541-984-3012 <P Reed> /PO Box 530; Pollocksville, NC-28573\n :+1-321-512-2222 <Paul Dive> Sequoia Alley PQ-67209\n" + "+1-741-984-3090 <Peter Reedgrave> _Chicago\n :+1-921-333-2222 <Anna Stevens> Haramburu_Street AA-67209\n" + "+1-111-544-8973 <Peter Pan> LA\n +1-921-512-2222 <Wilfrid Stevens> Wild Street AA-67209\n" + "<Peter Gone> LA ?+1-121-544-8974 \n <R Steell> Quora Street AB-47209 +1-481-512-2222\n" + "<Arthur Clarke> San Antonio $+1-121-504-8974 TT-45120\n <Ray Chandler> Teliman Pk. !+1-681-512-2222! AB-47209,\n" + "<Sophia Loren> +1-421-674-8974 Bern TP-46017\n <Peter O'Brien> High Street +1-908-512-2222; CC-47209\n" + "<Anastasia> +48-421-674-8974 Via Quirinal    Roma\n <P Salinger> Main Street, +1-098-512-2222, Denver\n" + "<C Powel> *+19-421-674-8974 Chateau des Fosses Strasbourg F-68000\n <Bernard Deltheil> +1-498-512-2222; Mount Av.  Eldorado\n" + "+1-099-500-8000 <Peter Crush> Labrador Bd.\n +1-931-512-4855 <William Saurin> Bison Street CQ-23071\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n"

	fmt.Println(Phone(dr, "48-421-674-8974"))
	fmt.Println(Phone(dr, "1-921-512-2222"))
	fmt.Println(Phone(dr, "1-908-512-2222"))
	fmt.Println(Phone(dr, "1-541-754-3010"))
	fmt.Println(Phone(dr, "1-121-504-8974"))
	fmt.Println(Phone(dr, "1-098-512-2222"))
	fmt.Println(Phone(dr, "5-555-555-5555"))
}
