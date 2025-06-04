## module XDPoS

### XDPoS_getBlockInfoByEpochNum

Parameters:

- epochNumber: integer, required, epoch number

Returns:

result: object EpochNumInfo:

- hash: hash of first block in this epoch
- round: round of epoch
- firstBlock: number of first block in this epoch
- lastBlock: number of last block in this epoch

Example:

```shell
epoch=89300

curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getBlockInfoByEpochNum",
  "params": [
    '"${epoch}"'
  ]
}' | jq
```

Response:

```shell
{
  "id": 1,
  "jsonrpc": "2.0",
  "result": {
    "hash": "0x5a701a8ba642a9b53475bb19cb9a313829f7afb4287caa76bebaea02f0219f89",
    "round": 1,
    "firstBlock": 80370001,
    "lastBlock": 80370838
  }
}
```

### XDPoS_getEpochNumbersBetween

Parameters:

- begin: string, required, block number
- end: string, required, block number

Returns:

result: array of uint64

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getEpochNumbersBetween",
  "params": [
    "0x5439860",
    "0x5439c48"
  ]
}' | jq
```

Response:

```shell
{
  "id": null,
  "jsonrpc": "2.0",
  "error": {
    "code": -32700,
    "message": "Failed to parse request",
    "data": {
      "trace_id": "3058fd0f3dca6cca3f7a7c8ff270ebe1"
    }
  }
}
```

### XDPoS_getLatestPoolStatus

The `XDPoS_getLatestPoolStatus` method retrieves current vote pool and timeout pool content and missing messages.

Parameters:

None

Returns:

result: object MessageStatus

- vote: object
- timeout: object

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getLatestPoolStatus"
}' | jq
```

Response:

```shell
TODO:
```

### XDPoS_getMasternodesByNumber

Parameters:

- number: string, required, BlockNumber

Returns:

result: object MasternodesStatus:

- Number: uint64
- Round: uint64
- MasternodesLen: int
- Masternodes: array of address
- PenaltyLen: int
- Penalty: array of address
- StandbynodesLen: int
- Standbynodes: array of address
- Error: string

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getMasternodesByNumber",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "Number": 88958955,
    "Round": 8705228,
    "MasternodesLen": 108,
    "Masternodes": [
      "0x047ffe1fc7f6d0b7168c4ccc312221089629f470",
      "0xb1693d224d4fd70ad6c9c9bdc44faa3ce21fb40d",
      "0x74d3ac0efc4c22ea05150d7501c95c610b130c4b",
      "0x49d3fa92eec838f644e8cacf2c93533d29b6c713",
      "0x90c87c9ff588e9268a1c7c79a5986886e98c2f04",
      "0x2a591f3d64f3ce6b1d2afeead839ad76aab9feb2",
      "0x5651290bdd3a952357066b324f9334b544100140",
      "0x65c90c2d3e99d8366f90db7f4d25f23a0a569d49",
      "0x595b8170eaf2e53e47cc20db47ad063a1e6e2c0f",
      "0x0b0afb41c14dd921a9a4511c226dc29254ba0d8f",
      "0xb9a3a97f6a02a86483bf02cb33f8b33d2d117708",
      "0x25c65b4b379ac37cf78357c4915f73677022eaff",
      "0xc67c2dec79da735d6587d8db3c23271d557196ab",
      "0x241ff0d3096e2e0b477780f9f551918a06827c05",
      "0x1d393817ab218ebade3e8ad591593ec3b132b1f6",
      "0x4b7076c988da8a0ef87f1af137f7abc39557b746",
      "0x619f838ea2a12cdb508e759c3e0697e021d52ceb",
      "0xf04f32c46f1a16663bc7f6409b35b33f3daa9a03",
      "0x8f2fb5da850042b7da5097061f098493f8ec6dd8",
      "0xd22fdac1459760f698618d927bbe22249e2b29b9",
      "0xa65010026b83368ca05df6e8b467985d6de3eac5",
      "0xe230905c99aaa7b68402af8611b89ceda743191e",
      "0xdccd99da1c942a36c5dcaf26d19b98c815dfcb73",
      "0x2f86476faa31c3f5a3d5b9376282e1b02b4dfa00",
      "0x03d9e17ae3ff2c6712e44e25b09ac5ee91f6c9ff",
      "0x9a3787688fd210ec8f8d0224c6c50b8178d75bc0",
      "0xe4710a854b24062ff37ab6636bd9a456e24c1635",
      "0x63443ffdb5c139c3bbab97d3b06ee0674d75ab7f",
      "0x1aff171caa8c1ea93bac4c27676f356d054412b4",
      "0x1e0103ba7665d15328b3886927d4f0a85f8b2299",
      "0x8e455bb521212bdcf8cb326e32dc1183cb3fd887",
      "0x70916660766043c23443479658008eeedad8cb94",
      "0x9043fafbe8833adafb449b559cef300a9632763f",
      "0x4398241671b3dd484fe3213a4fb7511f30e7d7c0",
      "0xe865a5b2bf699a4b498de8a8c55da14bb0d94b21",
      "0xb38aba47a5563fa4aedc2a649ef819b9fb160b6a",
      "0xa70b87c39a1237ba53953a1da04b23d4db28eace",
      "0x39ee018a85e4469a1b335efe014067e315bd6594",
      "0xa72ce94a09db26dce57a4852409abc2fff07a962",
      "0xdb2e141595d8edf6b1ce40b10d57ca5b83855ec5",
      "0x9fce52c5c451599235c17bc37e99f846d25ee6ec",
      "0xd2fefdd118aa2bea5991dc079fd341d3e7a92caf",
      "0x9c54bc764c2e00717fbbf31b4dc962a72c47bdf6",
      "0x72a0224245992792c75565bd3ab6caa6c8b95218",
      "0xa4657c02208797985adedcbd048efc82291dbdb2",
      "0xc7d49d0a2cf198deebd6ce581af465944ec8b2bb",
      "0x1a7c3d7481adc7e067f2f73f35f44143d7bf02d4",
      "0x4e111142fbf2edacc4ab122feb54b031dc3d95d9",
      "0x24c0832d9df8189166d8e4e2cedc40ad69077b17",
      "0x360b9e5870fc91e375f902fc134f91739f4a99c7",
      "0x2d74d0125982bdc3a9f54a10216d82509379e821",
      "0x609cc32e7513135349fbe9ebd78ea8cdcfd9f819",
      "0xe494fe5a38b7212ab999152d1d45adab8d84a8a7",
      "0xccab2f5267bc66b69f1688560c2e075216908a2c",
      "0xbbd2d417a8b6f1b1d7a267cd1d7402b443f35cfe",
      "0xc02aed857b01b4d60b378096221db3b60afadc37",
      "0x664c4a7b15d91b07c468162f535909114c038b91",
      "0x72fb467ef6da285b6b0d8f7a25abd6049385d5fd",
      "0xe2bc9a03d5ae35e130c8bf99ea50c022de375db7",
      "0xd5bb37763625bc16f2ad0791d87e2726004241e4",
      "0xdef4bd67f9e2627ca495b5f71794fab1bab40619",
      "0x3056a8bff9a17b12d08f1837d0b44cf1e2018fbe",
      "0x065551f0dcac6f00cae11192d462db709be3758c",
      "0xec4a5fbd2e46e97be98c212e668ef1d08d695440",
      "0x7525f63e5f739ef952fbea50c1dc97ef5d5bd984",
      "0x7aa125338be075260e77c6a66a56c90a5dec4c58",
      "0x450d714e65f4de007937c56fb1c4686efd4fb4ec",
      "0x497c44f4a22099109ad7d194aad4fbe78dceb788",
      "0x35898873c021f9f4fa01a76c7fa595159f5523e7",
      "0xcfccdea1006a5cfa7d9484b5b293b46964c265c0",
      "0x92d32364313c376dcae272e113fd04cea194d2c8",
      "0xe8e8194c0a8e5a32306c7a86052a4e8ca4b8729d",
      "0x1a1a8229e4a2fa06f73a564092976ba2dfe6aab8",
      "0x19a1e02fec5d52be689f59b22d9c0f1765bd7052",
      "0xfce02e190f722e790be06b3a6ec43e2f606dba65",
      "0x051572edba4a8290bf45dd2d7ab80eb6cd56ccc6",
      "0x1bc85db77617515381cbbb28fdd2b1a10264de11",
      "0x346ada489c70c85ca665428389f1a971abbee960",
      "0xf38f7402de55c5c7cf7a5a29244148b817d7bca3",
      "0x78fb8980d122c902ad1abeac523a92e9c24e75c4",
      "0x06fb5e6a7abcc241729f591ac96ee41a9f10decf",
      "0x64e2e07bb17d978108824219741b9cc9e52b4443",
      "0x6e5bbbf0d9d1851599a21676c99b1a8cc012753d",
      "0xeb0f1565b5d3bbf4489deb67972b814ee84c6cb7",
      "0x361faefbcd1c1b7f2d18723d178633202ccc5858",
      "0x0407b8f788fbd763280212be6f7a672351009487",
      "0xcc8247cfaed1950dfd6dd4bc23a0ecab014c3b6c",
      "0xeb76fb646c5ab9db8543719e74adadf101cd0616",
      "0x449bed0b27afe708a30a481329733bcbcba2bc21",
      "0x5a39debded5481a5e5b73a1551fd31208b88da4e",
      "0x4bcdd234c20da32ec34a7ff85d6c1e6c832604d4",
      "0x070eed0a2c77650dd41599920b5a0089dc60bfe6",
      "0x7c0f72d80d9d8b822e2fb217e56be732a12fc8c9",
      "0xc428369353c7de2782370ee634be1516d836ddad",
      "0x7b03e5e7912b0ce65f8be7549e23cb08392e369c",
      "0x29a66147543707ef83d5d15e0a1d9f65cbe90e7c",
      "0x0278c350152e15fa6ffc712a5a73d704ce73e2e1",
      "0x3bddd35ddf0945c0151a4ccd547d9e7a0b1fa4cd",
      "0x67a24b5821724419bbc0710cbbb3122d84703318",
      "0x7bda7752b3533b3b71346ff02f6314b8fe77ec43",
      "0x4466ca98bdcd08d54a9672d5ca76a69909c4b59e",
      "0xfcbddd58be4616042543637a84ad18cb065adba4",
      "0x0afe5a6890cf357022dfdaac7e2d6e248755cdfd",
      "0x8b9654195b7c9a56ae20deb36a29834763800b03",
      "0x3cc8e30f8be48e57dbac01d8c7b05fb18ef1f7c8",
      "0xd4ff534569c5bfd12d2280b8bd27aa8876f6b5fa",
      "0x555d4cd195bcb7d4196f2005db46a4b71c36c5a4",
      "0xfc5fea016be9ac2d23dfe18e16d3be79547e5466"
    ],
    "PenaltyLen": 10,
    "Penalty": [
      "0xaf3db8ad8154820303adc6be0d9b10ef8e33024f",
      "0x00633d6bc3c30349a7fe55999c5b18940bcc9f43",
      "0x55b7d6ca73944e2ab43d226d7ed1ff3d72516aff",
      "0x418c9839222e7105d227281f2a69a6c9d19c297e",
      "0xd9d4f525683b2b54508727ae44ac147e4f0db441",
      "0x95e58330cc798c8f079b8c91ba4d960d114cd785",
      "0x69a809f4d6e65792158c1bef35b316f6292eefb1",
      "0xb1847452a58a2e2a4ee658a192f0d5d511c6f3f2",
      "0xd9d474b83c514920f44c519f83e23e7b6927d552",
      "0x0b5ba833d689a2ec8a3f5275f2b033bdd2de337b"
    ],
    "StandbynodesLen": 136,
    "Standbynodes": [
      "0x48378844ab9486410ebf01d058a434e6a42e2ca5",
      "0x4e52755f091c4fb826b994666a3e207fdf3a0cc5",
      "0xa45457cc0153092a7e717e8e6d41221b827c28ae",
      "0xa478633e8aba4188c25159e1103206eeca98bced",
      "0x4cfd9048cfe93e02e30a84c7538c726e980c1d78",
      "0x1390a610e22f9ee1c0fc92c341a2b0c6f3f9d08a",
      "0xbb0830e2d006d557541935362a621d002b0d6b00",
      "0xac95f4af1e9efbc914895695e232fb8dbaa66275",
      "0xab5d74a461bb0f2231abcc5341ba7d89b24f0926",
      "0xef7ce03b7fdea2d0bf2ad6e0a1e9b0f2ed547447",
      "0x76e9285b4a804033a1d87b21c4e9d4e7028ce944",
      "0xba96294248d34b90f5b893ef82fd736cc1cc00ad",
      "0x61ddae7aefa4d69c868379416504b1537d1fd558",
      "0x75ea95b0e9a56417ce1434aac37b0afcba0c5221",
      "0x74f297c0197bab5ecf21add9ad9ced96ad42b61f",
      "0xa9c2b73ba170ddf37bc92fee1bafaf54e138ce87",
      "0xde1470c0f478ca7efd151be5a13c2addcfff328c",
      "0x83e12760c1e324b16fe68bc23ffc69b37d17ee23",
      "0xe02ca8d829b3e0185d168c62aeef3960e6c0882e",
      "0x9bbf9e3af8a473124a17656f0e84d99aa8b51864",
      "0x97c1117aec1b18a6b70a103eff30d2c55adf7bc0",
      "0x3dd28884ec3ea0904f0f6697c31b7d273a899480",
      "0xb594c78a71763d543d10358b2dde579ff041319a",
      "0x9682520376f26a733a75588c0bdede7645bdfc4a",
      "0x51aa6a6b6a738d44e6a4e00d6e02faa991cf44be",
      "0x1cad85be8d097ed7a9b74ef5fbc9c75f2e2aa3bb",
      "0x00021c85d8cf9ef6c73e8f95740f8e874e4c9d1f",
      "0x251aca934009f4dda1d904e28ad7f20b3da4d888",
      "0xf0dee298ac62b7e351148b4fbde1d45d7d9b1560",
      "0x555ff7828a0b435432308a4721e4d9f110324025",
      "0xa59318ba345b293f4ac5acb9f59a47365987b7da",
      "0x01df5f4986976e8cf1592cca7b792664510b6086",
      "0xd438128b57cf71ab2168a263f24708263775c54f",
      "0x935a34962ed8237d9700191320e3e0b66741133d",
      "0x84620482d52b71cb18bf84cd34f0cc0dcfd940cd",
      "0x783f0724f9dabd1bb52c7331bce7041804221d1b",
      "0xcd52996e7a6c75ddf85baab529af6c55f254ac15",
      "0x82183d37848f38c807e274b15b86d16e9ae60fd3",
      "0xe8b4fd78437831cb9a182771bef485b619cef31f",
      "0x200f1c0b9411dd0f1fad9a76cc47cf8d56e0a19d",
      "0xf11e012f56ac7341be939cb668e63a37e79db84e",
      "0xe495a6242b78fe35211116839ad1882e7c4fc564",
      "0x4eaff58c4a66839c6bbe2bca1dbd52a4d328aa5c",
      "0x7b94fcbd37f865892874e315c954062f936460fc",
      "0x3396803b0cb3b5790911bb51b6c5193e9940793f",
      "0x2d881b373edf297b495e8deedd9108e3d6870dee",
      "0xf48aa29de23c45c54a96c2a7b8ce5218abf46874",
      "0x844a6c40b34e2335fc6887db2975d55183082c0d",
      "0x7b6b545279bc9e262a9dc224ab724ad629c75b63",
      "0x44d21022e104d605235c863cf86c96045fad9505",
      "0xd5384163e03481b8b1c13e8da611e05e2612586c",
      "0xcac04ac16876a65d3bf797679f34fc8272a1f6b4",
      "0x417b06236ecc5bf200a59152e92cd4923867409c",
      "0x09bae325e54edf4dc8b79a8bb29f90deb90e8fe7",
      "0x82c85801a8b18c4b2701e552727cea8494bbe9ea",
      "0x5454edee66858dfcc14871cc8b26f57ef528bedc",
      "0xd6e9161fad50d2e697cda2b960fd26bf52a5d169",
      "0x8eeeeafaca49e507bc9bba121dac97ec75774ff2",
      "0x621996eb1d03a6e9614ddf150f6123c5602c8418",
      "0x391f3d21d9062ebbdd958e436b26d4b40928658c",
      "0x4288545139bb238d1671b96e8f8016b0b3793ad4",
      "0x6f356f6a7ff84c46f59fee691a2a161494871eec",
      "0xd97d7434af4ad873585c71df55fce07996e27a82",
      "0x7fca0e918c06ff0b4e51ecf76670e07714e44a9c",
      "0x9200e4fe8959f0268eb9daf06aba4756e595b99f",
      "0x2e32b4db98ae68ea725212040f5f29abc61ce18f",
      "0xa6d916b338291ae14402d67af2b9d9e19bf72fd5",
      "0x9e057cebf2101a5e3e57906b834a2a82893b7739",
      "0x58f25975086b8d23c028ebb651be064da0360367",
      "0x5f4ff7948061a605e9d64b4d5fde927b70bc3fa5",
      "0xb2610a596e2e34cf25adda2f05d867b3dc96b47d",
      "0xb95d3439a7d3654e7871e72213caaad472ec8325",
      "0x785076b971f2a90b029b680f90d4d609060c01d5",
      "0x43bb2e3bb47d0a12d45c87dcdbd6d16a47e1000a",
      "0x944dd93e0f6d213b0dce27590bd5a9a4351fafd6",
      "0x349563f9b0e3167f816cc39b272e3f888e23d3f3",
      "0xc7f52a1c275d470ff92718dfe27be7800f7bb697",
      "0xa43a99c9752166ec30c8369ce1e84de43be1e617",
      "0x3197b8ed3bcd7eea2c695951d9d57e833457072b",
      "0x7af6aedf3382bb3de2dba61769d96444bf660494",
      "0xf2a1d535f4bda9e810c067975d08dd8724f9066a",
      "0x934a05866a7a44b212b55484f201e0ada955ef41",
      "0xcb67c6c9ea422e33feef44aa5564bdf89997f57b",
      "0x58d877f1f8a61cf99192e3161429c8a78d7f1243",
      "0xebf8e4904fe2bd561475e9e0e65a3336c56f83d3",
      "0xa828c2381bb91fc96cc269de73b99a4314ac4b5b",
      "0x39ffd4aa42b02f1b55ed822267da57219f11fc6a",
      "0x2759e5176992c834e7c0da762ae6178de6e7e711",
      "0x966f1f7e10f9bd531e31ab6bba0936944e00ebf9",
      "0x05f4c54ffc16b943b98713248c751ce98f924cbd",
      "0xfa9a1f2ec3d92428fc6105a4e05832ae74170ac4",
      "0x695ce237df312f9a6c6a5386b91b37ae249d07dd",
      "0x226ddc05e1adc6872285d95d4277a498fdce375b",
      "0xe8bc0e78114ad8dc4290eaed78bcb9e3f5b91e32",
      "0xd64c15787bb52bbfe0d195df0d954725469a318e",
      "0x249f443d5c8400d7bc638a199d1f458b47a41a47",
      "0x21396ca79a433a20b5904c9a77aa9a4f110ef80e",
      "0x079d87904cd707143dd39c8be595bbe7a2181fb4",
      "0xb8ad6e0a7bffb5ea0b24c095e16bb4a2c936203c",
      "0x9d38187eccc49419c99af4b99dbb1494592b8a56",
      "0xb653aa2df1b710894bc1b94b3bd06799701d6d6b",
      "0xb41517b07efc01226332421c3fbf4e56807df5f1",
      "0x8e9976680f84fab8b1dff91716762855c1a975e1",
      "0xbd46f94ca626a3ac72e60ca7efde379e42603b98",
      "0x66bb8c2724ed2dc09ece6024045e90e65c29eaff",
      "0xe79ef2d92a8e4d758d78dfcc8e61b2fa377ee78d",
      "0x0e26a54bf70c51460cb13f4823027678672e4a0a",
      "0x618e706c531e57f44a739cef129bf08cbd912d6a",
      "0x7c1f132c335f04002178d0d29fe785ad329bdfb1",
      "0x2a4559e84ca53e00b05d9d53086005432926a992",
      "0x08d53a08b1d46ecb45bc3a299af9abdde9035338",
      "0xc7bf1cd426e6f47d879fbaba789d7c5740ce9b4b",
      "0xf7ff86c0d2c2795dbd2731ca122a7d13a802b7b8",
      "0xe7a5556c87fe764dcba9bcd31e1f4aa27b638117",
      "0x32e494f6f737bd419169a837c35e195a851bcb8d",
      "0xb55f4ca1a5741a91ff4ea4e15e9324eb973f4ccd",
      "0x26179a252e9de28c29a37fbeb60df72be5e3e27d",
      "0xb2e8a55ea7fccdd8cdfb1af8e497a6f793ddf63e",
      "0xec5273989dd0e18d658f037cae2eff6803369c5c",
      "0xc6b903af2816a2423a3300ed853e5e7d26ea0cf0",
      "0x1eca7cd6df02a177e10bc9769a544b49f13aa875",
      "0x19e945a26b15e4d6f96386a4f8596b6313311b7e",
      "0xfab275bfc83c3bc8aa17c16ed5a0b4e60caca299",
      "0x3b8a906602422fe9f316e6fa0c1728cfaa238766",
      "0x7b0704722600979fd7da53afb1618aab3b9ddeb4",
      "0x7f82ddbd91c5fafe6d5fdf417c681eed211fefec",
      "0xb570b13fa9ca093347448dc600ed0d09333ac9fa",
      "0xfc2f29a87a63e07a341d0bb9462e9a14af013a6c",
      "0x76df51309c6367678635a3ea76d1482666f55a05",
      "0xfe5b2f3e7fb6fc3b8c65f2af4d3f4cabe024bce9",
      "0x13afe4210b6340723883be07618d3a4f019c2adc",
      "0x1c908c36c4470c3f66860e7c5152cc17e480ff54",
      "0x823a47d0a1f4c5781ead9f0d5b361557a45355bd",
      "0x52c21a9daa75e6eb91d449cad9649013903b7a73",
      "0x5fca7280f7b3f743bc3248261d05f7b57651ec3c",
      "0xe05e3146327d8fb775f50fd8d8f943dbe200ef4b"
    ],
    "Error": null
  }
}
```

### XDPoS_getMissedRoundsInEpochByBlockNum

Parameters:

- number: string, required, BlockNumber

Returns:

result: object PublicApiMissedRoundsMetadata:

- EpochRound: uint64
- EpochBlockNumber: big.Int
- MissedRounds: array of MissedRoundInfo

MissedRoundInfo:

- Round: uint64
- Miner: address
- CurrentBlockHash: hash
- CurrentBlockNum: big.Int
- ParentBlockHash: hash
- ParentBlockNum: big.Int

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getMissedRoundsInEpochByBlockNum",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "EpochRound": 8704800,
    "EpochBlockNumber": 88958531,
    "MissedRounds": [
      {
        "Round": 8705243,
        "Miner": "0x25c65b4b379ac37cf78357c4915f73677022eaff",
        "CurrentBlockHash": "0x3e6cfe3cc93ad1abd83a3161e2d231423bdeb53c9c39bad7d69e56354e8bbcfb",
        "CurrentBlockNum": 88958970,
        "ParentBlockHash": "0x0122d58cbaff45c6612e2c3a9133286afa036ba4aca19533e16eb3b3fa7bac94",
        "ParentBlockNum": 88958969
      },
      {
        "Round": 8705135,
        "Miner": "0x25c65b4b379ac37cf78357c4915f73677022eaff",
        "CurrentBlockHash": "0x76351d68f818219bc0062dd4e6bc53a7906a83c16f0e835116418172aa92151f",
        "CurrentBlockNum": 88958863,
        "ParentBlockHash": "0xa0f7cbdaf1a740f181ff752dc3555c709d6fb3b4a4c566eb5765a9df3a77e9de",
        "ParentBlockNum": 88958862
      },
      {
        "Round": 8705027,
        "Miner": "0x25c65b4b379ac37cf78357c4915f73677022eaff",
        "CurrentBlockHash": "0xb024abd82163cba895e85acf4f036919b44fa24a51e77afd5ef5d1bdc674c376",
        "CurrentBlockNum": 88958756,
        "ParentBlockHash": "0xb093b1566da1c150b41d2e4ba688f2b7cb3a878acb03479820086d86bc3f4c83",
        "ParentBlockNum": 88958755
      },
      {
        "Round": 8704883,
        "Miner": "0xeb0f1565b5d3bbf4489deb67972b814ee84c6cb7",
        "CurrentBlockHash": "0x6af55200da7d7ae77a065c5d6c28454ea7b893718bf392da70b110f10e23d813",
        "CurrentBlockNum": 88958613,
        "ParentBlockHash": "0x6e716e6b375c27ad35b818dd094bf6a15d5af972ef414bc0bab70cc674a34e28",
        "ParentBlockNum": 88958612
      },
      {
        "Round": 8704811,
        "Miner": "0x25c65b4b379ac37cf78357c4915f73677022eaff",
        "CurrentBlockHash": "0x91546fa6ad53dce3ba69d2d453ae6e7b630f016ea461afccd90bc18e4d09aad0",
        "CurrentBlockNum": 88958542,
        "ParentBlockHash": "0xeaff38e3c6160e36ce8ed1092efbcab820f3aafaa674b847f601a6e2003f67a7",
        "ParentBlockNum": 88958541
      }
    ]
  }
}
```

