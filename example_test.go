package curseforge

import (
	"fmt"
	"os"

	"github.com/ASjet/go-curseforge/schema/enum"
)

var (
	apiKey string
)

func init() {
	apiKey = os.Getenv("CURSE_FORGE_APIKEY")
}

func Example_searchMod() {
	cli := NewClient(apiKey)

	rsp, err := cli.SearchMod(enum.MinecraftGameID,
		cli.SearchMod.WithGameVersion("1.19.2"),
		cli.SearchMod.WithModLoaderType(enum.ModLoaderForge),
		cli.SearchMod.WithSearchFilter("JourneyMap"),
		cli.SearchMod.WithSortField(enum.ModsSearchSortFieldPopularity),
		cli.SearchMod.WithSortOrder(enum.SortOrderDescending),
		cli.SearchMod.WithIndex(0),
		cli.SearchMod.WithPageSize(1),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ModID: %d\nName: %s\nSummary: %s\n",
		rsp.Data[0].ID, rsp.Data[0].Name, rsp.Data[0].Summary)

	// Output:
	// ModID: 32274
	// Name: JourneyMap
	// Summary: Real-time mapping in-game or your browser as you explore.
}

func Example_getLatestModFile() {
	cli := NewClient(apiKey)

	rsp, err := cli.ModFiles(32274,
		cli.ModFiles.WithGameVersion("1.19.2"),
		cli.ModFiles.WithModLoader(enum.ModLoaderForge),
		cli.ModFiles.WithIndex(0),
		cli.ModFiles.WithPageSize(1),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("FileName:\t%s\nFileDate:\t%s\nDownloadURL:\t%s\n",
		rsp.Data[0].FileName, rsp.Data[0].FileDate, rsp.Data[0].DownloadURL)

	// Output:
	// FileName:       journeymap-1.19.2-5.9.7-forge.jar
	// FileDate:       2023-05-11 15:42:02.777 +0000 UTC
	// DownloadURL:    https://edge.forgecdn.net/files/4532/924/journeymap-1.19.2-5.9.7-forge.jar
}
