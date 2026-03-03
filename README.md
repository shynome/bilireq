## bilireq

**自用, 接口尚未稳定, 随时可能变动**

使用浏览器代替用户发起请求, 实现自动化, 无需实现cookie刷新机制

## UserScript

```js
// ==UserScript==
// @name        well-jsnet
// @namespace   well-jsnet.remoon.net
// @match       https://message.bilibili.com/
// @grant       none
// @version     1.0
// @author      -
// @run-at      document-start
// @description 12/05/2025, 10:07:53
// ==/UserScript==

void (async function main() {
  console.log("加载程序")
  // const SaltLink = await import("http://127.0.0.1:4173/index.js")
  const SaltLink = await import("https://unpkg.com/well-net/index.js")
  console.log("连接开始")
  const net = await SaltLink.connect({
    Key: "CH7G4Uu+0hDnIVzcc0aN+iPwgKG/uGZbL9gJvZnSg3k=",
    Peer: "ws://127.0.0.1:7799/api/whip#xMjphMUyLIGExyJluSslD9tjaIcF9QS6ADyI8DOTzyg=",
    // LogLevel: "debug",
  })
  const srv = await net.listen("0.0.0.0:80", {
    async fetch(req) {
      const link = new URL(req.url)
      if (link.pathname === "/live-cookie") {
        const ids = await Promise.all(
          ["DedeUserID", "buvid3"].map((k) =>
            cookieStore.get(k).then((v) => v.value),
          ),
        )
        return new Response(JSON.stringify(ids))
      }
      const jct = await cookieStore.get("bili_jct")
      if (link.pathname === "/csrf2") {
        const uid = await cookieStore.get("DedeUserID").then((v) => v.value)
        const devId = localStorage.getItem("im_deviceid_" + uid)
        return new Response(JSON.stringify([jct.value, uid, devId]))
      }
      return new Response(jct.value)
    },
  })
  await net.http_proxy("0.0.0.0:1080", {})
  console.log("代理启动成功")
  console.log("连接成功")
  setTimeout(
    () => {
      location.reload()
    },
    24 * 60 * 60 * 1e3,
  )
})()
```

# Todo

- [x] 会话列表接口
- [x] 会话消息接口
- [x] 发送消息接口