### XDPoS_getSigners

The `getSigners` method retrieves the list of authorized signers at the specified block.

Parameters:

- number: string, required, BlockNumber

Returns:

result: array of address

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getSigners",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": [
    "0x047ffe1fc7f6d0b7168c4ccc312221089629f470",
    "0xb1693d224d4fd70ad6c9c9bdc44faa3ce21fb40d",
    "0x74d3ac0efc4c22ea05150d7501c95c610b130c4b",
    "0x49d3fa92eec838f644e8cacf2c93533d29b6c713",
    "0x90c87c9ff588e9268a1c7c79a5986886e98c2f04",
    "0x2a591f3d64f3ce6b1d2afeead839ad76aab9feb2",
    "0x5651290bdd3a952357066b324f9334b544100140",
    "0x65c90c2d3e99d8366f90db7f4d25f23a0a569d49",
    "0x595b8170eaf2e53e47cc20db47ad063a1e6e2c0f",
    "0x0b0afb41c14dd921a9a4511c226dc29254ba0d8f",
    "0xb9a3a97f6a02a86483bf02cb33f8b33d2d117708",
    "0x25c65b4b379ac37cf78357c4915f73677022eaff",
    "0xc67c2dec79da735d6587d8db3c23271d557196ab",
    "0x241ff0d3096e2e0b477780f9f551918a06827c05",
    "0x1d393817ab218ebade3e8ad591593ec3b132b1f6",
    "0x4b7076c988da8a0ef87f1af137f7abc39557b746",
    "0x619f838ea2a12cdb508e759c3e0697e021d52ceb",
    "0xf04f32c46f1a16663bc7f6409b35b33f3daa9a03",
    "0x8f2fb5da850042b7da5097061f098493f8ec6dd8",
    "0xd22fdac1459760f698618d927bbe22249e2b29b9",
    "0xa65010026b83368ca05df6e8b467985d6de3eac5",
    "0xe230905c99aaa7b68402af8611b89ceda743191e",
    "0xdccd99da1c942a36c5dcaf26d19b98c815dfcb73",
    "0x2f86476faa31c3f5a3d5b9376282e1b02b4dfa00",
    "0x03d9e17ae3ff2c6712e44e25b09ac5ee91f6c9ff",
    "0x9a3787688fd210ec8f8d0224c6c50b8178d75bc0",
    "0xe4710a854b24062ff37ab6636bd9a456e24c1635",
    "0x63443ffdb5c139c3bbab97d3b06ee0674d75ab7f",
    "0x1aff171caa8c1ea93bac4c27676f356d054412b4",
    "0x1e0103ba7665d15328b3886927d4f0a85f8b2299",
    "0x8e455bb521212bdcf8cb326e32dc1183cb3fd887",
    "0x70916660766043c23443479658008eeedad8cb94",
    "0x9043fafbe8833adafb449b559cef300a9632763f",
    "0x4398241671b3dd484fe3213a4fb7511f30e7d7c0",
    "0xe865a5b2bf699a4b498de8a8c55da14bb0d94b21",
    "0xb38aba47a5563fa4aedc2a649ef819b9fb160b6a",
    "0xa70b87c39a1237ba53953a1da04b23d4db28eace",
    "0x39ee018a85e4469a1b335efe014067e315bd6594",
    "0xa72ce94a09db26dce57a4852409abc2fff07a962",
    "0xdb2e141595d8edf6b1ce40b10d57ca5b83855ec5",
    "0x9fce52c5c451599235c17bc37e99f846d25ee6ec",
    "0xd2fefdd118aa2bea5991dc079fd341d3e7a92caf",
    "0x9c54bc764c2e00717fbbf31b4dc962a72c47bdf6",
    "0x72a0224245992792c75565bd3ab6caa6c8b95218",
    "0xa4657c02208797985adedcbd048efc82291dbdb2",
    "0xc7d49d0a2cf198deebd6ce581af465944ec8b2bb",
    "0xaf3db8ad8154820303adc6be0d9b10ef8e33024f",
    "0x1a7c3d7481adc7e067f2f73f35f44143d7bf02d4",
    "0x4e111142fbf2edacc4ab122feb54b031dc3d95d9",
    "0x24c0832d9df8189166d8e4e2cedc40ad69077b17",
    "0x360b9e5870fc91e375f902fc134f91739f4a99c7",
    "0x2d74d0125982bdc3a9f54a10216d82509379e821",
    "0x609cc32e7513135349fbe9ebd78ea8cdcfd9f819",
    "0xe494fe5a38b7212ab999152d1d45adab8d84a8a7",
    "0xccab2f5267bc66b69f1688560c2e075216908a2c",
    "0xbbd2d417a8b6f1b1d7a267cd1d7402b443f35cfe",
    "0xc02aed857b01b4d60b378096221db3b60afadc37",
    "0x664c4a7b15d91b07c468162f535909114c038b91",
    "0x72fb467ef6da285b6b0d8f7a25abd6049385d5fd",
    "0xe2bc9a03d5ae35e130c8bf99ea50c022de375db7",
    "0xd5bb37763625bc16f2ad0791d87e2726004241e4",
    "0xdef4bd67f9e2627ca495b5f71794fab1bab40619",
    "0x3056a8bff9a17b12d08f1837d0b44cf1e2018fbe",
    "0x065551f0dcac6f00cae11192d462db709be3758c",
    "0xec4a5fbd2e46e97be98c212e668ef1d08d695440",
    "0x7525f63e5f739ef952fbea50c1dc97ef5d5bd984",
    "0x7aa125338be075260e77c6a66a56c90a5dec4c58",
    "0x450d714e65f4de007937c56fb1c4686efd4fb4ec",
    "0x497c44f4a22099109ad7d194aad4fbe78dceb788",
    "0x35898873c021f9f4fa01a76c7fa595159f5523e7",
    "0xcfccdea1006a5cfa7d9484b5b293b46964c265c0",
    "0x92d32364313c376dcae272e113fd04cea194d2c8",
    "0xe8e8194c0a8e5a32306c7a86052a4e8ca4b8729d",
    "0x1a1a8229e4a2fa06f73a564092976ba2dfe6aab8",
    "0x19a1e02fec5d52be689f59b22d9c0f1765bd7052",
    "0xfce02e190f722e790be06b3a6ec43e2f606dba65",
    "0x051572edba4a8290bf45dd2d7ab80eb6cd56ccc6",
    "0x1bc85db77617515381cbbb28fdd2b1a10264de11",
    "0x346ada489c70c85ca665428389f1a971abbee960",
    "0xf38f7402de55c5c7cf7a5a29244148b817d7bca3",
    "0x78fb8980d122c902ad1abeac523a92e9c24e75c4",
    "0x06fb5e6a7abcc241729f591ac96ee41a9f10decf",
    "0x64e2e07bb17d978108824219741b9cc9e52b4443",
    "0x6e5bbbf0d9d1851599a21676c99b1a8cc012753d",
    "0xb1847452a58a2e2a4ee658a192f0d5d511c6f3f2",
    "0xeb0f1565b5d3bbf4489deb67972b814ee84c6cb7",
    "0x361faefbcd1c1b7f2d18723d178633202ccc5858",
    "0x0b5ba833d689a2ec8a3f5275f2b033bdd2de337b",
    "0x0407b8f788fbd763280212be6f7a672351009487",
    "0xcc8247cfaed1950dfd6dd4bc23a0ecab014c3b6c",
    "0xeb76fb646c5ab9db8543719e74adadf101cd0616",
    "0x418c9839222e7105d227281f2a69a6c9d19c297e",
    "0x55b7d6ca73944e2ab43d226d7ed1ff3d72516aff",
    "0x449bed0b27afe708a30a481329733bcbcba2bc21",
    "0x5a39debded5481a5e5b73a1551fd31208b88da4e",
    "0x4bcdd234c20da32ec34a7ff85d6c1e6c832604d4",
    "0x070eed0a2c77650dd41599920b5a0089dc60bfe6",
    "0x7c0f72d80d9d8b822e2fb217e56be732a12fc8c9",
    "0xc428369353c7de2782370ee634be1516d836ddad",
    "0x7b03e5e7912b0ce65f8be7549e23cb08392e369c",
    "0x29a66147543707ef83d5d15e0a1d9f65cbe90e7c",
    "0x0278c350152e15fa6ffc712a5a73d704ce73e2e1",
    "0x3bddd35ddf0945c0151a4ccd547d9e7a0b1fa4cd",
    "0x67a24b5821724419bbc0710cbbb3122d84703318",
    "0x7bda7752b3533b3b71346ff02f6314b8fe77ec43",
    "0x4466ca98bdcd08d54a9672d5ca76a69909c4b59e",
    "0xfcbddd58be4616042543637a84ad18cb065adba4",
    "0x0afe5a6890cf357022dfdaac7e2d6e248755cdfd",
    "0x8b9654195b7c9a56ae20deb36a29834763800b03",
    "0x3cc8e30f8be48e57dbac01d8c7b05fb18ef1f7c8",
    "0xd4ff534569c5bfd12d2280b8bd27aa8876f6b5fa",
    "0x555d4cd195bcb7d4196f2005db46a4b71c36c5a4",
    "0xfc5fea016be9ac2d23dfe18e16d3be79547e5466",
    "0x95e58330cc798c8f079b8c91ba4d960d114cd785",
    "0x48378844ab9486410ebf01d058a434e6a42e2ca5",
    "0x4e52755f091c4fb826b994666a3e207fdf3a0cc5",
    "0xa45457cc0153092a7e717e8e6d41221b827c28ae",
    "0xa478633e8aba4188c25159e1103206eeca98bced",
    "0x4cfd9048cfe93e02e30a84c7538c726e980c1d78",
    "0x1390a610e22f9ee1c0fc92c341a2b0c6f3f9d08a",
    "0xbb0830e2d006d557541935362a621d002b0d6b00",
    "0xac95f4af1e9efbc914895695e232fb8dbaa66275",
    "0xab5d74a461bb0f2231abcc5341ba7d89b24f0926",
    "0xef7ce03b7fdea2d0bf2ad6e0a1e9b0f2ed547447",
    "0x76e9285b4a804033a1d87b21c4e9d4e7028ce944",
    "0xba96294248d34b90f5b893ef82fd736cc1cc00ad",
    "0x61ddae7aefa4d69c868379416504b1537d1fd558",
    "0x75ea95b0e9a56417ce1434aac37b0afcba0c5221",
    "0x74f297c0197bab5ecf21add9ad9ced96ad42b61f",
    "0xa9c2b73ba170ddf37bc92fee1bafaf54e138ce87",
    "0xde1470c0f478ca7efd151be5a13c2addcfff328c",
    "0xd9d474b83c514920f44c519f83e23e7b6927d552",
    "0x83e12760c1e324b16fe68bc23ffc69b37d17ee23",
    "0xe02ca8d829b3e0185d168c62aeef3960e6c0882e",
    "0x9bbf9e3af8a473124a17656f0e84d99aa8b51864",
    "0x97c1117aec1b18a6b70a103eff30d2c55adf7bc0",
    "0x3dd28884ec3ea0904f0f6697c31b7d273a899480",
    "0xb594c78a71763d543d10358b2dde579ff041319a",
    "0xd9d4f525683b2b54508727ae44ac147e4f0db441",
    "0x9682520376f26a733a75588c0bdede7645bdfc4a",
    "0x51aa6a6b6a738d44e6a4e00d6e02faa991cf44be",
    "0x1cad85be8d097ed7a9b74ef5fbc9c75f2e2aa3bb",
    "0x00021c85d8cf9ef6c73e8f95740f8e874e4c9d1f",
    "0x251aca934009f4dda1d904e28ad7f20b3da4d888",
    "0xf0dee298ac62b7e351148b4fbde1d45d7d9b1560",
    "0x555ff7828a0b435432308a4721e4d9f110324025",
    "0xa59318ba345b293f4ac5acb9f59a47365987b7da",
    "0x01df5f4986976e8cf1592cca7b792664510b6086",
    "0xd438128b57cf71ab2168a263f24708263775c54f",
    "0x935a34962ed8237d9700191320e3e0b66741133d",
    "0x84620482d52b71cb18bf84cd34f0cc0dcfd940cd",
    "0x783f0724f9dabd1bb52c7331bce7041804221d1b",
    "0xcd52996e7a6c75ddf85baab529af6c55f254ac15",
    "0x82183d37848f38c807e274b15b86d16e9ae60fd3",
    "0x69a809f4d6e65792158c1bef35b316f6292eefb1",
    "0xe8b4fd78437831cb9a182771bef485b619cef31f",
    "0x200f1c0b9411dd0f1fad9a76cc47cf8d56e0a19d",
    "0xf11e012f56ac7341be939cb668e63a37e79db84e",
    "0xe495a6242b78fe35211116839ad1882e7c4fc564",
    "0x4eaff58c4a66839c6bbe2bca1dbd52a4d328aa5c",
    "0x7b94fcbd37f865892874e315c954062f936460fc",
    "0x3396803b0cb3b5790911bb51b6c5193e9940793f",
    "0x2d881b373edf297b495e8deedd9108e3d6870dee",
    "0xf48aa29de23c45c54a96c2a7b8ce5218abf46874",
    "0x844a6c40b34e2335fc6887db2975d55183082c0d",
    "0x7b6b545279bc9e262a9dc224ab724ad629c75b63",
    "0x44d21022e104d605235c863cf86c96045fad9505",
    "0xd5384163e03481b8b1c13e8da611e05e2612586c",
    "0xcac04ac16876a65d3bf797679f34fc8272a1f6b4",
    "0x417b06236ecc5bf200a59152e92cd4923867409c",
    "0x09bae325e54edf4dc8b79a8bb29f90deb90e8fe7",
    "0x82c85801a8b18c4b2701e552727cea8494bbe9ea",
    "0x5454edee66858dfcc14871cc8b26f57ef528bedc",
    "0xd6e9161fad50d2e697cda2b960fd26bf52a5d169",
    "0x8eeeeafaca49e507bc9bba121dac97ec75774ff2",
    "0x621996eb1d03a6e9614ddf150f6123c5602c8418",
    "0x391f3d21d9062ebbdd958e436b26d4b40928658c",
    "0x4288545139bb238d1671b96e8f8016b0b3793ad4",
    "0x6f356f6a7ff84c46f59fee691a2a161494871eec",
    "0xd97d7434af4ad873585c71df55fce07996e27a82",
    "0x7fca0e918c06ff0b4e51ecf76670e07714e44a9c",
    "0x9200e4fe8959f0268eb9daf06aba4756e595b99f",
    "0x2e32b4db98ae68ea725212040f5f29abc61ce18f",
    "0xa6d916b338291ae14402d67af2b9d9e19bf72fd5",
    "0x9e057cebf2101a5e3e57906b834a2a82893b7739",
    "0x58f25975086b8d23c028ebb651be064da0360367",
    "0x5f4ff7948061a605e9d64b4d5fde927b70bc3fa5",
    "0xb2610a596e2e34cf25adda2f05d867b3dc96b47d",
    "0xb95d3439a7d3654e7871e72213caaad472ec8325",
    "0x785076b971f2a90b029b680f90d4d609060c01d5",
    "0x43bb2e3bb47d0a12d45c87dcdbd6d16a47e1000a",
    "0x944dd93e0f6d213b0dce27590bd5a9a4351fafd6",
    "0x349563f9b0e3167f816cc39b272e3f888e23d3f3",
    "0xc7f52a1c275d470ff92718dfe27be7800f7bb697",
    "0xa43a99c9752166ec30c8369ce1e84de43be1e617",
    "0x3197b8ed3bcd7eea2c695951d9d57e833457072b",
    "0x7af6aedf3382bb3de2dba61769d96444bf660494",
    "0xf2a1d535f4bda9e810c067975d08dd8724f9066a",
    "0x934a05866a7a44b212b55484f201e0ada955ef41",
    "0xcb67c6c9ea422e33feef44aa5564bdf89997f57b",
    "0x58d877f1f8a61cf99192e3161429c8a78d7f1243",
    "0xebf8e4904fe2bd561475e9e0e65a3336c56f83d3",
    "0xa828c2381bb91fc96cc269de73b99a4314ac4b5b",
    "0x39ffd4aa42b02f1b55ed822267da57219f11fc6a",
    "0x2759e5176992c834e7c0da762ae6178de6e7e711",
    "0x966f1f7e10f9bd531e31ab6bba0936944e00ebf9",
    "0x05f4c54ffc16b943b98713248c751ce98f924cbd",
    "0xfa9a1f2ec3d92428fc6105a4e05832ae74170ac4",
    "0x695ce237df312f9a6c6a5386b91b37ae249d07dd",
    "0x226ddc05e1adc6872285d95d4277a498fdce375b",
    "0xe8bc0e78114ad8dc4290eaed78bcb9e3f5b91e32",
    "0xd64c15787bb52bbfe0d195df0d954725469a318e",
    "0x249f443d5c8400d7bc638a199d1f458b47a41a47",
    "0x21396ca79a433a20b5904c9a77aa9a4f110ef80e",
    "0x079d87904cd707143dd39c8be595bbe7a2181fb4",
    "0xb8ad6e0a7bffb5ea0b24c095e16bb4a2c936203c",
    "0x9d38187eccc49419c99af4b99dbb1494592b8a56",
    "0xb653aa2df1b710894bc1b94b3bd06799701d6d6b",
    "0xb41517b07efc01226332421c3fbf4e56807df5f1",
    "0x8e9976680f84fab8b1dff91716762855c1a975e1",
    "0xbd46f94ca626a3ac72e60ca7efde379e42603b98",
    "0x66bb8c2724ed2dc09ece6024045e90e65c29eaff",
    "0xe79ef2d92a8e4d758d78dfcc8e61b2fa377ee78d",
    "0x0e26a54bf70c51460cb13f4823027678672e4a0a",
    "0x618e706c531e57f44a739cef129bf08cbd912d6a",
    "0x7c1f132c335f04002178d0d29fe785ad329bdfb1",
    "0x2a4559e84ca53e00b05d9d53086005432926a992",
    "0x08d53a08b1d46ecb45bc3a299af9abdde9035338",
    "0xc7bf1cd426e6f47d879fbaba789d7c5740ce9b4b",
    "0xf7ff86c0d2c2795dbd2731ca122a7d13a802b7b8",
    "0xe7a5556c87fe764dcba9bcd31e1f4aa27b638117",
    "0x32e494f6f737bd419169a837c35e195a851bcb8d",
    "0xb55f4ca1a5741a91ff4ea4e15e9324eb973f4ccd",
    "0x26179a252e9de28c29a37fbeb60df72be5e3e27d",
    "0xb2e8a55ea7fccdd8cdfb1af8e497a6f793ddf63e",
    "0xec5273989dd0e18d658f037cae2eff6803369c5c",
    "0x00633d6bc3c30349a7fe55999c5b18940bcc9f43",
    "0xc6b903af2816a2423a3300ed853e5e7d26ea0cf0",
    "0x1eca7cd6df02a177e10bc9769a544b49f13aa875",
    "0x19e945a26b15e4d6f96386a4f8596b6313311b7e",
    "0xfab275bfc83c3bc8aa17c16ed5a0b4e60caca299",
    "0x3b8a906602422fe9f316e6fa0c1728cfaa238766",
    "0x7b0704722600979fd7da53afb1618aab3b9ddeb4",
    "0x7f82ddbd91c5fafe6d5fdf417c681eed211fefec",
    "0xb570b13fa9ca093347448dc600ed0d09333ac9fa",
    "0xfc2f29a87a63e07a341d0bb9462e9a14af013a6c",
    "0x76df51309c6367678635a3ea76d1482666f55a05",
    "0xfe5b2f3e7fb6fc3b8c65f2af4d3f4cabe024bce9",
    "0x13afe4210b6340723883be07618d3a4f019c2adc",
    "0x1c908c36c4470c3f66860e7c5152cc17e480ff54",
    "0x823a47d0a1f4c5781ead9f0d5b361557a45355bd",
    "0x52c21a9daa75e6eb91d449cad9649013903b7a73",
    "0x5fca7280f7b3f743bc3248261d05f7b57651ec3c",
    "0xe05e3146327d8fb775f50fd8d8f943dbe200ef4b"
  ]
}
```

### XDPoS_getSignersAtHash

The `getSignersAtHash` method retrieves the state snapshot at a given block.

Parameters:

- hash: string, required, block hash

Returns:

same as `XDPoS_getSigners`

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getSignersAtHash",
  "params": [
    "'"${hash}"'"
  ]
}' | jq
```

