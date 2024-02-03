package enum

// https://docs.curseforge.com/#tocS_ModLoaderInstallMethod
type ModLoaderInstallMethod int

const (
	ForgeInstaller    ModLoaderInstallMethod = 1
	ForgeJarInstall   ModLoaderInstallMethod = 2
	ForgeInstaller_v2 ModLoaderInstallMethod = 3
)

func (mlim ModLoaderInstallMethod) String() string {
	switch int(mlim) {
	case int(ForgeInstaller):
		return "ForgeInstaller"
	case int(ForgeJarInstall):
		return "ForgeJarInstall"
	case int(ForgeInstaller_v2):
		return "ForgeInstaller_v2"
	default:
		return "Unknown"
	}
}
