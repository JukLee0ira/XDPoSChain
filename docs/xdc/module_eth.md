## module eth

### eth_accounts

The `accounts` method returns a list of addresses owned by the client.

Parameters:

None

Returns

result: array of address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_accounts"
}' | jq
```

Response:

```shell
{
  "jsonrpc": "2.0",
  "id": 1001,
  "result": [
    "0x724da43371727bf1d7eed3741217ee03346de9ea",
    "0x1b1552127184fa0963ad910538c293d83ddfdd32",
    "0xb2f5f4b60e477486497e7ddab913fd24c1e81b4b",
    "0x9c3efc01f7ff6681b0bcb18ce5cd9f9ab0b2c2fa",
    "0x24da29f981944fce2a95da5d82eca943ad68f31d",
    "0x90bb08c3d474ba1438647ef936a343133184de45",
    "0x27dcfcacd95faefd3746d617884135277359fdce",
    "0x9f4b325593c3d8ab453d4ab96754b501f5ea1052",
    "0x0fb10bad35543c6e472af5d12499f679151e9457",
    "0xc05bb5b4f73df4626087152016b832b4d96bb14e",
    "0x13b90b7f16821dedbf2658be6024b1df3afa51e3",
    "0x32323ccabb3898760b21d7c4fbec34f4226d283d",
    "0x26f7afdf81afea4e587f6edf876b1c2fbb8ac802",
    "0x39fa212f8e19083f6dde1187e0357b2cd8138808",
    "0xb533167c0ee983ea166bc11795dfd79ed31c7d0a",
    "0x87db18fd32cd84166df77910e532296c59f03868",
    "0xab6492cfe5dd604420e72f17655bb086b23c4a6a",
    "0x2640e04fb446b8f84313cd9fa0c23f2d76bbea48",
    "0xde692ca132e2898eb91a759a349ef21295f96349",
    "0xc85ece9bac746166e8bada7f307e1f7c5d1035a4",
    "0x088f5051fdaebb895b84dd88b706bb3c217160de",
    "0x0ce287b0de17e78a870f0a7498604b72b4fc6856",
    "0x71634b300a38c002d948468a8bc3217adb74170e",
    "0xdaca1b4995e07c89e33ee49ab16c6904675021dd",
    "0xfac5307dd07b72bcbac131ded24c54b16b93335d",
    "0xbeecce90688140b2bf15cef7be741e0d676fa5b8",
    "0x86ca29948e01c263ff95104831a714f78336c0b0",
    "0x3cd4100d51b34ce94ad39af4667698e34e0bab4e",
    "0x83fdcda7d2fb846188a1342e6801782b8ae87348",
    "0x94578a99b383e1564f85c2dc593d53835058174e",
    "0xad55a0a9b1feae1613b53cd3e01c085d199b762a",
    "0xfafda28d9e8b422d63f0a25911d0102df0a69ef1",
    "0x8998ad1090322052f68788024f3513ba9ce9dbe5",
    "0xe454b141657a3d18dbcd5e1a93790e20c1b72368",
    "0x6c32ba44e927b2d21828bc27e02dad5f57fe8187",
    "0x85f850c54956ef081e37cd0e0703cf7cbd066962",
    "0x7b11167e2314561b3a1b7c9e7cb8a669c7104425",
    "0x96fdefc8f75c447b649003f4174ce210db607b9c",
    "0xc5a55126cd823d02007e245f5331348f56051354",
    "0x5bdcf254229ef75c70027b05081b60f36509bae7",
    "0x3fabaa5d858410c9fb471547755982cf2b10e43c",
    "0xf604cfd8b8a1594e37a7cc15b1cb7b746cc8eac2",
    "0x40f27b0a774d2053dc87d910f78441121d53f6bb",
    "0xe3aa031aab2a9490b9813096833c0f1813ed0258",
    "0x70eb3026c182858eaedcd7222cbbccd2782d5059",
    "0x946643bc588ec9cd4843d5345d28348367e6832c",
    "0x0e881629024fe5eee54bf5e808919458f71964ff",
    "0x2acc9d95ae9b1659c17839665c3e324416ad91e9",
    "0xdb95ecd043af59a2fb59cc5e442cade45895c16d",
    "0xc6341654d7049dc864591cefa5da1b24d2487d41",
    "0x81fededdf3ee449fc2b0f4376d3ab9e82c809f92",
    "0x80f5669925dcdbd2c36bff4812db69e4b95be143",
    "0x5c586fbe12fa4628c190c419d8e7cc931b19a65c",
    "0x3d0ee4c88b6b123eea614dcfa77e4cf70f2166a7",
    "0x83b09e4923508ba5e264d0775c1739d98fd570c9",
    "0x0d5a1390c46e0b1be30d6f5c71810d810ca73b29",
    "0x3008bdfc5c093b1522270324101fcd383477e9af",
    "0x814785a8cbe49921832107488262edb663147e90",
    "0x203237fd77955515abcb5c9a4c685c73938b979c",
    "0xdf509beec43cbb9b3bd9936fedcba5388e937a29",
    "0x76f144034fc41798b1eb87579096e9b157bd9534",
    "0x122ced1daca5f82131009e522dd56eff93a70b36",
    "0x23c63053483a34ad7a317f6f9af843dfada899dc",
    "0xc8109fa721a4a8020e3e3617586f5b0d396d7641",
    "0x313511481bb912da9c0d0a96fdac41fdfc8ce24e",
    "0x310266c6bcb2eb6f0401c533f0e766913263bb0f",
    "0xb52f5985d65c72ee37f1e551d0c965b93f687172",
    "0xacb3ed55f98164325503134363e66219b54363c5",
    "0x4f1b0b92a053e242a8a6e86f60b3c62359821662",
    "0xfa19585a0b4e29cade79560b45c0bf2043e42453",
    "0xdbd4279be7a36f2fb9f7a21f3d7adc68d89081e6",
    "0x9a635abb589ba8998b16fa700f853394db345c44",
    "0x44ec81a9ec9c8b3fc36b60d7241ee5585e516886",
    "0xd5f7be04fdbb0448efcb3971b2ea830ad9b726ec",
    "0x9b6a9a2cc809309b9a6db1f1a639f2860ace9bab",
    "0xe40ffa7e68d7ffdb2f49cc9df5330cf153b9624b",
    "0xf9ca0946edbc4eefd8d2d9b7133f303e368d95bf",
    "0xa158d1b2feb8f3f15a2e3c8612b49be45418ad8d",
    "0xe54f31aa6c6e11bec53e0007b683c7345ae5cf93",
    "0x8ece756bbbcf56139758af2dbc944bb1f1903a51",
    "0x6dcf2f2d0bb912d38d4bcab2d163f9671c5cb666",
    "0x91a824dfc5ff2eb34dcb396c36f31ea03f329f00",
    "0x38e73cd75e25a48d69f75bbdf76463f553b29fd2",
    "0xd235ab10105c316eb3b920646d3a6ca9bf22c8de",
    "0x856b7d152cf51efb424eef245949f0804a40dbf4",
    "0x18b387c16641047b06c22cc2696f53f28e43ab8f",
    "0xfcc86105f634ffec5d3653f6e45b467280594a5f",
    "0x2adccbb93e6a9ac7e9467db1b9098b2816abf4a8",
    "0x87c5ecb512769a6113e0b29d05c563634ec184c7",
    "0xa85f68b66c5c20a47e7bbbe1d5e7c68e14e5817e",
    "0x698b1dd89af48bd0f09159e2edab9ede167471e3",
    "0x50d0cbf116f0ab6c810bdb695c5d6336a1e8b5c9",
    "0x6eb57e3288fc7de842b2da77ad899f6cb30ed079",
    "0x648ff38a6efb4341738c6d990e74d5c8f0e9a9cb",
    "0x5e1d2f47704f47f80229b9379d593a173a2dccd4",
    "0x4adbdf91b776a67574f913d50eedbec7c807a69c",
    "0x2c4cc4f79670ae3007376ea8713babbedf6439b5",
    "0x139710365dde15e95abe7051017f5b527f5ed38a",
    "0x8ad6b266e10d7f8cd5dbc371d390d3f35f0a9493",
    "0xfaf09fe37b33df47874a57f6009d0694d195ab82",
    "0xa43da23aa1f73d01007519c3afff0a8ee2f14411",
    "0x6ff474634ef206187e138e580339ae392d4f0340",
    "0x5df96f606a5a799dadb0c8c44d190c43b24da0f9",
    "0x582bc5ab4d0a47b893625a0d5c8fe32f82a475cb",
    "0xba927181a191307e1ca505eb6f3e985b049abddd",
    "0x10eb4ba3a97357084aa5097de2eb7256f585aacd",
    "0xa598cde94cc24c390d497e98534695cfb65e6b5b",
    "0xba2cfdf6cbcb90a2ebb7c51a4e95b948cc1e7f0c",
    "0x9f3d7f5776e1dbb7db19f1c1124cfc2d415a1091",
    "0xee826e63f11e4099253617de82f77aa63e257d2f",
    "0x360560c2fca25160c6bb770f6f7fd7e33c8f5e14",
    "0xa5828825243819a46cd3d594537b561b14624e2c",
    "0xd4245690c6ebf1cab11ff6eb3ae01f0cbc8a6059",
    "0x1f776ce980ce81fba9f19b92d3b988e44577ed32",
    "0xc6cc3a499a5b9b527ae48f521faf7b2036156e10",
    "0x8a53ce1c4a23e9e7123be4979155fbb2cfafe837",
    "0xfe8129153c577afbc19f7f692b285142b7769f49",
    "0xeb916a921517d279a934969eb26517cce1cd28f3",
    "0x98905d90fa6fde50d92688799b3488420e77163d",
    "0x930b8234b73903c42d454bce23ba7937399ef6ba",
    "0xdd621fd15bbda32b30cb81c42980fffd9a51d9f8",
    "0xfc5ed957a682886b6c2ebdbae82ca8bb1bed9c1d",
    "0x2971454ead7f9dd73fd07c7958ce535ef69d90e0",
    "0xe1c3713007d9fe45544d237517c7666eab5e62da",
    "0x446cbc80911a705bbcb0ed98e2233178fd218958",
    "0x7ad89af734b2e822fc683a517e605048d862247e",
    "0x6259af8f40450d5e1dd9617d8681e64cd9258569",
    "0xee48d393af03ec09537d20278cfa904641a14627",
    "0x3814b89b1219a6ca709e5cf6e5e4f14c8e738af9",
    "0xc8af65573d4238baed7cf983c682ce0e055b133c",
    "0x91e88b8c6a22c97e6bb028e9a8e317602473732a",
    "0x59397ed01cd0dc43d27834a9fb4b1bd02d406d17",
    "0x4407e1ee49d26af0d53432ecbef955ad4a17bc12",
    "0x943a0174d66868d041e2a870bcd661255d0949ab",
    "0x13b21510b050a8405ad33ede724e28e58c211b64",
    "0x867f4ba06b61a67679c6d2d50a7c5c339fbbfefe",
    "0xb4bd3a4a5cd64c980bb9fd9671637f9e3f82ff14",
    "0x636f73985da0e3424a069b0cf2623426cf81ba08",
    "0xfc4acc547e7eda6868fc68667d5f51f38e794f08",
    "0x431b7bb06a889ddad8f45b57b45b9cd68b3bcc45",
    "0x43a90a32be48860c432f9ac3fdab39072e92f485",
    "0x01226bf1d9be7e5479088fc005c1107aeeca858d",
    "0xa1ea82b1c9f92c158f6420d7a0db01e8629e639c",
    "0x00af65bc042c073d758bca4c3cbab64db58d9fd5",
    "0x6e04ce0e366cdec84677ff4fa2b06461eb3e8c4b",
    "0x5e2f9c9b94f3f7f9026b60944aa288789db5faed",
    "0x6ad684cda672fafc40003921bbdd8903f4d39ae6",
    "0x5afd45da249c191683634b2b5d830ef42eb78206",
    "0x802e63811ed8d34346548d3cac9fc273b2ab547d",
    "0x944a383c151f29ab710c38cac90399f25f04159e",
    "0x6b41c4bbf80e7fecd51aa4a69d8091e083dd0e52",
    "0xe09e54d0ca0877400b0c4258cb0e231762775b04",
    "0xb939acddd396c320d20a9867802aeb79e17f56b5",
    "0xc8cc1fc1c85db24f11fc8bba0f5aae8892dc37eb",
    "0x74d60e479e6408e0129400ef38d173b3ce1c7a37",
    "0x56a8a9454aebdc9e637c56e2380aedaee52b748d",
    "0x6aa4a24cf2c78ac52197796b682477de837e5aed",
    "0x983fc283bd3761deb6380ff04aecaac7e2dac22e",
    "0xfa94a9f462680df64319cef2ba3a3c8ad422b99e",
    "0xe6187d5da69204e4d93c2cf429a510bd6754e0da",
    "0xd09a0a5533ad8d283638c6bd9228eed5ec41cef2",
    "0xa0b35624196403cb84382c1ce4e8b3d1035167b2",
    "0xb610fbc269568d31af5300706efdb7886f0d3f3a",
    "0xe70977c669d9b604a5440e6dcc33c32e6661e886",
    "0x9485e93b85bf9c8e289d550522fbf2c01574d668",
    "0x67e92deefd29b941fc39a681cffd0b19fc395fd7",
    "0x1ebbc838fdab8628ed4b69b46af8856df38c3442",
    "0x913ffdde90a483a324397a63311b078e89a66f2d",
    "0x41aa58da5b449bf301fcc93b533efda133209174",
    "0xc4249db2440056ff2121d84724bce54fffd2fe57",
    "0x5eb05322b5c7a115ec3987b416c6bc823bba4232",
    "0x1125d383ba08fd8a6a6446a40187cea72dfdaae0",
    "0x5a05468c3c299b5676e7d9aedad694b9a6fd98c9",
    "0x451fce3a8d695ab8b97297b950a4198160d6d7a1",
    "0xa5ada214b60ea311cf195c4ad73a87554f738c33",
    "0x92a78e82fb716376541cd4dde1bc8906e9b30a9c",
    "0x860db4ea830e9c2eeb2b37d4af6d98cd39bfb82d",
    "0xb19cc8bc4c575638f0dfb1cfd009d4b89ad355e6",
    "0x7f6a9f13e8d066b368338c46b1f8a18e7a9a395a",
    "0x371c6d563062e1a74bcea5a00223215821b7fe4e",
    "0x640dfa5cf45b00ac4448e901e990511340ab1158",
    "0xddcc2a3f2121f7cc688cbab05ba779b8b658f249",
    "0x258953e41ec7501f90e025ef2636e434a352388c",
    "0x7d71edc92c7bd533eb1538fb627e9360e9c58164",
    "0xabec9c2548a949b76180e59f12db17601f166258",
    "0x03768b3c78afdec1470c1e0dfdb4f8711ed27a7c",
    "0x9b32c5cd4474561482d5fd63bde49dadee2f8b87",
    "0x58266fdf59f8ce3008b7c9f52666b6a1b23bf3ce",
    "0xe64e376bb0d02197bc3924db0052a605f5092f15",
    "0x74c615a37b5c79ebadf82465647e38ac209c84bc",
    "0x91f20bd0e40efd1f1c31b711287e52a9cf634664",
    "0x49c4e317b8262f87012d6a6614c98066cf8b097a",
    "0x2789053be7dafc4dfe39dd0dc9d184c2be4d0a7b",
    "0xa707f53eb5064d0b3ae4f20f89580e2d7786ff62",
    "0x9fd4d20f89926fa3fa788b71a042ba14f8f68f41",
    "0x6429b0dd53c89cd629314f8a0ba54bd842125803",
    "0x1e8757451058388ed4c4c5c3aae4db3abf26a883",
    "0xdd4cf86fec603af280b9de7251fcbce40f3e9a18",
    "0x064379acfb0f8e39b5351302c0be24b094a2a800",
    "0x03af06240fd36ae257ee693c5d7a445a1a9d9950",
    "0x0f505c8e4b429e07dda1904846609ddf6ec0d4a4",
    "0xcdcd4e77d4881cf66f60482b830dfb7b3f221bc8",
    "0x0d05915936b479ca68a69ec6181446da7a2c12a2",
    "0xb9bd179a8f29d6c8af059b520a54e1af898a95dc",
    "0x3a0cd31b21d81765cd9e43f79b140da2e0fd8bbc",
    "0x0102fd479deae48b7dbbec5af9cc69638fb4fb45",
    "0xcaf5409d437d9127e168e440b723723ede6fe9c7",
    "0x7b27fd319ed471814358db12cf2db0c29e366fd1",
    "0x1bf823caa3be4a61b0dec06fa306d09f0733514a",
    "0xe3a244907f4cf7f31a788e155419b82399ff0f48",
    "0x8feb30d97fe021148176fbf35480e0c87134dd3f",
    "0x9b7297fe2d5d7de908767710a9113238581135aa",
    "0x6e10e15b31222064dc38c9a7ee2d41c51abc5cc6",
    "0x12e830d13827f408677141ae6ecf4c4b58997d2c",
    "0xc6e7a10ba32fa7ba06a7c6c4b23282d8924184cf",
    "0x428ba7db08d0e999a3af5a356709554e1958280e",
    "0x33118756f582a4f751ccd4dfd5f5dbd7f5932818",
    "0xa8ceab7cc7f105efddfed06ec6be80bf5d739c67",
    "0x65e119944efc6fba564c1e579ac0e678d581a9a2",
    "0xf92eb5ef8c1993e46bc6e1d3c9ad4605f9ed7b34",
    "0x2af593e5f17f2d9287a8fe73432585479198e0ec",
    "0xac7dd2a4e354705487c411a6d67fd00eaccb56bc",
    "0xf7c8589e33b153e9f7f4a8e0a85ba1d7bdf885e8",
    "0xd7d2d00180099478c8c970d96e1792dca19ac8e9",
    "0x17265eef6c0c11b602e5546dd3eaae5195ab3fdc",
    "0xe8f20a20eb94f7d6ee9f3a9b8f8221a334ee5bcc",
    "0x0b8ca6c68e50c74cf85329d727a798aec57717ea",
    "0x330dd840c76ded3447dd772581def66815c6d90e",
    "0xa3b39b84788ccd6ab6d78bf7bd7cb6f86f76d1a1",
    "0x580bb1ebad582008c793f2f09046bbe9e4713896",
    "0xe9e8df9444894ec6bf15aead4122446891a89930",
    "0xdbb3570c341fb6634609b8c742941c91163d16f8",
    "0x6880262527b3255a77d40ca7979689dcfcc7ca42",
    "0x09cf60061cd0ba29b38e4ccefd63d4433d2b4129",
    "0xebb9bb31ad7500b583f39ae9f763c1e4af50bd0a",
    "0xc2de56851f43958f07175fb9c7c3e7040eaccf81",
    "0x201fc7d3609f698b592978defdefadbf2cb1eb3d",
    "0x731b6616e656c4f4937d77f083caa55ab50b6e5e",
    "0xcac9a010abbc0cd8d21d1c8cf3e86fe9fe54f477",
    "0x1d8626e04c6bda93425612199b73c3c3cdeba7da",
    "0x942174bef0a6b024cf10bfdbc9fe52fcb58414ef",
    "0x01d337064074ba56dc6d9541448284814c59d523",
    "0xdc341721dff7998cbe8008396c798062f448d68d",
    "0xe08dc479683b83ef37182497cbdc2be6c72f653c",
    "0x156c6183dbce64d60a0299ba5e23ffc5d23845e3",
    "0x084554c2d9cdd87a4c2c2c32d2cfbc2ebd1c0917",
    "0x2579048ec530e043dc60fd5662c47640c0e15126",
    "0xc6b0326eb375b00d1adcaa1882d83d1be8b5837e",
    "0xcf81ee8cfa4aca6e304c9cade7a301ca9c2cf146",
    "0x3da61f49ef5c3e6a5523209f70d6fc6280326eb1",
    "0x762ffcd0c7eacb3c5bd747742ccd185b06625240",
    "0xe0a89c63e382de82cfdcab75c666b3d92ca9188c",
    "0xabe19f1bac4ac69b1782c87d21de5a920bbab249",
    "0xc13c12b31e9777f8c4d6fec7b6f6f3feff925b7f",
    "0x051776a593601fcc09f94000ab61e12a9a95f95a",
    "0x70f09c9f65ef67ee4b1587ad2845e1881a126d4b",
    "0xf8c9cf0ff6b5c052ea99904e44dfa4457eadc0b5",
    "0xbb1915c4289ec59dd64150f6a220ecec3b0ab552",
    "0xc1be8d415d0ea9e389ea5cefec4d5dd15c041540",
    "0xccb7609bba8801bc4d00947262b31c4f0478662b",
    "0x69291ddba25108e49c4cf9d9d19d3fb62320625e",
    "0xe559e779caaf709f7307319bf9797913caf0bc87",
    "0x54dda6b71eb4af6191daa05f88dca4489c5e574f",
    "0x527bc0ff54840d2c244142a7f7eeac8a58d41645",
    "0x90ed88c5341f37c326bd43426371b6024c2b165c",
    "0xbb7856e8721121ad712974c08842ea950a104b00",
    "0xf1c82ec177e3f99db3f39097da41f327ed395be3",
    "0x979e0f566a955b4934199e249caf9ed89452c9f8",
    "0x154612a77bae8120d6d1366922a1f93b02912a4f",
    "0x185508c994ee5d76de0a061cd4142d2b743886aa",
    "0x6eb7a645a5a429c6787696a0e26e35fcf12209ae",
    "0x2739b8c4c67afc6fc2834320e20c4661025201d7",
    "0xb5391f0f5adf22dc8bf4f5b8a4d516d0180d978a",
    "0x674900dcefce6de57f08c2355fe368703b90c8ee",
    "0x619d5edbb036bb3cfb99540d81275579f709a198",
    "0x6a8d3be6e882033f7ac7b55b0540135393454492",
    "0x4fd5e8bb9882fc5c7f5dd966e72f49f0e1cf8b8c",
    "0x8752aeb8435b11f2847fa71da2ef6c69c2c7f4a4",
    "0xf3f46b98f196d315fc649f680c9ff16c6cb2e1e1",
    "0xb2335053f16957ad4b347ebe0b7644c5c3b3ddf0",
    "0x355fd8675806ab126b2d3b30c36851af42d5832b",
    "0x8eebfd6609d9f45eb6018d0286429869913c705d",
    "0x920f2612a1a2b8db99000d33ed63c54ba7977439",
    "0x29e074043ace527d8f67675a656ed2be381dc16e",
    "0x8644a13a9b3978522d10e93bf0e2e31699e289b4",
    "0x7cbeea2adbca4b47b43138ab2f65b997668f10a6",
    "0x0516c69ee11ac84b9668bdd785c6c2d69cdafb75",
    "0xb91d4ed4ebb1662c99570d3d65bdc8d4b8d07d8e",
    "0x785bf3175cb90c6ea39c3c2e6289ca8b4c9d62ec",
    "0x64c87b400fdaaa66b9193fcdb80bfb9a57384c79",
    "0x7461bd146fa2473cb7a86952335bccb78580c12e",
    "0x30d28e4051b328bdc4ed05cb4d6fe43d276e3e66",
    "0x0737405fbd6aa4c93de56ba4c9d6cf1082e55279",
    "0xfa72dbd6c24830da6425d58f134d0f873a6b590b",
    "0x02953305cd8fce44ac882d116d84364ca2cc2020",
    "0x8cad2839988aafd61c4dcd5c4c10ba21a7536354",
    "0x2dc7273e4e50fe95cc89dc1e9cb1fe60e5640dc1",
    "0x7cb3d6d6779641f8ca7d3d4511fced923430fb7f",
    "0x45bc79ec8778a1ee402f44960a4ec58fbfa8e921",
    "0x6af5d0f3fd6ab74068a0cdcde3ea3f6026f88396",
    "0xb28203424e1921b6f3006ad5175da1a377f63159",
    "0xfa94b52a9ff0189d5c322e70057abd9a6ff2b9de",
    "0x0889f3f75b3046939e7626f03e99f39b134841db",
    "0x8b49e0ec57480fa31dbc1c4cdc4322f119a14314",
    "0x5b4af61d126f703cf46912336ead53db6e4b8b50",
    "0xfa1e1eac2a035c5e91f93eff3681c2bfce7f4104",
    "0xe739d97e9b8f168d18686f55adb0927788878345",
    "0xa29f5da5688e1d479bf66af4c69799bbd1ac6dd6",
    "0x301fd241c626463307058dfb8af54806a1f95141",
    "0xb4d1951829920cfb01cbf36282dd4b16135524bf",
    "0xb4e24678089ab9acff44c15b190fadfe2bc31f3b",
    "0xe89c454aac720088004aedfcb860e2bfdd299ff3",
    "0xce8dc2a35f0ccdd8d26fc9eacf5b696c373b3ac4",
    "0x8add7985189784e6445d20c3e20c93cc3480ae38",
    "0xd15d3fc1c006754644fae348ee77b19660bdedf3",
    "0x30b61b77e54524cbd1d2402df23175b5fabd4f4a",
    "0x965ba376d57a264c551e9b0904ff948640e9ecd7",
    "0xcf0cb0ed3a7571160d66facd0ab0b8e48ff66fd8",
    "0x8c94d85be1e4baf9aa4d5e9f30790ac58b528c9f",
    "0x9b1d3c5bf854f291722b6c09779cf294753a5ca1",
    "0x948e10a6e14313d4b7d10c4b39e018aa24c715dc",
    "0x52726ff64f3c471cdaab538e9e907766bb756dea",
    "0x9b3a5cefb4f1fa68b04dadaf9c3720cfd172e6de",
    "0x7e601189d394995f473df079b581297b719176b4",
    "0x75a8ba162bbe6ba0d795650147b84f770099e987",
    "0xc14ebf355fbc59017a9469224a341cf6c8a062d3",
    "0x6b1b7e2a5cb5850cc18eed3528339707881746d1",
    "0xeb00a32b5f7115b334180c488a480aae70fa678e",
    "0x5501abc44760b7f2dc44756e5fe30fd3754b35e0",
    "0x2d93b90933f0498daeda86c8c79d1c3b94609672",
    "0xd84310751005057a12bba270ed4d88fe0b760f5c",
    "0xd3a19d9bf34b118a72c7a17836726141d670c624",
    "0xce054d95dd029aa9e5612cf22939f068d4f757fb",
    "0x03a99bdd62879ca8685477b9bb470d2295622769",
    "0xdd330a10f6ca16ce2e334e4a6c36bb1d3e8c9514",
    "0xa8066df6d8ac700f4f7a282d7c3525ad9f3e501f",
    "0x4701099bb9089bec8ac90e0732f4e8d912ab6dfa",
    "0x52c9fe06a1d1a6869800a415eb6cdec1bfc1331a",
    "0x446e282a288f1845f4ecb2fc29f77f986f70e8fe",
    "0x542facfa0d1c836b9a1b0c862bf885db969ac0f6",
    "0xa388612847d1c43af6e1f34f6b5a37a34bf06af9",
    "0xe0392f9b0a6c4229bbf4e4243ced5b61ccd6d60f",
    "0xa197f6de874a16b515c69f44fbf8d9e46678909b",
    "0x194b29e85a5b464223744c494692c9966b253cee",
    "0x5738f18da697590eee13d3762674d11d71299e3d",
    "0x2e0e175673a99450cba2580df35783c6460bc0dc",
    "0x82c11e76a1f9d2d102acf9df89a06b60f723742e",
    "0x17ec7a931cfa4b2246b554d940d7b010db14e591",
    "0x9eea84027f68b6a0b01d4c015ed0fe9fffc0b88e",
    "0x92861681102120d9c1815b17fcc0f9ec1b7c0ea7",
    "0xd6f777a2dd0ad89688f4aa8e181dd1eb2417a20f",
    "0x2d918ed3a1372d64f79a1cbae661ec8882100801",
    "0xd1ba5511be1c21c882f549a0635943ddf76eddb7",
    "0x4d74eea220d2c2829fb65823fee87963ae0604b3",
    "0xe7bd0ab4ff45faadd63e8210888e8c23d88a563a",
    "0x9cc255d04bfe1af73e4b7fdfe91ec3228e41a06f",
    "0x975de6b32fb01cb800e7e7eeed51ead04b7031ad",
    "0x3c49829a61454f340168d99ef58c9f7c271fe20d",
    "0xd20e70821d465632c5c4d2c91a55952dc12a8662",
    "0x9f4d00a580c7846da6feeffe1d76b37d379f4b04",
    "0x7fc8325c3f1705c6648dafeaf100fe750ef3ae5a",
    "0x6c201167d8f9cc5759163ebdc7059c3a51e4d164",
    "0x7ac67cad12b9469b994119e02ce734c84cfe9477",
    "0xcb9b4e8eece82b9df55d2bf366db5fd35546b444",
    "0xd9f678dcb18b12aeeed92badb901290e107a830f",
    "0xe687cb908c0e32faa1f46da56f87205aff50c4cd",
    "0xda9f402691eadacae2d6b314b11155616ea25a44",
    "0x8e1c30995159047cbdc9f9de78890bf060f97c18",
    "0x2b66e96848fe524f3f03a1a6ff30892f3dec3587",
    "0xd037628297f8c56b179a4532514e312cf3bf22cb",
    "0x7d494aac55f14c1c0052a2f227ec3a9eb0eeffb2",
    "0x0a02435be5fca6dbc56dd3863de0dd58d834d851",
    "0xfdc9dedfdb86e119a80ef812182557fa59e29486",
    "0xa68dd9bc8fcb2602774c987e49b3fea7653e745e",
    "0x4ad6483e248393096ae6afcca2c998254ac48f35",
    "0xb4d4a7bbc9bb53a871b933f40ee1621fc2f05304",
    "0x8a8317704b8bbad8a677780731ab81bb793f3087",
    "0x206271c40348ec6c7f40349b2184a553b4ee9f07",
    "0x60b453f9659d9fe14c52f38a06d525f3f9247574",
    "0x0c17d9bed88de881ecdb6878e78dcd8a9b4321ed",
    "0x5624636a4f275bb2652b84325d6bdd7d0c9f3bfe"
  ]
}
```

### eth_blobBaseFee

The `blobBaseFee` method returns the expected base fee for blobs in the next block.

Parameters:

None

Returns:

result: big.Int, The expected base fee in wei, represented as a hexadecimal.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_blobBaseFee"
}' | jq
```

