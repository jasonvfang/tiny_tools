package pack

import (
	"log"
)

const (
	ToolVersion = "v0.1"
	PackTargetName = "sgw_avs"
	RootOtaDirNam = PackTargetName
	CmbsPrefixString = "Alexa_DCX81_MOD_"
	HandsetPrefixString = "AXH01H"
	BaseKernelPrefixString = "boot.img."	
	BaseRootfsPrefixString = "rootfs.img."
	ScriptName = "update.sh"

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

	HelperAESKeyFile = "helper/ota_key/aes.key"
	HelperRsaPrivateKeyFile = "helper/ota_key/rsa_2048_private.key"
	HelperRsaPublicKeyFile = "helper/ota_key/rsa_2048_pub.key"

	ReleaseLinkplaySdkVerFile = "release/linkplaySdk/MVver"
)


var Region string
var ReleaseVersion int
var IsForDebug bool
var DstOtaPkgPath string


func ShowVersion(){
	log.Println("SGW OTA pkg Pack Tool ", ToolVersion)
}

