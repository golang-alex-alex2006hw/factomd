package templates

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

type staticFilesFile struct {
	data  string
	mime  string
	mtime time.Time
	// size is the size before compression. If 0, it means the data is uncompressed
	size int
	// hash is a sha256 hash of the file contents. Used for the Etag, and useful for caching
	hash string
}

var staticFiles = map[string]*staticFilesFile{
	"general/footer.html": {
		data:  "{{define \"footer\"}}\n    </body>\n</html>\n{{end}}\n",
		hash:  "239cf4d42e3fd75d2fa4395fb31f7c6679f9c248afe058dc90d9b37280bc5aa2",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  0,
	},
	"general/header.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x9cT͎\xeb6\x0f\xdd\xcfS\xf0\xd3\xe6\xdbDq\x8b\xa2E1\x13g\xd3\xe2\x16\xdd\x15\xb8\x0fP0\x12m\xb3W\x96|)ʙt0\xef^8N\x1aw\xe2\xb9\xfdYٔt\x0e\x0f\x0f)\xbd\xbcxj8\x12\x98\x8eГ\x98\xd7ׇ\xdd\xff|rz\x1a\b:\xed\xc3\xfea7}\xc0\x05̹61\xd9߲\x81\x80\xb1\xad\rE\x03\x9e\xa56A\xc5\xec\x1f\x00\x00v\x13\xcf\xfc{\x0e{R\x04סd\xd2\xda\x14m\xec\xf7\xe6\xedv\xa7:X\xfa\\x\xacͳ-h]\xea\aT>\x042\xe0RT\x8aZ\x1b\xa6\x9a|Kw\xe8\x88=\xd5fd:\x0eIt\x018\xb2\u05ee\xf64\xb2#{\x0e6\xc0\x91\x951\xd8\xec0P\xfd\xf5\xf6\xab%]\xe0\xf8\t\x84Bmr\x97D]Q`\x97\xa2\x81Ɍ\xdap\x8f-U\xcfv^넚i\xad\xad\xa6\xb8jp\x9c\xbe[viI\xa9\xac\x81\xf6?\xa4\xa8\x92\x02\xfc\x82\x91\x02X\xf8\x80NS\x0f\x1fR\x89\x1e\x95S\xdcU\xf3\xc1U-z\n\x94;\"\xbd&u9W͟\xd8m\xcfq\xebr6\xff\n\x1d\xd5\xe2\x91r\xea\xe9\xbf\xe0q\x18\x16\x90]uk\xfa\xee\x90\xfc\xe9\xc6t\x9b\t\x92\x05\xbf\xe7\xf1:P\x92\x8e\x8b\xccow\x03JK\xf6\x1b\xe8\xc9s\xe9\xed\xb7\xe0R(}\xcc\x10R\x9b\xde\xe0\xceX\xbc\xa8\xac\xcc~\xc7}\vY\xdcܥ\xe6칽\xf9\xf6k\x97\x84\u007fOQ1l\xf3\xd8\x1a\xc0\xa0\xb5\xb9k\x8d\xd9\xef*|#\xb0\xf2<\xfe\x9d\xe6\xef\xee5gBq\x1d\\\x94̑\x9d\xc6\x159\x92\xacU\xb3`\xe58\x14\xb5\xad\xa42\xac\x9c<\x9f\xce\x03ƕ\xe36\xe0\x81\xc2\xe4\xc7u\xb3Ah\xf0\"\xc0\x00\n\xa3\xed\xd8{\x8a\xb5Q)4\x95\xcc\xfb]5\xf1\xbd\x93\xeaL\u007f\xb9\x17W\x9e\xf9\"~6k\x12\x1a\xa6\xe0\r\xb0\x9f\x92/\xca70\x04tԥ\xe0Ij\xf3qv\b\xa1#n;\xdd\x00z/\x94\xf3\x06T0ftSG\xe0\xe7\x1f7@\xea\xb6\xdb\xf7\x8cX\xb7\xcd\x1e\x8a\xea\xd4\xd0U\xd0}Y\xe5г\xaeh\xb6םK\x8a\x99\x16<\xca'\x03#\x86B\xb5\xf9)\xbd\xa7\xed~vn˰>\x01\xe7;X\x9b!e\x9e\xea\u007f\xc4CN\xa1(=\x9d\x1f\xb5G\x8e\x1d\t\xeb?0\xe3\xafe\x90H\x12p\x18B*\n\x18H\x14\x94\x9e\xd5:\x8aJb\xaey\xd3H҄t|\x9cg\xe4\xc9s\x1e\x02\x9e\x1ec\x8a\xf4d\xf6\x1f\x93\xc8i\x03\xa7T\xe4:\xe1\x9e}\xfc\xbf\x82\x90\x16\x89\x80\xf1\x04B\xb9\x04\xcd\xdb/\x96\xff\xa5K\xb6\b\xe7\xe7fzO^^(\xfa\xd7ׇ?\x02\x00\x00\xff\xff\x93P.\x04\xc6\x06\x00\x00",
		hash:  "c0b8bc411d656fcf964e2c07939211bc9c6ff43d9be8dcd6f5509d89c7981384",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  1734,
	},
	"general/scripts.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x8c\xcb;\x0e\xc3 \f\x80ὧ\xb0\xd8\x13.\x90\xe6.\x16\x0f\xd5H1\x14\x9b>\x84\xb8{\x87l\x1d\"\xc6_\xfa\xfe\xde}\x88\xc4\x01\x8c\xb8JEŌq\x03\x00\xd8\xce\x06\xa9\xeen\x92\xd8W`\x9f\xabM\xcf\x16\xeawMb\xf6͞d\xbf\xf2\xef\a\xeaB\\\x9a\xce?17\xf6\xa8\x94y=\x88'>,eBEt\x9a\x0f\xbf`\xc2\xcf\x1f\xef=\xb0\x1f\xe3\x17\x00\x00\xff\xff\xb1\x1b \xa8\r\x01\x00\x00",
		hash:  "e084eeb01388db75a5076e1e487fc761e3b7ece8ee41f6a104896b630431b97c",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  269,
	},
	"index/controlPanelScripts.html": {
		data:  "{{define \"controlPanelScripts\"}}\n\t<script src=\"js/controlPanel.js\"></script>\n{{end}}",
		hash:  "5fc312858cfbfba3160a20303376141a6fcfb759d1291c25dcb27fe9aecffae0",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  0,
	},
	"index/datadump.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xe4Y\xdfo\xdb6\x10~\xb6\xff\n\x8e{I\x80\xc9F\x9a\xfe\x00\x1aI\xc0\x96\x16C\x81\xb4+\xeaa\xef\xb4x\x92\x88P\xa4FRN\x8c\xc0\xff\xfb@J\xb2i[JRU\xc9\x024/\xa2x\xe4\xdd\xf1\xbe\xef\xa3I\xe5\xee\x8eB\xca\x04 L\x89!\xb4*J\xbc\xd9LC\r\x89aR F#g\xf8`\r(\xe1D\xeb\b\xe7\x8c\x02\x8e\xa7\b!\x14R\xb6j\xbb\x95\xbc\xc1\xf1tr\u061dH^\x15B\xb7&g\xce\xcf\xe2\xbfA\x15L\x10\x8e\xac{ĊR*\x13\xce\xf3\xb3x:\x99L\u008a\xb7\xd3\rYj\xec\x06\x05\xb6\xe92\x82[R\x94\x1c\\\av\x13&!g\xfe\x8c\xc00\xc3\x011\x1d\x90İ\x15\xe08$(W\x90F\xf8W\xbb\xc83\x8c\x88b$\xd0\xc0!1@#lT\x058^TEA\xd4:\x9c\x938\x9csv\x8f\xefC\x8f\xafp\xfcU\xc9\x04\xb4FWL\x9b\x01\x1eέ\a&\xccgR\x0e\x98\xfd\x1a\xc7\vP+Pz7\x19\x1d\xfc=\xca\xd1[\x1c\u007f\xe45\xfcz@\x1eop|)\x858r0(\x97w8\xbe\x92\x19Z\x801Ld\xfb\xe9\x84\xf3\x8a\xd7\r\x8fl\xceU\"\x85\x01a<ִ]\xdd\xd49\x9c_\x12\x01\xdc\xe3N\xad\x02ǚz\xc6!?\x915\x06\x8fd\xead\x12\xfe\x12\x04\xde\xf2\xdb\xc9\xe81\xc4]\xe4R\x99>\xf2\xe6NCM\x8d\x82`\x1bo`\xaco\xe4\x06\xc7\xdf\xc8M?\x88\x1d`vD8Z\xc3Z$Ld}\xab\xa8\xad\xfb\xd4ۡ\xdd\xe0\xd5B\xe2\x90%L\x80ڕ\x97\x15\x99\xb3\xa7\x15\xe7:Q\x00\"\x90\xa5\xa5\xe3v\a#K-ye \xe8\x18\xa2U\x12aVd\xf3\x9dm\xa6W\x19\x8e\xc39+2\x1b\xc4_y\r\xa6\x81[C\x14\x10D\x99&K\x0e\x14\xe9\x128OrH\xae#\x9c\x12\xae\x01?\x8eb5\xbeq8o]\xfa8\x8e\x17\xc6A\xeb\x05\xe9GvXLoA\r\xd8{\xd1\x1aH)[5\x12\xf4\x9a\x9dj\xdc\xf9{\xe5\xff\x90\xd4\b\xfc\x98\x18\x87\x8b\xc3\xee\xf7=$\xbe\xac\x94\x02a\xc6V\xceW\x05\xab\x9e\x88\xd64v\xb8/p۷\xd7X\xd3.\xdc\xf4\xe5\x89t|\xc18\xb8\x9fG4\xae\xf0\xcf\x13\xcaQj,q\x9e\xbf\x1cq\xf6\xf0v\xef\xe7\xec' \xedhо\x1e\xe9\xec3\x10\xd1\xdf+\x93\xf7@jMR1\xc3\xe0\xe0\xbc\xfaP\xb4\xc3\x18\x9f\xa8=5\xc6\xee1\x82\xbb\xcf\xeb/\xd2^\x94>\xaf\x91m\xbc\\\xda\xed\tvL\xfe9Ў\t8\xf9\xe1M\xabA\xea\t<\xb7\xa0\x8d\xa4\x9a7\xff\xafj.\xa5\xf8\x9e\xad\xf0;\t\xbepGտL\x0e\xea'$wSۧ \xe1\xe2\xf0\x0e\xd0G\xc1\xa3\xd3\xc0\x03t|\x8b\xbbO\x12C\xd99\xec\x98\xd9\xcf\xd7퇇\x1eʶ\xf6\xd1o\xa4\xac\xa881@\x1fJ\xe0\x1fi\xf6\xee\xa6\xd3\xeer:\xfawZz\xd5л\x90\xa7\xd5\xc7\xc8\xe7\xca.\xa1\xecj\xfaL7\xcec0\x1f\fܧ\xa7a2{\x87\xefA\u007f\xefkio\x11\xbc\x91\xba \x9c\a2M5\x98\xe0\fկo\x91\xf7]\xb5\xd9v\xf2\xf3\xf82'\"\x03\xc4e\x86\xf4\xf6\xa3Y~\xbe\x1b\xe3\xf9e\xa2\xacL\x90)Y\x95;'\x93P\x97Dt\f\t8Y\x02\xc7q\xb8\xd5UJPJ\x82\x94\xd9\xed`wQ\xcc\x19\xa5 Z\xbd\x84s\x16\x87s\xebҋ\xe0\xbc\"\xb3.!\xc2\x1a\x88Jr\x8c\x04) \xc2\xff\xe2\xae\xc8)\x03N\xeb\xea\xa6$1\xb2\b\xb8\xcc\x02\r\x06\xa3\x92\x93\x04r\xc9)\xa8\b/\xc0 \x93\u05ebW\x90\xc1\xedl6\xf3Wֽ\xf6`Y\x19#\x857\xf0 \xc3jY0s\x1c\xbe\xe9n\x1c\xd6N\x10%\xea\x1a\xa3\x15\xe1\x15D\xf8O\xe9G\xdfm\xde\xfd\xfbw\x1f\x05\xbc\xb8\xa0\x94T(!\x9c\xcb\xca \xc2A\x19d\xcb\x1f$ \f(\x8c\xb4Ys\x88\xb0\\\x81J\xb9\xbcy_\x03rA\x99.9Y\xbf\x17R\xc0\x05\x8e?\x80`@\x87\xa7\xa1\xab\xc4}\xebn\x13i\xdf\a\xa4r\xe5\x91\xd56\xf6*\xe5\xbf\xdc/R\u007fp\xdbh\x9e\xe1\xbc\xf9\u007fF<\x9d\x86:Q\xac4\xb5\x9b\x15Q\xa8\xc6:BT&U\x01\xc2\xcc20\x1f9\xd8\xe6\x1f\xebO\xf4\xe4\x90s\xa7\x17n\xaa\x9b6#\x94~\\\x810WL\x1b\x10\xa0N\xf05\xac\xab\x12\xff\x86\xd2J\xb8\x90'`\xed\xa7\xe8n\x9b\xb7똕\xca=?@J*nN\x1a\xaf\xces\x8a\xeaI\xb3kX_J\n(\x8a\"tv\xee\xfb\xb0\u007f\x8fJ\xb8f\xe9\xe9,\xe1,\xb9\xf6\xa3l\\ksza\xabӔdzw\a\x82n6\xff\x05\x00\x00\xff\xff\xefi!w\x15\x1a\x00\x00",
		hash:  "e7bc92d7f43e7325725c0342ea0358bbdfcac80f7790d4cbbace0a0d89941c9a",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1529338153, 0),
		size:  6677,
	},
	"index/index.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xfft\x8fA\xaa\x021\f\x86\xd7\uf762v_O .\x04\xf7\x03z\x81\xd0d\xb4\xd0&\xa5͈C\xe9݅q!:\xcc6\xf9\xbe\xe4\xff[C\x1a\x03\x93\xb1\x81\x91\x9e\x03\xdc\xc8\xf6\xfe\xffךR\xca\x11\x94\x8c\xbd\x13 \x95e|\xd89gN\x82\xb3q\xee\xf8M->\xc3c\xa5G\xf1\x10\xaf\x92\xad\xd9\xff\xae\xb4\x00W\xf0\x1a\x84\xeb\x94\x12\x94ye#(\xe0\x94\xf2\xe7\xfd\x99q#B\xf5%d\xad\xab\x1b^X\x8b\xc4\x01\x98\xe2e\x83\x19E\xf4]\xb25b\xec\xfd\x15\x00\x00\xff\xff)\xb2x\xeb\x1a\x01\x00\x00",
		hash:  "cc3a3da9da68933167db530679efab736e844d7d525da9b7f55fe484cfc74a3f",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  282,
	},
	"index/indexnav.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x84\x8e\xb1j\xc40\f\x86\xf7<\x85\xd0n2ui\x1dO]:\\)\xf4\t\x94Xw\b\x14;\xd8J\x9a\x12\xf2\xee\xe5B\x0fz]nѠ\x9f\xff\xfb\xfem\x83\xc8gI\f()\xf2\x9ah\xc1}o|\x94\x05\x06\xa5Z;,\xf9\v\xa1ڷr\x87#\x95\x8b$\xd7g\xb3<>\xc3Ӵ\xbe`h\x00\x00\xfc\xac\xb7\x82Q_\xe1zܐ\x93\x95\xacn\xa2Ċ\x10\xc9\xc8\x1d\xa9\xc4\x0ey\xa5qR>\x1e\xbf\x90\x03\xa4\xf2\x17\xe4LL\x19\xa4:\x1aL\x16ƣ{\xdb\xeaF\x92\x84\xc1\x13P\x11r\x95\x95\a\xe3ء\x95\x991\x9cH\x12|\x1a\xd9\\\xe1\x83.\xec[\n\xbeUyd3\xea\xffo\xbf\xb7\xe6\xc2Wk8\xe5\xc2\xf0\xcaF\xa2\x1c\xe1=G\x86\xb7t\xcee$\x93\x9c\xeeu\xbe\x9d54\xbe\x8d\xb2\x84f\xdb8\xc5}\xff\t\x00\x00\xff\xff\xa7\xffC\x19\u007f\x01\x00\x00",
		hash:  "dbe5cd4fb16f7f214469c016e8613fc72fd7f0bba689d5a13a5a99fac9615dc8",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  383,
	},
	"index/localTop.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xdcX_o\xdb6\x10\u007fϧ`\t\fH\x80\xb1\xb2\xb3`\v\\\x89@\x92\xa2ـ\xb5(\xe6b\xc0\x1ei\xf1l\x13\x95H\x8d<96\f\u007f\xf7\x81\x94d'\xaem\xc9N݇\xe5%6y\xf7\xbb\xbb\xdf\xfd#\xbc\\J\x18+\r\x84f&\x15\xd9\x17S\xd0\xd5\xea\x82\xd4\u007f\xb1\x83\x14\x95\xd1Dɤ\x12\xa0|}\x19\x04\xa4\x9a\x914\x13\xce%Ԛ\xa7\xad\xdbm\x89\xd4de\xae\xdd\x0e\xa9\xcaX.\xb2\x8cǮ\x10\x95A\av\x06\x969\x14X:\xca\xe3\xc8\xdf\xf8\u007fAn7ƴO\x1c.2H蓒8\x1d\xf4{\xbd\x9f\xdeQ\xfe\x8f)-\xf9d$4Vf\xcb\xe5ۿ\xc1:e\xf4j\xd5@Vw\r\xc083\x02\aVM\xa6\xf8\x8e\xf2G\x85\xe4\xbeT\x99\x1c\x90\xe5\xf2\xed\xa3\xc2\xf0\xe5\x99n4\xeds\xb2۩7\x8c\xd5v\x89\xd1i\xa6ү\t\xd50G\xef\xd0\xe5\x15\xe5\xb1\xe0\x9f`\x8e\xc1\xc18\x12\xeb\x10\u007fn\xbc}(\xad\x05\x8d\x83\r7iu´\x91\xc0t\x99\x8f\xc0R\xdeۢ\x880\xb6#!\x91T\xb3\xad,\xee8:*\xb1\x99\xb0\x13`\xbf\x90\x1c\xa4*svC\x82}ֿ&-)\u007f\x86\x91\x03Z\x95\xee\x11\f\u0099\x18AF\xc6\xc6&ԇ\xfd;\xf8\xd4Թ\xbd\xcfL\xfa\x95TG\x838\n\xa2\a\xa0\x94.J$\xb8( \xa1\bs\xa4\x81\xd4g\xa8D\x8b\x1c^\x9eH\xe5\xc4(\x03\x99P\xb4%P2\x13Y\t\te\xfbb\xfb\x96\xd4W\x87\xed\x16:\xfd\xa0\xacC\xcaC1\x0f\x17:%\xc3\xd0\x1f\xe4\xb2\xef\x90\x14¹\xab\x0e\xf1{\aB\x8b\xad\x01\x1b\u007f\nk&\x16\x9c\xa3\xc4\x1a\xdf\x05\xcd\xf7\x91\xb0\x94\xa0\x18)-a\x9e\xd0\x1e%\xc2*\xc1\x02\t\xda<%\xf4\xfa\xc5Q\xae\xf4\x96\x90\xa79\xa1\xfd^\x8f\x14`S\xd0\xf8B\\\xcc\xc3\xdd\x01\x1e\xaa\x11\xe1\xeb\u007f\xcbS\x96\x03\x82\xa5/\xfb\x9e\xf8\xc6oA\v\x88E\xe0\x01\xc1\xa1\xd2\x13\xba\x1b\x9b\x85\x12\xe1\x1e2P\x0e\x92\\\xf6\x88\x19\x93\xdeU\x1c\x15-.W-\xb9?\x15\a\xca\xe4L\x154\x84\xd4h\xb9\xab\x84\xae\xb5<\xa9\x84j\xc4\x1fTC\xb77?\xa6\x84no:V\xd0\xff\xb2hZ\x17\xc0>\xe9j\xf6\xff\xda2\xfa\xf7\x16hX\xfac\x90\xa9)5R\xfe\x01$X\x81 [+\xb2e\xb8o\x01\xd7\x03~\xfb\xf4\xc8!߁\xf5sQ$ʆ\xa2\xbbR*\xfc>\xf4\xacAkz\xceE\xc8\xf1\x05\xbc\xef\xf8\x9bW\xc8m\xf3\n\xf9\xed쯐\x02\xc0\xfe\xa9\xfc6\xde<\xccР\xc8>\x03؇*9u+\x93\a\xa3u\xf5\x98v\x1d\x86+z\xce\x03\xde\xc6\xc6a\xbeq\nBv\xc8>\xdav\xa1\x1apm\x9f\xa9\x82\xf2?>\x93X\xe5\x13\x12\x86\xa3\x9f\xb4\xebq\xef\x8c\xf5˓\xf9۩\x92@\x9f+2\u007f\xeb\xaf(\x89x\x1cᴳy^m\xa5\xe3t\x8e\x92\xde\xf8)K+|n(\u007f_\u007f:!\xd8\x06\xe4Ԑ\xdf0\xe6Ch<\b\x9a\xbb^\xf0\xad\xd18\xbf\x1a\xf9\x104\x9e\x10\x85W>=i\x1b\x1c\v)\xa8\x19H\xca\xff\xaa?\x9d\xe0L\x03\xf2\x8a*\xba\xab\x9a\xae\x9bR\x1c\xb5\xf5\x87\xc7i\xed\xb4\x18GF.Z\x81:\b\xe1\xd8\x18\xfc\xaem-\xf9\xc7\xe1\xe3\xf0\xf2\xfdݗ\xbb\xab8B\xd9]\xef(\xe9u\x0e\xff-E\xa6pA\xf9\xb9\x8d\x95\x05\xe5=\xff\xc6\xfax\u007fu\xbc\xb64O\xfaD\xfd\x8e\xbev\xaa\xad\xc3鎣\xb0\x18^\xbb8\xb7\x8e\xe2\xa8\xfe\x99\x87_,\x97\xa0\xe5ju\xf1_\x00\x00\x00\xff\xff\x14\x9c\x92}\x17\x12\x00\x00",
		hash:  "16d315754a43e6d3ce502fd3ebe791edc82ec3ebf3df33e7c378bfee73e58f78",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1526058549, 0),
		size:  4631,
	},
	"index/transactionsummary.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xc4V\xddn\xdb:\f\xbeN\x9e\x82P\x81\x83s.\f\x9f\xeerS\f\xaci\x8b\x16\xeb0`\xe8\v(\x163\v\x95%C\xa2\xbb\x1aA\xde}\x90\x1c\x1bn\xfe\xeat\xe8\xcfMe\x92\x9f\xfc\x91\x1fI\aV+\x89Ke\x10\x189a\xbc\xc8IY\xe3\xeb\xb2\x14\xaea\xeb\xf5\x94{\x8c&Pr\xf6,\x84eS\x00\x00.\xd5#\xe4Zx?c\xce\xfe\xdeX\xb7=\xb9\xd5u\xd9c\xfa\x88\xe2<\xbb\x1f\\\t\xff\x88\xb2\xfa\x02W\x86\x9cB\xcf\xd3\xe2<\x9bN&\x13^\xeb\xee\x1e\x12\v\xcf@\n\x12I8FR\xf8$\xcaJc4\xb0\b\x98p\xad\x86\x88\x84\x14i\x04\xe5\x93\xf0\xa2Gd\x19\x17P8\\\xce\xd8Y%̵\xc8\xc9*\xe9\x19\b\xa7D\xe2QcN\x18ӭ\x91e\x9d\x9b\xfbJ\xb4ep\x98\xa3\xa1d\xd9:\x12\xb2$4\xcb\xfe\xfd\xff?\x9e\x86\x98\x8c\xa7\"\xe3\xa9VG\xc8lQؤ̲y!\x94I\xc3c\x03s[\x96\x8a<\xec\xbc\x18\x83\xfb\xd8ka\xebo\f\x85K\xe50'\xeb\x9a\vm\xf3\a\x96\xdd\tO\xd0\x1b!Z!\xd9%#\xbb\x90d\xd1\x02\xf7\x15\x81\xa7\xb5n\x0f\x83\xa6\x88Trk\b\r\rD\xedL\xfb\x95\xdd\xc6W\u00a0\x1eH\x1b\xa9\rE\xddS\r\x12\x8b\xd0\x0emC?\xdd)O{\xa2\xda\xc8\x02\x85\xdc\xefk\xfd\xee\xb0ss\xc1\xb0\xc3\xe1\xf6\x92\xa7T\x8c\xc0\x04m\xe1\xd6T5\x8d\x03\x9c\xb5\xc1\xf0UJ\x87އ\xe9\x19\a\xfbQ\xd3\t8\x9e\x1e\xca8\xe0\x0e֊\xd3\xc2ʦ\xf3\x1d\xc2\x0fc\x9e9\x82\\\x1b\xf9S\xa9\x1e\x8fuB\xaf\u007f?Q\x1f-\u007f;\xcb7\xc2\x17\xe3$\x89\x1b`t\xa3\\\xcdan=\xbd\xa9j\u007f\xaf\xd7N\xccq\xed\xb6W\xd1\aK\b\xb9\xd5a\xa5\xcdا\x03k\xf1\xd6,\xad+E\x18\xf1w\x98\x9fWe!\xb3o\xd8|\xff\xf9\x99\xa7$_\x8c\x8d\x95\xbd\xbc\x88\x88\xf8\x99\b\xcf\xf1sW&\x1e\x85ˋD+\xf3\xc0\x80\x9a\ngL\xf6\x9b?\xac\xfcc\xf7\x1f\xce\u007ft\x1a\x17V6pz.\x01\xd6\xe5\xf3\xd6\x14\xafk\xad\xe3ğ\xc40\xa0\x02\xe8\x1d\bޫ\x12=\x89\xb2:\xad\x84A\xe5\x1e\xfa\x0e4\xdb\xe1\xbaA\xf5\xab\xa0ә\xb6\xb8\xd7\xd3|y\xc3\xed:\xba\xafS\u007f\xea\x0e\x9b\xff<\xdd\xfc\x9cΦ\xab\x15\x1a\xb9^\xff\t\x00\x00\xff\xfff\x97\xc7~\x81\v\x00\x00",
		hash:  "8a5c3afc8ccb05195c89284ff19c6370e4b5e7ab7889450aa99b31f429396e16",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  2945,
	},
	"searchresults/tools.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xffD\xcbA\n\xc3 \x10\x05\xd0u<\x85\xcc\x01\xe2\x05Ի\x94\xc9\x0f\x9a\x8a)\xf3\xdd\x14\xf1\xee\x85v\xd1\xfd{s\x1e8k\x87\x97qߍ\xb2\x96\xdb\"\xd5\xeakx\x9a&\xb9\x18\x88\x87i\x01×\xec\x17%\xc7\xf03\xd9m\xb1\xd5\xfe\U00106584\xe3\xdd\xc0\x02\f\xf1\xc5p&Q\xfe\xfb\xae\xa4d7'\xfa\xb1\xd6'\x00\x00\xff\xff\x01\xea%yy\x00\x00\x00",
		hash:  "a98871ca472c8a4b10c9916ad166c30071ff142d7e9e52290254dd7f836e036a",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  121,
	},
	"searchresults/type/EC.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xa4\x93An\xc20\x10E\xd7\xce)\xdc\xec\x83Ŷ\x1a\"\x01\xe5\x06\xbd\x80\xc9\f\u0092\xb1#{\xa0\x8d,\u07fd\"\xb8*U\x03EmV\xd1\xfc\xf9\xf6\x93\xf5\u007fJH;\xe3H֛u\x9ds%Rb:\xf4V3\xc9zO\x1a)\x8ccxj\x1a\xb9\xf28Ȧi+\x01\x91:6\xdeI\x83\x8b\x9a\xde{\xeb\x03\x85\xba\xad\x84\x004'\xd9Y\x1d\xe3\xa2\x0e\xfem\x9c}\x1bv\xde\x1e\x0f.^\x04\x01\xfby\xbbD\f\x14\xa3|ѬA\xed\xe7m%'>`\xbd\xb54\xad\t\xe0\xad\xc7aZ\xbc>\"\xfc\xb6R\xf6\xf0\x13\xea\x19\x14\xe3æ\x94fŗ\xf3#FP\xb7\x88\x84\x80ے\x18\x01\x0f\xfe\xe8X.O\xda\xd8\xf3\xcb\xdc!-\x8e\x94f+m\xb5\xeb(g\xb9q\x1c\x06\xb9\x0e\x84\x86\xe3=\xeb\u007f\x18_\x87\xfe\x01\xaek\x94?\x92\x80\xba\x13\x00P%:\xe7\xfb\x14\x9aӘ\xd3\xf2\x03\xaaD\xb9-!\xdf8\xfc\n\xfa迮D\xec\x82\xe99\x9e;\xf1Cc\xef\xed\xb4\xb2\xf3\x9e/EJ\x89\x1c\xe6\xfc\x11\x00\x00\xff\xff\x9d\x93a\x14w\x03\x00\x00",
		hash:  "5decf7ae2b1cc77506b4c190c36b0a297816c2dbbf15f1cfd7c93b2ad5c606f9",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  887,
	},
	"searchresults/type/FA.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xa4S\xe1j21\x10\xfc\x9d{\x8a|\xf7\xff\f\xfe\xfdX\x0f\x94\xd6'\xe8\v\xc4슁\x98\x1c\xc9j+!\xef^<S\xb8R\xb5\xb6ͯ\xb0\xb3\xb3\f\xc3L\xceH[\xebI\xb6\xebe[J#rf\xda\x0fN3\xc9vG\x1a)\x8ec\xf8\xd7ur\x15\xf0$\xbb\xaeo\x04$2l\x83\x97\x16\x17-\xbd\r.D\x8am\xdf\b\x01h\x8f\xd28\x9dҢ\x8d\xe1u\x9c}\x1a\x9a\xe0\x0e{\x9f.\x80\x80ݼ_\"FJI>i֠v\xf3\xbe\x91W\x1e\xb0\xde8\xba\x8e\t\xe0M\xc0\xd3upz\"~\xb7R\xf7\xf0C\xd4\u007fP\x8c\x0f\x93r\x9eU^)\x8f\x10A\xddR$\x04܆\xc4(p\x1f\x0e\x9e\xe5\xf2\xa8\xad;;sGie\xe4<[i\xa7\xbd\xa1R\xe4Z\x1b\x0e\x16\xd3=\xd6my\xbf\xf3\xf4\xe54\xd0\xcf\f\xad*\xff\xe8%\xa8;\xf1\x00U\x83uvI\xa1=\x8e)\xae\x1fP5\xe8}\xad\xc0\xb3\xc7I\r\xa6eI&ځӹ-\xe3\xdd)\xc6!\xb8\xf4\xa5^\xdb\x10\xf8R\xaf\x9c\xc9c)\xef\x01\x00\x00\xff\xff\xfc\xee\x95g\x8d\x03\x00\x00",
		hash:  "e5451723f5b66ce13c8bf4d3804aa8c178a63b67b7caa5aa11b5731a35a5cb01",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  909,
	},
	"searchresults/type/ablock.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xbcU\xdfj\xdb>\x14\xbeV\x9e\xe2\xfcD/\u007f\x8e)\xbd+\x8a\xa1i\v\x19ll\x8c\xbd\x80b\x9d\xc4\"\x8ad$%\x9b\x11~\xf7!Y-\xde\x12\xa7\xe9.\xaa+s\xce\xf7\xc9\xdf\xf9\xa7\x13\x82\xc0\x8d\xd4\b\x94\xaf\x95\xa9w\xb4\xefg$\x04\x8f\xfbVq\x8f@\x1b\xe4\x02m2\xb3\xff\x8a\x02\x96FtP\x14Ռ\x00sX{i4H\xb1\xa0\xf8\xabUƢ\xa5\xd5\f\xf2aB\x1e\xa1Vܹ\x05\xb5\xe6\xe7\xc8\xf3\xb7\xb76\xea\xb0\u05ceV\xf0\a$\xc1\x9a\xdb\xeaA쥆e\xd4\xc7\xca\xe6\xb6:\x05y\xbeVxj\x1f|k#\xba\xf3\xbe\xc1o\xa7\x9d\x03@T\x9f\x8d\xd9\x1dZXq\xd7ܳҋ\xb7\x19!\xcc\aR\xe4\xf4\xfde\x12+/\x89\xb8J\xe1\x92\xd7;\xf8\x8e\x1b\xb4\xa8k|\xa7\xd2H~\xe5~\x90\xe0\x15\xcam\xe3\xaf\u05f8J\x8d8\u007fZ\x0e\xc4\x0fP\xf8\xcd\xe2Q\x9a\x83\x83Q\xff]\xa9\xf7\" \x81x\x9a\x9a\r\xaf\xbd\xd9\x17\x0e\xb9\xad\x9bBI\xbd\xa3\xe0\xbb\x16\x17/\xe38\x8a<\xcaɅz)\x11\u007fC\xc9?f\x88\x95\x133\xc3ʉAc\xcd]\xf5\xac\xbd\x95\xe8\xe0\xd1hϥF\x01R\x8fS\a\xcc\xed\xb9R\xa3\x88\xbe\xa0s|\x8b\x8f\xe6\xa0}\xdfC\xbe\x80\x95\x03\x8e\x95\xcdݙ\a!\x04\xcb\xf5\x16\xe1F\xfe\x0f7\xa8\x10\xee\x170\u007fX>I\xd7*\xde\xf5\xfd+\x81\x90\xfc.\xa4L\xe7\x14'\x03=\r\x80L=\x13\x84\x9c\xed\x15\x12\xed\"\xc5\xdc\xc1\x8f\xaeŉ\xc6ȸ\x10\xa2\xd4y\x04N5.!\xe7\vBȴ&2\x9c\xf8\x87\x98u\xd4\xd7\xcfS\xd2\xf3\xd57h?\xe9\x8dy\xb7\xa8\xeb[\x84\x90\x10P\x8bQa\x12J\xc8c5\x8b\xd7\x0f\x1f\xac\xccˤ\xca{\xe6Y\x8bѮ\x19o$W[\xd9zG\xf3\x8dc\x977F\xb9\x93\x15\xb61\xc6\x0f+,+\xf9\x1d\x00\x00\xff\xff)\a>\xed\xf5\x06\x00\x00",
		hash:  "d885b10188970b8c7bd9093ee5516b9fbbde629f218f31d3bb61962936399f9b",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  1781,
	},
	"searchresults/type/chainhead.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xecV\xcdn\xe20\x10>\x87\xa7\x98\xb5\xd0jW\xda4\xdbkj,\xb5\x80T\xa4\xbd\xf5\tL<4\x16\xc6\xceƦm\x14\xe5\xddWN\x1c\x9a\xa6@iϛ\x03?3\x93o\xbe\xf9\xb3\xa7\xae\x05n\xa4F YΥΑ\v\xd24\x93\xa8\xae\x1d\xee\n\xc5\x1d\x02\xf1B,[1\xfd\x16\xc7pgD\x05q\xcc&\x11P\x8b\x99\x93F\x83\x143\x82/\x852%\x96\x84M <T\xc8'\xc8\x14\xb7vFJ\xf3<Ќ\xb5\x99Q\xfb\x9d\xb6#\x8b\xba.\xb9~D\x98\xca_0E\x85\x90\xce\xe0\xaaiZ\x9b(\x8a\xeaz\xba\xf5\"\xfc\vS\t\xbf\x87\n\xb9\x81\xe9\xf6U\x10Q\xbb\xe3J\xb1y\x1fe\n\x94\xb7\xac7<sf\x17[\xe4e\x96\xc7J\xea-\x01W\x158#\xb8V&\xdb\x12V\xd7\xde\xf5\xd5\xdch\x87\xda]\xdd#\x17MC\x13\xcehҁ\x1e(GQ\x044\xbf\xee\xbc@p\xb9Z\xa4?\x02\xc4J\x17{\xd74?\xfb\x17\xa9-\xb8\x06\xeb*\x853\xb2Q\x86\xbb\xb4\x94\x8f\xb9#\fFN\xff\xa0~ty\xd3\xc0R\xbbR\xa2\xa5\x89\u007f\x95\xd1$\xbf\x1e\xba\xafkT\x16C\u061d\b\xa8\xe3k\x85m\xac!\xc8V0\xcatkL\xddڈ\xea\x88\"\xa2\xae<&\x8e\xa8\x13\xccS\xaa\xe0\x9e\xdb<\xa5\x89\x13'\xed>N\xb8\a:\xe4;$\xabK\xf4qX\x9a\x1c\xa7u\x96\xed\x8b\xc3Rs\x05\xab\x85=\xcfw\x12\x85\x87\xee\xd5\xe1Oו\x10\xdar\xb5\xf0\xed\xf7\xa6R\xcb\x17\xb7ZX\xf0\xd3\xf2\xfaF\xdb\xedJvc⣌1Ј\xa5h#^-|\xa8J2\x18yB-\x06`4٫ㄿ\x96\xa2\x1e\xd6\t\x16\x028\x91\x937s\xfb\x91Ak\xd4\xf6\xf6k\xc0Y\a\x1f\xdb\xfdn\xc7}\x99\x83?x\xe8\x04~ Y\x18\x99\x87\xdc<íR\x879I\xf8\x05\x0e\xd7%$\fb\xb8\xab\x1c\xdat<A\xe1\xbb\x1f\xa4\xcb\xe1z\x9em\x83\x9f@\xf5\xba\xcf`.\xe707\xd6\xc1;\xbc\xe5\xdc\xcb/\x81\xea\x0e\x00\x00\xf8r%\xfc\xac\x93\xfe\xf4\x11\xd2\x16\x8aW\xa96\x1ao\b\xbbU\xaa\x0f<\xfd\x8e;[ܼ\xff\x1c\x14\xec^\n\xfc_\xb0\xf3Py\xf9qJ\x8e\x93\xbd\xbc\x1dΏ\xed\xe7\x8e\b\x9a\x9c\xb8\rh\xd2^ \xa3{G\x8b\x11\xcd {ӟ4\x11\xf2ɟ\xa5\xfd\x0f\x9a\x84\r\x82\x85\xe5b\xa9\xc5`\xc1\x18\xae!6+e\xe1,\tn\x86*g\x8c\xb2\xef\xf6\x96\x8d1\xae\xdb[\x02\x95\u007f\x01\x00\x00\xff\xff\xbf\xbd}\xb8\xed\b\x00\x00",
		hash:  "9cae6493220b9658790e7c71938d38470ca1f300364642f144844d5511d8de55",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  2285,
	},
	"searchresults/type/dblock.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xd4W\xcbn\xdb:\x10]\xcb_1W\xc8\xf2\xdaB\x90]@\v\xb8\xb1\x13$\xb8(P\xb4\xf9\x01Z\x1cGD(\xd2 鴂\xe0\u007f/\xf8\x88\xad\xf8!?\x9a\xb8\xa9V\x02\xe7px8g8\xe44\r\xc3)\x97\b)\x9b\bU<\xa7\x8bE/i\x1a\x8b\xd5LP\x8b\x90\x96H\x19j?L\xfe\xe9\xf7\xe1F\xb1\x1a\xfa\xfd\xbc\x97\x001XX\xae$p6L\xf1\xe7L(\x8d:\xcd{\x10?\xc2\xf8\v\x14\x82\x1a3L\xb5\xfaѲ\xac[\v%\xe6\x954i\x0eo \x1eV^\xe6c\xae\xb1\xb0J\xd7p\xe38\x92\xac\xbc\xcc7\x81\x96N\x04n\x8e\a\xdbD\xb1z\xbb-\xd8\xf5nc\x00\xb0\xfc\u007f\xac\xbf|\xbb&\x99e\xfb\xb1M3\xf0\xf0Ţ\x1bO\xb2\xae\x95\x0f\xa2\xe5\x059\x92۽\x17uঞ\x85\xe3\xdd\\\b\xb8\xa7\xa6<\x9c\xa2\x9b\xe2f\x9c\x81\xdd#\xaf\xf0\xbb\xa5\xd5\xec\xe8\x00\xde)]Q\x8bl\xe9\xe1\x1cz\xbb#\x00\xf7ȟJ{4\xe1\xf1M\x98x\x06\x9e_5\xbep57\xb0vz\x0f\xe4\xdc\tp\xdfҿ\xcf}\xb8v\xf3\xa8/ESZXU\xf5\rR]\x94}\xc1\xe5s\n\xb6\x9e\xe1\xf0\xb5Ƶ\x02\xe2\xbc,O*\xcd\xc9DCv\xc4\xda\xcb\xc4\xf6\xeb\xbfu\xbb\xca\xe0\xee͞(\x04\xc9v\x145\x92\xed\xa8\x84\xa4\xbc\n\xd9c`\xa4\xa4\xa5\\\"\x03.\x83,@LE\x85hW\a7<Rsi\x17\x8b\x801$\v \x92\x95W\xdbj\xb5_\xd8k\x10\x83\xef\a\xd2M*ɮ\x8a\x9c$[\x93+q\xe3,\u007f\xacg\xb8#\u007f\"\xe2?V\xbd\xeeh\x17n{H\xbb\x17v)\xa2\xbbWޛ|t\x95|\x9e\xa5'\xe9\xaf\t\x1d\x93\xef8\xc2'$\xc0\xc1\xfa\xbcۥ١\xd8:\xf4VZ]\xc3H#\xe3\xb6K\xc1w\xacQ]\xb2\xaec\xf7\xea\x8b\xc5J`\xbf\x97\xb0\x95\x03e\xfe\x90\x03\xff\xa9\xf5\xbes\x91\xe4\xeco\x94z\xbaR:\xee\xe2\x0f\xaa\xdc4\x9a\xca'\x84\v\xfe/\\\xa0@\xb8\x1e\xc2\xe06\x14\xec\xd6ݓ$\x9f\xa1>\x87C\xfe)\xeb3.EuQ\x1c\xb4_\x05G2\x85.\xaa\xfb\xde\xc2\x11\x16Y\xec{\x06\x9fHbTR.\xe1a\xfc\x9b!+\x9c\x1b\xd7!.\xa3\x16\xdf\x0e\xde\xff\xc3\xf8\xb4\xf0uf\x98\xb2T\x80\xcb\"\x8e\xe6\xa0\bFF\xa1\"\x87\xd7̇]\xb4I\xd24(\xd9\xdb\x17\x1f\xc9\x18\u007f\xc9{\xce}\xf8!Yl\x9c\xf3\xd8S\xdfJ\xd6\xea\xab\xdbݷ)4\x9fY\x93F\x8fm\x93UJ\x98\x8dv}\xaa\x94\r\xedzd\xf2+\x00\x00\xff\xffq\xd0\xe5\xf7\xe1\x0f\x00\x00",
		hash:  "c0b6677fc67ef5062e5f8d7fedd45e6545285789d84f217496e9902761af4400",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  4065,
	},
	"searchresults/type/eblock.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xbcV\xdbN\xe30\x10}N\xbfb\xd6\xe2q\xd3\b\xf1\xd6u#\x01eE\xb5BZ-_\xe0\xc6Sb\xe1\xda\xdd\xd8\x01\xaa\xa8\xff\xbe\xf2\xa5%-\xe9U\vyj\xed\xf1̙3\xe3\xe3i\x1a\x8eS\xa1\x10\bN\xa4.\x9e\xc9r\xd9K\x9a\xc6\xe2l.\x99E %2\x8e\x95_\xa6\xdf\xd2\x14n4_@\x9a\xe6\xbd\x04\xa8\xc1\xc2\n\xad@\xf0!\xc1\xb7\xb9\xd4\x15V$\xefA\xfc(\x17/PHf̐T\xfa\xb5\xb5\xb3\xbd[hYϔ!9l\x98x\xb3\xf22\xbfS\xb6Z\xc0\x8d\xc3G\xb3\xf22\xffhd\xd9D\xe2\xc7\xf5\xb07\xd1|ѽ\x17\xf6\xabݛ\xc1\x80\xe7\xbfp\xf1\xf0g@3\xcb\x0f\xdb6Mߛ/\x97\xfb\xedi\xb6/\xf2Q\xb0~\xd6R\xc2=3\xe5\xf1\xd0\xdc\x11w\xe2\v\xd0ݖL\xa8\xf1\xe8Hl\x94\xf9>\x9a\xb2\xc2\xeaYj\x90UE\x99J\xa1\x9e\t\xd8\xc5\x1c\x87\xa4p\xee\\;\x12\x97ǽ\xef\xcb~\x8c\xe1\xb2a\xf9\xa7g\xe4[\x10\xeeQ<\x95\xf6x\xca#\xd4\xd1M8\xf8\x05\xcc\xff\xae\xf0E\xe8\xda@\xeb\xe6\x1c\x89w\xaf\x81\xfb־}\x93\xfb\xa5\x81;{\xa8|Q_Z\x848O\xeb\x9b\xc2r:\xa9 ;!\xfe\xba\xf9}\xfcM\xb7\xef]\xbe?\xe13\vA\xb3\x1d\xa2B\xb3\x1dJD\xcb+/c\x02\r\xdcje\x99P\xc8A\xa8P\x1b\xa0fƤlq\xe3\vw\xabke\x97K\x88\ai\x16\xachV^u(e\xd3TL=!\\\x88\xefp\x81\x12a0\x84~<\xda\xc1Cӈ)\xe0_o\xda\xf74\x92\a\xa1j\x8b\xf0\xc0\xaa\xe7 \xf9\xd0-\xb4\xbeұ\xc4~\x81|\xa2\xf4n\x80:\xf6\xce\xf9\xa4Z\xda\xf0Uen\x1a\x94\x06[\xcc%\xc9\t\x9c%\xbb\bK\x92N\xaa\x12\xb7\xce\xe3\xf3\xb8\xe7\x19\x88v\x87\xaf\xa8sDV\xf4\xad^\x89]\xba\x9a$\xdd\xcc\x1d\x00\xfbf\xb1RL\xc2xd\xf6\xc3\xed%\xf1\xa3\xb5\\\xff\t$C\xec\xf4\xf1\xc85\xb9G{\xf7f\xc7#\x03nNy\xb7\xf4\xb5\x92\"\f(.\xb9\x14c\xf8T\xf8g\xe4\"\xb4\x87\x149lE@\xc5[\xceh\xe60t\x01=\x8f\x9a\x95[\xf7JjeQ\xd9\xff%\xcf\xd4̙j%\\\x04\xf7\xa9\xa9g3\xe6\xaa\x1b\xe3\xc1cX\x18\x00ey\x94\x9f\xc7R\xbfµ\x94\xefB\xb3R妱U\xad\n7\x11\x86\xab\x15\x9c8\xee\\\xbc\xf3a\xb9~'`\xecB\xe2\x90pa\xe6\x92-\x06J+\xfcA\xf2k)a\xc5\xce6ʈ\xbe\v\xe9\xe9\x00O,\xe2Y\xb2\xa0\xf8\x86*l\xaf\xf8\xd3\\\xbc\xb8F_\xfd\xa0Y\x1c\xb0\xf38{\xdf)ޚ\xbf\xdbS\xba)*1\xb7f\xa5\xd9\xed-\xab\xb54\x1f\xc6\xfa\xa9\xd66h|D\xf2/\x00\x00\xff\xff+Q\x86\xa5\t\f\x00\x00",
		hash:  "909c9767e9d7b7300eb04417ff26ad99ea315b9e7bad9f5c9f66c0ea34e62167",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  3081,
	},
	"searchresults/type/ecblock.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xbcUMo\x1a1\x10=ï\x98Z9vcE\xb9Ef\x0fP\xd4\x1cz\xe8_0\xeb\x01[16\xb2\x1dZ\xb4\xda\xff^\xf9\x83j[\u0605\x10\x89=y\xfdތ\x9e\x9fg<m+p\xad\f\x02\xc1f\xa5m\xf3F\xban:iۀ\u06dd\xe6\x01\x81H\xe4\x02]\xdaf_\xaa\n\xe6V\x1c\xa0\xaa\xea\xe9\x04\x98\xc7&(k@\x89\x19\xc1\xdf;m\x1d:RO\xa1|L\xa8=4\x9a{?#\xce\xfe\xea!\xff\xa3\x8d\xd5\xef[\xe3I\r\xffP\x12M>\xd5K\x13\xdc\x01\x16\x0e\x85\n0\x8f2\x19\x95O\xf5)7\xf0\x95\xc6\xd3\xfd\x8c\xad\xac8\x9c\xc72\xee\x86\xc1L\x10\xf5+\xf7\xf2\x85\xd1 .S\xdb\xf6q\xb9HZ\x1f\xbfc\x88\x81]7\x1e\xc9蘄\xab\xf4\xa5˹]d\xba긊y\xee$\xf9\x15\xd5F\x86O\xea\xfd6\xcfi\xee\xa0\xf7\xa7ý\xb2\xef\x1eN\x8b\xf2\xcaC\x8c\x12\x12\x89\xa7\x8eZ\xf3&\xd8m呻FVZ\x997\x02\xe1\xb0\xc3\xd9\xdf^\x1d\xf4#\x8a\xcc\u007f\xc7[\xe4\x17t\xddh\x1b\xa3\x03m\xc5\xe8P/2\xf9\x9c\x1aZ\xa1\x87\x855\x81+\x83\x02\x949c(0\xbf\xe5Z\xc7c\xfe@\xb3\t\xb2렄2\x9a!F\xe5\xf3\x99W\xa3<\x05\xc9\xc7b`\xda \xb7=\x0em\xeb\xb8\xd9 <\xa8\xaf\xf0\x80\x1a\xe1e\x06}狨\xae\x1bɠ\xd6)\xf4H/\xfd\xf5\xb9b̖}\xa0\xe1/\x97V\xcc\x18\v\xeb\x8c\xd6XEW\x14\n\xc0\x88\vhĨK\xc3\xf8\x87J\x8dQ\xa1\xf6\xf5t29.\x18-ê.sliDo\x96\xf5'\x9eo\x9c\xda\x05O\x8a\x8e>\x14\xac\xd5\xfedD\xae\xad\ryD\x16\xfd\u007f\x02\x00\x00\xff\xff\xfeґ\xadV\a\x00\x00",
		hash:  "5a92072cac8a8cafae3c30cd758ec30320b45160754be5716e3753b8da4fc2d6",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  1878,
	},
	"searchresults/type/ectransaction.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xc4TAN\xc30\x10<'\xafX|\x0fV\xaf\xc8\xcd\x01T\xc1\x03\xf8\x80ko\x15\v\u05ce\xec\xa5\x10Y\xf9;\xaa\x9d\x8a \xaa6\x87\n|\x8a\xd6\xe3\xdd\xc9hfSҸ3\x0e\x81\xa1\xa2 ]\x94\x8a\x8cwl\x1c\xeb*%\xc2}o%!\xb0\x0e\xa5Ɛ\xcb\xe2\xaei\xe0\xd1\xeb\x01\x9a\xa6\xad+\x111?\x01\xa3\xd7\f?{\xeb\x03\x06\xd6\xd6U%\xb49\x80\xb22\xc65\v\xfe#\xd7~\x14\x95\xb7\xef{\x17\xcb\x05\x80\xe8V\xed\xc6Q\x18\xe0)\xa06\x04\xaf߄\x04\xefVm\rg\x8e \xb9\xb5\x98\xa7G\x94AuM.\xb0\xf3\xe8ӛ\xad\xd7\xc3%D\x06\x85+\x88\x82\xd2\vP\xc73\xfb\x1bx\x91\xb1{XМ/꾜CJ\xf7\xcfH\xc7\xe9\xe3x\x93\xf1\x82_U\xe9\xd6:\x16\x8b\xfc\x97\x82Bf\xab\xed\xa4\"\xbfo&\xc7Y\xe3\xde\x18\xd0\xd0\xe3\x9a\xe1\x91\x1ek\x8bҙk\x91[p\xd9\xfe\x85\xe4\x82_2\xb7\xe09\x1d%\x8a\\\x9bC\x0e\xea\xf4!\xf8\x94\xe5vJ\xf9\xc6\xe9Y\xd2\xe7\xfb \xaa`z\x8a\xecd\xa3\xf9\x1dyo\xe3\xaf\r\xb2\xf3\x9e\xca\x06I\t\x9d\x1eǯ\x00\x00\x00\xff\xff\xaeĢ\x15{\x04\x00\x00",
		hash:  "7d590b3a71bd677ee47f433a0177e1d1de7eb5fee91da29a4e0a5e88fc34dbcd",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  1147,
	},
	"searchresults/type/entry.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xe4U\xc1n\xa3<\x10>ӧ\x98ߪ\xfe\x1bE\xbdR\a\xa9\rH\x8d\xb4\xb7>\x81\x8b'\xc5Zc#{\xd2\x06!\xde}eC\x14Z\x0eA{]\x0e\xd10\x13\xbe\xef\x9b\x19{f\x18$\x1e\x95A`h\xc8\xf5l\x1c\xef\x92a l;-\b\x815($\xba\xe8\xe6\xff\xa5)\xbcX\xd9C\x9a\x16w\t\xf7X\x93\xb2\x06\x94\xdc1<w\xda:t\xac\xb8K\x12.\xd5'\xd4Zx\xbfc\xce~E\xdf7gm\xf5\xa95~\n$\xbcy,\xaa@\x0e\xa5 \xc1\xb3\xe6q\xf6\x93x\xd78\xd9\t\xa7\xf7\xc0\x1c\xb8<\nW7i\x8c\xce\x10!\xee.f\xc2I\u0380\xaf\xc279\xcfH~\x8b\r\xc3C\b\x8c\xe32³+\xc2O\xb0}#\x94\x81C\xb9B\xe2\"*:\x8a\x9al\x9b\xce´2\xbf\x19P\xdf\xe1\x8e\xd5\xe1\xcbPC\x16X#Ρ\fĢ\xd8J^\x9d\t\x9d\x11\x1a\x0e\xa5_\xe7r\xb1\x13~\xd2\xd7\x17\x00\x80a\x00'\xcc\a\xc2\xfd\xa1\x84|\a\x0fՙ\x0e\xa5\x87\xd0\xca\xc5\xdf\xc2õ\x9a\xba\x18\x8a\x96\xe2L\x98\xaa\xa8\xfa~\x12\xacU\x01?\xf1\xd1\xc8%\x1c\xcf\x16\x1a6\xd7\xd6\x1aBC9\u007fw\x90\x15\xeb\xfc\xe0\xc6\xc3}'\xccB}=\xe1\xa5\xfeԶ\xc2\xf5\xecB\x00o\x93#\a.\n\xee[\xa1u\xf1\xd6\xd8/x֚g\xd3{h\xcbm\xc2(\x14Rx\xe9\t}\x0e\xa1\xaf\x13\xc5/4\x1fԌ\xe3v\x88\x8b\xb6xN\x17H\xd3\xf1\u070eS\xedao=AĨ\xf6\xc1\xde\xf2y\x16jW\x00\xc0_W9\\J\x06\x9ez\x8d;&\x95\xef\xb4\xe8sc\r>\xb1\xe2Y\xebK\x82\xf9\xff\xd8\xfa\xeei\xfd\xbbhƫ\x92\xf8/7\xa3q\xb7ӽ\x8a\xda\xde\xde\x1b\x17\x92gq\xb2\xce#7\xbb\xce\\\x9eI\xf5\x19\xc7\xf9l\xf0l\x9e\xf8ż\v*#\x17\xfb`\xb95|\xedTG~\xb5M\xc8Z\xbd\xf6\x1e\xad\xa5i\xc7\f\x03\x1a9\x8e\u007f\x02\x00\x00\xff\xff\x1a\x98\xf9\x91\x95\x06\x00\x00",
		hash:  "a9b9fe61fc7a2385e955d54519ae2bb7a9da16621e812e7030e44ee0fc63100b",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  1685,
	},
	"searchresults/type/entryack.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xbcTAO\xf30\f=\xa7\xbf\"_\xee]\xb5\xeb'/\a`\x12w\xf8\x03^\x93\xa9\xd1ڤJ\xbcA\x15忣5\x05\x15\xb1\x8e\x1d\x80\x9c*?\xd7\xcf~\xf2s\x8cJ\xef\x8d\xd5\\hK~\xc0\xfa R*X\x8c\xa4\xbb\xbeE\xd2\\4\x1a\x95\xf6c\x18\xfe\x95%\xbfsj\xe0e)\v\x06A\xd7d\x9c\xe5Fm\x84~\xed[\xe7\xb5\x17\xb2`\f\x949\xf1\xba\xc5\x106»\x971\xf6)X\xbb\xf6\xd8ِ\x01\x06\xcdZn\xcf\xfc\xfc٣\r\x98\xab>\x11\xd21@լe\xc1/< ܵ\xfa2ƀvN\r\v \x03\xf2K\x10\x03R\xb9\x99G\f\xcd\u007f\xa8H]M\x05\x1c\xc7\xdfcM\xae+\x83F_7ek\xecAp\x1az\xbd\xc9\xc2\n\x19\xe3\xea\xa3jJP\xa1\xbcV\x1a\xaa\xa5\x0e\xe7\xf3\u007f\x9b2\xe5)y\xef\xba\xce\xd0$镡.\xfczK\xde\xf8b\\e\x9a\a$\\e\xaa\x94n\xa2\xb9\xa1\x9f\x9fV$\xef\xdbo\v2\xb2\xfc\xb1\x1e\xe7\xe5Y^~\xa8&ۜ\xf7\xb7R\xe64\xfau\xfa\x80j\xb2\xb4\x9c̾\xb5jf\xf8\xf9Y\b\xb57=\x05\xf1>\xd1\x1c#\xe7\xda\xf0\xe5\x90읣|Hb\xd4V\xa5\xf4\x16\x00\x00\xff\xff\xe2\x95\b\x0e}\x04\x00\x00",
		hash:  "bdd438304a91ea0bbb47d4b261e0ec89fa0888139f2b8d47ff1353b21492e7d0",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  1149,
	},
	"searchresults/type/factoidack.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\x9c\x92Mn\xc20\x10\x85\xd7\xce)\\\xef\x83Ŷ\x1a\xbc\xa8h%\xd6\xe5\x02&\x1e\x14\vcG\xf6@AV\xee^\xe5\xa7j\x10)\xad\x9aU4Oy\xf9<\xfer6\xb8\xb7\x1e\xb9\xd8늂5\xba:\x88\xb6-X΄\xc7\xc6iB.j\xd4\x06c?\x86\xa7\xb2\xe4/\xc1\\yY\xaa\x82A\u008al\xf0ܚ\x95\xc0K\xe3B\xc4(T\xc1\x18\x18{\xe6\x95\xd3)\xadD\f\x1f\xfd\xecfX\x05w:\xfa4\x04\f\xea\xa5z\x1b\b\xf86j\x9f\xf4\xd0\xfbN\x9aN\td\xbdT\x05\x9fy\x80\xf4\xce\xe1|ƀv\xc1\\\u007f\b\x19P\xbc\x8fz\x162jʰY?\x83$3\xdfs\vc\x14\xe8~\x17\xfd6\x8feB\x1d\xab\xbat\xd6\x1f\x04\xa7k\x83CB\xdf\xedB\xe5\xbc\xd8^6\xeb\xb6\x05\xa9\xd5\xef?\x029\xc7}\x8b\xf1\xe0`_\v\xfd\xebyr^\f\x9ft|\xffgc \x1f\\\x06\xc8\xf1\x1a;Hi\xec\xb97h|\x019J\xa6F\xfd^\xbd\x99(8\x155U\xd16\x94:S\xbb\xdaiD!\xb8tg\xf6>\x04\x1a\xcc\xce\x19\xbdi\xdb\xcf\x00\x00\x00\xff\xff}\xcag\xb0\x10\x03\x00\x00",
		hash:  "d3778c57993f99c57a010e02aaf2088588c9f9e59abce76a0219877b96a1faf4",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  784,
	},
	"searchresults/type/facttransaction.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xdcW\xddN\xdc<\x10\xbd\xde}\n\u007f\xfe\xa8\xd4J\r\x01Do\x16\xc7\xd26\r-W\xbd\xa0/\xe0\xb5\xcd\xc6\xc2kG\xf6\x04XEy\xf7*N\x02\xa9\xf8\xcb\xf6\x87]ؕ\xa2h2\x1e\x8f\xcf\x19\x8d\xe7T\x95\x90\x17\xcaH\x84/\x18\ap\xccx\xc6AY\x83\xebz:\xa9*\x90\xabB3\x90\b\xe7\x92\t邙\xfc\x17E\xe8\xb3\x15k\x14Et:!^\x86%H\x89\x04˛B['\x1d\xa6\xd3Ʉ\bu\x85\xb8f\xde'\xd8\xd9\xeb`\xfb\xc5ȭ.WƷ\x1f\x10\"\xf9!=e\x1c\xac\x12\xe8\xc7].$\xce\x0f\xe9\x14=\xf0#\xc0\x16Z\x86\x8d\xbdd\x8e\xe7Q0\xe0\x87\xbd\xfb5\v+\xd6Oy\xb4^\x0eyXk\x99\xe0\x05\xe3\x97KgK#\"n\xb5u\xb3\xff\x8f\x0e\x9a?\xa6\x04\x04%q\xf3\xb8}\x89\xc1\x8d\b\xfd\x9cK\xf0\x12\x88[\xed\vf\x12|\x84\xe9\xe3\x9e\xcfǺ\x83\x8a\x8es\xee\x13\xe8\x10\xb8V\x02\xf2٧\x83w' o bZ-͌K\x03ҝ\xe0MB\xe6\xc7\xf4\xcc\x14%x\x12\xe7\xc7\x1b,\xac*\xc7\xccR\xa2=\xf5\x11\xed)\x83f\t\xda\xff*\xa1\x8dU\xd7\xe3\x03\x85,X(\x98\xa6\xe0\xed*\xea\xeaF+s\x89\x11\xac\v\x99\xe0\xd39\xa6U5\x17\xc2I\xefO\xe7\xa9uNrh6n6\xed\xec\xfb\xe7\xe0\x94Y\xd65\x89\x19Ed\xe1P\xbcс\xa4\x11\x9b$\x1e\xaa\xeb\xb7ȻG٦\x8c}/\xe1O)\xb3%\xf4\x9cu\xd1^\x8c4[\xc2\xf6X{\f\x84,\xfdW0d\xe9\x00\x86,\x1d\x03\xc3.\xd5.\x89Ƕ\xa9\x11AG4\xe3ѽxd\xfe\x83[\v\x9d}A\xef\xe1F\x89\x0f\xb3\xe9߂h|\"U\xd5\xd0}\xae\x96ߘ\xcfǐ\xb5\xfbx6'\xd9\x1e\x94\xaf\x1cG\vL\xa3p_n\a\xc2\x01\x91\xf3\x95-\r\xf4\x9di?\xa4\xd6\xdf\xe4\xa8\x1b\xff\xfc\xeb\x06\xba\x1fb\xdb.\xbf\x93\x88\xdf^@o\x04\xf2,\xdde\xb4\xb3\xf4M =\xe8\xc6\xe7\xc0\xa0\xf4[\xea\xc7\xed\xe6/ӎI\xfc\x94`\xbc\x1dX\x1ae\x1b\vu\x15to\xf7B\xe2N\x1a\xd3N4gF\f\x84\xf3P^{\xeeT\x01\x1e\xf7g\x1a~\x03k\xb5\xbf'\xc8/\xac\x85V\x90w\x03\xd9\xcf\x00\x00\x00\xff\xffy\xac\x81\xde\xcc\x0f\x00\x00",
		hash:  "87249c7c9dd50e7b313ec61605b35b41978f0a5946f208ad8db955ddb3bd0b82",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  4044,
	},
	"searchresults/type/fblock.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xff\xd4X\xcbn\xeb6\x10]\xdb_\xc1\xb2鮲\xf2\xeaơ\x058\xae\x92\x14m\xd0\"\xed\x0f\xd0\xe2\xd8\"B\x93\x069Jb\b\xfa\xf7\x82\x92|#_\xbf\x94\x97obo\x04ΐ<s\xce\fGb\x9e\v\x98H\r\x84N\xc6\xca$\xf7\xb4(\xba\x9d<G\x98\xcd\x15G 4\x05.\xc0\x96\xc3\xec\xa7  \x97F,H\x10D\xdd\x0ea\x0e\x12\x94F\x13)\x06\x14\x9e\xe6\xcaX\xb04\xea\x92\xfaǄ| \x89\xe2\xce\r\xa85\x8f\r\xcb\xf7\xd6Ĩl\xa6\x1d\x8dȊK閞DW<A#\x05\xb9\xf4\bY\x98\x9eD\xebn\xc8\xc7\n\xd6\xc7+\xdb؈\xc5f[e\xb7ۍ\x95\x83\x88ʸo\xef\xfa,D\xb1\xdf;\xcf{~\xc2\xed]Q\xec\x9e\xc0\xc2]\x9b\xb7Bv\x03r\x9ab{`\xbf_V3\x0e\x00-~JR\xae\xa7@\xee8B{\x84~\x9a\x9fq\x00\x84\xffXx\x90&sd%\xc7ZB\xdd\xe9P:\xf1\xb26&~\xedY\xe0\x80\xdb$\r\x94\xd4\xf7\x94\xe0b\x0e\x83e\xd1\xf9\xa0=\x92?\xa1N\x19\xbeg\xf3W\xd2\xc2\xc2-\xa5\xc0\xc2-\xf5\xc3ҳ\xe8?˵\xe3e\xa9;22\x1a\xb9\xd4 \x88ԫ\xa4\x11\xe6f\\)\x1f\xcb_\xa0\xa7\x98\x16\x05\x895Z\t\x8e\x85\x95\x89\x85\xe9ن\x1a\xcfs[fɑ\xfc\x95\x1c\x81\x02\xd2\x1f\x90^sע\xe8\x92\xcd\x15_\xf2[\x13[\x0e\xd0-\xbc\xec;\x04|\xba\x10\x87\v\x05\x03:\xe6\xc9\xfdԚL\x8b 1\xca\xd8\xfeϧ\xc7\xfeO#\xafzI\xfe\xf3\xc3\xce\x1cl\x95\x84$1\xca\u0379\x1e\xd0S\x1am\xf7ܟm;\xce\xc0M\xdb\xd6\xd1>J\x81i\xff\xb7\xe3_.\x10\x9e0\xe0JNu?\x01\x8d`/h\xcb\xd5\xd2\xf3\xe8\x0f=\xcfб0=o7gEt\xa9\xbd\xe6^\xfb\xde5`\xb5\xd4&\xd1_[gWC_cC!,8w5\x1c\x19k!A\xbf\xaf߯\x1e\xef\xfd\x8bV\xeaiU\u007f\x84\x8d-\tۆ\x02Z\xb4\x84\xbb\xffd\xd9 њ0/\xd0\xe5\xef\f\xdf \x8cɰ\xa9L\xbd\xd8!\xa41\x19\x1eZ\x9b]\x91ǣ\x0f\x88=\x1e5b\x8fGmb\xff\xa1i\xb9\xb5M\xbcGsjݳ\xf7\xc2l\xb4\x0er\xc3]\xda\xef\xbe-\xf4w\xea\xf5<A|FV6\xfdk@\x0f\xf0\x03[\xfe;\xb2j\x90+R\x9e\xcd\a!4\xcf\x1b:\x0eg&Ӹ,\x90^\x89e\xd9&\x96\xaf!\xeeK\x10\xb8|g\xaa\x8e\x93\xcf\xc1䷣\xedkQ\x19\x8f>\x15\x8b\xf1\xe8\xc3\x19|\xf9\v\xfc\x86&\xc0B!\x1f\xa2n\xa7\xb3|`a\xfd%\x1f\xd5\x1f\xf9\xb1\x16\x8d\x0f\xfd\xe6u\x80K\xac\x9c\xa3\xa3\xf5\x8aM\x13\x1a\xa3\xdc\xda\xfd\xc1\xc4\x18\xac\xee\x0fj$\xff\a\x00\x00\xff\xff\x8a\x8c\\/r\x10\x00\x00",
		hash:  "0e4acdd241874ffef73fb1ea1efd38e23daf97e38822a335bcf44216e4c1cffd",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  4210,
	},
	"searchresults/type/notfound.html": {
		data:  "\x1f\x8b\b\x00\x00\x00\x00\x00\x02\xffdR\xc1\x8e\x9c:\x10<3_Q\x8f\xf3̠\xbd>y9DJ\xa4\xbd\xe4\x92\xfc@\x83\x9bqk\xc1Fv\xc3,B\xfc{\x84\x87l&\xcaͪ\uebaer\u05faZ\xee\xc43J\x1f\xb4\v\x93\xb7嶝\x8auU\x1eƞ\x94Q:&\xcb1\xc3\xe6\xbf\xcb\x05_\x82]p\xb9ԧ\xc2$nU\x82\x87\xd8ג?\xc6>D\x8ee}*\nceF\xdbSJ\xafe\f\xf7\x8c\xfd\x05\xb6\xa1\x9f\x06\x9f\x1e\x85¸\x97\xfa{P|\xdb\x05\x98ʽ\x1c\xf0X\xbf)\x12\U000d080e#\xc3QB\xc3\xecA\x1e\x1cc\x88\xb8;\xe9\x19\x1a\x17\xf17h@\x1b\x86\xb1ge\xa8\xa3}\x96b\xeb\xae\xf8\xe9\x18\xa2<\xa0\xe1\xbd\xef\x01\xb3ŝ\x12|Pd\xe7gH\x87%L\xf04ˍ\x94\xedN\xa8N\x12F\xba1\x9a\x05m/\xed\xfb\xce@\xe8ſ\x9f\xf7\xc5\x10?N\xaay\xbf\xe3\x83\x1b\xcaq\x80\xf8L\xf0\t6\x14AM\x98\xf9\x8a\xb7\ue870#\xe9\xd3\xf9\xb9i\xa0\x05\x8ef\xce%\xb6\xb0\x13\xef::j5\f\xf6p`\xc3\xdd#\xc4<\x97\x8d\xed.\xf8C\x92d!\xe2sŒRC\x89\xaf\xa6\x1a?\xbf\xf4\xc7\x1f\x81\xff\xc34\xf5\xba^\xb7\xcdTM\xfd\xbb\xc9TV\xe6|\xc5\xe3a\xaa\xe3\xd0\xf5\x11\x81\xaf\xde>\xc5\xe09,\xa9\x8d2j\xfa'D]\b\xfa\bѺ\xb2\xb7\xdb\xf6+\x00\x00\xff\xfft8\xff\xe9y\x02\x00\x00",
		hash:  "f6de428c64d09d9ace9dd83a38ecbddf09b543f5747e06f555262ebf01ea4d67",
		mime:  "text/html; charset=utf-8",
		mtime: time.Unix(1516307821, 0),
		size:  633,
	},
}