Response:

```shell
TODO:
```

### XDPoS_getSnapshot

The `getSnapshot` method retrieves the state snapshot at a given block.

Parameters:

- number: string, required, BlockNumber

Returns:

result: object PublicApiSnapshot:

- number: block number where the snapshot was created
- hash: block hash where the snapshot was created
- signers: array of authorized signers at this moment
- recents: array of recent signers for spam protections
- votes: list of votes cast in chronological order
- tally: current vote tally to avoid recalculating

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getSnapshot",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "number": 88958250,
    "hash": "0xacbdf5f5c5b51bec8ecb0af78ede89a9504e6196be47835f8eb5e7d8e8cfee21",
    "signers": {
      "0x00021c85d8cf9ef6c73e8f95740f8e874e4c9d1f": {},
      "0x00633d6bc3c30349a7fe55999c5b18940bcc9f43": {},
      "0x01df5f4986976e8cf1592cca7b792664510b6086": {},
      "0x0278c350152e15fa6ffc712a5a73d704ce73e2e1": {},
      "0x03d9e17ae3ff2c6712e44e25b09ac5ee91f6c9ff": {},
      "0x0407b8f788fbd763280212be6f7a672351009487": {},
      "0x047ffe1fc7f6d0b7168c4ccc312221089629f470": {},
      "0x051572edba4a8290bf45dd2d7ab80eb6cd56ccc6": {},
      "0x05f4c54ffc16b943b98713248c751ce98f924cbd": {},
      "0x065551f0dcac6f00cae11192d462db709be3758c": {},
      "0x06fb5e6a7abcc241729f591ac96ee41a9f10decf": {},
      "0x070eed0a2c77650dd41599920b5a0089dc60bfe6": {},
      "0x079d87904cd707143dd39c8be595bbe7a2181fb4": {},
      "0x08d53a08b1d46ecb45bc3a299af9abdde9035338": {},
      "0x09bae325e54edf4dc8b79a8bb29f90deb90e8fe7": {},
      "0x0afe5a6890cf357022dfdaac7e2d6e248755cdfd": {},
      "0x0b0afb41c14dd921a9a4511c226dc29254ba0d8f": {},
      "0x0b5ba833d689a2ec8a3f5275f2b033bdd2de337b": {},
      "0x0e26a54bf70c51460cb13f4823027678672e4a0a": {},
      "0x1390a610e22f9ee1c0fc92c341a2b0c6f3f9d08a": {},
      "0x13afe4210b6340723883be07618d3a4f019c2adc": {},
      "0x19a1e02fec5d52be689f59b22d9c0f1765bd7052": {},
      "0x19e945a26b15e4d6f96386a4f8596b6313311b7e": {},
      "0x1a1a8229e4a2fa06f73a564092976ba2dfe6aab8": {},
      "0x1a7c3d7481adc7e067f2f73f35f44143d7bf02d4": {},
      "0x1aff171caa8c1ea93bac4c27676f356d054412b4": {},
      "0x1bc85db77617515381cbbb28fdd2b1a10264de11": {},
      "0x1c908c36c4470c3f66860e7c5152cc17e480ff54": {},
      "0x1cad85be8d097ed7a9b74ef5fbc9c75f2e2aa3bb": {},
      "0x1d393817ab218ebade3e8ad591593ec3b132b1f6": {},
      "0x1e0103ba7665d15328b3886927d4f0a85f8b2299": {},
      "0x1eca7cd6df02a177e10bc9769a544b49f13aa875": {},
      "0x200f1c0b9411dd0f1fad9a76cc47cf8d56e0a19d": {},
      "0x21396ca79a433a20b5904c9a77aa9a4f110ef80e": {},
      "0x226ddc05e1adc6872285d95d4277a498fdce375b": {},
      "0x241ff0d3096e2e0b477780f9f551918a06827c05": {},
      "0x249f443d5c8400d7bc638a199d1f458b47a41a47": {},
      "0x24c0832d9df8189166d8e4e2cedc40ad69077b17": {},
      "0x251aca934009f4dda1d904e28ad7f20b3da4d888": {},
      "0x25c65b4b379ac37cf78357c4915f73677022eaff": {},
      "0x26179a252e9de28c29a37fbeb60df72be5e3e27d": {},
      "0x2759e5176992c834e7c0da762ae6178de6e7e711": {},
      "0x29a66147543707ef83d5d15e0a1d9f65cbe90e7c": {},
      "0x2a4559e84ca53e00b05d9d53086005432926a992": {},
      "0x2a591f3d64f3ce6b1d2afeead839ad76aab9feb2": {},
      "0x2d74d0125982bdc3a9f54a10216d82509379e821": {},
      "0x2d881b373edf297b495e8deedd9108e3d6870dee": {},
      "0x2e32b4db98ae68ea725212040f5f29abc61ce18f": {},
      "0x2f86476faa31c3f5a3d5b9376282e1b02b4dfa00": {},
      "0x3056a8bff9a17b12d08f1837d0b44cf1e2018fbe": {},
      "0x3197b8ed3bcd7eea2c695951d9d57e833457072b": {},
      "0x32e494f6f737bd419169a837c35e195a851bcb8d": {},
      "0x3396803b0cb3b5790911bb51b6c5193e9940793f": {},
      "0x346ada489c70c85ca665428389f1a971abbee960": {},
      "0x349563f9b0e3167f816cc39b272e3f888e23d3f3": {},
      "0x35898873c021f9f4fa01a76c7fa595159f5523e7": {},
      "0x360b9e5870fc91e375f902fc134f91739f4a99c7": {},
      "0x361faefbcd1c1b7f2d18723d178633202ccc5858": {},
      "0x391f3d21d9062ebbdd958e436b26d4b40928658c": {},
      "0x39ee018a85e4469a1b335efe014067e315bd6594": {},
      "0x39ffd4aa42b02f1b55ed822267da57219f11fc6a": {},
      "0x3b8a906602422fe9f316e6fa0c1728cfaa238766": {},
      "0x3bddd35ddf0945c0151a4ccd547d9e7a0b1fa4cd": {},
      "0x3cc8e30f8be48e57dbac01d8c7b05fb18ef1f7c8": {},
      "0x3dd28884ec3ea0904f0f6697c31b7d273a899480": {},
      "0x417b06236ecc5bf200a59152e92cd4923867409c": {},
      "0x418c9839222e7105d227281f2a69a6c9d19c297e": {},
      "0x4288545139bb238d1671b96e8f8016b0b3793ad4": {},
      "0x4398241671b3dd484fe3213a4fb7511f30e7d7c0": {},
      "0x43bb2e3bb47d0a12d45c87dcdbd6d16a47e1000a": {},
      "0x4466ca98bdcd08d54a9672d5ca76a69909c4b59e": {},
      "0x449bed0b27afe708a30a481329733bcbcba2bc21": {},
      "0x44d21022e104d605235c863cf86c96045fad9505": {},
      "0x450d714e65f4de007937c56fb1c4686efd4fb4ec": {},
      "0x48378844ab9486410ebf01d058a434e6a42e2ca5": {},
      "0x497c44f4a22099109ad7d194aad4fbe78dceb788": {},
      "0x49d3fa92eec838f644e8cacf2c93533d29b6c713": {},
      "0x4b7076c988da8a0ef87f1af137f7abc39557b746": {},
      "0x4bcdd234c20da32ec34a7ff85d6c1e6c832604d4": {},
      "0x4cfd9048cfe93e02e30a84c7538c726e980c1d78": {},
      "0x4e111142fbf2edacc4ab122feb54b031dc3d95d9": {},
      "0x4e52755f091c4fb826b994666a3e207fdf3a0cc5": {},
      "0x4eaff58c4a66839c6bbe2bca1dbd52a4d328aa5c": {},
      "0x51aa6a6b6a738d44e6a4e00d6e02faa991cf44be": {},
      "0x52c21a9daa75e6eb91d449cad9649013903b7a73": {},
      "0x5454edee66858dfcc14871cc8b26f57ef528bedc": {},
      "0x555d4cd195bcb7d4196f2005db46a4b71c36c5a4": {},
      "0x555ff7828a0b435432308a4721e4d9f110324025": {},
      "0x55b7d6ca73944e2ab43d226d7ed1ff3d72516aff": {},
      "0x5651290bdd3a952357066b324f9334b544100140": {},
      "0x58d877f1f8a61cf99192e3161429c8a78d7f1243": {},
      "0x58f25975086b8d23c028ebb651be064da0360367": {},
      "0x595b8170eaf2e53e47cc20db47ad063a1e6e2c0f": {},
      "0x5a39debded5481a5e5b73a1551fd31208b88da4e": {},
      "0x5f4ff7948061a605e9d64b4d5fde927b70bc3fa5": {},
      "0x5fca7280f7b3f743bc3248261d05f7b57651ec3c": {},
      "0x609cc32e7513135349fbe9ebd78ea8cdcfd9f819": {},
      "0x618e706c531e57f44a739cef129bf08cbd912d6a": {},
      "0x619f838ea2a12cdb508e759c3e0697e021d52ceb": {},
      "0x61ddae7aefa4d69c868379416504b1537d1fd558": {},
      "0x621996eb1d03a6e9614ddf150f6123c5602c8418": {},
      "0x63443ffdb5c139c3bbab97d3b06ee0674d75ab7f": {},
      "0x64e2e07bb17d978108824219741b9cc9e52b4443": {},
      "0x65c90c2d3e99d8366f90db7f4d25f23a0a569d49": {},
      "0x664c4a7b15d91b07c468162f535909114c038b91": {},
      "0x66bb8c2724ed2dc09ece6024045e90e65c29eaff": {},
      "0x67a24b5821724419bbc0710cbbb3122d84703318": {},
      "0x695ce237df312f9a6c6a5386b91b37ae249d07dd": {},
      "0x69a809f4d6e65792158c1bef35b316f6292eefb1": {},
      "0x6e5bbbf0d9d1851599a21676c99b1a8cc012753d": {},
      "0x6f356f6a7ff84c46f59fee691a2a161494871eec": {},
      "0x70916660766043c23443479658008eeedad8cb94": {},
      "0x72a0224245992792c75565bd3ab6caa6c8b95218": {},
      "0x72fb467ef6da285b6b0d8f7a25abd6049385d5fd": {},
      "0x74d3ac0efc4c22ea05150d7501c95c610b130c4b": {},
      "0x74f297c0197bab5ecf21add9ad9ced96ad42b61f": {},
      "0x7525f63e5f739ef952fbea50c1dc97ef5d5bd984": {},
      "0x75ea95b0e9a56417ce1434aac37b0afcba0c5221": {},
      "0x76df51309c6367678635a3ea76d1482666f55a05": {},
      "0x76e9285b4a804033a1d87b21c4e9d4e7028ce944": {},
      "0x783f0724f9dabd1bb52c7331bce7041804221d1b": {},
      "0x785076b971f2a90b029b680f90d4d609060c01d5": {},
      "0x78fb8980d122c902ad1abeac523a92e9c24e75c4": {},
      "0x7aa125338be075260e77c6a66a56c90a5dec4c58": {},
      "0x7af6aedf3382bb3de2dba61769d96444bf660494": {},
      "0x7b03e5e7912b0ce65f8be7549e23cb08392e369c": {},
      "0x7b0704722600979fd7da53afb1618aab3b9ddeb4": {},
      "0x7b6b545279bc9e262a9dc224ab724ad629c75b63": {},
      "0x7b94fcbd37f865892874e315c954062f936460fc": {},
      "0x7bda7752b3533b3b71346ff02f6314b8fe77ec43": {},
      "0x7c0f72d80d9d8b822e2fb217e56be732a12fc8c9": {},
      "0x7c1f132c335f04002178d0d29fe785ad329bdfb1": {},
      "0x7f82ddbd91c5fafe6d5fdf417c681eed211fefec": {},
      "0x7fca0e918c06ff0b4e51ecf76670e07714e44a9c": {},
      "0x82183d37848f38c807e274b15b86d16e9ae60fd3": {},
      "0x823a47d0a1f4c5781ead9f0d5b361557a45355bd": {},
      "0x82c85801a8b18c4b2701e552727cea8494bbe9ea": {},
      "0x83e12760c1e324b16fe68bc23ffc69b37d17ee23": {},
      "0x844a6c40b34e2335fc6887db2975d55183082c0d": {},
      "0x84620482d52b71cb18bf84cd34f0cc0dcfd940cd": {},
      "0x8b9654195b7c9a56ae20deb36a29834763800b03": {},
      "0x8e455bb521212bdcf8cb326e32dc1183cb3fd887": {},
      "0x8e9976680f84fab8b1dff91716762855c1a975e1": {},
      "0x8eeeeafaca49e507bc9bba121dac97ec75774ff2": {},
      "0x8f2fb5da850042b7da5097061f098493f8ec6dd8": {},
      "0x9043fafbe8833adafb449b559cef300a9632763f": {},
      "0x90c87c9ff588e9268a1c7c79a5986886e98c2f04": {},
      "0x9200e4fe8959f0268eb9daf06aba4756e595b99f": {},
      "0x92d32364313c376dcae272e113fd04cea194d2c8": {},
      "0x934a05866a7a44b212b55484f201e0ada955ef41": {},
      "0x935a34962ed8237d9700191320e3e0b66741133d": {},
      "0x944dd93e0f6d213b0dce27590bd5a9a4351fafd6": {},
      "0x95e58330cc798c8f079b8c91ba4d960d114cd785": {},
      "0x966f1f7e10f9bd531e31ab6bba0936944e00ebf9": {},
      "0x9682520376f26a733a75588c0bdede7645bdfc4a": {},
      "0x97c1117aec1b18a6b70a103eff30d2c55adf7bc0": {},
      "0x9a3787688fd210ec8f8d0224c6c50b8178d75bc0": {},
      "0x9bbf9e3af8a473124a17656f0e84d99aa8b51864": {},
      "0x9c54bc764c2e00717fbbf31b4dc962a72c47bdf6": {},
      "0x9d38187eccc49419c99af4b99dbb1494592b8a56": {},
      "0x9e057cebf2101a5e3e57906b834a2a82893b7739": {},
      "0x9fce52c5c451599235c17bc37e99f846d25ee6ec": {},
      "0xa43a99c9752166ec30c8369ce1e84de43be1e617": {},
      "0xa45457cc0153092a7e717e8e6d41221b827c28ae": {},
      "0xa4657c02208797985adedcbd048efc82291dbdb2": {},
      "0xa478633e8aba4188c25159e1103206eeca98bced": {},
      "0xa59318ba345b293f4ac5acb9f59a47365987b7da": {},
      "0xa65010026b83368ca05df6e8b467985d6de3eac5": {},
      "0xa6d916b338291ae14402d67af2b9d9e19bf72fd5": {},
      "0xa70b87c39a1237ba53953a1da04b23d4db28eace": {},
      "0xa72ce94a09db26dce57a4852409abc2fff07a962": {},
      "0xa828c2381bb91fc96cc269de73b99a4314ac4b5b": {},
      "0xa9c2b73ba170ddf37bc92fee1bafaf54e138ce87": {},
      "0xab5d74a461bb0f2231abcc5341ba7d89b24f0926": {},
      "0xac95f4af1e9efbc914895695e232fb8dbaa66275": {},
      "0xaf3db8ad8154820303adc6be0d9b10ef8e33024f": {},
      "0xb1693d224d4fd70ad6c9c9bdc44faa3ce21fb40d": {},
      "0xb1847452a58a2e2a4ee658a192f0d5d511c6f3f2": {},
      "0xb2610a596e2e34cf25adda2f05d867b3dc96b47d": {},
      "0xb2e8a55ea7fccdd8cdfb1af8e497a6f793ddf63e": {},
      "0xb38aba47a5563fa4aedc2a649ef819b9fb160b6a": {},
      "0xb41517b07efc01226332421c3fbf4e56807df5f1": {},
      "0xb55f4ca1a5741a91ff4ea4e15e9324eb973f4ccd": {},
      "0xb570b13fa9ca093347448dc600ed0d09333ac9fa": {},
      "0xb594c78a71763d543d10358b2dde579ff041319a": {},
      "0xb653aa2df1b710894bc1b94b3bd06799701d6d6b": {},
      "0xb8ad6e0a7bffb5ea0b24c095e16bb4a2c936203c": {},
      "0xb95d3439a7d3654e7871e72213caaad472ec8325": {},
      "0xb9a3a97f6a02a86483bf02cb33f8b33d2d117708": {},
      "0xba96294248d34b90f5b893ef82fd736cc1cc00ad": {},
      "0xbb0830e2d006d557541935362a621d002b0d6b00": {},
      "0xbbd2d417a8b6f1b1d7a267cd1d7402b443f35cfe": {},
      "0xbd46f94ca626a3ac72e60ca7efde379e42603b98": {},
      "0xc02aed857b01b4d60b378096221db3b60afadc37": {},
      "0xc428369353c7de2782370ee634be1516d836ddad": {},
      "0xc67c2dec79da735d6587d8db3c23271d557196ab": {},
      "0xc6b903af2816a2423a3300ed853e5e7d26ea0cf0": {},
      "0xc7bf1cd426e6f47d879fbaba789d7c5740ce9b4b": {},
      "0xc7d49d0a2cf198deebd6ce581af465944ec8b2bb": {},
      "0xc7f52a1c275d470ff92718dfe27be7800f7bb697": {},
      "0xcac04ac16876a65d3bf797679f34fc8272a1f6b4": {},
      "0xcb67c6c9ea422e33feef44aa5564bdf89997f57b": {},
      "0xcc8247cfaed1950dfd6dd4bc23a0ecab014c3b6c": {},
      "0xccab2f5267bc66b69f1688560c2e075216908a2c": {},
      "0xcd52996e7a6c75ddf85baab529af6c55f254ac15": {},
      "0xcfccdea1006a5cfa7d9484b5b293b46964c265c0": {},
      "0xd22fdac1459760f698618d927bbe22249e2b29b9": {},
      "0xd2fefdd118aa2bea5991dc079fd341d3e7a92caf": {},
      "0xd438128b57cf71ab2168a263f24708263775c54f": {},
      "0xd4ff534569c5bfd12d2280b8bd27aa8876f6b5fa": {},
      "0xd5384163e03481b8b1c13e8da611e05e2612586c": {},
      "0xd5bb37763625bc16f2ad0791d87e2726004241e4": {},
      "0xd64c15787bb52bbfe0d195df0d954725469a318e": {},
      "0xd6e9161fad50d2e697cda2b960fd26bf52a5d169": {},
      "0xd97d7434af4ad873585c71df55fce07996e27a82": {},
      "0xd9d474b83c514920f44c519f83e23e7b6927d552": {},
      "0xd9d4f525683b2b54508727ae44ac147e4f0db441": {},
      "0xdb2e141595d8edf6b1ce40b10d57ca5b83855ec5": {},
      "0xdccd99da1c942a36c5dcaf26d19b98c815dfcb73": {},
      "0xde1470c0f478ca7efd151be5a13c2addcfff328c": {},
      "0xdef4bd67f9e2627ca495b5f71794fab1bab40619": {},
      "0xe02ca8d829b3e0185d168c62aeef3960e6c0882e": {},
      "0xe05e3146327d8fb775f50fd8d8f943dbe200ef4b": {},
      "0xe230905c99aaa7b68402af8611b89ceda743191e": {},
      "0xe2bc9a03d5ae35e130c8bf99ea50c022de375db7": {},
      "0xe4710a854b24062ff37ab6636bd9a456e24c1635": {},
      "0xe494fe5a38b7212ab999152d1d45adab8d84a8a7": {},
      "0xe495a6242b78fe35211116839ad1882e7c4fc564": {},
      "0xe79ef2d92a8e4d758d78dfcc8e61b2fa377ee78d": {},
      "0xe7a5556c87fe764dcba9bcd31e1f4aa27b638117": {},
      "0xe865a5b2bf699a4b498de8a8c55da14bb0d94b21": {},
      "0xe8b4fd78437831cb9a182771bef485b619cef31f": {},
      "0xe8bc0e78114ad8dc4290eaed78bcb9e3f5b91e32": {},
      "0xe8e8194c0a8e5a32306c7a86052a4e8ca4b8729d": {},
      "0xeb0f1565b5d3bbf4489deb67972b814ee84c6cb7": {},
      "0xeb76fb646c5ab9db8543719e74adadf101cd0616": {},
      "0xebf8e4904fe2bd561475e9e0e65a3336c56f83d3": {},
      "0xec4a5fbd2e46e97be98c212e668ef1d08d695440": {},
      "0xec5273989dd0e18d658f037cae2eff6803369c5c": {},
      "0xef7ce03b7fdea2d0bf2ad6e0a1e9b0f2ed547447": {},
      "0xf04f32c46f1a16663bc7f6409b35b33f3daa9a03": {},
      "0xf0dee298ac62b7e351148b4fbde1d45d7d9b1560": {},
      "0xf11e012f56ac7341be939cb668e63a37e79db84e": {},
      "0xf2a1d535f4bda9e810c067975d08dd8724f9066a": {},
      "0xf38f7402de55c5c7cf7a5a29244148b817d7bca3": {},
      "0xf48aa29de23c45c54a96c2a7b8ce5218abf46874": {},
      "0xf7ff86c0d2c2795dbd2731ca122a7d13a802b7b8": {},
      "0xfa9a1f2ec3d92428fc6105a4e05832ae74170ac4": {},
      "0xfab275bfc83c3bc8aa17c16ed5a0b4e60caca299": {},
      "0xfc2f29a87a63e07a341d0bb9462e9a14af013a6c": {},
      "0xfc5fea016be9ac2d23dfe18e16d3be79547e5466": {},
      "0xfcbddd58be4616042543637a84ad18cb065adba4": {},
      "0xfce02e190f722e790be06b3a6ec43e2f606dba65": {},
      "0xfe5b2f3e7fb6fc3b8c65f2af4d3f4cabe024bce9": {}
    },
    "recents": null,
    "votes": null,
    "tally": null
  }
}
```

### XDPoS_getSnapshotAtHash

The `getSnapshotAtHash` method retrieves the state snapshot at a given block.

Parameters:

- hash: string, required, block hash

Returns:

same as `XDPoS_getSnapshot`

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getSnapshotAtHash",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell
TODO:
```