Response:

```shell
{
  "id": 1001,
  "jsonrpc": "2.0",
  "error": {
    "code": -32601,
    "message": "the method eth_blobBaseFee does not exist/is not available",
    "data": {
      "trace_id": "5cdd0e1ca22f47d1fb98bf119fa0e51e"
    }
  }
}
```

### eth_blockNumber

The `blockNumber` method returns the current latest block number.

Parameters:

None

Returns:

result: uint64, A hexadecimal of an integer representing the current block number the client is on.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_blockNumber"
}' | jq
```

Response:

```shell
{
  "id": 1001,
  "jsonrpc": "2.0",
  "result": "0x54de0d7"
}
```

### eth_call

The `call` method executes a new message call immediately, without creating a transaction on the block chain. Often used for executing read-only smart contract functions, for example the balanceOf for an ERC-20 contract.

Parameters:

- args: object TransactionArgs, required
- blockNrOrHash: object BlockNumberOrHash, optional
- overrides: object StateOverride, optional

Returns:

result: array of byte, the return value of executed contract.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 8001,
  "method": "eth_call",
  "params": [
    {
      "to": "0x0000000000000000000000000000000000000088",
      "data": "0x0db02622"
    },
    "latest"
  ]
}' | jq
```

Response:

```shell
{
"id": 8001,
"jsonrpc": "2.0",
"result": "0x00000000000000000000000000000000000000000000000000000000000000d4"
}

```

### eth_chainId

The `chainId` method returns the currently configured chain ID, a value used in replay-protected transaction signing as introduced by EIP-155.

Parameters:

None

Returns:

result: uint64, a hexadecimal of the current chain ID.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_chainId"
}' | jq
```

Response:

```shell
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0x32"
}

```

### eth_coinbase

The `coinbase` method returns the client coinbase address. The coinbase address is the account to pay mining rewards to. This is the alias for `eth_etherbase`.

Parameters:

None

Returns:

result: address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_coinbase"
}' | jq
```

Response:

```shell
TODO:disabled
```

### eth_createAccessList

The `createAccessList` method creates an EIP2930 type accessList based on a given Transaction. The accessList contains all storage slots and addresses read and written by the transaction, except for the sender account and the precompiles. This method uses the same transaction call Transaction Call Object and blockNumberOrTag object as eth_call. An accessList can be used to unstuck contracts that became inaccessible due to gas cost increases.

Parameters:

- args: object transactionArgs, required
  - from: optional, 20 bytes. The address of the sender.
  - to: 20 bytes. address the transaction is directed to.
  - gas: optional, hexadecimal value of the gas provided for the transaction execution.
  - gasPrice: optional, hexadecimal value gas price, in wei, provided by the sender. The default is 0. Used only in non-EIP-1559 transactions.
  - maxPriorityFeePerGas: optional, maximum fee, in wei, the sender is willing to pay per gas above the base fee. See EIP-1559 transactions. If used, must specify maxFeePerGas.
  - maxFeePerGas: optional, maximum total fee (base fee + priority fee), in wei, the sender is willing to pay per gas. See EIP-1559 transactions. If used, must specify maxPriorityFeePerGas.
  - value: optional, hexadecimal of the value transferred, in wei.
  - data: optional, hash of the method signature and encoded parameters. See Ethereum contract ABI specification.
- blockNrOrHash: BlockNumberOrHash, optional, a string representing a block number, block hash, or one of the string tags
  - latest
  - earliest
  - pending
  - finalized.

Returns:

result: object accessListResult:

- accessList: A list of objects with the following fields:
  - address: Addresses to be accessed by the transaction.
  - storageKeys: Storage keys to be accessed by the transaction.
- gasUsed: A hexadecimal string representing the approximate gas cost for the transaction if the access list is included.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_createAccessList",
  "params": [
    {
      "from": "0x3bc5885c2941c5cda454bdb4a8c88aa7f248e312",
      "data": "0x20965255",
      "gasPrice": "0x3b9aca00",
      "gas": "0x3d0900",
      "to": "0x00f5f5f3a25f142fafd0af24a754fafa340f32c7"
    },
    "latest"
  ]
}' | jq
```

Response:

```shell

```

### eth_etherbase

The `etherbase` method returns the client coinbase address. The etherbase address is the account to pay mining rewards to.

Parameters:

None

Returns:

result: address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_etherbase"
}' | jq
```

Response:

```shell

```

### eth_estimateGas

The `estimateGas` method generates and returns an estimate of how much gas is necessary to allow the transaction to complete. The transaction will not be added to the blockchain. Note that the estimate may be significantly more than the amount of gas actually used by the transaction, for a variety of reasons including EVM mechanics and node performance.

Parameters:

- args: object TransactionArgs, required
- blockNrOrHash: object BlockNumberOrHash, optional
- overrides: object StateOverride, optional

Returns:

result: uint64

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_estimateGas",
  "params": [
    {
      "from": "0xD4CE02705041F04135f1949Bc835c1Fe0885513c",
      "to": "0x85f33E1242d87a875301312BD4EbaEe8876517BA",
      "value": "0x1"
    }
  ]
}' | jq
```

Response:

```shell

```

### eth_feeHistory

The `feeHistory` returns transaction base fee per gas and effective priority fee per gas for the requested block range.

Parameters:

- blockCount math.HexOrDecimal64, required, Number of blocks in the requested range. Between 1 and 1024 blocks can be requested in a single query. If blocks in the specified block range are not available, then only the fee history for available blocks is returned.
- lastBlock: BlockNumber, required, integer representing the highest number block of the requested range, or one of the string tags `latest`, `earliest`, or `pending`.
- rewardPercentiles: array of integers, optional, a monotonically increasing list of percentile values to sample from each block's effective priority fees per gas in ascending order, weighted by gas used.

Returns:

result: object feeHistoryResult

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_feeHistory",
  "params": [
    "0x3",
    "latest",
    [20,50]
  ]
}' | jq
```

Response:

```shell

```

### eth_gasPrice

The `gasPrice` method returns the current gas price in wei.

Parameters:

None.

Returns:

result: big.Int

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1003,
  "method": "eth_gasPrice"
}' | jq
```

Response:

```shell

```

### eth_getBalance

The `getBalance` returns the balance of the account of a given address. The balance is in wei.

Parameters:

- address: address, required, a string representing the address (20 bytes) to check for balance.
- blockNrOrHash: object BlockNumberOrHash, required, a hexadecimal block number, or one of the string tags latest, earliest, pending, or finalized.

Returns:

result: big.Int

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1003,
  "method": "eth_getBalance",
  "params": [
    "0xD4CE02705041F04135f1949Bc835c1Fe0885513c",
    "latest"
  ]
}' | jq
```

Response:

```shell

```

### eth_getBlockByHash

The `getBlockByHash` returns information about a block whose hash is in the request.

Parameters:

- blockHash: hash, required, block hash
- fullTx: bool, required, if true returns the full transaction objects, if false returns only the hashes of the transactions

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1003,
  "method": "eth_getBlockByHash",
  "params": [
    "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
    true
  ]
}' | jq
```

Response:

```shell

```

### eth_getBlockByNumber

The `getBlockByNumber` method returns information about a block by block number.

Parameters

- blockNr: BlockNumber, integer of a block number, or the string "earliest", "latest", "pending", or "finalized", as in the default block parameter.
- fullTx: bool, if true returns the full transaction objects, if false only the hashes of the transactions.

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC}  -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getBlockByNumber",
  "params": [
    "latest",
    true
  ]
}' | jq
```

Response:

```shell

```

### eth_getBlockReceipts

The `getBlockReceipts` returns the block receipts for the given block hash or number or tag.

Parameters:

- blockNrOrHash: BlockNumberOrHash, required, hexadecimal or decimal integer representing a block number, or one of the string tags:
  - latest
  - earliest
  - pending
  - finalized

note: pending returns the same data as latest.

Returns:

result: object, block object or null when there is no corresponding block.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_getBlockReceipts",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell

```

### eth_getBlockTransactionCountByHash

The `getBlockTransactionCountByHash` method returns the number of transactions in the block with the given block hash.

Parameters:

- blockHash: hash, required, block hash

Returns:

result: uint, block transaction count

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_getBlockTransactionCountByHash",
  "params": [
    "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce"
  ]
}' | jq
```

Response:

```shell

```

### eth_getBlockTransactionCountByNumber

The `getBlockTransactionCountByNumber` method returns the number of transactions in the block with the given block number.

