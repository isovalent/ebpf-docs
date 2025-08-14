# VMLinux blob update process

The current VMLinux blob is created from the v6.17-rc1 tag of the Linux kernel. The following steps are used to update the VMLinux blob:

1. Clone cilium/ci-kernels
2. Patch `config` with the following values (To compile in all Kfuncs defined in the kernel (so far)):
   ```
   CONFIG_HID=y
   CONFIG_HID_BPF=y
   CONFIG_TCP_CONG_BBR=y
   CONFIG_TCP_CONG_DCTCP=y
   CONFIG_XFRM=y
   CONFIG_XFRM_INTERFACE=y
   CONFIG_FS_VERITY=y
   CONFIG_MODULE_SIG=y
   CONFIG_MODULE_SIG_FORMAT=y
   CONFIG_SYSTEM_DATA_VERIFICATION=y
   CONFIG_CRYPTO=y
   CONFIG_NF_TABLES=y
   CONFIG_NF_FLOW_TABLE=y
   CONFIG_MMU=y
   CONFIG_64BIT=y
   CONFIG_CGROUP_SCHED=y
   CONFIG_SCHED_CLASS_EXT=y
   CONFIG_DMA_SHARED_BUFFER=y
   CONFIG_NET_SCH_BPF=y
   ```
3. Run `./buildx.sh {latest tag} amd64 vmlinux --tag foo:vmlinux`
4. Run `echo "FROM foo:vmlinux" | "$docker" buildx build --quiet --output="$tmp" - &> /dev/null`
5. Run `"/lib/modules/$(uname -r)/build/scripts/extract-vmlinux" "$tmp/boot/vmlinuz" > "$tmp/vmlinux.elf"`
6. Run `objcopy --dump-section .BTF=/dev/stdout "$tmp/vmlinux.elf" /dev/null > vmlinux`
