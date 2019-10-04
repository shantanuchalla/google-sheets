package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/api/sheets/v4"
)

func main() {

	credFilePath := flag.String("cred-file", "/Users/shantanuchalla/code/resources/googleapis/credentials_sheets_practice.json", "path to credentials file")
	tokenFilePath := flag.String("token-file", "/Users/shantanuchalla/code/resources/googleapis/token_sheets_practice.json", "path to token file")
	spreadsheetID := flag.String("spredsheet-id", "1dFXie_b-9-sodRajKEBPT89yEy8FQBZV5gGNsxa3hWk", "id of spredsheet to append data to")
	sheet := flag.String("range", "PD Analysis", "sheet to append to")

	flag.Parse()

	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"
	
	values := &sheets.ValueRange{
		Values: [][]interface{} {
			[]interface{} {"abc", "xyz"}, 
			[]interface{} {"123", "456"},
		},
	}

	ctx := context.Background()

	client := getClient(credFilePath, tokenFilePath)

	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := sheetsService.Spreadsheets.Values.Append(*spreadsheetID, *sheet, values).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", resp)

}
