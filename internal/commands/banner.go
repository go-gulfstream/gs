package commands

import "fmt"

var banner = []byte(`

/***
 *      ________      .__   _____         __                                 
 *     /  _____/ __ __|  |_/ ____\_______/  |________   ____ _____    _____  
 *    /   \  ___|  |  \  |\   __\/  ___/\   __\_  __ \_/ __ \\__  \  /     \ 
 *    \    \_\  \  |  /  |_|  |  \___ \  |  |  |  | \/\  ___/ / __ \|  Y Y  \
 *     \______  /____/|____/__| /____  > |__|  |__|    \___  >____  /__|_|  /
 *            \/                     \/                    \/     \/      \/ 
 */

`)

func drawBanner() {
	fmt.Println(string(banner))
}