Parameters:

- blockNr: BlockNumber, required, block number, or one of the string tags latest, earliest, pending, or finalized.

Returns:

result: uint, block transaction count

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_getBlockTransactionCountByNumber",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell

```

### eth_getCode

The `getCode` method returns the compiled byte code of a smart contract, if any, at a given address.

Parameters:

- address: address, required
- blockNrOrHash: BlockNumberOrHash, required

Returns:

result: array of byte

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_getCode",
  "params": [
    "0x0000000000000000000000000000000000000088",
    "latest"
  ]
}' | jq
```

Response:

```shell

```

### eth_getLogs

The `getLogs` method returns an array of all the logs matching the given filter object.

Parameters:

- crit: ojbect FilterCriteria, a filter object containing the following:

- address: optional, contract address (20 bytes) or a list of addresses from which logs should originate.
- fromBlock: optional, default is "latest", a hexadecimal block number, or one of the string tags latest, earliest, pending, safe, or finalized. See the default block parameter.
- toBlock: optional, default is "latest", a hexadecimal block number, or one of the string tags latest, earliest, pending, safe, or finalized. See the default block parameter.
- topics: optional, array of 32 bytes DATA topics. Topics are order-dependent.
- blockhash: optional, restricts the logs returned to the single block referenced in the 32-byte hash blockHash. Using blockHash is equivalent to setting fromBlock and toBlock to the block number referenced in the blockHash. If blockHash is present in the filter criteria, then neither fromBlock nor toBlock are allowed.

Returns:

result: array of Log

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_getLogs",
  "params": [
    {
      "address": "0x53350795c11cee781a7e174479778f848d76ab2a",
      "fromBlock": "0x22b2277",
      "toBlock": "0x22b2277",
      "topics": [
        [
          "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925",
          "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
          "0x6a12b3df6cba4203bd7fd06b816789f87de8c594299aed5717ae070fac781bac"
        ]
      ]
    }
  ]
}' | jq
```

Response:

```shell

```

### eth_getOwnerByCoinbase

The `getOwnerByCoinbase` return masternode owner of the given coinbase address.

Parameters:

- coinbase: address, required, account
- blockNr: BlockNumber, required, block number

Returns:

result: address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getOwnerByCoinbase",
  "params": [
    "0xD4CE02705041F04135f1949Bc835c1Fe0885513c",
    "latest"
  ]
}' | jq
```

Response:

```shell

```

### eth_getProof

The `getProof` returns the account and storage values of the specified account including the Merkle-proof. The block number can be nil, in which case the value is taken from the latest known block.

Parameters:

- account: address, required
- keys: array of string, required
- blockNumber: big.Int, optional

Returns:

result: object AccountResult

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getProof",
  "params": [
    "0xe5cB067E90D5Cd1F8052B83562Ae670bA4A211a8",
    [
      "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
      "0x283s34c8e2b1456f09832c71e5d6a0b4f8c9e1d3a2b5c7f0e6d4a8b2c1f3e5d7"
    ],
    "latest"
  ],
}' | jq
```

Response:

```shell

```

### eth_getStorageAt

The `getStorageAt` method returns the value from a storage position at a given address.

Parameters:

- address: address, required
- key: string, required
- blockNrOrHash: BlockNumberOrHash, required

Returns:

result: array of byte

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getStorageAt",
  "params": [
    "0xfe3b557e8fb62b89f4916b721be55ceb828dbd73",
    "0x0",
    "latest"
  ],
}' | jq
```

Response:

```shell

```

### eth_getRawTransactionByBlockHashAndIndex

Teh `getRawTransactionByBlockHashAndIndex` method returns the bytes of the transaction for the given block hash and index.

Parameters:

- blockHash: hash, required, block hash
- index: uint, required, transaction index

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getRawTransactionByBlockHashAndIndex",
  "params": [
    "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
    0
  ]
}' | jq
```

Response:

```shell

```

### eth_getRawTransactionByBlockNumberAndIndex

The `getRawTransactionByBlockNumberAndIndex` returns the bytes of the transaction for the given block number and index.

Parameters:

- blockNr: BlockNumber, required, blcok number
- index: uint, required, transaction index

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getRawTransactionByBlockNumberAndIndex",
  "params": [
    "latest",
    0
  ]
}' | jq
```

Response:

```shell

```

### eth_getRawTransactionByHash

The `getRawTransactionByHash` method returns the bytes of the transaction for the given hash.

Parameters:

- hash, required, transaction hash

Returns:

result: array of byte

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getRawTransactionByHash",
  "params": [
    "0x5bbcde52084defa9d1c7068a811363cc27a25c80d7e495180964673aa5f47687"
  ]
}' | jq
```

Response:

```shell

```

### eth_getRewardByHash

The `getRewardByHash` method returns the reward by block hash.

Parameters:

- hash, required, block hash

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getRewardByHash",
  "params": [
    "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce"
  ]
}' | jq
```

Response:

```shell

```

### eth_getTransactionAndReceiptProof

The `getTransactionAndReceiptProof` method returns the Trie transaction and receipt proof of the given transaction hash.

Parameters:

- hash, required, transaction hash

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getTransactionAndReceiptProof",
  "params": [
    "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9"
  ]
}' | jq
```

Response:

```shell

```

### eth_getTransactionByBlockHashAndIndex

The `getTransactionByBlockHashAndIndex` method returns information about a transaction given block hash and transaction index position.

Parameters:

- blockHash: hash, required, a string representing the hash (32 bytes) of a block
- index: uint, required, a hexadecimal of the integer representing the position in the block

Returns:

result: object RPCTransaction

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getTransactionByBlockHashAndIndex",
  "params": [
    "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
    "0x0"
  ]
}' | jq
```

Response:

```shell

```

### eth_getTransactionByBlockNumberAndIndex

The `getTransactionByBlockNumberAndIndex` method returns information about a transaction given block number and transaction index position.

Parameters:

- blockNr: BlockNumber, required, a hexadecimal block number, or one of the string tags latest, earliest, pending, or finalized
- index: uint, required, a hexadecimal of the integer representing the position in the block

Returns:

result: object RPCTransaction

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getTransactionByBlockNumberAndIndex",
  "params": [
    "0x548f4f1",
    "0x0"
  ]
}' | jq
```

Response:

```shell

```

### eth_getTransactionByHash

The `getTransactionByHash` method returns information about a transaction for a given hash.

Parameters:

- hash: hash, required, a string representing the hash (32 bytes) of a transaction

Returns:

result: object RPCTransaction

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getTransactionByHash",
  "params": [
    "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9"
  ]
}' | jq
```

Response:

```shell

```

### eth_getTransactionCount

The `getTransactionCount` method returns the number of transactions sent from an address.

Parameters:

- address: address, required, a string representing the address (20 bytes)
- blockNrOrHash: BlockNumberOrHash, required, a hexadecimal block number, or one of the string tags latest, earliest, pending, or finalized.

Returns:

result: uint64

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getTransactionCount",
  "params": [
    "0xD4CE02705041F04135f1949Bc835c1Fe0885513c",
    "latest"
  ]
}' | jq
```

Response:

```shell

```

### eth_getTransactionReceipt

The `getTransactionReceipt` method returns the receipt of a transaction given transaction hash. Note that the receipt is not available for pending transactions.

Parameters:

- hash: hash, required, a string representing the hash (32 bytes) of a transaction

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 5002,
  "method": "eth_getTransactionReceipt",
  "params": [
    "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9"
  ]
}' | jq
```

Response:

```shell

```

### eth_getUncleByBlockHashAndIndex

The `getUncleByBlockHashAndIndex` method returns information about an uncle of a block given the block hash and the uncle index position.

Parameters:

- blockHash: hash, required, a string representing the hash (32 bytes) of a block.
- index: uint, required, a hexadecimal equivalent of the integer indicating the uncle's index position.

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 5002,
  "method": "eth_getUncleByBlockHashAndIndex",
  "params": [
    "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
    "0x0"
  ]
}' | jq
```

Response:

```shell

```

### eth_getUncleByBlockNumberAndIndex

The `getUncleByBlockNumberAndIndex` method returns information about an uncle of a block given the block number and the uncle index position.

Parameters:

- blockNr: BlockNumber, required, a hexadecimal block number, or one of the string tags latest, earliest, pending, or finalized
- index: uint, required, a hexadecimal equivalent of the integer indicating the uncle's index position

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 5002,
  "method": "eth_getUncleByBlockNumberAndIndex",
  "params": [
    "0x548f4f1",
    "0x0"
  ]
}' | jq
```

Response:

```shell

```

### eth_getUncleCountByBlockHash

The `getUncleCountByBlockHash` method returns the number of uncles in a block from a block matching the given block hash.

Parameters:

- blockHash: hash, required, a string representing the hash (32 bytes) of a block

Returns:

result: uint

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 5002,
  "method": "eth_getUncleCountByBlockHash",
  "params": [
    "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce"
  ]
}' | jq
```

Response:

```shell

```

### eth_getUncleCountByBlockNumber

The `getUncleCountByBlockNumber` method returns the number of uncles in a block from a block matching the given block number.

Parameters:

- blockNr: BlockNumber, required, a hexadecimal block number, or one of the string tags latest, earliest, pending, or finalized

Returns:

result: uint

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 5002,
  "method": "eth_getUncleCountByBlockNumber",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell

```

### eth_getWork

The `getWork` method returns the hash of the current block, the seed hash, and the boundary condition to be met ("target").

Parameters:

None

Returns:

result: array of string, with the following properties:

- Current block header PoW-hash (32 bytes).
- The seed hash used for the DAG (32 bytes).
- The boundary condition ("target") (32 bytes), 2^256 / difficulty.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 5002,
  "method": "eth_getWork"
}' | jq
```

Response:

```shell

```

### eth_hashrate

The `hashrate` method returns the number of hashes per second that the node is mining with. Only applicable when the node is mining.

Parameters:

None

Returns:

result: uint64

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 5002,
  "method": "eth_hashrate"
}' | jq
```

Response:

```shell

```

### eth_maxPriorityFeePerGas

The `maxPriorityFeePerGas` method returns an estimate of how much priority fee, in wei, you need to be included in a block.

Parameters:

None

Returns

result: big.Int, a hexadecimal value of the priority fee, in wei, needed to be included in a block.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1002,
  "method": "eth_maxPriorityFeePerGas"
}' | jq
```

Response:

```shell

```

### eth_mining

The `mining` method returns true if client is actively mining new blocks.

Parameters:

None

Returns

result: bool

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 5002,
  "method": "eth_mining"
}' | jq
```

Response:

```shell

```

### eth_pendingTransactions

The `pendingTransactions` returns the transactions that are in the transaction pool and have a from address that is one of the accounts this node manages.

Parameters:

None

Returns:

result: array of RPCTransaction

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_pendingTransactions"
}' | jq
```

Response:

```shell

```

### eth_protocolVersion

The `protocolVersion` method returns the current Ethereum protocol version.

Parameters:

None

Returns:

result: uint

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1004,
  "method": "eth_protocolVersion"
}' | jq
```

Response:

```shell

```

### eth_resend

The `resend` method accepts an existing transaction and a new gas price and limit. It will remove the given transaction from the pool and reinsert it with the new gas price and limit.

Parameters:

- sendArgs: object TransactionArgs, required, the arguments to construct a new transaction
- gasPrice: big.Int, optional, gas price
- gasLimit: uint64, optional, gas limit

Returns:

result: hash, transaction hash

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_resend",
  "params":[
    {
      "from": "0xca7a99380131e6c76cfa622396347107aeedca2d",
      "to": "0x8c9f4468ae04fb3d79c80f6eacf0e4e1dd21deee",
      "value": "0x1",
      "gas": "0x9999",
      "maxFeePerGas": "0x5d21dba00",
      "maxPriorityPerGas": "0x5d21dba00"
    },
    "0x5d21dba99",
    "0x5d21dba99"
  ]
}' | jq
```

Response:

```shell

```

### eth_sendRawTransaction

The `sendRawTransaction` method submits a pre-signed transaction for broadcast to the Ethereum network.

Parameters:

- input: array of byte

Returns:

result: hash

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_sendRawTransaction",
  "params":[
    "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
  ]
}' | jq
```

Response:

```shell

```

### eth_sendTransaction

The `sendTransaction` method creates new message call transaction or a contract creation, if the data field contains code, and signs it using the account specified in from.

Parameters:

- args: object TransactionArgs

Returns:

result: hash

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_sendTransaction",
  "params":[
    {
      from: "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
      to: "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
      gas: "0x76c0",
      gasPrice: "0x9184e72a000",
      value: "0x9184e72a",
      input: "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
    }
  ]
}' | jq
```

Response:

```shell

```

### eth_sign

The `sign` method calculates an Ethereum specific signature with: `sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))`.

By adding a prefix to the message makes the calculated signature recognizable as an Ethereum specific signature. This prevents misuse where a malicious dapp can sign arbitrary data (e.g. transaction) and use the signature to impersonate the victim.

Note: the address to sign with must be unlocked.

Parameters:

- addr: address, required, account address
- data: array of byte, required, message to sign

Returns:

result: array of byte

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_sign",
  "params":[
    "0xD4CE02705041F04135f1949Bc835c1Fe0885513c",
    "0x1234abcd"
  ]
}' | jq
```

Response:

```shell

```

### eth_signTransaction

The `signTransaction` method signs a transaction that can be submitted to the network at a later time using with `eth_sendRawTransaction`.

Parameters:

- args: object TransactionArgs, required
  - nonce: uint64, optional, anti-replay parameter
  - to: address, optional, recipient address, or null if this is a contract creation transaction
  - from: address, required, sender address
  - value: big.Int, optional, value to be transferred, in wei
  - data: array of byte, optional, compiled code of a contract or hash of the invoked method signature and encoded parameters
  - input: same as data
  - gas: uint64, optional, gas provided by the sender
  - gasPrice: big.Int, optional, gas price, in wei, provided by the sender
  - maxPriorityFeePerGas: big.Int, optional, maximum fee, in wei, the sender is willing to pay per gas above the base fee
  - maxFeePerGas: big.Int, optional, maximum total fee (base fee + priority fee), in wei, the sender is willing to pay per gas.
  - accessList: array of object, optional, list of addresses and storage keys the transaction plans to access
  - chainId: big.Int, optional, chain ID

Returns:

result: object SignTransactionResult

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_signTransaction",
  "params": [
    {
      "data":"0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
      "from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
      "gas": "0x76c0",
      "gasPrice": "0x9184e72a000",
      "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
      "value": "0x9184e72a"
    }
  ]
}' | jq
```

Response:

```shell

```

### eth_submitWork

The `submitWork` method can be used by external miner to submit their POW solution. It returns an indication if the work was accepted.

Note, this is not an indication if the provided work was valid!

Parameters:

- nonce: BlockNonce, required
- solution: hash, required
- digest: hash, required

Returns:

result: bool

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_submitWork",
  "params": [
   "0x0000000000000001",
   "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
   "0xD1FE5700000000000000000000000000D1FE5700000000000000000000000000"
  ]
}' | jq
```

Response:

```shell

```

### eth_syncing

The `syncing` method returns an object with data about the sync status or false.

Parameters:

None

Returns:

result: bool

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_syncing"
}' | jq
```

Response:

```shell

```

### Filter methods

#### eth_getFilterChanges

The `getFilterChanges` method polling method for a filter, which returns an array of logs which occurred since the last poll. Filter must be created by calling either `eth_newFilter` or `eth_newBlockFilter`.

Parameters:

- id: string, required, a string denoting the filter ID

Returns:

result: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getFilterChanges",
  "params": [
   "0x68ce60ffdb0c9480c307b0c3d2ae9391"
  ]
}' | jq
```

Response:

```shell
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}

```

#### eth_getFilterLogs

The `getFilterLogs` method returns an array of all logs matching the filter with the given filter ID.

Parameters:

- id: string, required, a string denoting the filter ID

Returns:

result: array of Log, Log objects contain the following keys and their values:

- address: Address from which this log originated.
- blockHash: The hash of the block where this log was in. null when it's a pending log.
- blockNumber: The block number where this log was in. null when it's a pending log.
- data: DATA. Contains the non-indexed arguments of the log.
- logIndex: A hexadecimal of the log index position in the block. null when it is a pending log.
- removed: true when the log was removed, due to a chain reorganization. false if it's a valid log.
- topics: Array of DATA. An array of 0 to 4 32-bytes DATA of indexed log arguments. In Solidity the first topic is the hash of the signature of the event (for example, Deposit(address,bytes32,uint256)), except when you declared the event with the anonymous specifier.
- transactionHash: A hash of the transactions from which this log was created. null when it's a pending log.
- transactionIndex: A hexadecimal of the transactions index position from which this log was created. null when it's a pending log.

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_getFilterLogs",
  "params": [
   "0x68ce60ffdb0c9480c307b0c3d2ae9391"
  ]
}' | jq
```

Response:

```shell
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}

```

#### eth_newBlockFilter

The `newBlockFilter` method creates a filter in the node, to notify when a new block arrives. To check if the state has changed, call `eth_getFilterChanges`.

Parameters:

None

Returns:

result: string

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_newBlockFilter"
}' | jq
```

Response:

```shell
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}

```

#### eth_newFilter

The `newFilter` method creates a filter object based on the given filter options, to notify when the state changes (logs). To check if the state has changed, call `eth_getFilterChanges`.

Parameters:

- crit: ojbect FilterCriteria, a filter object with the following keys and their values:

- address: optional, a contract address or a list of addresses from which logs should originate.
- fromBlock: optional, default is latest, a hexadecimal block number, or one of the string tags latest, earliest, pending, safe, or finalized. See the default block parameter.
- toBlock: optional, default is latest, a hexadecimal block number, or one of the string tags latest, earliest, pending, safe, or finalized. See the default block parameter.
- topics: aoptional, an array of 32 bytes DATA topics. Topics are order-dependent.

Returns:

result: string

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_newFilter",
  "params": [
    {
      "fromBlock": "0x2bb7231",
      "toBlock": "0x2bb7233"
    }
  ]
}' | jq
```

Response:

```shell
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}

```

#### eth_newPendingTransactionFilter

The `newPendingTransactionFilter` method creates a filter in the node, to notify when new pending transactions arrive. To check if the state has changed, call `eth_getFilterChanges`.

Parameters:

None

Returns:

result: string

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_newPendingTransactionFilter"
}' | jq
```

Response:

```shell
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}

```

#### eth_uninstallFilter

The `uninstallFilter` method uninstalls a filter with given ID. This method should always be called when watching is no longer needed. Additionally, filters time out when they aren't requested with `eth_getFilterChanges` for a period of time.

Parameters:

- id: string, required, a string denoting the ID of the filter to be uninstalled.

Returns:

