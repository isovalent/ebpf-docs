# Resource limiting

## Rlimit

rlimit or "resource limit" is a system to track and limit the amount of certain resources you are allowed to use. One of the things it limits is the amount of "locked memory" https://man7.org/linux/man-pages/man2/getrlimit.2.html

Until kernel version v5.11 this mechanism was used for BPF resources, so you commonly would have to increase or disable this rlimit which requires an additional capability CAP_SYS_RESOURCE.

## cGroup memory limit

After v5.11, cgroup-based memory accounting is used for BPF resources which eliminates the need for the rlimit raising/disabling

