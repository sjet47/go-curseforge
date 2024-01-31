package curseforge

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/ASjet/go-curseforge/api"
	"github.com/ASjet/go-curseforge/schema"
	"github.com/ASjet/go-curseforge/schema/enum"
)

func TestGames(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.Games(
		cli.Games.WithIndex(0),
		cli.Games.WithPageSize(10),
	)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestGame(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.Game(enum.MinecraftGameID)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestGameVersions(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.GameVersions(enum.MinecraftGameID)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestGameVersionTypes(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.GameVersionTypes(enum.MinecraftGameID)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestGameVersionsV2(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.GameVersionsV2(enum.MinecraftGameID)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestCategories(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.Categories(enum.MinecraftGameID)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestSearchMod(t *testing.T) {
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
}

func TestMod(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.Mod(32274)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestMods(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.Mods([]schema.ModID{32274})
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestFeaturedMods(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.FeaturedMods(
		enum.MinecraftGameID,
		cli.FeaturedMods.WithExcludedModIDs(0),
		cli.FeaturedMods.WithGameVersionTypeID(1),
	)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestModDescription(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.ModDescription(32274,
		cli.ModDescription.WithStrippedContent(true),
	)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestModFile(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.ModFile(32274, 4759554)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestModFiles(t *testing.T) {
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
}

func TestFiles(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.Files([]schema.FileID{32274})
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestModFileChangelog(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.ModFileChangelog(32274, 4759554)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestModFileUrl(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.ModFileUrl(32274, 4759554)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestFingerprintMatchesByGame(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.FingerprintMatchesByGame(enum.MinecraftGameID,
		[]schema.Fingerprint{1})
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestFingerprintMatches(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.FingerprintMatches([]schema.Fingerprint{1})
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestFingerprintFuzzyMatchesByGame(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.FingerprintFuzzyMatchesByGame(enum.MinecraftGameID,
		cli.FingerprintFuzzyMatchesByGame.WithFolderFingerprint(schema.FolderFingerprint{
			FolderName:   "",
			Fingerprints: []schema.Fingerprint{1},
		}),
	)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestFingerprintFuzzyMatches(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.FingerprintFuzzyMatches(enum.MinecraftGameID,
		cli.FingerprintFuzzyMatches.WithFolderFingerprint(schema.FolderFingerprint{
			FolderName:   "",
			Fingerprints: []schema.Fingerprint{1},
		}),
	)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestMinecraftVersions(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.MinecraftVersions()
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestMinecraftVersion(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.MinecraftVersion("1.19.2")
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestModLoaders(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.ModLoaders(
		cli.ModLoaders.WithGameVersion("1.19.2"),
	)
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestModLoader(t *testing.T) {
	cli := NewClient(apiKey)
	rsp, err := cli.ModLoader("forge-43.3.7")
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func TestDefaultModLoader(t *testing.T) {
	InitDefault(apiKey)
	rsp, err := api.ModLoader("forge-43.3.7")
	if err != nil {
		panic(err)
	}
	printJson(rsp)
}

func printJson(v interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(v); err != nil {
		panic(err)
	}
}
