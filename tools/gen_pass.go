package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := flag.String("p", "", "Password")
	username := flag.String("u", "admin", "Username for SQL statement")
	hashInput := flag.String("hash", "", "Hash to verify (optional)")
	flag.Parse()

	// Debug: print received arguments to help diagnose shell escaping issues
	if *hashInput != "" {
		fmt.Printf("Debug: Received Hash: '%s' (Length: %d)\n", *hashInput, len(*hashInput))
	}

	if *password == "" {
		fmt.Println("Usage:")
		fmt.Println("  Generate: go run tools/gen_pass.go -p <password> [-u <username>]")
		fmt.Println("  Verify:   go run tools/gen_pass.go -p <password> -hash <hash_string>")
		os.Exit(1)
	}

	// Verify mode
	if *hashInput != "" {
		err := bcrypt.CompareHashAndPassword([]byte(*hashInput), []byte(*password))
		if err != nil {
			fmt.Printf("❌ Verification Failed: %v\n", err)
			fmt.Println("Possible reasons:")
			fmt.Println("1. The password does not match the hash.")
			fmt.Println("2. The hash string is truncated or corrupted.")
			os.Exit(1)
		}
		fmt.Println("✅ Verification Success! The password matches the hash.")
		return
	}

	// Generate mode
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	hashStr := string(hash)
	fmt.Printf("Password: %s\n", *password)
	fmt.Printf("Bcrypt Hash: %s\n", hashStr)
	fmt.Printf("Hash Length: %d\n", len(hashStr))
	fmt.Println("\nSQL Update:")
	fmt.Printf("UPDATE user SET password = '%s' WHERE username = '%s';\n", hashStr, *username)
	fmt.Println("\n⚠️  IMPORTANT CHECK:")
	fmt.Println("Please verify that your database 'password' column length is at least 60 characters.")
	fmt.Println("Run this SQL to check: SHOW COLUMNS FROM user LIKE 'password';")
	fmt.Println("If Type is varchar(32) or similar, the hash will be truncated and login will fail.")
}
