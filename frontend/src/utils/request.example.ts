/**
 * Token 过期处理测试示例
 * 
 * 这个文件展示了如何手动测试 token 过期处理逻辑
 */

// 测试步骤：
// 1. 打开浏览器开发者工具
// 2. 在 Console 中执行以下代码来模拟过期 token
localStorage.setItem('token', 'expired_token_example')

// 3. 刷新页面或发起任意 API 请求
// 4. 观察是否自动跳转到登录页面
// 5. 检查 localStorage 是否被清空

// 或者使用以下方式测试：
// 1. 正常登录获取有效 token
// 2. 等待 token 自然过期（根据 JWT 配置的时间）
// 3. 发起 API 请求
// 4. 验证是否自动跳转到登录页面

// 预期行为：
// - 当后端返回 { status: 400, msg: "token is expired" } 时
// - 前端应该清除 localStorage
// - 并重定向到 /login 页面