result: bool, true if the filter was successfully uninstalled, otherwise false

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1001,
  "method": "eth_uninstallFilter",
  "params": [
    "0x43f0c93bf463861b7c15a5d11d402d9b"
  ]
}' | jq
```

Response:

```shell
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}
```

---

# Using RPC Endpoint: https://rpc.ankr.com/xdc

Executing eth_accounts...

{
"jsonrpc": "2.0",
"id": 1001,
"result": [
"0x724da43371727bf1d7eed3741217ee03346de9ea",
"0x1b1552127184fa0963ad910538c293d83ddfdd32",
"0xb2f5f4b60e477486497e7ddab913fd24c1e81b4b",
"0x9c3efc01f7ff6681b0bcb18ce5cd9f9ab0b2c2fa",
"0x24da29f981944fce2a95da5d82eca943ad68f31d",
"0x90bb08c3d474ba1438647ef936a343133184de45",
"0x27dcfcacd95faefd3746d617884135277359fdce",
"0x9f4b325593c3d8ab453d4ab96754b501f5ea1052",
"0x0fb10bad35543c6e472af5d12499f679151e9457",
"0xc05bb5b4f73df4626087152016b832b4d96bb14e",
"0x13b90b7f16821dedbf2658be6024b1df3afa51e3",
"0x32323ccabb3898760b21d7c4fbec34f4226d283d",
"0x26f7afdf81afea4e587f6edf876b1c2fbb8ac802",
"0x39fa212f8e19083f6dde1187e0357b2cd8138808",
"0xb533167c0ee983ea166bc11795dfd79ed31c7d0a",
"0x87db18fd32cd84166df77910e532296c59f03868",
"0xab6492cfe5dd604420e72f17655bb086b23c4a6a",
"0x2640e04fb446b8f84313cd9fa0c23f2d76bbea48",
"0xde692ca132e2898eb91a759a349ef21295f96349",
"0xc85ece9bac746166e8bada7f307e1f7c5d1035a4",
"0x088f5051fdaebb895b84dd88b706bb3c217160de",
"0x0ce287b0de17e78a870f0a7498604b72b4fc6856",
"0x71634b300a38c002d948468a8bc3217adb74170e",
"0xdaca1b4995e07c89e33ee49ab16c6904675021dd",
"0xfac5307dd07b72bcbac131ded24c54b16b93335d",
"0xbeecce90688140b2bf15cef7be741e0d676fa5b8",
"0x86ca29948e01c263ff95104831a714f78336c0b0",
"0x3cd4100d51b34ce94ad39af4667698e34e0bab4e",
"0x83fdcda7d2fb846188a1342e6801782b8ae87348",
"0x94578a99b383e1564f85c2dc593d53835058174e",
"0xad55a0a9b1feae1613b53cd3e01c085d199b762a",
"0xfafda28d9e8b422d63f0a25911d0102df0a69ef1",
"0x8998ad1090322052f68788024f3513ba9ce9dbe5",
"0xe454b141657a3d18dbcd5e1a93790e20c1b72368",
"0x6c32ba44e927b2d21828bc27e02dad5f57fe8187",
"0x85f850c54956ef081e37cd0e0703cf7cbd066962",
"0x7b11167e2314561b3a1b7c9e7cb8a669c7104425",
"0x96fdefc8f75c447b649003f4174ce210db607b9c",
"0xc5a55126cd823d02007e245f5331348f56051354",
"0x5bdcf254229ef75c70027b05081b60f36509bae7",
"0x3fabaa5d858410c9fb471547755982cf2b10e43c",
"0xf604cfd8b8a1594e37a7cc15b1cb7b746cc8eac2",
"0x40f27b0a774d2053dc87d910f78441121d53f6bb",
"0xe3aa031aab2a9490b9813096833c0f1813ed0258",
"0x70eb3026c182858eaedcd7222cbbccd2782d5059",
"0x946643bc588ec9cd4843d5345d28348367e6832c",
"0x0e881629024fe5eee54bf5e808919458f71964ff",
"0x2acc9d95ae9b1659c17839665c3e324416ad91e9",
"0xdb95ecd043af59a2fb59cc5e442cade45895c16d",
"0xc6341654d7049dc864591cefa5da1b24d2487d41",
"0x81fededdf3ee449fc2b0f4376d3ab9e82c809f92",
"0x80f5669925dcdbd2c36bff4812db69e4b95be143",
"0x5c586fbe12fa4628c190c419d8e7cc931b19a65c",
"0x3d0ee4c88b6b123eea614dcfa77e4cf70f2166a7",
"0x83b09e4923508ba5e264d0775c1739d98fd570c9",
"0x0d5a1390c46e0b1be30d6f5c71810d810ca73b29",
"0x3008bdfc5c093b1522270324101fcd383477e9af",
"0x814785a8cbe49921832107488262edb663147e90",
"0x203237fd77955515abcb5c9a4c685c73938b979c",
"0xdf509beec43cbb9b3bd9936fedcba5388e937a29",
"0x76f144034fc41798b1eb87579096e9b157bd9534",
"0x122ced1daca5f82131009e522dd56eff93a70b36",
"0x23c63053483a34ad7a317f6f9af843dfada899dc",
"0xc8109fa721a4a8020e3e3617586f5b0d396d7641",
"0x313511481bb912da9c0d0a96fdac41fdfc8ce24e",
"0x310266c6bcb2eb6f0401c533f0e766913263bb0f",
"0xb52f5985d65c72ee37f1e551d0c965b93f687172",
"0xacb3ed55f98164325503134363e66219b54363c5",
"0x4f1b0b92a053e242a8a6e86f60b3c62359821662",
"0xfa19585a0b4e29cade79560b45c0bf2043e42453",
"0xdbd4279be7a36f2fb9f7a21f3d7adc68d89081e6",
"0x9a635abb589ba8998b16fa700f853394db345c44",
"0x44ec81a9ec9c8b3fc36b60d7241ee5585e516886",
"0xd5f7be04fdbb0448efcb3971b2ea830ad9b726ec",
"0x9b6a9a2cc809309b9a6db1f1a639f2860ace9bab",
"0xe40ffa7e68d7ffdb2f49cc9df5330cf153b9624b",
"0xf9ca0946edbc4eefd8d2d9b7133f303e368d95bf",
"0xa158d1b2feb8f3f15a2e3c8612b49be45418ad8d",
"0xe54f31aa6c6e11bec53e0007b683c7345ae5cf93",
"0x8ece756bbbcf56139758af2dbc944bb1f1903a51",
"0x6dcf2f2d0bb912d38d4bcab2d163f9671c5cb666",
"0x91a824dfc5ff2eb34dcb396c36f31ea03f329f00",
"0x38e73cd75e25a48d69f75bbdf76463f553b29fd2",
"0xd235ab10105c316eb3b920646d3a6ca9bf22c8de",
"0x856b7d152cf51efb424eef245949f0804a40dbf4",
"0x18b387c16641047b06c22cc2696f53f28e43ab8f",
"0xfcc86105f634ffec5d3653f6e45b467280594a5f",
"0x2adccbb93e6a9ac7e9467db1b9098b2816abf4a8",
"0x87c5ecb512769a6113e0b29d05c563634ec184c7",
"0xa85f68b66c5c20a47e7bbbe1d5e7c68e14e5817e",
"0x698b1dd89af48bd0f09159e2edab9ede167471e3",
"0x50d0cbf116f0ab6c810bdb695c5d6336a1e8b5c9",
"0x6eb57e3288fc7de842b2da77ad899f6cb30ed079",
"0x648ff38a6efb4341738c6d990e74d5c8f0e9a9cb",
"0x5e1d2f47704f47f80229b9379d593a173a2dccd4",
"0x4adbdf91b776a67574f913d50eedbec7c807a69c",
"0x2c4cc4f79670ae3007376ea8713babbedf6439b5",
"0x139710365dde15e95abe7051017f5b527f5ed38a",
"0x8ad6b266e10d7f8cd5dbc371d390d3f35f0a9493",
"0xfaf09fe37b33df47874a57f6009d0694d195ab82",
"0xa43da23aa1f73d01007519c3afff0a8ee2f14411",
"0x6ff474634ef206187e138e580339ae392d4f0340",
"0x5df96f606a5a799dadb0c8c44d190c43b24da0f9",
"0x582bc5ab4d0a47b893625a0d5c8fe32f82a475cb",
"0xba927181a191307e1ca505eb6f3e985b049abddd",
"0x10eb4ba3a97357084aa5097de2eb7256f585aacd",
"0xa598cde94cc24c390d497e98534695cfb65e6b5b",
"0xba2cfdf6cbcb90a2ebb7c51a4e95b948cc1e7f0c",
"0x9f3d7f5776e1dbb7db19f1c1124cfc2d415a1091",
"0xee826e63f11e4099253617de82f77aa63e257d2f",
"0x360560c2fca25160c6bb770f6f7fd7e33c8f5e14",
"0xa5828825243819a46cd3d594537b561b14624e2c",
"0xd4245690c6ebf1cab11ff6eb3ae01f0cbc8a6059",
"0x1f776ce980ce81fba9f19b92d3b988e44577ed32",
"0xc6cc3a499a5b9b527ae48f521faf7b2036156e10",
"0x8a53ce1c4a23e9e7123be4979155fbb2cfafe837",
"0xfe8129153c577afbc19f7f692b285142b7769f49",
"0xeb916a921517d279a934969eb26517cce1cd28f3",
"0x98905d90fa6fde50d92688799b3488420e77163d",
"0x930b8234b73903c42d454bce23ba7937399ef6ba",
"0xdd621fd15bbda32b30cb81c42980fffd9a51d9f8",
"0xfc5ed957a682886b6c2ebdbae82ca8bb1bed9c1d",
"0x2971454ead7f9dd73fd07c7958ce535ef69d90e0",
"0xe1c3713007d9fe45544d237517c7666eab5e62da",
"0x446cbc80911a705bbcb0ed98e2233178fd218958",
"0x7ad89af734b2e822fc683a517e605048d862247e",
"0x6259af8f40450d5e1dd9617d8681e64cd9258569",
"0xee48d393af03ec09537d20278cfa904641a14627",
"0x3814b89b1219a6ca709e5cf6e5e4f14c8e738af9",
"0xc8af65573d4238baed7cf983c682ce0e055b133c",
"0x91e88b8c6a22c97e6bb028e9a8e317602473732a",
"0x59397ed01cd0dc43d27834a9fb4b1bd02d406d17",
"0x4407e1ee49d26af0d53432ecbef955ad4a17bc12",
"0x943a0174d66868d041e2a870bcd661255d0949ab",
"0x13b21510b050a8405ad33ede724e28e58c211b64",
"0x867f4ba06b61a67679c6d2d50a7c5c339fbbfefe",
"0xb4bd3a4a5cd64c980bb9fd9671637f9e3f82ff14",
"0x636f73985da0e3424a069b0cf2623426cf81ba08",
"0xfc4acc547e7eda6868fc68667d5f51f38e794f08",
"0x431b7bb06a889ddad8f45b57b45b9cd68b3bcc45",
"0x43a90a32be48860c432f9ac3fdab39072e92f485",
"0x01226bf1d9be7e5479088fc005c1107aeeca858d",
"0xa1ea82b1c9f92c158f6420d7a0db01e8629e639c",
"0x00af65bc042c073d758bca4c3cbab64db58d9fd5",
"0x6e04ce0e366cdec84677ff4fa2b06461eb3e8c4b",
"0x5e2f9c9b94f3f7f9026b60944aa288789db5faed",
"0x6ad684cda672fafc40003921bbdd8903f4d39ae6",
"0x5afd45da249c191683634b2b5d830ef42eb78206",
"0x802e63811ed8d34346548d3cac9fc273b2ab547d",
"0x944a383c151f29ab710c38cac90399f25f04159e",
"0x6b41c4bbf80e7fecd51aa4a69d8091e083dd0e52",
"0xe09e54d0ca0877400b0c4258cb0e231762775b04",
"0xb939acddd396c320d20a9867802aeb79e17f56b5",
"0xc8cc1fc1c85db24f11fc8bba0f5aae8892dc37eb",
"0x74d60e479e6408e0129400ef38d173b3ce1c7a37",
"0x56a8a9454aebdc9e637c56e2380aedaee52b748d",
"0x6aa4a24cf2c78ac52197796b682477de837e5aed",
"0x983fc283bd3761deb6380ff04aecaac7e2dac22e",
"0xfa94a9f462680df64319cef2ba3a3c8ad422b99e",
"0xe6187d5da69204e4d93c2cf429a510bd6754e0da",
"0xd09a0a5533ad8d283638c6bd9228eed5ec41cef2",
"0xa0b35624196403cb84382c1ce4e8b3d1035167b2",
"0xb610fbc269568d31af5300706efdb7886f0d3f3a",
"0xe70977c669d9b604a5440e6dcc33c32e6661e886",
"0x9485e93b85bf9c8e289d550522fbf2c01574d668",
"0x67e92deefd29b941fc39a681cffd0b19fc395fd7",
"0x1ebbc838fdab8628ed4b69b46af8856df38c3442",
"0x913ffdde90a483a324397a63311b078e89a66f2d",
"0x41aa58da5b449bf301fcc93b533efda133209174",
"0xc4249db2440056ff2121d84724bce54fffd2fe57",
"0x5eb05322b5c7a115ec3987b416c6bc823bba4232",
"0x1125d383ba08fd8a6a6446a40187cea72dfdaae0",
"0x5a05468c3c299b5676e7d9aedad694b9a6fd98c9",
"0x451fce3a8d695ab8b97297b950a4198160d6d7a1",
"0xa5ada214b60ea311cf195c4ad73a87554f738c33",
"0x92a78e82fb716376541cd4dde1bc8906e9b30a9c",
"0x860db4ea830e9c2eeb2b37d4af6d98cd39bfb82d",
"0xb19cc8bc4c575638f0dfb1cfd009d4b89ad355e6",
"0x7f6a9f13e8d066b368338c46b1f8a18e7a9a395a",
"0x371c6d563062e1a74bcea5a00223215821b7fe4e",
"0x640dfa5cf45b00ac4448e901e990511340ab1158",
"0xddcc2a3f2121f7cc688cbab05ba779b8b658f249",
"0x258953e41ec7501f90e025ef2636e434a352388c",
"0x7d71edc92c7bd533eb1538fb627e9360e9c58164",
"0xabec9c2548a949b76180e59f12db17601f166258",
"0x03768b3c78afdec1470c1e0dfdb4f8711ed27a7c",
"0x9b32c5cd4474561482d5fd63bde49dadee2f8b87",
"0x58266fdf59f8ce3008b7c9f52666b6a1b23bf3ce",
"0xe64e376bb0d02197bc3924db0052a605f5092f15",
"0x74c615a37b5c79ebadf82465647e38ac209c84bc",
"0x91f20bd0e40efd1f1c31b711287e52a9cf634664",
"0x49c4e317b8262f87012d6a6614c98066cf8b097a",
"0x2789053be7dafc4dfe39dd0dc9d184c2be4d0a7b",
"0xa707f53eb5064d0b3ae4f20f89580e2d7786ff62",
"0x9fd4d20f89926fa3fa788b71a042ba14f8f68f41",
"0x6429b0dd53c89cd629314f8a0ba54bd842125803",
"0x1e8757451058388ed4c4c5c3aae4db3abf26a883",
"0xdd4cf86fec603af280b9de7251fcbce40f3e9a18",
"0x064379acfb0f8e39b5351302c0be24b094a2a800",
"0x03af06240fd36ae257ee693c5d7a445a1a9d9950",
"0x0f505c8e4b429e07dda1904846609ddf6ec0d4a4",
"0xcdcd4e77d4881cf66f60482b830dfb7b3f221bc8",
"0x0d05915936b479ca68a69ec6181446da7a2c12a2",
"0xb9bd179a8f29d6c8af059b520a54e1af898a95dc",
"0x3a0cd31b21d81765cd9e43f79b140da2e0fd8bbc",
"0x0102fd479deae48b7dbbec5af9cc69638fb4fb45",
"0xcaf5409d437d9127e168e440b723723ede6fe9c7",
"0x7b27fd319ed471814358db12cf2db0c29e366fd1",
"0x1bf823caa3be4a61b0dec06fa306d09f0733514a",
"0xe3a244907f4cf7f31a788e155419b82399ff0f48",
"0x8feb30d97fe021148176fbf35480e0c87134dd3f",
"0x9b7297fe2d5d7de908767710a9113238581135aa",
"0x6e10e15b31222064dc38c9a7ee2d41c51abc5cc6",
"0x12e830d13827f408677141ae6ecf4c4b58997d2c",
"0xc6e7a10ba32fa7ba06a7c6c4b23282d8924184cf",
"0x428ba7db08d0e999a3af5a356709554e1958280e",
"0x33118756f582a4f751ccd4dfd5f5dbd7f5932818",
"0xa8ceab7cc7f105efddfed06ec6be80bf5d739c67",
"0x65e119944efc6fba564c1e579ac0e678d581a9a2",
"0xf92eb5ef8c1993e46bc6e1d3c9ad4605f9ed7b34",
"0x2af593e5f17f2d9287a8fe73432585479198e0ec",
"0xac7dd2a4e354705487c411a6d67fd00eaccb56bc",
"0xf7c8589e33b153e9f7f4a8e0a85ba1d7bdf885e8",
"0xd7d2d00180099478c8c970d96e1792dca19ac8e9",
"0x17265eef6c0c11b602e5546dd3eaae5195ab3fdc",
"0xe8f20a20eb94f7d6ee9f3a9b8f8221a334ee5bcc",
"0x0b8ca6c68e50c74cf85329d727a798aec57717ea",
"0x330dd840c76ded3447dd772581def66815c6d90e",
"0xa3b39b84788ccd6ab6d78bf7bd7cb6f86f76d1a1",
"0x580bb1ebad582008c793f2f09046bbe9e4713896",
"0xe9e8df9444894ec6bf15aead4122446891a89930",
"0xdbb3570c341fb6634609b8c742941c91163d16f8",
"0x6880262527b3255a77d40ca7979689dcfcc7ca42",
"0x09cf60061cd0ba29b38e4ccefd63d4433d2b4129",
"0xebb9bb31ad7500b583f39ae9f763c1e4af50bd0a",
"0xc2de56851f43958f07175fb9c7c3e7040eaccf81",
"0x201fc7d3609f698b592978defdefadbf2cb1eb3d",
"0x731b6616e656c4f4937d77f083caa55ab50b6e5e",
"0xcac9a010abbc0cd8d21d1c8cf3e86fe9fe54f477",
"0x1d8626e04c6bda93425612199b73c3c3cdeba7da",
"0x942174bef0a6b024cf10bfdbc9fe52fcb58414ef",
"0x01d337064074ba56dc6d9541448284814c59d523",
"0xdc341721dff7998cbe8008396c798062f448d68d",
"0xe08dc479683b83ef37182497cbdc2be6c72f653c",
"0x156c6183dbce64d60a0299ba5e23ffc5d23845e3",
"0x084554c2d9cdd87a4c2c2c32d2cfbc2ebd1c0917",
"0x2579048ec530e043dc60fd5662c47640c0e15126",
"0xc6b0326eb375b00d1adcaa1882d83d1be8b5837e",
"0xcf81ee8cfa4aca6e304c9cade7a301ca9c2cf146",
"0x3da61f49ef5c3e6a5523209f70d6fc6280326eb1",
"0x762ffcd0c7eacb3c5bd747742ccd185b06625240",
"0xe0a89c63e382de82cfdcab75c666b3d92ca9188c",
"0xabe19f1bac4ac69b1782c87d21de5a920bbab249",
"0xc13c12b31e9777f8c4d6fec7b6f6f3feff925b7f",
"0x051776a593601fcc09f94000ab61e12a9a95f95a",
"0x70f09c9f65ef67ee4b1587ad2845e1881a126d4b",
"0xf8c9cf0ff6b5c052ea99904e44dfa4457eadc0b5",
"0xbb1915c4289ec59dd64150f6a220ecec3b0ab552",
"0xc1be8d415d0ea9e389ea5cefec4d5dd15c041540",
"0xccb7609bba8801bc4d00947262b31c4f0478662b",
"0x69291ddba25108e49c4cf9d9d19d3fb62320625e",
"0xe559e779caaf709f7307319bf9797913caf0bc87",
"0x54dda6b71eb4af6191daa05f88dca4489c5e574f",
"0x527bc0ff54840d2c244142a7f7eeac8a58d41645",
"0x90ed88c5341f37c326bd43426371b6024c2b165c",
"0xbb7856e8721121ad712974c08842ea950a104b00",
"0xf1c82ec177e3f99db3f39097da41f327ed395be3",
"0x979e0f566a955b4934199e249caf9ed89452c9f8",
"0x154612a77bae8120d6d1366922a1f93b02912a4f",
"0x185508c994ee5d76de0a061cd4142d2b743886aa",
"0x6eb7a645a5a429c6787696a0e26e35fcf12209ae",
"0x2739b8c4c67afc6fc2834320e20c4661025201d7",
"0xb5391f0f5adf22dc8bf4f5b8a4d516d0180d978a",
"0x674900dcefce6de57f08c2355fe368703b90c8ee",
"0x619d5edbb036bb3cfb99540d81275579f709a198",
"0x6a8d3be6e882033f7ac7b55b0540135393454492",
"0x4fd5e8bb9882fc5c7f5dd966e72f49f0e1cf8b8c",
"0x8752aeb8435b11f2847fa71da2ef6c69c2c7f4a4",
"0xf3f46b98f196d315fc649f680c9ff16c6cb2e1e1",
"0xb2335053f16957ad4b347ebe0b7644c5c3b3ddf0",
"0x355fd8675806ab126b2d3b30c36851af42d5832b",
"0x8eebfd6609d9f45eb6018d0286429869913c705d",
"0x920f2612a1a2b8db99000d33ed63c54ba7977439",
"0x29e074043ace527d8f67675a656ed2be381dc16e",
"0x8644a13a9b3978522d10e93bf0e2e31699e289b4",
"0x7cbeea2adbca4b47b43138ab2f65b997668f10a6",
"0x0516c69ee11ac84b9668bdd785c6c2d69cdafb75",
"0xb91d4ed4ebb1662c99570d3d65bdc8d4b8d07d8e",
"0x785bf3175cb90c6ea39c3c2e6289ca8b4c9d62ec",
"0x64c87b400fdaaa66b9193fcdb80bfb9a57384c79",
"0x7461bd146fa2473cb7a86952335bccb78580c12e",
"0x30d28e4051b328bdc4ed05cb4d6fe43d276e3e66",
"0x0737405fbd6aa4c93de56ba4c9d6cf1082e55279",
"0xfa72dbd6c24830da6425d58f134d0f873a6b590b",
"0x02953305cd8fce44ac882d116d84364ca2cc2020",
"0x8cad2839988aafd61c4dcd5c4c10ba21a7536354",
"0x2dc7273e4e50fe95cc89dc1e9cb1fe60e5640dc1",
"0x7cb3d6d6779641f8ca7d3d4511fced923430fb7f",
"0x45bc79ec8778a1ee402f44960a4ec58fbfa8e921",
"0x6af5d0f3fd6ab74068a0cdcde3ea3f6026f88396",
"0xb28203424e1921b6f3006ad5175da1a377f63159",
"0xfa94b52a9ff0189d5c322e70057abd9a6ff2b9de",
"0x0889f3f75b3046939e7626f03e99f39b134841db",
"0x8b49e0ec57480fa31dbc1c4cdc4322f119a14314",
"0x5b4af61d126f703cf46912336ead53db6e4b8b50",
"0xfa1e1eac2a035c5e91f93eff3681c2bfce7f4104",
"0xe739d97e9b8f168d18686f55adb0927788878345",
"0xa29f5da5688e1d479bf66af4c69799bbd1ac6dd6",
"0x301fd241c626463307058dfb8af54806a1f95141",
"0xb4d1951829920cfb01cbf36282dd4b16135524bf",
"0xb4e24678089ab9acff44c15b190fadfe2bc31f3b",
"0xe89c454aac720088004aedfcb860e2bfdd299ff3",
"0xce8dc2a35f0ccdd8d26fc9eacf5b696c373b3ac4",
"0x8add7985189784e6445d20c3e20c93cc3480ae38",
"0xd15d3fc1c006754644fae348ee77b19660bdedf3",
"0x30b61b77e54524cbd1d2402df23175b5fabd4f4a",
"0x965ba376d57a264c551e9b0904ff948640e9ecd7",
"0xcf0cb0ed3a7571160d66facd0ab0b8e48ff66fd8",
"0x8c94d85be1e4baf9aa4d5e9f30790ac58b528c9f",
"0x9b1d3c5bf854f291722b6c09779cf294753a5ca1",
"0x948e10a6e14313d4b7d10c4b39e018aa24c715dc",
"0x52726ff64f3c471cdaab538e9e907766bb756dea",
"0x9b3a5cefb4f1fa68b04dadaf9c3720cfd172e6de",
"0x7e601189d394995f473df079b581297b719176b4",
"0x75a8ba162bbe6ba0d795650147b84f770099e987",
"0xc14ebf355fbc59017a9469224a341cf6c8a062d3",
"0x6b1b7e2a5cb5850cc18eed3528339707881746d1",
"0xeb00a32b5f7115b334180c488a480aae70fa678e",
"0x5501abc44760b7f2dc44756e5fe30fd3754b35e0",
"0x2d93b90933f0498daeda86c8c79d1c3b94609672",
"0xd84310751005057a12bba270ed4d88fe0b760f5c",
"0xd3a19d9bf34b118a72c7a17836726141d670c624",
"0xce054d95dd029aa9e5612cf22939f068d4f757fb",
"0x03a99bdd62879ca8685477b9bb470d2295622769",
"0xdd330a10f6ca16ce2e334e4a6c36bb1d3e8c9514",
"0xa8066df6d8ac700f4f7a282d7c3525ad9f3e501f",
"0x4701099bb9089bec8ac90e0732f4e8d912ab6dfa",
"0x52c9fe06a1d1a6869800a415eb6cdec1bfc1331a",
"0x446e282a288f1845f4ecb2fc29f77f986f70e8fe",
"0x542facfa0d1c836b9a1b0c862bf885db969ac0f6",
"0xa388612847d1c43af6e1f34f6b5a37a34bf06af9",
"0xe0392f9b0a6c4229bbf4e4243ced5b61ccd6d60f",
"0xa197f6de874a16b515c69f44fbf8d9e46678909b",
"0x194b29e85a5b464223744c494692c9966b253cee",
"0x5738f18da697590eee13d3762674d11d71299e3d",
"0x2e0e175673a99450cba2580df35783c6460bc0dc",
"0x82c11e76a1f9d2d102acf9df89a06b60f723742e",
"0x17ec7a931cfa4b2246b554d940d7b010db14e591",
"0x9eea84027f68b6a0b01d4c015ed0fe9fffc0b88e",
"0x92861681102120d9c1815b17fcc0f9ec1b7c0ea7",
"0xd6f777a2dd0ad89688f4aa8e181dd1eb2417a20f",
"0x2d918ed3a1372d64f79a1cbae661ec8882100801",
"0xd1ba5511be1c21c882f549a0635943ddf76eddb7",
"0x4d74eea220d2c2829fb65823fee87963ae0604b3",
"0xe7bd0ab4ff45faadd63e8210888e8c23d88a563a",
"0x9cc255d04bfe1af73e4b7fdfe91ec3228e41a06f",
"0x975de6b32fb01cb800e7e7eeed51ead04b7031ad",
"0x3c49829a61454f340168d99ef58c9f7c271fe20d",
"0xd20e70821d465632c5c4d2c91a55952dc12a8662",
"0x9f4d00a580c7846da6feeffe1d76b37d379f4b04",
"0x7fc8325c3f1705c6648dafeaf100fe750ef3ae5a",
"0x6c201167d8f9cc5759163ebdc7059c3a51e4d164",
"0x7ac67cad12b9469b994119e02ce734c84cfe9477",
"0xcb9b4e8eece82b9df55d2bf366db5fd35546b444",
"0xd9f678dcb18b12aeeed92badb901290e107a830f",
"0xe687cb908c0e32faa1f46da56f87205aff50c4cd",
"0xda9f402691eadacae2d6b314b11155616ea25a44",
"0x8e1c30995159047cbdc9f9de78890bf060f97c18",
"0x2b66e96848fe524f3f03a1a6ff30892f3dec3587",
"0xd037628297f8c56b179a4532514e312cf3bf22cb",
"0x7d494aac55f14c1c0052a2f227ec3a9eb0eeffb2",
"0x0a02435be5fca6dbc56dd3863de0dd58d834d851",
"0xfdc9dedfdb86e119a80ef812182557fa59e29486",
"0xa68dd9bc8fcb2602774c987e49b3fea7653e745e",
"0x4ad6483e248393096ae6afcca2c998254ac48f35",
"0xb4d4a7bbc9bb53a871b933f40ee1621fc2f05304",
"0x8a8317704b8bbad8a677780731ab81bb793f3087",
"0x206271c40348ec6c7f40349b2184a553b4ee9f07",
"0x60b453f9659d9fe14c52f38a06d525f3f9247574",
"0x0c17d9bed88de881ecdb6878e78dcd8a9b4321ed",
"0x5624636a4f275bb2652b84325d6bdd7d0c9f3bfe"
]
}
eth_accounts executed successfully.

---

Executing eth_blobBaseFee...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32601,
"message": "the method eth_blobBaseFee does not exist/is not available",
"data": {
"trace_id": "5cdd0e1ca22f47d1fb98bf119fa0e51e"
}
}
}
eth_blobBaseFee executed successfully.

---

Executing eth_blockNumber...
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0x54de0d7"
}
eth_blockNumber executed successfully.

---

Executing eth_call...
{
"id": 8001,
"jsonrpc": "2.0",
"result": "0x00000000000000000000000000000000000000000000000000000000000000d4"
}
eth_call executed successfully.

---

Executing eth_chainId...
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0x32"
}
eth_chainId executed successfully.

---

Executing eth_coinbase...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled, reason: restricted by blockchain schema"
}
}
eth_coinbase executed successfully.

---

Executing eth_createAccessList...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32000,
"message": "execution reverted",
"data": {
"trace_id": "972b530211eef414ade519788bfd82d8"
}
}
}
eth_createAccessList executed successfully.

---

Executing eth_etherbase (using eth_coinbase)...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled, reason: restricted by blockchain schema"
}
}
eth_etherbase (using eth_coinbase) executed successfully.

---

Executing eth_estimateGas...
{
"id": 1004,
"jsonrpc": "2.0",
"result": "0x5208"
}
eth_estimateGas executed successfully.

---

Executing eth_feeHistory...
{
"id": 1004,
"jsonrpc": "2.0",
"error": {
"code": -32601,
"message": "the method eth_feeHistory does not exist/is not available",
"data": {
"trace_id": "f79e4c898ddae00c634d232c364af399"
}
}
}
eth_feeHistory executed successfully.

---

Executing eth_gasPrice...
{
"id": 1003,
"jsonrpc": "2.0",
"result": "0x3bd8c7240"
}
eth_gasPrice executed successfully.

---

Executing eth_getBalance...
{
"id": 1003,
"jsonrpc": "2.0",
"result": "0x38d972c8091d1125"
}
eth_getBalance executed successfully.

---

Executing eth_getBlockByHash...
{
"jsonrpc": "2.0",
"id": 1003,
"result": {
"difficulty": "0x1",
"extraData": "0x02f9135583805187f9134eeaa023d9b15b34339bed8c20f4cea2aea76194a0e670c1b412dec9a77991ee418c5083805186840548f4f0f9131bb8418a9a5a8fcd87c3cad67882c109a1c0943564c5905a01b955ecd1b1bfbf1f507a2280b5bb8360bd67cd2b5d5d3f98247a9e384d9cc2097056689522c6c71b4f0301b841522ece4249e8ee4b3082cace851d50c50114553985325c10cf29b9527af41fe52c5571631e5f2011657747c78809df30abd551e13bb6fe6278f73d2b8b75512b00b8410801db1ac986de42f36e7fb6d8fb538a62211aeed5d0c1a6f6b597178e53950b17cde2d3808b08f1ff36ad10e20c5108e6e817efc1100feb1e8101f2fc32a41101b84140d113edad78aa56fd110c4d4a1a39427a5210b05a2c0e340863bb22627b70fc1ad089e2f008d11dc62044fee04351fb91c893c427a41f6aa5344c91c8b5d4dd01b8414b2710459d77d9d842bde95c39aa03a642cbd120f784a51eb7323d88a93e98851517a457a564c4fcefadc3307adbd482696842552a0297fe835487931128185001b841aecce9a88402589e89607e3b5bb3fc21fbd5fde146238058e78a7b2115ab821329cd08cba450c76e5ab78cd05de9ca5023a7446970b8b10e5602e7ce7eb5bbec01b841e7f943508c97d6db6312f57f31a58414e5a750e552e0c33416d985ab20b53a541ab8cd858be92919792cbca10b630f3112bab25939f9f12f972155d3d6e44d7700b841a08d1ffdc33c6a8470dbe2e3c2e474992e8b7278e6259cc99d39d4ac49626bc44703588a1e2a80c4d871b38c1903774dd056e6abb68dcf5e3eeec310572d4c5501b8417d328a967965ca5f4130ad172a684476fa8788deee735384a39bcaa6e942a49043fab8eb4b356612cb7c8898762c50088b14721123bdb21b971b0b7f059ba57a01b841ba357d7bdc48591a7b60aa000896cbb08de74b10a0a25b5231686506ee0edcc2015e1fc80a07ca3a01152f977ca2d90503fac64ffe74e5cafa5da11d06c6642b01b84108bf27fb082dfa8e83147a4598d2253aff92fe0099ca234048c01d94d10d43201154b5a57c9ebd47aae6949fefb753f8fda604f1bee3fd7bbbb1f57b4fae888f00b8412016cb51e4b34e96ba1957bf835bb126b05426b051f5e183758d87551f6cc3015cfb96a1e7f9cbbbcbd3aeaa28fa899873cb6cba504ce56586ccea013cef216501b8411c0f16e4f47a192a47d39aa12bb7313dd2991fcf03b28a8da50f59dc4cc54b2351e66930ca1011d25e3ef61e9db909db0a1d4fe1db35ddef4ef0590b4f2ca9c101b841eac21f66c6d5eff307f46a4458748b7c464e83a11e3df6c8652988bd8b9826ac74884615fa49e6ea040eea7d21ad68ecdd8445ac8c1787d897a8cee9f9fff25900b8414db0c552793b5adc7001c61089c1daf45814734548d720f3166ef98f405a47895b36a138da243c64a3cfa5cff3ab8a37a2bc60ea8138b5f0e982fe594ad58ab400b8414f6e62f037bc81dd26ac657c8f449e751b79ec7cf2c596a87f852220b7488cad66600ec7fcc9334bccee022f350a43cd14e399413ee19beaf83a5e031436255200b84120cd13947a1c45ff039803ce157928afafda1784ac5a5ae592890f07b8a2b3b5105b403dc24b7ebfe39a24aa082c420a77e937fb888facd4c53fb5c029a25f6e01b841a8edecb818bc380b630e5f01c4f1ef0809a3836e83829cb9bd234ba373caf5cf72b1618d91c56c56c1d491bc50104a1330f57148c8cdd1233ca2a974a969130700b841fe9e4134a9789e0dd23095db384e42b12036afb0396062470907da162216ec2a030fd66978182381a51868bfa983fa1cb6937948576671768d2f4d6b2aefb33d01b841dbef788acf979f4248e12c2b3bdd79cda356ad63bb9a55953b5053e4a41f2c810974ca465837238ea489efdef90791baa8e64dfc336bd88148f82d7b684006ae00b8416e792fb0c29c086c28274028a35714d7b876ec07f203be1a9628bdeabc149d4e32a7de951132f4b7e2c96369fd3e4cf5448312c36275b614bb6d0ea20ce7c48d01b84122318490dfd07d7d1725238a547d61c11b76c75b3215b82153bd0c8f328d96e17f493627005c418b0babcbe5e985af7cfbee6388db813fd6713f13a84a055fa901b841ea3c7af315514adfacf0f4d578eaa9cb52362f974a061b231ed7f8c604bbd3177591b9e1208115abc3267009d324f66a4cbd71ab4ce73f65387846adfa5221f201b841c9bbf9b407b98994ba0c7e8c2ab650b90d5945e70194fcd5bb14af64555d18bf5fa2258f0b1d9df02af14397399e7214368983f04a076a9e5e3549102500cf9c01b8411f456c26c51eda415ab305f5b03b96f53b38712e7aed46d424c49cd55ca6f03610468d402f390e06390b7aeeb7014a4fd8ce9372c70d05a7df49481058bbad1a00b84197b4b75fd92acae857b66ebcc3ed0e53a887dee8aec97e73ac01395d90fbdb5e2f9c7708010e89b04168c1bbdeec816bdbfdfaf51395ca37d6d6de6e21401cb401b841c33d387fe750815323ae9fe2231beee0e958dc42b51725d7ce29bd1208a993fd5f28b1d16021d4bf9f1ac7b4c004c1205c2a247107461aa55f932b0564af9bfb00b841dcf5603717030ed9534519cb8ac126919fbc96534d8c7da8f9d680ec2e4791652381a3a34cf398c52f8e8be3595c27d409ef60cc1ffcf4179ab7d2b0a31badf800b841e9a73d8d16d2a95dd55e222d00a9adefa23aebf1b6551d00a9912c831991cd9a1cc5534fa6adacfb154b4513b0d512554aee8a1d70bb43d0c736de5fe65fe74201b8415c5bd6bd4a923a75c8e7187e70852c6bd2780bd4e2c597a0d9e61ff7485b96f061c36b09fd6c4c29fcc11f066f8d11b0fa7d8f0a2e636e739a9662c416ed03e201b841a15b749646a4f741f1d6bce21e43bb6a77f32098c07cd0d4d3157119d97f44d7488741c71706653e53e10dc95663118199eca64e2e9f0585d0b22afef41b840301b8412e30b6d9dadfbee03cdcf1154548bb8b04703b75b75c85b0c1f3b6a9ffcdd4bb2538c569dc1b87f87c749a46932efa47dee2aed553af8f710edd29b9cf2e269a01b841b5a035b695453639a5051229216249b53873152bb94491dd56aa7d4be21d939b42aa67333955c441b9e67b1ad91640ef2aace89ea6ff8f7eb2b9eedb0ce802b101b8414ac5c90a63b50f96685a041092b6493c6986e32f78b95d306a1373c676040b2d5567fa94be899f64f67aebd9e381bcec69063c20f7fb59b4063776a7b77ab41a00b841fabf9bc0204d3c50a9a0e25827cb7c3229e1f5e8705999036fb83e4b09f5fdc05c4514d2336e82b41cd8225b3cd6e19113622f0be7cf6fb7ff9db38594de596900b84130b3813f0811482d8ba7bcea41f21398f88ad3368ac4043ff3d7759524664e296aeae9524b7f316a7c5f0c18ef794ef4ea7add536c0a648f4d5093805e5c73a001b8419fa4e5d2c1f1012aa8e63430274ca080bbf198569750ccef1b0538880553241119a620efa6f2c8b588dbffe3386b8a591cdf4029f5bcb0d195a079606661c44400b84188a413146dde351eb094f04188425579bbdb23f1a5d3aa8a9c2f0625c6e643796a1bfd5ed60aa0c82e90b678fa32670a0f5f5f5ac5e8fa175c7905eba265f47701b841030520c37dc58557c26d26893174c318d25eb644fc48213419455d31e869ad627218d8f1e05deb44426bd9a887b339983056f428cd3aecd357b4be6c68f7cc9600b8419b271de370a3bc401fe06be9a265d8bf8539d0716b2fc1e84e2f56ead058baae65d2a981e92596763bbe4a9ecd5cc170808fae0b8c6e136ad7e403847d48e7cf01b841137821a32062e924ac8bf1cd4cb9fc1eb03764f32e4294364a7788ccd9eb59d047809030a6e9a542e5b8ac6226443f964ea604d8a0027eb832f0165ea61cc50c00b841a0908745c2597bbba125df7faee9a058ff07c1a5c15478c81031dfac013da6c06705fd9e2a7a8b4909396bfd2c941029b142bf02dd34585eb7b7cef1353f798901b841593be7ef0de0d6d6900de685be3c6531dbdecb5e08ec2ebbeed90518ce85db28143693b9c39acc122ea84edceafaf56a1cadab010765f8726838e388b1e0826201b8411a26d30ae1c7b787a3bf0939823dca9ee2645f34cde6404f72279c1d9abaf17761c727210d2d256b14e3d5cb8fee49b64492a42e6d4dbc13e95ab70bc12d13c601b84152d2c4ff0cf0ff41bc5c2dc17b67e3d86a2e5f848639c14e0a67b40cff8c755d2de001d8557542fddad9e35cf2afdd6fd2d3de64f6e437e347ee18568aa66ffa00b841c445cb31a277d0cd4b37423898ad81196a4b48b2c72b02ce7e596f84343389ee785d99dec5383fdc239cc7a81fc720e79190de8430f2199ab7b28c97f1daad8d00b8411da564492b2fe5c7bee19558b4af2fadf29b371501766d1d63b795e87b882c991873a0e7d679264630262b2c675573a27502cf14224dbcba27b057e5b2d8b71101b841bd54b0a4ace16df1cd63613887fef2fbf541c862b40e39499c7c9e30f6d685a926fbc1b11331ded89601e27a45c44f4c78c8e01253d9b0485c65270d5ec8be4200b8413fd044a51016af6d3c5b57ac7e6e4ea6c4d1af9349eec4a283d7f219c9738bc55741cfb8743b714aee31e5a1aaab8031008e89cf8c7e817cc020905c44112a2101b841eec29d3ba5611c300d58bf6614542dbd9ef229c08feae504be87627b1100f74476ae939e5f7418a5944e85b895b9596aa23431390ba1bb609a866fa039bdfa4e00b841fd7d6e9825f3c4bbf974eecd7766397c0a89744e7fed9c06d464b6332c2b39a87ea72c414429c6981b3bfad602c5853706d02393beaa514c319f03f639702ae400b841be91ca9cebd516100ba6495bdb1fb70b9ac7b51c90953105d5bca6b5a8a973f17c7951966c553c1b7f11cb6b058e1f7be69cd9672d9ca8429e66a2c117cc3bcf00b84128dc77e12847efc39d37765508310ad503c9f4a59e8185444c7d3b2b8a5c71b80ef25661e0b28ec45202dbd14e02b9769b73293a7567a8ccf71c98357657c03401b841829d56d93c2bc19fc65338f2edf8b5ec396c91d0cc25286fd6660fed8ae43e2a674eac0e382deb187ae2c3102df69013ecb19401503248a5a1039d81bb87a99700b841f314c144095c65c7055534d3e489201561d3905d3d84be5b65db5e384ff3706b2e2f9850cf03b57ae4685a836a82bff4e843706188f99d6a27594c8ea82440f200b841a97bf554f35da0cfca1b3fa9bf05e8083989dbb7c5908e19e7f44d567757c5e77840b3b4cd908f283a07cdbf7ec4fa9d557d342ef8d4edb1b46ba9bc934aa68d00b8410ffb1bbaaba80cb9907af53e0494f52622e5c206211cd67b0fcb664a7918337623fcc4c9e2188b77de65f1974b2ed3d416fd5a9e8894c7fb20355b9fc0b7b25901b84106daaeb5288cf178e5a890b127ad3cb510a073d569acb0054a02c8d63359f35c48fe6797a28f2daf17fddfe750b476aff92d2146131e2585564e18a5e6d5409100b841d7079a73286fb3b9d77e21e7513aec8620d8ba320b15e7e599b9e9b00d4b8bd8440e71e3972fa5a961f579fbe18217911a7d332d8e6f0c2e4cc836e0e374b04800b8412f673c1f6bf4cf080f684ba96cb80b2fcec7c40268aeba377b34b4bfcee02ab03b68c6def89e1dd4555e6ce0b19a06864e458c336a5a029d744c3be7ae061acd00b841fa8c2522c12b2e9b603d2a274903f740f37fdb3fd22187b99cc0a2bdb81eb5650e53260ceaab7070b9532454d47b040c57fa285bed508823f0d96704fdf6fc0d01b841b9b5d381284f64c147157d9d77f50bd623fb7a35489012cd570b6fc73750873b31de4258fdc393664a2d303719381b536fff167b6079ac84102df6e23539c88700b84165ff2b78e7ff47a76c047c752b2abc6a57eb0cdda12e3441b58ac40bbdce96ab29baa0910310dd48c9f5c4f4005c4732280618cb0dbdd283e9e7cdc12536213001b84190bd9eb24bdfa05dd33f999bcc7ea7c791ef5a0c287df8f89cb6d74b8ce0cba260ed6b53cc2759452514a4d9da7d42ffbe7719810d266315bb9e8b790e18962d00b841ce3a39ec4c959cf3752b99c1f962c01853be4a99779adcec45216a234f6b1ff83a46ec4757f6095cb5f843c0e20de8ad79e42ab8c2caaacc1bf769f3af00fc5200b841a35df1ac147aea722e1980d3687303b849f422b7038477a9c1d50e1fe1d0457a656cf08c12b7e37e3548f3f352a3c5e7c0bb3d9d8be0026ba16885d63a6a292f01b8413d6c2ece8a92e8238aba60bfe505cf2944401a0654b4bed777154d83f0517e7962d67acfdad410821d8947dd772637908a789bda0123c18b8ed937cd30aa92f701b84130fac3174a2e48e846e0980ea4ab78569171fd6f70c11bf3a7b06dac7dc6fb7d01b0457713bb4370d11bef960a76a6025af80f3c6196ee8dd3e0383ace3153be00b841572ea236267892280ce29f9d7e009372fad8acd48888fb632570d295401042934b1ed2939a201dcb1f05b9f90ab8f1a9372f2215ffde420cc34db5fa586a9e9400b8416666fbee524ced1633fe3fe25fcc49b85cf93fb0a30f89de5f54a72471bf33981eb9a42cbf0f91cca13d84992cf60ff435211c5bfa9b62e244f4976e3e53dbc200b841135591e433c4a59160152746c6947f3bd4c1375acc2b8863181f334e4fd367134fe3cf9c2a93aa163f0dd5d81806dd6b90559b0c9a2c58de6401afb1c7cc516f00b841ca703546b9893b1c53d09b068c0db93c7ebb456c4dd6e06f692a3478aad01b2c2dcbfc3f33a4ebb3861e196514f497fd29461107212992fb5632ce0dedb9c83500b8416d73b3028f72ebf0dab31609bb5d394f6c2e302285ecaa58eb8d57ea029af05a6f9df958a640dc69e439cbced7b3ebee34b2e1a6815f02df5c106e22613ec80b00840548ee96",
"gasLimit": "0x1908b100",
"gasUsed": "0x7ff74",
"hash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"logsBloom": "0x00000020000000002000000000000000000000000000000000000000000000000080000008020000000000800200000000000000000420000000000000200000000000000000000800000008000008010000000000000000000000000000080000000000000000100000000000000000000000000000000000000010000800000000001000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000480000000040000000002000000000004000000000000000000000040000000000000000000000030000000000000000000000000000000000000000000040000000000000010",
"miner": "0x2f86476faa31c3f5a3d5b9376282e1b02b4dfa00",
"mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
"nonce": "0x0000000000000000",
"number": "0x548f4f1",
"parentHash": "0x23d9b15b34339bed8c20f4cea2aea76194a0e670c1b412dec9a77991ee418c50",
"penalties": "0x",
"receiptsRoot": "0x473f2a6554298d1b72b3a73e266bee07f5f93509534f716428cf3d233bb25ec7",
"sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
"size": "0x17f5",
"stateRoot": "0xc8ab77ef1dbae46ec2791933df17332d3e9f82b0d12433daea286bed53a90758",
"timestamp": "0x68254b01",
"totalDifficulty": "0x1baffba5d",
"transactions": [
{
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"blockNumber": "0x548f4f1",
"from": "0xd844d373f4ef676665fcc18d9b6c89464a36f2d2",
"gas": "0xf4240",
"gasPrice": "0x37e11d600",
"hash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"input": "0x414bf3890000000000000000000000003fb46c4db76d8e9f69f3f8388f43a7ca7e140807000000000000000000000000951857744785e80e2de051c32ee7b25f9c458c420000000000000000000000000000000000000000000000000000000000000064000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d20000000000000000000000000000000000000000000000000000000068254fab000000000000000000000000000000000000000000000fccf103db7d22e000000000000000000000000000000000000000000000000000acee08ed426a20ae840000000000000000000000000000000000000000000000000000000000000000",
"nonce": "0x5e8f",
"to": "0xecf4ea7907e779b8a7d0f90cb95fe06f43b610fb",
"transactionIndex": "0x0",
"value": "0x0",
"type": "0x0",
"v": "0x88",
"r": "0x5b1d0e4565493e5f815e26d622c6b9f5c7f23e269eac6ed950446ff7a10e3197",
"s": "0x25df537fe481c63c7defaaaccdd91105a0088f720789ecf5728460a4c3c0fc2"
},
{
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"blockNumber": "0x548f4f1",
"from": "0x2d40b858b6805d591d7f3d3affba6ec56c8f44fa",
"gas": "0x54c563",
"gasPrice": "0x37e11d600",
"hash": "0x7d094a854f038470cd419062bd5de9dbc99aa967e298f09d564a682500490e06",
"input": "0x8a054ac2",
"nonce": "0xfd8c9",
"to": "0x7aeb485080024b786a19cc76ad8857b8d4673a78",
"transactionIndex": "0x1",
"value": "0x0",
"type": "0x0",
"v": "0x87",
"r": "0xe29bd1d3345902626deb2f4fee492cb6ab8a6c757e84ef16b461c9272e6b013c",
"s": "0x4510436c2e4758d2c2808222eb964e248ca2a236b5d220252c10a04d6e045196"
},
{
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"blockNumber": "0x548f4f1",
"from": "0x2d40b858b6805d591d7f3d3affba6ec56c8f44fa",
"gas": "0x54c563",
"gasPrice": "0x37e11d600",
"hash": "0x1a2f811942a1615e4f0e97c1ea1e031238624d73cddc4976d6324c174259de37",
"input": "0x8a054ac2",
"nonce": "0xfd8ca",
"to": "0x7aeb485080024b786a19cc76ad8857b8d4673a78",
"transactionIndex": "0x2",
"value": "0x0",
"type": "0x0",
"v": "0x88",
"r": "0x5c1f2f9db92bf433c54414df6683cd20d61ca019cea958a2c1171f67ceab236",
"s": "0x36331a3dd1fa0b79bb071f11233528001b14ce73ac2af03ce4d69225cc20395a"
}
],
"transactionsRoot": "0xa417362c60be564f32c58a2134c4a96ae7667d170537c58a207e95faaf37fae9",
"uncles": [],
"validator": "0xdf8c7a62018bc45cdc88b44f3e4e63c1a1bfb9a34faf0231f9c01383120096f450eefb39c2e78df6a5a72a0b260ed2f241153be560980984327a078832f3c45801",
"validators": "0x"
}
}
eth_getBlockByHash executed successfully.

---

Executing eth_getBlockByNumber...
{
"jsonrpc": "2.0",
"id": 1001,
"result": {
"difficulty": "0x1",
"extraData": "0x02f9135583854f35f9134eeaa0061fd499015e34d4feda9d2999d04fcfb62ef4b543838e87dc44e11f4824d6d883854f3484054de0d7f9131bb841bb918f80f314dc85b124bd02109505e55f03dd87e56bfd60c299c80cf4e2dc392a71741bca32a5d26d8fb7dbfe2f9c7c89ae4d0838a1c09838966a081d24a50901b841e16ac88d93083d978ef822bd80484e5df6a0bb21b2946d0692af10c00138a35031b0997d39298c539a00db7e558ecc72c612f7111e27af58d9b97af1dc2480d400b84108e0cc8d8ee8e2b186d093a62625893030b53f00d206b328dd80e21699c123c329811651ca749cabe8e2ec5bb2f4fb8ab0f955339c31f76ec7f592e5d34e47c101b841dcf84ce24340c01bd53a8817c47a85964053eb173ff3374e3c40ec0086a04c7673520998e63f3a4b08801281f092fec6d7b83c6f7f5bfa4b8039cf65c18ccd5901b841607b95d852c82eb950ecccde15bc79d51d0b2a83bdb104a4a5359de29c94302e2437ad68f757dbdc1814a5be80a279804a425e3ec8d0b54cb8995b08f5e160f500b8417f3c43698a82e637cc0848fc24711159267e94634181d7d7e73ab47852492caa1c117fb0986e975280f9d9fe787f80c668efd11cfb023ab45197d7b8a538ac1500b841f8d0e9d03c8178e8b148998b820593fd46614b2763bf21fd2de06a54febb896d6e57aab0d9f6ce5c2e6427e6883b651024d27fd75e37b9950cecdfd3f142aadd00b84129d06f970e7b84fbf995ca2f96c712cbfd5d6fb20f3f7bad3dd190c94389b8d83e7fb3baa3f9de77f2c427608f0b24b799e2b0596a2223c769358206c089c7cc00b841df1e894d6c0539efedb4bbf6d6a90e949bf92a98484dc1f5be0c106b7c27375e6639ec3a8b3a57564fecf288c943859c121144d88684ed4336c7e57d32db254800b841819d72c6059d4c10d05679742652b29601b65e3c7e380afab0eba4924018a3a718133b871f38f8de4d86155271ba79f403314e85c9e170a2849cd8f502e5ad5600b8418f7b9402f397f23dd64861e0b5119a06dce5dec2f677d016d61b38f0d36753291a45fa75f40ae51c00c16e74841de24d49bcf013eec4b12535c0f3d71591beb100b841e99d34c8dae9e633bc0f64f8d0e5550f863f1d2bcbe0fa4e1e0b6a9c393fe48937cfa8938525b8d81abba9c11ef9ca4d93027453295ed49c3799a5e981e794d801b841bba5e9fddbf1a270cca74b64b398cabbe249cbc0df562f22238070cb590ec26e1c5f7621cb885028468d7f384eff83b7726af2ad7249c600a19cd94b9723786400b841a177f8572af09ed2f26e72d1ddd2245c69d401975aa382b7396b0eb2719f01bb6f16ee93ea67c8ff9d67d54f81f6853fd4640c575b324553beb0d58ed245991201b8418fb9dd3c58410ad8262811359d74428b924eb27f650bf628edc7fe3e88aaf6ef45f13b458dc753476814951c6597cc6a5d67c80da9b5dea4f229124a76daa03200b8412b6fb4ad6324b6aff0a4a1a2f8e3bb22dafcf655552a54a58c83f0549070c1315057eea4bc3927578b5932717f41c9cdee89a3178be2e37a889719646f7775c301b8415841f300509e6c26686a599128d166be3fde533971a91067176855fefee5740642983057c0000020c749fa8e0a5f92306adb6701e032a65d95557e34f824004c01b841321be0f862f7afdf9c2411cd7af95e8d9e13c1dde8a6ffb4e630ef8e5845b5672053a5b2d39b15631d2f1abce23664c4ae53682dad6b5ed9f9f86176c43a9e6500b84108831b235497507f1e66a45930a4df199b693fbc523a4a4a9f92ad7d0aff53ed4ddb22072d132425ad6b28559dc304fc83242f7797f8c15d0c702be951daa5a201b84173324b8f1155a8d2d4e86d825b6cec7258d6d0a2900766218bf6cff6aca86d457c018c3c603e6e6bdfe4c6da1a944bffbbad71aaa0975c647ab62ccb6adfde2501b8415ee6dc6ca30d9aff02565991297608b140af423330989ac9bdd488d4c13755441df08af65ba59c8d787291c023820e1f563c2e6df88aadbe789216d20b8fa75f01b8419ab09916eda1ac7a7205dbaae83df38b0eedad544f97407bcb9b1f26b7505c8c257a0ed0060e83a6b507bde3da9dee5cc5f746977bbd9944cb4e955241d381c301b841c8f8cae6b6ecb82becb1b35b8c96e34f8e93060bc51b2550c752f7da347fb4537211d390b23681c7a3d90461e8ec4405a2dd21e1c1a5b35884869fdca065b1ef00b8418eae179cae70a3c00f193dca0667811232d1dd72d81c869211641eb33b8f1f51229fccb594820f124b306ec9254897511acbe33bd74dd16dd6bef477d7ee976f00b8413a617ab32db378afdb5cf3c6c9f63cb47678d085fc93d86f0effc9957d8314785bc23809ad65d1e9d3006e0fb323a2dbbf2942ab1c7243052c6c52a38a93cecc01b8413dc7c3b417fea551a5ad7cdd4df7569ec9817a2c188f9cac054e6e3b8f49e8de3821b7906fb89447778c5e0bf9121bef11b62d7456a67fdd5a3406ed3a631ad901b841e6d66960f4b339b1de2765af43c716fcb945e643cc3979e5cf70a2edde59217b6cb42e824ffe1fc2f9338104c0b739aef897d69c241b266f81aef4d4af6772c601b841bf1472bb8508e5492b955c1d8e85885a67fc4ab78eb9788dafabd9561b2f7ca9614d89ac37ffcf4a7011be7c7e8a0f1ba83c4b1be736da7b051bb31882ea44a300b8411a9746c747066cded1ddfa54020646f3aa0a2bed8f1ee2a5b0b7604e40ba7e0341e5361d1bdcc064c05e8501527b5e99a0adc4a05f6ce9fe40a5bbc29b0530c500b841e9d413c3cf810423f0f662384761b491037eead9e4bac4eb7c0b6fd47bf2b1805e210ef9cb62ba237dce4da9549b9ea8d38d97a5238f0eb51fdca8feac6078e501b841bc60926d3a71f9d7b4e1ef272163e655ff68164d2292bbf14918ca24884dd7fa4d15b7f11f89ba4b90d752fd4310b5c23b80677c0b08b589dd2623934971680600b841e66034f5f8d38970c3f1612b0367e4506861322df6ca8d4733bdf5a5027d22081870299a430fd87a508aeca5868ebc8ea95dd35d44dfa723f15789756319cee601b841563111795b2951e6c37dc6589e2d28126053ad3cf785a9f6305f28cb023bd8495063ecaad988da577fbd2d56d679ad1ad6317a10416812eb4b73a0a56f37202100b8415e995d237eca0dbe160d7d2096739507c0a0eb23134c7913881e75a55e64d05a4055501d036523073981c39644d281087ca140dde9df55f846e90287972b8c5e00b841d26c93f33e7f1de89b6ea290059649d4e3630c49b99762b9fdd8c7a12388439b51a6b413b393213fd6474889d2810f1857e9f8ae8b3510c1d557d63fa228983001b841b065a9015a602a08e71ba65a826b38b33e5095d0a34e3c64cba90813b1bd6b575a83e2283469f88fc8de079f10618c5fd4d478bbbfbb7a6077769813f3a0f14600b8412e8162189d6976d2e4d24d7b23a2ee63242ff8b138ad88f50965b68edc35355724920cbeb93b725db8906e5cfcb160142eee6821868ea821516b996a73a11fe101b8417bc786f8ec57acbc5d069470fb9fe67d134507d0a0a7431da8ffcae74a4b7b262e963f69bb580778f55791f69c1b4ca70956865e0caea481ec4383d3e67c854a00b841f20cccfe09fd5e1ff3004c013b1aa842149acf4bc8ed8703fdbe47311bc1e30630c547f501c6ca8b40b09e4c101b9a7acf141d1a78ed830e508f91504950303c00b8410821e6b2e1c895f2fd6b85c153a530a24f457c9a40da325d10c734ec060e19bf2d359b584c177537a8c8583bc6874822cc1c489b370bdef17899ffc30f14042c01b841484875795908fe5bdc0ed5d847a16690a306c8cccd04e8e05d50f2223464e3cd304e01efc20dedb36d5c6d69a5ebf53a551d1d63f143859a2ee7a63577043ca600b84178dc67868032bc6a405a55b6553d8da7b09ff2bbb3a888f380bec6042aef26686ec0ef9907388ed37425c87c2483cd56c530dbca5f196264e35c0ca5f65ce44601b84157ddcc1a6e6d66a75f1e6d0aeea9d7cd8c6c10b1a6d75959f1e7d45d9d26082328549a5500e7a08c447ed2e51f78c52f468ecacbd069fae92ee9da161e96bc1101b84162d6ba81ddddadc2dbaa9c7be74bb4c103224b655576f4a12e90f822124b26e85a45fea7bcab74913d4d4d7fc7689f1e968c4bf3901e7b80c0967e0dd9a7089c00b84175e5e7d6752eda933848ddbeb57f7f84ed88d59b5cbcb5abdd574faa8821a0ad721da68205127f7b5a6cb3f78c936d900a7be4da406f970d7e750e7b77f6df3800b84150ecd5a981fb85df9915c37b7bd0773070cc6144c5722c1626e79fb07e5adfec65f1e17a14a23cac6aefb6b9ebc34144d45f1ca1769690e0e17511f5a669967a01b8416fa846aa3ba78418980559af84d8720e3e1137ae6104df215fc1e0520e1305bc127d959edb0b5c3c34080ac5ef76e50c9f233671ca30f336035afd069bd7737c00b841dd554faaade8cfea5faf2ae55cfa49bddb936606a1629a309ee15d0346286f1c40b42864eaa6426e27f5ea7a66cdce7a23d1b0727c933efb29479843cfb2a57b00b8412113804eed512be8a9e8b365c1cbac0671294c841bf297c4cc2bac72b882af2310ed7f0866019a30a95060c855a5641ce829bb30dbb1c3795a2332296797dd6a01b8416ec27547261bcea2fe2a7952fa8f6d541152810888c0c56681957ae85fd69ee3727094cf1d3c7977f82932f23ff6329ed75afeaa6155f4ea8d149e7e2d01b06800b841e8a8c72e305b045d0b3b21b4fcc8780be3485ce5c9b7cc78b58d04a055b41f7954155fb4310bc411fa428b94155f40e68955c70435c7c32ce83d28065711e31b01b84160e8c0a6edfa062be06b8b1160993a03401eb4c83e8bf648c158dab9b3e296df0cfa27606ac2eba0fda862be86952174d1769c282d89e3ca420ae39f3aa96b9200b8417a7711d80d9e175e6261fd7d9bab2b1d5d80eca930f714d280176ad441b662096fecd17291a52aeb6cc8f98fa952472e609a39184100efbc74e5a917d922bfda00b8419c9b40a9829fb5d3148307138aff523847d433f8e3c796412bb6673db90647914cd78027af92cb8bc64c221e26912cc6e25e4ebb1a7eb7a516809dc0779e0cff01b84152ef7390bda52070483f216a1185b085992e48e964a0894fb625e02847f2baf25597859e3f232b0540b92b204c5b250901dcd6d6a5fd371954ff97581d7b9f6d00b841c6431d27fab8a7a02663cc8d2b07bd41eadbe658615e4e6279822925270953416638b656ab2142ea9bc02db67c5ba427664e14aa81cc0cd35ad200f55d3473e601b8413c209be1003bf4f36829e32662dc3c9e9c008bdd0b03c81133ebd7cbe6fe3f0877a4fbbe04fe9b5844b64816c8bb2731290049656bcd7802f42aed59bf831f6100b84166ec5cd606e40a08efbbbcb92ed1816abb31fa85d234ffbc7b7e8871202bcfb10143ac4ea51de22bacdae64031d5231ba16c50d32e8d883b55b76a1b2e53929b01b841a7bc5749cb57c5aaa7f1f33ba7899a7e91426e75cfac826e7125dd3b36d6c7884c0a340b426710045cd682f873d8dd0216749aa82b6f20fd98d48e52aa233ee601b841cc9a4417c16d660c0dd3885b8cf7ce0063f421a496f2b814082df2e5bbd2906f77f9850bbbd9e227ff1f761b129ca5825a9b0b1056f6c6b630d5a44f628897cc00b841bcf19c8f727c7e6b9e813563300eebfb103528252c1a4023e751c77050f00f396aaab3869ca6c8507c439910d655b85e1bab707228e0e66b258e67d7ad70e84f01b8416e57c412320281b684b625d94f7a9597827252863621611ccfe4282a9e1600f16aeb63cc53215f0681b3ee82f49452e6af3377f2cbfbbb1508be80cea29c254701b841f57466fdc258117a75c727941b5aa46352a57d10ad8a15db362b7f682df1aa341b3c62b72f7f207c4c6607ff6739a8bda1166970e42c28d7ab57228ae890897800b8418cca97e6226c39de135f9c06ae4d35545d0c78356a7ede27d90a10fd1a56898e3167fd5bbf3ebc8d8ecb7a74d91f0e9dff4807cd59eef0026b779ba7108d3ae300b841d9ceea06505d2ddee859b66909feffebc77ee8c0cf7929b5aa402a6c5889476607e1eaa64e92bd11d01d3cf7f8b50b386aaecb53fedec9c93f300f53b473b7c100b841f42a941a08007a41e7e5c0cca1702f8889da56b74dcc34c51926afd97bb40a9b70a242e8adb15f9f85dc3d7c8f8254282b647fae622b712901943e49e298b4ca00b841898567c9a022584e754624c6e7dbe8e62fee746b86d13c72f6cb99c1c488a4bd6f7864047e823f4e9006de46741ea440783afec68873c39bd7f13b885783067b00b8419ee4ad343cfc6ba02445ccf658044a72c94a6bbfcc240cc7f7e5d3455a896367462441d3226e9bf994891157635cab72b6e64217e6b6da002d0b8c1be90482c800b8417cbb0b8692d35d391dfe2456cbed764e0acef16e2b7a1065b5b49ad1cb197b71315b73ccb3c31e52197e44a42f472056c696c51c66f74bbdbd960a140f0a1f4d01b8414d70843ba36303dd70b3baf064eccd4bedd068a6e962a4e5fc0bea6b43720e1a29f9151777ea61c951151d9888bc3a853cb43e67d5b836399b379631f56999d700b841cb2d150f835fb2481c977403e67553a9019abf27e9c91258cb7000b4cd02f2a0126f69a0353db0d9ed0418ec3244935ca62fc8e3e10cd89b083f47d4480f4d1d00b84123bda1cc6787f85bbac5eeefebe4cfbdba4a9385ce32738e2192bf34fc9bc51c373a1094a035bd879d9f20394a8e8cafc55773eac726a126f1b9b5e5e027f4df01b841ea27069c2f663af3825d835482e1b860c8b5a723481433da277b3e75729d4b306ed730f8eb1dece93c97f65c03430843d5323493556b664fe9c46c994b468cc80184054ddcb2",
"gasLimit": "0x1908b100",
"gasUsed": "0x0",
"hash": "0xd44373810f545bafcd5036f60a5b70b4e1b4a402db1a1392eaa400872effe153",
"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
"miner": "0x9043fafbe8833adafb449b559cef300a9632763f",
"mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
"nonce": "0x0000000000000000",
"number": "0x54de0d8",
"parentHash": "0x061fd499015e34d4feda9d2999d04fcfb62ef4b543838e87dc44e11f4824d6d8",
"penalties": "0x",
"receiptsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
"sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
"size": "0x15a1",
"stateRoot": "0x7a7b91e79eb74368cb2e808ce44624997781319d42237868e6cfce16fda344be",
"timestamp": "0x682fea18",
"totalDifficulty": "0x1bb04a644",
"transactions": [],
"transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
"uncles": [],
"validator": "0x4afbd8eca5164b3cac0c35f2f9e2b3ee758d8db7e675b5b0f37a230996db32250077b57cd44e659f79613961f38ee5093b8e14c044a2561b476eba69667f06d000",
"validators": "0x"
}
}
eth_getBlockByNumber executed successfully.

---

Executing eth_getBlockReceipts...
{
"id": 1004,
"jsonrpc": "2.0",
"result": []
}
eth_getBlockReceipts executed successfully.

---

Executing eth_getBlockTransactionCountByHash...
{
"id": 1004,
"jsonrpc": "2.0",
"result": "0x3"
}
eth_getBlockTransactionCountByHash executed successfully.

---

Executing eth_getBlockTransactionCountByNumber...
{
"id": 1004,
"jsonrpc": "2.0",
"result": "0x0"
}
eth_getBlockTransactionCountByNumber executed successfully.

---

Executing eth_getCode...
{
"jsonrpc": "2.0",
"id": 1004,
"result": "0x606060405260043610610196576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063012679511461019b578063025e7c27146101c957806302aa9be21461022c57806306a49fce1461026e5780630db02622146102d85780630e3e4fb81461030157806315febd68146103715780632a3640b1146103a85780632d15cc041461042a5780632f9c4bba146104b8578063302b687214610522578063326586521461058e5780633477ee2e14610640578063441a3e70146106a357806358e7525f146106cf5780635b860d271461071c5780635b9cd8cc146107695780636dd7d8ea1461082457806372e44a3814610852578063a9a981a31461089f578063a9ff959e146108c8578063ae6e43f5146108f1578063b642facd1461092a578063c45607df146109a3578063d09f1ab4146109f0578063d161c76714610a19578063d51b9e9314610a42578063d55b7dff14610a93578063ef18374a14610abc578063f2ee3c7d14610ae5578063f5c9512514610b1e578063f8ac9dd514610b4c575b600080fd5b6101c7600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b75565b005b34156101d457600080fd5b6101ea60048080359060200190919050506111fc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561023757600080fd5b61026c600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061123b565b005b341561027957600080fd5b610281611796565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102c45780820151818401526020810190506102a9565b505050509050019250505060405180910390f35b34156102e357600080fd5b6102eb61182a565b6040518082815260200191505060405180910390f35b341561030c57600080fd5b610357600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611830565b604051808215151515815260200191505060405180910390f35b341561037c57600080fd5b610392600480803590602001909190505061185f565b6040518082815260200191505060405180910390f35b34156103b357600080fd5b6103e8600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506118bb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561043557600080fd5b610461600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611909565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156104a4578082015181840152602081019050610489565b505050509050019250505060405180910390f35b34156104c357600080fd5b6104cb6119dc565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561050e5780820151818401526020810190506104f3565b505050509050019250505060405180910390f35b341561052d57600080fd5b610578600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611a79565b6040518082815260200191505060405180910390f35b341561059957600080fd5b6105c5600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611b03565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106055780820151818401526020810190506105ea565b50505050905090810190601f1680156106325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561064b57600080fd5b6106616004808035906020019091905050611da2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156106ae57600080fd5b6106cd6004808035906020019091908035906020019091905050611de1565b005b34156106da57600080fd5b610706600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061208d565b6040518082815260200191505060405180910390f35b341561072757600080fd5b610753600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506120d9565b6040518082815260200191505060405180910390f35b341561077457600080fd5b6107a9600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506121a1565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107e95780820151818401526020810190506107ce565b50505050905090810190601f1680156108165780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610850600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061226a565b005b341561085d57600080fd5b610889600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612653565b6040518082815260200191505060405180910390f35b34156108aa57600080fd5b6108b261266b565b6040518082815260200191505060405180910390f35b34156108d357600080fd5b6108db612671565b6040518082815260200191505060405180910390f35b34156108fc57600080fd5b610928600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612677565b005b341561093557600080fd5b610961600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612c36565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156109ae57600080fd5b6109da600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612ca2565b6040518082815260200191505060405180910390f35b34156109fb57600080fd5b610a03612cee565b6040518082815260200191505060405180910390f35b3415610a2457600080fd5b610a2c612cf4565b6040518082815260200191505060405180910390f35b3415610a4d57600080fd5b610a79600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612cfa565b604051808215151515815260200191505060405180910390f35b3415610a9e57600080fd5b610aa6612d53565b6040518082815260200191505060405180910390f35b3415610ac757600080fd5b610acf612d59565b6040518082815260200191505060405180910390f35b3415610af057600080fd5b610b1c600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612d63565b005b3415610b2957600080fd5b610b4a600480803590602001908201803590602001919091929050506134f1565b005b3415610b5757600080fd5b610b5f6135f0565b6040518082815260200191505060405180910390f35b6000600b543410151515610b8857600080fd5b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050141580610c1c57506000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050115b1515610c2757600080fd5b81600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16151515610c8457600080fd5b610cd934600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101546135f690919063ffffffff16565b915060088054806001018281610cef919061362d565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506060604051908101604052803373ffffffffffffffffffffffffffffffffffffffff16815260200160011515815260200183815250600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160010155905050610eb834600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546135f690919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610f5160016009546135f690919063ffffffff16565b6009819055506000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905014156110185760078054806001018281610fb6919061362d565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600a600081548092919060010191905055505b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806001018281611069919061362d565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806001018281611109919061362d565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1338434604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b60078181548110151561120b57fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000828280600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156112cd57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561140657600b546113f882600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461361490919063ffffffff16565b1015151561140557600080fd5b5b61145b84600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461361490919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555061153384600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461361490919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506115cb43600f546135f690919063ffffffff16565b9250611632846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000868152602001908152602001600020546135f690919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000858152602001908152602001600020819055506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010180548060010182816116db9190613659565b9160005260206000209001600085909190915055507faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2338686604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050505050565b61179e613685565b600880548060200260200160405190810160405280929190818152602001828054801561182057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116117d6575b5050505050905090565b600a5481565b60056020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000838152602001908152602001600020549050919050565b6006602052816000526040600020818154811015156118d657fe5b90600052602060002090016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611911613685565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156119d057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611986575b50505050509050919050565b6119e4613699565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101805480602002602001604051908101604052809291908181526020018280548015611a6f57602002820191906000526020600020905b815481526020019060010190808311611a5b575b5050505050905090565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b611b0b6136ad565b611b1482612cfa565b15611c655760036000611b2684612c36565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600160036000611b6f86612c36565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905003815481101515611bba57fe5b90600052602060002090018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c595780601f10611c2e57610100808354040283529160200191611c59565b820191906000526020600020905b815481529060010190602001808311611c3c57829003601f168201915b50505050509050611d9d565b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905003815481101515611cf657fe5b90600052602060002090018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611d955780601f10611d6a57610100808354040283529160200191611d95565b820191906000526020600020905b815481529060010190602001808311611d7857829003601f168201915b505050505090505b919050565b600881815481101515611db157fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008282600082111515611df457600080fd5b814310151515611e0357600080fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600084815260200190815260200160002054111515611e6457600080fd5b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010182815481101515611eb357fe5b906000526020600020900154141515611ecb57600080fd5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008681526020019081526020016000205492506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000868152602001908152602001600020600090556000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010184815481101515611fc457fe5b9060005260206000209001600090553373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561201357600080fd5b7ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568338685604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a15050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050919050565b60008082600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16151561213857600080fd5b61214184612c36565b915061214b612d59565b6064600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540281151561219757fe5b0492505050919050565b6003602052816000526040600020818154811015156121bc57fe5b9060005260206000209001600091509150508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156122625780601f1061223757610100808354040283529160200191612262565b820191906000526020600020905b81548152906001019060200180831161224557829003601f168201915b505050505081565b600c54341015151561227b57600080fd5b80600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff1615156122d757600080fd5b61232c34600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101546135f690919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054141561249b57600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480600101828161244b919061362d565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b61252d34600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546135f690919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc338334604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b60046020528060005260406000206000915090505481565b60095481565b600f5481565b6000806000833373ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561271957600080fd5b84600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16151561277557600080fd5b6000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548160ff0219169083151502179055506127e6600160095461361490919063ffffffff16565b600981905550600094505b6008805490508510156128bb578573ffffffffffffffffffffffffffffffffffffffff1660088681548110151561282457fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156128ae5760088581548110151561287b57fe5b906000526020600020900160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556128bb565b84806001019550506127f1565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054935061299284600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461361490919063ffffffff16565b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612a7243600e546135f690919063ffffffff16565b9250612ad9846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000868152602001908152602001600020546135f690919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000858152602001908152602001600020819055506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018054806001018281612b829190613659565b9160005260206000209001600085909190915055507f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d33387604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b600d5481565b600e5481565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff169050919050565b600b5481565b6000600a54905090565b600080612d6e613685565b600080600033600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff161515612dcf57600080fd5b87600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff161515612e2b57600080fd5b612e3433612c36565b9750612e3f89612c36565b9650600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515612ed757600080fd5b6001600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550604b612fc4612d59565b6064600460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540281151561301057fe5b041015156134e65760016008805490500360405180591061302e5750595b9080825280602002602001820160405250955060009450600093505b600880549050841015613357578673ffffffffffffffffffffffffffffffffffffffff166130b160088681548110151561308057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612c36565b73ffffffffffffffffffffffffffffffffffffffff16141561334a576130e3600160095461361490919063ffffffff16565b6009819055506008848154811015156130f857fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16868680600101975081518110151561313857fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060088481548110151561318357fe5b906000526020600020900160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600160006008868154811015156131c457fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549060ff021916905560018201600090555050600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006132bb91906136c1565b600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061330691906136e2565b600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090555b838060010194505061304a565b600092505b600780549050831015613439578673ffffffffffffffffffffffffffffffffffffffff1660078481548110151561338f57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561342c576007838154811015156133e657fe5b906000526020600020900160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600a6000815480929190600190039190505550613439565b828060010193505061335c565b7fe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e8787604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156134d15780820151818401526020810190506134b6565b50505050905001935050505060405180910390a15b505050505050505050565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060010182816135429190613703565b91600052602060002090016000848490919290919250919061356592919061372f565b50507f949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a338383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a15050565b600c5481565b600080828401905083811015151561360a57fe5b8091505092915050565b600082821115151561362257fe5b818303905092915050565b8154818355818115116136545781836000526020600020918201910161365391906137af565b5b505050565b8154818355818115116136805781836000526020600020918201910161367f91906137af565b5b505050565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b50805460008255906000526020600020908101906136df91906137d4565b50565b508054600082559060005260206000209081019061370091906137af565b50565b81548183558181151161372a5781836000526020600020918201910161372991906137d4565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061377057803560ff191683800117855561379e565b8280016001018555821561379e579182015b8281111561379d578235825591602001919060010190613782565b5b5090506137ab91906137af565b5090565b6137d191905b808211156137cd5760008160009055506001016137b5565b5090565b90565b6137fd91905b808211156137f957600081816137f09190613800565b506001016137da565b5090565b90565b50805460018160011615610100020316600290046000825580601f106138265750613845565b601f01602090049060005260206000209081019061384491906137af565b5b505600a165627a7a72305820f5bbb127b52ce86c873faef85cff176563476a5e49a3d88eaa9a06a8f432c9080029"
}
eth_getCode executed successfully.

---

Executing eth_getLogs...
{
"id": 1004,
"jsonrpc": "2.0",
"result": []
}
eth_getLogs executed successfully.

---

Executing eth_getOwnerByCoinbase...
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0x0000000000000000000000000000000000000000"
}
eth_getOwnerByCoinbase executed successfully.

---

Executing eth_getProof...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32601,
"message": "the method eth_getProof does not exist/is not available",
"data": {
"trace_id": "fcd232dd232d919ae534ecc3337c3783"
}
}
}
eth_getProof executed successfully.

---

Executing eth_getStorageAt...
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0x0000000000000000000000000000000000000000000000000000000000000000"
}
eth_getStorageAt executed successfully.

---

Executing eth_getRawTransactionByBlockHashAndIndex...
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0xf9016e825e8f85037e11d600830f424094ecf4ea7907e779b8a7d0f90cb95fe06f43b610fb80b90104414bf3890000000000000000000000003fb46c4db76d8e9f69f3f8388f43a7ca7e140807000000000000000000000000951857744785e80e2de051c32ee7b25f9c458c420000000000000000000000000000000000000000000000000000000000000064000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d20000000000000000000000000000000000000000000000000000000068254fab000000000000000000000000000000000000000000000fccf103db7d22e000000000000000000000000000000000000000000000000000acee08ed426a20ae8400000000000000000000000000000000000000000000000000000000000000008188a05b1d0e4565493e5f815e26d622c6b9f5c7f23e269eac6ed950446ff7a10e3197a0025df537fe481c63c7defaaaccdd91105a0088f720789ecf5728460a4c3c0fc2"
}
eth_getRawTransactionByBlockHashAndIndex executed successfully.

---

Executing eth_getRawTransactionByBlockNumberAndIndex...
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0x"
}
eth_getRawTransactionByBlockNumberAndIndex executed successfully.

---

Executing eth_getRawTransactionByHash...
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0x"
}
eth_getRawTransactionByHash executed successfully.

---

Executing eth_getRewardByHash...
{
"id": 1001,
"jsonrpc": "2.0",
"result": {}
}
eth_getRewardByHash executed successfully.

---

Executing eth_getTransactionAndReceiptProof...
{
"jsonrpc": "2.0",
"id": 1001,
"result": {
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"key": "0x80",
"receiptProofKeys": [
"0x473f2a6554298d1b72b3a73e266bee07f5f93509534f716428cf3d233bb25ec7",
"0xf8487e8af578fb2cc3fef425f28cbb15fab131e3557eb67874f1a80e7fc1a9f5"
],
"receiptProofValues": [
"0xf851a068bb429e7d2063f286c6091187ce9c7d2147c2644efd9a8944934a6dc14f50be80808080808080a0f8487e8af578fb2cc3fef425f28cbb15fab131e3557eb67874f1a80e7fc1a9f58080808080808080",
"0xf904a530b904a1f9049e01830256bab9010000000020000000002000000000000000000000000000000000000000000000000080000008020000000000800200000000000000000420000000000000200000000000000000000800000008000008010000000000000000000000000000080000000000000000100000000000000000000000000000000000000010000800000000001000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000480000000040000000002000000000004000000000000000000000040000000000000000000000030000000000000000000000000000000000000000000040000000000000010f90393f89b943fb46c4db76d8e9f69f3f8388f43a7ca7e140807f863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa0000000000000000000000000712d30c0975386b043e09224433f400ea87cab7da000000000000000000000000061b7b0009fced05695ee811b7f8f78ba37c38344a00000000000000000000000000000000000000000000000000a5ae491119b0000f89b94951857744785e80e2de051c32ee7b25f9c458c42f863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa0000000000000000000000000712d30c0975386b043e09224433f400ea87cab7da0000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d2a00000000000000000000000000000000000000000000000b593895f85bc3bea71f89b943fb46c4db76d8e9f69f3f8388f43a7ca7e140807f863a0ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa0000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d2a0000000000000000000000000712d30c0975386b043e09224433f400ea87cab7da0000000000000000000000000000000000000000000000fccf103db7d22e00000f89b943fb46c4db76d8e9f69f3f8388f43a7ca7e140807f863a08c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a0000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d2a0000000000000000000000000ecf4ea7907e779b8a7d0f90cb95fe06f43b610fba0000000000000000000000000000000000000314dc3a0f41e5c923c6637280000f9011c94712d30c0975386b043e09224433f400ea87cab7df863a0c42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67a0000000000000000000000000ecf4ea7907e779b8a7d0f90cb95fe06f43b610fba0000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d2b8a0000000000000000000000000000000000000000000000fccf103db7d22e00000ffffffffffffffffffffffffffffffffffffffffffffff4a6c76a07a43c4158f0000000000000000000000000000000000000000363decaa6e4099f6e686ca910000000000000000000000000000000000000000220948b5cb625b7056562c32ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff86c3"
],
"receiptRoot": "0x473f2a6554298d1b72b3a73e266bee07f5f93509534f716428cf3d233bb25ec7",
"txProofKeys": [
"0xa417362c60be564f32c58a2134c4a96ae7667d170537c58a207e95faaf37fae9",
"0x254e83386c61197750483c05b48b21e2d356827082cade0d9107d6015e2f28a5"
],
"txProofValues": [
"0xf851a024f4fd8be1fcb465204601a37ba06036414fae5bdbf213c0d3e778da088e5d1280808080808080a0254e83386c61197750483c05b48b21e2d356827082cade0d9107d6015e2f28a58080808080808080",
"0xf9017530b90171f9016e825e8f85037e11d600830f424094ecf4ea7907e779b8a7d0f90cb95fe06f43b610fb80b90104414bf3890000000000000000000000003fb46c4db76d8e9f69f3f8388f43a7ca7e140807000000000000000000000000951857744785e80e2de051c32ee7b25f9c458c420000000000000000000000000000000000000000000000000000000000000064000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d20000000000000000000000000000000000000000000000000000000068254fab000000000000000000000000000000000000000000000fccf103db7d22e000000000000000000000000000000000000000000000000000acee08ed426a20ae8400000000000000000000000000000000000000000000000000000000000000008188a05b1d0e4565493e5f815e26d622c6b9f5c7f23e269eac6ed950446ff7a10e3197a0025df537fe481c63c7defaaaccdd91105a0088f720789ecf5728460a4c3c0fc2"
],
"txRoot": "0xa417362c60be564f32c58a2134c4a96ae7667d170537c58a207e95faaf37fae9"
}
}
eth_getTransactionAndReceiptProof executed successfully.

---

Executing eth_getTransactionByBlockHashAndIndex...
{
"jsonrpc": "2.0",
"id": 1001,
"result": {
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"blockNumber": "0x548f4f1",
"from": "0xd844d373f4ef676665fcc18d9b6c89464a36f2d2",
"gas": "0xf4240",
"gasPrice": "0x37e11d600",
"hash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"input": "0x414bf3890000000000000000000000003fb46c4db76d8e9f69f3f8388f43a7ca7e140807000000000000000000000000951857744785e80e2de051c32ee7b25f9c458c420000000000000000000000000000000000000000000000000000000000000064000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d20000000000000000000000000000000000000000000000000000000068254fab000000000000000000000000000000000000000000000fccf103db7d22e000000000000000000000000000000000000000000000000000acee08ed426a20ae840000000000000000000000000000000000000000000000000000000000000000",
"nonce": "0x5e8f",
"to": "0xecf4ea7907e779b8a7d0f90cb95fe06f43b610fb",
"transactionIndex": "0x0",
"value": "0x0",
"type": "0x0",
"v": "0x88",
"r": "0x5b1d0e4565493e5f815e26d622c6b9f5c7f23e269eac6ed950446ff7a10e3197",
"s": "0x25df537fe481c63c7defaaaccdd91105a0088f720789ecf5728460a4c3c0fc2"
}
}
eth_getTransactionByBlockHashAndIndex executed successfully.

---

Executing eth_getTransactionByBlockNumberAndIndex...
{
"jsonrpc": "2.0",
"id": 1001,
"result": {
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"blockNumber": "0x548f4f1",
"from": "0xd844d373f4ef676665fcc18d9b6c89464a36f2d2",
"gas": "0xf4240",
"gasPrice": "0x37e11d600",
"hash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"input": "0x414bf3890000000000000000000000003fb46c4db76d8e9f69f3f8388f43a7ca7e140807000000000000000000000000951857744785e80e2de051c32ee7b25f9c458c420000000000000000000000000000000000000000000000000000000000000064000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d20000000000000000000000000000000000000000000000000000000068254fab000000000000000000000000000000000000000000000fccf103db7d22e000000000000000000000000000000000000000000000000000acee08ed426a20ae840000000000000000000000000000000000000000000000000000000000000000",
"nonce": "0x5e8f",
"to": "0xecf4ea7907e779b8a7d0f90cb95fe06f43b610fb",
"transactionIndex": "0x0",
"value": "0x0",
"type": "0x0",
"v": "0x88",
"r": "0x5b1d0e4565493e5f815e26d622c6b9f5c7f23e269eac6ed950446ff7a10e3197",
"s": "0x25df537fe481c63c7defaaaccdd91105a0088f720789ecf5728460a4c3c0fc2"
}
}
eth_getTransactionByBlockNumberAndIndex executed successfully.

---

Executing eth_getTransactionByHash...
{
"jsonrpc": "2.0",
"id": 1001,
"result": {
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"blockNumber": "0x548f4f1",
"from": "0xd844d373f4ef676665fcc18d9b6c89464a36f2d2",
"gas": "0xf4240",
"gasPrice": "0x37e11d600",
"hash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"input": "0x414bf3890000000000000000000000003fb46c4db76d8e9f69f3f8388f43a7ca7e140807000000000000000000000000951857744785e80e2de051c32ee7b25f9c458c420000000000000000000000000000000000000000000000000000000000000064000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d20000000000000000000000000000000000000000000000000000000068254fab000000000000000000000000000000000000000000000fccf103db7d22e000000000000000000000000000000000000000000000000000acee08ed426a20ae840000000000000000000000000000000000000000000000000000000000000000",
"nonce": "0x5e8f",
"to": "0xecf4ea7907e779b8a7d0f90cb95fe06f43b610fb",
"transactionIndex": "0x0",
"value": "0x0",
"type": "0x0",
"v": "0x88",
"r": "0x5b1d0e4565493e5f815e26d622c6b9f5c7f23e269eac6ed950446ff7a10e3197",
"s": "0x25df537fe481c63c7defaaaccdd91105a0088f720789ecf5728460a4c3c0fc2"
}
}
eth_getTransactionByHash executed successfully.

---

Executing eth_getTransactionCount...
{
"id": 1001,
"jsonrpc": "2.0",
"result": "0x58"
}
eth_getTransactionCount executed successfully.

---

Executing eth_getTransactionReceipt...
{
"jsonrpc": "2.0",
"id": 5002,
"result": {
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"blockNumber": "0x548f4f1",
"contractAddress": null,
"cumulativeGasUsed": "0x256ba",
"from": "0xd844d373f4ef676665fcc18d9b6c89464a36f2d2",
"gasUsed": "0x256ba",
"logs": [
{
"address": "0x3fb46c4db76d8e9f69f3f8388f43a7ca7e140807",
"topics": [
"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
"0x000000000000000000000000712d30c0975386b043e09224433f400ea87cab7d",
"0x00000000000000000000000061b7b0009fced05695ee811b7f8f78ba37c38344"
],
"data": "0x0000000000000000000000000000000000000000000000000a5ae491119b0000",
"blockNumber": "0x548f4f1",
"transactionHash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"transactionIndex": "0x0",
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"logIndex": "0x0",
"removed": false
},
{
"address": "0x951857744785e80e2de051c32ee7b25f9c458c42",
"topics": [
"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
"0x000000000000000000000000712d30c0975386b043e09224433f400ea87cab7d",
"0x000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d2"
],
"data": "0x0000000000000000000000000000000000000000000000b593895f85bc3bea71",
"blockNumber": "0x548f4f1",
"transactionHash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"transactionIndex": "0x0",
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"logIndex": "0x1",
"removed": false
},
{
"address": "0x3fb46c4db76d8e9f69f3f8388f43a7ca7e140807",
"topics": [
"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
"0x000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d2",
"0x000000000000000000000000712d30c0975386b043e09224433f400ea87cab7d"
],
"data": "0x000000000000000000000000000000000000000000000fccf103db7d22e00000",
"blockNumber": "0x548f4f1",
"transactionHash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"transactionIndex": "0x0",
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"logIndex": "0x2",
"removed": false
},
{
"address": "0x3fb46c4db76d8e9f69f3f8388f43a7ca7e140807",
"topics": [
"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925",
"0x000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d2",
"0x000000000000000000000000ecf4ea7907e779b8a7d0f90cb95fe06f43b610fb"
],
"data": "0x000000000000000000000000000000000000314dc3a0f41e5c923c6637280000",
"blockNumber": "0x548f4f1",
"transactionHash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"transactionIndex": "0x0",
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"logIndex": "0x3",
"removed": false
},
{
"address": "0x712d30c0975386b043e09224433f400ea87cab7d",
"topics": [
"0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67",
"0x000000000000000000000000ecf4ea7907e779b8a7d0f90cb95fe06f43b610fb",
"0x000000000000000000000000d844d373f4ef676665fcc18d9b6c89464a36f2d2"
],
"data": "0x000000000000000000000000000000000000000000000fccf103db7d22e00000ffffffffffffffffffffffffffffffffffffffffffffff4a6c76a07a43c4158f0000000000000000000000000000000000000000363decaa6e4099f6e686ca910000000000000000000000000000000000000000220948b5cb625b7056562c32ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff86c3",
"blockNumber": "0x548f4f1",
"transactionHash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"transactionIndex": "0x0",
"blockHash": "0xb6fbeabaa5682445b825c5bb02faf9290a38be44d9a47834b65224478923ebce",
"logIndex": "0x4",
"removed": false
}
],
"logsBloom": "0x00000020000000002000000000000000000000000000000000000000000000000080000008020000000000800200000000000000000420000000000000200000000000000000000800000008000008010000000000000000000000000000080000000000000000100000000000000000000000000000000000000010000800000000001000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000480000000040000000002000000000004000000000000000000000040000000000000000000000030000000000000000000000000000000000000000000040000000000000010",
"status": "0x1",
"to": "0xecf4ea7907e779b8a7d0f90cb95fe06f43b610fb",
"transactionHash": "0xbf83342ccdd6592eff8e2acfed87e23e852d684a4e2cfade89ba3b304c2b66a9",
"transactionIndex": "0x0",
"type": "0x0"
}
}
eth_getTransactionReceipt executed successfully.

---

Executing eth_getUncleByBlockHashAndIndex...
{
"id": 5002,
"jsonrpc": "2.0",
"result": null
}
eth_getUncleByBlockHashAndIndex executed successfully.

---

Executing eth_getUncleByBlockNumberAndIndex...
{
"id": 5002,
"jsonrpc": "2.0",
"result": null
}
eth_getUncleByBlockNumberAndIndex executed successfully.

---

Executing eth_getUncleCountByBlockHash...
{
"id": 5002,
"jsonrpc": "2.0",
"result": "0x0"
}
eth_getUncleCountByBlockHash executed successfully.

---

Executing eth_getUncleCountByBlockNumber...
{
"id": 5002,
"jsonrpc": "2.0",
"result": "0x0"
}
eth_getUncleCountByBlockNumber executed successfully.

---

Executing eth_getWork...
{
"id": 5002,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled, reason: restricted by blockchain schema"
}
}
eth_getWork executed successfully.

---

Executing eth_hashrate...
{
"id": 5002,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled, reason: restricted by blockchain schema"
}
}
eth_hashrate executed successfully.

---

Executing eth_maxPriorityFeePerGas...
{
"id": 1002,
"jsonrpc": "2.0",
"error": {
"code": -32601,
"message": "the method eth_maxPriorityFeePerGas does not exist/is not available",
"data": {
"trace_id": "aba84f57878bda50889cd7240821ce6d"
}
}
}
eth_maxPriorityFeePerGas executed successfully.

---

Executing eth_mining...
{
"id": 5002,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled, reason: restricted by blockchain schema"
}
}
eth_mining executed successfully.

---

Executing eth_pendingTransactions...
{
"id": 1004,
"jsonrpc": "2.0",
"result": []
}
eth_pendingTransactions executed successfully.

---

Executing eth_protocolVersion...
{
"id": 1004,
"jsonrpc": "2.0",
"result": "0x64"
}
eth_protocolVersion executed successfully.

---

Executing eth_resend...
{
"id": null,
"jsonrpc": "2.0",
"error": {
"code": -32700,
"message": "Failed to parse request",
"data": {
"trace_id": "db0083822037fda2dd2bfb455232c6fd"
}
}
}
eth_resend executed successfully.

---

Executing eth_sendRawTransaction...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32000,
"message": "rlp: element is larger than containing list",
"data": {
"trace_id": "07028f1962179944d0a32ab4f908f08f"
}
}
}
eth_sendRawTransaction executed successfully.

---

Executing eth_sendTransaction...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32000,
"message": "unknown account",
"data": {
"trace_id": "ce1f33469bc8b4100853db974edf281e"
}
}
}
eth_sendTransaction executed successfully.

---

Executing eth_sign...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32000,
"message": "unknown account",
"data": {
"trace_id": "213059d7c7a34a91d9849cab1c9549bf"
}
}
}
eth_sign executed successfully.

---

Executing eth_signTransaction...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32000,
"message": "unknown account",
"data": {
"trace_id": "8717866861f1afd3960412c8de16285e"
}
}
}
eth_signTransaction executed successfully.

---

Executing eth_submitWork...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled, reason: restricted by blockchain schema"
}
}
eth_submitWork executed successfully.

---

Executing eth_syncing...
{
"id": 1001,
"jsonrpc": "2.0",
"result": false
}
eth_syncing executed successfully.

---

--- Filter Methods ---
Executing eth_getFilterChanges...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}
}
eth_getFilterChanges executed successfully.

---

Executing eth_getFilterLogs...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}
}
eth_getFilterLogs executed successfully.

---

Executing eth_newBlockFilter...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}
}
eth_newBlockFilter executed successfully.

---

Executing eth_newFilter...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}
}
eth_newFilter executed successfully.

---

Executing eth_newPendingTransactionFilter...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}
}
eth_newPendingTransactionFilter executed successfully.

---

Executing eth_uninstallFilter...
{
"id": 1001,
"jsonrpc": "2.0",
"error": {
"code": -32075,
"message": "Method disabled"
}
}
eth_uninstallFilter executed successfully.

---

# All RPC calls attempted.
