package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/FactomProject/factomd/common/directoryBlock"
	"github.com/FactomProject/factomd/common/entryBlock"
)

// HelpText is the help string
var HelpText string

// Indent is whether or not to indent Json return
var Indent bool

func main() {
	var (
		indent = flag.Bool("i", false, "Indent Json")
	)

	flag.Parse()
	Indent = *indent

	addHelp("|---[type]---|", "|---[text]---|")
	addHelp("eblock", "Entry-block serialized as hex string")
	addHelp("dblock", "Directory-block serialized as hex string")

	os.Args = flag.Args()

	if len(os.Args) < 2 {
		fmt.Println(help())
		return
	}

	// Serialize type
	switch os.Args[0] {
	case "eblock":
		jsonStr, err := serializeEblock(os.Args[1])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(jsonStr)
		}
	case "dblock":
		jsonStr, err := serializeEblock(os.Args[1])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(jsonStr)
		}
	default:
		fmt.Printf("%s is not a supported type.\n", os.Args[0])
		fmt.Println(help())
	}

}

func serializeEblock(hexString string) (jsonResp string, err error) {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	eb := entryBlock.NewEBlock()
	newdata, err := eb.UnmarshalBinaryData(data)
	if err != nil {
		return "", err
	}

	if len(newdata) > 0 {
		return "", fmt.Errorf("%d bytes remain, extra hex characters were given", len(newdata))
	}

	if Indent {
		var dst bytes.Buffer
		jB, err := eb.JSONByte()
		err = json.Indent(&dst, jB, "", "\t")
		if err != nil {
			return "", err
		}

		return dst.String(), nil
	}

	return eb.JSONString()
}

func serializeDblock(hexString string) (jsonResp string, err error) {
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	db := new(directoryBlock.DirectoryBlock)
	newdata, err := db.UnmarshalBinaryData(data)
	if err != nil {
		return "", err
	}

	if len(newdata) > 0 {
		return "", fmt.Errorf("%d bytes remain, extra hex characters were given", len(newdata))
	}

	if Indent {
		var dst bytes.Buffer
		jB, err := db.JSONByte()
		err = json.Indent(&dst, jB, "", "\t")
		if err != nil {
			return "", err
		}

		return dst.String(), nil
	}

	return db.JSONString()
}

//
func addHelp(command string, text string) {
	HelpText += fmt.Sprintf("|   %-30s%s\n", command, text)
}

func help() string {
	return HelpText
}
