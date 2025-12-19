使用 `rsrc` 将清单文件嵌入 Go 程序，可以让编译后的程序在 Windows 上自动请求管理员权限。以下是详细步骤。

### 第一步：安装 rsrc

在终端执行以下命令安装 `rsrc` 工具：
```bash
go install github.com/akavel/rsrc@latest
```
安装成功后，`rsrc.exe` 会出现在你的 `$GOPATH/bin` 目录下（通常是 `C:\Users\<用户名>\go\bin`），请确保该目录已添加到系统的 PATH 环境变量中。

### 第二步：准备清单文件

创建一个名为 `app.manifest` 的文本文件，内容如下。这个文件告诉 Windows：“此程序需要以管理员身份运行”。
```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
    <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
        <security>
            <requestedPrivileges>
                <!-- 关键设置：requireAdministrator 表示需要管理员权限 -->
                <requestedExecutionLevel level="requireAdministrator" uiAccess="false"/>
            </requestedPrivileges>
        </security>
    </trustInfo>
</assembly>
```
将文件保存在你的 Go 项目目录下。

### 第三步：生成 .syso 资源文件

打开命令行，**进入你的 Go 项目根目录**，然后执行 `rsrc` 命令：
```bash
rsrc -manifest app.manifest -o rsrc.syso
```
此命令会读取 `app.manifest`，并生成一个名为 `rsrc.syso` 的 Windows 资源文件。

**关键点**：
*   `-manifest`：指定清单文件路径。
*   `-o`：指定输出的 `.syso` 文件名。**强烈建议使用默认的 `rsrc.syso`**，因为这是 Go 工具链在 Windows 上查找资源文件时的默认名称，兼容性最好。
*   生成的 `rsrc.syso` 文件**必须放在与你的 `main` 包（即包含 `func main()` 的文件）相同的目录下**。

### 第四步：正常编译程序

现在，像往常一样编译你的 Go 程序即可：
```bash
go build -o your-program.exe
```
Go 编译器会自动发现并嵌入 `rsrc.syso` 文件中的资源。编译完成后，`your-program.exe` 运行时就会自动弹出 UAC 提示，请求管理员权限。

### 验证与排查

*   **验证清单是否嵌入**：你可以右键点击生成的 `.exe` 文件，选择 **属性 → 兼容性**，如果看到“以管理员身份运行此程序”的选项被勾选且无法取消，就说明清单已成功嵌入。
*   **清理旧资源**：如果之前生成过其他 `.syso` 文件（如 `app.syso`），请先删除它们，只保留最新的 `rsrc.syso`，避免多个资源文件造成冲突。
*   **文件位置**：这是最常见的错误。请务必确认 `rsrc.syso` 文件位于正确的目录（项目根目录，且与 `main.go` 同级）。

### 高级用法：嵌入图标等资源

`rsrc` 的功能不限于清单文件，它还可以将图标（`.ico`）等资源嵌入程序。命令格式如下：
```bash
rsrc -manifest app.manifest -ico your-icon.ico -o rsrc.syso
```
这样，程序的属性对话框中就会显示你自定义的图标。

### 总结流程与要点

整个流程可以总结为：**安装工具 → 编写清单 → 生成资源 → 编译程序**。

对于你的 Windows 服务管理工具，使用清单文件提权是最标准、最可靠的方式。如果后续你希望程序在非管理员权限下也能启动（比如先检查更新，再动态提权），可以再考虑实现运行时提权方案。

如果你在操作中遇到了错误（例如 `rsrc` 命令未找到，或者编译后程序没有请求提权），可以告诉我具体的错误信息，我能帮你进一步排查。