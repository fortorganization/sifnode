package types

import "fmt"

var WhitelistStorePrefix = []byte{0x01}
var AdminAccountStorePrefix = []byte{0x02}

func GetAdminAccountKey(adminAccount AdminAccount) []byte {
	key := []byte(fmt.Sprintf("%s_%s", adminAccount.AdminType.String(), adminAccount.AdminAddress))
	return append(AdminAccountStorePrefix, key...)
}
