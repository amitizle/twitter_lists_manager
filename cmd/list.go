package cmd

import (
	"github.com/amitizle/twitter_lists_manager/internal/printer"
	"github.com/amitizle/twitter_lists_manager/internal/twitter_client"
)

// ListLists lists the lists owned by the authenticated user
func ListLists(client *twitter_client.Client) {
	ownedLists, err := client.GetOwnedLists(nil)
	if err != nil {
		printer.Fatalf("Error: %v", err)
	}
	for _, l := range ownedLists {
		printer.NoColorf("Name: %s, Slug: %s, Description: %s", l.Name, l.Slug, l.Description)
	}
}
