package main

import (
	"flag"
	"fmt"
	"github.com/danyelmorales/weak_encryption_lib/pkg/cipher"
	"github.com/danyelmorales/weak_encryption_lib/pkg/symbol"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetOutput(ioutil.Discard)

	// argparse
	showLogs := flag.Bool("l", false, "show logs")
	doEncrypt := flag.String("e", "", "encrypt text")
	doDecrypt := flag.String("d", "", "decrypt text")
	//useKey := flag.Int64("k", 3, "the key to be used")
	showBadGoodKeys := flag.Bool("gbk", false, "Show good and bad keys within current alphabet")

	flag.Parse()

	if *showLogs {
		log.SetOutput(os.Stdout)
	}

	// start
	theAlphabet := symbol.StrToSymbolArray("$!?¿_-.;·#@ABCDEFGHIJKLMNOPQRSTUVWXY0123456789Zabcdefghijklmnopqrstuvxyz")
	weakCipher := cipher.New(theAlphabet)

	log.Printf(strings.Repeat("·", 8))
	log.Printf("Starting cipher suite: weakCipher")
	log.Printf("Key: %d", weakCipher.Key)
	log.Printf("Modulus: %d", weakCipher.Modulus)
	log.Printf(strings.Repeat("·", 8))

	var result string
	if *doDecrypt != "" {
		result = weakCipher.DecryptFromStr(*doDecrypt)
		fmt.Println(result)
	}

	if *doEncrypt != "" {
		result = weakCipher.EncryptFromStr(*doEncrypt)
		fmt.Println(result)
	}

	if *showBadGoodKeys{
		bad, good := weakCipher.GoodBadKeys()
		fmt.Printf("bad:%v \n good: %v", bad, good)
	}

}
