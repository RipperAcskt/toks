package netPackage

const (
	Flag = 0x7E
	Esc  = 0x1B
)

type byteNetPackage struct {
	flag       byte
	distAddr   []byte
	sourceAddr []byte
	data       []byte
}

type netPackage struct {
	flag       byte
	distAddr   string
	sourceAddr string
	data       string
}

func CreatePackage(adress, text string) []byte {
	pack := netPackage{Flag, adress, "", text}

	return pack.convert().toByteOreder()
}

func (n *netPackage) convert() byteNetPackage {
	var pack byteNetPackage

	pack.flag = n.flag
	pack.distAddr = []byte(n.distAddr)
	pack.sourceAddr = []byte(n.sourceAddr)
	pack.data = []byte(n.data)

	return pack
}

func (b byteNetPackage) toByteOreder() []byte {
	var pack []byte

	pack = append(pack, b.flag)

	for _, a := range b.distAddr {
		pack = append(pack, a)
	}
	for _, a := range b.sourceAddr {
		pack = append(pack, a)
	}
	for _, a := range byteStuffing(b.data) {
		pack = append(pack, a)
	}

	return pack
}

func byteStuffing(data []byte) []byte {
	var stuffedData []byte

	for _, b := range data {
		if b == Flag || b == Esc {
			stuffedData = append(stuffedData, Esc)
		}
		stuffedData = append(stuffedData, b)
	}
	return stuffedData
}
