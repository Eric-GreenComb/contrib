## crypto

### DES,3DES,Aes使用方法：

- DES,3DES,Aes

```

package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Eric-GreenComb/contrib/crypto"
)

func main() {
	// DES 加解密
	testDes()
	// 3DES加解密
	test3Des()

	testAes()
}

func testDes() {
	key := []byte("sfe023f_")
	result, err := crypto.DesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := crypto.DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func test3Des() {
	key := []byte("sfe023f_sefiel#fi32lf3e!")
	result, err := crypto.TripleDesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := crypto.TripleDesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9fd&fwflsfe023f_9fd&fwfl")
	// key := []byte("sfe023f_9fd&fwfl")
	result, err := crypto.AesEncrypt([]byte("polaris@studygolang"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := crypto.AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

```

### RSA使用方法：
- 生成RSA公私钥文件

```

package main

import (
	"github.com/banerwai/gommon/crypto"
	"log"
)

func main() {
	var bits int
	bits = 2048
	if err := crypto.RsaGenKey(bits); err != nil {
		log.Fatal("密钥文件生成失败！")
	}
	log.Println("密钥文件生成成功！")
}

```

- RsaEncrypt 公钥加密
- RsaDecrypt 私钥解密

```

package main

import (
	"fmt"
	"github.com/banerwai/gommon/crypto"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	initData := "abcdefghij_klmnopq_"
	init := []byte(initData)

	data, err := crypto.RsaEncrypt(init, publicKey)
	if err != nil {
		panic(err)
	}
	pre := time.Now()
	origData, err := crypto.RsaDecrypt(data, privateKey)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	fmt.Println(now.Sub(pre))
	fmt.Println(string(origData))

}

var decrypted string
var privateKey, publicKey []byte

func init() {
	var err error
	// flag.StringVar(&decrypted, "d", "", "加密过的数据")
	// flag.Parse()
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
}


```

- RsaSign,RsaVerify

```

package main

import (
	"crypto"
	"crypto/md5"
	"fmt"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"io/ioutil"
	"os"
)

func main() {

	initData := "abcdefghijklmnopq"
	init := []byte(initData)
	hashed := md5.Sum(init)
	_sign, err := banerwaicrypto.RsaSign(crypto.MD5, hashed[:], privateKey)

	init1 := []byte("abcdefghijklmnopq")
	hashed1 := md5.Sum(init1)

	err = banerwaicrypto.RsaVerify(crypto.MD5, hashed1[:], publicKey, _sign)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sign is right")
	}

}

var decrypted string
var privateKey, publicKey []byte

func init() {
	var err error
	// flag.StringVar(&decrypted, "d", "", "加密过的数据")
	// flag.Parse()
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
}


```