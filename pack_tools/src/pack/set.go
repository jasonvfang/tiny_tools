package pack

import (
	"log"
)

const (
	ToolVersion = "v0.1"
	RootOtaDirName = "sgw_avs"
	CmbsPrefixString = "Alexa_DCX81_MOD_"
	HandsetPrefixString = "AXH01H"
	BaseKernelPrefixString = "boot.img."	
	BaseRootfsPrefixString = "rootfs.img."

	ConfigCmbsImgDir = "release/cmbs/"
	ConfigHandsetImgDir = "release/handset/"
	ConfigReleaseR18ImgDir = "release/R18/release/"
	ConfigDebugR18ImgDir = "release/R18/debug/non-secure/"

	ConfigOutputOtaPackageDir = "out/"

	HelperDirPath = "helper/"
	HelperToolDirPath = "helper/tools/"
	HelperToolOpenSSL = "helper/tools/openssl"
	HelperToolPackImg = "helper/tools/pack_img"
	HelperToolRsaFile = "helper/tools/rsa_file_x86"
	GenVersionScript = "gen_version.sh"

)


var Region string
var ReleaseVersion int
var IsForDebug bool
var DstOtaPkgPath string


func ShowVersion(){
	log.Println("SGW OTA pkg Pack Tool ", ToolVersion)
}