### XDPoS_getV2BlockByHash

Parameters:

- hash: string, required, block hash

Returns:

same as `XDPoS_getV2BlockByNumber`

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getV2BlockByHash",
  "params": [
    "'"${hash}"'"
  ]
}' | jq
```

Response:

```shell
TODO:
```

### XDPoS_getV2BlockByNumber

Parameters:

- number: string, required, BlockNumber

Returns:

result: object V2BlockInfo:

- Hash: hash
- Round: uint64
- Number: big.Int
- ParentHash: hash
- Committed: bool
- Miner: common.Hash
- Timestamp: big.Int
- EncodedRLP: string
- Error: string

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_getV2BlockByNumber",
  "params": [
    "latest"
  ]
}' | jq
```

Response:

```shell
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "Hash": "0xbe0ef7d77e3193361d77b427de62b4562f595a5a2d05fc4afa4d1d2413127870",
    "Round": 8705315,
    "Number": 88959039,
    "ParentHash": "0x18bb9571872dc9f0c40b8b00f9b8f5774fa09bf143f483e3d5814ac4ee727697",
    "Committed": false,
    "Miner": "0x000000000000000000000000eb0f1565b5d3bbf4489deb67972b814ee84c6cb7",
    "Timestamp": 1747905328,
    "EncodedRLP": "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",
    "Error": ""
  }
}
```

