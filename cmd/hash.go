/*
Copyright © 2021 Piyush Saurabh admin@roguesecurity.in

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"crypto/sha256"
	//"crypto/sha384" not available in standard crypto
	"crypto/sha512"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/sha3"
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Returns the hash of a given input string",
	Long: `Returns the hash of a given input string. It accepts various inputs like hashing algorithm, and plain text. In the output it will return the hexadecimal encoded hash of the input. 
	
	For example: cryptoctl hash --algo="sha3-256" --input="hello"
	
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hash called")

		// Get the arguments
		algo, _ := cmd.Flags().GetString("algo")
		input, _ := cmd.Flags().GetString("input")

		processRequest(algo, input)

	},
}

func init() {
	rootCmd.AddCommand(hashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	hashCmd.PersistentFlags().StringP("algo", "a", "sha2", "Algorithm for hashing. Possible values: sha3-256, sha-256")

	hashCmd.PersistentFlags().StringP("input", "i", "", "Enter the text you want to hash")
}

// Process the incoming request and call the necessary function
func processRequest(algorithm, input string) {

	//fmt.Printf("Input is %s\n", input)

	switch algorithm {
	case "sha3-256":
		//fmt.Println("Selected algorithm is sha3-256")
		getSHA3(input)
	case "sha-256":
		getSHA2(input)
	default:
		fmt.Println("No algorithm matched. Selecting default algorithm sha2")
	}

}

// SHA3 implementation
// Ref: https://pkg.go.dev/golang.org/x/crypto/sha3#Sum256
func getSHA3(input string) {

	// SHA3-256 bits
	hash256 := sha3.Sum256([]byte(input))
	fmt.Printf("SHA3-256 output: %x\n", hash256)

	// SHA3-384 bits
	hash384 := sha3.Sum384([]byte(input))
	fmt.Printf("SHA3-384 output: %x\n", hash384)

	// SHA3-512 bits
	hash512 := sha3.Sum512([]byte(input))
	fmt.Printf("SHA3-512 output: %x\n", hash512)
}

// SHA2 256 implementation
// Ref: https://golang.org/pkg/crypto/sha256/#Sum256
func getSHA2(input string) {
	hash256 := sha256.Sum256([]byte(input))
	fmt.Printf("SHA2-256 output: %x\n", hash256)

	/*
	   Not available in standard crypto library
	   	hash384 := sha384.Sum384([]byte(input))
	   	fmt.Printf("SHA2-384 output: %x\n", hash384)
	*/

	hash512 := sha512.Sum512([]byte(input))
	fmt.Printf("SHA2-512 output: %x\n", hash512)
}