// NotFound is called when no asset is found.
// It defaults to http.NotFound but can be overwritten
var NotFound = http.NotFound

// ServeHTTP serves a request, attempting to reply with an embedded file.
func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	f, ok := staticFiles[strings.TrimPrefix(req.URL.Path, "/")]
	if !ok {
		NotFound(rw, req)
		return
	}
	header := rw.Header()
	if f.hash != "" {
		if hash := req.Header.Get("If-None-Match"); hash == f.hash {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("ETag", f.hash)
	}
	if !f.mtime.IsZero() {
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && f.mtime.Before(t.Add(1*time.Second)) {
			rw.WriteHeader(http.StatusNotModified)
			return
		}
		header.Set("Last-Modified", f.mtime.UTC().Format(http.TimeFormat))
	}
	header.Set("Content-Type", f.mime)

	// Check if the asset is compressed in the binary
	if f.size == 0 {
		header.Set("Content-Length", strconv.Itoa(len(f.data)))
		io.WriteString(rw, f.data)
	} else {
		if header.Get("Content-Encoding") == "" && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			header.Set("Content-Encoding", "gzip")
			header.Set("Content-Length", strconv.Itoa(len(f.data)))
			io.WriteString(rw, f.data)
		} else {
			header.Set("Content-Length", strconv.Itoa(f.size))
			reader, _ := gzip.NewReader(strings.NewReader(f.data))
			io.Copy(rw, reader)
			reader.Close()
		}
	}
}

