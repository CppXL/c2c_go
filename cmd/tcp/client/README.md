#### Client命令

* ListDir 从受感染系统返回目录列表
* ListProcesses 此命令返回在受感染系统上运行的进程列表
* 此命令会导致恶意软件针对受感染系统上运行的目标进程执行 taskkill
  命令
* DownloadFile - 此命令会导致恶意软件下载文件并将其保存到命令参数指定的目标
  位置
* RunCMD - 此命令使用 Go os/exec 包在受感染的系统上执行系统命令
* DLRUN - 此命令会导致恶意软件下载文件，将其保存到 %TEMP% 并执行该文件