### XDPoS_networkInformation

Parameters:

None

Returns:

result: object NetworkInformation:

- NetworkId: big.Int
- XDCValidatorAddress: address
- RelayerRegistrationAddress: address
- XDCXListingAddress: address
- XDCZAddress: address
- LendingAddress: address
- ConsensusConfigs: object of XDPoSConfig

Example:

```shell
curl -s -X POST -H "Content-Type: application/json" ${RPC} -d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "XDPoS_networkInformation"
}' | jq
```

Response:

```shell
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "NetworkId": 50,
    "XDCValidatorAddress": "0x0000000000000000000000000000000000000088",
    "RelayerRegistrationAddress": "0x16c63b79f9c8784168103c0b74e6a59ec2de4a02",
    "XDCXListingAddress": "0xde34dd0f536170993e8cff639ddffcf1a85d3e53",
    "XDCZAddress": "0x8c0faeb5c6bed2129b8674f262fd45c4e9468bee",
    "LendingAddress": "0x7d761afd7ff65a79e4173897594a194e3c506e57",
    "ConsensusConfigs": {
      "period": 2,
      "epoch": 900,
      "reward": 5000,
      "rewardCheckpoint": 900,
      "gap": 450,
      "foudationWalletAddr": "0x92a289fe95a85c53b8d0d113cbaef0c1ec98ac65",
      "SkipV1Validation": false,
      "v2": {
        "switchBlock": 80370000,
        "config": {
          "maxMasternodes": 108,
          "switchRound": 3200000,
          "minePeriod": 2,
          "timeoutSyncThreshold": 3,
          "timeoutPeriod": 10,
          "certificateThreshold": 0.667
        },
        "allConfigs": {
          "0": {
            "maxMasternodes": 108,
            "switchRound": 0,
            "minePeriod": 2,
            "timeoutSyncThreshold": 3,
            "timeoutPeriod": 30,
            "certificateThreshold": 0.667
          },
          "2000": {
            "maxMasternodes": 108,
            "switchRound": 2000,
            "minePeriod": 2,
            "timeoutSyncThreshold": 2,
            "timeoutPeriod": 600,
            "certificateThreshold": 0.667
          },
          "220000": {
            "maxMasternodes": 108,
            "switchRound": 220000,
            "minePeriod": 2,
            "timeoutSyncThreshold": 2,
            "timeoutPeriod": 30,
            "certificateThreshold": 0.667
          },
          "3200000": {
            "maxMasternodes": 108,
            "switchRound": 3200000,
            "minePeriod": 2,
            "timeoutSyncThreshold": 3,
            "timeoutPeriod": 10,
            "certificateThreshold": 0.667
          },
          "460000": {
            "maxMasternodes": 108,
            "switchRound": 460000,
            "minePeriod": 2,
            "timeoutSyncThreshold": 2,
            "timeoutPeriod": 20,
            "certificateThreshold": 0.667
          },
          "8000": {
            "maxMasternodes": 108,
            "switchRound": 8000,
            "minePeriod": 2,
            "timeoutSyncThreshold": 2,
            "timeoutPeriod": 60,
            "certificateThreshold": 0.667
          }
        },
        "SkipV2Validation": false
      }
    }
  }
}
```