// Server is simply ServeHTTP but wrapped in http.HandlerFunc so it can be passed into net/http functions directly.
var Server http.Handler = http.HandlerFunc(ServeHTTP)

// Open allows you to read an embedded file directly. It will return a decompressing Reader if the file is embedded in compressed format.
// You should close the Reader after you're done with it.
func Open(name string) (io.ReadCloser, error) {
	f, ok := staticFiles[name]
	if !ok {
		return nil, fmt.Errorf("Asset %s not found", name)
	}

	if f.size == 0 {
		return ioutil.NopCloser(strings.NewReader(f.data)), nil
	} else {
		return gzip.NewReader(strings.NewReader(f.data))
	}
}

// ModTime returns the modification time of the original file.
// Useful for caching purposes
// Returns zero time if the file is not in the bundle
func ModTime(file string) (t time.Time) {
	if f, ok := staticFiles[file]; ok {
		t = f.mtime
	}
	return
}

// Hash returns the hex-encoded SHA256 hash of the original file
// Used for the Etag, and useful for caching
// Returns an empty string if the file is not in the bundle
func Hash(file string) (s string) {
	if f, ok := staticFiles[file]; ok {
		s = f.hash
	}
	return
}

// Slower than Open as it must cycle through every element in map. Open all files that match glob.
func OpenGlob(name string) ([]io.ReadCloser, error) {
	readers := make([]io.ReadCloser, 0)
	for file := range staticFiles {
		matches, err := path.Match(name, file)
		if err != nil {
			continue
		}
		if matches {
			reader, err := Open(file)
			if err == nil && reader != nil {
				readers = append(readers, reader)
			}
		}
	}
	if len(readers) == 0 {
		return nil, fmt.Errorf("No assets found that match.")
	}
	return readers, nil
}
