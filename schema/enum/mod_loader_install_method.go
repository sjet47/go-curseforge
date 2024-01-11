package enum

// https://docs.curseforge.com/#tocS_ModLoaderInstallMethod
type ModLoaderInstallMethod int

const (
	ForgeInstaller    ModLoaderInstallMethod = 1
	ForgeJarInstall   ModLoaderInstallMethod = 2
	ForgeInstaller_v2 ModLoaderInstallMethod = 3
)
