package core

import "main/global"

func Initialize() {
	global.GP_Viper = getViper()
}
