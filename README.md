# GSA and external-CI integration lab

这是一个长期保留的公开实验仓库，用来验证 GitHub Repository Security
Advisory（GSA）的完整操作、权限边界，以及 GitHub Actions/Jenkins 等 CI 与
temporary private fork 的集成能力。

仓库中的漏洞、修复和 advisory 全部是合成测试数据，不对应真实安全问题，
不得申请 CVE，也不会发布成真实安全公告。

完整中文报告：[`docs/GSA-操作与外部-CI-验证报告.zh-CN.md`](docs/GSA-操作与外部-CI-验证报告.zh-CN.md)。
报告包含逐点击流程、CLI/API 命令、权限矩阵、生命周期实验、截图，以及 TiDB
Jenkins/其它 CI 与 temporary private fork 的实测结论。仓库还包含一个中文自定义
private vulnerability reporting 表单，用于验证必填字段和 API 行为。
